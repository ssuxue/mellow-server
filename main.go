package main

import "memo/initData"

func main() {
	router := initData.SetupRouter()
	_ = router.Run(":8080")
}
