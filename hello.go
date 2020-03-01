package hello // if you want to execute a package it should be named main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"os"
	"reflect"
	"regexp"
)

//constants
const (
	CounterMax  = 100
	IgnoreCase  = false
	Pi          = 3.14
	loopForever = true
	StartOfWeek = "Monday"
)

//variables
var (
	s string = "Hello World!"
	b bool   = false
)

//functions
func Hello1() string {
	return "Hello World!"
}

func increment(anInt int) {
	for i := 0; i < 100; i++ {
		anInt += 1
	}
}

func incrementByPointer(anInt *int) {
	for i := 0; i < 100; i++ {
		*anInt += 1
	}
}

func printInputArg() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s [filename]\n", os.Args[0])
		//os.Exit(1)
	}
}

//Demo errors, nil
var ErrInvalidTopicName = errors.New("Topic Name is not of expected format: projects/project_id/topics/topic_id")

func IsTopicValidFormat(topic string) (bool, error) {
	r, _ := regexp.Compile("projects\\/([a-zA-Z0-9\\W]+)\\/topics\\/([a-zA-Z0-9\\W]+)")
	match := r.FindStringSubmatch(topic)
	if match == nil {
		return false, ErrInvalidTopicName
	} else {
		return true, nil
	}
}

// program entry point
func main() {
	// Demo how to read input arguments to program
	printInputArg()

	fmt.Println("Call Hello1: " + Hello1())
	fmt.Println("print  constant CounterMax: ", CounterMax)
	fmt.Println("print  constant CounterMax type: ", reflect.TypeOf(CounterMax).String())
	fmt.Println("print  constant StartOfWeek type: ", reflect.TypeOf(StartOfWeek).String())
	fmt.Println("print  constant IgnoreCase: ", IgnoreCase)
	fmt.Println("print  constant Pi: ", Pi)

	fmt.Println("print  constant Pi type: ", reflect.TypeOf(Pi).String())

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println("My favorite number is", rand.Intn(10))

	s := "Test String"

	fmt.Println("print  type of s: ", reflect.TypeOf(s).String())

	l := len(s)
	fmt.Println("print  length of s: ", l)

	/*
	      Range. The range form of the for loop iterates over a slice or map. When ranging over a slice, two values are returned for each iteration.
	   The first is the index, and the second is a copy of the element at that index.
	*/
	for index, char := range s {
		fmt.Printf("character at index %d is %c\n", index, char)
	}

	//Demo map
	var isSpecialElementMap = map[string]bool{
		"address": true,
		"applet":  true,
		"area":    true,
	}

	for key, value := range isSpecialElementMap {
		fmt.Println("Key: ", key, " value: ", value)
	}

	// Demo if else statement also demo for doubling as an while loop
	counter := 0
	for loopForever {
		counter++
		if counter == CounterMax {
			fmt.Println("counter has reached max coun, so breaking now")
			break
		} else {

			// dont do anything
		}
	}

	// Demo errors and regex

	exampleGCPTopicName := "projects/mytest-us-nonprod/topics/mytopic"
	aBoolResult, err := IsTopicValidFormat(exampleGCPTopicName)
	fmt.Println("test isTopicValidFormatwith correctly formatted topic string: ", aBoolResult, err)
	exampleGCPTopicName = "projects/mytest-us-nonprod//invalid"
	aBoolResult, err = IsTopicValidFormat(exampleGCPTopicName)
	fmt.Println("test isTopicValidFormat with invalid formatted topic string: ", aBoolResult, err)

	//Demo arrays and slices
	//array
	a := [6]int{1, 2, 34, 4, 5}
	a[0] = 100
	fmt.Println(a)

	//slices
	aSlice := []int{2, 3, 5, 7, 11, 13}
	aSlice = append(aSlice, 100)
	fmt.Println(aSlice)

	// Demo pointers
	anInt := 100
	increment(anInt)
	fmt.Println("Value of anInt after incrementing is: ", anInt)
	incrementByPointer(&anInt)
	fmt.Println("Value of anInt is after incrementing by pointer: ", anInt)

	//Demo structs

	c := Cat{
		Name: "Zippy",
		Sex:  "F",
	}

	fmt.Println("Return of MakeSound: ", c.MakeSound())
	output := GetOutput(c)
	fmt.Println("Output: ", output)
}

type Cat struct {
	Name string
	Sex  string
}

func (c Cat) MakeSound() string {
	return "Meow!!"
}

func (c Cat) Sleep() {}

type Car struct {
	Make  string
	Model string
}

func (c Car) Drive() string {
	return "Driving!!"
}

type Automobile interface {
	SoundHorn() string
	Drive() string
	Brake() string
}

type Animal interface {
	MakeSound() string
	Sleep()
}

//A type switch lets you choose between types
func GetOutput(value interface{}) string {
	var output string
	switch v := value.(type) {
	case Automobile:
		output = v.Drive()
	case Animal:
		fmt.Println("type is Animal")
		output = v.MakeSound()
	default:
		fmt.Println("Unknown type", v)
	}
	return output
}

//Notes:
//Go provides C-style /* */ block comments and C++-style // line comments
