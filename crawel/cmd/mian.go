package main

import "crawel/crawel"

func main() {
	//url := "http://gityuan.com/archive/"
	url1 := "https://blog.csdn.net/Innost"

	//crawel.Getartices(url)
	crawel.GetarticesForCSDN(url1)
}
