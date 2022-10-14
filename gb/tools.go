package gb

import "github.com/sqweek/dialog"

func SetBit(n byte, pos uint) byte {
	n |= (1 << pos)
	return n
}

func ResetBit(n byte, pos uint) byte {
	n = n &^ (1 << pos)
	return n
}

func CheckBit(n byte, pos uint) bool {
	return (((n) & (1 << (pos))) > 0)
}

func GetBit(n byte, pos uint) byte {
	return (n >> pos) & 1
}

func LoadRomDialog() (string, error) {
	var err error
	var cartridgePath string
	cartridgePath, err = dialog.File().
		Filter("GameBoy ROM", "zip", "gb", "gbc", "bin").
		Title("Choose a ROM to load").Load()
	if err != nil {
		return "", err
	}
	return cartridgePath, nil
}
