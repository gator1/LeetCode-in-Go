package problem0004
import "math"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		panic("切片的长度为0，无法求解中位数。")
	}

	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	
    left := 0
	right := len(nums1)

    for left <=right {
		x := left + (right-left)/2
		y := (len(nums1)+len(nums2)+1) / 2 - x
	
		xL := -math.MaxInt32
		if x > 0 {
			xL = nums1[x-1]
		} 
		xR := math.MaxInt32
		if x < len(nums1)  {
			xR = nums1[x]
		}
		yL := -math.MaxInt32
		if y > 0 {
			yL = nums2[y-1]
		}  
		yR := math.MaxInt32
		if y < len(nums2)  {
			yR =nums2[y]
		}
		if xL <= yR && yL <= xR {
			if (len(nums1)+len(nums2)) % 2 != 0 {
				return float64(max(xL, yL))
			}
			return float64(max(xL, yL)+min(xR,yR)) / 2.0
		} else if xR > yL {
			right = x -1
		} else {
			left = x + 1
		}
			
	   
	}
	panic("切片的长度为0，无法求解中位数。")

}
