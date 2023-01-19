package analyst

import "fmt"
//this function executes the command that the user sets, 
//returning the command
func MyCommand()string{
	var command string;
  username := myOSUsername("root mode");
		fmt.Printf("%s🧠 ",username);
		fmt.Scanf("%s",&command);
		return command;
}
