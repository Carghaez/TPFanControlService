package main

// #cgo CFLAGS: -I${SRCDIR}/libs
// #cgo LDFLAGS: -L${SRCDIR}/libs -lTVicPort
// #include "TVicPort.h"
import "C"

import (
	"time"
)

// Registers of the embedded controller

// ECDataPort EC data io-port 0x62
const ECDataPort = 0x1600

// ECCtrlPort EC control io-port 0x66
const ECCtrlPort = 0x1604

// Embedded controller status register bits

// ECStatOBF Output buffer full
const ECStatOBF = 0x01

// ECStatIBF Input buffer full
const ECStatIBF = 0x02

// ECStatCmd Last write was a command write (0=data)
const ECStatCmd = 0x08

// ECCtrlPortRead Embedded controller commands (write to ECCtrlPort to initiate read/write operation)
const ECCtrlPortRead = 0x80

// ECCtrlPortWrite Embedded controller commands (write to ECCtrlPort to initiate read/write operation)
const ECCtrlPortWrite = 0x81

// ECCtrlPortQuery Embedded controller commands (write to ECCtrlPort to initiate read/write operation)
const ECCtrlPortQuery = 0x84

// ReadByteFromEC read a byte from the embedded controller (EC) via port io
func ReadByteFromEC(offset int, pdata *byte) int {
	var data byte
	iOK := false
	iTimeout := 100
	iTimeoutBuf := 1000
	iTime := 0
	iTick := 10

	for iTime < iTimeoutBuf {
		data = byte(C.ReadPort(ECCtrlPort) & 0xff) // or timeout iTimeoutBuf = 1000
		if data&(ECStatIBF|ECStatOBF) == 0 {
			break
		}
		time.Sleep(time.Duration(iTick) * time.Millisecond)
		iTime += iTick
	}

	if data&ECStatOBF != 0 {
		_ = byte(C.ReadPort(ECDataPort)) //clear OBF if full
	}

	C.WritePort(ECCtrlPort, ECCtrlPortRead) // tell 'em we want to "READ"

	iTime = 0
	for iTime < iTimeout { // wait for IBF and OBF to clear
		data = byte(C.ReadPort(ECCtrlPort) & 0xff)
		if data&(ECStatIBF|ECStatOBF) == 0 {
			iOK = true
			break
		}
		time.Sleep(time.Duration(iTick) * time.Millisecond)
		iTime += iTick
	} // try again after a moment

	if !iOK {
		return 0
	}
	iOK = false

	C.WritePort(ECDataPort, C.uchar(offset)) // tell 'em where we want to read from

	if data&ECStatOBF == 0 {
		iTime = 0
		for iTime < iTimeout { // wait for OBF
			data = byte(C.ReadPort(ECCtrlPort) & 0xff)
			if data&ECStatOBF != 0 {
				iOK = true
				break
			}
			time.Sleep(time.Duration(iTick) * time.Millisecond)
			iTime += iTick
		} // try again after a moment
		if !iOK {
			return 0
		}
	}

	*pdata = byte(C.ReadPort(ECDataPort))

	return 1
}

// WriteByteToEC write a byte to the embedded controller (EC) via port io
func WriteByteToEC(offset int, NewData byte) int {
	var data byte
	iOK := false
	iTimeout := 100
	iTimeoutBuf := 1000
	iTime := 0
	iTick := 10

	for iTime < iTimeoutBuf {
		data = byte(C.ReadPort(ECCtrlPort) & 0xff)
		if data&(ECStatIBF|ECStatOBF) == 0 {
			break
		}
		time.Sleep(time.Duration(iTick) * time.Millisecond)
		iTime += iTick
	}

	if data&ECStatOBF != 0 {
		_ = byte(C.ReadPort(ECDataPort)) //clear OBF if full
	}

	iTime = 0
	for iTime < iTimeout { // wait for IOBF to clear
		data = byte(C.ReadPort(ECCtrlPort) & 0xff)
		if data&ECStatOBF == 0 {
			iOK = true
			break
		}
		time.Sleep(time.Duration(iTick) * time.Millisecond)
		iTime += iTick
	} // try again after a moment

	if !iOK {
		return 0
	}
	iOK = false

	C.WritePort(ECCtrlPort, ECCtrlPortWrite) // tell 'em we want to "WRITE"

	iTime = 0
	for iTime < iTimeout { // wait for IOBF to clear
		data = byte(C.ReadPort(ECCtrlPort) & 0xff)
		if data&(ECStatIBF|ECStatOBF) == 0 {
			iOK = true
			break
		}
		time.Sleep(time.Duration(iTick) * time.Millisecond)
		iTime += iTick
	} // try again after a moment

	if !iOK {
		return 0
	}
	iOK = false

	C.WritePortL(ECDataPort, C.ulong(offset)) // tell 'em where we want to write to

	iTime = 0
	for iTime < iTimeout { // wait for IOBF to clear
		data = byte(C.ReadPort(ECCtrlPort) & 0xff)
		if data&(ECStatIBF|ECStatOBF) == 0 {
			iOK = true
			break
		}
		time.Sleep(time.Duration(iTick) * time.Millisecond)
		iTime += iTick
	} // try again after a moment

	if !iOK {
		return 0
	}
	iOK = false

	C.WritePort(ECDataPort, C.uchar(NewData)) // tell 'em what we want to write there

	return 1
}
