package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type InfoBases struct {
	Root struct {
		ClientID           string `json:"ClientID"`
		InfoBasesCheckCode string `json:"InfoBasesCheckCode"`
		InfoBases          string `json:"InfoBases"`
	} `json:"root"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var infoBases InfoBases
	infoBases.Root.ClientID = "000000"
	infoBases.Root.InfoBasesCheckCode = "sfdsdf"
	infoBases.Root.InfoBases = "sdfsdfsdfsfdsf"

	buf, _ := json.Marshal(infoBases)

	fmt.Fprintf(w, "%s", buf)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
