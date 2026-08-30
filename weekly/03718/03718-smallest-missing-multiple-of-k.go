// ============================================================================
// 题号: 3718
// 题目: Smallest Missing Multiple of K
// 难度: Easy
// 分类: Array, Hash Table
// 链接: https://leetcode.cn/problems/smallest-missing-multiple-of-k/
// 每日一题日期: 2026-08-25
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你一个整数数组 nums 和一个整数 k，请返回从 nums 中缺失的、最小的正整数 k 的倍数。
//
// 倍数 指能被 k 整除的任意正整数。
//
// 示例 1：
//
// 输入： nums = [8,2,3,4,6], k = 2
//
// 输出： 10
//
// 解释：
//
// 当 k = 2 时，其倍数为 2、4、6、8、10、12……，其中在 nums 中缺失的最小倍数是 10。
//
// 示例 2：
//
// 输入： nums = [1,4,7,10,15], k = 5
//
// 输出： 5
//
// 解释：
//
// 当 k = 5 时，其倍数为 5、10、15、20……，其中在 nums 中缺失的最小倍数是 5。
//
// 提示：
//
// 1 <= nums.length <= 100
//
// 1 <= nums[i] <= 100
//
// 1 <= k <= 100
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// Given an integer array nums and an integer k, return the smallest positive multiple of k that is missing from nums.
//
// A multiple of k is any positive integer divisible by k.
//
// Example 1:
//
// Input: nums = [8,2,3,4,6], k = 2
//
// Output: 10
//
// Explanation:
//
// The multiples of k = 2 are 2, 4, 6, 8, 10, 12... and the smallest multiple missing from nums is 10.
//
// Example 2:
//
// Input: nums = [1,4,7,10,15], k = 5
//
// Output: 5
//
// Explanation:
//
// The multiples of k = 5 are 5, 10, 15, 20... and the smallest multiple missing from nums is 5.
//
// Constraints:
//
// 1 <= nums.length <= 100
//
// 1 <= nums[i] <= 100
//
// 1 <= k <= 100
// ============================================================================

package leetcode

func missingMultiple(nums []int, k int) int {
    
}
