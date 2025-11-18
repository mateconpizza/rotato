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
//	░ ▒ ▒ ░ ▓
func WithSymbolsBlock() Option {
	return func(r *Rotato) {
		r.symbols = []string{"░", "▒", "▒", "░", "▓"}
	}
}

// WithSymbolsBarBlock returns an option function that sets the spinner
// unicode animation with bars.
//
//	█▒▒▒▒▒▒▒▒▒ ███▒▒▒▒▒▒▒ █████▒▒▒▒▒ ███████▒▒▒ ██████████
func WithSymbolsBarBlock() Option {
	return func(r *Rotato) {
		r.symbols = []string{"█▒▒▒▒▒▒▒▒▒", "███▒▒▒▒▒▒▒", "█████▒▒▒▒▒", "███████▒▒▒", "██████████"}
	}
}

// WithSymbolsBarBlock2 returns an option function that sets the spinner
// unicode animation with bars.
//
//	[|       ] [||      ] [|||     ] [||||    ] [|||||   ] [||||||  ] [||||||| ] [||||||||].
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
//	[=       ] [==      ] [===     ] [====    ] [=====   ] [======  ] [======= ] [========]
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
//	| || ||| |||| ||||| |||||| ||||||| ||||||||
//	||||||| |||||| ||||| |||| ||| || |
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
//	[*-------] [-*------] [--*-----] [---*----] [----*---]
//	[-----*--] [------*-] [-------*] [------*-] [-----*--]
//	[----*---] [---*----] [--*-----] [-*------] [*-------]
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
//	·----- -·---- --·--- ---·-- ----·- -----· ----·- ---·-- --·--- -·---- ·-----
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
//	      
func WithSymbolsBlockPretty() Option {
	return func(r *Rotato) {
		//  
		r.symbols = []string{"", "", "", "", "", "", ""}
	}
}

// WithSymbolsDots returns an option function that sets the spinner unicode
// animation with braille patterns.
//
//	⣾ ⣽ ⣻ ⢿ ⡿ ⣟ ⣯ ⣷
func WithSymbolsDots() Option {
	return func(r *Rotato) {
		r.symbols = []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
	}
}

// WithSymbolsDots2 returns an option function that sets the spinner unicode
// animation with dots.
//
//	'  . . . .'
//	'.   . . .'
//	'. .   . .'
//	'. . .   .'
//	'. . . .  '
//	'. . . . .'
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
//	⠄ ⠆ ⠇ ⠋ ⠙ ⠸ ⠰ ⠠ ⠰ ⠸ ⠙ ⠋ ⠇ ⠆
func WithSymbolsDots3() Option {
	return func(r *Rotato) {
		r.symbols = []string{"⠄", "⠆", "⠇", "⠋", "⠙", "⠸", "⠰", "⠠", "⠰", "⠸", "⠙", "⠋", "⠇", "⠆"}
	}
}

// WithSymbolsDots4 returns an option function that sets the spinner unicode
// animation with dots.
//
//	⠁ ⠂ ⠄ ⡀ ⢀ ⠠ ⠐ ⠈
func WithSymbolsDots4() Option {
	return func(r *Rotato) {
		r.symbols = []string{"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"}
	}
}

// WithSymbolsDots5 returns an option function that sets the spinner unicode
// animation with dots.
//
//	⠁⠁⠉⠙⠚⠒⠂⠂⠒⠲⠴⠤⠄⠄⠤⠠⠠⠤⠦⠖⠒⠐⠐⠒⠓⠋⠉⠈⠈
func WithSymbolsDots5() Option {
	return func(r *Rotato) {
		r.symbols = []string{
			"⠁", "⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄",
			"⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈",
			"⠈",
		}
	}
}

// WithSymbolsLines returns an option function that sets the spinner unicode
// animation with lines.
//
//	⠂ - – — – -
func WithSymbolsLines() Option {
	return func(r *Rotato) {
		r.symbols = []string{"⠂", "-", "–", "—", "–", "-"}
	}
}

// WithSymbolsWave returns an option function that sets the spinner unicode
// animation with wave patterns.
//
//	⢄ ⢂ ⢁ ⡀ ⠈ ⠘ ⠸
func WithSymbolsWave() Option {
	return func(r *Rotato) {
		r.symbols = []string{"⢄", "⢂", "⢁", "⡀", "⠈", "⠘", "⠸"}
	}
}

// WithSymbolsGrow returns an option function that sets the spinner unicode
// animation with growing bars.
//
//	▉ ▊ ▋ ▌ ▍ ▎ ▏
func WithSymbolsGrow() Option {
	return func(r *Rotato) {
		r.symbols = []string{"▉", "▊", "▋", "▌", "▍", "▎", "▏"}
	}
}

// WithSymbolsGrowVert returns an option function that sets the spinner unicode
// animation with growing bars.
//
//	▁ ▃ ▄ ▅ ▆ ▇ ▆ ▅ ▄ ▃
func WithSymbolsGrowVert() Option {
	return func(r *Rotato) {
		r.symbols = []string{"▁", "▃", "▄", "▅", "▆", "▇", "▆", "▅", "▄", "▃"}
	}
}

// WithSymbolsMoon returns an option function that sets the spinner unicode
// animation with moon phases.
//
//	🌑 🌒 🌓 🌔 🌕 🌖 🌗 🌘
func WithSymbolsMoon() Option {
	return func(r *Rotato) {
		r.symbols = []string{"🌑", "🌒", "🌓", "🌔", "🌕", "🌖", "🌗", "🌘"}
	}
}

// WithSymbolsPipe returns an option function that sets the spinner unicode
// animation with pipe characters.
//
//	| / - \\
func WithSymbolsPipe() Option {
	return func(r *Rotato) {
		r.symbols = []string{"|", "/", "-", "\\"}
	}
}

// WithSymbolsPipe2 returns an option function that sets the spinner unicode
// animation with pipe characters.
//
//	┤ ┘ ┴ └ ├ ┌ ┬ ┐
func WithSymbolsPipe2() Option {
	return func(r *Rotato) {
		r.symbols = []string{"┤", "┘", "┴", "└", "├", "┌", "┬", "┐"}
	}
}

// WithSymbolsSquare returns an option function that sets the spinner unicode
// animation with square segments.
//
//	▖ ▘ ▝ ▗
func WithSymbolsSquare() Option {
	return func(r *Rotato) {
		r.symbols = []string{"▖", "▘", "▝", "▗"}
	}
}

// WithSymbolsSquare2 returns an option function that sets the spinner unicode
// animation with square segments.
//
//	    
func WithSymbolsSquare2() Option {
	return func(r *Rotato) {
		r.symbols = []string{"", "", "", "", ""}
	}
}

// WithSymbolsClock returns an option function that sets the spinner unicode
// animation with clock symbols.
//
//	🕛 🕐 🕑 🕒 🕓 🕔 🕕 🕖 🕗 🕘 🕙 🕚
func WithSymbolsClock() Option {
	return func(r *Rotato) {
		r.symbols = []string{"🕛", "🕐", "🕑", "🕒", "🕓", "🕔", "🕕", "🕖", "🕗", "🕘", "🕙", "🕚"}
	}
}

// WithSymbolsDiamond returns an option function that sets the spinner unicode
// animation with diamond symbols.
//
//	◇ ◈ ⬟ ⬞
func WithSymbolsDiamond() Option {
	return func(r *Rotato) {
		r.symbols = []string{"◇", "◈", "⬟", "⬞"}
	}
}

// WithSymbolsDiamond2 returns an option function that sets the spinner unicode
// animation with diamond symbols.
//
//	   
func WithSymbolsDiamond2() Option {
	return func(r *Rotato) {
		r.symbols = []string{"", "", "", ""}
	}
}

// WithSymbolsPlusCross returns an option function that sets the spinner unicode
// animation with plus and cross symbols.
//
//   - x
func WithSymbolsPlusCross() Option {
	return func(r *Rotato) {
		r.symbols = []string{"+", "x"}
	}
}

// WithSymbolsArrows returns an option function that sets the spinner unicode
// animation with arrows.
//
//	< << <<< - > >> >>>
func WithSymbolsArrows() Option {
	return func(r *Rotato) {
		r.symbols = []string{"<", "<<", "<<<", "-", ">", ">>", ">>>"}
	}
}

// WithSymbolsArrows2 returns an option function that sets the spinner unicode
// animation with arrows.
//
//	>    >>   >>>  >>>>
func WithSymbolsArrows2() Option {
	return func(r *Rotato) {
		r.symbols = []string{">   ", ">>  ", ">>> ", ">>>>"}
	}
}

// WithSymbolsArrows3 returns an option function that sets the spinner unicode
// animation with arrows.
//
//	▹▹▹▹▹ ▸▹▹▹▹ ▹▸▹▹▹ ▹▹▸▹▹ ▹▹▹▸▹ ▹▹▹▹▸
func WithSymbolsArrows3() Option {
	return func(r *Rotato) {
		r.symbols = []string{"▹▹▹▹▹", "▸▹▹▹▹", "▹▸▹▹▹", "▹▹▸▹▹", "▹▹▹▸▹", "▹▹▹▹▸"}
	}
}

// WithSymbolsArrows4 returns an option function that sets the spinner unicode
// animation with arrows.
//
//	← ↖ ↑ ↗ → ↘ ↓ ↙
func WithSymbolsArrows4() Option {
	return func(r *Rotato) {
		r.symbols = []string{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"}
	}
}

// WithSymbolsCircles returns an option function that sets the spinner unicode
// animation with circles
//
//	o O @ *
func WithSymbolsCircles() Option {
	return func(r *Rotato) {
		r.symbols = []string{"o", "O", "@", "*"}
	}
}

// WithSymbolsCircles2 returns an option function that sets the spinner unicode
// animation with circles.
//
//	. o O ° O o .
func WithSymbolsCircles2() Option {
	return func(r *Rotato) {
		r.symbols = []string{".", "o", "O", "°", "O", "o", "."}
	}
}

// WithSymbolsCircles3 returns an option function that sets the spinner unicode
// animation with circles.
//
//	● ● ● ●
func WithSymbolsCircles3() Option {
	return func(r *Rotato) {
		r.symbols = []string{"●", "●", "●", "●"}
	}
}

// WithSymbolsCircles4 returns an option function that sets the spinner unicode
// animation with circles.
//
//	   
func WithSymbolsCircles4() Option {
	return func(r *Rotato) {
		r.symbols = []string{"", "", "", ""}
	}
}

// WithSymbolsCircles5 returns an option function that sets the spinner unicode
// animation with circles.
//
//	   
func WithSymbolsCircles5() Option {
	return func(r *Rotato) {
		r.symbols = []string{"", "", "", "", ""}
	}
}

// WithSymbolsCircles6 returns an option function that sets the spinner unicode
// animation with circles.
//
//	    
func WithSymbolsCircles6() Option {
	return func(r *Rotato) {
		r.symbols = []string{"", "", "", "", "", ""}
	}
}

// WithSymbolsCircles7 returns an option function that sets the spinner unicode
// animation with circles.
//
//	     
func WithSymbolsCircles7() Option {
	return func(r *Rotato) {
		r.symbols = []string{"", "", "", "", "", ""}
	}
}

// WithSymbolsBounce returns an option function that sets the spinner unicode
// animation with circles.
//
//	[    ] [=   ] [==  ] [=== ] [ ===] [  ==] [   =]
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
//	( ●    ) (  ●   ) (   ●  ) (    ● ) (     ●)
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
//	■ □ ▪ ▫
func WithSymbolsToggle() Option {
	return func(r *Rotato) {
		r.symbols = []string{"■", "□", "▪", "▫"}
	}
}

// WithSymbolsToggle2 returns an option function that sets the spinner unicode
// animation with toggle symbols.
//
//	= * -
func WithSymbolsToggle2() Option {
	return func(r *Rotato) {
		r.symbols = []string{"=", "*", "-"}
	}
}

// WithSymbolsToggle3 returns an option function that sets the spinner unicode
// animation with toggle symbols.
//
//	◉ ◎
func WithSymbolsToggle3() Option {
	return func(r *Rotato) {
		r.symbols = []string{"◉", "◎"}
	}
}

// WithSymbolsLoading returns an option function that sets the spinner unicode
// animation with loading symbols.
//
//	loading....
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

// WithSymbolsTriangles returns an option function that sets the spinner
// unicode animation with rotating triangles.
//
//	▲ ▶ ▼ ◀
func WithSymbolsTriangles() Option {
	return func(r *Rotato) {
		r.symbols = []string{"▲", "▶", "▼", "◀"}
	}
}

// WithSymbolsCubes returns an option function that sets the spinner
// unicode animation with cube rotation.
//
//	▖ ▘ ▝ ▗
func WithSymbolsCubes() Option {
	return func(r *Rotato) {
		r.symbols = []string{"▖", "▘", "▝", "▗"}
	}
}

// WithSymbolsThinking returns an option function that sets the spinner
// animation with growing question marks.
//
//	? ?? ??? ???? ?????
func WithSymbolsThinking() Option {
	return func(r *Rotato) {
		r.symbols = []string{"?", "??", "???", "????", "?????"}
	}
}

// WithSymbolsPingPong returns an option function that sets the spinner
// animation with expanding and contracting brackets.
//
//	<     > <    > <   > <  > < > <>< < > <  > <   > <    >
func WithSymbolsPingPong() Option {
	return func(r *Rotato) {
		r.symbols = []string{"<     >", "<    >", "<   >", "<  >", "< >", "<><", "< >", "<  >", "<   >", "<    >"}
	}
}

// WithSymbolsTimer returns an option function that sets the spinner
// animation with counting timer.
//
//	00:00 00:01 00:02 00:03 00:04 00:05
func WithSymbolsTimer() Option {
	return func(r *Rotato) {
		r.symbols = []string{"00:00", "00:01", "00:02", "00:03", "00:04", "00:05"}
	}
}

// WithSymbolsMatrix returns an option function that sets the spinner
// unicode animation with matrix-style loading.
//
//	╔═══╗ ║▓▓▓║ ║░▓▓║ ║░░▓║ ║░░░║ ╚═══╝
func WithSymbolsMatrix() Option {
	return func(r *Rotato) {
		r.symbols = []string{
			"╔═══╗",
			"║▓▓▓║",
			"║░▓▓║",
			"║░░▓║",
			"║░░░║",
			"╚═══╝",
		}
	}
}

// WithSymbolsHex returns an option function that sets the spinner
// animation with hexadecimal counting.
//
//	0x0 0x1 0x2 0x3 0x4 0x5 0x6 0x7 0x8 0x9 0xA 0xB 0xC 0xD 0xE 0xF
func WithSymbolsHex() Option {
	return func(r *Rotato) {
		r.symbols = []string{
			"0x0", "0x1", "0x2", "0x3",
			"0x4", "0x5", "0x6", "0x7",
			"0x8", "0x9", "0xA", "0xB",
			"0xC", "0xD", "0xE", "0xF",
		}
	}
}

// WithSymbolsPacman returns an option function that sets the spinner
// unicode animation with pacman movement.
//
//	ᗧ··· ·ᗧ·· ··ᗧ· ···ᗧ
func WithSymbolsPacman() Option {
	return func(r *Rotato) {
		r.symbols = []string{
			"ᗧ···",
			"·ᗧ··",
			"··ᗧ·",
			"···ᗧ",
		}
	}
}

// WithSymbolsBoxFill returns an option function that sets the spinner
// animation with progressively filling box.
//
//	[          ] [■         ] [■■        ] [■■■       ] [■■■■      ]
//	[■■■■■ ] [■■■■■■    ] [■■■■■■■   ] [■■■■■■■■  ] [■■■■■■■■■ ] [■■■■■■■■■■]
func WithSymbolsBoxFill() Option {
	return func(r *Rotato) {
		r.symbols = []string{
			"[          ]",
			"[■         ]",
			"[■■        ]",
			"[■■■       ]",
			"[■■■■      ]",
			"[■■■■■     ]",
			"[■■■■■■    ]",
			"[■■■■■■■   ]",
			"[■■■■■■■■  ]",
			"[■■■■■■■■■ ]",
			"[■■■■■■■■■■]",
		}
	}
}

// WithSymbolsSnail returns an option function that sets the spinner
// animation with growing snail trail.
//
//	@ @- @-- @--- @---- @-----
func WithSymbolsSnail() Option {
	return func(r *Rotato) {
		r.symbols = []string{
			"@     ",
			"@-    ",
			"@--   ",
			"@---  ",
			"@---- ",
			"@-----",
		}
	}
}

// WithSymbolsWorm returns an option function that sets the spinner
// animation with growing and shrinking worm.
//
//	~ ~~ ~~~ ~~~~ ~~~~~ ~~~~ ~~~ ~~ ~
func WithSymbolsWorm() Option {
	return func(r *Rotato) {
		r.symbols = []string{
			"~",
			"~~",
			"~~~",
			"~~~~",
			"~~~~~",
			"~~~~",
			"~~~",
			"~~",
			"~",
		}
	}
}

// WithSymbolsWorm2 returns an option function that sets the spinner
// animation with growing and shrinking worm.
//
//	~ ~~ ~~~ ~~~~ ~~~~~ ~~~~ ~~~ ~~ ~
func WithSymbolsWorm2() Option {
	return func(r *Rotato) {
		r.symbols = []string{
			"~    ",
			"~~   ",
			"~~~  ",
			"~~~~ ",
			"~~~~~",
			"~~~~ ",
			"~~~  ",
			"~~   ",
			"~    ",
		}
	}
}

// WithSymbolsMathOps returns an option function that sets the spinner
// unicode animation with mathematical operators.
//
//   - - × ÷ = ≠ ≈ ≤ ≥
func WithSymbolsMathOps() Option {
	return func(r *Rotato) {
		r.symbols = []string{"+", "-", "×", "÷", "=", "≠", "≈", "≤", "≥"}
	}
}

// WithSymbolsGreek returns an option function that sets the spinner
// unicode animation with Greek letters.
//
//	α β γ δ ε ζ η θ
func WithSymbolsGreek() Option {
	return func(r *Rotato) {
		r.symbols = []string{"α", "β", "γ", "δ", "ε", "ζ", "η", "θ"}
	}
}

// WithSymbolsCorners returns an option function that sets the spinner
// unicode animation with box corners.
//
//	┌ ┐ └ ┘
func WithSymbolsCorners() Option {
	return func(r *Rotato) {
		r.symbols = []string{"┌", "┐", "└", "┘"}
	}
}

// WithSymbolsBox returns an option function that sets the spinner
// unicode animation with opening and closing box.
//
//	╔════╗ ║    ║ ╚════╝
func WithSymbolsBox() Option {
	return func(r *Rotato) {
		r.symbols = []string{
			"╔════╗",
			"║    ║",
			"╚════╝",
		}
	}
}

// WithSymbolsSlash returns an option function that sets the spinner
// animation with increasing forward slashes.
//
//	/ // /// //// /////
func WithSymbolsSlash() Option {
	return func(r *Rotato) {
		r.symbols = []string{"/    ", "//   ", "///  ", "//// ", "/////"}
	}
}

// WithSymbolsBackslash returns an option function that sets the spinner
// animation with increasing backslashes.
//
//	\ \\ \\\ \\\\ \\\\\
func WithSymbolsBackslash() Option {
	return func(r *Rotato) {
		r.symbols = []string{"\\    ", "\\\\   ", "\\\\\\  ", "\\\\\\\\ ", "\\\\\\\\\\"}
	}
}

// WithSymbolsMarquee returns an option function that sets the spinner
// animation with moving marquee arrow.
//
//	[          ] [ >        ] [  >       ] [   >      ] [    >     ]
//	[     >    ] [      >   ] [       >  ] [        > ] [         >]
func WithSymbolsMarquee() Option {
	return func(r *Rotato) {
		r.symbols = []string{
			"[          ]",
			"[ >        ]",
			"[  >       ]",
			"[   >      ]",
			"[    >     ]",
			"[     >    ]",
			"[      >   ]",
			"[       >  ]",
			"[        > ]",
			"[         >]",
		}
	}
}

// WithSymbolsEyes returns an option function that sets the spinner
// unicode animation with blinking eyes.
//
//	◉◉ ●● ○○ ●●
func WithSymbolsEyes() Option {
	return func(r *Rotato) {
		r.symbols = []string{"◉◉", "●●", "○○", "●●"}
	}
}

// WithSymbolsBlink returns an option function that sets the spinner
// unicode animation with blinking dots.
//
//	◉ ◉ - - ◉ ◉
func WithSymbolsBlink() Option {
	return func(r *Rotato) {
		r.symbols = []string{"◉ ◉", "- -", "◉ ◉"}
	}
}

// WithSymbolsGradient returns an option function that sets the spinner
// unicode animation with density gradient.
//
//	░ ▒ ▓ █ ▓ ▒
func WithSymbolsGradient() Option {
	return func(r *Rotato) {
		r.symbols = []string{"░", "▒", "▓", "█", "▓", "▒"}
	}
}

// WithSymbolsFade returns an option function that sets the spinner
// unicode animation with fading block.
//
//	█ ▓ ▒ ░   ░ ▒ ▓
func WithSymbolsFade() Option {
	return func(r *Rotato) {
		r.symbols = []string{"█", "▓", "▒", "░", " ", "░", "▒", "▓"}
	}
}

// WithSymbolsMath returns an option function that sets the spinner
// unicode animation with mathematical symbols.
//
//	∀ ∃ ∈ ∉ ∋ ∌ ⊆ ⊂ ⊄ ⊇ ⊃ ⊅
func WithSymbolsMath() Option {
	return func(r *Rotato) {
		r.symbols = []string{"∀", "∃", "∈", "∉", "∋", "∌", "⊆", "⊂", "⊄", "⊇", "⊃", "⊅"}
	}
}

// WithSymbolsCurrency returns an option function that sets the spinner
// unicode animation with currency symbols.
//
//	$ € £ ¥ ₿ ₹
func WithSymbolsCurrency() Option {
	return func(r *Rotato) {
		r.symbols = []string{"$", "€", "£", "¥", "₿", "₹"}
	}
}

// WithSymbolsGeometric returns an option function that sets the spinner
// unicode animation with geometric shapes.
//
//	△ ◊ ◈ ◇ ○ ● ◐ ◑ ◒ ◓
func WithSymbolsGeometric() Option {
	return func(r *Rotato) {
		r.symbols = []string{"△", "◊", "◈", "◇", "○", "●", "◐", "◑", "◒", "◓"}
	}
}
