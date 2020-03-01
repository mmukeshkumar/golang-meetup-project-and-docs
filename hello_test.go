package hello

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello World!"
	if got := Hello1(); got != want {
		t.Errorf("Hello1() = %q, want %q", got, want)
	}
}

func TestIsTopicValidFormat(t *testing.T) {
	validGCPTopicName := "projects/mytest-us-nonprod/topics/mytopic"
	want := true
	if got, _ := IsTopicValidFormat(validGCPTopicName); got != want {
		fmt.Println("Hello1() Got: ", got)
		fmt.Println("Hello1() want: ", want)
	}
	want = false
	invalidGCPTopicName := "projects/mytest-us-nonprod//invalid"
	if got, err := IsTopicValidFormat(invalidGCPTopicName); got != want {
		fmt.Println("Hello1() Got: ", got)
		fmt.Println("Hello1() want: ", want)
		fmt.Println("error: ", err)

	}

}
