//给定一个二叉树，找出其最小深度。 
//
// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。 
//
// 说明: 叶子节点是指没有子节点的节点。 
//
// 示例: 
//
// 给定二叉树 [3,9,20,null,null,15,7], 
//
//     3
//   / \
//  9  20
//    /  \
//   15   7 
//
// 返回它的最小深度 2. 
// Related Topics 树 深度优先搜索 广度优先搜索
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

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	minLen := 0
	if root.Left == nil && root.Right == nil {
		return 1
	} else if root.Left != nil && root.Right != nil {
		minLen = Min(minDepth(root.Left), minDepth(root.Right)) + 1
	} else if root.Left != nil {
		minLen = minDepth(root.Left) + 1
	} else {
		minLen = minDepth(root.Right) + 1
	}

	return minLen
}
func Min(n1 int, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

//leetcode submit region end(Prohibit modification and deletion)
