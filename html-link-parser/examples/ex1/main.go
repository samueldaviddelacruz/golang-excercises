package main

import (
	"fmt"
	link "github.com/samueldaviddelacruz/golang-exercises/html-link-parser"
	"strings"
)

var exampleHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>
`

func main() {
	reader := strings.NewReader(exampleHtml)

	links, err := link.Parse(reader)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", links)
}
