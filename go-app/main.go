// Sampe-RestServer project main.go
// Mock_Server project main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	contenttypeJSON  = "application/json; charset=utf-8"
	maxpayloadlength = 1 << 12
)

type listValue struct {
	TotalPages int      `json:"totalpages"`
	Domains    []string `json:"domains"`
}

type GTI struct {
	URL       string `json:"url"`
	ID        string `json:"id"`
	Password  string `json:"password"`
	ClientID  string `json:"clientid"`
	Product   string `json:"product"`
	Certcheck bool   `json:"certcheck"`
	Timeout   string `json:"timeout"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requesting for getting gti cred from server.....!")
	gti := GTI{
		URL:       "https://sample.com/1",
		ID:        "is-beta",
		Password:  "vA1c1(92a%",
		ClientID:  "67d8d1082c2f2f",
		Product:   "SConnect",
		Certcheck: false,
		Timeout:   "5s",
	}
	w.Header().Set("Content-Type", contenttypeJSON)
	json.NewEncoder(w).Encode(gti)
}

func main() {
	fmt.Println("Hello World!")
	http.HandleFunc("/gti", handler)
	http.ListenAndServe(":8082", nil)

}
