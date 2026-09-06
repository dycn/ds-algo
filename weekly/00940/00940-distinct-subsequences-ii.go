// ============================================================================
// 题号: 940
// 题目: Distinct Subsequences II
// 难度: Hard
// 分类: String, Dynamic Programming
// 链接: https://leetcode.cn/problems/distinct-subsequences-ii/
// 每日一题日期: 2026-09-07
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给定一个字符串 s，计算 s 的 不同非空子序列 的个数。因为结果可能很大，所以返回答案需要对 10^9 + 7 取余 。
//
// 字符串的 子序列 是经由原字符串删除一些（也可能不删除）字符但不改变剩余字符相对位置的一个新字符串。
//
// 例如，"ace" 是 "abcde" 的一个子序列，但 "aec" 不是。
//
// 示例 1：
//
// 输入：s = "abc"
// 输出：7
// 解释：7 个不同的子序列分别是 "a", "b", "c", "ab", "ac", "bc", 以及 "abc"。
//
// 示例 2：
//
// 输入：s = "aba"
// 输出：6
// 解释：6 个不同的子序列分别是 "a", "b", "ab", "ba", "aa" 以及 "aba"。
//
// 示例 3：
//
// 输入：s = "aaa"
// 输出：3
// 解释：3 个不同的子序列分别是 "a", "aa" 以及 "aaa"。
//
// 提示：
//
// 1 <= s.length <= 2000
//
// s 仅由小写英文字母组成
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// Given a string s, return the number of distinct non-empty subsequences of s. Since the answer may be very large, return it modulo 109 + 7.
//
// A subsequence of a string is a new string that is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not.
//
// Example 1:
//
// Input: s = "abc"
// Output: 7
// Explanation: The 7 distinct subsequences are "a", "b", "c", "ab", "ac", "bc", and "abc".
//
// Example 2:
//
// Input: s = "aba"
// Output: 6
// Explanation: The 6 distinct subsequences are "a", "b", "ab", "aa", "ba", and "aba".
//
// Example 3:
//
// Input: s = "aaa"
// Output: 3
// Explanation: The 3 distinct subsequences are "a", "aa" and "aaa".
//
// Constraints:
//
// 1 <= s.length <= 2000
//
// s consists of lowercase English letters.
// ============================================================================

package leetcode

func distinctSubseqII(s string) int {
    
}
