CXX = g++-14

all: testbtree

testbtree: tests/testbtree.cpp
	$(CXX)  -I. tests/testbtree.cpp -o testbtree
	./testbtree 

clean:
	@rm testbtree