package main

import (
	"fmt"
	"encoding/json"
	"os"
	"bufio"
)

func main() {
	var m map[string]string
	m = make(map[string]string)
	fmt.Println("Please enter your name:")
	br := bufio.NewReader(os.Stdin)
	bname, _, _ := br.ReadLine()
	name := string(bname)
	m["name"] = name
	fmt.Println("Please enter your address:")
	br1 := bufio.NewReader(os.Stdin)
	baddress, _, _ := br1.ReadLine()
	address := string(baddress)
	m["address"] = address
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Encoding failed")
	} else {
		fmt.Println("Encoding data : ")
		fmt.Println(b)
		fmt.Println("Decoding data : ")
		fmt.Println(string(b))
	}
}