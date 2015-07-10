package controller

import (
	"fmt"
	"net/http"
)

type Dashboard struct{}

func (d *Dashboard) Action_Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "index")
}

func (d *Dashboard) Action_Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home")
}

func (d *Dashboard) Action_AboutUs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "about")
}

func (d *Dashboard) Action_DataAnalytic_GetData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "data-analytic/get-data")
}
