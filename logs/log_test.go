package logs

import "testing"

func TestLog(t *testing.T) {
	InitLog()
	Infof("a=%v, b=%v, c=%v", 1, 2, 3)
}
