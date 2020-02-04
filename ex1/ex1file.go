package ex1

import (
	"encoding/json"
	"fmt"
)

// ReverseRunes returns its argument string reversed rune-wise left to right.
func ReverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// ChanTest tests channel functionality
func ChanTest() {
	type Packet struct {
		Msg    string
		Status int
	}

	messages := make(chan Packet)

	go func() {
		messages <- Packet{
			Msg:    "begin",
			Status: 200,
		}

		received := <-messages
		out, err := json.Marshal(received)
		if err != nil {
			panic(err)
		}

		fmt.Println("worker received: " + string(out))
	}()

	received := <-messages
	out, err := json.Marshal(received)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T\n", out)
	fmt.Printf("%T\n", string(out))
	fmt.Printf("%s\n", string(out))

}
