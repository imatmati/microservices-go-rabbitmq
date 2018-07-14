package http

import (
	l "account/utils/language"
	"fmt"
	"net/http"
)

func ExtractParametersFrom(rq *http.Request, parameters ...string) []string {
	fmt.Println(rq.PostForm)

	parameters_values := make([]string, len(parameters))
	for i, parameter := range parameters {
		parameters_values[i] = rq.FormValue(parameter)
		if parameters_values[i] == "" {
			l.PanicIf(fmt.Errorf("Parameter %s absent", parameter))
		}
	}
	return parameters_values
}

func ErrorJSON(err string) string {
	return fmt.Sprintf("{\"error\":\"%s\"}", err)
}
