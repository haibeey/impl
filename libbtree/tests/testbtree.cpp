#include <chrono>
#include <cassert>
#include <iostream>
#include "src/btree.h"

using namespace std;

int main()
{
    Btree<long> tree = Btree<long>(3);
    for (int i = 0; i < 15; i++)
        tree.Insert(i);

    Node<long> *r = tree.Search(11);
    assert(r != nullptr);

    tree.Delete(6);
    r = tree.Search(6);
    assert(r == nullptr);

    tree.Delete(4);
    r = tree.Search(4);
    assert(r == nullptr);

    tree.printTree();

    return 0;
}