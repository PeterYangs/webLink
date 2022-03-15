package main

import (
	"fmt"
	"github.com/PeterYangs/request/v2"
)

func main() {

	c := request.NewClient()

	//ct, err := c.R().GetToContent("https://www.youxi369.com/down/api/3-89580-pc")
	rsp, err := c.R().Get("https://www.youxi369.com/down/api/3-89580-pc")

	if err != nil {

		fmt.Println(err)

		return
	}

	//fmt.Println(rsp.Header())
	//fmt.Println(ct.ToString())

	fmt.Println(rsp.Header().Get("Content-Type"))

	//for s, strings := range rsp.Header() {
	//
	//	fmt.Println(s)
	//}

}
