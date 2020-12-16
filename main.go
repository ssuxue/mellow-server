package main

import "memo/initData"

func main() {
	// initData.SetCasbin() // TODO Casbin策略
	router := initData.SetupRouter()
	_ = router.Run(":8080")
}
