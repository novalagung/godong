package main

import (
	"fmt"
	"github.com/novalagung/godong"
	"github.com/novalagung/godong/testing/controller"
	"net/http"
)

func main() {
	var dashboard controller.Dashboard
	godong.Debug = true
	godong.Route(&dashboard)

	fmt.Println("start at :3000")
	http.ListenAndServe(":3000", nil)
}
