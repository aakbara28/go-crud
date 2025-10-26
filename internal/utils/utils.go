package utils

import (
	"runtime"
	"strings"
)

func GetFuncName() string {
	pc, _, _, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	name := fn.Name()
	
	// Extract just the function name from the full path
	parts := strings.Split(name, ".")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return name
}