// ============================================================================
// 题号: 2058
// 题目: Find the Minimum and Maximum Number of Nodes Between Critical Points
// 难度: Medium
// 分类: Linked List
// 链接: https://leetcode.cn/problems/find-the-minimum-and-maximum-number-of-nodes-between-critical-points/
// 每日一题日期: 2026-08-31
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 链表中的 临界点 定义为一个 局部极大值点 或 局部极小值点 。
//
// 如果当前节点的值 严格大于 前一个节点和后一个节点，那么这个节点就是一个  局部极大值点 。
//
// 如果当前节点的值 严格小于 前一个节点和后一个节点，那么这个节点就是一个  局部极小值点 。
//
// 注意：节点只有在同时存在前一个节点和后一个节点的情况下，才能成为一个 局部极大值点 / 极小值点 。
//
// 给你一个链表 head ，返回一个长度为 2 的数组 [minDistance, maxDistance] ，其中 minDistance 是任意两个不同临界点之间的最小距离，maxDistance 是任意两个不同临界点之间的最大距离。如果临界点少于两个，则返回 [-1，-1] 。
//
// 示例 1：
//
// 输入：head = [3,1]
// 输出：[-1,-1]
// 解释：链表 [3,1] 中不存在临界点。
//
// 示例 2：
//
// 输入：head = [5,3,1,2,5,1,2]
// 输出：[1,3]
// 解释：存在三个临界点：
// - [5,3,1,2,5,1,2]：第三个节点是一个局部极小值点，因为 1 比 3 和 2 小。
// - [5,3,1,2,5,1,2]：第五个节点是一个局部极大值点，因为 5 比 2 和 1 大。
// - [5,3,1,2,5,1,2]：第六个节点是一个局部极小值点，因为 1 比 5 和 2 小。
// 第五个节点和第六个节点之间距离最小。minDistance = 6 - 5 = 1 。
// 第三个节点和第六个节点之间距离最大。maxDistance = 6 - 3 = 3 。
//
// 示例 3：
//
// 输入：head = [1,3,2,2,3,2,2,2,7]
// 输出：[3,3]
// 解释：存在两个临界点：
// - [1,3,2,2,3,2,2,2,7]：第二个节点是一个局部极大值点，因为 3 比 1 和 2 大。
// - [1,3,2,2,3,2,2,2,7]：第五个节点是一个局部极大值点，因为 3 比 2 和 2 大。
// 最小和最大距离都存在于第二个节点和第五个节点之间。
// 因此，minDistance 和 maxDistance 是 5 - 2 = 3 。
// 注意，最后一个节点不算一个局部极大值点，因为它之后就没有节点了。
//
// 示例 4：
//
// 输入：head = [2,3,3,2]
// 输出：[-1,-1]
// 解释：链表 [2,3,3,2] 中不存在临界点。
//
// 提示：
//
// 链表中节点的数量在范围 [2, 105] 内
//
// 1 <= Node.val <= 105
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// A critical point in a linked list is defined as either a local maxima or a local minima.
//
// A node is a local maxima if the current node has a value strictly greater than the previous node and the next node.
//
// A node is a local minima if the current node has a value strictly smaller than the previous node and the next node.
//
// Note that a node can only be a local maxima/minima if there exists both a previous node and a next node.
//
// Given a linked list head, return an array of length 2 containing [minDistance, maxDistance] where minDistance is the minimum distance between any two distinct critical points and maxDistance is the maximum distance between any two distinct critical points. If there are fewer than two critical points, return [-1, -1].
//
// Example 1:
//
// Input: head = [3,1]
// Output: [-1,-1]
// Explanation: There are no critical points in [3,1].
//
// Example 2:
//
// Input: head = [5,3,1,2,5,1,2]
// Output: [1,3]
// Explanation: There are three critical points:
// - [5,3,1,2,5,1,2]: The third node is a local minima because 1 is less than 3 and 2.
// - [5,3,1,2,5,1,2]: The fifth node is a local maxima because 5 is greater than 2 and 1.
// - [5,3,1,2,5,1,2]: The sixth node is a local minima because 1 is less than 5 and 2.
// The minimum distance is between the fifth and the sixth node. minDistance = 6 - 5 = 1.
// The maximum distance is between the third and the sixth node. maxDistance = 6 - 3 = 3.
//
// Example 3:
//
// Input: head = [1,3,2,2,3,2,2,2,7]
// Output: [3,3]
// Explanation: There are two critical points:
// - [1,3,2,2,3,2,2,2,7]: The second node is a local maxima because 3 is greater than 1 and 2.
// - [1,3,2,2,3,2,2,2,7]: The fifth node is a local maxima because 3 is greater than 2 and 2.
// Both the minimum and maximum distances are between the second and the fifth node.
// Thus, minDistance and maxDistance is 5 - 2 = 3.
// Note that the last node is not considered a local maxima because it does not have a next node.
//
// Constraints:
//
// The number of nodes in the list is in the range [2, 105].
//
// 1 <= Node.val <= 105
// ============================================================================

package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func nodesBetweenCriticalPoints(head *ListNode) []int {
    
}
