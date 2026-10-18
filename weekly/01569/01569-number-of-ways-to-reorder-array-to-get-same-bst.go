// ============================================================================
// 题号: 1569
// 题目: Number of Ways to Reorder Array to Get Same BST
// 难度: Hard
// 分类: Tree, Union Find, Binary Search Tree, Memoization, Array, Math, Divide and Conquer, Dynamic Programming, Binary Tree, Combinatorics, 费马小定理
// 链接: https://leetcode.cn/problems/number-of-ways-to-reorder-array-to-get-same-bst/
// 每日一题日期: 2026-10-18
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你一个数组 nums 表示 1 到 n 的一个排列。我们按照元素在 nums 中的顺序依次插入一个初始为空的二叉搜索树（BST）。请你统计将 nums 重新排序后，统计满足如下条件的方案数：重排后得到的二叉搜索树与 nums 原本数字顺序得到的二叉搜索树相同。
//
// 比方说，给你 nums = [2,1,3]，我们得到一棵 2 为根，1 为左孩子，3 为右孩子的树。数组 [2,3,1] 也能得到相同的 BST，但 [3,2,1] 会得到一棵不同的 BST 。
//
// 请你返回重排 nums 后，与原数组 nums 得到相同二叉搜索树的方案数。
//
// 由于答案可能会很大，请将结果对 109 + 7 取余数。
//
// 示例 1：
//
// 输入：nums = [2,1,3]
// 输出：1
// 解释：我们将 nums 重排， [2,3,1] 能得到相同的 BST 。没有其他得到相同 BST 的方案了。
//
// 示例 2：
//
// 输入：nums = [3,4,5,1,2]
// 输出：5
// 解释：下面 5 个数组会得到相同的 BST：
// [3,1,2,4,5]
// [3,1,4,2,5]
// [3,1,4,5,2]
// [3,4,1,2,5]
// [3,4,1,5,2]
//
// 示例 3：
//
// 输入：nums = [1,2,3]
// 输出：0
// 解释：没有别的排列顺序能得到相同的 BST 。
//
// 提示：
//
// 1 <= nums.length <= 1000
//
// 1 <= nums[i] <= nums.length
//
// nums 中所有数 互不相同 。
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// Given an array nums that represents a permutation of integers from 1 to n. We are going to construct a binary search tree (BST) by inserting the elements of nums in order into an initially empty BST. Find the number of different ways to reorder nums so that the constructed BST is identical to that formed from the original array nums.
//
// For example, given nums = [2,1,3], we will have 2 as the root, 1 as a left child, and 3 as a right child. The array [2,3,1] also yields the same BST but [3,2,1] yields a different BST.
//
// Return the number of ways to reorder nums such that the BST formed is identical to the original BST formed from nums.
//
// Since the answer may be very large, return it modulo 109 + 7.
//
// Example 1:
//
// Input: nums = [2,1,3]
// Output: 1
// Explanation: We can reorder nums to be [2,3,1] which will yield the same BST. There are no other ways to reorder nums which will yield the same BST.
//
// Example 2:
//
// Input: nums = [3,4,5,1,2]
// Output: 5
// Explanation: The following 5 arrays will yield the same BST:
// [3,1,2,4,5]
// [3,1,4,2,5]
// [3,1,4,5,2]
// [3,4,1,2,5]
// [3,4,1,5,2]
//
// Example 3:
//
// Input: nums = [1,2,3]
// Output: 0
// Explanation: There are no other orderings of nums that will yield the same BST.
//
// Constraints:
//
// 1 <= nums.length <= 1000
//
// 1 <= nums[i] <= nums.length
//
// All integers in nums are distinct.
// ============================================================================

package leetcode

func numOfWays(nums []int) int {
    
}
