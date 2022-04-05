package common

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

func ProcessCmdLineArguments() (bool, int) {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 2 {
		fmt.Println("Usage: main shouldDisplay nrOfRequests")
		fmt.Println("Example:")
		fmt.Println("    go run name 1 100")
		panic("Wrong nr of args")
	}
	shouldDisplay := argsWithoutProg[0] == "1"
	nrOfRequests, _ := strconv.Atoi(argsWithoutProg[1])

	return shouldDisplay, nrOfRequests
}

func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
    return math.Sqrt(f), nil
}