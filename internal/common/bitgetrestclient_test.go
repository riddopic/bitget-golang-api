package common

import (
	"github.com/riddopic/bitget-golang-api/internal"
	"testing"
)

func TestBitgetRestClient_HttpExecuter(t *testing.T) {
	restClient := new(BitgetRestClient).Init()
	params := internal.NewParams()
	params["productType"] = "umcbl"
	resp, _ := restClient.DoGet("/api/mix/v1/account/accounts", params)

	println(resp)
}
