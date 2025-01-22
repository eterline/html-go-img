
# html-go-img

HTML to Image/PDF converter in go.

`wkhtmltoimage` go wrapper linrary.


## Exmaple

code usage exmaple

```go
func main() {

	// Creates converter object
	c := convert.NewConverterImg()

	// Reading HTML page from file
	f, _ := os.ReadFile("./index.html")

	// Upload file content to converter
	c.BytesPayload(f)

	// Make convert
	err := c.Convert()
	if err != nil {
		log.Fatal(err)
	}

	// Save in new/exiting file with rewriting old data
	// output file: 'test.png'
	err = c.SaveFile("test", convert.PNG)
	if err != nil {
		log.Fatal(err)
	}

	// Base 64 input output
	fmt.Println(c.ToBase64())

	// HTML from string
	c.StringPayload("<h1>Text<h1/>")
}
```
## License

[MIT](https://choosealicense.com/licenses/mit/)

