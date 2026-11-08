// ============================================================================
// 题号: 3736
// 题目: Minimum Moves to Equal Array Elements III
// 难度: Easy
// 分类: Array, Math
// 链接: https://leetcode.cn/problems/minimum-moves-to-equal-array-elements-iii/
// 每日一题日期: 2026-11-03
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你一个整数数组 nums。
//
// 在一步操作中，你可以将任意单个元素 nums[i] 的值 增加 1。
//
// 返回使数组中的所有元素都 相等 所需的 最小总操作次数 。
//
// 示例 1:
//
// 输入: nums = [2,1,3]
//
// 输出: 3
//
// 解释:
//
// 使所有元素相等的操作如下:
//
// 将 nums[0] = 2 增加 1, 变为 3。
//
// 将 nums[1] = 1 增加 1, 变为 2。
//
// 将 nums[1] = 2 增加 1, 变为 3。
//
// 现在，nums 中的所有元素都等于 3。最小总操作次数为 3。
//
// 示例 2:
//
// 输入: nums = [4,4,5]
//
// 输出: 2
//
// 解释:
//
// 使所有元素相等的操作如下:
//
// 将 nums[0] = 4 增加 1, 变为 5。
//
// 将 nums[1] = 4 增加 1, 变为 5。
//
// 现在，nums 中的所有元素都等于 5。最小总操作次数为 2。
//
// 提示:
//
// 1 <= nums.length <= 100
//
// 1 <= nums[i] <= 100
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// You are given an integer array nums.
//
// In one move, you may increase the value of any single element nums[i] by 1.
//
// Return the minimum total number of moves required so that all elements in nums become equal.
//
// Example 1:
//
// Input: nums = [2,1,3]
//
// Output: 3
//
// Explanation:
//
// To make all elements equal:
//
// Increase nums[0] = 2 by 1 to make it 3.
//
// Increase nums[1] = 1 by 1 to make it 2.
//
// Increase the updated nums[1] = 2 by 1 to make it 3.
//
// Now, all elements of nums are equal to 3. The minimum total moves is 3.
//
// Example 2:
//
// Input: nums = [4,4,5]
//
// Output: 2
//
// Explanation:
//
// To make all elements equal:
//
// Increase nums[0] = 4 by 1 to make it 5.
//
// Increase nums[1] = 4 by 1 to make it 5.
//
// Now, all elements of nums are equal to 5. The minimum total moves is 2.
//
// Constraints:
//
// 1 <= nums.length <= 100
//
// 1 <= nums[i] <= 100
// ============================================================================

package leetcode

func minMoves(nums []int) int {
    
}
