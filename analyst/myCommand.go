package analyst

import "fmt"
//this function executes the command that the user sets, 
//returning the command
func MyCommand()string{
	var command string;
		fmt.Printf("🧠 ");
		fmt.Scanf("%s",&command);
		return command;
}
