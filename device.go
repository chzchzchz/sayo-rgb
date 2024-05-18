package sayo_hid

import (
	hid "github.com/sstallion/go-hid"
)

const Vid = uint16(0x8089)
const Pid = uint16(0x0004)

type Mode byte

const (
	ModeStatic Mode = iota
	ModeIndicator
	ModeBreath
	ModeBreathOnce
	ModeGradient
	ModeSwitch
	ModeSwitchOnce
	ModeFlash
	ModeFlashOnce
	ModeFade
	ModeFadeIn
)

type Event byte

const (
	EventNone Event = iota
	EventEventPressOnReleaseOff
	EventPressOffReleaseOn
	EventPressFadeOutReleaseFadeIn
	EventPressFadeInReleaseFadeOut
	EventPressToConfig
	EventReleaseToConfig
)

func makePacket(mode Mode, key int, r, g, b byte) []byte {
	pkt := make([]byte, 57)
	pkt[0] = 0x02
	pkt[1] = 0x10
	pkt[2] = 0x35
	pkt[3] = 0x01
	pkt[4] = byte(key)
	pkt[5] = byte(mode)
	pkt[6] = 0 // color mode
	pkt[7] = 0x00 // speed; 0xc0 is 8x speed
	pkt[8] = byte(EventNone)
	pkt[9] = r
	pkt[10] = g
	pkt[11] = b
	pkt[12] = 0 // keep light on
	pkt[13] = 0 // keep light off
	pkt[14] = 0 // color table
	chksum := byte(0)
	for _, v := range pkt {
		chksum = chksum + v
	}
	pkt[56] = chksum
	return pkt
}

type Device struct {
	hiddev *hid.Device
}

func NewDevice(path string) (*Device, error) {
	dev, err := hid.OpenPath(path)
	if err != nil {
		return nil, err
	}
	return &Device{dev}, nil
}

func (d *Device) Close() {
	d.hiddev.Close()
}

func (d *Device) Write(m Mode, key int, r, g, b byte) error {
	pkt := makePacket(m, key, r, g, b)
	_, err := d.hiddev.Write(pkt)
	return err
}
