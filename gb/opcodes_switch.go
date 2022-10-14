package gb

import (
	"fmt"
)

var Call []string
var ticks int

func (goboy *Goboy) ExecuteOpcode(opcode byte) int {
	ticks = OpcodeCycles[opcode] * 4
	goboy.OpcodeSwitch(opcode)
	return ticks
}

func (goboy *Goboy) OpcodeSwitch(opcode byte) int {

	switch opcode {
	// 8bit LD instructions
	case 0x3E:
		goboy.setA(goboy.getParameter8Bit())
		return 8
	case 0x06:
		goboy.setB(goboy.getParameter8Bit())
		return 8
	case 0x0E:
		goboy.setC(goboy.getParameter8Bit())
		return 8
	case 0x16:
		goboy.setD(goboy.getParameter8Bit())
		return 8
	case 0x1E:
		goboy.setE(goboy.getParameter8Bit())
		return 8
	case 0x26:
		goboy.setH(goboy.getParameter8Bit())
		return 8
	case 0x2E:
		goboy.setL(goboy.getParameter8Bit())
		return 8
	case 0x02:
		goboy.WriteMemory(goboy.CPU.GetBC(), goboy.CPU.Registers.A)
		return 8
	case 0x12:
		goboy.WriteMemory(goboy.CPU.GetDE(), goboy.CPU.Registers.A)
		return 8
	case 0xEA:
		goboy.WriteMemory(goboy.getParameter16Bit(), goboy.CPU.Registers.A)
		return 16
	case 0xF2:
		goboy.setA(goboy.ReadMemory(0xFF00 + uint16(goboy.CPU.Registers.C)))
		return 8
	case 0xE2:
		goboy.WriteMemory(0xFF00+uint16(goboy.CPU.Registers.C), goboy.CPU.Registers.A)
		return 8
	// register A LD
	case 0x7F:
		goboy.setA(goboy.CPU.Registers.A)
		return 4
	case 0x78:
		goboy.setA(goboy.CPU.Registers.B)
		return 4
	case 0x79:
		goboy.setA(goboy.CPU.Registers.C)
		return 4
	case 0x7A:
		goboy.setA(goboy.CPU.Registers.D)
		return 4
	case 0x7B:
		goboy.setA(goboy.CPU.Registers.E)
		return 4
	case 0x7C:
		goboy.setA(goboy.CPU.Registers.H)
		return 4
	case 0x7D:
		goboy.setA(goboy.CPU.Registers.L)
		return 4
	case 0x0A:
		goboy.setA(goboy.ReadMemory(goboy.CPU.GetBC()))
		return 8
	case 0x1A:
		goboy.setA(goboy.ReadMemory(goboy.CPU.GetDE()))
		return 8
	case 0x7E:
		goboy.setA(goboy.ReadMemory(goboy.CPU.GetHL()))
		return 8
	case 0xFA:
		goboy.setA(goboy.ReadMemory(goboy.getParameter16Bit()))
		return 8
	// register B LD
	case 0x40:
		goboy.setB(goboy.CPU.Registers.B)
		return 4
	case 0x41:
		goboy.setB(goboy.CPU.Registers.C)
		return 4
	case 0x42:
		goboy.setB(goboy.CPU.Registers.D)
		return 4
	case 0x43:
		goboy.setB(goboy.CPU.Registers.E)
		return 4
	case 0x44:
		goboy.setB(goboy.CPU.Registers.H)
		return 4
	case 0x45:
		goboy.setB(goboy.CPU.Registers.L)
		return 4
	case 0x46:
		goboy.setB(goboy.ReadMemory(goboy.CPU.GetHL()))
		return 8
	case 0x47:
		goboy.setB(goboy.CPU.Registers.A)
		return 4
	// register C LD
	case 0x48:
		goboy.setC(goboy.CPU.Registers.B)
		return 4
	case 0x49:
		goboy.setC(goboy.CPU.Registers.C)
		return 4
	case 0x4A:
		goboy.setC(goboy.CPU.Registers.D)
		return 4
	case 0x4B:
		goboy.setC(goboy.CPU.Registers.E)
		return 4
	case 0x4C:
		goboy.setC(goboy.CPU.Registers.H)
		return 4
	case 0x4D:
		goboy.setC(goboy.CPU.Registers.L)
		return 4
	case 0x4E:
		goboy.setC(goboy.ReadMemory(goboy.CPU.GetHL()))
		return 8
	case 0x4F:
		goboy.setC(goboy.CPU.Registers.A)
		return 4
	// register D LD
	case 0x50:
		goboy.setD(goboy.CPU.Registers.B)
		return 4
	case 0x51:
		goboy.setD(goboy.CPU.Registers.C)
		return 4
	case 0x52:
		goboy.setD(goboy.CPU.Registers.D)
		return 4
	case 0x53:
		goboy.setD(goboy.CPU.Registers.E)
		return 4
	case 0x54:
		goboy.setD(goboy.CPU.Registers.H)
		return 4
	case 0x55:
		goboy.setD(goboy.CPU.Registers.L)
		return 4
	case 0x56:
		goboy.setD(goboy.ReadMemory(goboy.CPU.GetHL()))
		return 8
	case 0x57:
		goboy.setD(goboy.CPU.Registers.A)
		return 4
	// register E LD
	case 0x58:
		goboy.setE(goboy.CPU.Registers.B)
		return 4
	case 0x59:
		goboy.setE(goboy.CPU.Registers.C)
		return 4
	case 0x5A:
		goboy.setE(goboy.CPU.Registers.D)
		return 4
	case 0x5B:
		goboy.setE(goboy.CPU.Registers.E)
		return 4
	case 0x5C:
		goboy.setE(goboy.CPU.Registers.H)
		return 4
	case 0x5D:
		goboy.setE(goboy.CPU.Registers.L)
		return 4
	case 0x5E:
		goboy.setE(goboy.ReadMemory(goboy.CPU.GetHL()))
		return 8
	case 0x5F:
		goboy.setE(goboy.CPU.Registers.A)
		return 4
	// register H LD
	case 0x60:
		goboy.setH(goboy.CPU.Registers.B)
		return 4
	case 0x61:
		goboy.setH(goboy.CPU.Registers.C)
		return 4
	case 0x62:
		goboy.setH(goboy.CPU.Registers.D)
		return 4
	case 0x63:
		goboy.setH(goboy.CPU.Registers.E)
		return 4
	case 0x64:
		goboy.setH(goboy.CPU.Registers.H)
		return 4
	case 0x65:
		goboy.setH(goboy.CPU.Registers.L)
		return 4
	case 0x66:
		goboy.setH(goboy.ReadMemory(goboy.CPU.GetHL()))
		return 8
	case 0x67:
		goboy.setH(goboy.CPU.Registers.A)
		return 4
	// register L LD
	case 0x68:
		goboy.setL(goboy.CPU.Registers.B)
		return 4
	case 0x69:
		goboy.setL(goboy.CPU.Registers.C)
		return 4
	case 0x6A:
		goboy.setL(goboy.CPU.Registers.D)
		return 4
	case 0x6B:
		goboy.setL(goboy.CPU.Registers.E)
		return 4
	case 0x6C:
		goboy.setL(goboy.CPU.Registers.H)
		return 4
	case 0x6D:
		goboy.setL(goboy.CPU.Registers.L)
		return 4
	case 0x6E:
		goboy.setL(goboy.ReadMemory(goboy.CPU.GetHL()))
		return 8
	case 0x6F:
		goboy.setL(goboy.CPU.Registers.A)
		return 4
	// memory value (HL) LD
	case 0x70:
		goboy.WriteMemory(goboy.CPU.GetHL(), goboy.CPU.Registers.B)
		return 8
	case 0x71:
		goboy.WriteMemory(goboy.CPU.GetHL(), goboy.CPU.Registers.C)
		return 8
	case 0x72:
		goboy.WriteMemory(goboy.CPU.GetHL(), goboy.CPU.Registers.D)
		return 8
	case 0x73:
		goboy.WriteMemory(goboy.CPU.GetHL(), goboy.CPU.Registers.E)
		return 8
	case 0x74:
		goboy.WriteMemory(goboy.CPU.GetHL(), goboy.CPU.Registers.H)
		return 8
	case 0x75:
		goboy.WriteMemory(goboy.CPU.GetHL(), goboy.CPU.Registers.L)
		return 8
	case 0x77:
		goboy.WriteMemory(goboy.CPU.GetHL(), goboy.CPU.Registers.A)
		return 8
	case 0x36:
		goboy.WriteMemory(goboy.CPU.GetHL(), goboy.getParameter8Bit())
		return 12

	// END 8bit LD instructions

	// 16bit LD instructions
	case 0x01:
		goboy.CPU.setBC(goboy.getParameter16Bit())
		return 12
	case 0x11:
		goboy.CPU.setDE(goboy.getParameter16Bit())
		return 12
	case 0x21:
		goboy.CPU.setHL(goboy.getParameter16Bit())
		return 12
	case 0x31:
		goboy.CPU.Registers.SP = (goboy.getParameter16Bit())
		return 12
	case 0xF9:
		goboy.CPU.Registers.SP = goboy.CPU.GetHL()
		return 8
	case 0xF8:
		val1 := int32(goboy.CPU.Registers.SP)
		val2 := int32(int8(goboy.getParameter8Bit()))
		result := val1 + val2
		goboy.CPU.setHL(uint16(result))
		tempVal := val1 ^ val2 ^ result
		goboy.CPU.Flags.Zero = false
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = ((tempVal & 0x10) == 0x10)
		goboy.CPU.Flags.Carry = ((tempVal & 0x100) == 0x100)
		goboy.CPU.UpdateFlags()
		return 12
	case 0x08:
		address := goboy.getParameter16Bit()
		goboy.WriteMemory(address, byte(goboy.CPU.Registers.SP&0xff))
		goboy.WriteMemory(address+1, byte((goboy.CPU.Registers.SP&0xff00)>>8))
		return 20
	// PUSH section
	case 0xF5:
		goboy.CPU.UpdateFlags()
		goboy.StackPush(goboy.CPU.GetAF())
		return 16
	case 0xC5:
		goboy.StackPush(goboy.CPU.GetBC())
		return 16
	case 0xD5:
		goboy.StackPush(goboy.CPU.GetDE())
		return 16
	case 0xE5:
		goboy.StackPush(goboy.CPU.GetHL())
		return 16
	// POP section
	case 0xF1:
		goboy.CPU.setAF(goboy.StackPop() & 0xFFF0)
		goboy.CPU.Flags.Zero = CheckBit(goboy.CPU.Registers.F, 7)
		goboy.CPU.Flags.Sub = CheckBit(goboy.CPU.Registers.F, 6)
		goboy.CPU.Flags.HalfCarry = CheckBit(goboy.CPU.Registers.F, 5)
		goboy.CPU.Flags.Carry = CheckBit(goboy.CPU.Registers.F, 4)
		return 12
	case 0xC1:
		goboy.CPU.setBC(goboy.StackPop())
		return 12
	case 0xE1:
		goboy.CPU.setHL(goboy.StackPop())
		return 12
	case 0xD1:
		goboy.CPU.setDE(goboy.StackPop())
		return 12

	// END 16bit LD instructions

	// LDD LDI LDH instructions
	// LDD LDI section
	case 0x2A:
		goboy.setA(goboy.ReadMemory(goboy.CPU.GetHL()))
		goboy.CPU.setHL(goboy.CPU.GetHL() + 1)
		return 8
	case 0x3A:
		goboy.setA(goboy.ReadMemory(goboy.CPU.GetHL()))
		goboy.CPU.setHL(goboy.CPU.GetHL() - 1)
		return 8
	case 0x22:
		goboy.WriteMemory(goboy.CPU.GetHL(), goboy.CPU.Registers.A)
		goboy.CPU.setHL(goboy.CPU.GetHL() + 1)
		return 8
	case 0x32:
		goboy.WriteMemory(goboy.CPU.GetHL(), goboy.CPU.Registers.A)
		goboy.CPU.setHL(goboy.CPU.GetHL() - 1)
		return 8
	// LDH section
	case 0xE0:
		goboy.WriteMemory(0xFF00+uint16(goboy.getParameter8Bit()), goboy.CPU.Registers.A)
		return 12
	case 0xF0:
		goboy.setA(goboy.ReadMemory(0xFF00 + uint16(goboy.getParameter8Bit())))
		return 12
	// END LDD LDI LDH instructions

	// 8bit ALU
	// ADD ADC instructions
	// ADD
	case 0x87:
		origin1 := goboy.CPU.Registers.A
		origin2 := goboy.CPU.Registers.A
		res := int16(goboy.CPU.Registers.A) + int16(origin2)
		goboy.CPU.Registers.A = byte(res)
		goboy.CPU.Flags.Zero = byte(res) == 0
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = (origin1&0xF)+(origin2&0xF) > 0xF
		goboy.CPU.Flags.Carry = res > 0xFF
		goboy.CPU.UpdateFlags()
		return 4
	case 0x80:
		origin1 := goboy.CPU.Registers.A
		origin2 := goboy.CPU.Registers.B
		res := int16(goboy.CPU.Registers.A) + int16(origin2)
		goboy.CPU.Registers.A = byte(res)
		goboy.CPU.Flags.Zero = byte(res) == 0
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = (origin1&0xF)+(origin2&0xF) > 0xF
		goboy.CPU.Flags.Carry = res > 0xFF
		goboy.CPU.UpdateFlags()
		return 4
	case 0x81:
		origin1 := goboy.CPU.Registers.A
		origin2 := goboy.CPU.Registers.C
		res := int16(goboy.CPU.Registers.A) + int16(origin2)
		goboy.CPU.Registers.A = byte(res)
		goboy.CPU.Flags.Zero = byte(res) == 0
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = (origin1&0xF)+(origin2&0xF) > 0xF
		goboy.CPU.Flags.Carry = res > 0xFF
		goboy.CPU.UpdateFlags()
		return 4
	case 0x82:
		origin1 := goboy.CPU.Registers.A
		origin2 := goboy.CPU.Registers.D
		res := int16(goboy.CPU.Registers.A) + int16(origin2)
		goboy.CPU.Registers.A = byte(res)
		goboy.CPU.Flags.Zero = byte(res) == 0
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = (origin1&0xF)+(origin2&0xF) > 0xF
		goboy.CPU.Flags.Carry = res > 0xFF
		goboy.CPU.UpdateFlags()
		return 4
	case 0x83:
		origin1 := goboy.CPU.Registers.A
		origin2 := goboy.CPU.Registers.E
		res := int16(goboy.CPU.Registers.A) + int16(origin2)
		goboy.CPU.Registers.A = byte(res)
		goboy.CPU.Flags.Zero = byte(res) == 0
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = (origin1&0xF)+(origin2&0xF) > 0xF
		goboy.CPU.Flags.Carry = res > 0xFF
		goboy.CPU.UpdateFlags()
		return 4
	case 0x84:
		origin1 := goboy.CPU.Registers.A
		origin2 := goboy.CPU.Registers.H
		res := int16(goboy.CPU.Registers.A) + int16(origin2)
		goboy.CPU.Registers.A = byte(res)
		goboy.CPU.Flags.Zero = byte(res) == 0
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = (origin1&0xF)+(origin2&0xF) > 0xF
		goboy.CPU.Flags.Carry = res > 0xFF
		goboy.CPU.UpdateFlags()
		return 4
	case 0x85:
		origin1 := goboy.CPU.Registers.A
		origin2 := goboy.CPU.Registers.L
		res := int16(goboy.CPU.Registers.A) + int16(origin2)
		goboy.CPU.Registers.A = byte(res)
		goboy.CPU.Flags.Zero = byte(res) == 0
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = (origin1&0xF)+(origin2&0xF) > 0xF
		goboy.CPU.Flags.Carry = res > 0xFF
		goboy.CPU.UpdateFlags()
		return 4
	case 0x86:
		origin1 := goboy.CPU.Registers.A
		origin2 := goboy.ReadMemory(goboy.CPU.GetHL())
		res := int16(goboy.CPU.Registers.A) + int16(origin2)
		goboy.CPU.Registers.A = byte(res)
		goboy.CPU.Flags.Zero = byte(res) == 0
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = (origin1&0xF)+(origin2&0xF) > 0xF
		goboy.CPU.Flags.Carry = res > 0xFF
		goboy.CPU.UpdateFlags()
		return 8
	case 0xC6:
		origin1 := goboy.CPU.Registers.A
		origin2 := goboy.getParameter8Bit()
		res := int16(goboy.CPU.Registers.A) + int16(origin2)
		goboy.CPU.Registers.A = byte(res)
		goboy.CPU.Flags.Zero = byte(res) == 0
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = (origin1&0xF)+(origin2&0xF) > 0xF
		goboy.CPU.Flags.Carry = res > 0xFF
		goboy.CPU.UpdateFlags()
		return 8
	//ADC
	case 0x8F:
		carry := int16(0)
		if goboy.CPU.Flags.Carry {
			carry = 1
		}
		origin1 := goboy.CPU.Registers.A
		origin2 := goboy.CPU.Registers.A
		res := int16(goboy.CPU.Registers.A) + int16(origin2) + carry
		goboy.CPU.Registers.A = byte(res)
		goboy.CPU.Flags.Zero = byte(res) == 0
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = (origin1&0xF)+(origin2&0xF)+byte(carry) > 0xF
		goboy.CPU.Flags.Carry = res > 0xFF
		goboy.CPU.UpdateFlags()
		return 4
	case 0x88:
		carry := int16(0)
		if goboy.CPU.Flags.Carry {
			carry = 1
		}
		origin1 := goboy.CPU.Registers.A
		origin2 := goboy.CPU.Registers.B
		res := int16(goboy.CPU.Registers.A) + int16(origin2) + carry
		goboy.CPU.Registers.A = byte(res)
		goboy.CPU.Flags.Zero = byte(res) == 0
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = (origin1&0xF)+(origin2&0xF)+byte(carry) > 0xF
		goboy.CPU.Flags.Carry = res > 0xFF
		goboy.CPU.UpdateFlags()
		return 4
	case 0x89:
		carry := int16(0)
		if goboy.CPU.Flags.Carry {
			carry = 1
		}
		origin1 := goboy.CPU.Registers.A
		origin2 := goboy.CPU.Registers.C
		res := int16(goboy.CPU.Registers.A) + int16(origin2) + carry
		goboy.CPU.Registers.A = byte(res)
		goboy.CPU.Flags.Zero = byte(res) == 0
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = (origin1&0xF)+(origin2&0xF)+byte(carry) > 0xF
		goboy.CPU.Flags.Carry = res > 0xFF
		goboy.CPU.UpdateFlags()
		return 4
	case 0x8A:
		carry := int16(0)
		if goboy.CPU.Flags.Carry {
			carry = 1
		}
		origin1 := goboy.CPU.Registers.A
		origin2 := goboy.CPU.Registers.D
		res := int16(goboy.CPU.Registers.A) + int16(origin2) + carry
		goboy.CPU.Registers.A = byte(res)
		goboy.CPU.Flags.Zero = byte(res) == 0
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = (origin1&0xF)+(origin2&0xF)+byte(carry) > 0xF
		goboy.CPU.Flags.Carry = res > 0xFF
		goboy.CPU.UpdateFlags()
		return 4
	case 0x8B:
		carry := int16(0)
		if goboy.CPU.Flags.Carry {
			carry = 1
		}
		origin1 := goboy.CPU.Registers.A
		origin2 := goboy.CPU.Registers.E
		res := int16(goboy.CPU.Registers.A) + int16(origin2) + carry
		goboy.CPU.Registers.A = byte(res)
		goboy.CPU.Flags.Zero = byte(res) == 0
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = (origin1&0xF)+(origin2&0xF)+byte(carry) > 0xF
		goboy.CPU.Flags.Carry = res > 0xFF
		goboy.CPU.UpdateFlags()
		return 4
	case 0x8C:
		carry := int16(0)
		if goboy.CPU.Flags.Carry {
			carry = 1
		}
		origin1 := goboy.CPU.Registers.A
		origin2 := goboy.CPU.Registers.H
		res := int16(goboy.CPU.Registers.A) + int16(origin2) + carry
		goboy.CPU.Registers.A = byte(res)
		goboy.CPU.Flags.Zero = byte(res) == 0
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = (origin1&0xF)+(origin2&0xF)+byte(carry) > 0xF
		goboy.CPU.Flags.Carry = res > 0xFF
		goboy.CPU.UpdateFlags()
		return 4
	case 0x8D:
		carry := int16(0)
		if goboy.CPU.Flags.Carry {
			carry = 1
		}
		origin1 := goboy.CPU.Registers.A
		origin2 := goboy.CPU.Registers.L
		res := int16(goboy.CPU.Registers.A) + int16(origin2) + carry
		goboy.CPU.Registers.A = byte(res)
		goboy.CPU.Flags.Zero = byte(res) == 0
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = (origin1&0xF)+(origin2&0xF)+byte(carry) > 0xF
		goboy.CPU.Flags.Carry = res > 0xFF
		goboy.CPU.UpdateFlags()
		return 4
	case 0x8E:
		carry := int16(0)
		if goboy.CPU.Flags.Carry {
			carry = 1
		}
		origin1 := goboy.CPU.Registers.A
		origin2 := goboy.ReadMemory(goboy.CPU.GetHL())
		res := int16(goboy.CPU.Registers.A) + int16(origin2) + carry
		goboy.CPU.Registers.A = byte(res)
		goboy.CPU.Flags.Zero = byte(res) == 0
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = (origin1&0xF)+(origin2&0xF)+byte(carry) > 0xF
		goboy.CPU.Flags.Carry = res > 0xFF
		goboy.CPU.UpdateFlags()
		return 8
	case 0xCE:
		carry := int16(0)
		if goboy.CPU.Flags.Carry {
			carry = 1
		}
		origin1 := goboy.CPU.Registers.A
		origin2 := goboy.getParameter8Bit()
		res := int16(goboy.CPU.Registers.A) + int16(origin2) + carry
		goboy.CPU.Registers.A = byte(res)
		goboy.CPU.Flags.Zero = byte(res) == 0
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = (origin1&0xF)+(origin2&0xF)+byte(carry) > 0xF
		goboy.CPU.Flags.Carry = res > 0xFF
		goboy.CPU.UpdateFlags()
		return 8
	// SUB SBC instructions
	// SUB
	case 0x97:
		val := goboy.CPU.Registers.A
		origin := goboy.CPU.Registers.A
		res := int16(goboy.CPU.Registers.A) - int16(val)
		final := byte(res)
		goboy.CPU.Registers.A = final
		goboy.CPU.Flags.Zero = final == 0
		goboy.CPU.Flags.Sub = true
		goboy.CPU.Flags.HalfCarry = int16(origin&0x0F)-int16(val&0xF) < 0
		goboy.CPU.Flags.Carry = res < 0
		goboy.CPU.UpdateFlags()
		return 4
	case 0x90:
		val := goboy.CPU.Registers.B
		origin := goboy.CPU.Registers.A
		res := int16(goboy.CPU.Registers.A) - int16(val)
		final := byte(res)
		goboy.CPU.Registers.A = final
		goboy.CPU.Flags.Zero = final == 0
		goboy.CPU.Flags.Sub = true
		goboy.CPU.Flags.HalfCarry = int16(origin&0x0F)-int16(val&0xF) < 0
		goboy.CPU.Flags.Carry = res < 0
		goboy.CPU.UpdateFlags()
		return 4
	case 0x91:
		val := goboy.CPU.Registers.C
		origin := goboy.CPU.Registers.A
		res := int16(goboy.CPU.Registers.A) - int16(val)
		final := byte(res)
		goboy.CPU.Registers.A = final
		goboy.CPU.Flags.Zero = final == 0
		goboy.CPU.Flags.Sub = true
		goboy.CPU.Flags.HalfCarry = int16(origin&0x0F)-int16(val&0xF) < 0
		goboy.CPU.Flags.Carry = res < 0
		goboy.CPU.UpdateFlags()
		return 4
	case 0x92:
		val := goboy.CPU.Registers.D
		origin := goboy.CPU.Registers.A
		res := int16(goboy.CPU.Registers.A) - int16(val)
		final := byte(res)
		goboy.CPU.Registers.A = final
		goboy.CPU.Flags.Zero = final == 0
		goboy.CPU.Flags.Sub = true
		goboy.CPU.Flags.HalfCarry = int16(origin&0x0F)-int16(val&0xF) < 0
		goboy.CPU.Flags.Carry = res < 0
		goboy.CPU.UpdateFlags()
		return 4
	case 0x93:
		val := goboy.CPU.Registers.E
		origin := goboy.CPU.Registers.A
		res := int16(goboy.CPU.Registers.A) - int16(val)
		final := byte(res)
		goboy.CPU.Registers.A = final
		goboy.CPU.Flags.Zero = final == 0
		goboy.CPU.Flags.Sub = true
		goboy.CPU.Flags.HalfCarry = int16(origin&0x0F)-int16(val&0xF) < 0
		goboy.CPU.Flags.Carry = res < 0
		goboy.CPU.UpdateFlags()
		return 4
	case 0x94:
		val := goboy.CPU.Registers.H
		origin := goboy.CPU.Registers.A
		res := int16(goboy.CPU.Registers.A) - int16(val)
		final := byte(res)
		goboy.CPU.Registers.A = final
		goboy.CPU.Flags.Zero = final == 0
		goboy.CPU.Flags.Sub = true
		goboy.CPU.Flags.HalfCarry = int16(origin&0x0F)-int16(val&0xF) < 0
		goboy.CPU.Flags.Carry = res < 0
		goboy.CPU.UpdateFlags()
		return 4
	case 0x95:
		val := goboy.CPU.Registers.L
		origin := goboy.CPU.Registers.A
		res := int16(goboy.CPU.Registers.A) - int16(val)
		final := byte(res)
		goboy.CPU.Registers.A = final
		goboy.CPU.Flags.Zero = final == 0
		goboy.CPU.Flags.Sub = true
		goboy.CPU.Flags.HalfCarry = int16(origin&0x0F)-int16(val&0xF) < 0
		goboy.CPU.Flags.Carry = res < 0
		goboy.CPU.UpdateFlags()
		return 4
	case 0x96:
		val := goboy.ReadMemory(goboy.CPU.GetHL())
		origin := goboy.CPU.Registers.A
		res := int16(goboy.CPU.Registers.A) - int16(val)
		final := byte(res)
		goboy.CPU.Registers.A = final
		goboy.CPU.Flags.Zero = final == 0
		goboy.CPU.Flags.Sub = true
		goboy.CPU.Flags.HalfCarry = int16(origin&0x0F)-int16(val&0xF) < 0
		goboy.CPU.Flags.Carry = res < 0
		goboy.CPU.UpdateFlags()
		return 8
	case 0xD6:
		val := goboy.getParameter8Bit()
		origin := goboy.CPU.Registers.A
		res := int16(goboy.CPU.Registers.A) - int16(val)
		final := byte(res)
		goboy.CPU.Registers.A = final
		goboy.CPU.Flags.Zero = final == 0
		goboy.CPU.Flags.Sub = true
		goboy.CPU.Flags.HalfCarry = int16(origin&0x0F)-int16(val&0xF) < 0
		goboy.CPU.Flags.Carry = res < 0
		goboy.CPU.UpdateFlags()
		return 8
	// SBC
	case 0x9F:
		carry := int16(0)
		if goboy.CPU.Flags.Carry {
			carry = 1
		}
		val := goboy.CPU.Registers.A
		origin := goboy.CPU.Registers.A
		res := int16(goboy.CPU.Registers.A) - int16(val) - carry
		final := byte(res)
		goboy.CPU.Registers.A = final
		goboy.CPU.Flags.Zero = final == 0
		goboy.CPU.Flags.Sub = true
		goboy.CPU.Flags.HalfCarry = int16(origin&0x0F)-int16(val&0xF)-carry < 0
		goboy.CPU.Flags.Carry = res < 0
		goboy.CPU.UpdateFlags()
		return 4
	case 0x98:
		carry := int16(0)
		if goboy.CPU.Flags.Carry {
			carry = 1
		}
		val := goboy.CPU.Registers.B
		origin := goboy.CPU.Registers.A
		res := int16(goboy.CPU.Registers.A) - int16(val) - carry
		final := byte(res)
		goboy.CPU.Registers.A = final
		goboy.CPU.Flags.Zero = final == 0
		goboy.CPU.Flags.Sub = true
		goboy.CPU.Flags.HalfCarry = int16(origin&0x0F)-int16(val&0xF)-carry < 0
		goboy.CPU.Flags.Carry = res < 0
		goboy.CPU.UpdateFlags()
		return 4
	case 0x99:
		carry := int16(0)
		if goboy.CPU.Flags.Carry {
			carry = 1
		}
		val := goboy.CPU.Registers.C
		origin := goboy.CPU.Registers.A
		res := int16(goboy.CPU.Registers.A) - int16(val) - carry
		final := byte(res)
		goboy.CPU.Registers.A = final
		goboy.CPU.Flags.Zero = final == 0
		goboy.CPU.Flags.Sub = true
		goboy.CPU.Flags.HalfCarry = int16(origin&0x0F)-int16(val&0xF)-carry < 0
		goboy.CPU.Flags.Carry = res < 0
		goboy.CPU.UpdateFlags()
		return 4
	case 0x9A:
		carry := int16(0)
		if goboy.CPU.Flags.Carry {
			carry = 1
		}
		val := goboy.CPU.Registers.D
		origin := goboy.CPU.Registers.A
		res := int16(goboy.CPU.Registers.A) - int16(val) - carry
		final := byte(res)
		goboy.CPU.Registers.A = final
		goboy.CPU.Flags.Zero = final == 0
		goboy.CPU.Flags.Sub = true
		goboy.CPU.Flags.HalfCarry = int16(origin&0x0F)-int16(val&0xF)-carry < 0
		goboy.CPU.Flags.Carry = res < 0
		goboy.CPU.UpdateFlags()
		return 4
	case 0x9B:
		carry := int16(0)
		if goboy.CPU.Flags.Carry {
			carry = 1
		}
		val := goboy.CPU.Registers.E
		origin := goboy.CPU.Registers.A
		res := int16(goboy.CPU.Registers.A) - int16(val) - carry
		final := byte(res)
		goboy.CPU.Registers.A = final
		goboy.CPU.Flags.Zero = final == 0
		goboy.CPU.Flags.Sub = true
		goboy.CPU.Flags.HalfCarry = int16(origin&0x0F)-int16(val&0xF)-carry < 0
		goboy.CPU.Flags.Carry = res < 0
		goboy.CPU.UpdateFlags()
		return 4
	case 0x9C:
		carry := int16(0)
		if goboy.CPU.Flags.Carry {
			carry = 1
		}
		val := goboy.CPU.Registers.H
		origin := goboy.CPU.Registers.A
		res := int16(goboy.CPU.Registers.A) - int16(val) - carry
		final := byte(res)
		goboy.CPU.Registers.A = final
		goboy.CPU.Flags.Zero = final == 0
		goboy.CPU.Flags.Sub = true
		goboy.CPU.Flags.HalfCarry = int16(origin&0x0F)-int16(val&0xF)-carry < 0
		goboy.CPU.Flags.Carry = res < 0
		goboy.CPU.UpdateFlags()
		return 4
	case 0x9D:
		carry := int16(0)
		if goboy.CPU.Flags.Carry {
			carry = 1
		}
		val := goboy.CPU.Registers.L
		origin := goboy.CPU.Registers.A
		res := int16(goboy.CPU.Registers.A) - int16(val) - carry
		final := byte(res)
		goboy.CPU.Registers.A = final
		goboy.CPU.Flags.Zero = final == 0
		goboy.CPU.Flags.Sub = true
		goboy.CPU.Flags.HalfCarry = int16(origin&0x0F)-int16(val&0xF)-carry < 0
		goboy.CPU.Flags.Carry = res < 0
		goboy.CPU.UpdateFlags()
		return 4
	case 0x9E:
		carry := int16(0)
		if goboy.CPU.Flags.Carry {
			carry = 1
		}
		val := goboy.ReadMemory(goboy.CPU.GetHL())
		origin := goboy.CPU.Registers.A
		res := int16(goboy.CPU.Registers.A) - int16(val) - carry
		final := byte(res)
		goboy.CPU.Registers.A = final
		goboy.CPU.Flags.Zero = final == 0
		goboy.CPU.Flags.Sub = true
		goboy.CPU.Flags.HalfCarry = int16(origin&0x0F)-int16(val&0xF)-carry < 0
		goboy.CPU.Flags.Carry = res < 0
		goboy.CPU.UpdateFlags()
		return 8
	case 0xDE:
		carry := int16(0)
		if goboy.CPU.Flags.Carry {
			carry = 1
		}
		val := goboy.getParameter8Bit()
		origin := goboy.CPU.Registers.A
		res := int16(goboy.CPU.Registers.A) - int16(val) - carry
		final := byte(res)
		goboy.CPU.Registers.A = final
		goboy.CPU.Flags.Zero = final == 0
		goboy.CPU.Flags.Sub = true
		goboy.CPU.Flags.HalfCarry = int16(origin&0x0F)-int16(val&0xF)-carry < 0
		goboy.CPU.Flags.Carry = res < 0
		goboy.CPU.UpdateFlags()
		return 8
	// AND
	case 0xA7:
		goboy.CPU.Registers.A = goboy.CPU.Registers.A & goboy.CPU.Registers.A
		goboy.CPU.Flags.Zero = (goboy.CPU.Registers.A == 0)
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = true
		goboy.CPU.Flags.Carry = false
		goboy.CPU.UpdateFlags()
		return 4
	case 0xA0:
		goboy.CPU.Registers.A = goboy.CPU.Registers.B & goboy.CPU.Registers.A
		goboy.CPU.Flags.Zero = (goboy.CPU.Registers.A == 0)
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = true
		goboy.CPU.Flags.Carry = false
		goboy.CPU.UpdateFlags()
		return 4
	case 0xA1:
		goboy.CPU.Registers.A = goboy.CPU.Registers.C & goboy.CPU.Registers.A
		goboy.CPU.Flags.Zero = (goboy.CPU.Registers.A == 0)
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = true
		goboy.CPU.Flags.Carry = false
		goboy.CPU.UpdateFlags()
		return 4
	case 0xA2:
		goboy.CPU.Registers.A = goboy.CPU.Registers.D & goboy.CPU.Registers.A
		goboy.CPU.Flags.Zero = (goboy.CPU.Registers.A == 0)
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = true
		goboy.CPU.Flags.Carry = false
		goboy.CPU.UpdateFlags()
		return 4
	case 0xA3:
		goboy.CPU.Registers.A = goboy.CPU.Registers.E & goboy.CPU.Registers.A
		goboy.CPU.Flags.Zero = (goboy.CPU.Registers.A == 0)
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = true
		goboy.CPU.Flags.Carry = false
		goboy.CPU.UpdateFlags()
		return 4
	case 0xA4:
		goboy.CPU.Registers.A = goboy.CPU.Registers.H & goboy.CPU.Registers.A
		goboy.CPU.Flags.Zero = (goboy.CPU.Registers.A == 0)
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = true
		goboy.CPU.Flags.Carry = false
		goboy.CPU.UpdateFlags()
		return 4
	case 0xA5:
		goboy.CPU.Registers.A = goboy.CPU.Registers.L & goboy.CPU.Registers.A
		goboy.CPU.Flags.Zero = (goboy.CPU.Registers.A == 0)
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = true
		goboy.CPU.Flags.Carry = false
		goboy.CPU.UpdateFlags()
		return 4
	case 0xA6:
		goboy.CPU.Registers.A = goboy.ReadMemory(goboy.CPU.GetHL()) & goboy.CPU.Registers.A
		goboy.CPU.Flags.Zero = (goboy.CPU.Registers.A == 0)
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = true
		goboy.CPU.Flags.Carry = false
		goboy.CPU.UpdateFlags()
		return 8
	case 0xE6:
		goboy.CPU.Registers.A = goboy.getParameter8Bit() & goboy.CPU.Registers.A
		goboy.CPU.Flags.Zero = (goboy.CPU.Registers.A == 0)
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = true
		goboy.CPU.Flags.Carry = false
		goboy.CPU.UpdateFlags()
		return 8
	// OR
	case 0xB7:
		goboy.CPU.Registers.A = goboy.CPU.Registers.A | goboy.CPU.Registers.A
		goboy.CPU.Flags.Zero = (goboy.CPU.Registers.A == 0)
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = false
		goboy.CPU.Flags.Carry = false
		goboy.CPU.UpdateFlags()
		return 4
	case 0xB0:
		goboy.CPU.Registers.A = goboy.CPU.Registers.B | goboy.CPU.Registers.A
		goboy.CPU.Flags.Zero = (goboy.CPU.Registers.A == 0)
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = false
		goboy.CPU.Flags.Carry = false
		goboy.CPU.UpdateFlags()
		return 4
	case 0xB1:
		goboy.CPU.Registers.A = goboy.CPU.Registers.C | goboy.CPU.Registers.A
		goboy.CPU.Flags.Zero = (goboy.CPU.Registers.A == 0)
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = false
		goboy.CPU.Flags.Carry = false
		goboy.CPU.UpdateFlags()
		return 4
	case 0xB2:
		goboy.CPU.Registers.A = goboy.CPU.Registers.D | goboy.CPU.Registers.A
		goboy.CPU.Flags.Zero = (goboy.CPU.Registers.A == 0)
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = false
		goboy.CPU.Flags.Carry = false
		goboy.CPU.UpdateFlags()
		return 4
	case 0xB3:
		goboy.CPU.Registers.A = goboy.CPU.Registers.E | goboy.CPU.Registers.A
		goboy.CPU.Flags.Zero = (goboy.CPU.Registers.A == 0)
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = false
		goboy.CPU.Flags.Carry = false
		goboy.CPU.UpdateFlags()
		return 4
	case 0xB4:
		goboy.CPU.Registers.A = goboy.CPU.Registers.H | goboy.CPU.Registers.A
		goboy.CPU.Flags.Zero = (goboy.CPU.Registers.A == 0)
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = false
		goboy.CPU.Flags.Carry = false
		goboy.CPU.UpdateFlags()
		return 4
	case 0xB5:
		goboy.CPU.Registers.A = goboy.CPU.Registers.L | goboy.CPU.Registers.A
		goboy.CPU.Flags.Zero = (goboy.CPU.Registers.A == 0)
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = false
		goboy.CPU.Flags.Carry = false
		goboy.CPU.UpdateFlags()
		return 4
	case 0xB6:
		goboy.CPU.Registers.A = goboy.ReadMemory(goboy.CPU.GetHL()) | goboy.CPU.Registers.A
		goboy.CPU.Flags.Zero = (goboy.CPU.Registers.A == 0)
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = false
		goboy.CPU.Flags.Carry = false
		goboy.CPU.UpdateFlags()
		return 8
	case 0xF6:
		goboy.CPU.Registers.A = goboy.getParameter8Bit() | goboy.CPU.Registers.A
		goboy.CPU.Flags.Zero = (goboy.CPU.Registers.A == 0)
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = false
		goboy.CPU.Flags.Carry = false
		goboy.CPU.UpdateFlags()
		return 8
	// XOR
	case 0xAF:
		goboy.CPU.Registers.A = goboy.CPU.Registers.A ^ goboy.CPU.Registers.A
		goboy.CPU.Flags.Zero = goboy.CPU.Registers.A == 0
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = false
		goboy.CPU.Flags.Carry = false
		goboy.CPU.UpdateFlags()
		return 4
	case 0xA8:
		goboy.CPU.Registers.A = goboy.CPU.Registers.B ^ goboy.CPU.Registers.A
		goboy.CPU.Flags.Zero = goboy.CPU.Registers.A == 0
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = false
		goboy.CPU.Flags.Carry = false
		goboy.CPU.UpdateFlags()
		return 4
	case 0xA9:
		goboy.CPU.Registers.A = goboy.CPU.Registers.C ^ goboy.CPU.Registers.A
		goboy.CPU.Flags.Zero = goboy.CPU.Registers.A == 0
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = false
		goboy.CPU.Flags.Carry = false
		goboy.CPU.UpdateFlags()
		return 4
	case 0xAA:
		goboy.CPU.Registers.A = goboy.CPU.Registers.D ^ goboy.CPU.Registers.A
		goboy.CPU.Flags.Zero = goboy.CPU.Registers.A == 0
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = false
		goboy.CPU.Flags.Carry = false
		goboy.CPU.UpdateFlags()
		return 4
	case 0xAB:
		goboy.CPU.Registers.A = goboy.CPU.Registers.E ^ goboy.CPU.Registers.A
		goboy.CPU.Flags.Zero = goboy.CPU.Registers.A == 0
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = false
		goboy.CPU.Flags.Carry = false
		goboy.CPU.UpdateFlags()
		return 4
	case 0xAC:
		goboy.CPU.Registers.A = goboy.CPU.Registers.H ^ goboy.CPU.Registers.A
		goboy.CPU.Flags.Zero = goboy.CPU.Registers.A == 0
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = false
		goboy.CPU.Flags.Carry = false
		goboy.CPU.UpdateFlags()
		return 4
	case 0xAD:
		goboy.CPU.Registers.A = goboy.CPU.Registers.L ^ goboy.CPU.Registers.A
		goboy.CPU.Flags.Zero = goboy.CPU.Registers.A == 0
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = false
		goboy.CPU.Flags.Carry = false
		goboy.CPU.UpdateFlags()
		return 4
	case 0xAE:
		goboy.CPU.Registers.A = goboy.ReadMemory(goboy.CPU.GetHL()) ^ goboy.CPU.Registers.A
		goboy.CPU.Flags.Zero = goboy.CPU.Registers.A == 0
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = false
		goboy.CPU.Flags.Carry = false
		goboy.CPU.UpdateFlags()
		return 8
	case 0xEE:
		goboy.CPU.Registers.A = goboy.getParameter8Bit() ^ goboy.CPU.Registers.A
		goboy.CPU.Flags.Zero = goboy.CPU.Registers.A == 0
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = false
		goboy.CPU.Flags.Carry = false
		goboy.CPU.UpdateFlags()
		return 8
	// CP
	case 0xBF:
		goboy.CPU.Compare(goboy.CPU.Registers.A, goboy.CPU.Registers.A)
		return 4
	case 0xB8:
		goboy.CPU.Compare(goboy.CPU.Registers.B, goboy.CPU.Registers.A)
		return 4
	case 0xB9:
		goboy.CPU.Compare(goboy.CPU.Registers.C, goboy.CPU.Registers.A)
		return 4
	case 0xBA:
		goboy.CPU.Compare(goboy.CPU.Registers.D, goboy.CPU.Registers.A)
		return 4
	case 0xBB:
		goboy.CPU.Compare(goboy.CPU.Registers.E, goboy.CPU.Registers.A)
		return 4
	case 0xBC:
		goboy.CPU.Compare(goboy.CPU.Registers.H, goboy.CPU.Registers.A)
		return 4
	case 0xBD:
		goboy.CPU.Compare(goboy.CPU.Registers.L, goboy.CPU.Registers.A)
		return 4
	case 0xBE:
		goboy.CPU.Compare(goboy.ReadMemory(goboy.CPU.GetHL()), goboy.CPU.Registers.A)
		return 8
	case 0xFE:
		goboy.CPU.Compare(goboy.getParameter8Bit(), goboy.CPU.Registers.A)
		return 8
	// INC
	case 0x3C:
		origin := goboy.CPU.Registers.A
		newVal := origin + 1
		goboy.CPU.Registers.A = newVal
		goboy.CPU.Flags.Zero = (newVal == 0)
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = ((origin&0xF)+(1&0xF) > 0xF)
		goboy.CPU.UpdateFlags()
		return 4
	case 0x04:
		origin := goboy.CPU.Registers.B
		newVal := origin + 1
		goboy.CPU.Registers.B = newVal
		goboy.CPU.Flags.Zero = (newVal == 0)
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = ((origin&0xF)+(1&0xF) > 0xF)
		goboy.CPU.UpdateFlags()
		return 4
	case 0x0C:
		origin := goboy.CPU.Registers.C
		newVal := origin + 1
		goboy.CPU.Registers.C = newVal
		goboy.CPU.Flags.Zero = (newVal == 0)
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = ((origin&0xF)+(1&0xF) > 0xF)
		goboy.CPU.UpdateFlags()
		return 4
	case 0x14:
		origin := goboy.CPU.Registers.D
		newVal := origin + 1
		goboy.CPU.Registers.D = newVal
		goboy.CPU.Flags.Zero = (newVal == 0)
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = ((origin&0xF)+(1&0xF) > 0xF)
		goboy.CPU.UpdateFlags()
		return 4
	case 0x1C:
		origin := goboy.CPU.Registers.E
		newVal := origin + 1
		goboy.CPU.Registers.E = newVal
		goboy.CPU.Flags.Zero = (newVal == 0)
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = ((origin&0xF)+(1&0xF) > 0xF)
		goboy.CPU.UpdateFlags()
		return 4
	case 0x24:
		origin := goboy.CPU.Registers.H
		newVal := origin + 1
		goboy.CPU.Registers.H = newVal
		goboy.CPU.Flags.Zero = (newVal == 0)
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = ((origin&0xF)+(1&0xF) > 0xF)
		goboy.CPU.UpdateFlags()
		return 4
	case 0x2C:
		origin := goboy.CPU.Registers.L
		newVal := origin + 1
		goboy.CPU.Registers.L = newVal
		goboy.CPU.Flags.Zero = (newVal == 0)
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = ((origin&0xF)+(1&0xF) > 0xF)
		goboy.CPU.UpdateFlags()
		return 4
	case 0x34:
		origin := goboy.ReadMemory(goboy.CPU.GetHL())
		newVal := origin + 1
		goboy.WriteMemory(goboy.CPU.GetHL(), newVal)
		goboy.CPU.Flags.Zero = (newVal == 0)
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = ((origin&0xF)+(1&0xF) > 0xF)
		goboy.CPU.UpdateFlags()
		return 12
	// DEC
	case 0x3D:
		origin := goboy.CPU.Registers.A
		goboy.CPU.Registers.A--
		goboy.CPU.Flags.Zero = (goboy.CPU.Registers.A == 0)
		goboy.CPU.Flags.Sub = true
		goboy.CPU.Flags.HalfCarry = (origin&0x0F == 0)
		goboy.CPU.UpdateFlags()
		return 4
	case 0x05:
		origin := goboy.CPU.Registers.B
		goboy.CPU.Registers.B--
		goboy.CPU.Flags.Zero = (goboy.CPU.Registers.B == 0)
		goboy.CPU.Flags.Sub = true
		goboy.CPU.Flags.HalfCarry = (origin&0x0F == 0)
		goboy.CPU.UpdateFlags()
		return 4
	case 0x0D:
		origin := goboy.CPU.Registers.C
		goboy.CPU.Registers.C--
		goboy.CPU.Flags.Zero = (goboy.CPU.Registers.C == 0)
		goboy.CPU.Flags.Sub = true
		goboy.CPU.Flags.HalfCarry = (origin&0x0F == 0)
		goboy.CPU.UpdateFlags()
		return 4
	case 0x15:
		origin := goboy.CPU.Registers.D
		goboy.CPU.Registers.D--
		goboy.CPU.Flags.Zero = (goboy.CPU.Registers.D == 0)
		goboy.CPU.Flags.Sub = true
		goboy.CPU.Flags.HalfCarry = (origin&0x0F == 0)
		goboy.CPU.UpdateFlags()
		return 4
	case 0x1D:
		origin := goboy.CPU.Registers.E
		goboy.CPU.Registers.E--
		goboy.CPU.Flags.Zero = (goboy.CPU.Registers.E == 0)
		goboy.CPU.Flags.Sub = true
		goboy.CPU.Flags.HalfCarry = (origin&0x0F == 0)
		goboy.CPU.UpdateFlags()
		return 4
	case 0x25:
		origin := goboy.CPU.Registers.H
		goboy.CPU.Registers.H--
		goboy.CPU.Flags.Zero = (goboy.CPU.Registers.H == 0)
		goboy.CPU.Flags.Sub = true
		goboy.CPU.Flags.HalfCarry = (origin&0x0F == 0)
		goboy.CPU.UpdateFlags()
		return 4
	case 0x2D:
		origin := goboy.CPU.Registers.L
		goboy.CPU.Registers.L--
		goboy.CPU.Flags.Zero = (goboy.CPU.Registers.L == 0)
		goboy.CPU.Flags.Sub = true
		goboy.CPU.Flags.HalfCarry = (origin&0x0F == 0)
		goboy.CPU.UpdateFlags()
		return 4
	case 0x35:
		origin := goboy.ReadMemory(goboy.CPU.GetHL())
		newVal := origin - 1
		goboy.WriteMemory(goboy.CPU.GetHL(), newVal)
		goboy.CPU.Flags.Zero = (newVal == 0)
		goboy.CPU.Flags.Sub = true
		goboy.CPU.Flags.HalfCarry = (origin&0x0F == 0)
		goboy.CPU.UpdateFlags()
		return 12
	// END 8bit ALU

	// 16bit ALU
	// ADD
	case 0x09:
		originHL := goboy.CPU.GetHL()
		originBC := goboy.CPU.GetBC()
		res := int32(originBC) + int32(originHL)
		goboy.CPU.setHL(uint16(res))
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = int32(originHL&0xFFF) > (res & 0xFFF)
		goboy.CPU.Flags.Carry = res > 0xFFFF
		goboy.CPU.UpdateFlags()
		return 8
	case 0x19:
		originHL := goboy.CPU.GetHL()
		originDE := goboy.CPU.GetDE()
		res := int32(originDE) + int32(originHL)
		goboy.CPU.setHL(uint16(res))
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = int32(originHL&0xFFF) > (res & 0xFFF)
		goboy.CPU.Flags.Carry = res > 0xFFFF
		goboy.CPU.UpdateFlags()
		return 8
	case 0x29:
		originHL := goboy.CPU.GetHL()
		originHL2 := goboy.CPU.GetHL()
		res := int32(originHL2) + int32(originHL)
		goboy.CPU.setHL(uint16(res))
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = int32(originHL&0xFFF) > (res & 0xFFF)
		goboy.CPU.Flags.Carry = res > 0xFFFF
		goboy.CPU.UpdateFlags()
		return 8
	case 0x39:
		originHL := goboy.CPU.GetHL()
		originSP := goboy.CPU.Registers.SP
		res := int32(originSP) + int32(originHL)
		goboy.CPU.setHL(uint16(res))
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = int32(originHL&0xFFF) > (res & 0xFFF)
		goboy.CPU.Flags.Carry = res > 0xFFFF
		goboy.CPU.UpdateFlags()
		return 8
	case 0xE8:
		origin1 := goboy.CPU.Registers.SP
		origin2 := int8(goboy.getParameter8Bit())
		res := uint16(int32(goboy.CPU.Registers.SP) + int32(origin2))
		tmpVal := origin1 ^ uint16(origin2) ^ res
		goboy.CPU.Registers.SP = res
		goboy.CPU.Flags.Zero = false
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = (tmpVal & 0x10) == 0x10
		goboy.CPU.Flags.Carry = ((tmpVal & 0x100) == 0x100)
		goboy.CPU.UpdateFlags()
		return 16
	// INC
	case 0x03:
		goboy.CPU.setBC(goboy.CPU.GetBC() + 1)
		return 8
	case 0x13:
		goboy.CPU.setDE(goboy.CPU.GetDE() + 1)
		return 8
	case 0x23:
		goboy.CPU.setHL(goboy.CPU.GetHL() + 1)
		return 8
	case 0x33:
		goboy.CPU.Registers.SP = goboy.CPU.Registers.SP + 1
		return 8

	// DEC
	case 0x0B:
		goboy.CPU.setBC(goboy.CPU.GetBC() - 1)
		return 8
	case 0x1B:
		goboy.CPU.setDE(goboy.CPU.GetDE() - 1)
		return 8
	case 0x2B:
		goboy.CPU.setHL(goboy.CPU.GetHL() - 1)
		return 8
	case 0x3B:
		goboy.CPU.Registers.SP = goboy.CPU.Registers.SP - 1
		return 8

	// END 16bit ALU

	// Jumps
	case 0xC3:
		addr := goboy.getParameter16Bit()
		goboy.CPU.Registers.PC = addr
		return 12
	case 0xC2:
		addr := goboy.getParameter16Bit()
		if !goboy.CPU.Flags.Zero {
			goboy.CPU.Registers.PC = addr
			ticks += 4
		}
		return 12
	case 0xCA:
		addr := goboy.getParameter16Bit()
		if goboy.CPU.Flags.Zero {
			goboy.CPU.Registers.PC = addr
			ticks += 4
		}
		return 12
	case 0xD2:
		addr := goboy.getParameter16Bit()
		if !goboy.CPU.Flags.Carry {
			goboy.CPU.Registers.PC = addr
			ticks += 4
		}
		return 12
	case 0xDA:
		addr := goboy.getParameter16Bit()
		if goboy.CPU.Flags.Carry {
			goboy.CPU.Registers.PC = addr
			ticks += 4
		}
		return 12
	case 0xE9:
		goboy.CPU.Registers.PC = goboy.CPU.GetHL()
		return 4
	case 0x18:
		addr := int8(goboy.getParameter8Bit())
		goboy.CPU.Registers.PC = uint16(int32(goboy.CPU.Registers.PC) + int32(addr))
		return 8
	case 0x20:
		addr := int8(goboy.getParameter8Bit())
		if !goboy.CPU.Flags.Zero {
			goboy.CPU.Registers.PC = uint16(int32(goboy.CPU.Registers.PC) + int32(addr))
			ticks += 4
		}
		return 8
	case 0x28:
		addr := int8(goboy.getParameter8Bit())
		if goboy.CPU.Flags.Zero {
			goboy.CPU.Registers.PC = uint16(int32(goboy.CPU.Registers.PC) + int32(addr))
			ticks += 4
		}
		return 8
	case 0x30:
		addr := int8(goboy.getParameter8Bit())
		if !goboy.CPU.Flags.Carry {
			goboy.CPU.Registers.PC = uint16(int32(goboy.CPU.Registers.PC) + int32(addr))
			ticks += 4
		}
		return 8
	case 0x38:
		addr := int8(goboy.getParameter8Bit())
		if goboy.CPU.Flags.Carry {
			goboy.CPU.Registers.PC = uint16(int32(goboy.CPU.Registers.PC) + int32(addr))
			ticks += 4
		}
		return 8
	// END Jumps

	case 0x27:
		if !goboy.CPU.Flags.Sub {

			if goboy.CPU.Flags.Carry || goboy.CPU.Registers.A > 0x99 {
				goboy.CPU.Registers.A = goboy.CPU.Registers.A + 0x60
				goboy.CPU.Flags.Carry = true
			}
			if goboy.CPU.Flags.HalfCarry || goboy.CPU.Registers.A&0xF > 0x9 {
				goboy.CPU.Registers.A = goboy.CPU.Registers.A + 0x06
				goboy.CPU.Flags.HalfCarry = false
			}
		} else if goboy.CPU.Flags.Carry && goboy.CPU.Flags.HalfCarry {
			goboy.CPU.Registers.A = goboy.CPU.Registers.A + 0x9A
			goboy.CPU.Flags.HalfCarry = false
		} else if goboy.CPU.Flags.Carry {
			goboy.CPU.Registers.A = goboy.CPU.Registers.A + 0xA0
		} else if goboy.CPU.Flags.HalfCarry {
			goboy.CPU.Registers.A += 0xFA
			goboy.CPU.Flags.HalfCarry = false
		}
		goboy.CPU.Flags.Zero = goboy.CPU.Registers.A == 0
		goboy.CPU.UpdateFlags()
		return 4
	case 0x2F:
		goboy.CPU.Registers.A = 0xFF ^ goboy.CPU.Registers.A
		goboy.CPU.Flags.Sub = true
		goboy.CPU.Flags.HalfCarry = true
		goboy.CPU.UpdateFlags()
		return 4
	case 0x3F:
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = false
		goboy.CPU.Flags.Carry = !goboy.CPU.Flags.Carry
		goboy.CPU.UpdateFlags()
		return 4
	case 0x37:
		goboy.CPU.Flags.Carry = true
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = false
		goboy.CPU.UpdateFlags()
		return 4
	case 0x00:
		return 4
	case 0x76:
		goboy.CPU.Halted = true
		return 4
	case 0x10: //TODO improve this opcode
		goboy.CPU.Halted = true
		goboy.CPU.Registers.PC += 1
		//display black line to the screen
		return 4
	case 0xF3:
		goboy.CPU.Flags.IME = false
		return 4
	case 0xFB:
		goboy.CPU.Flags.PendingInterruptEnabled = true
		return 4
	// END Miscellaneous

	// Calls
	case 0xCD:
		nextAddress := goboy.getParameter16Bit()
		goboy.StackPush(goboy.CPU.Registers.PC)
		goboy.CPU.Registers.PC = nextAddress
		Call = append(Call, fmt.Sprintf("call : %04x", nextAddress))
		return 12
	case 0xC4:
		val := goboy.getParameter16Bit()
		if !goboy.CPU.Flags.Zero {
			goboy.StackPush(goboy.CPU.Registers.PC)
			goboy.CPU.Registers.PC = val
			Call = append(Call, fmt.Sprintf("call : %04x", val))
			ticks += 12
		}
		return 12
	case 0xCC:
		val := goboy.getParameter16Bit()
		if goboy.CPU.Flags.Zero {
			goboy.StackPush(goboy.CPU.Registers.PC)
			goboy.CPU.Registers.PC = val
			Call = append(Call, fmt.Sprintf("call : %04x", val))
			ticks += 12
		}
		return 12
	case 0xD4:
		val := goboy.getParameter16Bit()
		if !goboy.CPU.Flags.Carry {
			goboy.StackPush(goboy.CPU.Registers.PC)
			goboy.CPU.Registers.PC = val
			Call = append(Call, fmt.Sprintf("call : %04x", val))
			ticks += 12
		}
		return 12
	case 0xDC:
		val := goboy.getParameter16Bit()
		if goboy.CPU.Flags.Carry {
			goboy.StackPush(goboy.CPU.Registers.PC)
			goboy.CPU.Registers.PC = val
			Call = append(Call, fmt.Sprintf("call : %04x", val))
			ticks += 12
		}
		return 12
	// END Calls

	// Returns
	case 0xC9:
		goboy.CPU.Registers.PC = goboy.StackPop()
		return 8
	case 0xC0:
		if !goboy.CPU.Flags.Zero {
			goboy.CPU.Registers.PC = goboy.StackPop()
			ticks += 12
		}
		return 8
	case 0xC8:
		if goboy.CPU.Flags.Zero {
			goboy.CPU.Registers.PC = goboy.StackPop()
			ticks += 12
		}
		return 8
	case 0xD0:
		if !goboy.CPU.Flags.Carry {
			goboy.CPU.Registers.PC = goboy.StackPop()
			ticks += 12
		}
		return 8
	case 0xD8:
		if goboy.CPU.Flags.Carry {
			goboy.CPU.Registers.PC = goboy.StackPop()
			ticks += 12
		}
		return 8
	case 0xD9:
		goboy.CPU.Registers.PC = goboy.StackPop()
		goboy.CPU.Flags.IME = true
		return 8
	// END returns

	// Restarts
	case 0xC7:
		goboy.StackPush(goboy.CPU.Registers.PC)
		goboy.CPU.Registers.PC = 0x0000
		return 16
	case 0xCF:
		goboy.StackPush(goboy.CPU.Registers.PC)
		goboy.CPU.Registers.PC = 0x0008
		return 16
	case 0xD7:
		goboy.StackPush(goboy.CPU.Registers.PC)
		goboy.CPU.Registers.PC = 0x0010
		return 16
	case 0xDF:
		goboy.StackPush(goboy.CPU.Registers.PC)
		goboy.CPU.Registers.PC = 0x0018
		return 16
	case 0xE7:
		goboy.StackPush(goboy.CPU.Registers.PC)
		goboy.CPU.Registers.PC = 0x0020
		return 16
	case 0xEF:
		goboy.StackPush(goboy.CPU.Registers.PC)
		goboy.CPU.Registers.PC = 0x0028
		return 16
	case 0xF7:
		goboy.StackPush(goboy.CPU.Registers.PC)
		goboy.CPU.Registers.PC = 0x0030
		return 16
	case 0xFF:
		goboy.StackPush(goboy.CPU.Registers.PC)
		goboy.CPU.Registers.PC = 0x0038
		return 16
	// END Restarts

	// Rotates and Shifts
	case 0x07:
		origin := goboy.CPU.Registers.A
		goboy.CPU.Registers.A = byte(goboy.CPU.Registers.A<<1) | (goboy.CPU.Registers.A >> 7)
		goboy.CPU.Flags.Zero = false
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = false
		goboy.CPU.Flags.Carry = origin > 0x7F
		goboy.CPU.UpdateFlags()
		return 4
	case 0x17:
		carryFlag := byte(0)
		if goboy.CPU.Flags.Carry {
			carryFlag = 1
		}
		goboy.CPU.Flags.Carry = ((goboy.CPU.Registers.A & 0x80) == 0x80)
		goboy.CPU.Registers.A = ((goboy.CPU.Registers.A << 1) & 0xFF) | carryFlag
		goboy.CPU.Flags.Zero = false
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = false
		goboy.CPU.UpdateFlags()
		return 4
	case 0x0F:
		value := goboy.CPU.Registers.A
		goboy.CPU.Registers.A = byte(value>>1) | byte((value&1)<<7)
		goboy.CPU.Flags.Zero = false
		goboy.CPU.Flags.Sub = false
		goboy.CPU.Flags.HalfCarry = false
		goboy.CPU.Flags.Carry = goboy.CPU.Registers.A > 0x7F
		goboy.CPU.UpdateFlags()
		return 4
	case 0x1F:
		value := goboy.CPU.Registers.A
		var carry byte
		if goboy.CPU.Flags.Carry {
			carry = 0x80
		}
		result := byte(value>>1) | carry
		goboy.CPU.Registers.A = result
		goboy.CPU.Flags.Zero = false
		goboy.CPU.Flags.Carry = (1 & value) == 1
		goboy.CPU.Flags.HalfCarry = false
		goboy.CPU.Flags.Sub = false
		goboy.CPU.UpdateFlags()
		return 4
	// END Rotates and Shifts

	//CB Prefixes
	case 0xCB:
		nextIns := goboy.getParameter8Bit()
		ticks += CBCycles[nextIns] * 4
		goboy.cbMap[nextIns]()
		return CBCycles[nextIns] * 4

	//END CB Prefixes

	default:
		fmt.Printf("%02x\n", goboy.ReadMemory(goboy.CPU.Registers.PC-1))
		fmt.Printf("%04x\n", goboy.CPU.Registers.PC-1)
		goboy.Running = false
		return 0
	}
}
