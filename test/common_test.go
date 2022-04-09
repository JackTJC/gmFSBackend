package test

import (
	"runtime"
	"testing"
)

func TestCaller(t *testing.T) {
	_, b, _, _ := runtime.Caller(0)
	t.Log(b)
}
