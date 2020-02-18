package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter target url: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}


/*
	TODO: populate this
*/
type PhisingResponse struct {
	
}

/*
	Query VirusTotal API for url
	TODO: find suitable API to use
*/
func query_api() {
	vt_Api_Key := os.Getenv("VirusTotalKey")
}