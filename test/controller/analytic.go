package controller

import (
	"fmt"
	"net/http"
)

type Analytic struct{}

func (a *Analytic) Action_Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "index")
}
