package SliceOperation

func Add(Slice []int, index, num int) []int {
	var newSlice []int
	if len(Slice) != cap(Slice) {
		newSlice = make([]int, len(Slice)+1, cap(Slice))
	} else {
		newSlice = make([]int, len(Slice)+1, cap(Slice)*2)
	}
	for i := 0; i < len(newSlice); i++ {
		if i == index {
			newSlice[i] = num

		} else if i > index {
			newSlice[i] = Slice[i-1]
		} else {
			newSlice[i] = Slice[i]
		}
	}
	return newSlice
}
func Delete(Slice []int, index int) []int {

	newSlice := make([]int, len(Slice)-1, cap(Slice))
	for i := 0; i < len(Slice); i++ {
		if i < index {
			newSlice[i] = Slice[i]
		} else if i == index {
			continue
		} else {
			newSlice[i-1] = Slice[i]
		}
	}
	return newSlice

}
