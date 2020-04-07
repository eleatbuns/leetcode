//根据一棵树的中序遍历与后序遍历构造二叉树。 
//
// 注意: 
//你可以假设树中没有重复的元素。 
//
// 例如，给出 
//
// 中序遍历 inorder = [9,3,15,20,7]
//后序遍历 postorder = [9,15,7,20,3] 
//
// 返回如下的二叉树： 
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
// 
// Related Topics 树 深度优先搜索 数组
package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0{
		return nil
	}
	root := &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  nil,
		Right: nil,
	}
	rootPos := 0
	for k, v := range inorder {
		if v == root.Val {
			rootPos = k
			break
		}
	}
	root.Left = buildTree(inorder[:rootPos],postorder[:rootPos])
	root.Right = buildTree(inorder[rootPos+1:],postorder[rootPos:len(postorder)-1])
	return root
}
//leetcode submit region end(Prohibit modification and deletion)
