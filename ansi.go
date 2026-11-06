package rotato

type ColorFunc func(s string) string

type Colorizer struct {
	spinner     ColorFunc
	message     ColorFunc
	prefixMesg  ColorFunc
	delimiter   ColorFunc
	doneMessage ColorFunc
	doneSymbol  ColorFunc
	failMessage ColorFunc
	failSymbol  ColorFunc
}

func newColorizer() *Colorizer { return &Colorizer{} }

func (c *Colorizer) Spinner(s string) string    { return c.format(c.spinner, s) }     // Spinner color
func (c *Colorizer) Message(s string) string    { return c.format(c.message, s) }     // Spinner message color
func (c *Colorizer) Prefix(s string) string     { return c.format(c.prefixMesg, s) }  // Prefix message color
func (c *Colorizer) Delimiter(s string) string  { return c.format(c.delimiter, s) }   // Delimiter color
func (c *Colorizer) DoneMsg(s string) string    { return c.format(c.doneMessage, s) } // Done channel message color
func (c *Colorizer) DoneSymbol(s string) string { return c.format(c.doneSymbol, s) }  // Done symbol color
func (c *Colorizer) FailMsg(s string) string    { return c.format(c.failMessage, s) } // Fail message color
func (c *Colorizer) FailSymbol(s string) string { return c.format(c.failSymbol, s) }  // Fail symbol color

func (c *Colorizer) format(color ColorFunc, text string) string {
	if text == "" || color == nil {
		return text
	}
	return color(text)
}
