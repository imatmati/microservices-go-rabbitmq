package data

import "account/utils/json"

type Result struct {
	Account string `json:"account"`
	Result  string `json:"result"`
}

func (r Result) ToJSON() (string, error) {

	return json.ToJSON(r)

}
