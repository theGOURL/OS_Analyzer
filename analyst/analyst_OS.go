package analyst

import (
	"fmt"
)

// this function is the operating system analyst
// directing to the specific implementation for the operating system
func OSAnalyzer() {
	switch MyOS() {
	case "linux":
		fmt.Println(MyOS());
			LinuxOS();
	case "windows":
		fmt.Println(MyOS());
			WindowsOS();
	case "darwin":
		fmt.Println(MyOS());
			MACOS();
	default:
		fmt.Println("UNDEFINIED");
	}

}
