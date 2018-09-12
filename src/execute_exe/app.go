package main

import "os/exec"

func main() {
	path := "D:/Program Files (x86)/Adobe/Photoshop CS6/ps.exe"
	run_ps(path)
}

func run_ps(path string)  {
	command := exec.Command(path)
	re := command.Run()
	print(re)
}