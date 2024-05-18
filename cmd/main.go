package main

import (
	"flag"
	"math/rand"
	"time"

	sayo "github.com/chzchzchz/sayo-rgb"
)

func main() {
	pathFlag := flag.String("path", "/dev/hidraw4", "path to hid device")
	flag.Parse()
	d, err := sayo.NewDevice(*pathFlag)
	if err != nil {
		panic(err)
	}
	colors := [][3]byte{{0xff, 0, 0}, {0, 0xff, 0}, {0, 0, 0xff}}
	for {
		for row := 0; row < 6; row++ {
			for col := 0; col < 4; col++ {
				r := row //rand.Int()
				i := row*4 + col
				c := colors[r%len(colors)]
				if rand.Int()&1 == 0 {
					c = [3]byte{0, 0, 0}
				}
				d.Write(sayo.ModeSwitchOnce, i, c[0], c[1], c[2])
				time.Sleep(10 * time.Millisecond)
			}
		}
	}
}
