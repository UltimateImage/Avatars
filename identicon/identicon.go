package identicon

type imgFormat int

const (
	PNG imgFormat = iota
	SVG
)

type Option struct {
	foreGround [4]byte
	backGround [4]byte
	size       int
	margin     float32
	saturation float32
	brightness float32
	format     imgFormat
}

type Identicon struct {
	hash   string
	option *Option
}

func NewIdenticon(hash string, option *Option) *Identicon {
	if len(hash) < 15 {
		panic("hash must has at least 15 characters")
	}
	return &Identicon{
		hash: hash,
	}
}

func (i *Identicon) hslToRgb() [4]byte {
	return [4]byte{}
}
