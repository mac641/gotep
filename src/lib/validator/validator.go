package validator

import (
	"net/http"
)

type Validator struct{}

func (v *Validator) Send(requests []http.Request) ([]*http.Response, error) {
	// TODO: collect requests and print ALL results on console, rather than exiting at every failed request
	responses := []*http.Response{}
	client := &http.Client{}
	for i := range requests {
		r, err := client.Do(&requests[i])
		if err != nil {
			return responses, err
		}
		responses = append(responses, r)
	}

	return responses, nil
}
