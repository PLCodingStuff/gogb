package main

import "fmt"

type Opcode byte

const (
	NOP    Opcode = 0x00
	JPa16  Opcode = 0xC3
	LDAn8  Opcode = 0x3E
	LDHa8A Opcode = 0xE0
	HALT   Opcode = 0x76
)

const KB = 1024

var RAM [8 * KB]byte // 8 KB RAM

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
	opcode := Opcode(RAM[cpu.PC])
	cpu.PC++

	switch opcode {
	case NOP:
		return 4
	case JPa16:
		lo := RAM[cpu.PC]
		hi := RAM[cpu.PC+1]
		cpu.PC = uint16(hi)<<8 | uint16(lo)
		return 16
	case LDAn8:
		cpu.A = RAM[cpu.PC]
		return 8
	case LDHa8A:
		offset := RAM[cpu.PC]
		RAM[0xFF00+uint16(offset)] = byte(cpu.A)
		return 12
	case HALT:
		cpu.Halted = true
		return 4
	default:
		panic(fmt.Sprintf("Unknown opcode: 0x%02X", opcode))
	}
}

func main() {

}
