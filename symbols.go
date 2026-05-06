package rotato

import "sort"

func Spinners() []SpinnerStyle {
	return registry
}

func Names() []string {
	names := make([]string, len(registry))
	for i, s := range registry {
		names[i] = s.Name
	}
	sort.Strings(names)
	return names
}

func ByName(name string) (SpinnerStyle, bool) {
	for _, s := range registry {
		if s.Name == name {
			return s, true
		}
	}
	return SpinnerStyle{}, false
}

func ByGroup(group SpinnerGroup) []SpinnerStyle {
	var out []SpinnerStyle
	for _, s := range registry {
		if s.Group == group {
			out = append(out, s)
		}
	}
	return out
}

func setSymbols(s []string) Option {
	return func(r *Rotato) {
		r.symbols = s
	}
}

// WithSymbols returns an option function that sets the spinner unicode
// animation.
func WithSymbols(symbols ...string) Option { return setSymbols(symbols) }

// WithSymbolsDefault returns an option function that sets the spinner unicode
// animation.
//
//	⠋ ⠙ ⠹ ⠸ ⠼ ⠴ ⠦ ⠧ ⠇ ⠏
func WithSymbolsDefault() Option { return setSymbols(defaultSymbols) }

// WithSymbolsBlock returns an option function that sets the spinner unicode
// animation with blocks.
//
//	░ ▒ ▒ ░ ▓
func WithSymbolsBlock() Option { return setSymbols(block) }

// WithSymbolsBarBlock returns an option function that sets the spinner
// unicode animation with bars.
//
//	█▒▒▒▒▒▒▒▒▒ ███▒▒▒▒▒▒▒ █████▒▒▒▒▒ ███████▒▒▒ ██████████
func WithSymbolsBarBlock() Option { return setSymbols(blockbar) }

// WithSymbolsBarBlock2 returns an option function that sets the spinner
// unicode animation with bars.
//
//	[|       ] [||      ] [|||     ] [||||    ] [|||||   ] [||||||  ] [||||||| ] [||||||||].
func WithSymbolsBarBlock2() Option { return setSymbols(blockbar2) }

// WithSymbolsBarBlock3 returns an option function that sets the spinner
// unicode animation with bars.
//
//	[=       ] [==      ] [===     ] [====    ] [=====   ] [======  ] [======= ] [========]
func WithSymbolsBarBlock3() Option { return setSymbols(blockbar3) }

// WithSymbolsBarBlock4 returns an option function that sets the spinner
// unicode animation with bars.
//
//	| || ||| |||| ||||| |||||| ||||||| ||||||||
//	||||||| |||||| ||||| |||| ||| || |
func WithSymbolsBarBlock4() Option { return setSymbols(blockbar4) }

// WithSymbolsBarBlock5 returns an option function that sets the spinner
// unicode animation with bars.
//
//	[*-------] [-*------] [--*-----] [---*----] [----*---]
//	[-----*--] [------*-] [-------*] [------*-] [-----*--]
//	[----*---] [---*----] [--*-----] [-*------] [*-------]
func WithSymbolsBarBlock5() Option { return setSymbols(blockbar5) }

// WithSymbolsBarBlock6 returns an option function that sets the spinner
// unicode animation with bars.
//
//	·----- -·---- --·--- ---·-- ----·- -----· ----·- ---·-- --·--- -·---- ·-----
func WithSymbolsBarBlock6() Option { return setSymbols(blockbar6) }

// WithSymbolsBarBlock7 returns an option function that sets the spinner
// unicode animation with bars.
//
//	■      ■■     ■■■    ■■■■   ■■■■■  ■■■■■■
func WithSymbolsBarBlock7() Option { return setSymbols(blockbar7) }

// WithSymbolsBlockPretty returns an option function that sets the spinner
// unicode animation with pretty blocks.
//
//	      
func WithSymbolsBlockPretty() Option { return setSymbols(blockbarpretty) }

// WithSymbolsDots returns an option function that sets the spinner unicode
// animation with braille patterns.
//
//	⣾ ⣽ ⣻ ⢿ ⡿ ⣟ ⣯ ⣷
func WithSymbolsDots() Option { return setSymbols(dots) }

// WithSymbolsDots3 returns an option function that sets the spinner unicode
// animation with dots.
//
//	⠄ ⠆ ⠇ ⠋ ⠙ ⠸ ⠰ ⠠ ⠰ ⠸ ⠙ ⠋ ⠇ ⠆
func WithSymbolsDots3() Option { return setSymbols(dots3) }

// WithSymbolsDots4 returns an option function that sets the spinner unicode
// animation with dots.
//
//	⠁ ⠂ ⠄ ⡀ ⢀ ⠠ ⠐ ⠈
func WithSymbolsDots4() Option { return setSymbols(dots4) }

// WithSymbolsDots5 returns an option function that sets the spinner unicode
// animation with dots.
//
//	⠁⠁⠉⠙⠚⠒⠂⠂⠒⠲⠴⠤⠄⠄⠤⠠⠠⠤⠦⠖⠒⠐⠐⠒⠓⠋⠉⠈⠈
func WithSymbolsDots5() Option { return setSymbols(dots5) }

// WithSymbolsLines returns an option function that sets the spinner unicode
// animation with lines.
//
//	⠂ - – — – -
func WithSymbolsLines() Option { return setSymbols(lines) }

// WithSymbolsWave returns an option function that sets the spinner unicode
// animation with wave patterns.
//
//	⢄ ⢂ ⢁ ⡀ ⠈ ⠘ ⠸
func WithSymbolsWave() Option { return setSymbols(wave) }

// WithSymbolsGrow returns an option function that sets the spinner unicode
// animation with growing bars.
//
//	▉ ▊ ▋ ▌ ▍ ▎ ▏
func WithSymbolsGrow() Option { return setSymbols(grow) }

// WithSymbolsGrowVert returns an option function that sets the spinner unicode
// animation with growing bars.
//
//	▁ ▃ ▄ ▅ ▆ ▇ ▆ ▅ ▄ ▃
func WithSymbolsGrowVert() Option { return setSymbols(growvert) }

// WithSymbolsMoon returns an option function that sets the spinner unicode
// animation with moon phases.
//
//	🌑 🌒 🌓 🌔 🌕 🌖 🌗 🌘
func WithSymbolsMoon() Option { return setSymbols(moon) }

// WithSymbolsPipe returns an option function that sets the spinner unicode
// animation with pipe characters.
//
//	| / - \\
func WithSymbolsPipe() Option { return setSymbols(pipe) }

// WithSymbolsPipe2 returns an option function that sets the spinner unicode
// animation with pipe characters.
//
//	┤ ┘ ┴ └ ├ ┌ ┬ ┐
func WithSymbolsPipe2() Option { return setSymbols(pipe2) }

// WithSymbolsSquare returns an option function that sets the spinner unicode
// animation with square segments.
//
//	▖ ▘ ▝ ▗
func WithSymbolsSquare() Option { return setSymbols(square) }

// WithSymbolsSquare2 returns an option function that sets the spinner unicode
// animation with square segments.
//
//	    
func WithSymbolsSquare2() Option { return setSymbols(square2) }

// WithSymbolsClock returns an option function that sets the spinner unicode
// animation with clock symbols.
//
//	🕛 🕐 🕑 🕒 🕓 🕔 🕕 🕖 🕗 🕘 🕙 🕚
func WithSymbolsClock() Option { return setSymbols(clock) }

// WithSymbolsDiamond returns an option function that sets the spinner unicode
// animation with diamond symbols.
//
//	◇ ◈ ⬟ ⬞
func WithSymbolsDiamond() Option { return setSymbols(diamond) }

// WithSymbolsDiamond2 returns an option function that sets the spinner unicode
// animation with diamond symbols.
//
//	   
func WithSymbolsDiamond2() Option { return setSymbols(diamond2) }

// WithSymbolsPlusCross returns an option function that sets the spinner unicode
// animation with plus and cross symbols.
//
//   - x
func WithSymbolsPlusCross() Option { return setSymbols(pluscross) }

// WithSymbolsArrows returns an option function that sets the spinner unicode
// animation with arrows.
//
//	< << <<< - > >> >>>
func WithSymbolsArrows() Option { return setSymbols(arrow) }

// WithSymbolsArrows2 returns an option function that sets the spinner unicode
// animation with arrows.
//
//	>    >>   >>>  >>>>
func WithSymbolsArrows2() Option { return setSymbols(arrow2) }

// WithSymbolsArrows3 returns an option function that sets the spinner unicode
// animation with arrows.
//
//	▹▹▹▹▹ ▸▹▹▹▹ ▹▸▹▹▹ ▹▹▸▹▹ ▹▹▹▸▹ ▹▹▹▹▸
func WithSymbolsArrows3() Option { return setSymbols(arrow3) }

// WithSymbolsArrows4 returns an option function that sets the spinner unicode
// animation with arrows.
//
//	← ↖ ↑ ↗ → ↘ ↓ ↙
func WithSymbolsArrows4() Option { return setSymbols(arrow4) }

// WithSymbolsCircles returns an option function that sets the spinner unicode
// animation with circles
//
//	o O @ *
func WithSymbolsCircles() Option { return setSymbols(circle) }

// WithSymbolsCircles2 returns an option function that sets the spinner unicode
// animation with circles.
//
//	. o O ° O o .
func WithSymbolsCircles2() Option { return setSymbols(circle2) }

// WithSymbolsCircles3 returns an option function that sets the spinner unicode
// animation with circles.
//
//	● ● ● ●
func WithSymbolsCircles3() Option { return setSymbols(circle3) }

// WithSymbolsCircles4 returns an option function that sets the spinner unicode
// animation with circles.
//
//	   
func WithSymbolsCircles4() Option { return setSymbols(circle4) }

// WithSymbolsCircles5 returns an option function that sets the spinner unicode
// animation with circles.
//
//	   
func WithSymbolsCircles5() Option { return setSymbols(circle5) }

// WithSymbolsCircles6 returns an option function that sets the spinner unicode
// animation with circles.
//
//	    
func WithSymbolsCircles6() Option { return setSymbols(circle6) }

// WithSymbolsCircles7 returns an option function that sets the spinner unicode
// animation with circles.
//
//	     
func WithSymbolsCircles7() Option { return setSymbols(circle7) }

// WithSymbolsBounce returns an option function that sets the spinner unicode
// animation with circles.
//
//	[    ] [=   ] [==  ] [=== ] [ ===] [  ==] [   =]
func WithSymbolsBounce() Option { return setSymbols(bounce) }

// WithSymbolsBounceBall returns an option function that sets the spinner unicode
// animation with circles.
//
//	( ●    ) (  ●   ) (   ●  ) (    ● ) (     ●)
func WithSymbolsBounceBall() Option { return setSymbols(bounceball) }

// WithSymbolsToggle returns an option function that sets the spinner unicode
// animation with toggle symbols.
//
//	■ □ ▪ ▫
func WithSymbolsToggle() Option { return setSymbols(toggle) }

// WithSymbolsToggle2 returns an option function that sets the spinner unicode
// animation with toggle symbols.
//
//	= * -
func WithSymbolsToggle2() Option { return setSymbols(toggle2) }

// WithSymbolsToggle3 returns an option function that sets the spinner unicode
// animation with toggle symbols.
//
//	◉ ◎
func WithSymbolsToggle3() Option { return setSymbols(toggle3) }

// WithSymbolsLoading returns an option function that sets the spinner unicode
// animation with loading symbols.
//
//	loading....
func WithSymbolsLoading() Option { return setSymbols(loading) }

// WithSymbolsTriangles returns an option function that sets the spinner
// unicode animation with rotating triangles.
//
//	▲ ▶ ▼ ◀
func WithSymbolsTriangles() Option { return setSymbols(triangles) }

// WithSymbolsCubes returns an option function that sets the spinner
// unicode animation with cube rotation.
//
//	▖ ▘ ▝ ▗
func WithSymbolsCubes() Option { return setSymbols(cubes) }

// WithSymbolsThinking returns an option function that sets the spinner
// animation with growing question marks.
//
//	? ?? ??? ???? ?????
func WithSymbolsThinking() Option { return setSymbols(question) }

// WithSymbolsPingPong returns an option function that sets the spinner
// animation with expanding and contracting brackets.
//
//	<     > <    > <   > <  > < > <>< < > <  > <   > <    >
func WithSymbolsPingPong() Option { return setSymbols(pingpong) }

// WithSymbolsMatrix returns an option function that sets the spinner
// unicode animation with matrix-style loading.
//
//	╔═══╗ ║▓▓▓║ ║░▓▓║ ║░░▓║ ║░░░║ ╚═══╝
func WithSymbolsMatrix() Option { return setSymbols(matrix) }

// WithSymbolsHex returns an option function that sets the spinner
// animation with hexadecimal counting.
//
//	0x0 0x1 0x2 0x3 0x4 0x5 0x6 0x7 0x8 0x9 0xA 0xB 0xC 0xD 0xE 0xF
func WithSymbolsHex() Option { return setSymbols(hexsymbols) }

// WithSymbolsPacman returns an option function that sets the spinner
// unicode animation with pacman movement.
//
//	󰮯··· ·󰮯·· ··󰮯· ···󰮯
func WithSymbolsPacman() Option { return setSymbols(pacman) }

// WithSymbolsBoxFill returns an option function that sets the spinner
// animation with progressively filling box.
//
//	[          ] [■         ] [■■        ] [■■■       ] [■■■■      ]
//	[■■■■■ ] [■■■■■■    ] [■■■■■■■   ] [■■■■■■■■  ] [■■■■■■■■■ ] [■■■■■■■■■■]
func WithSymbolsBoxFill() Option { return setSymbols(boxfill) }

// WithSymbolsBoxFillShort returns an option function that sets the spinner
// animation with progressively filling box.
//
//	[      ] [■     ] [■■    ] [■■■   ] [■■■■  ] [■■■■■ ] [■■■■■■]
func WithSymbolsBoxFillShort() Option { return setSymbols(boxfillshort) }

// WithSymbolsSnail returns an option function that sets the spinner
// animation with growing snail trail.
//
//	@ @- @-- @--- @---- @-----
func WithSymbolsSnail() Option { return setSymbols(snail) }

// WithSymbolsWorm returns an option function that sets the spinner
// animation with growing and shrinking worm.
//
//	~ ~~ ~~~ ~~~~ ~~~~~ ~~~~ ~~~ ~~ ~
func WithSymbolsWorm() Option { return setSymbols(worm) }

// WithSymbolsWorm2 returns an option function that sets the spinner
// animation with growing and shrinking worm.
//
//	~ ~~ ~~~ ~~~~ ~~~~~ ~~~~ ~~~ ~~ ~
func WithSymbolsWorm2() Option { return setSymbols(worm2) }

// WithSymbolsMathOps returns an option function that sets the spinner
// unicode animation with mathematical operators.
//
//   - - × ÷ = ≠ ≈ ≤ ≥
func WithSymbolsMathOps() Option { return setSymbols(mathops) }

// WithSymbolsGreek returns an option function that sets the spinner
// unicode animation with Greek letters.
//
//	α β γ δ ε ζ η θ
func WithSymbolsGreek() Option { return setSymbols(greek) }

// WithSymbolsCorners returns an option function that sets the spinner
// unicode animation with box corners.
//
//	┌ ┐ └ ┘
func WithSymbolsCorners() Option { return setSymbols(corners) }

// WithSymbolsSlash returns an option function that sets the spinner
// animation with increasing forward slashes.
//
//	/ // /// //// /////
func WithSymbolsSlash() Option { return setSymbols(slash) }

// WithSymbolsBackslash returns an option function that sets the spinner
// animation with increasing backslashes.
//
//	\ \\ \\\ \\\\ \\\\\
func WithSymbolsBackslash() Option { return setSymbols(backslash) }

// WithSymbolsMarquee returns an option function that sets the spinner
// animation with moving marquee arrow.
//
//	[          ] [ >        ] [  >       ] [   >      ] [    >     ]
//	[     >    ] [      >   ] [       >  ] [        > ] [         >]
func WithSymbolsMarquee() Option { return setSymbols(marquee) }

// WithSymbolsFade returns an option function that sets the spinner
// unicode animation with fading block.
//
//	█ ▓ ▒ ░   ░ ▒ ▓
func WithSymbolsFade() Option { return setSymbols(fade) }

// WithSymbolsMath returns an option function that sets the spinner
// unicode animation with mathematical symbols.
//
//	∀ ∃ ∈ ∉ ∋ ∌ ⊆ ⊂ ⊄ ⊇ ⊃ ⊅
func WithSymbolsMath() Option { return setSymbols(logicsymbols) }

// WithSymbolsCurrency returns an option function that sets the spinner
// unicode animation with currency symbols.
//
//	$ € £ ¥ ₿ ₹
func WithSymbolsCurrency() Option { return setSymbols(currency) }

// WithSymbolsGeometric returns an option function that sets the spinner
// unicode animation with geometric shapes.
//
//	△ ◊ ◈ ◇ ○ ● ◐ ◑ ◒ ◓
func WithSymbolsGeometric() Option { return setSymbols(geometric) }

// WithSymbolsRunner returns an option function that sets the spinner unicode animation.
//
//	▁▁▁▁▁ ▂▁▁▁▁ ▃▂▁▁▁ ▄▃▂▁▁ ▅▄▃▂▁ ▆▅▄▃▂ ▇▆▅▄▃ █▇▆▅▄
func WithSymbolsRunner() Option { return setSymbols(runner) }

// WithSymbolsCursorBlink returns a option function that sets the spinner
// unicode animation
//
//	"_", " ", "_", " "
func WithSymbolsCursorBlink() Option { return setSymbols(cursorBlink) }

// WithSymbolsEllipsis returns a option function that sets the spinner
// unicode animation
//
//	".  " ".. " "..." " .." "  ." "   "
func WithSymbolsEllipsis() Option { return setSymbols(ellipsis) }

// WithSymbolsBrailleWave returns a option function that sets the spinner
// unicode animation
//
//	⡀ ⡄ ⡆ ⡇ ⡏ ⡟ ⡿ ⣿ ⡿ ⡟ ⡏ ⡇ ⡆ ⡄
func WithSymbolsBrailleWave() Option { return setSymbols(brailleWave) }

// WithSymbolsOrbit returns a option function that sets the spinner
// unicode animation
//
//	◐ ◓ ◑ ◒
func WithSymbolsOrbit() Option { return setSymbols(orbit) }

// WithSymbolsSweep returns a option function that sets the spinner
// unicode animation
//
//	←──── ─←─── ──←── ───←─ ────← ───→─ ──→── ─→─── →────
func WithSymbolsSweep() Option { return setSymbols(sweep) }

// WithSymbolsPulse returns a option function that sets the spinner
// unicode animation
//
//	░ ▒ ▓ █ ▓ ▒
func WithSymbolsPulse() Option { return setSymbols(pulse) }
