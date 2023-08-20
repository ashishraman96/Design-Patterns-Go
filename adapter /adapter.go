package main

import "fmt"

// we will create a adapter pattern where
// target : inputUSB
// existing: macUSB
// client : usbCable
// adaptee: windowsUSB

// to add two unrelated objects or functionalitys

type USB interface {
	insertUSB()
}

type mac struct{}

func (m *mac) insertUSB() {
	fmt.Println("insert into mac USB")
}

type client struct{}

func (c *client) insertIntoComputer(i USB) {
	i.insertUSB()
}

func main() {
	m := &mac{}

	cl := &client{}
	cl.insertIntoComputer(m)

	w := &windows{}
	a := &adaptor{win: w}

	cl.insertIntoComputer(a)
}

// now lets add a functionality to adapt windows USB as well
type windows struct{}
type adaptor struct {
	win *windows
}

func (w *windows) insertIntoWindows() {
	fmt.Println("insert into windows USB")
}

func (a *adaptor) insertUSB() {
	a.win.insertIntoWindows()
}
