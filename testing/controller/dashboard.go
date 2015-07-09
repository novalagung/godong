package controller

import (
	"fmt"
	"net/http"
)

type Dashboard struct{}

func (d *Dashboard) Action_Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("index")
}

func (d *Dashboard) Action_Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("home")
}

func (d *Dashboard) Action_About(w http.ResponseWriter, r *http.Request) {
	fmt.Println("about")
}
