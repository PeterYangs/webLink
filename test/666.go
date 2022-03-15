package main

import (
	"fmt"
	"net/url"
)

func main() {

	u, err := url.Parse("https://www.925g.com")

	if err != nil {

		fmt.Println(err)

		return
	}

	fmt.Println(u.Scheme)

}
