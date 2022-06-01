package validator

import (
	"net/http"
)

type Validator struct{}

// Checks if request was successful, if so, increments the success count which will be returned
func (v *Validator) Analyze(responses []*http.Response) int {
	successful := 0

	for _, response := range responses {
		if v.isStatusSuccessful(response.StatusCode) {
			successful++
			continue
		}
	}

	return successful
}

// Sends given requests, collects responses and returns them
// If an error occurs, the function will return the error immediately without continuing
func (v *Validator) Send(requests []http.Request) ([]*http.Response, error) {
	responses := []*http.Response{}
	client := &http.Client{}
	for _, request := range requests {
		r, err := client.Do(&request)
		if err != nil {
			return responses, err
		}
		responses = append(responses, r)
	}

	return responses, nil
}

// Returns true if status code is in range successful (2xx), otherwise returns false
func (v *Validator) isStatusSuccessful(status int) bool {
	return status == http.StatusOK ||
		status == http.StatusCreated ||
		status == http.StatusAccepted ||
		status == http.StatusNonAuthoritativeInfo ||
		status == http.StatusNoContent ||
		status == http.StatusResetContent ||
		status == http.StatusPartialContent ||
		status == http.StatusMultiStatus ||
		status == http.StatusAlreadyReported ||
		status == http.StatusIMUsed
}
