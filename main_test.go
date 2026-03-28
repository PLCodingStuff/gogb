package main

import (
	"testing"
)

func TestNOP(t *testing.T) {
	RAM[0x0100] = byte(NOP)
	cpu := InitCPU()

	cycles := Step(cpu)
	if cpu.PC != 0x0102 {
		t.Errorf("cpu.PC = %04x, expected 0x0102", cpu.PC)
	} else if cycles != 4 {
		t.Errorf("Step({A: %04x, PC: %04x, Halted: %t}) = %d, expected 4", 0x00, 0x0100, false, cycles)
	}
}

func TestJP_a16(t *testing.T) {

}

func TestJP_a16_LittleEndian(t *testing.T) {

}

func TestLD_A_d8(t *testing.T) {

}

func TestLDH_a8_A(t *testing.T) {

}

func TestHALT(t *testing.T) {

}
