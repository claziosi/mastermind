package constants

type Colour int64

const (
	NbColumns   int    = 10
	NbLines     int    = 5
	NbColors    int    = 9
	Green       Colour = 0
	Yellow             = 1
	Red                = 2
	Blue               = 3
	Transparent        = 4
	Brown              = 5
	Purple             = 6
	Gray               = 7
	White              = 8
)

func (e Colour) String() string {
	switch e {
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Transparent:
		return "Transparent"
	case Brown:
		return "Brown"
	case Purple:
		return "Purple"
	case Gray:
		return "Gray"
	case White:
		return "White"
	default:
		return "Transparent"
	}
}
