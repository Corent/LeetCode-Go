package algorithm870

import "sort"

/*
class Solution {
    public int[] advantageCount(int[] nums1, int[] nums2) {
        int n = nums1.length;
        Integer[] idx = new Integer[n]; // nums2索引
        for (int i = 0; i < n; i++) idx[i] = i;
        Arrays.sort(nums1);
        Arrays.sort(idx, (i, j) -> nums2[i] - nums2[j]); //根据nums2数据升序，排序nums2(索引)
        int L = 0, R = n - 1; // nums2(索引)的左右指针
        for (int i : nums1) { // 遍历nums1
            if (i > nums2[idx[L]]) nums2[idx[L++]] = i; // 可以满足 nums1[i] > nums2[i]
            else nums2[idx[R--]] = i; //丢到数组最后
        }
        return nums2;
    }
}
*/

func advantageCount(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	idx := make([]int, n)
	for i := 0; i < n; i++ {
		idx[i] = i
	}
	sort.Ints(nums1)
	sort.Slice(idx, func(i, j int) bool {
		return nums2[idx[i]] < nums2[idx[j]]
	})
	left, right := 0, n-1
	for _, num := range nums1 {
		if num > nums2[idx[left]] {
			nums2[idx[left]] = num
			left++
		} else {
			nums2[idx[right]] = num
			right--
		}
	}
	return nums2
}
