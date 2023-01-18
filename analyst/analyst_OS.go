package analyst

import (
	"fmt"
  "github.com/theGOURL/go_url/pkg/system"
  "github.com/theGOURL/go_url/pkg/system/analyzer"
)

// this function is the operating system analyst
// directing to the specific implementation for the operating system
func OSChecking() {
	switch analyzer.MyOS() {
	case "linux":
		fmt.Println(analyzer.MyOS());
		system.LinuxOS();
	case "windows":
		fmt.Println(analyzer.MyOS());
		system.WindowsOS();
	case "darwin":
		fmt.Println(analyzer.MyOS());
		system.MACOS();
	default:
		fmt.Println("UNDEFINIED");
	}

}
