package helpers

import (
	"net/http"
	"strings"
)

func GetFormData(r *http.Request) map[string]interface{} {
	r.ParseForm()
	datas := map[string]interface{}{}
	for key, value := range r.Form {
		if len(value) > 1 {
			datas[key] = value
		} else {
			if key == "password" {
				datas[key] = value[0]
			} else {
				datas[key] = strings.TrimSpace(value[0])
			}
		}
	}
	return datas
}
