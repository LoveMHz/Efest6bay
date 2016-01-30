package main

import (
	"fmt"
)

// Decodes Bluetooth LE message. Based on orginal 'Efest 6bay charger' Android app.
func decodeMessage(reciverBytes []byte) string {
    var decodeData string
    var oNum byte = reciverBytes[0] ^ reciverBytes[2]

    for i := 3; i < len(reciverBytes); i++ {
        decodeData += string(reciverBytes[i] ^ oNum)
    }

    return decodeData
}

func main() {

}