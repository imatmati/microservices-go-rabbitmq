package http

import "net/http"

func ExtractParametersFrom(rq *http.Request, parameters ...string) []string {
	parameters_values := make([]string, len(parameters))
	for i, parameter := range parameters {
		parameters_values[i] = rq.FormValue(parameter)
	}
	return parameters_values
}
