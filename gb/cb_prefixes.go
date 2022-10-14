package gb

var CBCycles = []int{
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // 0
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // 1
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // 2
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // 3
	2, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 2, 2, 2, 3, 2, // 4
	2, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 2, 2, 2, 3, 2, // 5
	2, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 2, 2, 2, 3, 2, // 6
	2, 2, 2, 2, 2, 2, 3, 2, 2, 2, 2, 2, 2, 2, 3, 2, // 7
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // 8
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // 9
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // A
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // B
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // C
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // D
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // E
	2, 2, 2, 2, 2, 2, 4, 2, 2, 2, 2, 2, 2, 2, 4, 2, // F
} //0  1  2  3  4  5  6  7  8  9  a  b  c  d  e  f

func (goboy *Goboy) initCB() {

	var getters = [8]func() byte{
		func() byte { return goboy.CPU.Registers.B },
		func() byte { return goboy.CPU.Registers.C },
		func() byte { return goboy.CPU.Registers.D },
		func() byte { return goboy.CPU.Registers.E },
		func() byte { return goboy.CPU.Registers.H },
		func() byte { return goboy.CPU.Registers.L },
		func() byte { return goboy.ReadMemory(goboy.CPU.GetHL()) },
		func() byte { return goboy.CPU.Registers.A },
	}

	var setters = [8]func(byte){
		goboy.setB,
		goboy.setC,
		goboy.setD,
		goboy.setE,
		goboy.setH,
		goboy.setL,
		func(val byte) { goboy.WriteMemory(goboy.CPU.GetHL(), val) },
		goboy.setA,
	}

	//Every 8 instructions is a group,which use different registers
	for i := 0; i < 8; i++ {

		registerID := i

		goboy.cbMap[0x00+i] = func() {
			goboy.RLC(getters[registerID], setters[registerID])
		}

		goboy.cbMap[0x08+i] = func() {
			goboy.RRC(getters[registerID], setters[registerID])
		}

		goboy.cbMap[0x10+i] = func() {
			goboy.RL(getters[registerID], setters[registerID])
		}

		goboy.cbMap[0x18+i] = func() {
			goboy.RR(getters[registerID], setters[registerID])
		}

		goboy.cbMap[0x20+i] = func() {
			goboy.SLA(getters[registerID], setters[registerID])
		}
		goboy.cbMap[0x28+i] = func() {
			goboy.SRA(getters[registerID], setters[registerID])
		}

		goboy.cbMap[0x30+i] = func() {
			goboy.SWAP(getters[registerID], setters[registerID])
		}

		goboy.cbMap[0x38+i] = func() {
			goboy.SRL(getters[registerID], setters[registerID])
		}

		/*
			RES commands
		*/
		goboy.cbMap[0x80+i] = func() {
			setters[registerID](ResetBit(getters[registerID](), 0))
		}
		goboy.cbMap[0x88+i] = func() {
			setters[registerID](ResetBit(getters[registerID](), 1))
		}
		goboy.cbMap[0x90+i] = func() {
			setters[registerID](ResetBit(getters[registerID](), 2))
		}
		goboy.cbMap[0x98+i] = func() {
			setters[registerID](ResetBit(getters[registerID](), 3))
		}
		goboy.cbMap[0xA0+i] = func() {
			setters[registerID](ResetBit(getters[registerID](), 4))
		}
		goboy.cbMap[0xA8+i] = func() {
			setters[registerID](ResetBit(getters[registerID](), 5))
		}
		goboy.cbMap[0xB0+i] = func() {
			setters[registerID](ResetBit(getters[registerID](), 6))
		}
		goboy.cbMap[0xB8+i] = func() {
			setters[registerID](ResetBit(getters[registerID](), 7))
		}

		/*
			BIT commands
		*/
		goboy.cbMap[0x40+i] = func() {
			goboy.BIT(0, getters[registerID])
		}
		goboy.cbMap[0x48+i] = func() {
			goboy.BIT(1, getters[registerID])
		}
		goboy.cbMap[0x50+i] = func() {
			goboy.BIT(2, getters[registerID])
		}
		goboy.cbMap[0x58+i] = func() {
			goboy.BIT(3, getters[registerID])
		}
		goboy.cbMap[0x60+i] = func() {
			goboy.BIT(4, getters[registerID])
		}
		goboy.cbMap[0x68+i] = func() {
			goboy.BIT(5, getters[registerID])
		}
		goboy.cbMap[0x70+i] = func() {
			goboy.BIT(6, getters[registerID])
		}
		goboy.cbMap[0x78+i] = func() {
			goboy.BIT(7, getters[registerID])
		}

		/*
			Set commands
		*/
		goboy.cbMap[0xC0+i] = func() {
			setters[registerID](SetBit(getters[registerID](), 0))
		}
		goboy.cbMap[0xC8+i] = func() {
			setters[registerID](SetBit(getters[registerID](), 1))
		}
		goboy.cbMap[0xD0+i] = func() {
			setters[registerID](SetBit(getters[registerID](), 2))
		}
		goboy.cbMap[0xD8+i] = func() {
			setters[registerID](SetBit(getters[registerID](), 3))
		}
		goboy.cbMap[0xE0+i] = func() {
			setters[registerID](SetBit(getters[registerID](), 4))
		}
		goboy.cbMap[0xE8+i] = func() {
			setters[registerID](SetBit(getters[registerID](), 5))
		}
		goboy.cbMap[0xF0+i] = func() {
			setters[registerID](SetBit(getters[registerID](), 6))
		}
		goboy.cbMap[0xF8+i] = func() {
			setters[registerID](SetBit(getters[registerID](), 7))
		}
	}
}

func (goboy *Goboy) SRL(getter func() byte, setter func(byte)) {
	val := getter()
	carry := val & 1
	res := val >> 1
	setter(res)

	goboy.CPU.Flags.Zero = (res == 0)
	goboy.CPU.Flags.Sub = false
	goboy.CPU.Flags.HalfCarry = false
	goboy.CPU.Flags.Carry = (carry == 1)
	goboy.CPU.UpdateFlags()
}

func (goboy *Goboy) BIT(pos byte, getter func() byte) {
	val := getter()
	goboy.CPU.Flags.Zero = (val>>pos)&1 == 0
	goboy.CPU.Flags.Sub = false
	goboy.CPU.Flags.HalfCarry = true
	goboy.CPU.UpdateFlags()
}

func (goboy *Goboy) SLA(getter func() byte, setter func(byte)) {
	val := getter()
	carry := val >> 7
	res := (val << 1) & 0xFF
	setter(res)

	goboy.CPU.Flags.Zero = (res == 0)
	goboy.CPU.Flags.Sub = false
	goboy.CPU.Flags.HalfCarry = false
	goboy.CPU.Flags.Carry = (carry == 1)
	goboy.CPU.UpdateFlags()
}

func (goboy *Goboy) SRA(getter func() byte, setter func(byte)) {

	val := getter()

	rot := (val & 128) | (val >> 1)
	setter(rot)

	goboy.CPU.Flags.Zero = (rot == 0)
	goboy.CPU.Flags.Sub = false
	goboy.CPU.Flags.HalfCarry = false
	goboy.CPU.Flags.Carry = (val&1 == 1)
	goboy.CPU.UpdateFlags()
}

func (goboy *Goboy) RLC(getter func() byte, setter func(byte)) {
	val := getter()
	var carry byte
	var rot byte
	carry = val >> 7
	rot = (val<<1)&0xFF | carry
	setter(rot)
	goboy.CPU.Flags.Zero = (rot == 0)
	goboy.CPU.Flags.Sub = false
	goboy.CPU.Flags.HalfCarry = false
	goboy.CPU.Flags.Carry = (carry == 1)
	goboy.CPU.UpdateFlags()
}

func (goboy *Goboy) SWAP(getter func() byte, setter func(byte)) {
	val := getter()
	res := val<<4&240 | val>>4
	setter(res)
	goboy.CPU.Flags.Zero = (res == 0)
	goboy.CPU.Flags.Sub = false
	goboy.CPU.Flags.HalfCarry = false
	goboy.CPU.Flags.Carry = false
	goboy.CPU.UpdateFlags()
}

func (goboy *Goboy) RRC(getter func() byte, setter func(byte)) {
	val := getter()
	carry := val & 1
	rot := (val >> 1) | (carry << 7)
	setter(rot)
	goboy.CPU.Flags.Zero = (rot == 0)
	goboy.CPU.Flags.Sub = false
	goboy.CPU.Flags.HalfCarry = false
	goboy.CPU.Flags.Carry = (carry == 1)
	goboy.CPU.UpdateFlags()
}

func (goboy *Goboy) RR(getter func() byte, setter func(byte)) {
	val := getter()
	carry := val & 1
	oldCarry := byte(0)
	if goboy.CPU.Flags.Carry {
		oldCarry = 1
	}
	rot := (val >> 1) | (oldCarry << 7)
	setter(rot)

	goboy.CPU.Flags.Zero = (rot == 0)
	goboy.CPU.Flags.Sub = false
	goboy.CPU.Flags.HalfCarry = false
	goboy.CPU.Flags.Carry = (carry == 1)
	goboy.CPU.UpdateFlags()
}

func (goboy *Goboy) RL(getter func() byte, setter func(byte)) {
	val := getter()
	oldCarry := byte(0)
	if goboy.CPU.Flags.Carry {
		oldCarry = 1
	}

	newCarry := val >> 7
	rot := (val<<1)&0xFF | oldCarry
	setter(rot)

	goboy.CPU.Flags.Zero = (rot == 0)
	goboy.CPU.Flags.Sub = false
	goboy.CPU.Flags.HalfCarry = false
	goboy.CPU.Flags.Carry = (newCarry == 1)
	goboy.CPU.UpdateFlags()
}
