//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。 
// 左括号必须以正确的顺序闭合。
// 注意空字符串可被认为是有效字符串。
// 示例 1:
// 输入: "()"
//输出: true
// 示例 2:
// 输入: "()[]{}"
//输出: true
// 示例 3:
// 输入: "(]"
//输出: false
// 示例 4:
// 输入: "([)]"
//输出: false
// 示例 5:
// 输入: "{[]}"
//输出: true 
// Related Topics 栈 字符串
package main

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	arr := make([]byte, 0)
	index := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			arr = append(arr, s[i])
			index++
		} else if index > 0 && ((s[i] == ')' && arr[index-1] == '(') ||
			(s[i] == ']' && arr[index-1] == '[') ||
			(s[i] == '}' && arr[index-1] == '{')) {
			arr = arr[0:index-1]
			index--
		} else {
			return false
		}
	}
	if len(arr) == 0 {
		return true
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)
