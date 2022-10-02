package main

import (
	"com.caiflower/commons/thirdpart/internal/config"
	"com.caiflower/commons/thirdpart/internal/route"
	"flag"
	"log"
)

func main() {
	flag.Parse()

	// 初始化配置信息
	if e := config.Init(); e != nil {
		log.Fatal("config init fail")
	}

	// 初始化路由
	if e := route.Init(); e != nil {
		log.Fatal("init route fail")
	}

}
