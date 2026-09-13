// ============================================================================
// 题号: 835
// 题目: Image Overlap
// 难度: Medium
// 分类: Array, Matrix
// 链接: https://leetcode.cn/problems/image-overlap/
// 每日一题日期: 2026-09-13
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 给你两个图像 img1 和 img2 ，两个图像的大小都是 n x n ，用大小相同的二进制正方形矩阵表示。二进制矩阵仅由若干 0 和若干 1 组成。
//
// 转换 其中一个图像，将所有的 1 向左，右，上，或下滑动任何数量的单位；然后把它放在另一个图像的上面。该转换的 重叠 是指两个图像 都 具有 1 的位置的数目。
//
// 请注意，转换 不包括 向任何方向旋转。越过矩阵边界的 1 都将被清除。
//
// 最大可能的重叠数量是多少？
//
// 示例 1：
//
// 输入：img1 = [[1,1,0],[0,1,0],[0,1,0]], img2 = [[0,0,0],[0,1,1],[0,0,1]]
// 输出：3
// 解释：将 img1 向右移动 1 个单位，再向下移动 1 个单位。
//
// 两个图像都具有 1 的位置的数目是 3（用红色标识）。
//
// 示例 2：
//
// 输入：img1 = [[1]], img2 = [[1]]
// 输出：1
//
// 示例 3：
//
// 输入：img1 = [[0]], img2 = [[0]]
// 输出：0
//
// 提示：
//
// n == img1.length == img1[i].length
//
// n == img2.length == img2[i].length
//
// 1 <= n <= 30
//
// img1[i][j] 为 0 或 1
//
// img2[i][j] 为 0 或 1
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// You are given two images, img1 and img2, represented as binary, square matrices of size n x n. A binary matrix has only 0s and 1s as values.
//
// We translate one image however we choose by sliding all the 1 bits left, right, up, and/or down any number of units. We then place it on top of the other image. We can then calculate the overlap by counting the number of positions that have a 1 in both images.
//
// Note also that a translation does not include any kind of rotation. Any 1 bits that are translated outside of the matrix borders are erased.
//
// Return the largest possible overlap.
//
// Example 1:
//
// Input: img1 = [[1,1,0],[0,1,0],[0,1,0]], img2 = [[0,0,0],[0,1,1],[0,0,1]]
// Output: 3
// Explanation: We translate img1 to right by 1 unit and down by 1 unit.
//
// The number of positions that have a 1 in both images is 3 (shown in red).
//
// Example 2:
//
// Input: img1 = [[1]], img2 = [[1]]
// Output: 1
//
// Example 3:
//
// Input: img1 = [[0]], img2 = [[0]]
// Output: 0
//
// Constraints:
//
// n == img1.length == img1[i].length
//
// n == img2.length == img2[i].length
//
// 1 <= n <= 30
//
// img1[i][j] is either 0 or 1.
//
// img2[i][j] is either 0 or 1.
// ============================================================================

package leetcode

func largestOverlap(img1 [][]int, img2 [][]int) int {
    
}
