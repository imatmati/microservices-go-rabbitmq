package data

import "account/utils/json"

type Result struct {
	Account string `json:"account"`
	Check   string `json:"check"`
}

func (r Result) ToJSON() (string, error) {

	return json.ToJSON(r)

}
