package main
import "fmt"
import "time"
func main(){
	c:=make(chan string)
	
	go count("Sheep",c)
	for{	// deaklock occrs because count func ends after i==5 but msg waits to receive data from channe

		msg,open:=<-c
		if !open{
			break
		}
		fmt.Println(msg)
	}
	
}
func count(thing string,c chan string){
	for i:=1;i<=5;i++{
		c<-thing
		time.Sleep(time.Millisecond*500)
	}
	close(c)
}