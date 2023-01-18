package analyst

import (
	"fmt"
    "github.com/theGOURL/go_url/pkg/system"
)

// this function is the operating system analyst
// directing to the specific implementation for the operating system
func OSAnalyzer() {
	switch MyOS() {
	case "linux":
		fmt.Println(MyOS());
		system.LinuxOS();
	case "windows":
		fmt.Println(MyOS());
		system.WindowsOS();
	case "darwin":
		fmt.Println(MyOS());
		system.MACOS();
	default:
		fmt.Println("UNDEFINIED");
	}

}
