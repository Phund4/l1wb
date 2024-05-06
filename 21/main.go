package main

import "fmt"

type Client struct {}

// Функция подключения лайтнинга к компьютеру
func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
    fmt.Println("Client inserts Lightning connector into computer.")
    com.InsertIntoLightningPort()
}

type Computer interface {
    InsertIntoLightningPort()
}

type Mac struct {}

// Функция вставки в лайтнинг порт
func (m *Mac) InsertIntoLightningPort() {
    fmt.Println("Lightning connector is plugged into mac machine.")
}

type Windows struct{}

// Функция вставки в usb порт
func (w *Windows) insertIntoUSBPort() {
    fmt.Println("USB connector is plugged into windows machine.")
}

// Адаптер windows
type WindowsAdapter struct {
    windowMachine *Windows
}

// При помощи адаптера связываем lightning и usb
func (w *WindowsAdapter) InsertIntoLightningPort() {
    fmt.Println("Adapter converts Lightning signal to USB.")
    w.windowMachine.insertIntoUSBPort()
}

func main() {
	client := &Client{}
    mac := &Mac{}

    client.InsertLightningConnectorIntoComputer(mac)

    windowsMachine := &Windows{}
    windowsMachineAdapter := &WindowsAdapter{
        windowMachine: windowsMachine,
    }

    client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
