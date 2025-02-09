#include <vector>

using namespace std;

template <typename T>
class Node
{
public:

    vector<T> keys;
    vector<unique_ptr<Node<T>>> children;
    bool leaf;
    int size = 0;

    Node(int keys_count, bool is_leaf = false)
        : keys(keys_count), children(keys_count + 1), leaf(is_leaf) {}
    void showKeys();
    void showChildren();
    void printNode();
    void printKeys(int depth = 0);
};