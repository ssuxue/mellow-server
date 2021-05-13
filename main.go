package main

import router "github.com/ssuxue/mellow-server/internal/router"

func main() {
	r := router.SetupRouter()
	_ = r.Run(":8080")
}
