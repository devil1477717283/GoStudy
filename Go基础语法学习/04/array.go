package array

func SumAllTails(arrays ...[]int)(res int){
	for _,array := range arrays{
		if(len(array) == 0){
			res+=0
		}else{
			res+=array[len(array)-1]
		}
	}
	return 
}
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
    for _, numbers := range numbersToSum {
        sums = append(sums, Sum(numbers))
    }

    return sums
}

func Sum(arr []int) (res int) {
	for _, number := range arr {
		res += number
	}
	return
}
