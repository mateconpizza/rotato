package rotato

type SpinnerGroup string

const (
	GroupArrows   SpinnerGroup = "arrows"
	GroupBlocks   SpinnerGroup = "blocks"
	GroupBraille  SpinnerGroup = "braille"
	GroupCircular SpinnerGroup = "circular"
	GroupEffects  SpinnerGroup = "effects"
	GroupFramed   SpinnerGroup = "framed"
	GroupFun      SpinnerGroup = "fun"
	GroupLines    SpinnerGroup = "lines"
	GroupMinimal  SpinnerGroup = "minimal"
	GroupMotion   SpinnerGroup = "motion"
	GroupShapes   SpinnerGroup = "shapes"
	GroupSymbols  SpinnerGroup = "symbols"
	GroupText     SpinnerGroup = "text"
)

type SpinnerStyle struct {
	Name    string
	Frames  []string
	Group   SpinnerGroup
	OptName string
}

var Groups = []SpinnerGroup{
	GroupArrows,
	GroupBlocks,
	GroupBraille,
	GroupCircular,
	GroupEffects,
	GroupFramed,
	GroupFun,
	GroupLines,
	GroupMinimal,
	GroupMotion,
	GroupShapes,
	GroupSymbols,
	GroupText,
}

var (
	// braille-style spinners.
	defaultSymbols = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	brailleWave    = []string{"⡀", "⡄", "⡆", "⡇", "⡏", "⡟", "⡿", "⣿", "⡿", "⡟", "⡏", "⡇", "⡆", "⡄"}
	dots           = []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
	dots3          = []string{"⠄", "⠆", "⠇", "⠋", "⠙", "⠸", "⠰", "⠠", "⠰", "⠸", "⠙", "⠋", "⠇", "⠆"}
	dots4          = []string{"⠁", "⠃", "⠇", "⠧", "⠷", "⠿", "⠷", "⠧", "⠇", "⠃"}
	dots5          = []string{"⠁", "⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈", "⠈"}

	// arrow and directional spinners.
	arrow  = []string{"<", "<<", "<<<", "-", ">", ">>", ">>>"}
	arrow2 = []string{">   ", ">>  ", ">>> ", ">>>>"}
	arrow3 = []string{"▹▹▹▹▹", "▸▹▹▹▹", "▹▸▹▹▹", "▹▹▸▹▹", "▹▹▹▸▹", "▹▹▹▹▸"}
	arrow4 = []string{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"}
	sweep  = []string{"←────", "─←───", "──←──", "───←─", "────←", "───→─", "──→──", "─→───", "→────"}

	// line and pipe spinners.
	slash     = []string{"/    ", "//   ", "///  ", "//// ", "/////"}
	backslash = []string{"\\    ", "\\\\   ", "\\\\\\  ", "\\\\\\\\ ", "\\\\\\\\\\"}
	pipe      = []string{"|", "/", "-", "\\"}
	pipe2     = []string{"┤", "┘", "┴", "└", "├", "┌", "┬", "┐"}
	lines     = []string{"⠂", "-", "–", "—", "–", "-"}

	// block and progress bar spinners.
	block          = []string{"░", "▒", "▒", "░", "▓"}
	blockbar       = []string{"█▒▒▒▒▒▒▒▒▒", "███▒▒▒▒▒▒▒", "█████▒▒▒▒▒", "███████▒▒▒", "██████████"}
	blockbar2      = []string{"[       ]", "[|      ]", "[||     ]", "[|||    ]", "[||||   ]", "[|||||  ]", "[|||||| ]", "[|||||||]"}
	blockbar3      = []string{"[       ]", "[=      ]", "[==     ]", "[===    ]", "[====   ]", "[=====  ]", "[====== ]", "[=======]"}
	blockbar4      = []string{"|", "||", "|||", "||||", "|||||", "||||||", "|||||||", "||||||||", "|||||||", "||||||", "|||||", "||||", "|||", "||", "|"}
	blockbar5      = []string{"[*-------]", "[-*------]", "[--*-----]", "[---*----]", "[----*---]", "[-----*--]", "[------*-]", "[-------*]", "[------*-]", "[-----*--]", "[----*---]", "[---*----]", "[--*-----]", "[-*------]", "[*-------]"}
	blockbar6      = []string{"·-----", "-·----", "--·---", "---·--", "----·-", "-----·", "----·-", "---·--", "--·---", "-·----", "·-----"}
	blockbarpretty = []string{"", "", "", "", "", "", ""}
	boxfill        = []string{"[          ]", "[■         ]", "[■■        ]", "[■■■       ]", "[■■■■      ]", "[■■■■■     ]", "[■■■■■■    ]", "[■■■■■■■   ]", "[■■■■■■■■  ]", "[■■■■■■■■■ ]", "[■■■■■■■■■■]"}
	boxfillshort   = []string{"[      ]", "[■     ]", "[■■    ]", "[■■■   ]", "[■■■■  ]", "[■■■■■ ]", "[■■■■■■]"}

	// bouncing and motion spinners.
	bounce     = []string{"[    ]", "[=   ]", "[==  ]", "[=== ]", "[ ===]", "[  ==]", "[   =]", "[    ]", "[   =]", "[  ==]", "[ ===]", "[====]", "[=== ]", "[==  ]", "[=   ]"}
	bounceball = []string{"( ●    )", "(  ●   )", "(   ●  )", "(    ● )", "(     ●)", "(    ● )", "(   ●  )", "(  ●   )", "( ●    )", "(●     )"}
	pingpong   = []string{"<     >", "<    >", "<   >", "<  >", "< >", "<><", "< >", "<  >", "<   >", "<    >"}
	runner     = []string{"▁▁▁▁▁", "▂▁▁▁▁", "▃▂▁▁▁", "▄▃▂▁▁", "▅▄▃▂▁", "▆▅▄▃▂", "▇▆▅▄▃", "█▇▆▅▄"}

	// circular and orbit spinners.
	circle  = []string{"o", "O", "@", "*"}
	circle2 = []string{".", "o", "O", "°", "O", "o", "."}
	circle3 = []string{"●", "●", "●", "●"}
	circle4 = []string{"", "", "", ""}
	circle5 = []string{"", "", "", "", ""}
	circle6 = []string{"", "", "", "", "", ""}
	circle7 = []string{"", "", "", "", "", ""}
	orbit   = []string{"◐", "◓", "◑", "◒"}
	moon    = []string{"🌑", "🌒", "🌓", "🌔", "🌕", "🌖", "🌗", "🌘"}
	clock   = []string{"🕛", "🕐", "🕑", "🕒", "🕓", "🕔", "🕕", "🕖", "🕗", "🕘", "🕙", "🕚"}

	// shape and geometric spinners.
	square    = []string{"▖", "▘", "▝", "▗"}
	square2   = []string{"", "", "", "", ""}
	cubes     = []string{"▖", "▘", "▝", "▗"}
	triangles = []string{"▲", "▶", "▼", "◀"}
	diamond   = []string{"◇", "◈", "⬟", "⬞"}
	diamond2  = []string{"", "", "", ""}
	geometric = []string{"△", "◊", "◈", "◇", "○", "●", "◐", "◑", "◒", "◓"}

	// text and semantic spinners.
	loading    = []string{"l      ", "lo     ", "loa    ", "load   ", "loadi  ", "loadin ", "loading", "loading.", "loading..", "loading...", "loading...."}
	ellipsis   = []string{".  ", ".. ", "...", " ..", "  .", "   "}
	question   = []string{"?", "??", "???", "????", "?????"}
	hexsymbols = []string{"0x0", "0x1", "0x2", "0x3", "0x4", "0x5", "0x6", "0x7", "0x8", "0x9", "0xA", "0xB", "0xC", "0xD", "0xE", "0xF"}
	// symbol and themed spinners.
	currency     = []string{"$", "€", "£", "¥", "₿", "₹"}
	mathops      = []string{"+", "-", "×", "÷", "=", "≠", "≈", "≤", "≥"}
	logicsymbols = []string{"∀", "∃", "∈", "∉", "∋", "∌", "⊆", "⊂", "⊄", "⊇", "⊃", "⊅"}
	greek        = []string{"α", "β", "γ", "δ", "ε", "ζ", "η", "θ"}

	// fun and novelty spinners.
	pacman = []string{"󰮯···", "·󰮯··", "··󰮯·", "···󰮯"}
	snail  = []string{"@     ", "@-    ", "@--   ", "@---  ", "@---- ", "@-----"}
	worm   = []string{"~", "~~", "~~~", "~~~~", "~~~~~", "~~~~", "~~~", "~~", "~"}
	worm2  = []string{"~    ", "~~   ", "~~~  ", "~~~~ ", "~~~~~", "~~~~ ", "~~~  ", "~~   ", "~    "}

	// minimal and toggle spinners.
	toggle      = []string{"■", "□", "▪", "▫"}
	toggle2     = []string{"=", "*", "-"}
	toggle3     = []string{"◉", "◎"}
	cursorBlink = []string{"_", " ", "_", " "}
	pluscross   = []string{"+", "x"}

	// visual effects spinners.
	fade     = []string{"█", "▓", "▒", "░", " ", "░", "▒", "▓"}
	pulse    = []string{"░", "▒", "▓", "█", "▓", "▒"}
	grow     = []string{"▉", "▊", "▋", "▌", "▍", "▎", "▏"}
	growvert = []string{"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▂"}
	wave     = []string{"⢄", "⢂", "⢁", "⡀", "⠈", "⠘", "⠸"}

	// framed and container spinners.
	marquee = []string{"[          ]", "[ >        ]", "[  >       ]", "[   >      ]", "[    >     ]", "[     >    ]", "[      >   ]", "[       >  ]", "[        > ]", "[         >]"}
	matrix  = []string{"╔═══╗", "║▓▓▓║", "║░▓▓║", "║░░▓║", "║░░░║", "╚═══╝"}
	corners = []string{"┌", "┐", "└", "┘"}
)

var registry = []SpinnerStyle{
	// braille
	{"default", defaultSymbols, GroupBraille, "WithSymbolsDefault"},
	{"brailleWave", brailleWave, GroupBraille, "WithSymbolsBrailleWave"},
	{"dots", dots, GroupBraille, "WithSymbolsDots"},
	{"dots3", dots3, GroupBraille, "WithSymbolsDots3"},
	{"dots4", dots4, GroupBraille, "WithSymbolsDots4"},
	{"dots5", dots5, GroupBraille, "WithSymbolsDots5"},

	// arrows
	{"arrow", arrow, GroupArrows, "WithSymbolsArrows"},
	{"arrow2", arrow2, GroupArrows, "WithSymbolsArrows2"},
	{"arrow3", arrow3, GroupArrows, "WithSymbolsArrows3"},
	{"arrow4", arrow4, GroupArrows, "WithSymbolsArrows4"},
	{"sweep", sweep, GroupArrows, "WithSymbolsSweep"},

	// lines
	{"slash", slash, GroupLines, "WithSymbolsSlash"},
	{"backslash", backslash, GroupLines, "WithSymbolsBackslash"},
	{"pipe", pipe, GroupLines, "WithSymbolsPipe"},
	{"pipe2", pipe2, GroupLines, "WithSymbolsPipe2"},
	{"lines", lines, GroupLines, "WithSymbolsLines"},

	// blocks
	{"block", block, GroupBlocks, "WithSymbolsBlock"},
	{"blockbar", blockbar, GroupBlocks, "WithSymbolsBarBlock"},
	{"blockbar2", blockbar2, GroupBlocks, "WithSymbolsBarBlock2"},
	{"blockbar3", blockbar3, GroupBlocks, "WithSymbolsBarBlock3"},
	{"blockbar4", blockbar4, GroupBlocks, "WithSymbolsBarBlock4"},
	{"blockbar5", blockbar5, GroupBlocks, "WithSymbolsBarBlock5"},
	{"blockbar6", blockbar6, GroupBlocks, "WithSymbolsBarBlock6"},
	{"blockbarpretty", blockbarpretty, GroupBlocks, "WithSymbolsBlockPretty"},
	{"boxfill", boxfill, GroupBlocks, "WithSymbolsBoxFill"},
	{"boxfillshort", boxfillshort, GroupBlocks, "WithSymbolsBoxFill"},

	// motion
	{"bounce", bounce, GroupMotion, "WithSymbolsBounce"},
	{"bounceball", bounceball, GroupMotion, "WithSymbolsBounceBall"},
	{"pingpong", pingpong, GroupMotion, "WithSymbolsPingPong"},
	{"runner", runner, GroupMotion, "WithSymbolsRunner"},

	// circular
	{"circle", circle, GroupCircular, "WithSymbolsCircles"},
	{"circle2", circle2, GroupCircular, "WithSymbolsCircles2"},
	{"circle3", circle3, GroupCircular, "WithSymbolsCircles3"},
	{"circle4", circle4, GroupCircular, "WithSymbolsCircles4"},
	{"circle5", circle5, GroupCircular, "WithSymbolsCircles5"},
	{"circle6", circle6, GroupCircular, "WithSymbolsCircles6"},
	{"circle7", circle7, GroupCircular, "WithSymbolsCircles7"},
	{"orbit", orbit, GroupCircular, "WithSymbolsOrbit"},
	{"moon", moon, GroupCircular, "WithSymbolsMoon"},
	{"clock", clock, GroupCircular, "WithSymbolsClock"},

	// shapes
	{"square", square, GroupShapes, "WithSymbolsSquare"},
	{"square2", square2, GroupShapes, "WithSymbolsSquare2"},
	{"cubes", cubes, GroupShapes, "WithSymbolsCubes"},
	{"triangles", triangles, GroupShapes, "WithSymbolsTriangles"},
	{"diamond", diamond, GroupShapes, "WithSymbolsDiamond"},
	{"diamond2", diamond2, GroupShapes, "WithSymbolsDiamond2"},
	{"geometric", geometric, GroupShapes, "WithSymbolsGeometric"},

	// text
	{"loading", loading, GroupText, "WithSymbolsLoading"},
	{"ellipsis", ellipsis, GroupText, "WithSymbolsEllipsis"},
	{"question", question, GroupText, "WithSymbolsThinking"},
	{"hexsymbols", hexsymbols, GroupText, "WithSymbolsHex"},

	// symbols
	{"currency", currency, GroupSymbols, "WithSymbolsCurrency"},
	{"mathops", mathops, GroupSymbols, "WithSymbolsMathOps"},
	{"logicsymbols", logicsymbols, GroupSymbols, "WithSymbolsMath"},
	{"greek", greek, GroupSymbols, "WithSymbolsGreek"},

	// fun
	{"pacman", pacman, GroupFun, "WithSymbolsPacman"},
	{"snail", snail, GroupFun, "WithSymbolsSnail"},
	{"worm", worm, GroupFun, "WithSymbolsWorm"},
	{"worm2", worm2, GroupFun, "WithSymbolsWorm2"},

	// minimal
	{"toggle", toggle, GroupMinimal, "WithSymbolsToggle"},
	{"toggle2", toggle2, GroupMinimal, "WithSymbolsToggle2"},
	{"toggle3", toggle3, GroupMinimal, "WithSymbolsToggle3"},
	{"cursorBlink", cursorBlink, GroupMinimal, "WithSymbolsCursorBlink"},
	{"pluscross", pluscross, GroupMinimal, "WithSymbolsPlusCross"},

	// effects
	{"fade", fade, GroupEffects, "WithSymbolsFade"},
	{"pulse", pulse, GroupEffects, "WithSymbolsPulse"},
	{"grow", grow, GroupEffects, "WithSymbolsGrow"},
	{"growvert", growvert, GroupEffects, "WithSymbolsGrowVert"},
	{"wave", wave, GroupEffects, "WithSymbolsWave"},

	// framed
	{"marquee", marquee, GroupFramed, "WithSymbolsMarquee"},
	{"matrix", matrix, GroupFramed, "WithSymbolsMatrix"},
	{"corners", corners, GroupFramed, "WithSymbolsCorners"},
}
