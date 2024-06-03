# 刷题内容
> 原题链接
- https://leetcode.com/problems/two-sum
- https://leetcode-cn.com/problems/two-sum
> 内容描述

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

# 解题方案
## 思路1

我们在找两个的和是否等于某一个值。 

如果是两两组合的笛卡尔积，相当于闭着眼碰运气，时间复杂度为: O(n^2)

例如：
```txt
2+7     2+11    2+15
7+11    7+15
11+15
```


```golang []
func twoSum(nums []int, target int) []int {
    for i, x := range nums {
        for j := i + 1; j < len(nums); j++ {
            if x+nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return nil
}
```
```python []
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        for i in range(len(nums)):
            for j in range(i+1, len(nums)):
                if nums[i] + nums[j] == target:
                    return [i, j]
        return []
```

## 思路2
哈希表

注意到方法一的时间复杂度较高的原因是寻找 target - x 的时间复杂度过高。因此，我们需要一种更优秀的方法，能够快速寻找数组中是否存在目标元素。如果存在，我们需要找出它的索引。

使用哈希表，可以将target - x 的时间复杂度降低到 O(1)。

我们遍历数组，对于每一个 x，我们首先查询哈希表中是否存在 target - x，然后将 x 插入到哈希表中，即可保证插入的 x 和查询的 target - x 不是同一个元素。如果存在，则找到答案，否则继续遍历。

> 需注意是先判断后存储，因为如果先存储，再判断，会出现相同的情况。例如nums=[2,3,4], target=4。会输出[0,0]

```golang
func twoSum(nums []int, target int) []int {
    haseTable := map[int]int{}
    for i, v := range nums{
        if j, ok := haseTable[target-v]; ok {
            return []int{i,j}
        }
    }
    return nil
}

```