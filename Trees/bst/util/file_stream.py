import sys
import os
import pickle

def binary_search_tree(file_name):
 file_path = os.path.join(os.path.dirname(__file__), '..', 'data', file_name)
 fh = open(file_path, 'r')
 root = pickle.load(fh)
 fh.close()
 return root
 
