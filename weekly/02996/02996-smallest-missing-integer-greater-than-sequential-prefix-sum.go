// ============================================================================
// 题号: 2996
// 题目: Smallest Missing Integer Greater Than Sequential Prefix Sum
// 难度: Easy
// 分类: Array, Hash Table, Sorting
// 链接: https://leetcode.cn/problems/smallest-missing-integer-greater-than-sequential-prefix-sum/
// 每日一题日期: 2026-08-11
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你一个下标从 0 开始的整数数组 nums 。
//
// 如果一个前缀 nums[0..i] 满足对于 1 <= j <= i 的所有元素都有 nums[j] = nums[j - 1] + 1 ，那么我们称这个前缀是一个 顺序前缀 。特殊情况是，只包含 nums[0] 的前缀也是一个 顺序前缀 。
//
// 请你返回 nums 中没有出现过的 最小 整数 x ，满足 x 大于等于 最长 顺序前缀的和。
//
// 示例 1：
//
// 输入：nums = [1,2,3,2,5]
// 输出：6
// 解释：nums 的最长顺序前缀是 [1,2,3] ，和为 6 ，6 不在数组中，所以 6 是大于等于最长顺序前缀和的最小整数。
//
// 示例 2：
//
// 输入：nums = [3,4,5,1,12,14,13]
// 输出：15
// 解释：nums 的最长顺序前缀是 [3,4,5] ，和为 12 ，12、13 和 14 都在数组中，但 15 不在，所以 15 是大于等于最长顺序前缀和的最小整数。
//
// 提示：
//
// 1 <= nums.length <= 50
//
// 1 <= nums[i] <= 50
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// You are given a 0-indexed array of integers nums.
//
// A prefix nums[0..i] is sequential if, for all 1 <= j <= i, nums[j] = nums[j - 1] + 1. In particular, the prefix consisting only of nums[0] is sequential.
//
// Return the smallest integer x missing from nums such that x is greater than or equal to the sum of the longest sequential prefix.
//
// Example 1:
//
// Input: nums = [1,2,3,2,5]
// Output: 6
// Explanation: The longest sequential prefix of nums is [1,2,3] with a sum of 6. 6 is not in the array, therefore 6 is the smallest missing integer greater than or equal to the sum of the longest sequential prefix.
//
// Example 2:
//
// Input: nums = [3,4,5,1,12,14,13]
// Output: 15
// Explanation: The longest sequential prefix of nums is [3,4,5] with a sum of 12. 12, 13, and 14 belong to the array while 15 does not. Therefore 15 is the smallest missing integer greater than or equal to the sum of the longest sequential prefix.
//
// Constraints:
//
// 1 <= nums.length <= 50
//
// 1 <= nums[i] <= 50
// ============================================================================

package leetcode

func missingInteger(nums []int) int {
    
}
