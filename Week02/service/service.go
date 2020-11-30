package main

import (
	"fmt"
	"learn-go/Week02/biz"
)

func main() {
	userBiz, err := biz.GetUserInfo(1)
	if err != nil {
		fmt.Printf("%+v\n", err)
	} else {
		println(userBiz)
	}

}
