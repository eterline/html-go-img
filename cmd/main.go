package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eterline/html-go-img/convert"
)

func main() {
	f, _ := os.ReadFile("./index.html")
	c := convert.NewConverterImg()
	c.BytesPayload(f)
	err := c.Convert()
	if err != nil {
		log.Fatal(err)
	}
	err = c.CreateFile("test", convert.PNG)
	if err != nil {
		log.Fatal(err)
	}
	err = c.CreateFile("test", convert.JPG)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c.ToBase64(convert.JPG))
}
