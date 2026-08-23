// ============================================================================
// 题号: 1386
// 题目: Cinema Seat Allocation
// 难度: Medium
// 分类: Greedy, Bit Manipulation, Array, Hash Table
// 链接: https://leetcode.cn/problems/cinema-seat-allocation/
// 每日一题日期: 2026-08-19
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 如上图所示，电影院的观影厅中有 n 行座位，行编号从 1 到 n ，且每一行内总共有 10 个座位，列编号从 1 到 10 。
//
// 给定一个二维数组 reservedSeats ，其中 reservedSeats[i] = [rowi, seati] 表示第 rowi 行的座位 seati 已经被预定。
//
// 四人小组必须被安排在同一排的四个座位上。该小组可以坐在以下座位块之一：
//
// 座位 2, 3, 4, 5
//
// 座位 4, 5, 6, 7
//
// 座位 6, 7, 8, 9
//
// 只有当该块中的所有座位都 没有 被预订时，才能使用该块。每个座位 最多 只能分配给一个小组。
//
// 返回一个整数，表示可以分配的 最大 四人小组数量。
//
// 示例 1：
//
// 输入：n = 3, reservedSeats = [[1,2],[1,3],[1,8],[2,6],[3,1],[3,10]]
// 输出：4
// 解释：上图所示是最优的安排方案，总共可以安排 4 个家庭。蓝色的叉表示被预约的座位，橙色的连续座位表示一个 4 人家庭。
//
// 示例 2：
//
// 输入：n = 2, reservedSeats = [[2,1],[1,8],[2,6]]
// 输出：2
//
// 示例 3：
//
// 输入：n = 4, reservedSeats = [[4,3],[1,4],[4,6],[1,7]]
// 输出：4
//
// 提示：
//
// 1 <= n <= 109
//
// 1 <= reservedSeats.length <= min(10 * n, 104)
//
// reservedSeats[i] == [rowi, seati]
//
// 1 <= rowi <= n
//
// 1 <= seati <= 10
//
// 所有 reservedSeats[i] 都是互不相同的。
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// A cinema has n rows of seats, numbered from 1 to n. Each row has 10 seats, numbered from 1 to 10.
//
// You are given a 2D integer array reservedSeats, where reservedSeats[i] = [rowi, seati] means that seat seati in row rowi is already reserved.
//
// A four-person group must be assigned to four seats in the same row. The group can be seated in one of the following seat blocks:
//
// seats 2, 3, 4, 5
//
// seats 4, 5, 6, 7
//
// seats 6, 7, 8, 9
//
// A block can be used only if none of its seats are reserved. Each seat can be assigned to at most one group.
//
// Return an integer denoting the maximum number of four-person groups that can be assigned.
//
// Example 1:
//
// Input: n = 3, reservedSeats = [[1,2],[1,3],[1,8],[2,6],[3,1],[3,10]]
// Output: 4
// Explanation: The figure above shows an optimal allocation of four groups. Seats marked in blue are already reserved, and each set of four contiguous seats marked in orange is assigned to one group.
//
// Example 2:
//
// Input: n = 2, reservedSeats = [[2,1],[1,8],[2,6]]
// Output: 2
//
// Example 3:
//
// Input: n = 4, reservedSeats = [[4,3],[1,4],[4,6],[1,7]]
// Output: 4
//
// Constraints:
//
// 1 <= n <= 109
//
// 1 <= reservedSeats.length <= min(10 * n, 104)
//
// reservedSeats[i] == [rowi, seati]
//
// 1 <= rowi <= n
//
// 1 <= seati <= 10
//
// All reservedSeats[i] are distinct.
// ============================================================================

package leetcode

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
    
}
