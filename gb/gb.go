package gb

import (
	"fmt"
	sc "github.com/Let-sCodeSomething/goboy/screen"
	"log"
	"time"
)

// Goboy Represents the Game boy
type Goboy struct {
	Cartridge        Cartridge
	ROMCartridgePath string
	RAMCartridgePath string
	Memory           Memory
	CPU              CPU
	PPU              PPU
	APU              APU
	Timer            Timer
	ClockMultiplier  float32
	FPS              int
	Clock            int
	Paused           bool
	Running          bool
	cbMap            [0x100](func())
	Screen           [160][144][3]uint8
	ScanLineBG       [160]bool
	Joypad           byte
	mode             string
	Control          sc.Control
}

// WindowMode open in WindowMode
func (goboy *Goboy) WindowMode() {
	screen := sc.CreateScreen(goboy.Cartridge.CartridgeName, &goboy.Screen, &goboy.Paused)
	goboy.Control = &screen.Control
	screen.DebuggingData = sc.DebugData{CPU: &goboy.CPU, FF: func() byte { return goboy.ReadMemory(0xFF00) }, Joypad: &goboy.Joypad}
	fmt.Println(goboy.mode)
	sc.RunScreen(screen, func() bool { return goboy.mode == "debug" }())
}

//API

// SetControl set the key control
func (goboy *Goboy) SetControl(control sc.Control) {
	goboy.Control = control
}

// APIMode is for get current screen
func (goboy *Goboy) APIMode() *[160][144][3]uint8 {
	return &goboy.Screen
}

func (goboy *Goboy) StopGoboy() {
	goboy.Running = false
}

// Init initializes the Game Boy and loads the ROM
func (goboy *Goboy) Init(cartridgePath, mode string) error {
	if cartridgePath == "" {
		var err error
		cartridgePath, err = LoadRomDialog()
		if err != nil {
			return err
		}
	}
	goboy.ROMCartridgePath = cartridgePath
	goboy.RAMCartridgePath = goboy.ROMCartridgePath + ".sav"
	goboy.initRom(goboy.ROMCartridgePath)
	goboy.initMemory()
	goboy.initCPU()
	goboy.initCB()
	goboy.initPPU()
	goboy.initAPU()
	goboy.initTimer()
	goboy.mode = mode
	goboy.ClockMultiplier = 1
	goboy.FPS = 60
	goboy.Clock = 4194304

	return nil
}

var LastInstructions []string

func (goboy *Goboy) Delay(ms int) {
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

// Run starts
func (goboy *Goboy) Run() {
	goboy.Running = true
	ticker := time.NewTicker(time.Second / time.Duration(goboy.FPS))
	for range ticker.C {
		goboy.Update()
		// goboy.ASCII.Run(goboy.DrawSignal, func() { fmt.Println("exit") })
		if !goboy.Running {
			break
		}
	}
}

// Debug print the debug info
func (goboy *Goboy) Debug() {
	fmt.Printf("%04X\n", goboy.CPU.GetAF())
	fmt.Printf("%04X\n", goboy.CPU.GetBC())
	fmt.Printf("%04X\n", goboy.CPU.GetDE())
	fmt.Printf("%04X\n", goboy.CPU.GetHL())
	fmt.Printf("%04X\n", goboy.CPU.Registers.SP)
	fmt.Printf("%04X\n", goboy.CPU.Registers.PC)
}

// Update updates the gameboy
func (goboy *Goboy) Update() {
	goboy.Joypad = 0xFF
	if goboy.Control.GetReload() {
		goboy.Paused = true
		err := goboy.Init("", goboy.mode)
		if err != nil {
			fmt.Println(err)
		}
		goboy.Paused = false
		return
	}
	if goboy.Control.GetStart() {
		goboy.Joypad = ResetBit(goboy.Joypad, 7)
	}
	if goboy.Control.GetSelect() {
		goboy.Joypad = ResetBit(goboy.Joypad, 6)
	}
	if goboy.Control.GetA() {
		goboy.Joypad = ResetBit(goboy.Joypad, 4)
	}
	if goboy.Control.GetB() {
		goboy.Joypad = ResetBit(goboy.Joypad, 5)
	}
	if goboy.Control.GetUp() {
		goboy.Joypad = ResetBit(goboy.Joypad, 2)
	}
	if goboy.Control.GetDown() {
		goboy.Joypad = ResetBit(goboy.Joypad, 3)
	}
	if goboy.Control.GetLeft() {
		goboy.Joypad = ResetBit(goboy.Joypad, 1)
	}
	if goboy.Control.GetRigth() {
		goboy.Joypad = ResetBit(goboy.Joypad, 0)
	}
	allcycles := 0
	if !goboy.Paused {
		for allcycles < int(goboy.Clock/goboy.FPS) {
			cycles := 4
			if !goboy.CPU.Halted {
				cycles = goboy.CPUnextOPcode()
			}
			allcycles += cycles
			goboy.UpdateGraphics(cycles)
			goboy.UpdateTimers(cycles)
			allcycles += goboy.CheckInterrupt()
			// if len(LastInstructions) == 100 {
			// 	LastInstructions = LastInstructions[1:]
			// }
			// LastInstructions = append(LastInstructions, fmt.Sprintf("%04X with value %02X %02X %02X \t flag values Z: %t N: %t H: %t C: %t\n\t Registers :\n%v%v%v%v%v%v\n", goboy.CPU.Registers.PC, goboy.ReadMemory(goboy.CPU.Registers.PC), goboy.ReadMemory(goboy.CPU.Registers.PC+1), goboy.ReadMemory(goboy.CPU.Registers.PC+2), goboy.CPU.Flags.Zero, goboy.CPU.Flags.Sub, goboy.CPU.Flags.HalfCarry, goboy.CPU.Flags.Carry, fmt.Sprintf("\t%04X\n", goboy.CPU.getAF()), fmt.Sprintf("\t%04X\n", goboy.CPU.getBC()), fmt.Sprintf("\t%04X\n", goboy.CPU.getDE()), fmt.Sprintf("\t%04X\n", goboy.CPU.getHL()), fmt.Sprintf("\t%04X\n", goboy.CPU.Registers.SP), fmt.Sprintf("\t%04X\n", goboy.CPU.Registers.PC)))
		}
	}
}

func (goboy *Goboy) CheckInterrupt() int {
	if goboy.CPU.Flags.PendingInterruptEnabled {
		goboy.CPU.Flags.PendingInterruptEnabled = false
		goboy.CPU.Flags.IME = true
		return 0
	}
	if !goboy.CPU.Flags.IME && !goboy.CPU.Halted {
		return 0
	}
	req := goboy.ReadMemory(0xFF0F)
	enabled := goboy.ReadMemory(0xFFFF)
	if req > 0 {
		for i := 0; i < 5; i++ {
			if CheckBit(req, uint(i)) && CheckBit(enabled, uint(i)) {
				goboy.InterruptCPU(i)
				return 20
			}
		}
	}
	return 0
}

func (goboy *Goboy) InterruptCPU(interuptID int) {
	if goboy.CPU.Flags.PendingInterruptEnabled {
		goboy.CPU.Flags.PendingInterruptEnabled = false
		goboy.CPU.Flags.IME = true
		return
	}
	if !goboy.CPU.Flags.IME && goboy.CPU.Halted {
		goboy.CPU.Halted = false
		return
	}
	goboy.CPU.Flags.IME = false
	goboy.CPU.Halted = false
	req := goboy.ReadMemory(0xFF0F)
	req = ResetBit(req, uint(interuptID))
	goboy.WriteMemory(0xFF0F, req)
	goboy.StackPush(goboy.CPU.Registers.PC)
	switch interuptID {
	case 0:
		goboy.CPU.Registers.PC = 0x40
	case 1:
		goboy.CPU.Registers.PC = 0x48
	case 2:
		goboy.CPU.Registers.PC = 0x50
	case 3:
		goboy.CPU.Registers.PC = 0x58
	case 4:
		goboy.CPU.Registers.PC = 0x60
	default:
		log.Fatalf("Unknown Interrupt: %d", interuptID)
	}
}

func (goboy *Goboy) RequestInterrupt(id int) {
	req := goboy.ReadMemory(0xFF0F) | 0xE0
	req = SetBit(req, uint(id))
	goboy.WriteMemory(0xFF0F, req)
}

func (goboy *Goboy) UpdateTimers(cycles int) {
	goboy.DoDividerRegister(cycles)
	if goboy.IsClockEnabled() {
		goboy.Timer.TimerCounter += cycles
		freq := goboy.GetClockFreqCount()
		for goboy.Timer.TimerCounter >= freq {
			goboy.Timer.TimerCounter -= freq
			if goboy.ReadMemory(0xFF05) == 0xFF {
				goboy.WriteMemory(0xFF05, goboy.ReadMemory(0xFF06))
				goboy.RequestInterrupt(2)
			} else {
				goboy.WriteMemory(0xFF05, goboy.ReadMemory(0xFF05)+1)
			}
		}
	}
}
func (goboy *Goboy) DoDividerRegister(cycles int) {
	goboy.Timer.DividerRegister += cycles
	if goboy.Timer.DividerRegister >= 255 {
		goboy.Timer.DividerRegister -= 255
		goboy.Memory.MainMemory[0xFF04]++
	}
}

func (goboy *Goboy) SetClockFreq() {
	goboy.Timer.TimerCounter = 0
}

func (goboy *Goboy) GetClockFreqCount() int {
	switch goboy.GetClockFreq() {
	case 0:
		return 1024
	case 1:
		return 16
	case 2:
		return 64
	default:
		return 256
	}
}

func (goboy *Goboy) GetClockFreq() byte {
	return goboy.ReadMemory(0xFF07) & 0x3
}

func (goboy *Goboy) IsClockEnabled() bool {
	return CheckBit(goboy.ReadMemory(0xFF07), 2)
}

func (goboy *Goboy) initRom(romPath string) {
	RamPath := romPath + ".sav"
	romData := goboy.readRomFile(romPath)
	ramData := goboy.readRamFile(RamPath)
	goboy.Cartridge.CartridgeName = string(romData[0x134:0x143])
	if ramData == nil {
		ramData = make([]byte, 0x8000)
	}
	CartridgeType := romData[0x147]
	if _, ok := cartridgeTypeMap[CartridgeType]; !ok {
		log.Fatalf("[Cartridge] Unknown cartridge type: %x\n", CartridgeType)
	}
	switch CartridgeType {
	case 0x00, 0x08, 0x09, 0x0B, 0x0C, 0x0D:
		goboy.Cartridge.MBC = &MBCRom{
			Rom:            romData,
			CurrentROMBank: 1,
			CurrentRAMBank: 1,
		}
		goboy.Cartridge.MBCType = "rom"
		goboy.Cartridge.ROMLength = len(romData)
	case 0x01, 0x02, 0x03:
		MBC := MBC1{
			Rom:            romData,
			CurrentROMBank: 1,
			CurrentRAMBank: 0,
			RAMBank:        ramData,
		}
		goboy.Cartridge.MBC = &MBC
		goboy.Cartridge.MBCType = "MBC1"
		goboy.Cartridge.ROMLength = len(romData)
	case 0x05, 0x06:
		MBC := &MBC2{
			rom:            romData,
			CurrentROMBank: 1,
			CurrentRAMBank: 0,
			RAMBank:        ramData,
		}
		goboy.Cartridge.MBC = MBC
		goboy.Cartridge.MBCType = "MBC2"
		goboy.Cartridge.ROMLength = len(romData)
	case 0x0F, 0x10, 0x11, 0x12, 0x13:
		MBC := &MBC3{
			rom:            romData,
			CurrentROMBank: 1,
			CurrentRAMBank: 0,
			RAMBank:        ramData,
		}
		goboy.Cartridge.MBC = MBC
		goboy.Cartridge.MBCType = "MBC3"
		goboy.Cartridge.ROMLength = len(romData)
	default:
		log.Fatal("[Cartridge] Unsupported MBC type")
	}
	if _, ok := RomBankMap[romData[0x148]]; !ok {
		log.Fatalf("[Cartridge] Unknown ROM size byte : %x\n", romData[0x148])
	}
	goboy.Cartridge.ROMBank = RomBankMap[romData[0x148]]
	if _, ok := RamBankMap[romData[0x149]]; !ok {
		log.Fatalf("[Cartridge] Unknown RAM size byte : %x\n", romData[0x149])
	}
	goboy.Cartridge.RAMBank = RamBankMap[romData[0x148]]
}
