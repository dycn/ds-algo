// ============================================================================
// 题号: 1614
// 题目: Maximum Nesting Depth of the Parentheses
// 难度: Easy
// 分类: Stack, String, 括号序列
// 链接: https://leetcode.cn/problems/maximum-nesting-depth-of-the-parentheses/
// 每日一题日期: 2026-09-28
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给定 有效括号字符串 s，返回 s 的 嵌套深度。嵌套深度是嵌套括号的 最大 数量。
//
// 示例 1：
//
// 输入：s = "(1+(2*3)+((8)/4))+1"
//
// 输出：3
//
// 解释：数字 8 在嵌套的 3 层括号中。
//
// 示例 2：
//
// 输入：s = "(1)+((2))+(((3)))"
//
// 输出：3
//
// 解释：数字 3 在嵌套的 3 层括号中。
//
// 示例 3：
//
// 输入：s = "()(())((()()))"
//
// 输出：3
//
// 提示：
//
// 1 <= s.length <= 100
//
// s 由数字 0-9 和字符 '+'、'-'、'*'、'/'、'('、')' 组成
//
// 题目数据保证括号字符串 s 是 有效的括号字符串
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// Given a valid parentheses string s, return the nesting depth of s. The nesting depth is the maximum number of nested parentheses.
//
// Example 1:
//
// Input: s = "(1+(2*3)+((8)/4))+1"
//
// Output: 3
//
// Explanation:
//
// Digit 8 is inside of 3 nested parentheses in the string.
//
// Example 2:
//
// Input: s = "(1)+((2))+(((3)))"
//
// Output: 3
//
// Explanation:
//
// Digit 3 is inside of 3 nested parentheses in the string.
//
// Example 3:
//
// Input: s = "()(())((()()))"
//
// Output: 3
//
// Constraints:
//
// 1 <= s.length <= 100
//
// s consists of digits 0-9 and characters '+', '-', '*', '/', '(', and ')'.
//
// It is guaranteed that parentheses expression s is a VPS.
// ============================================================================

package leetcode

func maxDepth(s string) int {
    
}
