package main

import (
	"fmt"
	"regexp"
)

func main() {

	urls := "https://www.youxi369.com/s"

	s, _ := regexp.MatchString(`^/.*$`, urls)

	fmt.Println(s)

}
