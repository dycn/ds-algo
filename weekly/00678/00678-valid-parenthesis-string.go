// ============================================================================
// 题号: 678
// 题目: Valid Parenthesis String
// 难度: Medium
// 分类: Stack, Greedy, String, Dynamic Programming, 括号序列
// 链接: https://leetcode.cn/problems/valid-parenthesis-string/
// 每日一题日期: 2026-10-04
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你一个只包含三种字符的字符串，支持的字符类型分别是 '('、')' 和 '*'。请你检验这个字符串是否为有效字符串，如果是 有效 字符串返回 true 。
//
// 有效 字符串符合如下规则：
//
// 任何左括号 '(' 必须有相应的右括号 ')'。
//
// 任何右括号 ')' 必须有相应的左括号 '(' 。
//
// 左括号 '(' 必须在对应的右括号之前 ')'。
//
// '*' 可以被视为单个右括号 ')' ，或单个左括号 '(' ，或一个空字符串 ""。
//
// 示例 1：
//
// 输入：s = "()"
// 输出：true
//
// 示例 2：
//
// 输入：s = "(*)"
// 输出：true
//
// 示例 3：
//
// 输入：s = "(*))"
// 输出：true
//
// 提示：
//
// 1 <= s.length <= 100
//
// s[i] 为 '('、')' 或 '*'
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// Given a string s containing only three types of characters: '(', ')' and '*', return true if s is valid.
//
// The following rules define a valid string:
//
// Any left parenthesis '(' must have a corresponding right parenthesis ')'.
//
// Any right parenthesis ')' must have a corresponding left parenthesis '('.
//
// Left parenthesis '(' must go before the corresponding right parenthesis ')'.
//
// '*' could be treated as a single right parenthesis ')' or a single left parenthesis '(' or an empty string "".
//
// Example 1:
//
// Input: s = "()"
// Output: true
//
// Example 2:
//
// Input: s = "(*)"
// Output: true
//
// Example 3:
//
// Input: s = "(*))"
// Output: true
//
// Example 4:
//
// Input: s = "("
// Output: false
//
// Constraints:
//
// 1 <= s.length <= 100
//
// s[i] is '(', ')' or '*'.
// ============================================================================

package leetcode

func checkValidString(s string) bool {
    
}
