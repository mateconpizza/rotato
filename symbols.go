package rotato

import "sort"

func Spinners() []SpinnerStyle {
	return registry
}

func Names() []string {
	names := make([]string, len(registry))
	for i, s := range registry {
		names[i] = string(s.Name)
	}
	sort.Strings(names)
	return names
}

func ByName(name SpinnerName) (SpinnerStyle, bool) {
	sp, ok := byName[name]
	return sp, ok
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

// WithSymbols returns an option function that sets the spinner unicode
// animation.
func WithSymbols(symbols ...string) Option {
	return func(r *Rotato) {
		r.symbols = symbols
	}
}

// WithSymbolsDefault returns an option function that sets the spinner unicode
// animation.
//
//	⠋ ⠙ ⠹ ⠸ ⠼ ⠴ ⠦ ⠧ ⠇ ⠏
func WithSymbolsDefault() Option { return WithSpinnerStyle(SpinnerDefault) }

// WithSymbolsBlock returns an option function that sets the spinner unicode
// animation with blocks.
//
//	░ ▒ ▒ ░ ▓
func WithSymbolsBlock() Option { return WithSpinnerStyle(SpinnerBlock) }

// WithSymbolsBarBlock returns an option function that sets the spinner
// unicode animation with bars.
//
//	█▒▒▒▒▒▒▒▒▒ ███▒▒▒▒▒▒▒ █████▒▒▒▒▒ ███████▒▒▒ ██████████
func WithSymbolsBarBlock() Option { return WithSpinnerStyle(SpinnerBlockbar) }

// WithSymbolsBarBlock2 returns an option function that sets the spinner
// unicode animation with bars.
//
//	[|       ] [||      ] [|||     ] [||||    ] [|||||   ] [||||||  ] [||||||| ] [||||||||].
func WithSymbolsBarBlock2() Option { return WithSpinnerStyle(SpinnerBlockbar2) }

// WithSymbolsBarBlock3 returns an option function that sets the spinner
// unicode animation with bars.
//
//	[=       ] [==      ] [===     ] [====    ] [=====   ] [======  ] [======= ] [========]
func WithSymbolsBarBlock3() Option { return WithSpinnerStyle(SpinnerBlockbar3) }

// WithSymbolsBarBlock4 returns an option function that sets the spinner
// unicode animation with bars.
//
//	| || ||| |||| ||||| |||||| ||||||| ||||||||
//	||||||| |||||| ||||| |||| ||| || |
func WithSymbolsBarBlock4() Option { return WithSpinnerStyle(SpinnerBlockbar4) }

// WithSymbolsBarBlock5 returns an option function that sets the spinner
// unicode animation with bars.
//
//	[*-------] [-*------] [--*-----] [---*----] [----*---]
//	[-----*--] [------*-] [-------*] [------*-] [-----*--]
//	[----*---] [---*----] [--*-----] [-*------] [*-------]
func WithSymbolsBarBlock5() Option { return WithSpinnerStyle(SpinnerBlockbar5) }

// WithSymbolsBarBlock6 returns an option function that sets the spinner
// unicode animation with bars.
//
//	·----- -·---- --·--- ---·-- ----·- -----· ----·- ---·-- --·--- -·---- ·-----
func WithSymbolsBarBlock6() Option { return WithSpinnerStyle(SpinnerBlockbar6) }

// WithSymbolsBarBlock7 returns an option function that sets the spinner
// unicode animation with bars.
//
//	■      ■■     ■■■    ■■■■   ■■■■■  ■■■■■■
func WithSymbolsBarBlock7() Option { return WithSpinnerStyle(SpinnerBlockbar7) }

// WithSymbolsBlockPretty returns an option function that sets the spinner
// unicode animation with pretty blocks.
//
//	      
func WithSymbolsBlockPretty() Option { return WithSpinnerStyle(SpinnerBlockbarPretty) }

// WithSymbolsDots returns an option function that sets the spinner unicode
// animation with braille patterns.
//
//	⣾ ⣽ ⣻ ⢿ ⡿ ⣟ ⣯ ⣷
func WithSymbolsDots() Option { return WithSpinnerStyle(SpinnerDots) }

// WithSymbolsDots3 returns an option function that sets the spinner unicode
// animation with dots.
//
//	⠄ ⠆ ⠇ ⠋ ⠙ ⠸ ⠰ ⠠ ⠰ ⠸ ⠙ ⠋ ⠇ ⠆
func WithSymbolsDots3() Option { return WithSpinnerStyle(SpinnerDots3) }

// WithSymbolsDots4 returns an option function that sets the spinner unicode
// animation with dots.
//
//	⠁ ⠂ ⠄ ⡀ ⢀ ⠠ ⠐ ⠈
func WithSymbolsDots4() Option { return WithSpinnerStyle(SpinnerDots4) }

// WithSymbolsDots5 returns an option function that sets the spinner unicode
// animation with dots.
//
//	⠁⠁⠉⠙⠚⠒⠂⠂⠒⠲⠴⠤⠄⠄⠤⠠⠠⠤⠦⠖⠒⠐⠐⠒⠓⠋⠉⠈⠈
func WithSymbolsDots5() Option { return WithSpinnerStyle(SpinnerDots5) }

// WithSymbolsDots6 returns an option function that sets the spinner unicode
// animation with dots.
//
//	⢀⠀ ⡀⠀ ⠄⠀ ⢂⠀ ⡂⠀ ⠅⠀ ⢃⠀ ⡃⠀ ⠍⠀ ⢋⠀ ⡋⠀ ⠍⠁ ⢋⠁ ⡋⠁ ⠍⠉ ⠋⠉ ⠋⠉ ⠉⠙ ⠉⠙ ⠉⠩ ⠈⢙ ⠈⡙ ⢈⠩ ⡀⢙ ⠄⡙ ⢂⠩
func WithSymbolsDots6() Option { return WithSpinnerStyle(SpinnerDots6) }

// WithSymbolsDots7 returns an option function that sets the spinner unicode
// animation with dots.
//
//	⠁ ⠂ ⠃ ⠄ ⠅ ⠆ ⠇ ⡀ ⡁ ⡂ ⡃ ⡄ ⡅ ⡆ ⡇ ⠈ ⠉ ⠊ ⠋ ⠌ ⠍ ⠎ ⠏ ⡈ ⡉ ⡊ ⡋ ⡌ ⡍ ⡎ ⡏
func WithSymbolsDots7() Option { return WithSpinnerStyle(SpinnerDots7) }

// WithSymbolsLines returns an option function that sets the spinner unicode
// animation with lines.
//
//	⠂ - – — – -
func WithSymbolsLines() Option { return WithSpinnerStyle(SpinnerLines) }

// WithSymbolsWave returns an option function that sets the spinner unicode
// animation with wave patterns.
//
//	⢄ ⢂ ⢁ ⡀ ⠈ ⠘ ⠸
func WithSymbolsWave() Option { return WithSpinnerStyle(SpinnerWave) }

// WithSymbolsGrow returns an option function that sets the spinner unicode
// animation with growing bars.
//
//	▉ ▊ ▋ ▌ ▍ ▎ ▏
func WithSymbolsGrow() Option { return WithSpinnerStyle(SpinnerGrow) }

// WithSymbolsGrowVert returns an option function that sets the spinner unicode
// animation with growing bars.
//
//	▁ ▃ ▄ ▅ ▆ ▇ ▆ ▅ ▄ ▃
func WithSymbolsGrowVert() Option { return WithSpinnerStyle(SpinnerGrowvert) }

// WithSymbolsMoon returns an option function that sets the spinner unicode
// animation with moon phases.
//
//	🌑 🌒 🌓 🌔 🌕 🌖 🌗 🌘
func WithSymbolsMoon() Option { return WithSpinnerStyle(SpinnerMoon) }

// WithSymbolsPipe returns an option function that sets the spinner unicode
// animation with pipe characters.
//
//	| / - \\
func WithSymbolsPipe() Option { return WithSpinnerStyle(SpinnerPipe) }

// WithSymbolsPipe2 returns an option function that sets the spinner unicode
// animation with pipe characters.
//
//	┤ ┘ ┴ └ ├ ┌ ┬ ┐
func WithSymbolsPipe2() Option { return WithSpinnerStyle(SpinnerPipe2) }

// WithSymbolsSquare returns an option function that sets the spinner unicode
// animation with square segments.
//
//	▖ ▘ ▝ ▗
func WithSymbolsSquare() Option { return WithSpinnerStyle(SpinnerSquare) }

// WithSymbolsSquare2 returns an option function that sets the spinner unicode
// animation with square segments.
//
//	    
func WithSymbolsSquare2() Option { return WithSpinnerStyle(SpinnerSquare2) }

// WithSymbolsClock returns an option function that sets the spinner unicode
// animation with clock symbols.
//
//	🕛 🕐 🕑 🕒 🕓 🕔 🕕 🕖 🕗 🕘 🕙 🕚
func WithSymbolsClock() Option { return WithSpinnerStyle(SpinnerClock) }

// WithSymbolsDiamond returns an option function that sets the spinner unicode
// animation with diamond symbols.
//
//	◇ ◈ ⬟ ⬞
func WithSymbolsDiamond() Option { return WithSpinnerStyle(SpinnerDiamond) }

// WithSymbolsDiamond2 returns an option function that sets the spinner unicode
// animation with diamond symbols.
//
//	   
func WithSymbolsDiamond2() Option { return WithSpinnerStyle(SpinnerDiamond2) }

// WithSymbolsPlusCross returns an option function that sets the spinner unicode
// animation with plus and cross symbols.
//
//   - x
func WithSymbolsPlusCross() Option { return WithSpinnerStyle(SpinnerPluscross) }

// WithSymbolsArrows returns an option function that sets the spinner unicode
// animation with arrows.
//
//	< << <<< - > >> >>>
func WithSymbolsArrows() Option { return WithSpinnerStyle(SpinnerArrow) }

// WithSymbolsArrows2 returns an option function that sets the spinner unicode
// animation with arrows.
//
//	>    >>   >>>  >>>>
func WithSymbolsArrows2() Option { return WithSpinnerStyle(SpinnerArrow2) }

// WithSymbolsArrows3 returns an option function that sets the spinner unicode
// animation with arrows.
//
//	▹▹▹▹▹ ▸▹▹▹▹ ▹▸▹▹▹ ▹▹▸▹▹ ▹▹▹▸▹ ▹▹▹▹▸
func WithSymbolsArrows3() Option { return WithSpinnerStyle(SpinnerArrow3) }

// WithSymbolsArrows4 returns an option function that sets the spinner unicode
// animation with arrows.
//
//	← ↖ ↑ ↗ → ↘ ↓ ↙
func WithSymbolsArrows4() Option { return WithSpinnerStyle(SpinnerArrow4) }

// WithSymbolsCircles returns an option function that sets the spinner unicode
// animation with circles
//
//	o O @ *
func WithSymbolsCircles() Option { return WithSpinnerStyle(SpinnerCircle) }

// WithSymbolsCircles2 returns an option function that sets the spinner unicode
// animation with circles.
//
//	. o O ° O o .
func WithSymbolsCircles2() Option { return WithSpinnerStyle(SpinnerCircle2) }

// WithSymbolsCircles3 returns an option function that sets the spinner unicode
// animation with circles.
//
//	● ● ● ●
func WithSymbolsCircles3() Option { return WithSpinnerStyle(SpinnerCircle3) }

// WithSymbolsCircles4 returns an option function that sets the spinner unicode
// animation with circles.
//
//	   
func WithSymbolsCircles4() Option { return WithSpinnerStyle(SpinnerCircle4) }

// WithSymbolsCircles5 returns an option function that sets the spinner unicode
// animation with circles.
//
//	   
func WithSymbolsCircles5() Option { return WithSpinnerStyle(SpinnerCircle5) }

// WithSymbolsCircles6 returns an option function that sets the spinner unicode
// animation with circles.
//
//	    
func WithSymbolsCircles6() Option { return WithSpinnerStyle(SpinnerCircle6) }

// WithSymbolsCircles7 returns an option function that sets the spinner unicode
// animation with circles.
//
//	     
func WithSymbolsCircles7() Option { return WithSpinnerStyle(SpinnerCircle7) }

// WithSymbolsBounce returns an option function that sets the spinner unicode
// animation with circles.
//
//	[    ] [=   ] [==  ] [=== ] [ ===] [  ==] [   =]
func WithSymbolsBounce() Option { return WithSpinnerStyle(SpinnerBounce) }

// WithSymbolsBounceBall returns an option function that sets the spinner unicode
// animation with circles.
//
//	( ●    ) (  ●   ) (   ●  ) (    ● ) (     ●)
func WithSymbolsBounceBall() Option { return WithSpinnerStyle(SpinnerBounceball) }

// WithSymbolsToggle returns an option function that sets the spinner unicode
// animation with toggle symbols.
//
//	■ □ ▪ ▫
func WithSymbolsToggle() Option { return WithSpinnerStyle(SpinnerToggle) }

// WithSymbolsToggle2 returns an option function that sets the spinner unicode
// animation with toggle symbols.
//
//	= * -
func WithSymbolsToggle2() Option { return WithSpinnerStyle(SpinnerToggle2) }

// WithSymbolsToggle3 returns an option function that sets the spinner unicode
// animation with toggle symbols.
//
//	◉ ◎
func WithSymbolsToggle3() Option { return WithSpinnerStyle(SpinnerToggle3) }

// WithSymbolsLoading returns an option function that sets the spinner unicode
// animation with loading symbols.
//
//	loading....
func WithSymbolsLoading() Option { return WithSpinnerStyle(SpinnerLoading) }

// WithSymbolsTriangles returns an option function that sets the spinner
// unicode animation with rotating triangles.
//
//	▲ ▶ ▼ ◀
func WithSymbolsTriangles() Option { return WithSpinnerStyle(SpinnerTriangles) }

// WithSymbolsCubes returns an option function that sets the spinner
// unicode animation with cube rotation.
//
//	▖ ▘ ▝ ▗
func WithSymbolsCubes() Option { return WithSpinnerStyle(SpinnerCubes) }

// WithSymbolsThinking returns an option function that sets the spinner
// animation with growing question marks.
//
//	? ?? ??? ???? ?????
func WithSymbolsThinking() Option { return WithSpinnerStyle(SpinnerQuestion) }

// WithSymbolsPingPong returns an option function that sets the spinner
// animation with expanding and contracting brackets.
//
//	<     > <    > <   > <  > < > <>< < > <  > <   > <    >
func WithSymbolsPingPong() Option { return WithSpinnerStyle(SpinnerPingpong) }

// WithSymbolsPingPong2 returns an option function that sets the spinner
// animation with expanding and contracting brackets.
func WithSymbolsPingPong2() Option { return WithSpinnerStyle(SpinnerPingpong2) }

// WithSymbolsMatrix returns an option function that sets the spinner
// unicode animation with matrix-style loading.
//
//	╔═══╗ ║▓▓▓║ ║░▓▓║ ║░░▓║ ║░░░║ ╚═══╝
func WithSymbolsMatrix() Option { return WithSpinnerStyle(SpinnerMatrix) }

// WithSymbolsHex returns an option function that sets the spinner
// animation with hexadecimal counting.
//
//	0x0 0x1 0x2 0x3 0x4 0x5 0x6 0x7 0x8 0x9 0xA 0xB 0xC 0xD 0xE 0xF
func WithSymbolsHex() Option { return WithSpinnerStyle(SpinnerHexsymbols) }

// WithSymbolsPacman returns an option function that sets the spinner
// unicode animation with pacman movement.
//
//	󰮯··· ·󰮯·· ··󰮯· ···󰮯
func WithSymbolsPacman() Option { return WithSpinnerStyle(SpinnerPacman) }

// WithSymbolsBoxFill returns an option function that sets the spinner
// animation with progressively filling box.
//
//	[          ] [■         ] [■■        ] [■■■       ] [■■■■      ]
//	[■■■■■ ] [■■■■■■    ] [■■■■■■■   ] [■■■■■■■■  ] [■■■■■■■■■ ] [■■■■■■■■■■]
func WithSymbolsBoxFill() Option { return WithSpinnerStyle(SpinnerBoxFill) }

// WithSymbolsBoxFillShort returns an option function that sets the spinner
// animation with progressively filling box.
//
//	[      ] [■     ] [■■    ] [■■■   ] [■■■■  ] [■■■■■ ] [■■■■■■]
func WithSymbolsBoxFillShort() Option { return WithSpinnerStyle(SpinnerBoxFillshort) }

func WithSymbolsBoxBounce() Option { return WithSpinnerStyle(SpinnerBoxBounce) }

// WithSymbolsSnail returns an option function that sets the spinner
// animation with growing snail trail.
//
//	@ @- @-- @--- @---- @-----
func WithSymbolsSnail() Option { return WithSpinnerStyle(SpinnerSnail) }

// WithSymbolsWorm returns an option function that sets the spinner
// animation with growing and shrinking worm.
//
//	~ ~~ ~~~ ~~~~ ~~~~~ ~~~~ ~~~ ~~ ~
func WithSymbolsWorm() Option { return WithSpinnerStyle(SpinnerWorm) }

// WithSymbolsWorm2 returns an option function that sets the spinner
// animation with growing and shrinking worm.
//
//	~ ~~ ~~~ ~~~~ ~~~~~ ~~~~ ~~~ ~~ ~
func WithSymbolsWorm2() Option { return WithSpinnerStyle(SpinnerWorm2) }

// WithSymbolsMathOps returns an option function that sets the spinner
// unicode animation with mathematical operators.
//
//   - - × ÷ = ≠ ≈ ≤ ≥
func WithSymbolsMathOps() Option { return WithSpinnerStyle(SpinnerMathops) }

// WithSymbolsGreek returns an option function that sets the spinner
// unicode animation with Greek letters.
//
//	α β γ δ ε ζ η θ
func WithSymbolsGreek() Option { return WithSpinnerStyle(SpinnerGreek) }

// WithSymbolsCorners returns an option function that sets the spinner
// unicode animation with box corners.
//
//	┌ ┐ └ ┘
func WithSymbolsCorners() Option { return WithSpinnerStyle(SpinnerCorners) }

// WithSymbolsSlash returns an option function that sets the spinner
// animation with increasing forward slashes.
//
//	/ // /// //// /////
func WithSymbolsSlash() Option { return WithSpinnerStyle(SpinnerSlash) }

// WithSymbolsBackslash returns an option function that sets the spinner
// animation with increasing backslashes.
//
//	\ \\ \\\ \\\\ \\\\\
func WithSymbolsBackslash() Option { return WithSpinnerStyle(SpinnerBackslash) }

// WithSymbolsMarquee returns an option function that sets the spinner
// animation with moving marquee arrow.
//
//	[          ] [ >        ] [  >       ] [   >      ] [    >     ]
//	[     >    ] [      >   ] [       >  ] [        > ] [         >]
func WithSymbolsMarquee() Option { return WithSpinnerStyle(SpinnerMarquee) }

// WithSymbolsFade returns an option function that sets the spinner
// unicode animation with fading block.
//
//	█ ▓ ▒ ░   ░ ▒ ▓
func WithSymbolsFade() Option { return WithSpinnerStyle(SpinnerFade) }

// WithSymbolsMath returns an option function that sets the spinner
// unicode animation with mathematical symbols.
//
//	∀ ∃ ∈ ∉ ∋ ∌ ⊆ ⊂ ⊄ ⊇ ⊃ ⊅
func WithSymbolsMath() Option { return WithSpinnerStyle(SpinnerLogicsymbols) }

// WithSymbolsCurrency returns an option function that sets the spinner
// unicode animation with currency symbols.
//
//	$ € £ ¥ ₿ ₹
func WithSymbolsCurrency() Option { return WithSpinnerStyle(SpinnerCurrency) }

// WithSymbolsGeometric returns an option function that sets the spinner
// unicode animation with geometric shapes.
//
//	△ ◊ ◈ ◇ ○ ● ◐ ◑ ◒ ◓
func WithSymbolsGeometric() Option { return WithSpinnerStyle(SpinnerGeometric) }

// WithSymbolsRunner returns an option function that sets the spinner unicode animation.
//
//	▁▁▁▁▁ ▂▁▁▁▁ ▃▂▁▁▁ ▄▃▂▁▁ ▅▄▃▂▁ ▆▅▄▃▂ ▇▆▅▄▃ █▇▆▅▄
func WithSymbolsRunner() Option { return WithSpinnerStyle(SpinnerRunner) }

// WithSymbolsCursorBlink returns a option function that sets the spinner
// unicode animation
//
//	"_", " ", "_", " "
func WithSymbolsCursorBlink() Option { return WithSpinnerStyle(SpinnerCursorBlink) }

// WithSymbolsEllipsis returns a option function that sets the spinner
// unicode animation
//
//	".  " ".. " "..." " .." "  ." "   "
func WithSymbolsEllipsis() Option { return WithSpinnerStyle(SpinnerEllipsis) }

// WithSymbolsBrailleWave returns a option function that sets the spinner
// unicode animation
//
//	⡀ ⡄ ⡆ ⡇ ⡏ ⡟ ⡿ ⣿ ⡿ ⡟ ⡏ ⡇ ⡆ ⡄
func WithSymbolsBrailleWave() Option { return WithSpinnerStyle(SpinnerBrailleWave) }

// WithSymbolsOrbit returns a option function that sets the spinner
// unicode animation
//
//	◐ ◓ ◑ ◒
func WithSymbolsOrbit() Option { return WithSpinnerStyle(SpinnerOrbit) }

// WithSymbolsSweep returns a option function that sets the spinner
// unicode animation
//
//	←──── ─←─── ──←── ───←─ ────← ───→─ ──→── ─→─── →────
func WithSymbolsSweep() Option { return WithSpinnerStyle(SpinnerSweep) }

// WithSymbolsPulse returns a option function that sets the spinner
// unicode animation
//
//	░ ▒ ▓ █ ▓ ▒
func WithSymbolsPulse() Option { return WithSpinnerStyle(SpinnerPulse) }

// WithSymbolsFlip returns a option function that sets the spinner
// unicode animation
//
//	_ _ _ - ` ` ' ´ - _ _ _
func WithSymbolsFlip() Option { return WithSpinnerStyle(SpinnerFlip) }

// WithSymbolsMaterial returns a option function that sets the spinner
// unicode animation.
func WithSymbolsMaterial() Option { return WithSpinnerStyle(SpinnerMaterial) }

// WithSymbolsShark returns a option function that sets the spinner
// unicode animation
//
//	▐________|\\____▌
func WithSymbolsShark() Option { return WithSpinnerStyle(SpinnerShark) }

// WithSymbolsBetawave returns a option function that sets the spinner
// unicode animation
//
//	ρββββββ βρβββββ ββρββββ βββρβββ ββββρββ βββββρβ ββββββρ
func WithSymbolsBetawave() Option { return WithSpinnerStyle(SpinnerBetawave) }

// WithSymbolsFistbump returns a option function that sets the spinner
// unicode animation
//
//	🤜✨🤛
func WithSymbolsFistbump() Option { return WithSpinnerStyle(SpinnerFistbump) }

// WithSymbolsFutbol returns a option function that sets the spinner
// unicode animation
//
//	🧑   ⚽️     🧑
func WithSymbolsFutbol() Option { return WithSpinnerStyle(SpinnerFutbolHead) }

// WithSymbolsMindBlown returns a option function that sets the spinner
// unicode animation
//
//	😐  😐  😮  😮  😦  😦  😧  😧  🤯  💥  ✨
func WithSymbolsMindBlown() Option { return WithSpinnerStyle(SpinnerMindblown) }

// WithSymbolsSpeaker returns a option function that sets the spinner
// unicode animation
//
//	🔈 🔉  🔊  🔉
func WithSymbolsSpeaker() Option { return WithSpinnerStyle(SpinnerSpeaker) }

// WithSymbolsStar returns a option function that sets the spinner
// unicode animation
//
//	🔈 🔉  🔊  🔉
func WithSymbolsStar() Option { return WithSpinnerStyle(SpinnerStar) }
