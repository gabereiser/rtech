package types

type RColor struct {
	Red   uint8
	Green uint8
	Blue  uint8
	Alpha uint8
}

func (c *RColor) RedF() float32 {
	return float32(c.Red) / 255.0
}
func (c *RColor) GreenF() float32 {
	return float32(c.Red) / 255.0
}
func (c *RColor) BlueF() float32 {
	return float32(c.Red) / 255.0
}
func (c *RColor) AlphaF() float32 {
	return float32(c.Red) / 255.0
}
func (c *RColor) ToArrayF32() []float32 {
	return []float32{c.RedF(), c.GreenF(), c.BlueF(), c.AlphaF()}
}
func NewColorFromInt32(color int32) RColor {
	alpha := color & 0xFF
	blue := (color >> 8) & 0xFF
	green := (color >> 16) & 0xFF
	red := (color >> 24) & 0xFF
	return RColor{
		Red:   uint8(red),
		Green: uint8(green),
		Blue:  uint8(blue),
		Alpha: uint8(alpha),
	}
}
func (c *RColor) ToInt32() int32 {
	r := int32(c.Red) & 0xFF
	g := int32(c.Green) & 0xFF
	b := int32(c.Blue) & 0xFF
	a := int32(c.Alpha) & 0xFF
	return (r << 24) + (g << 16) + (b << 8) + (a)
}
