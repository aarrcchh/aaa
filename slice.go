/*
Write a program which prompts the user to enter integers and stores the 
integers in a sorted slice. The program should be written as a loop. 
Before entering the loop, the program should create an empty integer 
slice of size (length) 3. During each pass through the loop, the 
program prompts the user to enter an integer to be added to the slice. 
The program adds the integer to the slice, sorts the slice, and prints the 
contents of the slice in sorted order. The slice must grow in size to 
accommodate any number of integers which the user decides to enter. 
The program should only quit (exiting the loop) when the user enters 
the character ‘X’ instead of an integer.
*/
package main
import "fmt"
import "strconv"
import "sort"
func main() {
	a := make([]int, 3)
	var integer string
	var i int
	i=0
	n:=len(a)
	fmt.Scan(&integer)
	for integer != "X" {
		if n== cap(a){
			newSlice := make([]int, len(a), 2*len(a)+1)
			copy(newSlice, a)
			a = newSlice
			
		}
		b,err := strconv.Atoi(integer)
		
		if err== nil{
			a = a[0 : i+1]
			a[i]=b
			fmt.Scan(&integer)
			i++
			n=len(a)
		}else{
			break
		}
	}
	sort.Ints(a) 
	fmt.Println(a)
}
