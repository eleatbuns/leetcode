//给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。 
//
// candidates 中的数字可以无限制重复被选取。 
//
// 说明： 
//
// 
// 所有数字（包括 target）都是正整数。 
// 解集不能包含重复的组合。 
// 
//
// 示例 1: 
//
// 输入: candidates = [2,3,6,7], target = 7,
//所求解集为:
//[
//  [7],
//  [2,2,3]
//]
// 
//
// 示例 2: 
//
// 输入: candidates = [2,3,5], target = 8,
//所求解集为:
//[
//  [2,2,2,2],
//  [2,3,3],
//  [3,5]
//] 
// Related Topics 数组 回溯算法
package main

import (
	"sort"
)

//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return combination(candidates,target,[]int{},[][]int{})
}

func combination(candidates []int, remain int, currentPath []int, result [][]int) [][]int {
	if remain == 0 {
		abc := make([]int,0)
		abc = append(abc, currentPath...)
		result := append(result, abc)
		return result
	}
	for k, v := range candidates {
		if remain - v < 0 {
			break
		}
		currentPath := append(currentPath, v)
		result=combination(candidates[k:],remain - v,currentPath,result)
		currentPath = currentPath[:len(currentPath)-1]
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
