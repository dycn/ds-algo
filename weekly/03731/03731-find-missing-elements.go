// ============================================================================
// 题号: 3731
// 题目: Find Missing Elements
// 难度: Easy
// 分类: Array, Hash Table, Sorting
// 链接: https://leetcode.cn/problems/find-missing-elements/
// 每日一题日期: 2026-08-04
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你一个整数数组 nums ，数组由若干 互不相同 的整数组成。
//
// 数组 nums 原本包含了某个范围内的 所有整数 。但现在，其中可能 缺失 部分整数。
//
// 该范围内的 最小 整数和 最大 整数仍然存在于 nums 中。
//
// 返回一个 有序 列表，包含该范围内缺失的所有整数，并 按从小到大排序。如果没有缺失的整数，返回一个 空 列表。
//
// 示例 1：
//
// 输入： nums = [1,4,2,5]
//
// 输出： [3]
//
// 解释：
//
// 最小整数为 1，最大整数为 5，因此完整的范围应为 [1,2,3,4,5]。其中只有 3 缺失。
//
// 示例 2：
//
// 输入： nums = [7,8,6,9]
//
// 输出： []
//
// 解释：
//
// 最小整数为 6，最大整数为 9，因此完整的范围为 [6,7,8,9]。所有整数均已存在，因此没有缺失的整数。
//
// 示例 3：
//
// 输入： nums = [5,1]
//
// 输出： [2,3,4]
//
// 解释：
//
// 最小整数为 1，最大整数为 5，因此完整的范围应为 [1,2,3,4,5]。缺失的整数为 2、3 和 4。
//
// 提示：
//
// 2 <= nums.length <= 100
//
// 1 <= nums[i] <= 100
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// You are given an integer array nums consisting of unique integers.
//
// Originally, nums contained every integer within a certain range. However, some integers might have gone missing from the array.
//
// The smallest and largest integers of the original range are still present in nums.
//
// Return a sorted list of all the missing integers in this range. If no integers are missing, return an empty list.
//
// Example 1:
//
// Input: nums = [1,4,2,5]
//
// Output: [3]
//
// Explanation:
//
// The smallest integer is 1 and the largest is 5, so the full range should be [1,2,3,4,5]. Among these, only 3 is missing.
//
// Example 2:
//
// Input: nums = [7,8,6,9]
//
// Output: []
//
// Explanation:
//
// The smallest integer is 6 and the largest is 9, so the full range is [6,7,8,9]. All integers are already present, so no integer is missing.
//
// Example 3:
//
// Input: nums = [5,1]
//
// Output: [2,3,4]
//
// Explanation:
//
// The smallest integer is 1 and the largest is 5, so the full range should be [1,2,3,4,5]. The missing integers are 2, 3, and 4.
//
// Constraints:
//
// 2 <= nums.length <= 100
//
// 1 <= nums[i] <= 100
// ============================================================================

package leetcode

func findMissingElements(nums []int) []int {
    
}
