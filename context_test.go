package go_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
)

func TestContext(t *testing.T) {
	// context background dan todo adalah context kosong
	// perbedaannya adalah, context todo biasanya digunakan
	// ketika kita belum tau akan menggunakan context apa

	// context itu immutable
	// setelah context dibuat tidak bisa diubah lagi
	// ketika kita melakukan perubahan value pada context (value, timeout, dll)
	// itu akan membuat context baru (child)
	// parent context nya tidak akan berubah

	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	// contextA mempunyai child context yaitu contextB dan contextC
	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	// contextB mempunyai child context yaitu contexD dan contextE
	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	// ketika kita memberi sebuah value pada parent context
	// maka child context dari parent tersebut akan mendapatkan data juga

	// ketika kita membatalkan sebuah proses pada sebuah context
	// maka proses pada child context tersebut juga akan dibatalkan
	// namun tidak untuk parent context dan context tersebut

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	// cara untuk mendapatkan value dari sebuah context
	// untuk mendapatkan value, sebuah context akan mengambil data dari context dirinya sendiri
	// jika tidak ada, maka context nya akan mengambil data dari parent nya
	// begitu seterusnya
	// jika hingga parent teratas masih belum mendapatkan data yang diinginkan
	// maka akan return nil
	fmt.Println(contextF.Value("f"))
}

func CreateCounter() chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			destination <- counter
			counter++
		}
	}()

	return destination
}

func TestContextWithCancle(t *testing.T) {
	fmt.Println("Total Goroutine", runtime.NumGoroutine())

	destination := CreateCounter()
	for n := range destination {
		fmt.Println("Counter", n)

		// ini akan mengakibatkan goroutine leak
		// karena, ketika n == 10, maka perulangan ini akan di break
		// namun goroutine akan tetap berjalan dan mengirim data ke channel
		// karena masih belum ada sinyal cancle
		// sinyal cancle akan membuat proses goroutine berhenti

		if n == 10 {
			break
		}
	}

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}
