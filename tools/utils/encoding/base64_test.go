package encoding

import (
	"testing"
)

func TestBase64EncodeAndDecode(t *testing.T) {
	k := "This is EOS forked from openIM"
	p := Base64Encode(k)
	l,err  := Base64Decode(p)
	if err != nil {
		t.Fatal("error decode")
	}
	if l != k {
		t.Fatal("not same")
	}
}