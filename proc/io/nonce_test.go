package io

import "testing"

func TestGetNonceReturnsStringOfLength(t *testing.T) {
	nonce := getNonce(10)
	if len(nonce) != 10 {
		t.Errorf("Nonce %v should be 10 chars long, but is %v chars long.\n", nonce, len(nonce))
	}

	nonce = getNonce(43)
	if len(nonce) != 43 {
		t.Errorf("Nonce %v should be 43 chars long, but is %v chars long.\n", nonce, len(nonce))
	}
}
