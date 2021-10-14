package handlers

import (
	"fmt"
	"net/http"
)

func HealthHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "%s", "ok")
}
