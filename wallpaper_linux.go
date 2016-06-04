package wallpaper

import "os/exec"

func Set(name string) {
	exec.Command("feh", "--bg-center", name)
}
