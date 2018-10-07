
## Usage

Example: TCP packet

```
package main

import "github.com/subaruf/shkep/packet"

func main() {
  tcpFields := []packet.Field{
    packet.Field{16, "src port"},
    packet.Field{16, "dst port"},
    packet.Field{32, "seq number"},
    packet.Field{32, "ack number"},
    packet.Field{4, "offset"},
    packet.Field{6, "unused"},
    packet.Field{6, "ctrl flag"},
    packet.Field{16, "window size"},
    packet.Field{16, "checksum"},
    packet.Field{16, "urgent pointer"},
  }
  tcpPacket := packet.NewPacket(tcpFields, 4)
  tcpPacket.Show()
}

```
