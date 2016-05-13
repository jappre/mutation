package main
import "fmt"
func main() {
  xs := []float64{1, 2, 3, 4, 5, 6}
  start := make(chan bool)
	for i,_ := range xs {
		go worker(start, i)
	}
  close(start)
  select{}
}

func worker(start chan bool, index int) {
	<-start
	fmt.Println("this is the :", index)
}
