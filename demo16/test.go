package main

import "fmt"

func main() {
	var name string
	var l string

	scanf, err2 := fmt.Scanf("%s %s", &name, &l)
	if err2 != nil {
		return
	}
	fmt.Println(scanf)

	scanln, err := fmt.Scanln(&name, &l)
	fmt.Println(scanln)
	n, err := fmt.Sscanln("GO LANG", &name, &l)
	if err != nil {
		return
	}
	fmt.Println(n)
	sscanf, err := fmt.Sscanf("GO LANG", "%s %s", &name, &l)

	if err != nil {
		return

	}
	fmt.Println(sscanf)
}
