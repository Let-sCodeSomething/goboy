package gb

import (
	"log"
)

type PPU struct {
}

func (goboy *Goboy) initPPU() {

}

func (goboy *Goboy) SetLCDStatus() {
	status := goboy.ReadMemory(0xFF41)
	if !goboy.IsLCDEnabled() {
		// set the mode to 1 during lcd disabled and reset scanline
		goboy.Timer.ScanlineCounter = 456
		goboy.Memory.MainMemory[0xFF44] = 0
		status &= 252
		status = ResetBit(status, 0)
		status = ResetBit(status, 1)
		goboy.WriteMemory(0xFF41, status)
		return
	}

	currentLine := goboy.ReadMemory(0xFF44)
	currentMode := status & 0x3
	mode := byte(0)
	reqInt := false

	// in vblank so set mode to 1
	if currentLine >= 144 {
		mode = 1
		status = SetBit(status, 0)
		status = ResetBit(status, 1)
		reqInt = CheckBit(status, 4)
	} else {
		mode2bounds := 456 - 80
		mode3bounds := mode2bounds - 172

		// mode 2
		if goboy.Timer.ScanlineCounter >= mode2bounds {
			mode = 2
			status = SetBit(status, 1)
			status = ResetBit(status, 0)
			reqInt = CheckBit(status, 5)
		} else if goboy.Timer.ScanlineCounter >= mode3bounds {
			mode = 3
			// mode 3
			status = SetBit(status, 1)
			status = SetBit(status, 0)
		} else {
			// mode 0
			mode = 0
			status = ResetBit(status, 1)
			status = ResetBit(status, 0)
			reqInt = CheckBit(status, 3)
		}
	}

	// just entered a new mode so request interupt
	if reqInt && (mode != currentMode) {
		goboy.RequestInterrupt(1)
	}

	// check the conincidence flag
	if currentLine == goboy.ReadMemory(0xFF45) {
		status = SetBit(status, 2)
		if CheckBit(status, 6) {
			goboy.RequestInterrupt(1)
		}
	} else {
		status = ResetBit(status, 2)
	}
	goboy.WriteMemory(0xFF41, status)

}

/*
	Check if LCD is enabled
*/
func (goboy *Goboy) IsLCDEnabled() bool {
	return CheckBit(goboy.ReadMemory(0xFF40), 7)
}

/*
	Scan line and draw line
*/
func (goboy *Goboy) UpdateGraphics(cycles int) {
	//Set the LCD status register
	goboy.SetLCDStatus()

	//A complete cycle through Scan Line states takes 456 clks.
	//We use a counter to mark this.
	if goboy.IsLCDEnabled() {
		goboy.Timer.ScanlineCounter -= cycles
	} else {
		return
	}

	if goboy.Timer.ScanlineCounter <= 0 {
		// time to move onto next scanline
		goboy.Memory.MainMemory[0xFF44]++

		currentLine := goboy.ReadMemory(0xFF44)

		//Reset the counter
		goboy.Timer.ScanlineCounter += 456

		// we have entered vertical blank period
		if currentLine == 144 {
			goboy.DrawScanLine()
			goboy.RequestInterrupt(0)
		} else if currentLine > 153 {
			// if gone past scanline 153 reset to 0
			goboy.Memory.MainMemory[0xFF44] = 0
			goboy.DrawScanLine()
		} else if currentLine < 144 {
			//log.Println(currentLine)
			goboy.DrawScanLine()
		}

	}

}

func (goboy *Goboy) DrawScanLine() {
	control := goboy.ReadMemory(0xFF40)
	if CheckBit(control, 0) {
		goboy.RenderTiles()
	}
	if CheckBit(control, 1) {
		goboy.RenderSprites()
	}
}

// Render Sprites is not working
func (goboy *Goboy) RenderSprites() {
	use8x16 := false
	lcdControl := goboy.ReadMemory(0xFF40)

	if CheckBit(lcdControl, 2) {
		use8x16 = true
	}
	for sprite := 0; sprite < 40; sprite++ {
		// sprite occupies 4 bytes in the sprite attributes table
		index := sprite * 4
		yPos := goboy.ReadMemory(0xFE00+uint16(index)) - 16
		xPos := goboy.ReadMemory(0xFE00+uint16(index)+1) - 8
		tileLocation := goboy.ReadMemory(uint16(0xFE00 + index + 2))
		attributes := goboy.ReadMemory(0xFE00 + uint16(index) + 3)

		yFlip := CheckBit(attributes, 6)
		xFlip := CheckBit(attributes, 5)
		priority := !CheckBit(attributes, 7)
		scanline := goboy.ReadMemory(0xFF44)

		ysize := 8
		if use8x16 {
			ysize = 16
		}

		// does this sprite intercept with the scanline?
		if (scanline >= yPos) && (scanline < (yPos + byte(ysize))) {
			line := int(scanline - yPos)
			// read the sprite in backwards in the y axis
			if yFlip {
				line -= ysize
				line *= -1
			}
			line *= 2 // same as for tiles
			dataAddress := (uint16(int(tileLocation)*16 + line))
			data1 := goboy.ReadMemory(0x8000 + dataAddress)
			data2 := goboy.ReadMemory(0x8000 + dataAddress + 1)

			// its easier to read in from right to left as pixel 0 is
			// bit 7 in the colour data, pixel 1 is bit 6 etc...
			for tilePixel := 7; tilePixel >= 0; tilePixel-- {
				colourbit := tilePixel

				// read the sprite in backwards for the x axis
				if xFlip {
					colourbit -= 7
					colourbit *= -1
				}

				colourNum := GetBit(data2, uint(colourbit))
				colourNum <<= 1
				colourNum |= GetBit(data1, uint(colourbit))
				colourAddress := uint16(0xFF48)
				if CheckBit(attributes, 4) {
					colourAddress = 0xFF49
				}
				// now we have the colour id get the actual
				// colour from palette 0xFF47
				colour := goboy.GetColour(colourNum, colourAddress)

				// white is transparent for sprites.
				if colourNum == 0 {
					continue
				}

				red := uint8(0)
				green := uint8(0)
				blue := uint8(0)

				switch colour {
				case 0:
					red = 255
					green = 255
					blue = 255
				case 1:
					red = 0xCC
					green = 0xCC
					blue = 0xCC
				case 2:
					red = 0x77
					green = 0x77
					blue = 0x77
				default:
					red = 0
					green = 0
					blue = 0
				}

				xPix := 0 - tilePixel
				xPix += 7

				pixel := int(xPos) + xPix

				// sanity check
				if (scanline < 0) || (scanline > 143) || (pixel < 0) || (pixel > 159) {
					continue
				}

				if goboy.ScanLineBG[pixel] || priority {
					goboy.Screen[pixel][scanline][0] = red
					goboy.Screen[pixel][scanline][1] = green
					goboy.Screen[pixel][scanline][2] = blue
				}

			}
		}

	}
}

var last uint16

/*
	Render Tiles for the current scan line
*/
func (goboy *Goboy) RenderTiles() {
	var tileData uint16 = 0
	var backgroundMemory uint16 = 0
	var unsig bool = true
	lcdControl := goboy.ReadMemory(0xFF40)

	//	FF42 - SCY - Scroll Y (R/W)
	//	FF43 - SCX - Scroll X (R/W)
	//		Specifies the position in the 256x256 pixels BG map (32x32 tiles)
	//		which is to be displayed at the upper/left LCD display position.
	//		Values in range from 0-255 may be used for X/Y each, the video
	//		controller automatically wraps back to the upper (left) position
	//		in BG map when drawing exceeds the lower (right) border of the BG
	//		map area.
	scrollY := goboy.ReadMemory(0xFF42)
	scrollX := goboy.ReadMemory(0xFF43)

	//	FF4A - WY - Window Y Position (R/W)
	//	FF4B - WX - Window X Position minus 7 (R/W)
	//		Specifies the upper/left positions of the Window area.
	//		(The window is an alternate background area which can be
	//		displayed above of the normal background. OBJs (sprites)
	//		may be still displayed above or behinf the window, just as
	//		for normal BG.)
	//
	//		The window becomes visible (if enabled) when positions are set
	//		in range WX=0..166, WY=0..143. A postion of WX=7, WY=0 locates
	//		the window at upper left, it is then completly covering normal
	//		background.
	windowY := goboy.ReadMemory(0xFF4A)
	windowX := goboy.ReadMemory(0xFF4B) - 7

	usingWindow := false

	// is the window enabled?
	if CheckBit(lcdControl, 5) {

		// is the current scan line we're drawing
		// within the windows Y pos?,
		if windowY <= goboy.ReadMemory(0xFF44) {
			usingWindow = true
		}
	}

	// which tile data are we using?
	if CheckBit(lcdControl, 4) {
		tileData = 0x8000
	} else {
		// IMPORTANT: This memory region uses signed
		// bytes as tile identifiers
		tileData = 0x8800
		unsig = false
	}

	// which background mem?
	if !usingWindow {
		if CheckBit(lcdControl, 3) {
			backgroundMemory = 0x9C00
		} else {
			backgroundMemory = 0x9800
		}
	} else {
		if CheckBit(lcdControl, 6) {
			backgroundMemory = 0x9C00
		} else {
			backgroundMemory = 0x9800
		}
	}

	// yPos is used to calculate which of 32 vertical tiles the
	// current scanline is drawing
	var yPos byte = 0
	if !usingWindow {
		yPos = scrollY + goboy.ReadMemory(0xFF44)
	} else {
		yPos = goboy.ReadMemory(0xFF44) - windowY
	}

	// which of the 8 vertical pixels of the current
	// tile is the scanline on?
	var tileRow = ((uint16(yPos / 8)) * 32)

	// time to start drawing the 160 horizontal pixels
	// for this scanline
	for pixel := byte(0); pixel < 160; pixel++ {

		xPos := byte(pixel) + scrollX

		// translate the current x pos to window space if necessary
		if usingWindow {
			if pixel >= windowX {
				xPos = pixel - windowX
			}
		}

		// which of the 32 horizontal tiles does this xPos fall within?
		tileCol := uint16(xPos / 8)
		var tileNum int16

		// get the tile identity number. Remember it can be signed
		// or unsigned
		tileAddress := backgroundMemory + tileRow + tileCol
		if unsig {
			tileNum = int16(goboy.ReadMemory(tileAddress))
		} else {
			tileNum = int16(int8(goboy.ReadMemory(tileAddress)))
		}

		// deduce where this tile identifier is in memory.
		tileLocation := tileData
		if unsig {
			tileLocation += uint16(tileNum * 16)
		} else {
			tileLocation = uint16(int32(tileLocation) + int32((tileNum+128)*16))
		}

		// find the correct vertical line we're on of the
		// tile to get the tile data
		//	from in memory
		line := yPos % 8
		// each vertical line takes up two bytes of memory
		line *= 2
		data1 := goboy.ReadMemory(tileLocation + uint16(line))
		data2 := goboy.ReadMemory(tileLocation + uint16(line) + 1)
		if last == 0x86F0 && (tileLocation+uint16(line) == 0x8000) {
			log.Printf("%X\n", tileNum)

		}
		last = tileLocation + uint16(line)

		// pixel 0 in the tile is it 7 of data 1 and data2.
		// Pixel 1 is bit 6 etc..
		var colourBit int = int(xPos % 8)
		colourBit -= 7
		colourBit *= -1

		// combine data 2 and data 1 to get the colour id for this pixel
		// in the tile
		colourNum := GetBit(data2, uint(colourBit))
		colourNum <<= 1
		colourNum |= GetBit(data1, uint(colourBit))

		// now we have the colour id get the actual
		// colour from palette 0xFF47
		colour := goboy.GetColour(colourNum, 0xFF47)

		red := uint8(0)
		green := uint8(0)
		blue := uint8(0)

		switch colour {
		case 0:
			red = 255
			green = 255
			blue = 255
		case 1:
			red = 0xCC
			green = 0xCC
			blue = 0xCC
		case 2:
			red = 0x77
			green = 0x77
			blue = 0x77
		default:
			red = 0
			green = 0
			blue = 0
		}
		finally := int(goboy.ReadMemory(0xFF44))
		// safety check to make sure what im about
		// to set is int the 160x144 bounds
		if (finally < 0) || (finally > 143) || (pixel < 0) || (pixel > 159) {
			continue
		}

		// Store whether the background is white
		if colour == 0 {
			goboy.ScanLineBG[pixel] = true
		} else {
			goboy.ScanLineBG[pixel] = false
		}

		goboy.Screen[pixel][finally][0] = red
		goboy.Screen[pixel][finally][1] = green
		goboy.Screen[pixel][finally][2] = blue
	}

}

/*
	Get colour id via colour palette and colour Num
		0 - WHITE
		1 - LIGHT_GRAY
		2 - DARK_GRAY
		3 - BLACK
	TODO: GCB mode
*/
func (goboy *Goboy) GetColour(colourNum byte, address uint16) int {
	res := 0
	palette := goboy.ReadMemory(address)

	// which bits of the colour palette does the colour id map to?
	var hi uint
	var lo uint
	switch colourNum {
	case 0:
		hi = 1
		lo = 0
	case 1:
		hi = 3
		lo = 2
	case 2:
		hi = 5
		lo = 4
	case 3:
		hi = 7
		lo = 6
	default:
		hi = 1
		lo = 0
	}

	// use the palette to get the colour
	colour := byte(0)
	colour = GetBit(palette, hi) << 1
	colour |= GetBit(palette, lo)

	switch colour {
	case 0:
		res = 0
	case 1:
		res = 1
	case 2:
		res = 2
	case 3:
		res = 3
	default:
		res = 0
	}

	return res
}
