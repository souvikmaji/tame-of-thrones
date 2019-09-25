package main

import (
	"errors"
	"fmt"
	"os"
)

func showInstructions() {
	fmt.Println("1. Who is the ruler of Southeros?")
	fmt.Println("2. Allies of Ruler?")
	fmt.Println("3. Input Messages to kingdoms from King Shan. (Format: Name, \"msg\")")
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

func parseMessages() (string, string, error) {
	var kingdomName, msg string
	_, err := fmt.Scanf("%s %s", &kingdomName, &msg)
	if err != nil {
		fmt.Println(err)
		return kingdomName, msg, err
	}

	return kingdomName, msg, nil
}

func main() {

	kingShan := newKing("King Shan")

	for {
		option, err := parseOptions()
		if err != nil {
			fmt.Println(err)
			continue
		}

		switch option {
		case 1:
			kingShan.printIsRuler()
		case 2:
			kingShan.Kingdom.printAllies()
		case 3:
			kingdomName, msg, err := parseMessages()
			if err != nil {
				fmt.Println(err)
				continue
			}
			kingShan.Kingdom.makeAlliance(kingdomName, msg)
		}
	}
}
