/*
Write a program which prompts the user to first enter a name, and then enter an address. 
Your program should create a map and add the name and address to the map using the keys 
“name” and “address”, respectively. Your program should use Marshal() to create a JSON 
object from the map, and then your program should print the JSON object.
*/
package main
import "fmt"
import "encoding/json"
var name1,address1 string
func main() {
	
    fmt.Scan(&name1)
	fmt.Scan(&address1)
	commits := map[string]string{
    "name": name1,
    "address":   address1,
    
	}
	jsonData,err:=json.Marshal(commits)
	fmt.Println(string(jsonData))
	if err!=nil{
		fmt.Println(err)
	}
}