// ============================================================================
// 题号: 3963
// 题目: Create Grid With Exactly One Path
// 难度: Easy
// 分类: Uncategorized
// 链接: https://leetcode.cn/problems/create-grid-with-exactly-one-path/
// 每日一题日期: 2026-10-27
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你两个整数 m 和 n，分别表示网格的行数和列数。
//
// 请你构造 任意 一个只包含字符 '.' 和 '#' 的 m x n 网格，其中：
//
// '.' 表示空单元格。
//
// '#' 表示障碍物单元格。
//
// 有效路径 是满足以下条件的空单元格序列：
//
// 从左上角单元格 (0, 0) 开始。
//
// 在右下角单元格 (m - 1, n - 1) 结束。
//
// 只能向：
//
// 右移动，从 (i, j) 到 (i, j + 1)，或者
//
// 下移动，从 (i, j) 到 (i + 1, j)。
//
// 返回任意一个从左上角到右下角 恰好只有一条有效路径 的网格。
//
// 示例 1：
//
// 输入： m = 2, n = 3
//
// 输出： ["..#","#.."]
//
// 解释：
//
// 唯一的有效路径是：(0,0) → (0,1) → (1,1) → (1,2)
//
// 示例 2：
//
// 输入： m = 3, n = 3
//
// 输出： ["..#","#..","##."]
//
// 解释：
//
// 唯一的有效路径是：(0,0) → (0,1) → (1,1) → (1,2) → (2,2)
//
// 示例 3：
//
// 输入： m = 1, n = 4
//
// 输出： ["...."]
//
// 解释：
//
// 唯一的有效路径是：(0,0) → (0,1) → (0,2) → (0,3)
//
// 提示：
//
// 1 <= m, n <= 25
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// You are given two integers m and n, representing the number of rows and columns of a grid.
//
// Construct any m x n grid consisting only of the characters '.' and '#', where:
//
// '.' represents a free cell.
//
// '#' represents an obstacle cell.
//
// A valid path is a sequence of free cells that:
//
// Starts at the top-left cell (0, 0).
//
// Ends at the bottom-right cell (m - 1, n - 1).
//
// Moves only:
//
// Right, from (i, j) to (i, j + 1), or
//
// Down, from (i, j) to (i + 1, j).
//
// Return any grid such that there is exactly one valid path from the top-left cell to the bottom-right cell.
//
// Example 1:
//
// Input: m = 2, n = 3
//
// Output: ["..#","#.."]
//
// Explanation:
//
// The only valid path is: (0,0) → (0,1) → (1,1) → (1,2)
//
// Example 2:
//
// Input: m = 3, n = 3
//
// Output: ["..#","#..","##."]
//
// Explanation:
//
// The only valid path is: (0,0) → (0,1) → (1,1) → (1,2) → (2,2)
//
// Example 3:
//
// Input: m = 1, n = 4
//
// Output: ["...."]
//
// Explanation:
//
// The only valid path is: (0,0) → (0,1) → (0,2) → (0,3)
//
// Constraints:
//
// 1 <= m, n <= 25
// ============================================================================

package leetcode

func createGrid(m int, n int) []string {
    
}
