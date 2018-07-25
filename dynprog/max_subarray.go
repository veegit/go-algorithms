package dynprog

//MaxSum - returns maxSum
func MaxSum(nums []int) int {
	var maxEnding = 0
	var maxSoFar = 0
	for _, num := range nums {
		maxEnding += num
		if maxEnding < 0 {
			maxEnding = 0
		} else {
			if maxSoFar < maxEnding {
				maxSoFar = maxEnding
			}
		}
	}
	return maxSoFar
}
