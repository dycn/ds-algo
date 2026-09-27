// ============================================================================
// 题号: 1190
// 题目: Reverse Substrings Between Each Pair of Parentheses
// 难度: Medium
// 分类: Stack, String, 括号序列
// 链接: https://leetcode.cn/problems/reverse-substrings-between-each-pair-of-parentheses/
// 每日一题日期: 2026-09-27
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给出一个字符串 s（仅含有小写英文字母和括号）。
//
// 请你按照从括号内到外的顺序，逐层反转每对匹配括号中的字符串，并返回最终的结果。
//
// 注意，您的结果中 不应 包含任何括号。
//
// 示例 1：
//
// 输入：s = "(abcd)"
// 输出："dcba"
//
// 示例 2：
//
// 输入：s = "(u(love)i)"
// 输出："iloveu"
// 解释：先反转子字符串 "love" ，然后反转整个字符串。
//
// 示例 3：
//
// 输入：s = "(ed(et(oc))el)"
// 输出："leetcode"
// 解释：先反转子字符串 "oc" ，接着反转 "etco" ，然后反转整个字符串。
//
// 提示：
//
// 1 <= s.length <= 2000
//
// s 中只有小写英文字母和括号
//
// 题目测试用例确保所有括号都是成对出现的
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// You are given a string s that consists of lower case English letters and brackets.
//
// Reverse the strings in each pair of matching parentheses, starting from the innermost one.
//
// Your result should not contain any brackets.
//
// Example 1:
//
// Input: s = "(abcd)"
// Output: "dcba"
//
// Example 2:
//
// Input: s = "(u(love)i)"
// Output: "iloveu"
// Explanation: The substring "love" is reversed first, then the whole string is reversed.
//
// Example 3:
//
// Input: s = "(ed(et(oc))el)"
// Output: "leetcode"
// Explanation: First, we reverse the substring "oc", then "etco", and finally, the whole string.
//
// Constraints:
//
// 1 <= s.length <= 2000
//
// s only contains lower case English characters and parentheses.
//
// It is guaranteed that all parentheses are balanced.
// ============================================================================

package leetcode

func reverseParentheses(s string) string {
    
}
