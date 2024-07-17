package handlers

import (
	"encoding/json"
	"net/http"
)

type problemsResponse struct {
	Problems []string `json:"problems"`
}

var response, _ = json.Marshal(&problemsResponse{
	Problems: []string{
		"Basin Function Problem",
	},
})

func GetProblems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
