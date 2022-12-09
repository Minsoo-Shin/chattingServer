package test

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func Test_ReadInputStdin(t *testing.T) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	fmt.Println("Enter text: ")
	text2 := ""
	fmt.Scanln(text2)
	fmt.Println(text2)

	ln := ""
	fmt.Sscanln("%v", ln)
	fmt.Println(ln)
}
