// ============================================================================
// 题号: 115
// 题目: Distinct Subsequences
// 难度: Hard
// 分类: String, Dynamic Programming
// 链接: https://leetcode.cn/problems/distinct-subsequences/
// 每日一题日期: 2026-09-06
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你两个字符串 s 和 t ，统计并返回在 s 的 子序列 中 t 出现的个数。
//
// 测试用例保证结果在 32 位有符号整数范围内。
//
// 示例 1：
//
// 输入：s = "rabbbit", t = "rabbit"
// 输出：3
// 解释：
// 如下所示, 有 3 种可以从 s 中得到 "rabbit" 的方案。
// rabbbit
// rabbbit
// rabbbit
//
// 示例 2：
//
// 输入：s = "babgbag", t = "bag"
// 输出：5
// 解释：
// 如下所示, 有 5 种可以从 s 中得到 "bag" 的方案。
// babgbag
// babgbag
// babgbag
// babgbag
// babgbag
//
// 提示：
//
// 1 <= s.length, t.length <= 1000
//
// s 和 t 由英文字母组成
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// Given two strings s and t, return the number of distinct subsequences of s which equals t.
//
// The test cases are generated so that the answer fits on a 32-bit signed integer.
//
// Example 1:
//
// Input: s = "rabbbit", t = "rabbit"
// Output: 3
// Explanation:
// As shown below, there are 3 ways you can generate "rabbit" from s.
// rabbbit
// rabbbit
// rabbbit
//
// Example 2:
//
// Input: s = "babgbag", t = "bag"
// Output: 5
// Explanation:
// As shown below, there are 5 ways you can generate "bag" from s.
// babgbag
// babgbag
// babgbag
// babgbag
// babgbag
//
// Constraints:
//
// 1 <= s.length, t.length <= 1000
//
// s and t consist of English letters.
// ============================================================================

package leetcode

func numDistinct(s string, t string) int {
    
}
