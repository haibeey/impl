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
public:
    void insertNonFull(Node<T> *node, T key);
    void splitChild(Node<T> *node, int index);
    void deleteHelper(Node<T> *node, T key);
    void deleteInternalNode(Node<T> *node, T key, int index);
    void takeFromSibling(Node<T> *node, int A, int B);
    void merge(Node<T> *node, int A, int B);
    T deletePredecessor(Node<T> *node);
    T deleteSuccessor(Node<T> *node);
    Node<T> *searchHelper(Node<T> *node, T k);


    Btree(int keys_count);
    Node<T> *Search(T k);
    void Insert(T key);
    void Delete(T key);
    void printTree();
    int getKeyCount() const;
    Node<T> *getRoot();
};
