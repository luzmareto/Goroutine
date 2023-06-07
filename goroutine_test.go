package goroutines

import (
	"fmt"
	"testing"
	"time"
)

/*
default go routine adalah menambahkan kata go tepat sebelum pemanggilan nama function
maka setiap func akan memproses kode yang paling mudah dulu atau asyn
*/

func RunHelloWorld() {
	fmt.Println("Hello Wolrd")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
