package go_context

import (
	"context"
	"fmt"
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
