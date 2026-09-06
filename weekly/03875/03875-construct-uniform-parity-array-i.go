// ============================================================================
// 题号: 3875
// 题目: Construct Uniform Parity Array I
// 难度: Easy
// 分类: Array, Math
// 链接: https://leetcode.cn/problems/construct-uniform-parity-array-i/
// 每日一题日期: 2026-09-02
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你一个长度为 n 的数组 nums1，其中包含 互不相同 的整数。
//
// 你需要构造另一个长度为 n 的数组 nums2，使得 nums2 中的元素要么全部为 奇数，要么全部为 偶数。
//
// 对于每个下标 i，你必须从以下两种选择中 任选其一（顺序不限）：
//
// nums2[i] = nums1[i]
//
// nums2[i] = nums1[i] - nums1[j]，其中 j != i
//
// 如果能够构造出满足条件的数组，则返回 true；否则，返回 false。
//
// 示例 1：
//
// 输入： nums1 = [2,3]
//
// 输出： true
//
// 解释：
//
// 选择 nums2[0] = nums1[0] - nums1[1] = 2 - 3 = -1。
//
// 选择 nums2[1] = nums1[1] = 3。
//
// nums2 = [-1, 3]，两个元素均为奇数。因此答案为 true。
//
// 示例 2：
//
// 输入： nums1 = [4,6]
//
// 输出： true
//
// 解释：​​​​​​​
//
// 选择 nums2[0] = nums1[0] = 4。
//
// 选择 nums2[1] = nums1[1] = 6。
//
// nums2 = [4, 6]，两个元素均为偶数。因此答案为 true。
//
// 提示：
//
// 1 <= n == nums1.length <= 100
//
// 1 <= nums1[i] <= 100
//
// nums1 中的所有整数互不相同。
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// You are given an array nums1 of n distinct integers.
//
// You want to construct another array nums2 of length n such that the elements in nums2 are either all odd or all even.
//
// For each index i, you must choose exactly one of the following (in any order):
//
// nums2[i] = nums1[i]
//
// nums2[i] = nums1[i] - nums1[j], for an index j != i
//
// Return true if it is possible to construct such an array, otherwise, return false.
//
// Example 1:
//
// Input: nums1 = [2,3]
//
// Output: true
//
// Explanation:
//
// Choose nums2[0] = nums1[0] - nums1[1] = 2 - 3 = -1.
//
// Choose nums2[1] = nums1[1] = 3.
//
// nums2 = [-1, 3], and both elements are odd. Thus, the answer is true​​​​​​​.
//
// Example 2:
//
// Input: nums1 = [4,6]
//
// Output: true
//
// Explanation:​​​​​​​
//
// Choose nums2[0] = nums1[0] = 4.
//
// Choose nums2[1] = nums1[1] = 6.
//
// nums2 = [4, 6], and all elements are even. Thus, the answer is true.
//
// Constraints:
//
// 1 <= n == nums1.length <= 100
//
// 1 <= nums1[i] <= 100
//
// nums1 consists of distinct integers.
// ============================================================================

package leetcode

func uniformArray(nums1 []int) bool {
    
}
