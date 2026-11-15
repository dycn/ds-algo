// ============================================================================
// 题号: 2243
// 题目: Calculate Digit Sum of a String
// 难度: Easy
// 分类: String, Simulation
// 链接: https://leetcode.cn/problems/calculate-digit-sum-of-a-string/
// 每日一题日期: 2026-11-16
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你一个由若干数字（0 - 9）组成的字符串 s ，和一个整数。
//
// 如果 s 的长度大于 k ，则可以执行一轮操作。在一轮操作中，需要完成以下工作：
//
// 将 s 拆分 成长度为 k 的若干 连续数字组 ，使得前 k 个字符都分在第一组，接下来的 k 个字符都分在第二组，依此类推。注意，最后一个数字组的长度可以小于 k 。
//
// 用表示每个数字组中所有数字之和的字符串来 替换 对应的数字组。例如，"346" 会替换为 "13" ，因为 3 + 4 + 6 = 13 。
//
// 合并 所有组以形成一个新字符串。如果新字符串的长度大于 k 则重复第一步。
//
// 返回在完成所有轮操作后的 s 。
//
// 示例 1：
//
// 输入：s = "11111222223", k = 3
// 输出："135"
// 解释：
// - 第一轮，将 s 分成："111"、"112"、"222" 和 "23" 。
//   接着，计算每一组的数字和：1 + 1 + 1 = 3、1 + 1 + 2 = 4、2 + 2 + 2 = 6 和 2 + 3 = 5 。
//   这样，s 在第一轮之后变成 "3" + "4" + "6" + "5" = "3465" 。
// - 第二轮，将 s 分成："346" 和 "5" 。
//   接着，计算每一组的数字和：3 + 4 + 6 = 13 、5 = 5 。
//   这样，s 在第二轮之后变成 "13" + "5" = "135" 。
// 现在，s.length <= k ，所以返回 "135" 作为答案。
//
// 示例 2：
//
// 输入：s = "00000000", k = 3
// 输出："000"
// 解释：
// 将 "000", "000", and "00".
// 接着，计算每一组的数字和：0 + 0 + 0 = 0 、0 + 0 + 0 = 0 和 0 + 0 = 0 。
// s 变为 "0" + "0" + "0" = "000" ，其长度等于 k ，所以返回 "000" 。
//
// 提示：
//
// 1 <= s.length <= 100
//
// 2 <= k <= 100
//
// s 仅由数字（0 - 9）组成。
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// You are given a string s consisting of digits and an integer k.
//
// A round can be completed if the length of s is greater than k. In one round, do the following:
//
// Divide s into consecutive groups of size k such that the first k characters are in the first group, the next k characters are in the second group, and so on. Note that the size of the last group can be smaller than k.
//
// Replace each group of s with a string representing the sum of all its digits. For example, "346" is replaced with "13" because 3 + 4 + 6 = 13.
//
// Merge consecutive groups together to form a new string. If the length of the string is greater than k, repeat from step 1.
//
// Return s after all rounds have been completed.
//
// Example 1:
//
// Input: s = "11111222223", k = 3
// Output: "135"
// Explanation:
// - For the first round, we divide s into groups of size 3: "111", "112", "222", and "23".
//   ​​​​​Then we calculate the digit sum of each group: 1 + 1 + 1 = 3, 1 + 1 + 2 = 4, 2 + 2 + 2 = 6, and 2 + 3 = 5.
//   So, s becomes "3" + "4" + "6" + "5" = "3465" after the first round.
// - For the second round, we divide s into "346" and "5".
//   Then we calculate the digit sum of each group: 3 + 4 + 6 = 13, 5 = 5.
//   So, s becomes "13" + "5" = "135" after second round.
// Now, s.length <= k, so we return "135" as the answer.
//
// Example 2:
//
// Input: s = "00000000", k = 3
// Output: "000"
// Explanation:
// We divide s into "000", "000", and "00".
// Then we calculate the digit sum of each group: 0 + 0 + 0 = 0, 0 + 0 + 0 = 0, and 0 + 0 = 0.
// s becomes "0" + "0" + "0" = "000", whose length is equal to k, so we return "000".
//
// Constraints:
//
// 1 <= s.length <= 100
//
// 2 <= k <= 100
//
// s consists of digits only.
// ============================================================================

package leetcode

func digitSum(s string, k int) string {
    
}
