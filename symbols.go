package rotato

var defaultSymbols = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

// WithSymbols returns an option function that sets the spinner unicode
// animation.
func WithSymbols(symbols ...string) Option {
	return func(r *Rotato) {
		r.symbols = symbols
	}
}

// WithSymbolsBlock returns an option function that sets the spinner unicode
// animation with blocks.
//
//	"░", "▒", "▒", "░", "▓".
func WithSymbolsBlock() Option {
	return func(r *Rotato) {
		r.symbols = []string{"░", "▒", "▒", "░", "▓"}
	}
}

// WithSymbolsBarBlock returns an option function that sets the spinner
// unicode animation with bars.
//
//	"█▒▒▒▒▒▒▒▒▒", "███▒▒▒▒▒▒▒", "█████▒▒▒▒▒", "███████▒▒▒", "██████████".
func WithSymbolsBarBlock() Option {
	return func(r *Rotato) {
		r.symbols = []string{"█▒▒▒▒▒▒▒▒▒", "███▒▒▒▒▒▒▒", "█████▒▒▒▒▒", "███████▒▒▒", "██████████"}
	}
}

// WithSymbolsBarBlock2 returns an option function that sets the spinner
// unicode animation with bars.
//
//	"[|       ]", "[||      ]", "[|||     ]", "[||||    ]", "[|||||   ]", "[||||||  ]", "[||||||| ]", "[||||||||]".
func WithSymbolsBarBlock2() Option {
	return func(r *Rotato) {
		r.symbols = []string{
			"[       ]",
			"[|      ]",
			"[||     ]",
			"[|||    ]",
			"[||||   ]",
			"[|||||  ]",
			"[|||||| ]",
			"[|||||||]",
		}
	}
}

// WithSymbolsBarBlock3 returns an option function that sets the spinner
// unicode animation with bars.
//
//	"[=       ]", "[==      ]", "[===     ]", "[====    ]", "[=====   ]", "[======  ]", "[======= ]", "[========]".
func WithSymbolsBarBlock3() Option {
	return func(r *Rotato) {
		r.symbols = []string{
			"[       ]",
			"[=      ]",
			"[==     ]",
			"[===    ]",
			"[====   ]",
			"[=====  ]",
			"[====== ]",
			"[=======]",
		}
	}
}

// WithSymbolsBarBlock4 returns an option function that sets the spinner
// unicode animation with bars.
//
//	"|", "||", "|||", "||||", "|||||", "||||||", "|||||||", "||||||||",
//	"|||||||", "||||||", "|||||", "||||", "|||", "||", "|".
func WithSymbolsBarBlock4() Option {
	return func(r *Rotato) {
		r.symbols = []string{
			"|",
			"||",
			"|||",
			"||||",
			"|||||",
			"||||||",
			"|||||||",
			"||||||||",
			"|||||||",
			"||||||",
			"|||||",
			"||||",
			"|||",
			"||",
			"|",
		}
	}
}

// WithSymbolsBarBlock5 returns an option function that sets the spinner
// unicode animation with bars.
//
//	"[*-------]", "[-*------]", "[--*-----]", "[---*----]", "[----*---]".
//	"[-----*--]", "[------*-]", "[-------*]", "[------*-]", "[-----*--]".
//	"[----*---]", "[---*----]", "[--*-----]", "[-*------]", "[*-------]".
func WithSymbolsBarBlock5() Option {
	return func(r *Rotato) {
		r.symbols = []string{
			"[*-------]",
			"[-*------]",
			"[--*-----]",
			"[---*----]",
			"[----*---]",
			"[-----*--]",
			"[------*-]",
			"[-------*]",
			"[------*-]",
			"[-----*--]",
			"[----*---]",
			"[---*----]",
			"[--*-----]",
			"[-*------]",
			"[*-------]",
		}
	}
}

// WithSymbolsBarBlock6 returns an option function that sets the spinner
// unicode animation with bars.
//
//	"·-----", "-·----", "--·---", "---·--", "----·-", "-----·",
//	"----·-", "---·--", "--·---", "-·----", "·-----",
func WithSymbolsBarBlock6() Option {
	return func(r *Rotato) {
		r.symbols = []string{
			"·-----",
			"-·----",
			"--·---",
			"---·--",
			"----·-",
			"-----·",
			"----·-",
			"---·--",
			"--·---",
			"-·----",
			"·-----",
		}
	}
}

// WithSymbolsBlockPretty returns an option function that sets the spinner
// unicode animation with pretty blocks.
//
//	"", "", "", "", "", "", "".
func WithSymbolsBlockPretty() Option {
	return func(r *Rotato) {
		//  
		r.symbols = []string{"", "", "", "", "", "", ""}
	}
}

// WithSymbolsDots returns an option function that sets the spinner unicode
// animation with braille patterns.
//
//	"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷".
func WithSymbolsDots() Option {
	return func(r *Rotato) {
		r.symbols = []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
	}
}

// WithSymbolsDots2 returns an option function that sets the spinner unicode
// animation with dots.
//
//	"  . . . .", ".   . . .", ". .   . .", ". . .   .", ". . . .  ", ". . . . .".
func WithSymbolsDots2() Option {
	return func(r *Rotato) {
		r.symbols = []string{
			"  . . . .",
			".   . . .",
			". .   . .",
			". . .   .",
			". . . .  ",
			". . . . .",
		}
	}
}

// WithSymbolsDots3 returns an option function that sets the spinner unicode
// animation with dots.
//
//	"⠄", "⠆", "⠇", "⠋", "⠙", "⠸", "⠰", "⠠", "⠰", "⠸", "⠙", "⠋", "⠇", "⠆".
func WithSymbolsDots3() Option {
	return func(r *Rotato) {
		r.symbols = []string{"⠄", "⠆", "⠇", "⠋", "⠙", "⠸", "⠰", "⠠", "⠰", "⠸", "⠙", "⠋", "⠇", "⠆"}
	}
}

// WithSymbolsDots4 returns an option function that sets the spinner unicode
// animation with dots.
//
//	"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈".
func WithSymbolsDots4() Option {
	return func(r *Rotato) {
		r.symbols = []string{"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"}
	}
}

// WithSymbolsDots5 returns an option function that sets the spinner unicode
// animation with dots.
//
//	"⠁", "⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒".
func WithSymbolsDots5() Option {
	return func(r *Rotato) {
		r.symbols = []string{
			"⠁",
			"⠁",
			"⠉",
			"⠙",
			"⠚",
			"⠒",
			"⠂",
			"⠂",
			"⠒",
			"⠲",
			"⠴",
			"⠤",
			"⠄",
			"⠄",
			"⠤",
			"⠠",
			"⠠",
			"⠤",
			"⠦",
			"⠖",
			"⠒",
			"⠐",
			"⠐",
			"⠒",
			"⠓",
			"⠋",
			"⠉",
			"⠈",
			"⠈",
		}
	}
}

// WithSymbolsLines returns an option function that sets the spinner unicode
// animation with lines.
//
//	"⠂", "-", "–", "—", "–", "-".
func WithSymbolsLines() Option {
	return func(r *Rotato) {
		r.symbols = []string{"⠂", "-", "–", "—", "–", "-"}
	}
}

// WithSymbolsWave returns an option function that sets the spinner unicode
// animation with wave patterns.
//
//	"⢄", "⢂", "⢁", "⡀", "⠈", "⠘", "⠸".
func WithSymbolsWave() Option {
	return func(r *Rotato) {
		r.symbols = []string{"⢄", "⢂", "⢁", "⡀", "⠈", "⠘", "⠸"}
	}
}

// WithSymbolsGrow returns an option function that sets the spinner unicode
// animation with growing bars.
//
//	"▉", "▊", "▋", "▌", "▍", "▎", "▏".
func WithSymbolsGrow() Option {
	return func(r *Rotato) {
		r.symbols = []string{"▉", "▊", "▋", "▌", "▍", "▎", "▏"}
	}
}

// WithSymbolsGrowVert returns an option function that sets the spinner unicode
// animation with growing bars.
//
//	"▁", "▃", "▄", "▅", "▆", "▇", "▆", "▅", "▄", "▃".
func WithSymbolsGrowVert() Option {
	return func(r *Rotato) {
		r.symbols = []string{"▁", "▃", "▄", "▅", "▆", "▇", "▆", "▅", "▄", "▃"}
	}
}

// WithSymbolsMoon returns an option function that sets the spinner unicode
// animation with moon phases.
//
//	"🌑", "🌒", "🌓", "🌔", "🌕", "🌖", "🌗", "🌘".
func WithSymbolsMoon() Option {
	return func(r *Rotato) {
		r.symbols = []string{"🌑", "🌒", "🌓", "🌔", "🌕", "🌖", "🌗", "🌘"}
	}
}

// WithSymbolsPipe returns an option function that sets the spinner unicode
// animation with pipe characters.
//
//	"|", "/", "-", "\\".
func WithSymbolsPipe() Option {
	return func(r *Rotato) {
		r.symbols = []string{"|", "/", "-", "\\"}
	}
}

// WithSymbolsPipe2 returns an option function that sets the spinner unicode
// animation with pipe characters.
//
//	"┤", "┘", "┴", "└", "├", "┌", "┬", "┐".
func WithSymbolsPipe2() Option {
	return func(r *Rotato) {
		r.symbols = []string{"┤", "┘", "┴", "└", "├", "┌", "┬", "┐"}
	}
}

// WithSymbolsSquare returns an option function that sets the spinner unicode
// animation with square segments.
//
//	"▖", "▘", "▝", "▗".
func WithSymbolsSquare() Option {
	return func(r *Rotato) {
		r.symbols = []string{"▖", "▘", "▝", "▗"}
	}
}

// WithSymbolsSquare2 returns an option function that sets the spinner unicode
// animation with square segments.
//
//	"", "", "", "", "".
func WithSymbolsSquare2() Option {
	return func(r *Rotato) {
		r.symbols = []string{"", "", "", "", ""}
	}
}

// WithSymbolsClock returns an option function that sets the spinner unicode
// animation with clock symbols.
//
//	"🕛", "🕐", "🕑", "🕒", "🕓", "🕔", "🕕", "🕖", "🕗", "🕘", "🕙", "🕚".
func WithSymbolsClock() Option {
	return func(r *Rotato) {
		r.symbols = []string{"🕛", "🕐", "🕑", "🕒", "🕓", "🕔", "🕕", "🕖", "🕗", "🕘", "🕙", "🕚"}
	}
}

// WithSymbolsDiamond returns an option function that sets the spinner unicode
// animation with diamond symbols.
//
//	"◇", "◈", "⬟", "⬞".
func WithSymbolsDiamond() Option {
	return func(r *Rotato) {
		r.symbols = []string{"◇", "◈", "⬟", "⬞"}
	}
}

// WithSymbolsDiamond2 returns an option function that sets the spinner unicode
// animation with diamond symbols.
//
//	"", "", "", "".
func WithSymbolsDiamond2() Option {
	return func(r *Rotato) {
		r.symbols = []string{"", "", "", ""}
	}
}

// WithSymbolsPlusCross returns an option function that sets the spinner unicode
// animation with plus and cross symbols.
//
//	"+", "x".
func WithSymbolsPlusCross() Option {
	return func(r *Rotato) {
		r.symbols = []string{"+", "x"}
	}
}

// WithSymbolsArrows returns an option function that sets the spinner unicode
// animation with arrows.
//
//	"<", "<<", "<<<", "-", ">", ">>", ">>>".
func WithSymbolsArrows() Option {
	return func(r *Rotato) {
		r.symbols = []string{"<", "<<", "<<<", "-", ">", ">>", ">>>"}
	}
}

// WithSymbolsArrows2 returns an option function that sets the spinner unicode
// animation with arrows.
//
//	">   ", ">>  ", ">>> ", ">>>>".
func WithSymbolsArrows2() Option {
	return func(r *Rotato) {
		r.symbols = []string{">   ", ">>  ", ">>> ", ">>>>"}
	}
}

// WithSymbolsArrows3 returns an option function that sets the spinner unicode
// animation with arrows.
//
//	"▹▹▹▹▹", "▸▹▹▹▹", "▹▸▹▹▹", "▹▹▸▹▹", "▹▹▹▸▹", "▹▹▹▹▸".
func WithSymbolsArrows3() Option {
	return func(r *Rotato) {
		r.symbols = []string{"▹▹▹▹▹", "▸▹▹▹▹", "▹▸▹▹▹", "▹▹▸▹▹", "▹▹▹▸▹", "▹▹▹▹▸"}
	}
}

// WithSymbolsArrows4 returns an option function that sets the spinner unicode
// animation with arrows.
//
//	"←", "↖", "↑", "↗", "→", "↘", "↓", "↙".
func WithSymbolsArrows4() Option {
	return func(r *Rotato) {
		r.symbols = []string{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"}
	}
}

// WithSymbolsCircles returns an option function that sets the spinner unicode
// animation with circles
//
//	"o", "O", "@", "*".
func WithSymbolsCircles() Option {
	return func(r *Rotato) {
		r.symbols = []string{"o", "O", "@", "*"}
	}
}

// WithSymbolsCircles2 returns an option function that sets the spinner unicode
// animation with circles.
//
//	".", "o", "O", "°", "O", "o", ".".
func WithSymbolsCircles2() Option {
	return func(r *Rotato) {
		r.symbols = []string{".", "o", "O", "°", "O", "o", "."}
	}
}

// WithSymbolsCircles3 returns an option function that sets the spinner unicode
// animation with circles.
//
//	"●", "●", "●", "●".
func WithSymbolsCircles3() Option {
	return func(r *Rotato) {
		r.symbols = []string{"●", "●", "●", "●"}
	}
}

// WithSymbolsCircles4 returns an option function that sets the spinner unicode
// animation with circles.
//
//	"", "", "", "".
func WithSymbolsCircles4() Option {
	return func(r *Rotato) {
		r.symbols = []string{"", "", "", ""}
	}
}

// WithSymbolsCircles5 returns an option function that sets the spinner unicode
// animation with circles.
//
//	"", "", "", "".
func WithSymbolsCircles5() Option {
	return func(r *Rotato) {
		r.symbols = []string{"", "", "", "", ""}
	}
}

// WithSymbolsCircles6 returns an option function that sets the spinner unicode
// animation with circles.
//
//	"", "", "", "", "".
func WithSymbolsCircles6() Option {
	return func(r *Rotato) {
		r.symbols = []string{"", "", "", "", "", ""}
	}
}

// WithSymbolsCircles7 returns an option function that sets the spinner unicode
// animation with circles.
//
//	"", "", "", "", "", "".
func WithSymbolsCircles7() Option {
	return func(r *Rotato) {
		r.symbols = []string{"", "", "", "", "", ""}
	}
}

// WithSymbolsBounce returns an option function that sets the spinner unicode
// animation with circles.
//
//	"[    ]", "[=   ]", "[==  ]", "[=== ]", "[ ===]", "[  ==]", "[   =]".
func WithSymbolsBounce() Option {
	return func(r *Rotato) {
		r.symbols = []string{
			"[    ]",
			"[=   ]",
			"[==  ]",
			"[=== ]",
			"[ ===]",
			"[  ==]",
			"[   =]",
			"[    ]",
			"[   =]",
			"[  ==]",
			"[ ===]",
			"[====]",
			"[=== ]",
			"[==  ]",
			"[=   ]",
		}
	}
}

// WithSymbolsBounceBall returns an option function that sets the spinner unicode
// animation with circles.
//
//	"( ●    )", "(  ●   )", "(   ●  )", "(    ● )", "(     ●)".
func WithSymbolsBounceBall() Option {
	return func(r *Rotato) {
		r.symbols = []string{
			"( ●    )",
			"(  ●   )",
			"(   ●  )",
			"(    ● )",
			"(     ●)",
			"(    ● )",
			"(   ●  )",
			"(  ●   )",
			"( ●    )",
			"(●     )",
		}
	}
}

// WithSymbolsToggle returns an option function that sets the spinner unicode
// animation with toggle symbols.
//
//	"■", "□", "▪", "▫".
func WithSymbolsToggle() Option {
	return func(r *Rotato) {
		r.symbols = []string{"■", "□", "▪", "▫"}
	}
}

// WithSymbolsToggle2 returns an option function that sets the spinner unicode
// animation with toggle symbols.
//
//	"=", "*", "-".
func WithSymbolsToggle2() Option {
	return func(r *Rotato) {
		r.symbols = []string{"=", "*", "-"}
	}
}

// WithSymbolsToggle3 returns an option function that sets the spinner unicode
// animation with toggle symbols.
//
//	"◉", "◎".
func WithSymbolsToggle3() Option {
	return func(r *Rotato) {
		r.symbols = []string{"◉", "◎"}
	}
}

// WithSymbolsLoading returns an option function that sets the spinner unicode
// animation with loading symbols.
//
//	"loading....".
func WithSymbolsLoading() Option {
	return func(r *Rotato) {
		r.symbols = []string{
			"l      ",
			"lo     ",
			"loa    ",
			"load   ",
			"loadi  ",
			"loadin ",
			"loading",
			"loading.",
			"loading..",
			"loading...",
			"loading....",
		}
	}
}
