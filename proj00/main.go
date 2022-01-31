package main

import (
	"fewpudo/howtogolang/proj00/client"
	"fmt"
	"log"
)

func main() {
	crypto, err := client.FetchCrypto()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(crypto)
}
