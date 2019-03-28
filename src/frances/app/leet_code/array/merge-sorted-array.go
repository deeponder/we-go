package main

//sorted, 所以从后往前比较
func merge(nums1 []int, m int, nums2 []int, n int)  {
	i :=  m+n-1
	j := m-1
	k := n-1
	for j > 0 && k >0 {
		if nums1[j] > nums2[k] {
			nums1[i] = nums1[j]
			j--
		} else {
			nums1[i] = nums2[k]
			k--
		}
		i--
	}

	for k >= 0 {
		nums1[i] = nums2[k]
		k--
		i--
	}
}

func main() {

}
