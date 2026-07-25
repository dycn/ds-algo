// ============================================================================
// 题号: 3499
// 题目: Maximize Active Section with Trade I
// 难度: Medium
// 分类: String, Enumeration
// 链接: https://leetcode.cn/problems/maximize-active-section-with-trade-i/
// 每日一题日期: 2026-07-21
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你一个长度为 n 的二进制字符串 s，其中：
//
// '1' 表示一个 活跃 区段。
//
// '0' 表示一个 非活跃 区段。
//
// 你可以执行 最多一次操作 来最大化 s 中的活跃区段数量。在一次操作中，你可以：
//
// 将一个被 '0' 包围的连续 '1' 区块转换为全 '0'。
//
// 然后，将一个被 '1' 包围的连续 '0' 区块转换为全 '1'。
//
// 返回在执行最优操作后，s 中的 最大 活跃区段数。
//
// 注意：处理时需要在 s 的两侧加上 '1' ，即 t = '1' + s + '1'。这些加上的 '1' 不会影响最终的计数。
//
// 示例 1：
//
// 输入： s = "01"
//
// 输出： 1
//
// 解释：
//
// 因为没有被 '0' 包围的 '1' 区块，因此无法进行有效操作。最大活跃区段数为 1。
//
// 示例 2：
//
// 输入： s = "0100"
//
// 输出： 4
//
// 解释：
//
// 字符串 "0100" → 两端加上 '1' 后得到 "101001" 。
//
// 选择 "0100"，"101001" → "100001" → "111111" 。
//
// 最终的字符串去掉两端的 '1' 后为 "1111" 。最大活跃区段数为 4。
//
// 示例 3：
//
// 输入： s = "1000100"
//
// 输出： 7
//
// 解释：
//
// 字符串 "1000100" → 两端加上 '1' 后得到 "110001001" 。
//
// 选择 "000100"，"110001001" → "110000001" → "111111111"。
//
// 最终的字符串去掉两端的 '1' 后为 "1111111"。最大活跃区段数为 7。
//
// 示例 4：
//
// 输入： s = "01010"
//
// 输出： 4
//
// 解释：
//
// 字符串 "01010" → 两端加上 '1' 后得到 "1010101"。
//
// 选择 "010"，"1010101" → "1000101" → "1111101"。
//
// 最终的字符串去掉两端的 '1' 后为 "11110"。最大活跃区段数为 4。
//
// 提示：
//
// 1 <= n == s.length <= 105
//
// s[i] 仅包含 '0' 或 '1'
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// You are given a binary string s of length n, where:
//
// '1' represents an active section.
//
// '0' represents an inactive section.
//
// You can perform at most one trade to maximize the number of active sections in s. In a trade, you:
//
// Convert a contiguous block of '1's that is surrounded by '0's to all '0's.
//
// Afterward, convert a contiguous block of '0's that is surrounded by '1's to all '1's.
//
// Return the maximum number of active sections in s after making the optimal trade.
//
// Note: Treat s as if it is augmented with a '1' at both ends, forming t = '1' + s + '1'. The augmented '1's do not contribute to the final count.
//
// Example 1:
//
// Input: s = "01"
//
// Output: 1
//
// Explanation:
//
// Because there is no block of '1's surrounded by '0's, no valid trade is possible. The maximum number of active sections is 1.
//
// Example 2:
//
// Input: s = "0100"
//
// Output: 4
//
// Explanation:
//
// String "0100" → Augmented to "101001".
//
// Choose "0100", convert "101001" → "100001" → "111111".
//
// The final string without augmentation is "1111". The maximum number of active sections is 4.
//
// Example 3:
//
// Input: s = "1000100"
//
// Output: 7
//
// Explanation:
//
// String "1000100" → Augmented to "110001001".
//
// Choose "000100", convert "110001001" → "110000001" → "111111111".
//
// The final string without augmentation is "1111111". The maximum number of active sections is 7.
//
// Example 4:
//
// Input: s = "01010"
//
// Output: 4
//
// Explanation:
//
// String "01010" → Augmented to "1010101".
//
// Choose "010", convert "1010101" → "1000101" → "1111101".
//
// The final string without augmentation is "11110". The maximum number of active sections is 4.
//
// Constraints:
//
// 1 <= n == s.length <= 105
//
// s[i] is either '0' or '1'
// ============================================================================

package leetcode

func maxActiveSectionsAfterTrade(s string) int {
    
}
