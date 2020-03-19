package main

import "fmt"

type Car struct {
	Brand           string
	ManufactureYear uint
	BootVolume      uint
	EngineStart     bool
	WindowsOpen     bool
	TrunkFilling    uint
}

func main() {

	car := Car{
		Brand:           "Mercedes",
		ManufactureYear: 2020,
		BootVolume:      50,
		EngineStart:     true,
		WindowsOpen:     false,
		TrunkFilling:    25,
	}

	trunk := Car{
		Brand:           "Zil",
		ManufactureYear: 2020,
		BootVolume:      250,
		EngineStart:     false,
		WindowsOpen:     true,
		TrunkFilling:    200,
	}

	fmt.Println("Engine start: ", car.EngineStart)
	fmt.Println("ManufactureYear: ", trunk.ManufactureYear)

	trunk.ManufactureYear = 2019
	fmt.Println("ManufactureYear: ", trunk.ManufactureYear)

}
