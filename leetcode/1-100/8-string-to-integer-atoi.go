// ============================================================================
// 题号: 8
// 题目: String to Integer (atoi)
// 难度: Medium
// 分类: String
// 链接: https://leetcode.cn/problems/string-to-integer-atoi/
// 拉取日期: 2026-07-26
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 请你来实现一个 myAtoi(string s) 函数，使其能将字符串转换成一个 32 位有符号整数。
//
// 函数 myAtoi(string s) 的算法如下：
//
// 空格：读入字符串并丢弃无用的前导空格（" "）
//
// 符号：检查下一个字符（假设还未到字符末尾）为 '-' 还是 '+'。如果两者都不存在，则假定结果为正。
//
// 转换：通过跳过前置零来读取该整数，直到遇到非数字字符或到达字符串的结尾。如果没有读取数字，则结果为0。
//
// 舍入：如果整数数超过 32 位有符号整数范围 [−231,  231 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −231 的整数应该被舍入为 −231 ，大于 231 − 1 的整数应该被舍入为 231 − 1 。
//
// 返回整数作为最终结果。
//
// 示例 1：
//
// 输入：s = "42"
//
// 输出：42
//
// 解释：加粗的字符串为已经读入的字符，插入符号是当前读取的字符。
//
// 带下划线线的字符是所读的内容，插入符号是当前读入位置。
// 第 1 步："42"（当前没有读入字符，因为没有前导空格）
//          ^
// 第 2 步："42"（当前没有读入字符，因为这里不存在 '-' 或者 '+'）
//          ^
// 第 3 步："42"（读入 "42"）
//            ^
//
// 示例 2：
//
// 输入：s = " -042"
//
// 输出：-42
//
// 解释：
//
// 第 1 步："   -042"（读入前导空格，但忽视掉）
//             ^
// 第 2 步："   -042"（读入 '-' 字符，所以结果应该是负数）
//              ^
// 第 3 步："   -042"（读入 "042"，在结果中忽略前导零）
//                ^
//
// 示例 3：
//
// 输入：s = "1337c0d3"
//
// 输出：1337
//
// 解释：
//
// 第 1 步："1337c0d3"（当前没有读入字符，因为没有前导空格）
//          ^
// 第 2 步："1337c0d3"（当前没有读入字符，因为这里不存在 '-' 或者 '+'）
//          ^
// 第 3 步："1337c0d3"（读入 "1337"；由于下一个字符不是一个数字，所以读入停止）
//              ^
//
// 示例 4：
//
// 输入：s = "0-1"
//
// 输出：0
//
// 解释：
//
// 第 1 步："0-1" (当前没有读入字符，因为没有前导空格)
//          ^
// 第 2 步："0-1" (当前没有读入字符，因为这里不存在 '-' 或者 '+')
//          ^
// 第 3 步："0-1" (读入 "0"；由于下一个字符不是一个数字，所以读入停止)
//           ^
//
// 示例 5：
//
// 输入：s = "words and 987"
//
// 输出：0
//
// 解释：
//
// 读取在第一个非数字字符“w”处停止。
//
// 提示：
//
// 0 <= s.length <= 200
//
// s 由英文字母（大写和小写）、数字（0-9）、' '、'+'、'-' 和 '.' 组成
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer.
//
// The algorithm for myAtoi(string s) is as follows:
//
// Whitespace: Ignore any leading whitespace (" ").
//
// Signedness: Determine the sign by checking if the next character is '-' or '+', assuming positivity if neither present.
//
// Conversion: Read the integer by skipping leading zeros until a non-digit character is encountered or the end of the string is reached. If no digits were read, then the result is 0.
//
// Rounding: If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then round the integer to remain in the range. Specifically, integers less than -231 should be rounded to -231, and integers greater than 231 - 1 should be rounded to 231 - 1.
//
// Return the integer as the final result.
//
// Example 1:
//
// Input: s = "42"
//
// Output: 42
//
// Explanation:
//
// The underlined characters are what is read in and the caret is the current reader position.
// Step 1: "42" (no characters read because there is no leading whitespace)
//          ^
// Step 2: "42" (no characters read because there is neither a '-' nor '+')
//          ^
// Step 3: "42" ("42" is read in)
//            ^
//
// Example 2:
//
// Input: s = " -042"
//
// Output: -42
//
// Explanation:
//
// Step 1: "   -042" (leading whitespace is read and ignored)
//             ^
// Step 2: "   -042" ('-' is read, so the result should be negative)
//              ^
// Step 3: "   -042" ("042" is read in, leading zeros ignored in the result)
//                ^
//
// Example 3:
//
// Input: s = "1337c0d3"
//
// Output: 1337
//
// Explanation:
//
// Step 1: "1337c0d3" (no characters read because there is no leading whitespace)
//          ^
// Step 2: "1337c0d3" (no characters read because there is neither a '-' nor '+')
//          ^
// Step 3: "1337c0d3" ("1337" is read in; reading stops because the next character is a non-digit)
//              ^
//
// Example 4:
//
// Input: s = "0-1"
//
// Output: 0
//
// Explanation:
//
// Step 1: "0-1" (no characters read because there is no leading whitespace)
//          ^
// Step 2: "0-1" (no characters read because there is neither a '-' nor '+')
//          ^
// Step 3: "0-1" ("0" is read in; reading stops because the next character is a non-digit)
//           ^
//
// Example 5:
//
// Input: s = "words and 987"
//
// Output: 0
//
// Explanation:
//
// Reading stops at the first non-digit character 'w'.
//
// Constraints:
//
// 0 <= s.length <= 200
//
// s consists of English letters (lower-case and upper-case), digits (0-9), ' ', '+', '-', and '.'.
// ============================================================================

package leetcode

func myAtoi(s string) int {
    
}
