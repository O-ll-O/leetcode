package main

import "testing"

func TestIsErrMask(t *testing.T) {
	maskSlice := strToInts("255.254.255.0")
	if !isErrMask(maskSlice) {
		t.Fail()
	}
}
