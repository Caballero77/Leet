package main

func intersection(nums1 []int, nums2 []int) []int {
	dict := make(map[int]bool)
	for i := 0; i < len(nums1); i++ {
		dict[nums1[i]] = true
	}

	res := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		if dict[nums2[i]] {
			res = append(res, nums2[i])
			dict[nums2[i]] = false
		}
	}

	return res
}

func main() {

}
