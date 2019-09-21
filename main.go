package main

import (
	"errors"
	"fmt"
	"os"
)

func showInstructions() {
	fmt.Println("1. Who is the ruler of Southeros?")
	fmt.Println("2. Allies of Ruler?")
	fmt.Println("3. Input Messages to kingdoms from King Shan.")
	fmt.Print("> ")
}

func getOption(inputStream *os.File) (int, error) {
	var i int
	_, err := fmt.Scanf("%d", &i)

	if err != nil {
		return 0, err
	}

	return i, nil
}

func parseOptions() (int, error) {
	showInstructions()
	option, err := getOption(os.Stdin)
	if err != nil {
		return 0, err
	}

	if option < 1 || option > 3 {
		return 0, errors.New("invalid option")
	}

	return option, nil
}

func main() {

	// shan :=

	for {
		option, err := parseOptions()
		if err != nil {
			fmt.Println(err)
		}

		switch option {
		case 1:
			getRuler()
		case 2:
			getAllies()
		case 3:
			parseMessages()
		}
	}
}
