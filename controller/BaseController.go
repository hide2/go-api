package controller

import "encoding/json"

func ResponseJSON(data interface{}) ([]byte, error) {
	js := make(map[string]interface{}, 0)
	js["code"] = 200
	js["message"] = "OK"
	js["data"] = data
	j, err := json.Marshal(js)
	if err != nil {
		js["code"] = 500
		js["message"] = err.Error()
		js["data"] = ""
		j, _ = json.Marshal(js)
	}
	return j, nil
}
