package Sort

func FindMaximumSubarray(a []int, low ,high int) (int, int, int){
	mid := (low + high)/2
	leftLow, leftHigh, leftSum, rightLow, rightHigh, rightSum, crossLow, crossHigh, crossSum := 0,0,0,0,0,0,0,0,0
	if high == low {
		return low, high, a[low]
	}else{
		leftLow, leftHigh, leftSum = FindMaximumSubarray(a,  low, mid)
		rightLow, rightHigh, rightSum = FindMaximumSubarray(a, mid+1, high)
		crossLow, crossHigh, crossSum = FindMaxCrossingSubarray(a, low, mid, high)
	}
	if leftSum >= rightSum && leftSum >= crossSum {
		return leftLow, leftHigh, leftSum
	}else if rightSum >= leftSum && leftHigh >= crossSum{
		return rightLow, rightHigh, rightSum
	}else{
		return crossLow, crossHigh, crossSum
	}

}

func FindMaxCrossingSubarray(a []int, low, mid, high int) (int, int, int){
	leftSum := -2147483648
	sum := 0
	var maxLeft, maxRight int

	for i := mid; i >= low; i--{
		sum = sum + a[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}
	rightSum := -2147483648
	sum = 0

	for j := mid +1; j <= high; j++{
		sum = sum + a[j]
		if sum > rightSum {
			rightSum = sum
			maxRight = j
		}
	}
	return maxLeft, maxRight, leftSum + rightSum
}