package main

import (
	"fmt"
	"os"

	"github.com/DDRBoxman/wallpaper"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please specify a file")
		return
	}
	wallpaper.Set(os.Args[1])
}
