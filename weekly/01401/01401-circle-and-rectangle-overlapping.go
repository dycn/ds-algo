// ============================================================================
// 题号: 1401
// 题目: Circle and Rectangle Overlapping
// 难度: Medium
// 分类: Geometry, Math
// 链接: https://leetcode.cn/problems/circle-and-rectangle-overlapping/
// 每日一题日期: 2026-09-19
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你一个以 (radius, xCenter, yCenter) 表示的圆和一个与坐标轴平行的矩形 (x1, y1, x2, y2) ，其中 (x1, y1) 是矩形左下角的坐标，而 (x2, y2) 是右上角的坐标。
//
// 如果圆和矩形有重叠的部分，请你返回 true ，否则返回 false 。
//
// 换句话说，请你检测是否 存在 点 (xi, yi) ，它既在圆上也在矩形上（两者都包括点落在边界上的情况）。
//
// 示例 1 ：
//
// 输入：radius = 1, xCenter = 0, yCenter = 0, x1 = 1, y1 = -1, x2 = 3, y2 = 1
// 输出：true
// 解释：圆和矩形存在公共点 (1,0) 。
//
// 示例 2 ：
//
// 输入：radius = 1, xCenter = 1, yCenter = 1, x1 = 1, y1 = -3, x2 = 2, y2 = -1
// 输出：false
//
// 示例 3 ：
//
// 输入：radius = 1, xCenter = 0, yCenter = 0, x1 = -1, y1 = 0, x2 = 0, y2 = 1
// 输出：true
//
// 提示：
//
// 1 <= radius <= 2000
//
// -104 <= xCenter, yCenter <= 104
//
// -104 <= x1 < x2 <= 104
//
// -104 <= y1 < y2 <= 104
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// You are given a circle represented as (radius, xCenter, yCenter) and an axis-aligned rectangle represented as (x1, y1, x2, y2), where (x1, y1) are the coordinates of the bottom-left corner, and (x2, y2) are the coordinates of the top-right corner of the rectangle.
//
// Return true if the circle and rectangle are overlapped otherwise return false. In other words, check if there is any point (xi, yi) that belongs to the circle and the rectangle at the same time.
//
// Example 1:
//
// Input: radius = 1, xCenter = 0, yCenter = 0, x1 = 1, y1 = -1, x2 = 3, y2 = 1
// Output: true
// Explanation: Circle and rectangle share the point (1,0).
//
// Example 2:
//
// Input: radius = 1, xCenter = 1, yCenter = 1, x1 = 1, y1 = -3, x2 = 2, y2 = -1
// Output: false
//
// Example 3:
//
// Input: radius = 1, xCenter = 0, yCenter = 0, x1 = -1, y1 = 0, x2 = 0, y2 = 1
// Output: true
//
// Constraints:
//
// 1 <= radius <= 2000
//
// -104 <= xCenter, yCenter <= 104
//
// -104 <= x1 < x2 <= 104
//
// -104 <= y1 < y2 <= 104
// ============================================================================

package leetcode

func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
    
}
