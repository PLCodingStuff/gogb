package main

const KB = 1024

var RAM [8 * KB]byte // 8 KB RAM

type Opcode byte

const (
	NOP    Opcode = 0x00
	JPa16  Opcode = 0xC3
	LDAn8  Opcode = 0x3E
	LDHa8A Opcode = 0xE0
	HATL   Opcode = 0x76
)

type CPU struct {
	A      byte
	PC     uint16
	Halted bool
}

func InitCPU() *CPU {
	cpu := CPU{A: 0x0000, PC: 0x0100, Halted: false}

	return &cpu
}

func Step(cpu CPU) uint {
	return 0
}

func main() {

}
