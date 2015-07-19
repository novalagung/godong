package main

import (
	"github.com/novalagung/godong"
	"github.com/novalagung/godong/test/controller"
	"net/http"
)

func main() {
	godong.Debug = true
	godong.DefaultAction = "Dashboard.Action_Index"
	godong.HiddenIndex = true
	godong.Route(&controller.Dashboard{})
	godong.Route(&controller.Analytic{})

	http.ListenAndServe(":3000", nil)
}
