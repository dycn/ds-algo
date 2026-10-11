// ============================================================================
// 题号: 921
// 题目: Minimum Add to Make Parentheses Valid
// 难度: Medium
// 分类: Stack, Greedy, String, 括号序列
// 链接: https://leetcode.cn/problems/minimum-add-to-make-parentheses-valid/
// 每日一题日期: 2026-10-06
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 只有满足下面几点之一，括号字符串才是有效的：
//
// 它是一个空字符串，或者
//
// 它可以被写成 AB （A 与 B 连接）, 其中 A 和 B 都是有效字符串，或者
//
// 它可以被写作 (A)，其中 A 是有效字符串。
//
// 给定一个括号字符串 s ，在每一次操作中，你都可以在字符串的任何位置插入一个括号
//
// 例如，如果 s = "()))" ，你可以插入一个开始括号为 "(()))" 或结束括号为 "())))" 。
//
// 返回 为使结果字符串 s 有效而必须添加的最少括号数。
//
// 示例 1：
//
// 输入：s = "())"
// 输出：1
//
// 示例 2：
//
// 输入：s = "((("
// 输出：3
//
// 提示：
//
// 1 <= s.length <= 1000
//
// s 只包含 '(' 和 ')' 字符。
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// A parentheses string is valid if and only if:
//
// It is the empty string,
//
// It can be written as AB (A concatenated with B), where A and B are valid strings, or
//
// It can be written as (A), where A is a valid string.
//
// You are given a parentheses string s. In one move, you can insert a parenthesis at any position of the string.
//
// For example, if s = "()))", you can insert an opening parenthesis to be "(()))" or a closing parenthesis to be "())))".
//
// Return the minimum number of moves required to make s valid.
//
// Example 1:
//
// Input: s = "())"
// Output: 1
//
// Example 2:
//
// Input: s = "((("
// Output: 3
//
// Constraints:
//
// 1 <= s.length <= 1000
//
// s[i] is either '(' or ')'.
// ============================================================================

package leetcode

func minAddToMakeValid(s string) int {
    
}
