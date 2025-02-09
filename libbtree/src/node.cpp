#include <memory>
#include <vector>
#include <iostream>
#include <fstream>
#include <string>
#include <cstdio>

#include "node.h"

using namespace std;

template <typename T>
void Node<T>::showKeys()
{
    cout << "node keys ";
    for (int i = 0; i < size; i++)
        cout << keys[i] << " ";
    cout << "\n";
}

template <typename T>
void Node<T>::showChildren()
{
    cout << "node child keys ";
    for (int i = 0; i <= size; i++)
        for (int j = 0; j < children[i].get()->size; j++)
            cout << children[i].get()->keys[j] << " ";
    cout << "\n";
}

template <typename T>
void Node<T>::printNode()
{
    printKeys(0);
    ifstream inputFile("out.txt");
    string line = "";

    if (!inputFile)
    {
        cerr << "File could not be opened!" << endl;
        return;
    }

    while (getline(inputFile, line))
        cout << line << "\n";

    remove("out.txt");
}

template <typename T>
void Node<T>::printKeys(int depth)
{
    if (depth == 0)
    {
        ofstream outputFile("out.txt");
        if (outputFile)
        {
            outputFile << "";
            outputFile.close();
        }
    }

    ifstream inputFile("out.txt");
    vector<string> lines;
    string line = "";

    if (!inputFile)
    {
        cerr << "File could not be opened!" << endl;
        return;
    }

    while (getline(inputFile, line))
        lines.push_back(line);

    inputFile.close();

    string newLine = "";
    for (int i = 0; i < size; i++)
        newLine = newLine + to_string(this->keys[i]) + " ";

    string updateLine;
    if (depth < lines.size())
        updateLine = lines[depth];
    else
    {
        updateLine = "";
        lines.push_back("");
    }

    if (!updateLine.empty() && updateLine.back() == '\n')
        updateLine.pop_back();

    updateLine = updateLine + newLine;
    lines[depth] = updateLine;

    ofstream outputFile("out.txt");
    if (!outputFile)
    {
        cerr << "Error opening file for writing!" << endl;
        return;
    }

    for (const auto &l : lines)
        outputFile << l << endl;
    outputFile.close();
    if (!leaf)
        for (int i = 0; i <= size; i++)
            children[i].get()->printKeys(depth + 1);
}

template class Node<long>;
template class Node<int>;