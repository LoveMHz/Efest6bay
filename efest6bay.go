package main

import (
	"fmt"
	"math/rand"
	"encoding/hex"
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

func encodeMessage(bytes []byte) []byte {
    encodeByte := make([]byte, len(bytes) + 3)

    encodeByte[0] = 84
    encodeByte[1] = encodeByte[0] ^ byte(len(bytes) + 1)

    num := byte(rand.Int())
    encodeByte[2] = encodeByte[0] ^ num
    
    for i := 0; i < len(bytes); i++ {
        encodeByte[i + 3] = bytes[i] ^ num
    }
    
    return encodeByte
}

// TODO-dustin: Find a proper place for this function to live.
func decodeMessageTest() {
	test01 := []byte{0x54, 0x46, 0xdb, 0xcc, 0xc7, 0xbe, 0xb5, 0xbc, 0xa3, 0xbb, 0xa1, 0xbd, 0xbf, 0xd9, 0xa3, 0xbf, 0xe2, 0xce, 0xa3, 0xbe}
	test02 := []byte{0x54, 0x46, 0x4b, 0x2f, 0x2f, 0x6f, 0x7a, 0x6d, 0x33, 0x2c, 0x2e, 0x33, 0x5c, 0x57, 0x2d, 0x25, 0x2c, 0x33, 0x2b, 0x31}
	test03 := []byte{0x54, 0x46, 0x6b, 0x0d, 0x0f, 0x69, 0x13, 0x0f, 0x52, 0x7e, 0x13, 0x0e, 0x0f, 0x0f, 0x4f, 0x5a, 0x4d, 0x13, 0x0d, 0x0e}
	test04 := []byte{0x54, 0x46, 0x62, 0x1a, 0x75, 0x7e, 0x05, 0x0c, 0x05, 0x1a, 0x02, 0x18, 0x04, 0x06, 0x60, 0x1a, 0x06, 0x5b, 0x77, 0x1a}
	test05 := []byte{0x54, 0x46, 0xba, 0xdf, 0xde, 0xde, 0x9e, 0x8b, 0x9c, 0xc2, 0xdc, 0xdf, 0xc2, 0xad, 0xa6, 0xda, 0xd4, 0xdd, 0xc2, 0xda}
	test06 := []byte{0x54, 0x46, 0x17, 0x6d, 0x71, 0x73, 0x15, 0x6f, 0x73, 0x2e, 0x02, 0x6f, 0x72, 0x73, 0x73, 0x33, 0x26, 0x31, 0x6f, 0x71}
	test07 := []byte{0x54, 0x46, 0x2c, 0x49, 0x54, 0x3b, 0x30, 0x4d, 0x42, 0x4b, 0x54, 0x4c, 0x56, 0x4a, 0x48, 0x2e, 0x54, 0x48, 0x15, 0x39}
	test08 := []byte{0x54, 0x46, 0x19, 0x61, 0x7c, 0x7d, 0x7d, 0x3d, 0x28, 0x3f, 0x61, 0x7f, 0x7c, 0x61, 0x0e, 0x05, 0x7b, 0x77, 0x7e, 0x61}
	test09 := []byte{0x54, 0x46, 0x59, 0x39, 0x23, 0x3f, 0x3d, 0x5b, 0x21, 0x3d, 0x60, 0x4c, 0x21, 0x3c, 0x3d, 0x3d, 0x7d, 0x68, 0x7f, 0x21}
	test0A := []byte{0x54, 0x51, 0xae, 0xc9, 0xcb, 0xf7, 0xf0}

    fmt.Print( decodeMessage(test01) )
    fmt.Print( decodeMessage(test02) )
    fmt.Print( decodeMessage(test03) )
    fmt.Print( decodeMessage(test04) )
    fmt.Print( decodeMessage(test05) )
    fmt.Print( decodeMessage(test06) )
    fmt.Print( decodeMessage(test07) )
    fmt.Print( decodeMessage(test08) )
    fmt.Print( decodeMessage(test09) )
    fmt.Print( decodeMessage(test0A) )
    fmt.Print("Completed")
}

func main() {

}