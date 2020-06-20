// Paul Wagner 2020/06/14 Module 2 Activity: findian.go
/*Write a program which prompts the user to enter a string.
The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
The program should print “Found!” if the entered string starts with the character ‘i’,
ends with the character ‘n’, and contains the character ‘a’.
The program should print “Not Found!” otherwise.
The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings:
“ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”.
The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”. */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var myString string

	fmt.Println("Enter a string, I'll try to find an 'i' as the first letter, 'n' as the last letter and 'a' in between.")
	//	fmt.Scanln(&myString)
	myReader := bufio.NewReader(os.Stdin)   // returns a Reader-object
	myString, _ = myReader.ReadString('\n') // read until end-of-line, ignore errors

	myString = strings.TrimSpace(myString) // get rid of nagging spaces.
	myString = strings.ToLower(myString)   // transform the whole line to lower case
	fmt.Println("Working on :", myString)

	if strings.HasPrefix(myString, "i") && strings.HasSuffix(myString, "n") && strings.Contains(myString, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
