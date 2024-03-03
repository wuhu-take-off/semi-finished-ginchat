package main

import (
	"fmt"
	"ginchat/router"
	"ginchat/utils"
	"strconv"
	"time"
)

func test() {
	fmt.Println(strconv.FormatInt(time.Now().Unix(), 10))
}
func main() {
	//test()
	utils.InitConfig()
	utils.InitRedis()
	defer utils.CloseRedis()
	utils.InitMySQL()

	////
	router.Router().Run()
}
