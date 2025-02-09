CXX = g++-14

all: testbtree

testbtree: tests/testbtree.cpp src/btree.cpp src/node.cpp
	$(CXX)  -I. src/btree.cpp src/node.cpp tests/testbtree.cpp -o testbtree
	./testbtree 

clean:
	@rm testbtree