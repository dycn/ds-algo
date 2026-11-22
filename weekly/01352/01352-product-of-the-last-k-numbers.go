// ============================================================================
// 题号: 1352
// 题目: Product of the Last K Numbers
// 难度: Medium
// 分类: Design, Array, Math, Data Stream, Prefix Sum
// 链接: https://leetcode.cn/problems/product-of-the-last-k-numbers/
// 每日一题日期: 2026-11-23
// ============================================================================

// ============================================================================
// 题目描述(中文)
// ============================================================================
// 设计一个算法，该算法接受一个整数流并检索该流中最后 k 个整数的乘积。
//
// 实现 ProductOfNumbers 类：
//
// ProductOfNumbers() 用一个空的流初始化对象。
//
// void add(int num) 将数字 num 添加到当前数字列表的最后面。
//
// int getProduct(int k) 返回当前数字列表中，最后 k 个数字的乘积。你可以假设当前列表中始终 至少 包含 k 个数字。
//
// 题目数据保证：任何时候，任一连续数字序列的乘积都在 32 位整数范围内，不会溢出。
//
// 示例：
//
// 输入：
// ["ProductOfNumbers","add","add","add","add","add","getProduct","getProduct","getProduct","add","getProduct"]
// [[],[3],[0],[2],[5],[4],[2],[3],[4],[8],[2]]
//
// 输出：
// [null,null,null,null,null,null,20,40,0,null,32]
//
// 解释：
// ProductOfNumbers productOfNumbers = new ProductOfNumbers();
// productOfNumbers.add(3);        // [3]
// productOfNumbers.add(0);        // [3,0]
// productOfNumbers.add(2);        // [3,0,2]
// productOfNumbers.add(5);        // [3,0,2,5]
// productOfNumbers.add(4);        // [3,0,2,5,4]
// productOfNumbers.getProduct(2); // 返回 20 。最后 2 个数字的乘积是 5 * 4 = 20
// productOfNumbers.getProduct(3); // 返回 40 。最后 3 个数字的乘积是 2 * 5 * 4 = 40
// productOfNumbers.getProduct(4); // 返回  0 。最后 4 个数字的乘积是 0 * 2 * 5 * 4 = 0
// productOfNumbers.add(8);        // [3,0,2,5,4,8]
// productOfNumbers.getProduct(2); // 返回 32 。最后 2 个数字的乘积是 4 * 8 = 32
//
// 提示：
//
// 0 <= num <= 100
//
// 1 <= k <= 4 * 104
//
// add 和 getProduct 最多被调用 4 * 104 次。
//
// 在任何时间点流的乘积都在 32 位整数范围内。
//
// 进阶：您能否 同时 将 GetProduct 和 Add 的实现改为 O(1) 时间复杂度，而不是 O(k) 时间复杂度？
// ============================================================================

// ============================================================================
// 题目描述(英文)
// ============================================================================
// Design an algorithm that accepts a stream of integers and retrieves the product of the last k integers of the stream.
//
// Implement the ProductOfNumbers class:
//
// ProductOfNumbers() Initializes the object with an empty stream.
//
// void add(int num) Appends the integer num to the stream.
//
// int getProduct(int k) Returns the product of the last k numbers in the current list. You can assume that always the current list has at least k numbers.
//
// The test cases are generated so that, at any time, the product of any contiguous sequence of numbers will fit into a single 32-bit integer without overflowing.
//
// Example:
//
// Input
// ["ProductOfNumbers","add","add","add","add","add","getProduct","getProduct","getProduct","add","getProduct"]
// [[],[3],[0],[2],[5],[4],[2],[3],[4],[8],[2]]
//
// Output
// [null,null,null,null,null,null,20,40,0,null,32]
//
// Explanation
// ProductOfNumbers productOfNumbers = new ProductOfNumbers();
// productOfNumbers.add(3);        // [3]
// productOfNumbers.add(0);        // [3,0]
// productOfNumbers.add(2);        // [3,0,2]
// productOfNumbers.add(5);        // [3,0,2,5]
// productOfNumbers.add(4);        // [3,0,2,5,4]
// productOfNumbers.getProduct(2); // return 20. The product of the last 2 numbers is 5 * 4 = 20
// productOfNumbers.getProduct(3); // return 40. The product of the last 3 numbers is 2 * 5 * 4 = 40
// productOfNumbers.getProduct(4); // return 0. The product of the last 4 numbers is 0 * 2 * 5 * 4 = 0
// productOfNumbers.add(8);        // [3,0,2,5,4,8]
// productOfNumbers.getProduct(2); // return 32. The product of the last 2 numbers is 4 * 8 = 32
//
// Constraints:
//
// 0 <= num <= 100
//
// 1 <= k <= 4 * 104
//
// At most 4 * 104 calls will be made to add and getProduct.
//
// The product of the stream at any point in time will fit in a 32-bit integer.
//
// Follow-up: Can you implement both GetProduct and Add to work in O(1) time complexity instead of O(k) time complexity?
// ============================================================================

package leetcode

type ProductOfNumbers struct {
    
}


func Constructor() ProductOfNumbers {
    
}


func (this *ProductOfNumbers) Add(num int)  {
    
}


func (this *ProductOfNumbers) GetProduct(k int) int {
    
}


/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
