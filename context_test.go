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

	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}
