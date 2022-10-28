package util

import "testing"

func TestStrToBytes(t *testing.T) {
	expect := "zxlwansui"
	bytes := StrToBytes(expect)
	if string(bytes) != expect {
		t.Error("unexpect")
	}
	t.Log(bytes)
}
