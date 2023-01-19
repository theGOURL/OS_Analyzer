package analyst

import (
	"fmt"

	"github.com/theGOURL/web/browser"
	"github.com/theGOURL/system/commands"
	"github.com/theGOURL/URL/urls"
)

// this function should only be run on linux :)
func LinuxOS() {
	command := MyCommand()
	switch command {
	case "man":
		fmt.Println(`
    
                            WELCOME TO theGOURL CLI 💻
                               Commands Manual 🧠
                         -------------------------------------------------------
                        | ls       ->  run ls linux command                     |
                        | i        ->  run whoami linux command                 |
                        | hello    ->  say hello                                |
                        | firefox  ->  open firefox browser                     |
                        | github   ->  prints the Github   link on the console  |
                        | whatsapp ->  prints the Whatsapp link on the console  |
                        | google   ->  prints the Google   link on the console  |
                        | youtube  ->  prints the Youtube  link on the console  |
                         -------------------------------------------------------

      `);
      OSAnalyzer();
	case "ls":
		commands.Execute("ls")
	case "i":
		commands.Execute("whoami")
	case "hello":
		fmt.Println("Hello ", myOSUsername("[ERROR]: Cannot get OS username\n[RESPONSE]: "))
	case "firefox":
		browser.FIREFOX()
	case "github":
		fmt.Println("[CTRL CLICK] =>", "[", urls.Github(), "]")
	case "whatsapp":
		fmt.Println("[CTRL CLICK] =>", "[", urls.Whatsapp(), "]")
	case "google":
		fmt.Println("[CTRL CLICK] =>", "[", urls.Google(), "]")
	case "youtube":
		fmt.Println("[CTRL CLICK] =>", "[", urls.Youtube(), "]")
	default:
		fmt.Println("UNDEFINIED COMMAND")
	}

}
