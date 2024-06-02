package utils

import (
	"net/http"
	"io/ioutils"
	"encoding/json"
)

func ParseBody(r *http.Request, x interface{}){
	if body, err := ioutils.ReadAll(r.Body); err == nil{
		if err := json.Unmarshal([]byte(body), x); err != nil{
			return
		}
	}
}