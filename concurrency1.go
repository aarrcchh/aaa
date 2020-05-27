package main
import "fmt"
import "time"
func main(){
	count("Sheep")
	count("Fish")
}
func count(thing string){
	for i:=1;true;i++{
		fmt.Println(i,thing)
		time.Sleep(time.Millisecond*500)
	}
}