package menu

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func DisplayMenuOptions(options []string) string {
	fmt.Println("\nSelect an option:")
	for i, opt := range options {
		fmt.Printf("%d. %s\n", i+1, opt)
	}
	fmt.Print("Choice: ")

	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	return strings.TrimSpace(choice)
}

func getNewEntityInput() string {
	fmt.Println("Enter new entity as JSON:")
	reader := bufio.NewReader(os.Stdin)
	json, _ := reader.ReadString('\n')

	json = strings.TrimSpace(json)
	return json
}
