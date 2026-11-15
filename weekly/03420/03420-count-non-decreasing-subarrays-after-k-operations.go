// ============================================================================
// 题号: 3420
// 题目: Count Non-Decreasing Subarrays After K Operations
// 难度: Hard
// 分类: Stack, Segment Tree, Queue, Array, Sliding Window, Monotonic Queue, Monotonic Stack
// 链接: https://leetcode.cn/problems/count-non-decreasing-subarrays-after-k-operations/
// 每日一题日期: 2026-11-13
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你一个长度为 n 的数组 nums 和一个整数 k 。
//
// 对于 nums 中的每一个子数组，你可以对它进行 至多 k 次操作。每次操作中，你可以将子数组中的任意一个元素增加 1 。
//
// 注意 ，每个子数组都是独立的，也就是说你对一个子数组的修改不会保留到另一个子数组中。
//
// 请你返回最多 k 次操作以内，有多少个子数组可以变成 非递减 的。
//
// 如果一个数组中的每一个元素都大于等于前一个元素（如果前一个元素存在），那么我们称这个数组是 非递减 的。
//
// 示例 1：
//
// 输入：nums = [6,3,1,2,4,4], k = 7
//
// 输出：17
//
// 解释：
//
// nums 的所有 21 个子数组中，只有子数组 [6, 3, 1] ，[6, 3, 1, 2] ，[6, 3, 1, 2, 4] 和 [6, 3, 1, 2, 4, 4] 无法在 k = 7 次操作以内变为非递减的。所以非递减子数组的数目为 21 - 4 = 17 。
//
// 示例 2：
//
// 输入：nums = [6,3,1,3,6], k = 4
//
// 输出：12
//
// 解释：
//
// 子数组 [3, 1, 3, 6] 和 nums 中所有小于等于三个元素的子数组中，除了 [6, 3, 1] 以外，都可以在 k 次操作以内变为非递减子数组。总共有 5 个包含单个元素的子数组，4 个包含两个元素的子数组，除 [6, 3, 1] 以外有 2 个包含三个元素的子数组，所以总共有 1 + 5 + 4 + 2 = 12 个子数组可以变为非递减的。
//
// 提示：
//
// 1 <= nums.length <= 105
//
// 1 <= nums[i] <= 109
//
// 1 <= k <= 109
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// You are given an array nums of n integers and an integer k.
//
// For each subarray of nums, you can apply up to k operations on it. In each operation, you increment any element of the subarray by 1.
//
// Note that each subarray is considered independently, meaning changes made to one subarray do not persist to another.
//
// Return the number of subarrays that you can make non-decreasing ​​​​​after performing at most k operations.
//
// An array is said to be non-decreasing if each element is greater than or equal to its previous element, if it exists.
//
// Example 1:
//
// Input: nums = [6,3,1,2,4,4], k = 7
//
// Output: 17
//
// Explanation:
//
// Out of all 21 possible subarrays of nums, only the subarrays [6, 3, 1], [6, 3, 1, 2], [6, 3, 1, 2, 4] and [6, 3, 1, 2, 4, 4] cannot be made non-decreasing after applying up to k = 7 operations. Thus, the number of non-decreasing subarrays is 21 - 4 = 17.
//
// Example 2:
//
// Input: nums = [6,3,1,3,6], k = 4
//
// Output: 12
//
// Explanation:
//
// The subarray [3, 1, 3, 6] along with all subarrays of nums with three or fewer elements, except [6, 3, 1], can be made non-decreasing after k operations. There are 5 subarrays of a single element, 4 subarrays of two elements, and 2 subarrays of three elements except [6, 3, 1], so there are 1 + 5 + 4 + 2 = 12 subarrays that can be made non-decreasing.
//
// Constraints:
//
// 1 <= nums.length <= 105
//
// 1 <= nums[i] <= 109
//
// 1 <= k <= 109
// ============================================================================

package leetcode

func countNonDecreasingSubarrays(nums []int, k int) int64 {
    
}
