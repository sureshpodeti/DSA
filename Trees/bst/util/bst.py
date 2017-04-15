import sys
import os
import pickle

class Node(object):
 def __init__(self, x):
  self.left = None
  self.data = x
  self.right = None

class BST(object):
 def insert(self, x, rt=None):
  if not rt:
   rt = Node(x)
  elif rt.data > x:
   rt.left = self.insert(x, rt.left) 
  elif rt.data < x:
   rt.right  = self.insert(x, rt.right)
  return rt

 
 def preorder(self, rt):
  if not rt:
   return
  print rt.data
  self.preorder(rt.left)
  self.preorder(rt.right) 
    
 def inorder(self, rt):
  if not rt:
   return 
  self.inorder(rt.left)
  print rt.data
  self.inorder(rt.right)

 def postorder(self, rt):
  if not rt:
   return
  self.postorder(rt.left)
  self.postorder(rt.right)
  print rt.data

 def search(self, rt, x):
  if not rt:
   return False
  if rt.data > x:
   return self.search(rt.left, x)
  elif rt.data < x:
   return self.search(rt.right, x)
  if rt.data == x:
   return True


 def print_leaf_nodes(self, rt):
  if not rt:
   return
  if not rt.left and not rt.right:
   print rt.data
  self.print_leaf_nodes(rt.left)
  self.print_leaf_nodes(rt.right)

 def nodes_count(self, rt):
  if not rt:
    return 0
  return 1+ self.nodes_count(rt.left) + self.nodes_count(rt.right) 
  
 def node_children(self, rt, x):
  if not rt:
   return
  if rt.data == x:
    if rt.left:
     print rt.left.data,
    if rt.right:
     print rt.right.data
  if rt.data > x: 
   self.node_children(rt.left, x)
  elif rt.data < x:
   self.node_children(rt.right, x)

 def height_of_tree(self, rt):
  '''
   height of the tree = number of edges on the longest path between that node and a leaf. 
   The height of a tree is the height of its root node.
   The depth of a node is the number of edges from the tree's root node to the node.
   let x be the node of binary search tree
   l(x) be the height of left sub tree of node x 
   r(x) be the height of right sub tree of node x 
   h(x) = 1 + max(l(x), r(x)) 
  '''
  if not rt.left and not rt.right:
   return 0
  if not rt.left:
   return 1 + self.height_of_tree(rt.right)
  elif not rt.right:
   return 1 + self.height_of_tree(rt.left)
  else:
   return 1 + max(self.height_of_tree(rt.left), self.height_of_tree(rt.right))
  
 def search_for_node(self, rt, x):
  if not rt:
   return
  if rt.data > x:
   return self.search_for_node(rt.left, x) 
  elif rt.data < x:
   return self.search_for_node(rt.right, x)
  if rt.data == x:
   return rt

 def height_of_each_node(self, root,  x):
  '''
   let height of any node x be h(x)
   h(x) = 1 + max(l(x), r(x))
  ''' 
  rt = self.search_for_node(root, x)
  return self.height_of_tree(rt)

 def type_of_bst(self, rt):#TODO
  '''
    let c represent child of a node, and B represent the set of binary search tree items. 
   
  '''
  if not rt:
   return
  if (not rt.left and  rt.right) or (not rt.right and rt.left):
     print "binary tree" 
     return
  self.type_of_bst(rt.left)
  self.type_of_bst(rt.right)


'''
bst = BST()

root = None
for item  in sys.stdin:
 root = bst.insert(int(item), root)


# save the root object to file
# create a file called bst_tree and open it for writing
fh = open('/opt/DSA/Trees/bst/data/bst_0', 'wb')
pickle.dump(root, fh)
fh.close()



fh = open('/opt/DSA/Trees/bst/data/bst_0', 'r')
root = pickle.load(fh)
fh.close()


print "Preorder traversal:"
bst.preorder(root)

print "Inorder traversal:"
bst.inorder(root)

print "Postorder traversal:"
bst.postorder(root)

for line in sys.stdin:
 print bst.search(root, int(line))


bst.print_leaf_nodes(root)

print bst.nodes_count(root)


for line in sys.stdin:
 line = int(line)
 print "children of {}:".format(line),
 bst.node_children(root, line)
 print

 
print "height of tree:{}".format(bst.height_of_tree(root))


for line in sys.stdin:
 line = int(line)
 print "height of node:{} -> ".format(line),
 print bst.height_of_each_node(root, line) 
 print 
'''
