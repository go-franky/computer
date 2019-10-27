package computer

import "fmt"

// Size represent a unit for computer storage
type Size int64

// Common Size for computer hardware
const (
	Byte Size = 1

	Kilobyte = Byte * 1000     // KB
	Megabyte = Kilobyte * 1000 // MB
	Gigabyte = Megabyte * 1000 // GB
	Terabyte = Gigabyte * 1000 // TB
	Petabyte = Terabyte * 1000 // PB
	Exabyte  = Petabyte * 1000 // EB
	// Zettabyte = Exabyte * 1000   // ZB
	// Yottabyte = Zettabyte * 1000 // YB

	Kibibyte = Byte * 1024     // KiB
	Mebibyte = Kibibyte * 1024 // MiB
	Gibibyte = Mebibyte * 1024 // GiB
	Tebibyte = Gibibyte * 1024 // TiB
	Pebibyte = Tebibyte * 1024 // PiB
	Exbibyte = Pebibyte * 1024 // EiB
	// Zebibyte = Exbibyte * 1024
	// Yobibyte = Zebibyte * 1024
)

func (s Size) Bytes() Size {
	return s
}

func (s Size) String() string {
	val := s
	str := ""
	if val >= Exbibyte {
		str = fmt.Sprintf("%s%d%s", str, val/Exbibyte, "EiB")
		val = val % Exbibyte
	}
	if val >= Pebibyte {
		str = fmt.Sprintf("%s%d%s", str, val/Pebibyte, "PiB")
		val = val % Pebibyte
	}
	if val >= Gibibyte {
		str = fmt.Sprintf("%s%d%s", str, val/Gibibyte, "GiB")
		val = val % Gibibyte
	}
	if val >= Mebibyte {
		str = fmt.Sprintf("%s%d%s", str, val/Mebibyte, "MiB")
		val = val % Mebibyte
	}
	if val >= Kibibyte {
		str = fmt.Sprintf("%s%d%s", str, val/Kibibyte, "KiB")
		val = val % Kibibyte
	}
	if val > 0 {
		str = fmt.Sprintf("%s%d%s", str, val, "B")
	}

	return str
}

// Kibibytes returns the size in KiB.
func (s Size) Kibibytes() float64 {
	return float64(int64(s)) / float64(Kibibyte)
}

// Mebibytes returns the size in KiB.
func (s Size) Mebibytes() float64 {
	return float64(int64(s)) / float64(Mebibyte)
}
