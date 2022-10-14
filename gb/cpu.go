package gb

type CPU struct {
	Registers Registers
	Flags     Flags
	Halted    bool
}

type Registers struct {
	A  byte
	F  byte
	B  byte
	C  byte
	D  byte
	E  byte
	H  byte
	L  byte
	PC uint16
	SP uint16
}

type Flags struct {
	Zero                    bool
	Sub                     bool
	HalfCarry               bool
	Carry                   bool
	PendingInterruptEnabled bool
	IME                     bool //interrupt master enable
}

func (goboy *Goboy) initCPU() {
	goboy.CPU.Flags.Zero = true
	goboy.CPU.Flags.Sub = false
	goboy.CPU.Flags.HalfCarry = true
	goboy.CPU.Flags.Carry = true
	goboy.CPU.Flags.IME = false
	goboy.CPU.Registers.A = 0x01
	goboy.CPU.Registers.B = 0x00
	goboy.CPU.Registers.C = 0x13
	goboy.CPU.Registers.D = 0x00
	goboy.CPU.Registers.E = 0xD8
	goboy.CPU.Registers.F = 0xB0
	goboy.CPU.Registers.H = 0x01
	goboy.CPU.Registers.L = 0x4D
	goboy.CPU.Registers.PC = 0x0100
	goboy.CPU.Registers.SP = 0xFFFE

}

/*

	Opcode execution

*/
var OpcodeCycles = []int{
	1, 3, 2, 2, 1, 1, 2, 1, 5, 2, 2, 2, 1, 1, 2, 1, // 0
	0, 3, 2, 2, 1, 1, 2, 1, 3, 2, 2, 2, 1, 1, 2, 1, // 1
	2, 3, 2, 2, 1, 1, 2, 1, 2, 2, 2, 2, 1, 1, 2, 1, // 2
	2, 3, 2, 2, 3, 3, 3, 1, 2, 2, 2, 2, 1, 1, 2, 1, // 3
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // 4
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // 5
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // 6
	2, 2, 2, 2, 2, 2, 0, 2, 1, 1, 1, 1, 1, 1, 2, 1, // 7
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // 8
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // 9
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // a
	1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2, 1, // b
	2, 3, 3, 4, 3, 4, 2, 4, 2, 4, 3, 0, 3, 6, 2, 4, // c
	2, 3, 3, 0, 3, 4, 2, 4, 2, 4, 3, 0, 3, 0, 2, 4, // d
	3, 3, 2, 0, 0, 4, 2, 4, 4, 1, 4, 0, 0, 0, 2, 4, // e
	3, 3, 2, 1, 0, 4, 2, 4, 3, 2, 4, 1, 0, 0, 2, 4, // f
} //0  1  2  3  4  5  6  7  8  9  a  b  c  d  e  f

func (goboy *Goboy) CPUnextOPcode() int {
	opcode := goboy.ReadMemory(goboy.CPU.Registers.PC)
	goboy.CPU.Registers.PC += 1
	return goboy.ExecuteOpcode(opcode)
}

/*

	Parameter getters

*/

func (goboy *Goboy) getParameter8Bit() byte {
	rByte := goboy.ReadMemory(goboy.CPU.Registers.PC)
	goboy.CPU.Registers.PC += 1
	return rByte
}

func (goboy *Goboy) getParameter16Bit() uint16 {
	lowerbyte := uint16(goboy.ReadMemory(goboy.CPU.Registers.PC))
	upperByte := uint16(goboy.ReadMemory(goboy.CPU.Registers.PC + 1))
	goboy.CPU.Registers.PC += 2
	return upperByte<<8 | lowerbyte
}

/*

	Register Setters

*/
// 8 bit setters
func (goboy *Goboy) setA(val byte) {
	goboy.CPU.Registers.A = val
}

func (goboy *Goboy) setB(val byte) {
	goboy.CPU.Registers.B = val
}

func (goboy *Goboy) setC(val byte) {
	goboy.CPU.Registers.C = val
}

func (goboy *Goboy) setD(val byte) {
	goboy.CPU.Registers.D = val
}

func (goboy *Goboy) setE(val byte) {
	goboy.CPU.Registers.E = val
}

func (goboy *Goboy) setH(val byte) {
	goboy.CPU.Registers.H = val
}

func (goboy *Goboy) setL(val byte) {
	goboy.CPU.Registers.L = val
}

// 16 bit setters

func (cpu *CPU) setAF(val uint16) {
	cpu.Registers.A = byte((val & 0xFF00) >> 8)
	cpu.Registers.F = byte(val & 0xFF)
}

func (cpu *CPU) setBC(val uint16) {
	cpu.Registers.B = byte((val & 0xFF00) >> 8)
	cpu.Registers.C = byte(val & 0xFF)
}

func (cpu *CPU) setDE(val uint16) {
	cpu.Registers.D = byte((val & 0xFF00) >> 8)
	cpu.Registers.E = byte(val & 0xFF)
}

func (cpu *CPU) setHL(val uint16) {
	cpu.Registers.H = byte((val & 0xFF00) >> 8)
	cpu.Registers.L = byte(val & 0xFF)
}

/*

	Register Getters

*/

// 16 bit getters

func (cpu *CPU) GetAF() uint16 {
	return uint16(cpu.Registers.A)<<8 | uint16(cpu.Registers.F)
}

func (cpu *CPU) GetBC() uint16 {
	return uint16(cpu.Registers.B)<<8 | uint16(cpu.Registers.C)
}

func (cpu *CPU) GetDE() uint16 {
	return uint16(cpu.Registers.D)<<8 | uint16(cpu.Registers.E)
}

func (cpu *CPU) GetHL() uint16 {
	return uint16(cpu.Registers.H)<<8 | uint16(cpu.Registers.L)
}

func (cpu *CPU) Compare(val1 byte, val2 byte) {
	cpu.Flags.Zero = (val1 == val2)
	cpu.Flags.Carry = (val1 > val2)
	cpu.Flags.HalfCarry = ((val1 & 0x0f) > (val2 & 0x0f))
	cpu.Flags.Sub = true
	cpu.UpdateFlags()
}

func (cpu *CPU) UpdateFlags() {
	updatedFreg := cpu.Registers.F
	if cpu.Flags.Zero {
		updatedFreg = SetBit(updatedFreg, 7)
	} else {
		updatedFreg = ResetBit(updatedFreg, 7)
	}
	if cpu.Flags.Sub {
		updatedFreg = SetBit(updatedFreg, 6)
	} else {
		updatedFreg = ResetBit(updatedFreg, 6)
	}
	if cpu.Flags.HalfCarry {
		updatedFreg = SetBit(updatedFreg, 5)
	} else {
		updatedFreg = ResetBit(updatedFreg, 5)
	}
	if cpu.Flags.Carry {
		updatedFreg = SetBit(updatedFreg, 4)
	} else {
		updatedFreg = ResetBit(updatedFreg, 4)
	}
	cpu.Registers.F = updatedFreg
}

//Register Getter

func (cpu *CPU) GetRegisterSP() uint16 {
	return cpu.Registers.SP
}

func (cpu *CPU) GetRegisterPC() uint16 {
	return cpu.Registers.PC
}

//Getter Halted

func (cpu *CPU) GetHalted() bool {
	return cpu.Halted
}
