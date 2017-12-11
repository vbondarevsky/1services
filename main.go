package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type InfoBases struct {
	ClientID           string
	InfoBasesCheckCode string
	InfoBases          string
}

func handler(w http.ResponseWriter, r *http.Request) {
	var infoBases InfoBases
	infoBases.ClientID = "000000"
	infoBases.InfoBasesCheckCode = "sfdsdf"
	infoBases.InfoBases = "sdfsdfsdfsfdsf"

	ib := make(map[string]interface{})
	ib["root"] = infoBases
	buf, _ := json.Marshal(ib)

	fmt.Fprintf(w, "%s", buf)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
