// chap1/manual-parse/main.go
package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

//This is an object of type config
type config struct {
	numTimes   int
	printUsage bool
}

var usageString = fmt.Sprintf(`Usage: %s <integer> [-h|--help]
A greeter application which prints the name you entered <integer> number
of times.`, os.Args[0])

func printUsage(w io.Writer) {
	fmt.Fprintf(w, usageString)
}

//The io.reader and io.writer is declared this way to have control the types of variable
//uses the Reader interface in
func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "Your name please? Press the Enter key when done.\n"
	fmt.Fprintf(w, msg)
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("You didn't enter your name")
	}
	return name, nil
}

func parseArgs(args []string) (config, error) {
	var numTimes int
	var err error
	c := config{}

	//validates that num of arguments is correct
	if len(args) != 1 {
		return c, errors.New("Invalid number of arguments")
	}
	//validates if the user wants helps
	if args[0] == "-h" || args[0] == "--help" {
		c.printUsage = true
		return c, nil
	}
	//strconv with the atoi function converts from a string to an integer
	numTimes, err = strconv.Atoi(args[0])
	if err != nil {
		return c, err
	}
	c.numTimes = numTimes
	return c, nil
}

//validate num of repetion of name is greater than 0
func validateArgs(c config) error {
	if !(c.numTimes > 0) {
		return errors.New("Must specify a number greater than 0")
	}
	return nil
}

func runCmd(r io.Reader, w io.Writer, c config) error {
	if c.printUsage {
		printUsage(w)
		return nil
	}
	name, err := getName(r, w)
	if err != nil {
		return err
	}
	greetUser(c, name, w)
	return nil
}

func greetUser(c config, name string, w io.Writer) {
	msg := fmt.Sprintf("Nice to meet you %s\n", name)
	for i := 0; i < c.numTimes; i++ {
		fmt.Fprintf(w, msg)
	}
}

//to run code use go run main.go num
//to compile app 
//go build -o application

func main() {
	c, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}
	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}
	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}
