package model

import (
	"fmt"
)

// Cryptoresponse is exported, it models the data we receive.
type Cryptoresponse []struct {
	Eyes       string `json:"Eyes"`
	Mouth      string `json:"Mouth"`
	Fur        string `json:"Fur"`
	Background string `json:"Background"`
	Hat        string `json:"Hat"`
	Clothes    string `json:"Clothes"`
}

//TextOutput is exported,it formats the data to plain text.
func (c Cryptoresponse) TextOutput() string {
	p := fmt.Sprintf(
		"Eyes: %s\nMouth : %s\nFur: %s\nBackground: %s\nHat: %s\nClothes: %s\n",
		c[0].Eyes, c[0].Mouth, c[0].Fur, c[0].Background, c[0].Hat, c[0].Clothes)
	return p
}
