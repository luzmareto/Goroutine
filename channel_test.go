package goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	//var penampung channel
	channel := make(chan string)
	defer close(channel)

	//goroutines with anonym func
	go func() {
		time.Sleep(2 * time.Second)

		//mengirim data ke channel
		channel <- "Luz Mareto"

		fmt.Println("Selesai Mengirim Data ke Channel")
	}()

	data := <-channel //data dari channel
	fmt.Println(data) //Luzmareto Rusmadian

	time.Sleep(5 * time.Second)
}

// channel parameter
func GiveMeRespone(channel chan string) {
	time.Sleep(2 * time.Second)

	//memberi data/value ke channel
	channel <- "Mareto Luz"
}

func TestChannelAsParameter(t *testing.T) {
	//var penampung channel
	channel := make(chan string)
	defer close(channel)

	go GiveMeRespone(channel)

	data := <-channel //mengirim channel ke var penampung
	fmt.Println(data) //Luzmareto Rusmadian

	time.Sleep(5 * time.Second)
}

// func khusus mengirim data
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Mareto Luz"
}

// func khusus menerima data
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestOnlyInOut(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	//jumlah data tidak boleh melebihi buffer
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		//mengirim data
		channel <- "Luz"
		channel <- "mareto"
	}()

	go func() {
		fmt.Println(<-channel) //Luz
		fmt.Println(<-channel) //mareto

	}()

	time.Sleep(2 * time.Second)

	fmt.Println("selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	//mengambill func GiveMeRespone
	go GiveMeRespone(channel1)
	go GiveMeRespone(channel2)

	//mencetak channel menggunakan perulangan
	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari cannnel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelect(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	//mengambill func GiveMeRespone
	go GiveMeRespone(channel1)
	go GiveMeRespone(channel2)

	//mencetak channel menggunakan perulangan
	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari cannnel 2", data)
			counter++
		default:
			fmt.Println("menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}
