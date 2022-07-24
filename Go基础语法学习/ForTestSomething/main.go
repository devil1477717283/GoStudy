package main

import "fmt"

func update_slice(arr []int) []int {
	fmt.Println("Update Before....")
	fmt.Printf("%v,%p\n", arr, arr)
	arr = append(arr, 10)
	fmt.Println("Update After....")
	fmt.Printf("%v,%p\n", arr, arr)
	return arr
}
func main() {
	arr := make([]int, 0, 1)
	arr = append(arr, 5)
	fmt.Println("Use Function Before...")
	fmt.Printf("%v,%p\n", arr, arr)
	arr = update_slice(arr)
	fmt.Println("Use Function After...")
	fmt.Printf("%v,%p\n", arr, arr)
}
