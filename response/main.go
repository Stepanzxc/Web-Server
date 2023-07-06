package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type Error struct {
	Error string `json:"err"`
}

// Response ...
func Response(w http.ResponseWriter, a any) {
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(a)
	if err != nil {
		ErrorFun(w, err)
		return
	}
	if _, err := w.Write(response); err != nil {
		ErrorFun(w, err)
		return
	}
}

func ErrorFun(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	a, err := json.Marshal(Error{err.Error()})
	if err != nil {
		log.Println(err)
	}
	_, err = w.Write(a)
	if err != nil {
		log.Println(err)
	}
}
