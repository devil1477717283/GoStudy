package doublepointer

func MaxArea(height []int) int {
	start, end, max := 0, len(height)-1, -1
	for start < end {
		lower := 0
		width := end - start
		if height[start] < height[end] {
			lower = height[start]
			start++
		} else {
			lower = height[end]
			end--
		}
		if lower*width > max {
			max = lower * width
		}
	}
	return max
}
