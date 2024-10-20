# 刷题内容
> 原题链接
- https://leetcode.cn/problems/merge-sorted-array/?envType=study-plan-v2&envId=top-interview-150
> 内容描述

给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。

请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

注意：最终，合并后数组不应由函数返回，而是存储在数组 nums1 中。为了应对这种情况，nums1 的初始长度为 m + n，其中前 m 个元素表示应合并的元素，后 n 个元素为 0 ，应忽略。nums2 的长度为 n 。

示例 1：

输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
输出：[1,2,2,3,5,6]
解释：需要合并 [1,2,3] 和 [2,5,6] 。
合并结果是 [1,2,2,3,5,6] ，其中斜体加粗标注的为 nums1 中的元素。
# 解题方案
## 思路1
> 直接合并后排序
> 
```golang
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func merge(nums1 []int, m int, nums2 []int, n int) {
    copy(nums1[m:], nums2)
    sort.Ints(nums1)
}
```
> 复杂度分析
> - 时间复杂度：O((m+n)log(m+n))。
> \
> 排序序列长度为 m+n，套用快速排序的时间复杂度即可，平均情况为 O((m+n)log(m+n))。
> - 空间复杂度：O(log(m+n))。
> \
排序序列长度为 m+n，套用快速排序的空间复杂度即可，平均情况为 O(log(m+n))。

## 思路2
> 双指针法：方法一没有利用数组 nums1 与 nums2 已经被排序的性质。为了利用这一性质，我们可以使用双指针方法。这一方法将两个数组看作队列，每次从两个数组头部取出比较小的数字放到结果中。

我们为两个数组分别设置一个指针 p1 与 p2 来作为队列的头部指针。代码实现如下：
> 
```golang
func merge(nums1 []int, m int, nums2 []int, n int) {
    sorted := make([]int, 0, m+n)
    p1, p2 := 0, 0
    for {
        if p1 == m {
            sorted = append(sorted, nums2[p2:]...)
            break
        }
        if p2 == n {
            sorted = append(sorted, nums1[p1:]...)
            break
        }
        if nums1[p1] < nums2[p2] {
            sorted = append(sorted, nums1[p1])
            p1++
        } else {
            sorted = append(sorted, nums2[p2])
            p2++
        }
    }
    copy(nums1, sorted)
}
```
> 复杂度分析
> - 时间复杂度：O(m+n)。
> \
指针移动单调递增，最多移动 m+n 次，因此时间复杂度为 O(m+n)。
> - 空间复杂度：O(m+n)。
> \
> 需要建立长度为 m+n 的中间数组 sorted。
