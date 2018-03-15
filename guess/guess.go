package guess

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func Run() {
	var num = rand.New(rand.NewSource(time.Now().UnixNano()))
	var s = num.Intn(100)
	fmt.Printf("i have a number between %d, please guess\nguess: ", 100)
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("oops, something got wrong....quit!\n")
			return
		}

		var replacer = strings.NewReplacer("\r", "", "\n", "")
		input = replacer.Replace(input)

		if guessNum, err := strconv.Atoi(strings.Replace(input, "\r", "", -1)); err != nil {
			fmt.Println(err)
			fmt.Printf("oops, please input an Integer\n")
		} else {
			switch {
			case guessNum > s:
				fmt.Printf("oops, greater...\nguess: ")
			case guessNum < s:
				fmt.Printf("oops, less...\nguess: ")
			default:
				fmt.Printf("bingo, ur got the right num! want to play again(y/n)?\n")
				return
			}
		}

	}
}
