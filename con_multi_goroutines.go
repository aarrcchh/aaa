package main
import "fmt"
import "time"
func main(){
	c1:=make(chan string)
	c2:=make(chan string)
	
	go func(){
		for{
			time.Sleep(time.Millisecond*500)
			c1<-"Every 500ms"
		}
		
		
	}()
	
	go func(){
		for{
			time.Sleep(time.Second*2)
			c2<-"Every 2 seconds"
		}
		
	}()
	
	for{
		select{
			case msg1:=<-c1:
				fmt.Println(msg1)
			case msg2:=<-c2:
				fmt.Println(msg2)
		}
	}
	
	
}
