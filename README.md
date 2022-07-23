package main

import (
	"fmt"

	cache "github.com/AleksandrAgelskii/newcache"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42)
	userId, err := cache.Get("userId")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(userId)
	}

	err = cache.Delete("userId")
	if err != nil {
		fmt.Println(err.Error())
	}

}
