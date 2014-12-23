package main

import (
	"fmt"
	//"user/newmath"
	// "strconv"
	// "sort"
)

type USB interface {
	Name() string
	Connect()
}

type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect:", pc.name)
}

func main() {
	var a USB
	a = PhoneConnecter{"PhoneConnecter"}
	a.Connect()
	Disconnect(a)
}

func Disconnect(usb interface{}) {
	if pc, ok := usb.(PhoneConnecter); ok {
		fmt.Println("Disconnected", pc.name)
		return
	}

	switch v: = usb.(type){
		case PhoneConnecter:
			fmt.Print("dddd")
	}
	fmt.Println("Unknown device. ")

}
