package main

import "gorm-learn/route"

func main() {
	engine := route.InitRoute()

	engine.Run("192.168.124.7:1080") //开启
}
