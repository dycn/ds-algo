// ============================================================================
// 题号: 2958
// 题目: Length of Longest Subarray With at Most K Frequency
// 难度: Medium
// 分类: Array, Hash Table, Sliding Window
// 链接: https://leetcode.cn/problems/length-of-longest-subarray-with-at-most-k-frequency/
// 每日一题日期: 2026-08-12
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你一个整数数组 nums 和一个整数 k 。
//
// 一个元素 x 在数组中的 频率 指的是它在数组中的出现次数。
//
// 如果一个数组中所有元素的频率都 小于等于 k ，那么我们称这个数组是 好 数组。
//
// 请你返回 nums 中 最长好 子数组的长度。
//
// 子数组 指的是一个数组中一段连续非空的元素序列。
//
// 示例 1：
//
// 输入：nums = [1,2,3,1,2,3,1,2], k = 2
// 输出：6
// 解释：最长好子数组是 [1,2,3,1,2,3] ，值 1 ，2 和 3 在子数组中的频率都没有超过 k = 2 。[2,3,1,2,3,1] 和 [3,1,2,3,1,2] 也是好子数组。
// 最长好子数组的长度为 6 。
//
// 示例 2：
//
// 输入：nums = [1,2,1,2,1,2,1,2], k = 1
// 输出：2
// 解释：最长好子数组是 [1,2] ，值 1 和 2 在子数组中的频率都没有超过 k = 1 。[2,1] 也是好子数组。
// 最长好子数组的长度为 2 。
//
// 示例 3：
//
// 输入：nums = [5,5,5,5,5,5,5], k = 4
// 输出：4
// 解释：最长好子数组是 [5,5,5,5] ，值 5 在子数组中的频率没有超过 k = 4 。
// 最长好子数组的长度为 4 。
//
// 提示：
//
// 1 <= nums.length <= 105
//
// 1 <= nums[i] <= 109
//
// 1 <= k <= nums.length
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// You are given an integer array nums and an integer k.
//
// The frequency of an element x is the number of times it occurs in an array.
//
// An array is called good if the frequency of each element in this array is less than or equal to k.
//
// Return the length of the longest good subarray of nums.
//
// A subarray is a contiguous non-empty sequence of elements within an array.
//
// Example 1:
//
// Input: nums = [1,2,3,1,2,3,1,2], k = 2
// Output: 6
// Explanation: The longest possible good subarray is [1,2,3,1,2,3] since the values 1, 2, and 3 occur at most twice in this subarray. Note that the subarrays [2,3,1,2,3,1] and [3,1,2,3,1,2] are also good.
// It can be shown that there are no good subarrays with length more than 6.
//
// Example 2:
//
// Input: nums = [1,2,1,2,1,2,1,2], k = 1
// Output: 2
// Explanation: The longest possible good subarray is [1,2] since the values 1 and 2 occur at most once in this subarray. Note that the subarray [2,1] is also good.
// It can be shown that there are no good subarrays with length more than 2.
//
// Example 3:
//
// Input: nums = [5,5,5,5,5,5,5], k = 4
// Output: 4
// Explanation: The longest possible good subarray is [5,5,5,5] since the value 5 occurs 4 times in this subarray.
// It can be shown that there are no good subarrays with length more than 4.
//
// Constraints:
//
// 1 <= nums.length <= 105
//
// 1 <= nums[i] <= 109
//
// 1 <= k <= nums.length
// ============================================================================

package leetcode

func maxSubarrayLength(nums []int, k int) int {
    
}
