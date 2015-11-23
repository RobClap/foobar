package main

import (
	"./easier"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const PATH = "./sentences-trimmed"
const MAX_ATTEMPTS = 3

var console *bufio.Reader
var lines []string

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
		fmt.Println("FATAL: unable to open file")
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func welcome(username string) {
	fmt.Println("    Hello " + username + "!")
	fmt.Println(`
	Welcome to the easy crypto tutorial, section 1
	The rules for this part are simple:`)
	rules()
	fmt.Println(`
	To play the game it is suggested to use a terminal
	of horizontal size greater than 75 characters and
	horizontal size greater than 50.
	To check your current size run "stty size" in the
	shell prompt.
	`)

	fmt.Println(`
			HAVE FUN!
			`)
}
func rules() {
	fmt.Println(`
	I will give you a string that I want and a character set
	and you have 3 tries to give me a string that,
	after the proper encoding of the level, will 
	become my string.

	Here is an example:
	I want "Hello!"
	and the encoding is shifting the word by one position
	you have to send me "ello!H".
	You can use hints by typing a single "H" as input
	but beware, hints should be used ONLY if you
	think you can't get past the problem without
	them.`)
}
func getNewString() string {
	return strings.TrimSpace(lines[rand.Int()%len(lines)])
}
func readConsole() string {
	//TODO handle err
	text, _ := console.ReadString('\n')
	return text
}
func riddleMe(cur_level level) bool { // title string, charset string, encrypt func(string) string) bool {
	fmt.Println("Welcome to " + cur_level.title)
	fmt.Println("The character set I am considering is: ")
	fmt.Println(cur_level.charset)
	currentLine := getNewString()
	fmt.Println("The string I expect is: \n" + currentLine)
	for i := 0; i < MAX_ATTEMPTS; i++ {
		fmt.Printf("Attempts left: %d\n", MAX_ATTEMPTS-i)
		fmt.Print("?> ")
		input := strings.TrimSpace(readConsole())
		if input == "H" {
			fmt.Println("The hint is: ")
			fmt.Println(cur_level.hint)
			i--
			continue
		}
		answ := cur_level.encrypt(input)
		if answ == currentLine {
			fmt.Println("Well done!\n")
			return true
		}
		fmt.Println("That is not what I wanted! I wanted:\n" + currentLine + "\nbut I got:\n" + answ)
		fmt.Println("Type H to have the hint")
		fmt.Println()
	}
	fmt.Println("OH DARN! You finished your attempts!")
	return false
}

type level struct {
	title   string
	charset string
	hint    string
	encrypt func(string) string
}

var levels = []level{
	{"level0, actually not a level at all", easier.PRINTABLE_ASCII, "Just give me back what i want, be my echo", nil},
	{"level1, easy as eating a salad!", easier.PRINTABLE_ASCII, "More precisely: a Caesar salad", nil},
	{"level2, harder, but it is still a classic one.", easier.PRINTABLE_ASCII, "try typing in the alphabet, or 5 times the same character", nil},
}

//To understand this see Closures in golang
func levelGenerator(index int) func(string) string {
	switch index {
	case 0:
		return func(toencode string) string {
			return toencode
		}
	case 1:
		key := rand.Int() % len(easier.PRINTABLE_ASCII)
		return func(toencode string) string {
			return easier.Rot(toencode, easier.PRINTABLE_ASCII, key)
		}
	case 2:
		subs := make([]rune, len(easier.PRINTABLE_ASCII))
		perm := rand.Perm(len(easier.PRINTABLE_ASCII))
		for i, v := range perm {
			subs[v] = rune(easier.PRINTABLE_ASCII[i])
		}
		substiute_alpha := string(subs)
		return func(toencode string) string {
			return easier.Substitute(toencode, easier.PRINTABLE_ASCII, substiute_alpha)
		}
	}
	return nil
}
func Play(username string) {
	rand.Seed(time.Now().Unix())
	lines, _ = readLines(PATH)
	welcome(username)
	console = bufio.NewReader(os.Stdin)
	for index, level := range levels {
		level.encrypt = levelGenerator(index)
		if !riddleMe(level) {
			return
		}
	}
}
func main() {
	//TODO implement login
	Play("default")
}
