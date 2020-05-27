/*
Write a program to sort an array of integers. 
The program should partition the array into 4 parts, 
each of which is sorted by a different goroutine. 
Each partition should be of approximately equal size. 
Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. 
Each goroutine which sorts ¼ of the array should print the subarray that it will sort. 
When sorting is complete, the main goroutine should print the entire sorted list.
*/
package main

import (
	"fmt"
	"sort"
	"sync"
)
func main() {

	var numInts int
	var userNum int
	var waitGroup sync.WaitGroup
	sliceOfInts := make([]int, 0, 3)
	fmt.Scan(&numInts)

	for i := 0; i < numInts; i++ {

		fmt.Println("Please Enter a number to fill the array")
		_, err := fmt.Scan(&userNum)
		if err != nil {
			fmt.Println("Invalid user input")
		}

		sliceOfInts = append(sliceOfInts, userNum)

	}

	sliceSize := numInts / 4
	slice1 := sliceOfInts[:sliceSize]
	slice2 := sliceOfInts[sliceSize : 2*(sliceSize)]
	slice3 := sliceOfInts[2*(sliceSize) : 3*(sliceSize)]
	slice4 := sliceOfInts[3*(sliceSize):]


	waitGroup.Add(4) 
	go sortList(slice1)
	waitGroup.Done()
	go sortList(slice2)
	waitGroup.Done()
	go sortList(slice3)
	waitGroup.Done()
	go sortList(slice4)
	waitGroup.Done()
	waitGroup.Wait()

	newSlice := mergeAndSort(slice1, slice2, slice3, slice4)
	fmt.Println(newSlice)

}

func sortList(unsortedSlice []int) []int {
	sort.Ints(unsortedSlice)
	return unsortedSlice

}

func mergeAndSort(list1 []int, list2 []int, list3 []int, list4 []int) []int {
	newSlice := []int{}
	newSlice = append(list1, list2...)
	newSlice = append(newSlice, list3...)
	newSlice = append(newSlice, list4...)
	sort.Ints(newSlice)
	return newSlice

}