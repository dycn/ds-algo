// ============================================================================
// 题号: 2213
// 题目: Longest Substring of One Repeating Character
// 难度: Hard
// 分类: Segment Tree, Array, String, Ordered Set
// 链接: https://leetcode.cn/problems/longest-substring-of-one-repeating-character/
// 每日一题日期: 2026-08-13
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你一个下标从 0 开始的字符串 s 。另给你一个下标从 0 开始、长度为 k 的字符串 queryCharacters ，一个下标从 0 开始、长度也是 k 的整数 下标 数组 queryIndices ，这两个都用来描述 k 个查询。
//
// 第 i 个查询会将 s 中位于下标 queryIndices[i] 的字符更新为 queryCharacters[i] 。
//
// 返回一个长度为 k 的数组 lengths ，其中 lengths[i] 是在执行第 i 个查询 之后 s 中仅由 单个字符重复 组成的 最长子字符串 的 长度 。
//
// 示例 1：
//
// 输入：s = "babacc", queryCharacters = "bcb", queryIndices = [1,3,3]
// 输出：[3,3,4]
// 解释：
// - 第 1 次查询更新后 s = "bbbacc" 。由单个字符重复组成的最长子字符串是 "bbb" ，长度为 3 。
// - 第 2 次查询更新后 s = "bbbccc" 。由单个字符重复组成的最长子字符串是 "bbb" 或 "ccc"，长度为 3 。
// - 第 3 次查询更新后 s = "bbbbcc" 。由单个字符重复组成的最长子字符串是 "bbbb" ，长度为 4 。
// 因此，返回 [3,3,4] 。
//
// 示例 2：
//
// 输入：s = "abyzz", queryCharacters = "aa", queryIndices = [2,1]
// 输出：[2,3]
// 解释：
// - 第 1 次查询更新后 s = "abazz" 。由单个字符重复组成的最长子字符串是 "zz" ，长度为 2 。
// - 第 2 次查询更新后 s = "aaazz" 。由单个字符重复组成的最长子字符串是 "aaa" ，长度为 3 。
// 因此，返回 [2,3] 。
//
// 提示：
//
// 1 <= s.length <= 105
//
// s 由小写英文字母组成
//
// k == queryCharacters.length == queryIndices.length
//
// 1 <= k <= 105
//
// queryCharacters 由小写英文字母组成
//
// 0 <= queryIndices[i] < s.length
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// You are given a 0-indexed string s. You are also given a 0-indexed string queryCharacters of length k and a 0-indexed array of integer indices queryIndices of length k, both of which are used to describe k queries.
//
// The ith query updates the character in s at index queryIndices[i] to the character queryCharacters[i].
//
// Return an array lengths of length k where lengths[i] is the length of the longest substring of s consisting of only one repeating character after the ith query is performed.
//
// Example 1:
//
// Input: s = "babacc", queryCharacters = "bcb", queryIndices = [1,3,3]
// Output: [3,3,4]
// Explanation:
// - 1st query updates s = "bbbacc". The longest substring consisting of one repeating character is "bbb" with length 3.
// - 2nd query updates s = "bbbccc".
//   The longest substring consisting of one repeating character can be "bbb" or "ccc" with length 3.
// - 3rd query updates s = "bbbbcc". The longest substring consisting of one repeating character is "bbbb" with length 4.
// Thus, we return [3,3,4].
//
// Example 2:
//
// Input: s = "abyzz", queryCharacters = "aa", queryIndices = [2,1]
// Output: [2,3]
// Explanation:
// - 1st query updates s = "abazz". The longest substring consisting of one repeating character is "zz" with length 2.
// - 2nd query updates s = "aaazz". The longest substring consisting of one repeating character is "aaa" with length 3.
// Thus, we return [2,3].
//
// Constraints:
//
// 1 <= s.length <= 105
//
// s consists of lowercase English letters.
//
// k == queryCharacters.length == queryIndices.length
//
// 1 <= k <= 105
//
// queryCharacters consists of lowercase English letters.
//
// 0 <= queryIndices[i] < s.length
// ============================================================================

package leetcode

func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {
    
}
