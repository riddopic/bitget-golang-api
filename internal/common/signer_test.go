package common

import (
	"fmt"
	"github.com/riddopic/bitget-golang-api/internal"
	"testing"
)

func TestSigner_Sign(t *testing.T) {
	signer := new(Signer)
	result := signer.Sign("GET", "www.bitget.com", "aaaa", internal.TimesStamp())
	fmt.Print(result)
}
