// ============================================================================
// 题号: 2333
// 题目: Minimum Sum of Squared Difference
// 难度: Medium
// 分类: Greedy, Array, Binary Search, Sorting, Heap (Priority Queue)
// 链接: https://leetcode.cn/problems/minimum-sum-of-squared-difference/
// 每日一题日期: 2026-10-10
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你两个下标从 0 开始的整数数组 nums1 和 nums2 ，长度为 n 。
//
// 数组 nums1 和 nums2 的 差值平方和 定义为所有满足 0 <= i < n 的 (nums1[i] - nums2[i])2 之和。
//
// 同时给你两个正整数 k1 和 k2 。你可以将 nums1 中的任意元素 +1 或者 -1 至多 k1 次。类似的，你可以将 nums2 中的任意元素 +1 或者 -1 至多 k2 次。
//
// 请你返回修改数组 nums1 至多 k1 次且修改数组 nums2 至多 k2 次后的最小 差值平方和 。
//
// 注意：你可以将数组中的元素变成 负 整数。
//
// 示例 1：
//
// 输入：nums1 = [1,2,3,4], nums2 = [2,10,20,19], k1 = 0, k2 = 0
// 输出：579
// 解释：nums1 和 nums2 中的元素不能修改，因为 k1 = 0 和 k2 = 0 。
// 差值平方和为：(1 - 2)2 + (2 - 10)2 + (3 - 20)2 + (4 - 19)2 = 579 。
//
// 示例 2：
//
// 输入：nums1 = [1,4,10,12], nums2 = [5,8,6,9], k1 = 1, k2 = 1
// 输出：43
// 解释：一种得到最小差值平方和的方式为：
// - 将 nums1[0] 增加一次。
// - 将 nums2[2] 增加一次。
// 最小差值平方和为：
// (2 - 5)2 + (4 - 8)2 + (10 - 7)2 + (12 - 9)2 = 43 。
// 注意，也有其他方式可以得到最小差值平方和，但没有得到比 43 更小答案的方案。
//
// 提示：
//
// n == nums1.length == nums2.length
//
// 1 <= n <= 105
//
// 0 <= nums1[i], nums2[i] <= 105
//
// 0 <= k1, k2 <= 109
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// You are given two positive 0-indexed integer arrays nums1 and nums2, both of length n.
//
// The sum of squared difference of arrays nums1 and nums2 is defined as the sum of (nums1[i] - nums2[i])2 for each 0 <= i < n.
//
// You are also given two positive integers k1 and k2. You can modify any of the elements of nums1 by +1 or -1 at most k1 times. Similarly, you can modify any of the elements of nums2 by +1 or -1 at most k2 times.
//
// Return the minimum sum of squared difference after modifying array nums1 at most k1 times and modifying array nums2 at most k2 times.
//
// Note: You are allowed to modify the array elements to become negative integers.
//
// Example 1:
//
// Input: nums1 = [1,2,3,4], nums2 = [2,10,20,19], k1 = 0, k2 = 0
// Output: 579
// Explanation: The elements in nums1 and nums2 cannot be modified because k1 = 0 and k2 = 0.
// The sum of square difference will be: (1 - 2)2 + (2 - 10)2 + (3 - 20)2 + (4 - 19)2 = 579.
//
// Example 2:
//
// Input: nums1 = [1,4,10,12], nums2 = [5,8,6,9], k1 = 1, k2 = 1
// Output: 43
// Explanation: One way to obtain the minimum sum of square difference is:
// - Increase nums1[0] once.
// - Increase nums2[2] once.
// The minimum of the sum of square difference will be:
// (2 - 5)2 + (4 - 8)2 + (10 - 7)2 + (12 - 9)2 = 43.
// Note that, there are other ways to obtain the minimum of the sum of square difference, but there is no way to obtain a sum smaller than 43.
//
// Constraints:
//
// n == nums1.length == nums2.length
//
// 1 <= n <= 105
//
// 0 <= nums1[i], nums2[i] <= 105
//
// 0 <= k1, k2 <= 109
// ============================================================================

package leetcode

func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
    
}
