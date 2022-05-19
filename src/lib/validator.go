package lib

import (
	"net/http"

	"github.com/spf13/cobra"
)

type Validator struct{}

func (v *Validator) Send(requests []http.Request) []*http.Response {
	// TODO: collect requests and print ALL results on console, rather than exiting at every failed request
	responses := []*http.Response{}
	client := &http.Client{}
	for i := range requests {
		r, err := client.Do(&requests[i])
		cobra.CheckErr(err)
		responses = append(responses, r)
	}

	return responses
}
