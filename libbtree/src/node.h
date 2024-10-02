#include <memory>
#include <vector>
#include <iostream>
#include <fstream>
#include <string>
#include <cstdio>

using namespace std;

template <typename T>
struct Node
{
    vector<T> keys;
    vector<unique_ptr<Node<T>>> children;
    bool leaf;
    int size = 0;
    Node(int keys_count, bool is_leaf = false)
        : keys(keys_count), children(keys_count + 1), leaf(is_leaf) {}

    void showKeys()
    {
        cout << "node keys ";
        for (int i = 0; i < size; i++)
            cout << keys[i] << " ";
        cout << "\n";
    }

    void showChildren()
    {
        cout << "node child keys ";
        for (int i = 0; i <= size; i++)
            for (int j = 0; j < children[i].get()->size; j++)
                cout << children[i].get()->keys[j] << " ";
        cout << "\n";
    }

    void printNode() const
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

    void printKeys(int depth = 0) const
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
};