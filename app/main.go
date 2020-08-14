package main

import (
	"fmt"

	h "common-server/handlers"

	"github.com/urfave/negroni"
)

func main() {

	n := negroni.New()
	routes := h.Context()

	n.UseHandler(routes)
	n.Use(negroni.NewLogger())
	n.Run(":8880")
	fmt.Println("Serving On 8880")
}
