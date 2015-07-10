package main

import (
	"github.com/novalagung/godong"
	"github.com/novalagung/godong/testing/controller"
	"net/http"
)

func main() {
	godong.Debug = true
	godong.Route(&controller.Dashboard{})

	http.ListenAndServe(":3000", nil)
}
