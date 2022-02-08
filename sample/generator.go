package sample

import (
	"github.com/duongnln96/go-grpc-practice/pb/pcbook"
	"github.com/golang/protobuf/ptypes"
)

func NewKeyboard() *pcbook.Keyboard {
	return &pcbook.Keyboard{
		Layout:  randomKeyBoardLayout(),
		Backlit: randomBool(),
	}
}

// NewCPU returns a new sample CPU
func NewCPU() *pcbook.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)

	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)

	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)

	cpu := &pcbook.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}

	return cpu
}

// NewGPU returns a new sample GPU
func NewGPU() *pcbook.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)

	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)
	memGB := randomInt(2, 6)

	gpu := &pcbook.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: &pcbook.Memory{
			Value: uint64(memGB),
			Unit:  pcbook.Memory_GIGABYTE,
		},
	}

	return gpu
}

// NewRAM returns a new sample RAM
func NewRAM() *pcbook.Memory {
	memGB := randomInt(4, 64)

	ram := &pcbook.Memory{
		Value: uint64(memGB),
		Unit:  pcbook.Memory_GIGABYTE,
	}

	return ram
}

// NewSSD returns a new sample SSD
func NewSSD() *pcbook.Storage {
	memGB := randomInt(128, 1024)

	ssd := &pcbook.Storage{
		Driver: pcbook.Storage_SSD,
		Memory: &pcbook.Memory{
			Value: uint64(memGB),
			Unit:  pcbook.Memory_GIGABYTE,
		},
	}

	return ssd
}

// NewHDD returns a new sample HDD
func NewHDD() *pcbook.Storage {
	memTB := randomInt(1, 6)

	hdd := &pcbook.Storage{
		Driver: pcbook.Storage_HDD,
		Memory: &pcbook.Memory{
			Value: uint64(memTB),
			Unit:  pcbook.Memory_TERABYTE,
		},
	}

	return hdd
}

// NewScreen returns a new sample Screen
func NewScreen() *pcbook.Screen {
	screen := &pcbook.Screen{
		SizeInch:   randomFloat32(13, 17),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}

	return screen
}

// NewLaptop returns a new sample Laptop
func NewLaptop() *pcbook.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)

	laptop := &pcbook.Laptop{
		Id:       randomID(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Ram:      NewRAM(),
		Gpus:     []*pcbook.GPU{NewGPU()},
		Storages: []*pcbook.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pcbook.Laptop_WeightKg{
			WeightKg: randomFloat64(1.0, 3.0),
		},
		PriceUsd:    randomFloat64(1500, 3500),
		ReleaseYear: uint32(randomInt(2015, 2019)),
		UpdatedAt:   ptypes.TimestampNow(),
	}

	return laptop
}
