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
	Fields    []Field
	ByteWidth int
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
			byteStr = strings.Repeat(" ", pck.ByteWidth/2-1) + strconv.Itoa(i) + strings.Repeat(" ", pck.ByteWidth-(pck.ByteWidth/2-1)-1)
		} else {
			byteStr = strings.Repeat(" ", pck.ByteWidth/2-1) + strconv.Itoa(i) + strings.Repeat(" ", pck.ByteWidth-(pck.ByteWidth/2-1)-2)
		}
		fmt.Printf("%s", byteStr)
		fmt.Printf("|")
	}
	fmt.Printf("\n|%s|\n", strings.Repeat("-", pck.ByteWidth*32+31))
}
func (pck *Packet) writeFreeFields() {
	cnt := 0
	fmt.Printf("|")
	for _, v := range pck.Fields {
		for i := 0; i < v.Length; i++ {
			if i == v.Length-1 {
				fmt.Printf("%s", strings.Repeat(" ", pck.ByteWidth)+"|")
			} else {
				fmt.Printf("%s", strings.Repeat(" ", pck.ByteWidth+1))
			}
			cnt++
			if cnt == 32 {
				fmt.Printf("\n|")
				cnt = 0
			}
		}
	}
}
func (pck *Packet) writeFields() {
	cnt := 0
	fmt.Printf("|")
	fieldStr := ""
	for _, v := range pck.Fields {
		fieldStr = strings.Repeat(" ", (v.Length*(pck.ByteWidth+1)-len(v.Name))/2) + v.Name + strings.Repeat(" ", (v.Length*(pck.ByteWidth+1)-(v.Length*(pck.ByteWidth+1)-len(v.Name))/2))
		for i := 0; i < (v.Length * (pck.ByteWidth + 1)); i += (pck.ByteWidth + 1) {
			if (i + pck.ByteWidth + 1) == (v.Length * (pck.ByteWidth + 1)) {
				fmt.Printf(fieldStr[i : i+pck.ByteWidth])
				fmt.Printf("|")
			} else {
				fmt.Printf(fieldStr[i : i+pck.ByteWidth+1])
			}
			cnt++
			if cnt == 32 {
				fmt.Printf("\n|%s|\n|", strings.Repeat("-", pck.ByteWidth*32+31))
				cnt = 0
			}
		}
	}
	fmt.Printf("\b")
}
func (pck *Packet) Show() {
	pck.writeFirstline()
	pck.writeFields()
}
