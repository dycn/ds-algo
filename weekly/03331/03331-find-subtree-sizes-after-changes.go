// ============================================================================
// 题号: 3331
// 题目: Find Subtree Sizes After Changes
// 难度: Medium
// 分类: Tree, Depth-First Search, Array, Hash Table, String
// 链接: https://leetcode.cn/problems/find-subtree-sizes-after-changes/
// 每日一题日期: 2026-11-21
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你一棵 n 个节点且根节点为编号 0 的树，节点编号为 0 到 n - 1 。这棵树用一个长度为 n 的数组 parent 表示，其中 parent[i] 是第 i 个节点的父亲节点的编号。由于节点 0 是根，parent[0] == -1 。
//
// 给你一个长度为 n 的字符串 s ，其中 s[i] 是节点 i 对应的字符。
//
// 对于节点编号从 1 到 n - 1 的每个节点 x ，我们 同时 执行以下操作 一次 ：
//
// 找到距离节点 x 最近 的祖先节点 y ，且 s[x] == s[y] 。
//
// 如果节点 y 不存在，那么不做任何修改。
//
// 否则，将节点 x 与它父亲节点之间的边 删除 ，在 x 与 y 之间连接一条边，使 y 变为 x 新的父节点。
//
// 请你返回一个长度为 n 的数组 answer ，其中 answer[i] 是 最终 树中，节点 i 为根的 子树 的 大小 。
//
// 示例 1：
//
// 输入：parent = [-1,0,0,1,1,1], s = "abaabc"
//
// 输出：[6,3,1,1,1,1]
//
// 解释：
//
// 节点 3 的父节点从节点 1 变为节点 0 。
//
// 示例 2：
//
// 输入：parent = [-1,0,4,0,1], s = "abbba"
//
// 输出：[5,2,1,1,1]
//
// 解释：
//
// 以下变化会同时发生：
//
// 节点 4 的父节点从节点 1 变为节点 0 。
//
// 节点 2 的父节点从节点 4 变为节点 1 。
//
// 提示：
//
// n == parent.length == s.length
//
// 1 <= n <= 105
//
// 对于所有的 i >= 1 ，都有 0 <= parent[i] <= n - 1 。
//
// parent[0] == -1
//
// parent 表示一棵合法的树。
//
// s 只包含小写英文字母。
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// You are given a tree rooted at node 0 that consists of n nodes numbered from 0 to n - 1. The tree is represented by an array parent of size n, where parent[i] is the parent of node i. Since node 0 is the root, parent[0] == -1.
//
// You are also given a string s of length n, where s[i] is the character assigned to node i.
//
// We make the following changes on the tree one time simultaneously for all nodes x from 1 to n - 1:
//
// Find the closest node y to node x such that y is an ancestor of x, and s[x] == s[y].
//
// If node y does not exist, do nothing.
//
// Otherwise, remove the edge between x and its current parent and make node y the new parent of x by adding an edge between them.
//
// Return an array answer of size n where answer[i] is the size of the subtree rooted at node i in the final tree.
//
// Example 1:
//
// Input: parent = [-1,0,0,1,1,1], s = "abaabc"
//
// Output: [6,3,1,1,1,1]
//
// Explanation:
//
// The parent of node 3 will change from node 1 to node 0.
//
// Example 2:
//
// Input: parent = [-1,0,4,0,1], s = "abbba"
//
// Output: [5,2,1,1,1]
//
// Explanation:
//
// The following changes will happen at the same time:
//
// The parent of node 4 will change from node 1 to node 0.
//
// The parent of node 2 will change from node 4 to node 1.
//
// Constraints:
//
// n == parent.length == s.length
//
// 1 <= n <= 105
//
// 0 <= parent[i] <= n - 1 for all i >= 1.
//
// parent[0] == -1
//
// parent represents a valid tree.
//
// s consists only of lowercase English letters.
// ============================================================================

package leetcode

func findSubtreeSizes(parent []int, s string) []int {
    
}
