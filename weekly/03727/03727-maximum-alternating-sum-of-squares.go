// ============================================================================
// 题号: 3727
// 题目: Maximum Alternating Sum of Squares
// 难度: Medium
// 分类: Greedy, Array, Sorting
// 链接: https://leetcode.cn/problems/maximum-alternating-sum-of-squares/
// 每日一题日期: 2026-10-13
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你一个整数数组 nums。你可以以任意顺序 重新排列元素 。
//
// 数组 arr 的 交替得分 定义为：
//
// score = arr[0]2 - arr[1]2 + arr[2]2 - arr[3]2 + ...
//
// 在对 nums 重新排列后，返回其 最大可能的交替得分。
//
// 示例 1：
//
// 输入： nums = [1,2,3]
//
// 输出： 12
//
// 解释：
//
// nums 的一种可行重排为 [2,1,3]，该排列在所有可能重排中给出了最大交替得分。
//
// 交替得分计算如下：
//
// score = 22 - 12 + 32 = 4 - 1 + 9 = 12
//
// 示例 2：
//
// 输入： nums = [1,-1,2,-2,3,-3]
//
// 输出： 16
//
// 解释：
//
// nums 的一种可行重排为 [-3,-1,-2,1,3,2]，该排列在所有可能重排中给出了最大交替得分。
//
// 交替得分计算如下：
//
// score = (-3)2 - (-1)2 + (-2)2 - (1)2 + (3)2 - (2)2 = 9 - 1 + 4 - 1 + 9 - 4 = 16
//
// 提示：
//
// 1 <= nums.length <= 105
//
// -4 * 104 <= nums[i] <= 4 * 104
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// You are given an integer array nums. You may rearrange the elements in any order.
//
// The alternating score of an array arr is defined as:
//
// score = arr[0]2 - arr[1]2 + arr[2]2 - arr[3]2 + ...
//
// Return an integer denoting the maximum possible alternating score of nums after rearranging its elements.
//
// Example 1:
//
// Input: nums = [1,2,3]
//
// Output: 12
//
// Explanation:
//
// A possible rearrangement for nums is [2,1,3], which gives the maximum alternating score among all possible rearrangements.
//
// The alternating score is calculated as:
//
// score = 22 - 12 + 32 = 4 - 1 + 9 = 12
//
// Example 2:
//
// Input: nums = [1,-1,2,-2,3,-3]
//
// Output: 16
//
// Explanation:
//
// A possible rearrangement for nums is [-3,-1,-2,1,3,2], which gives the maximum alternating score among all possible rearrangements.
//
// The alternating score is calculated as:
//
// score = (-3)2 - (-1)2 + (-2)2 - (1)2 + (3)2 - (2)2 = 9 - 1 + 4 - 1 + 9 - 4 = 16
//
// Constraints:
//
// 1 <= nums.length <= 105
//
// -4 * 104 <= nums[i] <= 4 * 104
// ============================================================================

package leetcode

func maxAlternatingSum(nums []int) int64 {
    
}
