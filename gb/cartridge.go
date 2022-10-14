package gb

import (
	"bufio"
	"log"
	"os"
)

type Cartridge struct {
	CartridgeName string
	MBCType       string
	ROMLength     int
	ROMBank       uint8
	RAMBank       uint8
	MBC           MBC
}

func (goboy *Goboy) readRamFile(ramPath string) []byte {
	return loadFile(ramPath, true)
}
func (goboy *Goboy) readRomFile(romPath string) []byte {
	return loadFile(romPath, false)
}

func loadFile(path string, ram bool) []byte {
	romFile, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) && ram {
			return nil
		}

		log.Fatal(err)
	}
	defer romFile.Close()

	stats, statsErr := romFile.Stat()
	if statsErr != nil {
		log.Fatal(statsErr)
	}
	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufReader := bufio.NewReader(romFile)
	_, err = bufReader.Read(bytes)

	return bytes
}

var cartridgeTypeMap = map[byte]string{
	byte(0x00): "ROM ONLY",
	byte(0x01): "MBC1",
	byte(0x02): "MBC1+RAM",
	byte(0x03): "MBC1+RAM+BATTERY",
	byte(0x05): "MBC2",
	byte(0x06): "MBC2+BATTERY",
	byte(0x08): "ROM+RAM",
	byte(0x09): "ROM+RAM+BATTERY",
	byte(0x0B): "MMM01",
	byte(0x0C): "MMM01+RAM",
	byte(0x0D): "MMM01+RAM+BATTERY",
	byte(0x0F): "MBC3+TIMER+BATTERY",
	byte(0x10): "MBC3+TIMER+RAM+BATTERY",
	byte(0x11): "MBC3",
	byte(0x12): "MBC3+RAM",
	byte(0x13): "MBC3+RAM+BATTERY",
	byte(0x15): "MBC4",
	byte(0x16): "MBC4+RAM",
	byte(0x17): "MBC4+RAM+BATTERY",
	byte(0x19): "MBC5",
	byte(0x1A): "MBC5+RAM",
	byte(0x1B): "MBC5+RAM+BATTERY",
	byte(0x1C): "MBC5+RUMBLE",
	byte(0x1D): "MBC5+RUMBLE+RAM",
	byte(0x1E): "MBC5+RUMBLE+RAM+BATTERY",
	byte(0xFC): "POCKET CAMERA",
	byte(0xFD): "BANDAI TAMA5",
	byte(0xFE): "HuC3",
	byte(0xFF): "HuC1+RAM+BATTERY",
}

var RomBankMap = map[byte]uint8{
	byte(0x00): 2,
	byte(0x01): 4,
	byte(0x02): 8,
	byte(0x03): 16,
	byte(0x04): 32,
	byte(0x05): 64,
	byte(0x06): 128,
	byte(0x52): 72,
	byte(0x53): 80,
	byte(0x54): 96,
}

var RamBankMap = map[byte]uint8{
	byte(0x00): 0,
	byte(0x01): 1,
	byte(0x02): 1,
	byte(0x03): 4,
}

type MBC interface {
	ReadRom(uint16) byte
	ReadRomBank(uint16) byte
	ReadRamBank(uint16) byte
	WriteRamBank(uint16, byte)
	HandleBanking(uint16, byte)
	SaveRam(string)
}

type MBCRom struct {
	Rom            []byte
	CurrentROMBank byte
	RAMBank        [0x8000]byte
	CurrentRAMBank byte
	EnableRAM      bool
}

func (mbc *MBCRom) ReadRamBank(address uint16) byte {
	return byte(0x00)
}

func (mbc *MBCRom) WriteRamBank(address uint16, data byte) {

}

func (mbc *MBCRom) ReadRomBank(address uint16) byte {
	return mbc.Rom[address]
}

func (mbc *MBCRom) ReadRom(address uint16) byte {
	return mbc.Rom[address]
}

func (mbc *MBCRom) HandleBanking(address uint16, val byte) {
}

func (mbc *MBCRom) SaveRam(path string) {
}

type MBC1 struct {
	Rom            []byte
	CurrentROMBank byte
	RAMBank        []byte
	CurrentRAMBank byte
	EnableRAM      bool
	ROMBankingMode bool
}

func (mbc *MBC1) ReadRomBank(address uint16) byte {
	newAddress := uint32(address - 0x4000)
	return mbc.Rom[newAddress+uint32(mbc.CurrentROMBank)*0x4000]
}

func (mbc *MBC1) ReadRamBank(address uint16) byte {
	newAddress := uint32(address - 0xA000)
	return mbc.RAMBank[newAddress+(uint32(mbc.CurrentRAMBank)*0x2000)]
}

func (mbc *MBC1) WriteRamBank(address uint16, data byte) {
	if mbc.EnableRAM {
		newAddress := uint32(address - 0xA000)
		mbc.RAMBank[newAddress+(uint32(mbc.CurrentRAMBank)*0x2000)] = data
	}
}

func (mbc *MBC1) ReadRom(address uint16) byte {
	return mbc.Rom[address]
}

func (mbc *MBC1) HandleBanking(address uint16, val byte) {
	// do RAM enabling
	if address < 0x2000 {
		mbc.DoRamBankEnable(address, val)
	} else if (address >= 0x2000) && (address < 0x4000) {
		mbc.DoChangeLoROMBank(val)
	} else if (address >= 0x4000) && (address < 0x6000) {
		if mbc.ROMBankingMode {
			mbc.DoChangeHiRomBank(val)
		} else {
			mbc.DoRAMBankChange(val)
		}

	} else if (address >= 0x6000) && (address < 0x8000) {
		mbc.DoChangeROMRAMMode(val)
	}
}

func (mbc *MBC1) DoRamBankEnable(address uint16, val byte) {
	testData := val & 0xF
	if testData == 0xA {
		mbc.EnableRAM = true
	} else if testData == 0x0 {
		mbc.EnableRAM = false
	}
}

func (mbc *MBC1) DoChangeLoROMBank(val byte) {
	lower5 := val & 31
	mbc.CurrentROMBank &= 224
	mbc.CurrentROMBank |= lower5
	if mbc.CurrentROMBank == 0 {
		mbc.CurrentROMBank++
	}
}

func (mbc *MBC1) DoChangeHiRomBank(val byte) {
	mbc.CurrentROMBank &= 31
	val &= 224
	mbc.CurrentROMBank |= val
	if mbc.CurrentROMBank == 0 {
		mbc.CurrentROMBank++
	}
}

func (mbc *MBC1) DoRAMBankChange(val byte) {
	mbc.CurrentRAMBank = val & 0x3
}

func (mbc *MBC1) DoChangeROMRAMMode(val byte) {
	newData := val & 0x1
	if newData == 0 {
		mbc.ROMBankingMode = true
	} else {
		mbc.ROMBankingMode = false
	}
	if mbc.ROMBankingMode {
		mbc.CurrentRAMBank = 0
	}
}
func (mbc *MBC1) SaveRam(path string) {
}

type MBC2 struct {
	rom            []byte
	CurrentROMBank byte
	RAMBank        []byte
	CurrentRAMBank byte
	EnableRAM      bool
	ROMBankingMode bool
}

func (mbc *MBC2) ReadRomBank(address uint16) byte {
	newAddress := uint32(address - 0x4000)
	return mbc.rom[newAddress+uint32(mbc.CurrentROMBank)*0x4000]
}

func (mbc *MBC2) ReadRamBank(address uint16) byte {
	newAddress := uint32(address - 0xA000)
	return mbc.RAMBank[newAddress+(uint32(mbc.CurrentRAMBank)*0x2000)]
}

func (mbc *MBC2) WriteRamBank(address uint16, data byte) {
	if mbc.EnableRAM {
		newAddress := uint32(address - 0xA000)
		mbc.RAMBank[newAddress+(uint32(mbc.CurrentRAMBank)*0x2000)] = data
	}
}

func (mbc *MBC2) ReadRom(address uint16) byte {
	return mbc.rom[address]
}
func (mbc *MBC2) HandleBanking(address uint16, val byte) {
	// do RAM enabling
	if address < 0x2000 {
		mbc.DoRamBankEnable(address, val)
	} else if (address >= 0x2000) && (address < 0x4000) {
		mbc.DoChangeLoROMBank(val)

	} else if (address >= 0x4000) && (address < 0x6000) {
		if mbc.ROMBankingMode {
			mbc.DoChangeHiRomBank(val)
		} else {
			mbc.DoRAMBankChange(val)
		}

	} else if (address >= 0x6000) && (address < 0x8000) {
		mbc.DoChangeROMRAMMode(val)
	}
}

func (mbc *MBC2) DoRamBankEnable(address uint16, val byte) {

	if CheckBit(byte(address&0xFF), 4) == true {
		return
	}

	testData := val & 0xF
	if testData == 0xA {
		mbc.EnableRAM = true
	} else if testData == 0x0 {
		mbc.EnableRAM = false
	}
}

func (mbc *MBC2) DoChangeLoROMBank(val byte) {
	mbc.CurrentROMBank = val & 0xF
	if mbc.CurrentROMBank == 0 {
		mbc.CurrentROMBank++
	}
}

func (mbc *MBC2) DoChangeHiRomBank(val byte) {
	mbc.CurrentROMBank &= 31
	val &= 224
	mbc.CurrentROMBank |= val
	if mbc.CurrentROMBank == 0 {
		mbc.CurrentROMBank++
	}
}

func (mbc *MBC2) DoRAMBankChange(val byte) {
	mbc.CurrentRAMBank = val & 0x3
}

func (mbc *MBC2) DoChangeROMRAMMode(val byte) {
	newData := val & 0x1
	if newData == 0 {
		mbc.ROMBankingMode = true
	} else {
		mbc.ROMBankingMode = false
	}
	if mbc.ROMBankingMode {
		mbc.CurrentRAMBank = 0
	}
}

func (mbc *MBC2) SaveRam(path string) {
	writeRamFile(path, mbc.RAMBank)
}

type MBC3 struct {
	rom            []byte
	CurrentROMBank byte
	RAMBank        []byte
	CurrentRAMBank byte
	EnableRAM      bool

	rtc        []byte
	latchedRtc []byte
	latched    bool
}

func (mbc *MBC3) ReadRomBank(address uint16) byte {
	newAddress := uint32(address - 0x4000)
	return mbc.rom[newAddress+uint32(mbc.CurrentROMBank)*0x4000]
}

func (mbc *MBC3) ReadRamBank(address uint16) byte {
	if mbc.CurrentRAMBank >= 0x4 {
		if mbc.latched {
			return mbc.latchedRtc[mbc.CurrentRAMBank]
		}
		return mbc.rtc[mbc.CurrentRAMBank]
	}
	newAddress := uint32(address - 0xA000)
	return mbc.RAMBank[newAddress+(uint32(mbc.CurrentRAMBank)*0x2000)]
}

func (mbc *MBC3) WriteRamBank(address uint16, data byte) {
	if mbc.EnableRAM {
		if mbc.CurrentRAMBank >= 0x4 {
			mbc.rtc[mbc.CurrentRAMBank] = data
		} else {
			newAddress := uint32(address - 0xA000)
			mbc.RAMBank[newAddress+(uint32(mbc.CurrentRAMBank)*0x2000)] = data
		}
	}
}

func (mbc *MBC3) ReadRom(address uint16) byte {
	return mbc.rom[address]
}
func (mbc *MBC3) HandleBanking(address uint16, val byte) {
	if address < 0x2000 {
		mbc.DoRamBankEnable(address, val)
	} else if (address >= 0x2000) && (address < 0x4000) {
		mbc.DoChangeLoROMBank(val)

	} else if (address >= 0x4000) && (address < 0x6000) {
		mbc.DoRAMBankChange(val)

	} else if (address >= 0x6000) && (address < 0x8000) {
		mbc.DoChangeROMRAMMode(val)
	}
}

func (mbc *MBC3) DoRamBankEnable(address uint16, val byte) {
	testData := val & 0xA
	if testData != 0 {
		mbc.EnableRAM = true
	} else if testData == 0x0 {
		mbc.EnableRAM = false
	}
}

func (mbc *MBC3) DoChangeLoROMBank(val byte) {
	lower5 := val & 0x7F
	mbc.CurrentROMBank = lower5
	mbc.CurrentROMBank |= lower5
	if mbc.CurrentROMBank == 0x00 {
		mbc.CurrentROMBank++
	}
}

func (mbc *MBC3) DoChangeHiRomBank(val byte) {
	mbc.CurrentROMBank &= 31
	val &= 224
	mbc.CurrentROMBank |= val
}

func (mbc *MBC3) DoRAMBankChange(val byte) {
	mbc.CurrentRAMBank = val
}

func (mbc *MBC3) DoChangeROMRAMMode(val byte) {
	if val == 0x1 {
		mbc.latched = false
	} else if val == 0x0 {
		mbc.latched = true
		copy(mbc.rtc, mbc.latchedRtc)
	}
}

func (mbc *MBC3) SaveRam(path string) {
	writeRamFile(path, mbc.RAMBank)
}

func writeRamFile(ramPath string, data []byte) {
	ramFile, err := os.Create(ramPath)
	if err != nil {
		log.Fatal(err)
	}
	defer ramFile.Close()

	bufWriter := bufio.NewWriter(ramFile)
	_, err = bufWriter.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}
