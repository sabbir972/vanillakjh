package handler

import (
	"fmt"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	currentTime := new Date()
	fmt.Fprintf(w, currentTime)
}
