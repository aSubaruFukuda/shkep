package packet

import (
	"fmt"
	"strconv"
	"strings"
)

type Field struct {
	Length int
	Name   string
}
type Packet struct {
	Fields   []Field
	BitWidth int
}

func NewPacket(fi []Field, bw int) *Packet {
	pck := Packet{fi, bw}
	return &pck
}
func (pck *Packet) writeFirstline() {
	fmt.Printf("|")
	byteStr := ""
	for i := 0; i < 32; i++ {
		if i < 10 {
			byteStr = strings.Repeat(" ", pck.BitWidth/2-1) + strconv.Itoa(i) + strings.Repeat(" ", pck.BitWidth-(pck.BitWidth/2-1)-1)
		} else {
			byteStr = strings.Repeat(" ", pck.BitWidth/2-1) + strconv.Itoa(i) + strings.Repeat(" ", pck.BitWidth-(pck.BitWidth/2-1)-2)
		}
		fmt.Printf("%s", byteStr)
		fmt.Printf("|")
	}
	//fmt.Printf("\n|%s|\n", strings.Repeat("-", pck.BitWidth*32+31))
	fmt.Printf("\n|%s|\n", pck.line())
}

func (pck *Packet) line() string {
	foo := strings.Repeat(strings.Repeat("-", pck.BitWidth)+"+", 32)
	return foo[:len(foo)-1]
}

func (pck *Packet) writeFields() {
	fmt.Printf("|")
	cnt := 0
	for _, v := range pck.Fields {
		firstHalfSpace := strings.Repeat(" ", (v.Length*(pck.BitWidth+1)-len(v.Name))/2)
		latterHalfSpace := strings.Repeat(" ", v.Length*(pck.BitWidth+1)-1-len(firstHalfSpace)-len(v.Name))
		fieldStr := firstHalfSpace + v.Name + latterHalfSpace + "|"
		for i := 0; i < len(fieldStr); i += (pck.BitWidth + 1) {
			fmt.Printf(fieldStr[i : i+pck.BitWidth+1])
			cnt++
			if cnt%32 == 0 && (i+pck.BitWidth+1) == len(fieldStr) {
				//fmt.Printf("\n|%s|\n|", strings.Repeat("-", pck.BitWidth*32+31))
				fmt.Printf("\n|%s|\n|", pck.line())
			} else if cnt%32 == 0 {
				fmt.Printf("\b|\n|")
			}
		}
	}
	fmt.Printf("\b")
}
func (pck *Packet) Show() {
	pck.writeFirstline()
	pck.writeFields()
}
