package main

import (
	"fmt"
	"os"
)

func showInstructions() {
	fmt.Println("1. Who is the ruler of Southeros?")
	fmt.Println("2. Allies of Ruler?")
	fmt.Println("3. Input Messages to kingdoms from King Shan.")
	fmt.Print("> ")
}

func rcvInput(inputStream *os.File) (int, error) {
	// reader := bufio.NewReader(inputStream)
	// optionsStr, err := reader.ReadString('\n')
	// if err != nil {
	// 	return 0, err
	// }

	// i, err := strconv.Atoi(optionsStr)
	// if err != nil {
	// 	fmt.Println("i: ", err)
	// 	return 0, err
	// }
	// fmt.Println("i: ", i)

	// return i, nil
	var i int
	_, err := fmt.Scanf("%d", &i)

	if err != nil {
		return 0, err
	}

	return i, nil
}

func parseInput() {
	showInstructions()
	option, err := rcvInput(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}

	if option < 1 || option > 3 {
		fmt.Println("invalid option")
		return
	}

	action(option)
}

func action(option int) {
	switch option {
	case 1:

	}
}

func main() {
	for {
		parseInput()
	}
}
