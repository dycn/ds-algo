// ============================================================================
// 题号: 3364
// 题目: Minimum Positive Sum Subarray 
// 难度: Easy
// 分类: Array, Prefix Sum, Sliding Window
// 链接: https://leetcode.cn/problems/minimum-positive-sum-subarray/
// 每日一题日期: 2026-11-14
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你一个整数数组 nums 和 两个 整数 l 和 r。你的任务是找到一个长度在 l 和 r 之间（包含）且和大于 0 的 子数组 的 最小 和。
//
// 返回满足条件的子数组的 最小 和。如果不存在这样的子数组，则返回 -1。
//
// 子数组 是数组中的一个连续 非空 元素序列。
//
// 示例 1：
//
// 输入： nums = [3, -2, 1, 4], l = 2, r = 3
//
// 输出： 1
//
// 解释：
//
// 长度在 l = 2 和 r = 3 之间且和大于 0 的子数组有：
//
// [3, -2] 和为 1
//
// [1, 4] 和为 5
//
// [3, -2, 1] 和为 2
//
// [-2, 1, 4] 和为 3
//
// 其中，子数组 [3, -2] 的和为 1，是所有正和中最小的。因此，答案为 1。
//
// 示例 2：
//
// 输入： nums = [-2, 2, -3, 1], l = 2, r = 3
//
// 输出： -1
//
// 解释：
//
// 不存在长度在 l 和 r 之间且和大于 0 的子数组。因此，答案为 -1。
//
// 示例 3：
//
// 输入： nums = [1, 2, 3, 4], l = 2, r = 4
//
// 输出： 3
//
// 解释：
//
// 子数组 [1, 2] 的长度为 2，和为 3，是所有正和中最小的。因此，答案为 3。
//
// 提示：
//
// 1 <= nums.length <= 100
//
// 1 <= l <= r <= nums.length
//
// -1000 <= nums[i] <= 1000
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// You are given an integer array nums and two integers l and r. Your task is to find the minimum sum of a subarray whose size is between l and r (inclusive) and whose sum is greater than 0.
//
// Return the minimum sum of such a subarray. If no such subarray exists, return -1.
//
// A subarray is a contiguous non-empty sequence of elements within an array.
//
// Example 1:
//
// Input: nums = [3, -2, 1, 4], l = 2, r = 3
//
// Output: 1
//
// Explanation:
//
// The subarrays of length between l = 2 and r = 3 where the sum is greater than 0 are:
//
// [3, -2] with a sum of 1
//
// [1, 4] with a sum of 5
//
// [3, -2, 1] with a sum of 2
//
// [-2, 1, 4] with a sum of 3
//
// Out of these, the subarray [3, -2] has a sum of 1, which is the smallest positive sum. Hence, the answer is 1.
//
// Example 2:
//
// Input: nums = [-2, 2, -3, 1], l = 2, r = 3
//
// Output: -1
//
// Explanation:
//
// There is no subarray of length between l and r that has a sum greater than 0. So, the answer is -1.
//
// Example 3:
//
// Input: nums = [1, 2, 3, 4], l = 2, r = 4
//
// Output: 3
//
// Explanation:
//
// The subarray [1, 2] has a length of 2 and the minimum sum greater than 0. So, the answer is 3.
//
// Constraints:
//
// 1 <= nums.length <= 100
//
// 1 <= l <= r <= nums.length
//
// -1000 <= nums[i] <= 1000
// ============================================================================

package leetcode

func minimumSumSubarray(nums []int, l int, r int) int {
    
}
