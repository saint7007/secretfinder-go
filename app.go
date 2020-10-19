package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {

	data, err := ioutil.ReadFile("private.key")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
	match, _ := regexp.MatchString("([-]+BEGIN [^\\s]+ PRIVATE KEY[-]+[\\s]*[^-]*[-]+END [^\\s]+ PRIVATE KEY[-]+)", string(data))
	fmt.Println(match)

}
