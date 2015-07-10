package main

import (
	"github.com/novalagung/godong"
	"github.com/novalagung/godong/testing/controller"
	"net/http"
)

func main() {
	var dashboard controller.Dashboard
	godong.Debug = true
	godong.Route(&dashboard)

	http.ListenAndServe(":3000", nil)
}
