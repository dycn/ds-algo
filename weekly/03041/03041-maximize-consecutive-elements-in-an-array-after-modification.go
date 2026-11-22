// ============================================================================
// 题号: 3041
// 题目: Maximize Consecutive Elements in an Array After Modification
// 难度: Hard
// 分类: Array, Dynamic Programming, Sorting
// 链接: https://leetcode.cn/problems/maximize-consecutive-elements-in-an-array-after-modification/
// 每日一题日期: 2026-11-20
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你一个下标从 0 开始只包含 正 整数的数组 nums 。
//
// 一开始，你可以将数组中 任意数量 元素增加 至多 1 。
//
// 修改后，你可以从最终数组中选择 一个或者更多 元素，并确保这些元素升序排序后是 连续 的。比方说，[3, 4, 5] 是连续的，但是 [3, 4, 6] 和 [1, 1, 2, 3] 不是连续的。
//
// 请你返回 最多 可以选出的元素数目。
//
// 示例 1：
//
// 输入：nums = [2,1,5,1,1]
// 输出：3
// 解释：我们将下标 0 和 3 处的元素增加 1 ，得到结果数组 nums = [3,1,5,2,1] 。
// 我们选择元素 [3,1,5,2,1] 并将它们排序得到 [1,2,3] ，是连续元素。
// 最多可以得到 3 个连续元素。
//
// 示例 2：
//
// 输入：nums = [1,4,7,10]
// 输出：1
// 解释：我们可以选择的最多元素数目是 1 。
//
// 提示：
//
// 1 <= nums.length <= 105
//
// 1 <= nums[i] <= 106
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// You are given a 0-indexed array nums consisting of positive integers.
//
// Initially, you can increase the value of any element in the array by at most 1.
//
// After that, you need to select one or more elements from the final array such that those elements are consecutive when sorted in increasing order. For example, the elements [3, 4, 5] are consecutive while [3, 4, 6] and [1, 1, 2, 3] are not.
//
// Return the maximum number of elements that you can select.
//
// Example 1:
//
// Input: nums = [2,1,5,1,1]
// Output: 3
// Explanation: We can increase the elements at indices 0 and 3. The resulting array is nums = [3,1,5,2,1].
// We select the elements [3,1,5,2,1] and we sort them to obtain [1,2,3], which are consecutive.
// It can be shown that we cannot select more than 3 consecutive elements.
//
// Example 2:
//
// Input: nums = [1,4,7,10]
// Output: 1
// Explanation: The maximum consecutive elements that we can select is 1.
//
// Constraints:
//
// 1 <= nums.length <= 105
//
// 1 <= nums[i] <= 106
// ============================================================================

package leetcode

func maxSelectedElements(nums []int) int {
    
}
