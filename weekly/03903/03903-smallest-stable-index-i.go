// ============================================================================
// 题号: 3903
// 题目: Smallest Stable Index I
// 难度: Easy
// 分类: Array, Prefix Sum
// 链接: https://leetcode.cn/problems/smallest-stable-index-i/
// 每日一题日期: 2026-09-04
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你一个长度为 n 的整数数组 nums 和一个整数 k。
//
// 对于每个下标 i，定义它的 不稳定值 为 max(nums[0..i]) - min(nums[i..n - 1])。
//
// 换句话说：
//
// max(nums[0..i]) 表示从下标 0 到下标 i 的元素中的 最大值 。
//
// min(nums[i..n - 1]) 表示从下标 i 到下标 n - 1 的元素中的 最小值 。
//
// 如果某个下标 i 的不稳定值 小于等于 k，则称该下标为 稳定下标 。
//
// 返回 最小 的稳定下标。如果不存在这样的下标，则返回 -1。
//
// 示例 1：
//
// 输入： nums = [5,0,1,4], k = 3
//
// 输出： 3
//
// 解释：
//
// 在下标 0 处：[5] 中的最大值是 5，[5, 0, 1, 4] 中的最小值是 0，因此不稳定值为 5 - 0 = 5。
//
// 在下标 1 处：[5, 0] 中的最大值是 5，[0, 1, 4] 中的最小值是 0，因此不稳定值为 5 - 0 = 5。
//
// 在下标 2 处：[5, 0, 1] 中的最大值是 5，[1, 4] 中的最小值是 1，因此不稳定值为 5 - 1 = 4。
//
// 在下标 3 处：[5, 0, 1, 4] 中的最大值是 5，[4] 中的最小值是 4，因此不稳定值为 5 - 4 = 1。
//
// 这是第一个不稳定值小于等于 k = 3 的下标，因此答案是 3。
//
// 示例 2：
//
// 输入： nums = [3,2,1], k = 1
//
// 输出： -1
//
// 解释：
//
// 在下标 0 处，不稳定值为 3 - 1 = 2。
//
// 在下标 1 处，不稳定值为 3 - 1 = 2。
//
// 在下标 2 处，不稳定值为 3 - 1 = 2。
//
// 这些值都不小于等于 k = 1，因此答案是 -1。
//
// 示例 3：
//
// 输入： nums = [0], k = 0
//
// 输出： 0
//
// 解释：
//
// 在下标 0 处，不稳定值为 0 - 0 = 0，它小于等于 k = 0。因此答案是 0。
//
// 提示：
//
// 1 <= nums.length <= 100
//
// 0 <= nums[i] <= 109
//
// 0 <= k <= 109
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// You are given an integer array nums of length n and an integer k.
//
// For each index i, define its instability score as max(nums[0..i]) - min(nums[i..n - 1]).
//
// In other words:
//
// max(nums[0..i]) is the largest value among the elements from index 0 to index i.
//
// min(nums[i..n - 1]) is the smallest value among the elements from index i to index n - 1.
//
// An index i is called stable if its instability score is less than or equal to k.
//
// Return the smallest stable index. If no such index exists, return -1.
//
// Example 1:
//
// Input: nums = [5,0,1,4], k = 3
//
// Output: 3
//
// Explanation:
//
// At index 0: The maximum in [5] is 5, and the minimum in [5, 0, 1, 4] is 0, so the instability score is 5 - 0 = 5.
//
// At index 1: The maximum in [5, 0] is 5, and the minimum in [0, 1, 4] is 0, so the instability score is 5 - 0 = 5.
//
// At index 2: The maximum in [5, 0, 1] is 5, and the minimum in [1, 4] is 1, so the instability score is 5 - 1 = 4.
//
// At index 3: The maximum in [5, 0, 1, 4] is 5, and the minimum in [4] is 4, so the instability score is 5 - 4 = 1.
//
// This is the first index with an instability score less than or equal to k = 3. Thus, the answer is 3.
//
// Example 2:
//
// Input: nums = [3,2,1], k = 1
//
// Output: -1
//
// Explanation:
//
// At index 0, the instability score is 3 - 1 = 2.
//
// At index 1, the instability score is 3 - 1 = 2.
//
// At index 2, the instability score is 3 - 1 = 2.
//
// None of these values is less than or equal to k = 1, so the answer is -1.
//
// Example 3:
//
// Input: nums = [0], k = 0
//
// Output: 0
//
// Explanation:
//
// At index 0, the instability score is 0 - 0 = 0, which is less than or equal to k = 0. Therefore, the answer is 0.
//
// Constraints:
//
// 1 <= nums.length <= 100
//
// 0 <= nums[i] <= 109
//
// 0 <= k <= 109
// ============================================================================

package leetcode

func firstStableIndex(nums []int, k int) int {
    
}
