package main

import "fmt"

func main() {
	/*
		interface (接口)

		當某個類型為這個 interface 中的所有方法提供實現，他被稱為 "實現interface"

		1. 當需要用 interface 對象時，可以任意實現類對象代替

	*/

	m1 := Mouse{"技嘉"}
	fmt.Println(m1.name)

	f1 := FlashDisk{"金士頓64G"}
	fmt.Println(f1.name)

	fmt.Println("-------------")

	testInterface(m1)
	fmt.Println("-------------")

	testInterface(f1)

	fmt.Println("-------------")

	var usb USB
	usb = m1
	usb.start()
	usb.end()

	// this is error
	// fmt.Println(usb.name)
}

// USB 定義
type USB interface {
	start() // USB 開始工作
	end()   // USB 結束工作
}

// Mouse 結構
type Mouse struct {
	name string
}

// FlashDisk 結構
type FlashDisk struct {
	name string
}

func (m Mouse) start() {
	fmt.Println(m.name, "滑鼠")
}

func (m Mouse) end() {
	fmt.Println("結束工作，安全退出")
}

func (f FlashDisk) start() {
	fmt.Println("準備開始工作，可以進行資料儲存")
}

func (f FlashDisk) end() {
	fmt.Println("可以退出工作囉....")

}

// 測試方法
func testInterface(usb USB) {
	usb.start()
	usb.end()
}
