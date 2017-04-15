
import sys
import os

sys.path.insert(0, os.path.join(os.path.dirname(__file__), '..', 'util'))
from bst import BST, Node
from file_stream import binary_search_tree

# create an object of the bineary search tree
b = BST()

# specify the name of the binary search tree you want to use 
file_name = 'bst_0'
root = binary_search_tree(file_name)
print root


b.preorder(root)

