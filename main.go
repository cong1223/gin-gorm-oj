package main

import (
	"gin_orm_oj/router"
)

func main() {
	r := router.Router()

	r.Run()
}
