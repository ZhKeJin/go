package main

import "fmt"

type Usb interface {
	start()
	stop()
}

type Phone struct{

}

func (p *Phone) stop() {
	fmt.Println("start : this is my phone ")
}

func (p Phone) start(){
	fmt.Println("start : this is my phone ")
}


type Car struct {

}

func (c Car) start() {
	panic("start : car  me")
}

func (c Car) stop() {
	panic("stop : car me")
}

type Computer struct{

}

func (c Computer) work(usb Usb) {
	usb.start()
	usb.stop()
}
func main() {
    computer := Computer{}
	phone := Phone{}
	car := Car{}

	computer.work(&phone)
	fmt.Println()

	computer.work(car)
}
