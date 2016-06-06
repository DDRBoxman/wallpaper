# wallpaper

[![travis build status](https://api.travis-ci.org/DDRBoxman/wallpaper.svg)](https://travis-ci.org/DDRBoxman/wallpaper)
[![Appveyor Build status](https://ci.appveyor.com/api/projects/status/o66hiqhib9b43k6a?svg=true)](https://ci.appveyor.com/project/DDRBoxman/wallpaper)


Cross platform Go library to set the wallpaper

```
package main

import (
	"github.com/DDRBoxman/wallpaper"
)

func main() {
	wallpaper.Set("/Users/User/Pictures/test.png")
}
```
