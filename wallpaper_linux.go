package wallpaper

import "os/exec"

func Set(name string) error {
	cmd := exec.Command("feh", "--bg-center", name)
	err := cmd.Run()
	return err
}
