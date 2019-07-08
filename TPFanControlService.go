package main

// #cgo CFLAGS: -I${SRCDIR}/libs
// #cgo LDFLAGS: -L${SRCDIR}/libs -lTVicPort
// #include <windows.h>
// #include "TVicPort.h"
import "C"
import (
	"fmt"
	"time"
)

func main() {
	tryOpen := C.OpenTVicPort()
	checkOpenend := C.IsDriverOpened()
	if tryOpen == 0 || checkOpenend == 0 {
		fmt.Println("Something goes wrong.")
		return
	}
	isHardAccessAvailable := C.TestHardAccess()
	if isHardAccessAvailable != 0 {
		C.SetHardAccess(1)
	}

	// ok := WriteByteToEC(0x31, 0x0000)
	// fmt.Println(ok)
	// ok = WriteByteToEC(0x2F, 0x080)
	// fmt.Println(ok)

	for true {
		var sensors = [2]byte{}
		var fanState byte
		var fanspeedLowByte byte
		var fanspeedHighByte byte
		var fanspeed int
		ReadByteFromEC(0x78, &sensors[0])
		ReadByteFromEC(0x78+1, &sensors[1])
		ReadByteFromEC(0x2F, &fanState)
		ReadByteFromEC(0x84, &fanspeedLowByte)
		ReadByteFromEC(0x84+1, &fanspeedHighByte)
		fanspeed = int((int(fanspeedHighByte) << 8) | int(fanspeedLowByte))

		out := fmt.Sprintf("CPU: %d°, GPU: %d°, Fanspeed: %drpm, FanState: %d at", sensors[0], sensors[1], fanspeed, fanState)
		fmt.Println(out, time.Now().Format(time.Stamp))

		time.Sleep(time.Duration(5*1000) * time.Millisecond)
	}
}
