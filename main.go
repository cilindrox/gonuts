package gonuts

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(FizzBuzz(i))
	}
}
