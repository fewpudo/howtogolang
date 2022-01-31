package client

import (
	"encoding/json"
	"fewpudo/howtogolang/proj00/model"
	"log"
	"net/http"
)

//Fetch is exported ...
func FetchCrypto() (string, error) {
	//Build The URL string
	URL := "https://nonfungible.com/project/boredapeclub/BAYC/2"
	//We make HTTP request using the Get function
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal("Erro no httpGet")
	}
	defer resp.Body.Close()

	//Create a variable of the same type as our model
	var cResp model.Cryptoresponse

	//Decode the data
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		log.Fatal("Erro no decodificador")
	}
	//Invoke the text output function & return it with nil as the error value
	return cResp.TextOutput(), nil
}
