package main

import "github.com/toribecher/gojos/gojos"

func main() {
	a := gojos.App{}

	a.Initialize()

	a.Run(":8000")
}
