package main

import (
	"errors"
	"fmt"
	"log"
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

func parseMessages() (string, string, error) {
	var kingdomName, msg string
	_, err := fmt.Scanf("%s %s", &kingdomName, &msg)
	if err != nil {
		log.Fatalln(err)
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
			if kingShan.isRuler() {
				fmt.Println(kingShan.Name)
			} else {
				fmt.Println("NONE")
			}
		case 2:
			for _, allie := range kingShan.Kingdom.Allies {
				fmt.Printf("%s ", allie.Name)
			}
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
