package screen

import (
	"fmt"
	_ "image/png"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type DebugData struct {
	CPU    CPU
	FF     func() byte
	Joypad *uint8
}

type Control interface {
	GetUp() bool
	GetDown() bool
	GetLeft() bool
	GetRigth() bool
	GetA() bool
	GetB() bool
	GetStart() bool
	GetSelect() bool
	GetReload() bool
}

type CPU interface {
	GetAF() uint16
	GetBC() uint16
	GetDE() uint16
	GetHL() uint16
	GetRegisterSP() uint16
	GetRegisterPC() uint16
	GetHalted() bool
}

type ControlWindow struct {
	EventTouche [9]bool
}

func (c *ControlWindow) GetUp() bool {
	return c.EventTouche[0]
}

func (c *ControlWindow) GetDown() bool {
	return c.EventTouche[1]
}

func (c *ControlWindow) GetLeft() bool {
	return c.EventTouche[2]
}

func (c *ControlWindow) GetRigth() bool {
	return c.EventTouche[3]
}

func (c *ControlWindow) GetA() bool {
	return c.EventTouche[4]
}

func (c *ControlWindow) GetB() bool {
	return c.EventTouche[5]
}

func (c *ControlWindow) GetStart() bool {
	return c.EventTouche[6]
}

func (c *ControlWindow) GetSelect() bool {
	return c.EventTouche[7]
}

func (c *ControlWindow) GetReload() bool {
	return c.EventTouche[8]
}

// var s *Screen

type Screen struct {
	Title         string
	Width, Height int
	Screen        *[160][144][3]uint8
	// EventTouche is Table of touche  pressed , with 0 = up , 1 = down , 2 = left , 3 = right , 4 = A , 5 = B, 6 = Start , 7 = Selecte
	Control       ControlWindow
	Debug         bool
	DebuggingData DebugData
	Pause         *bool
}

func (s *Screen) setPixel() []uint8 {
	temp := []uint8{}
	for y := len(s.Screen[0]) - 1; y > -1; y-- {
		for x := range s.Screen {
			colorEnd := s.Screen[x][y]
			temp = append(temp, colorEnd[0], colorEnd[1], colorEnd[2], 0xff)
		}
	}
	return temp
}

func (s *Screen) GenerateIMG() []uint8 {
	return s.setPixel()
}

func (s *Screen) Loop() {
	pixelgl.Run(func() {
		cfg := pixelgl.WindowConfig{
			Title:     s.Title,
			Bounds:    pixel.R(0, 0, 640, 144*4),
			VSync:     true,
			Resizable: true,
		}
		win, err := pixelgl.NewWindow(cfg)
		if err != nil {
			panic(err)
		}

		// canvas := win.Canvas()
		canvas := pixelgl.NewCanvas(pixel.R(0, 0, 160, 144))
		//Debuging
		if s.Debug {
			for !win.Closed() {
				s.Control.EventTouche = [9]bool{win.Pressed(pixelgl.KeyUp), win.Pressed(pixelgl.KeyDown), win.Pressed(pixelgl.KeyLeft), win.Pressed(pixelgl.KeyRight), win.Pressed(pixelgl.KeyQ), win.Pressed(pixelgl.KeyW), win.Pressed(pixelgl.KeyEnter), win.Pressed(pixelgl.KeySpace), win.Pressed(pixelgl.KeyR)}
				if win.Bounds().W() == 0 && win.Bounds().H() == 0 {
					win.Clear(colornames.Black)
					*s.Pause = true
				} else {
					*s.Pause = false
					canvas.SetPixels(s.GenerateIMG())
				}
				win.Clear(colornames.Black)
				win.SetMatrix(pixel.IM.Scaled(pixel.ZV,
					math.Min(
						win.Bounds().W()/canvas.Bounds().W(),
						win.Bounds().H()/canvas.Bounds().H(),
					),
				).Moved(win.Bounds().Center()))
				canvas.Draw(win, pixel.IM.Moved(canvas.Bounds().Min))
				win.Update()
				fmt.Printf("AF : %04X\n", s.DebuggingData.CPU.GetAF())
				fmt.Printf("BC : %04X\n", s.DebuggingData.CPU.GetBC())
				fmt.Printf("DE : %04X\n", s.DebuggingData.CPU.GetDE())
				fmt.Printf("HL : %04X\n", s.DebuggingData.CPU.GetHL())
				fmt.Printf("SP : %04X\n", s.DebuggingData.CPU.GetRegisterSP())
				fmt.Printf("PC : %04X\n", s.DebuggingData.CPU.GetRegisterPC())
				fmt.Printf("JP : %08b\n", s.DebuggingData.FF())
				fmt.Printf("INPUTS : %08b\n", *s.DebuggingData.Joypad)
				fmt.Printf("HALTED : %t\n", s.DebuggingData.CPU.GetHalted())
			}
		} else {
			for !win.Closed() {
				s.Control.EventTouche = [9]bool{win.Pressed(pixelgl.KeyUp), win.Pressed(pixelgl.KeyDown), win.Pressed(pixelgl.KeyLeft), win.Pressed(pixelgl.KeyRight), win.Pressed(pixelgl.KeyQ), win.Pressed(pixelgl.KeyW), win.Pressed(pixelgl.KeyEnter), win.Pressed(pixelgl.KeySpace), win.Pressed(pixelgl.KeyR)}
				if win.Bounds().W() == 0 && win.Bounds().H() == 0 {
					win.Clear(colornames.Black)
					*s.Pause = true
				} else {
					*s.Pause = false
					canvas.SetPixels(s.GenerateIMG())
				}
				win.Clear(colornames.Black)
				win.SetMatrix(pixel.IM.Scaled(pixel.ZV,
					math.Min(
						win.Bounds().W()/canvas.Bounds().W(),
						win.Bounds().H()/canvas.Bounds().H(),
					),
				).Moved(win.Bounds().Center()))
				canvas.Draw(win, pixel.IM.Moved(canvas.Bounds().Min))
				win.Update()

			}
		}
	})
}

func CreateScreen(Title string, screen *[160][144][3]uint8, pause *bool) *Screen {
	return &Screen{Title: Title, Screen: screen, Control: ControlWindow{EventTouche: [9]bool{}}, Pause: pause}
}

func RunScreen(s *Screen, Debug bool) {
	s.Debug = Debug
	s.Loop()
}
