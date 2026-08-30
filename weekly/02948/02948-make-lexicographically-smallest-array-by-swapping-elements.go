// ============================================================================
// 题号: 2948
// 题目: Make Lexicographically Smallest Array by Swapping Elements
// 难度: Medium
// 分类: Union Find, Array, Sorting
// 链接: https://leetcode.cn/problems/make-lexicographically-smallest-array-by-swapping-elements/
// 每日一题日期: 2026-08-29
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你一个下标从 0 开始的 正整数 数组 nums 和一个 正整数 limit 。
//
// 在一次操作中，你可以选择任意两个下标 i 和 j，如果 满足 |nums[i] - nums[j]| <= limit ，则交换 nums[i] 和 nums[j] 。
//
// 返回执行任意次操作后能得到的 字典序最小的数组 。
//
// 如果在数组 a 和数组 b 第一个不同的位置上，数组 a 中的对应元素比数组 b 中的对应元素的字典序更小，则认为数组 a 就比数组 b 字典序更小。例如，数组 [2,10,3] 比数组 [10,2,3] 字典序更小，下标 0 处是两个数组第一个不同的位置，且 2 < 10 。
//
// 示例 1：
//
// 输入：nums = [1,5,3,9,8], limit = 2
// 输出：[1,3,5,8,9]
// 解释：执行 2 次操作：
// - 交换 nums[1] 和 nums[2] 。数组变为 [1,3,5,9,8] 。
// - 交换 nums[3] 和 nums[4] 。数组变为 [1,3,5,8,9] 。
// 即便执行更多次操作，也无法得到字典序更小的数组。
// 注意，执行不同的操作也可能会得到相同的结果。
//
// 示例 2：
//
// 输入：nums = [1,7,6,18,2,1], limit = 3
// 输出：[1,6,7,18,1,2]
// 解释：执行 3 次操作：
// - 交换 nums[1] 和 nums[2] 。数组变为 [1,6,7,18,2,1] 。
// - 交换 nums[0] 和 nums[4] 。数组变为 [2,6,7,18,1,1] 。
// - 交换 nums[0] 和 nums[5] 。数组变为 [1,6,7,18,1,2] 。
// 即便执行更多次操作，也无法得到字典序更小的数组。
//
// 示例 3：
//
// 输入：nums = [1,7,28,19,10], limit = 3
// 输出：[1,7,28,19,10]
// 解释：[1,7,28,19,10] 是字典序最小的数组，因为不管怎么选择下标都无法执行操作。
//
// 提示：
//
// 1 <= nums.length <= 105
//
// 1 <= nums[i] <= 109
//
// 1 <= limit <= 109
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// You are given a 0-indexed array of positive integers nums and a positive integer limit.
//
// In one operation, you can choose any two indices i and j and swap nums[i] and nums[j] if |nums[i] - nums[j]| <= limit.
//
// Return the lexicographically smallest array that can be obtained by performing the operation any number of times.
//
// An array a is lexicographically smaller than an array b if in the first position where a and b differ, array a has an element that is less than the corresponding element in b. For example, the array [2,10,3] is lexicographically smaller than the array [10,2,3] because they differ at index 0 and 2 < 10.
//
// Example 1:
//
// Input: nums = [1,5,3,9,8], limit = 2
// Output: [1,3,5,8,9]
// Explanation: Apply the operation 2 times:
// - Swap nums[1] with nums[2]. The array becomes [1,3,5,9,8]
// - Swap nums[3] with nums[4]. The array becomes [1,3,5,8,9]
// We cannot obtain a lexicographically smaller array by applying any more operations.
// Note that it may be possible to get the same result by doing different operations.
//
// Example 2:
//
// Input: nums = [1,7,6,18,2,1], limit = 3
// Output: [1,6,7,18,1,2]
// Explanation: Apply the operation 3 times:
// - Swap nums[1] with nums[2]. The array becomes [1,6,7,18,2,1]
// - Swap nums[0] with nums[4]. The array becomes [2,6,7,18,1,1]
// - Swap nums[0] with nums[5]. The array becomes [1,6,7,18,1,2]
// We cannot obtain a lexicographically smaller array by applying any more operations.
//
// Example 3:
//
// Input: nums = [1,7,28,19,10], limit = 3
// Output: [1,7,28,19,10]
// Explanation: [1,7,28,19,10] is the lexicographically smallest array we can obtain because we cannot apply the operation on any two indices.
//
// Constraints:
//
// 1 <= nums.length <= 105
//
// 1 <= nums[i] <= 109
//
// 1 <= limit <= 109
// ============================================================================

package leetcode

func lexicographicallySmallestArray(nums []int, limit int) []int {
    
}
