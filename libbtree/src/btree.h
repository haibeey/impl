#include <memory>
#include "node.h"

using namespace std;

template <typename T>
class Btree
{
private:
    int keys_count;
    int t;
    unique_ptr<Node<T>> root;

    void insertNonFull(Node<T> *node, T key);
    void splitChild(Node<T> *node, int index);
    void deleteHelper(Node<T> *node, T key);
    void deleteInternalNode(Node<T> *node, T key, int index);
    void takeFromSibling(Node<T> *node, int A, int B);
    void merge(Node<T> *node, int A, int B);
    T deletePredecessor(Node<T> *node);
    T deleteSuccessor(Node<T> *node);
    Node<T> *searchHelper(Node<T> *node, T k);

public:
    Btree(int keys_count)
    {
        this->t = keys_count;
        this->keys_count = (2 * keys_count) - 1;
        this->root = make_unique<Node<T>>(this->keys_count, true);
    }

    Node<T> *Search(T k);
    void Insert(T key);
    void Delete(T key);

    void printTree()
    {
        root->printNode();
    }

    int getKeyCount() const
    {
        return keys_count;
    }

    Node<T> *getRoot()
    {
        return root.get();
    }
};

template <typename T>
Node<T> *Btree<T>::Search(T k)
{
    return searchHelper(this->root.get(), k);
}

template <typename T>
Node<T> *Btree<T>::searchHelper(Node<T> *node, T key)
{
    int i = 0;
    while (i < node->size && key > node->keys[i])
        i++;
    if (i <= node->size && key == node->keys[i])
        return node;
    else if (node->leaf)
        return nullptr;
    else
        return searchHelper(node->children[i].get(), key);
}

template <typename T>
void Btree<T>::splitChild(Node<T> *x, int index)
{
    Node<T> *y = x->children[index].get();
    unique_ptr<Node<T>> z = make_unique<Node<T>>(this->getKeyCount());
    int t = this->t;
    z.get()->leaf = y->leaf;
    z.get()->size = t - 1;

    for (int i = 0; i < t - 1; i++)
        z->keys[i] = y->keys[i + t];

    if (!y->leaf)
    {
        for (int i = 0; i < t; i++)
            z->children[i] = move(y->children[i + t]);
    }

    y->size = t - 1;

    for (int i = x->size; i > index; i--)
        x->children[i] = move(x->children[i - 1]);

    x->children[index + 1] = move(z);

    for (int i = x->size; i > index; i--)
        x->keys[i] = x->keys[i - 1];

    x->keys[index] = y->keys[t - 1];
    ++x->size;
}

template <typename T>
void Btree<T>::Insert(T key)
{
    Node<T> *r = this->root.get();
    if (r->size == getKeyCount())
    {
        unique_ptr<Node<T>> new_node = make_unique<Node<T>>(this->getKeyCount());
        new_node.get()->children[0] = move(this->root);
        this->root = move(new_node);
        splitChild(this->root.get(), 0);
        insertNonFull(this->root.get(), key);
    }
    else
        insertNonFull(r, key);
}

template <typename T>
void Btree<T>::insertNonFull(Node<T> *node, T k)
{
    int index = node->size;
    if (node->leaf)
    {
        while (index > 0 && k < node->keys[index - 1])
        {
            node->keys[index] = node->keys[index - 1];
            --index;
        }
        node->keys[index] = k;
        ++node->size;
    }
    else
    {
        while (index > 0 && k < node->keys[index - 1])
            --index;

        if (node->children[index]->size == getKeyCount())
        {
            splitChild(node, index);
            if (k > node->keys[index])
                ++index;
        }

        insertNonFull(node->children[index].get(), k);
    }
}

template <typename T>
void Btree<T>::Delete(T key)
{
    deleteHelper(getRoot(), key);
}

template <typename T>
void Btree<T>::deleteHelper(Node<T> *node, T key)
{
    int index = 0;
    while (index < node->size && key > node->keys[index])
        index++;

    if (node->leaf)
    {
        if (node->keys[index] == key)
        {
            node->keys.erase(node->keys.begin() + index);
            --node->size;
        }

        return;
    }

    if (node->keys[index] == key)
        deleteInternalNode(node, key, index);
    else if (node->children[index]->size >= t)
        deleteHelper(node->children[index].get(), key);
    else
    {
        if (index != 0 && index + 2 <= node->size)
            if (node->children[index - 1]->size >= t)
                takeFromSibling(node, index, index - 1);
            else if (node->children[index + 1]->size >= t)
                takeFromSibling(node, index, index + 1);
            else
                merge(node, index, index + 1);
        else if (index == 0)
            if (node->children[1]->size >= t)
                takeFromSibling(node, index, index + 1);
            else
                merge(node, index, index + 1);
        else if (index + 1 == node->size + 1)
        {
            if (node->children[index - 1]->size >= t)
                takeFromSibling(node, index, index - 1);
            else
                merge(node, index, index - 1);
        }
        deleteHelper(node->children[index].get(), key);
    }
}

template <typename T>
void Btree<T>::deleteInternalNode(Node<T> *node, T key, int index)
{
    if (node->keys[index] == key)
        if (node->leaf)
        {
            node->keys.erase(node->keys.begin() + index);
            --node->size;
            return;
        }

    if (node->children[index]->size >= t)
        node->keys[index] = deletePredecessor(node->children[index].get());
    else if (node->children[index + 1]->size >= t)
        node->keys[index] = deleteSuccessor(node->children[index + 1].get());
    else
    {
        merge(node, index, index + 1);
        deleteInternalNode(node->children[index].get(), key, t - 1);
    }
}

template <typename T>
T Btree<T>::deletePredecessor(Node<T> *node)
{
    if (node->leaf)
    {
        --node->size;
        return node->keys[node->size + 1];
    }

    if (node->children[node->size - 1]->size >= t)
        takeFromSibling(node, node->size, node->size - 1);
    else
        merge(node, node->size - 1, node->size);

    return deletePredecessor(node->children[node->size - 1].get());
}

template <typename T>
T Btree<T>::deleteSuccessor(Node<T> *node)
{
    if (node->leaf)
    {
        T return_value = node->keys[0];
        node->keys.erase(node->keys.begin());
        --node->size;
        return return_value;
    }

    if (node->children[1]->size >= t)
        takeFromSibling(node, 0, 1);
    else
        merge(node, 0, 1);

    return deleteSuccessor(node->children[0].get());
}

template <typename T>
void Btree<T>::merge(Node<T> *node, int A, int B)
{

    Node<T> *nodeA = node->children[A].get();
    Node<T> *nodeB = node->children[B].get();

    if (B > A)
    {
        // Add all node b keys and children to A
        nodeA->keys[nodeA->size] = node->keys[A];
        ++nodeA->size;
        for (int i = 0; i <= nodeB->size; i++)
        {
            if (i < nodeB->size)
                nodeA->keys[nodeA->size + i] = nodeB->keys[i];

            if (!nodeB->leaf)
                nodeA->children[nodeA->size + i] = move(nodeB->children[i]);
        }
        nodeA->size += nodeB->size;
        nodeB->size = 0;
        node->keys.erase(node->keys.begin() + A);
        node->children.erase(node->children.begin() + B);
        --node->size;

        if (node == getRoot() && node->size <= 0)
            this->root = move(node->children[A]);
    }
    else
    {
        // Add all node A keys and children to B
        nodeB->keys[nodeB->size] = node->keys[A];
        ++nodeB->size;
        for (int i = 0; i <= nodeA->size; i++)
        {
            if (i < nodeB->size)
                nodeB->keys[nodeB->size + i] = nodeA->keys[i];

            if (!nodeA->leaf)
                nodeB->children[nodeB->size + i] = move(nodeA->children[i]);
        }

        nodeB->size += nodeA->size;
        nodeA->size = 0;
        node->keys.erase(node->keys.begin() + B);
        node->children.erase(node->children.begin() + A);
        --node->size;

        if (node == getRoot() && node->size <= 0)
            this->root = move(node->children[B]);
    }
}

template <typename T>
void Btree<T>::takeFromSibling(Node<T> *node, int A, int B)
{
    Node<T> *nodeA = node->children[A].get();
    Node<T> *nodeB = node->children[B].get();
    if (A < B)
    {
        // Taking right child key from node B
        nodeA->keys[nodeA->size] = node->keys[A];
        if (!nodeB->leaf)
            nodeA->children[nodeA->size] = move(nodeB->children[0]);
        node->keys[A] = nodeB->keys[0];
        nodeB->keys.erase(nodeB->keys.begin());
        if (!nodeB->leaf)
            nodeB->children.erase(nodeB->children.begin());
        --nodeB->size;
        ++nodeA->size;
    }
    else
    {
        // Taking left child key from node B
        nodeA->keys.insert(nodeA->keys.begin(), node->keys[A - 1]);
        node->keys[A - 1] = nodeB->keys[nodeB->size - 1];
        if (!nodeB->leaf)
        {
            nodeA->children.insert(nodeA->children.begin(), nullptr);
            nodeA->children[0] = move(nodeB->children[nodeB->size - 1]);
        }

        ++nodeA->size;
        --nodeB->size;
    }
}
