package util

import (
	"runtime/debug"

	"github.com/JackTJC/gmFS_backend/logs"
)

func RecoverFromPanic() {
	if err := recover(); err != nil {
		logs.Sugar.Fatalf("panic:%v, stack:%v", err, string(debug.Stack()))
	}
}
