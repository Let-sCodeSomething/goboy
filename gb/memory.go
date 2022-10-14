package gb

import "time"

type Memory struct {
	NewData    bool
	MainMemory [0x10000]byte
}

var Mem string

func (goboy *Goboy) ReadMemory(addr uint16) byte {
	if (addr >= 0x4000) && (addr <= 0x7FFF) {
		// are we reading from the rom memory bank?
		return goboy.Cartridge.MBC.ReadRomBank(addr)
	} else if (addr >= 0xA000) && (addr <= 0xBFFF) {
		// are we reading from ram memory bank?
		return goboy.Cartridge.MBC.ReadRamBank(addr)
	} else if addr == 0xFF00 {
		res := goboy.Memory.MainMemory[0xFF00]
		// flip all the bits
		res ^= 0xFF

		if !CheckBit(res, 4) {
			topJoypad := goboy.Joypad >> 4
			topJoypad |= 0xF0
			res &= topJoypad
		} else if !CheckBit(res, 5) {
			bottomJoypad := goboy.Joypad & 0xF
			bottomJoypad |= 0xF0
			res &= bottomJoypad
		}
		return res
	}
	return goboy.Memory.MainMemory[addr]

}

func (goboy *Goboy) WriteMemory(addr uint16, data byte) {
	if addr < 0x8000 {
		goboy.Cartridge.MBC.HandleBanking(addr, data)
	} else if (addr >= 0xA000) && (addr < 0xC000) {
		goboy.Cartridge.MBC.WriteRamBank(addr, data)
		goboy.Memory.NewData = true
	} else if (addr >= 0xE000) && (addr < 0xFE00) {
		// writing to ECHO ram also writes in RAM
		goboy.Memory.MainMemory[addr] = data
		goboy.WriteMemory(addr-0x2000, data)
	} else if addr == 0xFF44 {
		goboy.Memory.MainMemory[0xFF44] = 0
	} else if addr == 0xFF46 {
		destination := uint16(data) << 8
		for i := 0; i < 0xA0; i++ {
			goboy.WriteMemory(0xFE00+uint16(i), goboy.ReadMemory(destination+uint16(i)))
		}
	} else if addr == 0xFF04 {
		goboy.Memory.MainMemory[0xFF04] = 0
		goboy.SetClockFreq()
	} else if addr == 0xFF05 {
		goboy.Memory.MainMemory[0xFF05] = data
	} else if addr == 0xFF06 {
		goboy.Memory.MainMemory[0xFF06] = data
	} else if addr == 0xFF07 {
		origin := goboy.Memory.MainMemory[0xFF07]
		currentFreq := goboy.GetClockFreq()
		goboy.Memory.MainMemory[0xFF07] = data | 0xF8
		newFreq := goboy.GetClockFreq()
		if CheckBit(origin, 2) && !CheckBit(goboy.Memory.MainMemory[0xFF07], 2) && CheckBit(goboy.Memory.MainMemory[0xFF04], 1) {
			goboy.Timer.TimerCounter = goboy.GetClockFreqCount()
		}
		if currentFreq != newFreq {
			goboy.SetClockFreq()
		}
	} else if addr == 0xFF00 {
		actual := goboy.Memory.MainMemory[0xFF00]
		if CheckBit(data, 4) {
			actual = SetBit(actual, 4)
		}
		if CheckBit(data, 5) {
			actual = SetBit(actual, 5)
		}
		if !CheckBit(data, 4) {
			actual = ResetBit(actual, 4)
		}
		if !CheckBit(data, 5) {
			actual = ResetBit(actual, 5)
		}
		goboy.Memory.MainMemory[0xFF00] = actual
	} else if addr == 0xFF02 {
		// later handle serial transfer using $FF01 and $FF02
		actual := goboy.Memory.MainMemory[0xFF02]
		actual |= 0x7E
		if CheckBit(data, 7) {
			actual = SetBit(actual, 7)
		}
		if CheckBit(data, 0) {
			actual = SetBit(actual, 0)
		}
		if !CheckBit(data, 7) {
			actual = ResetBit(actual, 7)
		}
		if !CheckBit(data, 0) {
			actual = ResetBit(actual, 0)
		}
		goboy.Memory.MainMemory[0xFF02] = actual
	} else if addr == 0xFF0F {
		actual := goboy.Memory.MainMemory[0xFF0F]
		actual |= 0xE0
		if CheckBit(data, 0) {
			actual = SetBit(actual, 0)
		}
		if CheckBit(data, 1) {
			actual = SetBit(actual, 1)
		}
		if CheckBit(data, 2) {
			actual = SetBit(actual, 2)
		}
		if CheckBit(data, 3) {
			actual = SetBit(actual, 3)
		}
		if CheckBit(data, 4) {
			actual = SetBit(actual, 4)
		}
		if !CheckBit(data, 0) {
			actual = ResetBit(actual, 0)
		}
		if !CheckBit(data, 1) {
			actual = ResetBit(actual, 1)
		}
		if !CheckBit(data, 2) {
			actual = ResetBit(actual, 2)
		}
		if !CheckBit(data, 3) {
			actual = ResetBit(actual, 3)
		}
		if !CheckBit(data, 4) {
			actual = ResetBit(actual, 4)
		}
		goboy.Memory.MainMemory[0xFF0F] = actual
	} else if addr == 0xFF41 {
		actual := goboy.Memory.MainMemory[0xFF41]
		actual |= 0x80
		if CheckBit(data, 0) {
			actual = SetBit(actual, 0)
		}
		if CheckBit(data, 1) {
			actual = SetBit(actual, 1)
		}
		if CheckBit(data, 2) {
			actual = SetBit(actual, 2)
		}
		if CheckBit(data, 3) {
			actual = SetBit(actual, 3)
		}
		if CheckBit(data, 4) {
			actual = SetBit(actual, 4)
		}
		if CheckBit(data, 5) {
			actual = SetBit(actual, 5)
		}
		if CheckBit(data, 6) {
			actual = SetBit(actual, 6)
		}
		if !CheckBit(data, 0) {
			actual = ResetBit(actual, 0)
		}
		if !CheckBit(data, 1) {
			actual = ResetBit(actual, 1)
		}
		if !CheckBit(data, 2) {
			actual = ResetBit(actual, 2)
		}
		if !CheckBit(data, 3) {
			actual = ResetBit(actual, 3)
		}
		if !CheckBit(data, 4) {
			actual = ResetBit(actual, 4)
		}
		if !CheckBit(data, 5) {
			actual = ResetBit(actual, 5)
		}
		if !CheckBit(data, 6) {
			actual = ResetBit(actual, 6)
		}
		goboy.Memory.MainMemory[0xFF41] = actual
	} else if addr == 0xFF10 {
		actual := goboy.Memory.MainMemory[0xFF10]
		actual |= 0x80
		if CheckBit(data, 0) {
			actual = SetBit(actual, 0)
		}
		if CheckBit(data, 1) {
			actual = SetBit(actual, 1)
		}
		if CheckBit(data, 2) {
			actual = SetBit(actual, 2)
		}
		if CheckBit(data, 3) {
			actual = SetBit(actual, 3)
		}
		if CheckBit(data, 4) {
			actual = SetBit(actual, 4)
		}
		if CheckBit(data, 5) {
			actual = SetBit(actual, 5)
		}
		if CheckBit(data, 6) {
			actual = SetBit(actual, 6)
		}
		if !CheckBit(data, 0) {
			actual = ResetBit(actual, 0)
		}
		if !CheckBit(data, 1) {
			actual = ResetBit(actual, 1)
		}
		if !CheckBit(data, 2) {
			actual = ResetBit(actual, 2)
		}
		if !CheckBit(data, 3) {
			actual = ResetBit(actual, 3)
		}
		if !CheckBit(data, 4) {
			actual = ResetBit(actual, 4)
		}
		if !CheckBit(data, 5) {
			actual = ResetBit(actual, 5)
		}
		if !CheckBit(data, 6) {
			actual = ResetBit(actual, 6)
		}
		goboy.Memory.MainMemory[0xFF10] = actual
	} else if addr == 0xFF1A {
		actual := goboy.Memory.MainMemory[0xFF1A]
		actual |= 0x7F
		if CheckBit(data, 7) {
			actual = SetBit(actual, 7)
		}
		if !CheckBit(data, 7) {
			actual = ResetBit(actual, 7)
		}
		goboy.Memory.MainMemory[0xFF1A] = actual
	} else if addr == 0xFF1C {
		actual := goboy.Memory.MainMemory[0xFF1C]
		actual |= 0x9F
		if CheckBit(data, 5) {
			actual = SetBit(actual, 5)
		}
		if CheckBit(data, 6) {
			actual = SetBit(actual, 6)
		}
		if !CheckBit(data, 5) {
			actual = ResetBit(actual, 5)
		}
		if !CheckBit(data, 6) {
			actual = ResetBit(actual, 6)
		}
		goboy.Memory.MainMemory[0xFF1C] = actual
	} else if addr == 0xFF20 {
		actual := goboy.Memory.MainMemory[0xFF20]
		actual |= 0xC0
		if CheckBit(data, 0) {
			actual = SetBit(actual, 0)
		}
		if CheckBit(data, 1) {
			actual = SetBit(actual, 1)
		}
		if CheckBit(data, 2) {
			actual = SetBit(actual, 2)
		}
		if CheckBit(data, 3) {
			actual = SetBit(actual, 3)
		}
		if CheckBit(data, 4) {
			actual = SetBit(actual, 4)
		}
		if CheckBit(data, 5) {
			actual = SetBit(actual, 5)
		}
		if !CheckBit(data, 0) {
			actual = ResetBit(actual, 0)
		}
		if !CheckBit(data, 1) {
			actual = ResetBit(actual, 1)
		}
		if !CheckBit(data, 2) {
			actual = ResetBit(actual, 2)
		}
		if !CheckBit(data, 3) {
			actual = ResetBit(actual, 3)
		}
		if !CheckBit(data, 4) {
			actual = ResetBit(actual, 4)
		}
		if !CheckBit(data, 5) {
			actual = ResetBit(actual, 5)
		}
		goboy.Memory.MainMemory[0xFF20] = actual
	} else if addr == 0xFF23 {
		actual := goboy.Memory.MainMemory[0xFF23]
		actual |= 0x3F
		if CheckBit(data, 7) {
			actual = SetBit(actual, 7)
		}
		if CheckBit(data, 6) {
			actual = SetBit(actual, 6)
		}
		if !CheckBit(data, 7) {
			actual = ResetBit(actual, 7)
		}
		if !CheckBit(data, 6) {
			actual = ResetBit(actual, 6)
		}
		goboy.Memory.MainMemory[0xFF23] = actual
	} else if addr == 0xFF26 {
		actual := goboy.Memory.MainMemory[0xFF26]
		actual |= 0x70
		if CheckBit(data, 0) {
			actual = SetBit(actual, 0)
		}
		if CheckBit(data, 1) {
			actual = SetBit(actual, 1)
		}
		if CheckBit(data, 2) {
			actual = SetBit(actual, 2)
		}
		if CheckBit(data, 3) {
			actual = SetBit(actual, 3)
		}
		if CheckBit(data, 7) {
			actual = SetBit(actual, 7)
		}
		if !CheckBit(data, 0) {
			actual = ResetBit(actual, 0)
		}
		if !CheckBit(data, 1) {
			actual = ResetBit(actual, 1)
		}
		if !CheckBit(data, 2) {
			actual = ResetBit(actual, 2)
		}
		if !CheckBit(data, 3) {
			actual = ResetBit(actual, 3)
		}
		if !CheckBit(data, 7) {
			actual = ResetBit(actual, 7)
		}
		goboy.Memory.MainMemory[0xFF26] = actual
	} else if addr == 0xFF03 || (addr >= 0xFF08 && addr <= 0xFF0E) || addr == 0xFF15 || addr == 0xFF1F || (addr >= 0xFF27 && addr <= 0xFF2F) || (addr >= 0xFF4C && addr <= 0xFF7F) {
		// special addresses or protected ones
		goboy.Memory.MainMemory[addr] = 0xFF
	} else {
		goboy.Memory.MainMemory[addr] = data
	}

}

func (goboy *Goboy) initMemory() {
	for i := 0x0000; i < goboy.Cartridge.ROMLength && i < 0x8000; i++ {
		goboy.Memory.MainMemory[i] = goboy.Cartridge.MBC.ReadRom(uint16(i))
	}
	for i := 0xFF00; i <= 0xFF7F; i++ {
		goboy.Memory.MainMemory[i] = 0xFF
	}

	goboy.Memory.MainMemory[0xFF00] = 0x00
	goboy.Memory.MainMemory[0xFF05] = 0x00
	goboy.Memory.MainMemory[0xFF06] = 0x00
	goboy.Memory.MainMemory[0xFF07] = 0x00
	goboy.Memory.MainMemory[0xFF0F] = 0xE1
	goboy.Memory.MainMemory[0xFF10] = 0x80
	goboy.Memory.MainMemory[0xFF11] = 0xBF
	goboy.Memory.MainMemory[0xFF12] = 0xF3
	goboy.Memory.MainMemory[0xFF14] = 0xBF
	goboy.Memory.MainMemory[0xFF16] = 0x3F
	goboy.Memory.MainMemory[0xFF17] = 0x00
	goboy.Memory.MainMemory[0xFF19] = 0xBF
	goboy.Memory.MainMemory[0xFF1A] = 0x7F
	goboy.Memory.MainMemory[0xFF1B] = 0xFF
	goboy.Memory.MainMemory[0xFF1C] = 0x9F
	goboy.Memory.MainMemory[0xFF1E] = 0xBF
	goboy.Memory.MainMemory[0xFF20] = 0xFF
	goboy.Memory.MainMemory[0xFF21] = 0x00
	goboy.Memory.MainMemory[0xFF22] = 0x00
	goboy.Memory.MainMemory[0xFF23] = 0xBF
	goboy.Memory.MainMemory[0xFF24] = 0x77
	goboy.Memory.MainMemory[0xFF25] = 0xF3
	goboy.Memory.MainMemory[0xFF26] = 0xF1
	goboy.Memory.MainMemory[0xFF40] = 0x91
	goboy.Memory.MainMemory[0xFF41] = 0x85
	goboy.Memory.MainMemory[0xFF42] = 0x00
	goboy.Memory.MainMemory[0xFF43] = 0x00
	goboy.Memory.MainMemory[0xFF44] = 0x00
	goboy.Memory.MainMemory[0xFF45] = 0x00
	goboy.Memory.MainMemory[0xFF46] = 0xFC
	goboy.Memory.MainMemory[0xFF47] = 0xFC
	goboy.Memory.MainMemory[0xFF48] = 0xFF
	goboy.Memory.MainMemory[0xFF49] = 0xFF
	goboy.Memory.MainMemory[0xFF4A] = 0x00
	goboy.Memory.MainMemory[0xFF4B] = 0x00
	goboy.Memory.MainMemory[0xFFFF] = 0x00

	saveTimer := time.Tick(time.Second)
	go func() {
		for range saveTimer {
			if goboy.Memory.NewData {
				goboy.Memory.NewData = false
				goboy.Cartridge.MBC.SaveRam(goboy.RAMCartridgePath)
			}
		}
	}()
}

func (goboy *Goboy) StackPush(val uint16) {
	hi := val >> 8
	lo := val & 0xFF
	goboy.CPU.Registers.SP--
	goboy.WriteMemory(goboy.CPU.Registers.SP, byte(hi))
	goboy.CPU.Registers.SP--
	goboy.WriteMemory(goboy.CPU.Registers.SP, byte(lo))
}

func (goboy *Goboy) StackPop() uint16 {
	lo := goboy.ReadMemory(goboy.CPU.Registers.SP)
	hi := goboy.ReadMemory(goboy.CPU.Registers.SP + 1)
	goboy.CPU.Registers.SP += 2
	return uint16(lo) + (uint16(hi) << 8)
}
