function InvertBinaryTree(root) {
  if (root === null) {
    return;
  }
  let tmp = root.Left;
  root.Left = root.Right;
  root.Right = tmp;
  InvertBinaryTree(root.Left);
  InvertBinaryTree(root.Right);
}
