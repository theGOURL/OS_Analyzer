package analyst

import (
	"runtime"
)
//this function returns the operating system name
func MyOS()string{
	return runtime.GOOS;
}
