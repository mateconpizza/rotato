package rotato

type SpinnerName string

const (
	SpinnerDefault     SpinnerName = "default"
	SpinnerBrailleWave SpinnerName = "brailleWave"
	SpinnerDots        SpinnerName = "dots"
	SpinnerDots3       SpinnerName = "dots3"
	SpinnerDots4       SpinnerName = "dots4"
	SpinnerDots5       SpinnerName = "dots5"
	SpinnerDots6       SpinnerName = "dots6"
	SpinnerDots7       SpinnerName = "dots7"

	SpinnerArrow  SpinnerName = "arrow"
	SpinnerArrow2 SpinnerName = "arrow2"
	SpinnerArrow3 SpinnerName = "arrow3"
	SpinnerArrow4 SpinnerName = "arrow4"
	SpinnerSweep  SpinnerName = "sweep"

	SpinnerSlash     SpinnerName = "slash"
	SpinnerBackslash SpinnerName = "backslash"
	SpinnerPipe      SpinnerName = "pipe"
	SpinnerPipe2     SpinnerName = "pipe2"
	SpinnerLines     SpinnerName = "lines"

	SpinnerBlock          SpinnerName = "block"
	SpinnerBlockbar       SpinnerName = "blockbar"
	SpinnerBlockbar2      SpinnerName = "blockbar2"
	SpinnerBlockbar3      SpinnerName = "blockbar3"
	SpinnerBlockbar4      SpinnerName = "blockbar4"
	SpinnerBlockbar5      SpinnerName = "blockbar5"
	SpinnerBlockbar6      SpinnerName = "blockbar6"
	SpinnerBlockbar7      SpinnerName = "blockbar7"
	SpinnerBlockbarPretty SpinnerName = "blockbarpretty"
	SpinnerBoxFill        SpinnerName = "boxfill"
	SpinnerBoxFillshort   SpinnerName = "boxfillshort"
	SpinnerBoxBounce      SpinnerName = "boxbounce"

	SpinnerBounce     SpinnerName = "bounce"
	SpinnerBounceball SpinnerName = "bounceball"
	SpinnerPingpong   SpinnerName = "pingpong"
	SpinnerPingpong2  SpinnerName = "pingpong2"
	SpinnerRunner     SpinnerName = "runner"

	SpinnerCircle  SpinnerName = "circle"
	SpinnerCircle2 SpinnerName = "circle2"
	SpinnerCircle3 SpinnerName = "circle3"
	SpinnerCircle4 SpinnerName = "circle4"
	SpinnerCircle5 SpinnerName = "circle5"
	SpinnerCircle6 SpinnerName = "circle6"
	SpinnerCircle7 SpinnerName = "circle7"
	SpinnerOrbit   SpinnerName = "orbit"
	SpinnerMoon    SpinnerName = "moon"
	SpinnerClock   SpinnerName = "clock"

	SpinnerSquare    SpinnerName = "square"
	SpinnerSquare2   SpinnerName = "square2"
	SpinnerCubes     SpinnerName = "cubes"
	SpinnerTriangles SpinnerName = "triangles"
	SpinnerDiamond   SpinnerName = "diamond"
	SpinnerDiamond2  SpinnerName = "diamond2"
	SpinnerGeometric SpinnerName = "geometric"

	SpinnerLoading    SpinnerName = "loading"
	SpinnerEllipsis   SpinnerName = "ellipsis"
	SpinnerQuestion   SpinnerName = "question"
	SpinnerHexsymbols SpinnerName = "hexsymbols"

	SpinnerCurrency     SpinnerName = "currency"
	SpinnerMathops      SpinnerName = "mathops"
	SpinnerLogicsymbols SpinnerName = "logicsymbols"
	SpinnerGreek        SpinnerName = "greek"

	SpinnerPacman SpinnerName = "pacman"
	SpinnerSnail  SpinnerName = "snail"
	SpinnerWorm   SpinnerName = "worm"
	SpinnerWorm2  SpinnerName = "worm2"

	SpinnerToggle      SpinnerName = "toggle"
	SpinnerToggle2     SpinnerName = "toggle2"
	SpinnerToggle3     SpinnerName = "toggle3"
	SpinnerCursorBlink SpinnerName = "cursorBlink"
	SpinnerPluscross   SpinnerName = "pluscross"

	SpinnerFade     SpinnerName = "fade"
	SpinnerPulse    SpinnerName = "pulse"
	SpinnerGrow     SpinnerName = "grow"
	SpinnerGrowvert SpinnerName = "growvert"
	SpinnerWave     SpinnerName = "wave"

	SpinnerMarquee SpinnerName = "marquee"
	SpinnerMatrix  SpinnerName = "matrix"
	SpinnerCorners SpinnerName = "corners"

	SpinnerFlip       SpinnerName = "flip"
	SpinnerMaterial   SpinnerName = "material"
	SpinnerShark      SpinnerName = "shark"
	SpinnerBetawave   SpinnerName = "betawave"
	SpinnerFistbump   SpinnerName = "fistbump"
	SpinnerFutbolHead SpinnerName = "futbolHead"
	SpinnerMindblown  SpinnerName = "mindblown"
	SpinnerSpeaker    SpinnerName = "speaker"
	SpinnerStar       SpinnerName = "star"
)

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
	GroupMisc     SpinnerGroup = "misc"
)

type SpinnerStyle struct {
	Name   SpinnerName
	Frames []string
	Group  SpinnerGroup
}

var byName = func() map[SpinnerName]SpinnerStyle {
	m := make(map[SpinnerName]SpinnerStyle, len(registry))

	for _, spinner := range registry {
		m[spinner.Name] = spinner
	}

	return m
}()

var Groups = func() []SpinnerGroup {
	seen := make(map[SpinnerGroup]struct{})
	groups := make([]SpinnerGroup, 0)

	for _, spinner := range registry {
		if _, ok := seen[spinner.Group]; ok {
			continue
		}

		seen[spinner.Group] = struct{}{}
		groups = append(groups, spinner.Group)
	}

	return groups
}()

var (
	// braille-style spinners.
	defaultSymbols = dots
	brailleWave    = []string{"⡀", "⡄", "⡆", "⡇", "⡏", "⡟", "⡿", "⣿", "⡿", "⡟", "⡏", "⡇", "⡆", "⡄"}
	dots           = []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
	dots3          = []string{"⠄", "⠆", "⠇", "⠋", "⠙", "⠸", "⠰", "⠠", "⠰", "⠸", "⠙", "⠋", "⠇", "⠆"}
	dots4          = []string{"⠁", "⠃", "⠇", "⠧", "⠷", "⠿", "⠷", "⠧", "⠇", "⠃"}
	dots5          = []string{"⠁", "⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈", "⠈"}
	dots6          = []string{
		"⢀⠀",
		"⡀⠀",
		"⠄⠀",
		"⢂⠀",
		"⡂⠀",
		"⠅⠀",
		"⢃⠀",
		"⡃⠀",
		"⠍⠀",
		"⢋⠀",
		"⡋⠀",
		"⠍⠁",
		"⢋⠁",
		"⡋⠁",
		"⠍⠉",
		"⠋⠉",
		"⠋⠉",
		"⠉⠙",
		"⠉⠙",
		"⠉⠩",
		"⠈⢙",
		"⠈⡙",
		"⢈⠩",
		"⡀⢙",
		"⠄⡙",
		"⢂⠩",
		"⡂⢘",
		"⠅⡘",
		"⢃⠨",
		"⡃⢐",
		"⠍⡐",
		"⢋⠠",
		"⡋⢀",
		"⠍⡁",
		"⢋⠁",
		"⡋⠁",
		"⠍⠉",
		"⠋⠉",
		"⠋⠉",
		"⠉⠙",
		"⠉⠙",
		"⠉⠩",
		"⠈⢙",
		"⠈⡙",
		"⠈⠩",
		"⠀⢙",
		"⠀⡙",
		"⠀⠩",
		"⠀⢘",
		"⠀⡘",
		"⠀⠨",
		"⠀⢐",
		"⠀⡐",
		"⠀⠠",
		"⠀⢀",
		"⠀⡀",
	}
	dots7 = []string{
		"⠀",
		"⠁",
		"⠂",
		"⠃",
		"⠄",
		"⠅",
		"⠆",
		"⠇",
		"⡀",
		"⡁",
		"⡂",
		"⡃",
		"⡄",
		"⡅",
		"⡆",
		"⡇",
		"⠈",
		"⠉",
		"⠊",
		"⠋",
		"⠌",
		"⠍",
		"⠎",
		"⠏",
		"⡈",
		"⡉",
		"⡊",
		"⡋",
		"⡌",
		"⡍",
		"⡎",
		"⡏",
		"⠐",
		"⠑",
		"⠒",
		"⠓",
		"⠔",
		"⠕",
		"⠖",
		"⠗",
		"⡐",
		"⡑",
		"⡒",
		"⡓",
		"⡔",
		"⡕",
		"⡖",
		"⡗",
		"⠘",
		"⠙",
		"⠚",
		"⠛",
		"⠜",
		"⠝",
		"⠞",
		"⠟",
		"⡘",
		"⡙",
		"⡚",
		"⡛",
		"⡜",
		"⡝",
		"⡞",
		"⡟",
		"⠠",
		"⠡",
		"⠢",
		"⠣",
		"⠤",
		"⠥",
		"⠦",
		"⠧",
		"⡠",
		"⡡",
		"⡢",
		"⡣",
		"⡤",
		"⡥",
		"⡦",
		"⡧",
		"⠨",
		"⠩",
		"⠪",
		"⠫",
		"⠬",
		"⠭",
		"⠮",
		"⠯",
		"⡨",
		"⡩",
		"⡪",
		"⡫",
		"⡬",
		"⡭",
		"⡮",
		"⡯",
		"⠰",
		"⠱",
		"⠲",
		"⠳",
		"⠴",
		"⠵",
		"⠶",
		"⠷",
		"⡰",
		"⡱",
		"⡲",
		"⡳",
		"⡴",
		"⡵",
		"⡶",
		"⡷",
		"⠸",
		"⠹",
		"⠺",
		"⠻",
		"⠼",
		"⠽",
		"⠾",
		"⠿",
		"⡸",
		"⡹",
		"⡺",
		"⡻",
		"⡼",
		"⡽",
		"⡾",
		"⡿",
		"⢀",
		"⢁",
		"⢂",
		"⢃",
		"⢄",
		"⢅",
		"⢆",
		"⢇",
		"⣀",
		"⣁",
		"⣂",
		"⣃",
		"⣄",
		"⣅",
		"⣆",
		"⣇",
		"⢈",
		"⢉",
		"⢊",
		"⢋",
		"⢌",
		"⢍",
		"⢎",
		"⢏",
		"⣈",
		"⣉",
		"⣊",
		"⣋",
		"⣌",
		"⣍",
		"⣎",
		"⣏",
		"⢐",
		"⢑",
		"⢒",
		"⢓",
		"⢔",
		"⢕",
		"⢖",
		"⢗",
		"⣐",
		"⣑",
		"⣒",
		"⣓",
		"⣔",
		"⣕",
		"⣖",
		"⣗",
		"⢘",
		"⢙",
		"⢚",
		"⢛",
		"⢜",
		"⢝",
		"⢞",
		"⢟",
		"⣘",
		"⣙",
		"⣚",
		"⣛",
		"⣜",
		"⣝",
		"⣞",
		"⣟",
		"⢠",
		"⢡",
		"⢢",
		"⢣",
		"⢤",
		"⢥",
		"⢦",
		"⢧",
		"⣠",
		"⣡",
		"⣢",
		"⣣",
		"⣤",
		"⣥",
		"⣦",
		"⣧",
		"⢨",
		"⢩",
		"⢪",
		"⢫",
		"⢬",
		"⢭",
		"⢮",
		"⢯",
		"⣨",
		"⣩",
		"⣪",
		"⣫",
		"⣬",
		"⣭",
		"⣮",
		"⣯",
		"⢰",
		"⢱",
		"⢲",
		"⢳",
		"⢴",
		"⢵",
		"⢶",
		"⢷",
		"⣰",
		"⣱",
		"⣲",
		"⣳",
		"⣴",
		"⣵",
		"⣶",
		"⣷",
		"⢸",
		"⢹",
		"⢺",
		"⢻",
		"⢼",
		"⢽",
		"⢾",
		"⢿",
		"⣸",
		"⣹",
		"⣺",
		"⣻",
		"⣼",
		"⣽",
		"⣾",
		"⣿",
	}

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
	blockbar7      = []string{"■     ", "■■    ", "■■■   ", "■■■■  ", "■■■■■ ", "■■■■■■"}
	blockbarpretty = []string{"", "", "", "", "", "", ""}
	boxfill        = []string{"[          ]", "[■         ]", "[■■        ]", "[■■■       ]", "[■■■■      ]", "[■■■■■     ]", "[■■■■■■    ]", "[■■■■■■■   ]", "[■■■■■■■■  ]", "[■■■■■■■■■ ]", "[■■■■■■■■■■]"}
	boxfillshort   = []string{"[      ]", "[■     ]", "[■■    ]", "[■■■   ]", "[■■■■  ]", "[■■■■■ ]", "[■■■■■■]"}
	boxBounce      = []string{"▌", "▀", "▐", "▄"}

	// bouncing and motion spinners.
	bounce     = []string{"[    ]", "[=   ]", "[==  ]", "[=== ]", "[ ===]", "[  ==]", "[   =]", "[    ]", "[   =]", "[  ==]", "[ ===]", "[====]", "[=== ]", "[==  ]", "[=   ]"}
	bounceball = []string{"( ●    )", "(  ●   )", "(   ●  )", "(    ● )", "(     ●)", "(    ● )", "(   ●  )", "(  ●   )", "( ●    )", "(●     )"}
	pingpong   = []string{"<     >", "<    >", "<   >", "<  >", "< >", "<><", "< >", "<  >", "<   >", "<    >"}
	pingpong2  = []string{"▐⠂       ▌", "▐⠈       ▌", "▐ ⠂      ▌", "▐ ⠠      ▌", "▐  ⡀     ▌", "▐  ⠠     ▌", "▐   ⠂    ▌", "▐   ⠈    ▌", "▐    ⠂   ▌", "▐    ⠠   ▌", "▐     ⡀  ▌", "▐     ⠠  ▌", "▐      ⠂ ▌", "▐      ⠈ ▌", "▐       ⠂▌", "▐       ⠠▌", "▐       ⡀▌", "▐      ⠠ ▌", "▐      ⠂ ▌", "▐     ⠈  ▌", "▐     ⠂  ▌", "▐    ⠠   ▌", "▐    ⡀   ▌", "▐   ⠠    ▌", "▐   ⠂    ▌", "▐  ⠈     ▌", "▐  ⠂     ▌", "▐ ⠠      ▌", "▐ ⡀      ▌", "▐⠠       ▌"}
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
	pluscross   = []string{"+", "x", "*"}

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

	flip     = []string{"_", "_", "_", "-", "`", "`", "'", "´", "-", "_", "_", "_"}
	material = []string{
		"█▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁",
		"██▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁",
		"███▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁",
		"████▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁",
		"██████▁▁▁▁▁▁▁▁▁▁▁▁▁▁",
		"██████▁▁▁▁▁▁▁▁▁▁▁▁▁▁",
		"███████▁▁▁▁▁▁▁▁▁▁▁▁▁",
		"████████▁▁▁▁▁▁▁▁▁▁▁▁",
		"█████████▁▁▁▁▁▁▁▁▁▁▁",
		"█████████▁▁▁▁▁▁▁▁▁▁▁",
		"██████████▁▁▁▁▁▁▁▁▁▁",
		"███████████▁▁▁▁▁▁▁▁▁",
		"█████████████▁▁▁▁▁▁▁",
		"██████████████▁▁▁▁▁▁",
		"██████████████▁▁▁▁▁▁",
		"▁██████████████▁▁▁▁▁",
		"▁██████████████▁▁▁▁▁",
		"▁██████████████▁▁▁▁▁",
		"▁▁██████████████▁▁▁▁",
		"▁▁▁██████████████▁▁▁",
		"▁▁▁▁█████████████▁▁▁",
		"▁▁▁▁██████████████▁▁",
		"▁▁▁▁██████████████▁▁",
		"▁▁▁▁▁██████████████▁",
		"▁▁▁▁▁██████████████▁",
		"▁▁▁▁▁██████████████▁",
		"▁▁▁▁▁▁██████████████",
		"▁▁▁▁▁▁██████████████",
		"▁▁▁▁▁▁▁█████████████",
		"▁▁▁▁▁▁▁█████████████",
		"▁▁▁▁▁▁▁▁████████████",
		"▁▁▁▁▁▁▁▁████████████",
		"▁▁▁▁▁▁▁▁▁███████████",
		"▁▁▁▁▁▁▁▁▁███████████",
		"▁▁▁▁▁▁▁▁▁▁██████████",
		"▁▁▁▁▁▁▁▁▁▁██████████",
		"▁▁▁▁▁▁▁▁▁▁▁▁████████",
		"▁▁▁▁▁▁▁▁▁▁▁▁▁███████",
		"▁▁▁▁▁▁▁▁▁▁▁▁▁▁██████",
		"▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁█████",
		"▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁█████",
		"▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁████",
		"▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁███",
		"▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁███",
		"▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁███",
		"▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁██",
		"▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁█",
		"▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁█",
		"▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁█",
	}

	shark = []string{
		"▐|\\____________▌",
		"▐_|\\___________▌",
		"▐__|\\__________▌",
		"▐___|\\_________▌",
		"▐____|\\________▌",
		"▐_____|\\_______▌",
		"▐______|\\______▌",
		"▐_______|\\_____▌",
		"▐________|\\____▌",
		"▐_________|\\___▌",
		"▐__________|\\__▌",
		"▐___________|\\_▌",
		"▐____________|\\▌",
		"▐____________/|▌",
		"▐___________/|_▌",
		"▐__________/|__▌",
		"▐_________/|___▌",
		"▐________/|____▌",
		"▐_______/|_____▌",
		"▐______/|______▌",
		"▐_____/|_______▌",
		"▐____/|________▌",
		"▐___/|_________▌",
		"▐__/|__________▌",
		"▐_/|___________▌",
		"▐/|____________▌",
	}

	betawave = []string{"ρββββββ", "βρβββββ", "ββρββββ", "βββρβββ", "ββββρββ", "βββββρβ", "ββββββρ"}
	fistbump = []string{
		"🤜\u3000\u3000\u3000\u3000🤛 ",
		"🤜\u3000\u3000\u3000\u3000🤛 ",
		"🤜\u3000\u3000\u3000\u3000🤛 ",
		"\u3000🤜\u3000\u3000🤛\u3000 ",
		"\u3000\u3000🤜🤛\u3000\u3000 ",
		"\u3000🤜✨🤛\u3000\u3000 ",
		"🤜\u3000✨\u3000🤛\u3000 ",
	}
	futbolHead = []string{
		" 🧑⚽️       🧑 ",
		"🧑  ⚽️      🧑 ",
		"🧑   ⚽️     🧑 ",
		"🧑    ⚽️    🧑 ",
		"🧑     ⚽️   🧑 ",
		"🧑      ⚽️  🧑 ",
		"🧑       ⚽️🧑  ",
		"🧑      ⚽️  🧑 ",
		"🧑     ⚽️   🧑 ",
		"🧑    ⚽️    🧑 ",
		"🧑   ⚽️     🧑 ",
		"🧑  ⚽️      🧑 ",
	}
	mindblown = []string{
		"😐 ",
		"😐 ",
		"😮 ",
		"😮 ",
		"😦 ",
		"😦 ",
		"😧 ",
		"😧 ",
		"🤯 ",
		"💥 ",
		"✨ ",
		"\u3000 ",
		"\u3000 ",
		"\u3000 ",
	}
	speaker = []string{"🔈 ", "🔉 ", "🔊 ", "🔉 "}
	star    = []string{
		"✶",
		"✸",
		"✹",
		"✺",
		"✹",
		"✷",
	}
)

var registry = []SpinnerStyle{
	// braille
	{SpinnerDefault, defaultSymbols, GroupBraille},
	{SpinnerBrailleWave, brailleWave, GroupBraille},
	{SpinnerDots, dots, GroupBraille},
	{SpinnerDots3, dots3, GroupBraille},
	{SpinnerDots4, dots4, GroupBraille},
	{SpinnerDots5, dots5, GroupBraille},
	{SpinnerDots6, dots6, GroupBraille},
	{SpinnerDots7, dots7, GroupBraille},

	// arrows
	{SpinnerArrow, arrow, GroupArrows},
	{SpinnerArrow2, arrow2, GroupArrows},
	{SpinnerArrow3, arrow3, GroupArrows},
	{SpinnerArrow4, arrow4, GroupArrows},
	{SpinnerSweep, sweep, GroupArrows},

	// lines
	{SpinnerSlash, slash, GroupLines},
	{SpinnerBackslash, backslash, GroupLines},
	{SpinnerPipe, pipe, GroupLines},
	{SpinnerPipe2, pipe2, GroupLines},
	{SpinnerLines, lines, GroupLines},

	// blocks
	{SpinnerBlock, block, GroupBlocks},
	{SpinnerBlockbar, blockbar, GroupBlocks},
	{SpinnerBlockbar2, blockbar2, GroupBlocks},
	{SpinnerBlockbar3, blockbar3, GroupBlocks},
	{SpinnerBlockbar4, blockbar4, GroupBlocks},
	{SpinnerBlockbar5, blockbar5, GroupBlocks},
	{SpinnerBlockbar6, blockbar6, GroupBlocks},
	{SpinnerBlockbar7, blockbar7, GroupBlocks},
	{SpinnerBlockbarPretty, blockbarpretty, GroupBlocks},
	{SpinnerBoxFill, boxfill, GroupBlocks},
	{SpinnerBoxFillshort, boxfillshort, GroupBlocks},
	{SpinnerBoxBounce, boxBounce, GroupBlocks},

	// motion
	{SpinnerBounce, bounce, GroupMotion},
	{SpinnerBounceball, bounceball, GroupMotion},
	{SpinnerPingpong, pingpong, GroupMotion},
	{SpinnerPingpong2, pingpong2, GroupMotion},
	{SpinnerRunner, runner, GroupMotion},

	// circular
	{SpinnerCircle, circle, GroupCircular},
	{SpinnerCircle2, circle2, GroupCircular},
	{SpinnerCircle3, circle3, GroupCircular},
	{SpinnerCircle4, circle4, GroupCircular},
	{SpinnerCircle5, circle5, GroupCircular},
	{SpinnerCircle6, circle6, GroupCircular},
	{SpinnerCircle7, circle7, GroupCircular},
	{SpinnerOrbit, orbit, GroupCircular},
	{SpinnerMoon, moon, GroupCircular},
	{SpinnerClock, clock, GroupCircular},

	// shapes
	{SpinnerSquare, square, GroupShapes},
	{SpinnerSquare2, square2, GroupShapes},
	{SpinnerCubes, cubes, GroupShapes},
	{SpinnerTriangles, triangles, GroupShapes},
	{SpinnerDiamond, diamond, GroupShapes},
	{SpinnerDiamond2, diamond2, GroupShapes},
	{SpinnerGeometric, geometric, GroupShapes},

	// text
	{SpinnerLoading, loading, GroupText},
	{SpinnerEllipsis, ellipsis, GroupText},
	{SpinnerQuestion, question, GroupText},
	{SpinnerHexsymbols, hexsymbols, GroupText},

	// symbols
	{SpinnerCurrency, currency, GroupSymbols},
	{SpinnerMathops, mathops, GroupSymbols},
	{SpinnerLogicsymbols, logicsymbols, GroupSymbols},
	{SpinnerGreek, greek, GroupSymbols},

	// fun
	{SpinnerPacman, pacman, GroupFun},
	{SpinnerSnail, snail, GroupFun},
	{SpinnerWorm, worm, GroupFun},
	{SpinnerWorm2, worm2, GroupFun},

	// minimal
	{SpinnerToggle, toggle, GroupMinimal},
	{SpinnerToggle2, toggle2, GroupMinimal},
	{SpinnerToggle3, toggle3, GroupMinimal},
	{SpinnerCursorBlink, cursorBlink, GroupMinimal},
	{SpinnerPluscross, pluscross, GroupMinimal},

	// effects
	{SpinnerFade, fade, GroupEffects},
	{SpinnerPulse, pulse, GroupEffects},
	{SpinnerGrow, grow, GroupEffects},
	{SpinnerGrowvert, growvert, GroupEffects},
	{SpinnerWave, wave, GroupEffects},

	// framed
	{SpinnerMarquee, marquee, GroupFramed},
	{SpinnerMatrix, matrix, GroupFramed},
	{SpinnerCorners, corners, GroupFramed},

	// misc
	{SpinnerFlip, flip, GroupMisc},
	{SpinnerMaterial, material, GroupMisc},
	{SpinnerShark, shark, GroupMisc},
	{SpinnerBetawave, betawave, GroupMisc},
	{SpinnerFistbump, fistbump, GroupMisc},
	{SpinnerFutbolHead, futbolHead, GroupMisc},
	{SpinnerMindblown, mindblown, GroupMisc},
	{SpinnerSpeaker, speaker, GroupMisc},
	{SpinnerStar, star, GroupMisc},
}
