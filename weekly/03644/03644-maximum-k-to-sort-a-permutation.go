// ============================================================================
// 题号: 3644
// 题目: Maximum K to Sort a Permutation
// 难度: Medium
// 分类: Bit Manipulation, Array
// 链接: https://leetcode.cn/problems/maximum-k-to-sort-a-permutation/
// 每日一题日期: 2026-10-23
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你一个长度为 n 的整数数组 nums，其中 nums 是范围 [0..n - 1] 内所有数字的一个 排列 。
//
// 你可以在满足条件 nums[i] AND nums[j] == k 的情况下交换下标 i 和 j 的元素，其中 AND 表示按位与操作，k 是一个非负整数。
//
// 返回可以使数组按 非递减 顺序排序的最大值 k（允许进行任意次这样的交换）。如果 nums 已经是有序的，返回 0。
//
// 排列 是数组所有元素的一种重新排列。
//
// 示例 1：
//
// 输入：nums = [0,3,2,1]
//
// 输出：1
//
// 解释：
//
// 选择 k = 1。交换 nums[1] = 3 和 nums[3] = 1，因为 nums[1] AND nums[3] == 1，从而得到一个排序后的排列：[0, 1, 2, 3]。
//
// 示例 2：
//
// 输入：nums = [0,1,3,2]
//
// 输出：2
//
// 解释：
//
// 选择 k = 2。交换 nums[2] = 3 和 nums[3] = 2，因为 nums[2] AND nums[3] == 2，从而得到一个排序后的排列：[0, 1, 2, 3]。
//
// 示例 3：
//
// 输入：nums = [3,2,1,0]
//
// 输出：0
//
// 解释：
//
// 只有当 k = 0 时，才能进行排序，因为没有更大的 k 能够满足 nums[i] AND nums[j] == k 的交换条件。
//
// 提示：
//
// 1 <= n == nums.length <= 105
//
// 0 <= nums[i] <= n - 1
//
// nums 是从 0 到 n - 1 的一个排列。
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// You are given an integer array nums of length n, where nums is a permutation of the numbers in the range [0..n - 1].
//
// You may swap elements at indices i and j only if nums[i] AND nums[j] == k, where AND denotes the bitwise AND operation and k is a non-negative integer.
//
// Return the maximum value of k such that the array can be sorted in non-decreasing order using any number of such swaps. If nums is already sorted, return 0.
//
// Example 1:
//
// Input: nums = [0,3,2,1]
//
// Output: 1
//
// Explanation:
//
// Choose k = 1. Swapping nums[1] = 3 and nums[3] = 1 is allowed since nums[1] AND nums[3] == 1, resulting in a sorted permutation: [0, 1, 2, 3].
//
// Example 2:
//
// Input: nums = [0,1,3,2]
//
// Output: 2
//
// Explanation:
//
// Choose k = 2. Swapping nums[2] = 3 and nums[3] = 2 is allowed since nums[2] AND nums[3] == 2, resulting in a sorted permutation: [0, 1, 2, 3].
//
// Example 3:
//
// Input: nums = [3,2,1,0]
//
// Output: 0
//
// Explanation:
//
// Only k = 0 allows sorting since no greater k allows the required swaps where nums[i] AND nums[j] == k.
//
// Constraints:
//
// 1 <= n == nums.length <= 105
//
// 0 <= nums[i] <= n - 1
//
// nums is a permutation of integers from 0 to n - 1.
// ============================================================================

package leetcode

func sortPermutation(nums []int) int {
    
}
