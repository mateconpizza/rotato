package rotato

import "time"

type SpinnerName string

const (
	SpinnerDefault SpinnerName = "default"

	SpinnerBrailleWave    SpinnerName = "brailleWave"
	SpinnerBraillePulse   SpinnerName = "braillePulse"
	SpinnerBrailleSpin    SpinnerName = "brailleSpin"
	SpinnerBrailleOrbit   SpinnerName = "brailleOrbit"
	SpinnerBrailleBounce  SpinnerName = "brailleBounce"
	SpinnerBrailleScanner SpinnerName = "brailleScanner"
	SpinnerBrailleFire    SpinnerName = "brailleFire"
	SpinnerBrailleSpark   SpinnerName = "brailleSpark"

	SpinnerDots  SpinnerName = "dots"
	SpinnerDots3 SpinnerName = "dots3"
	SpinnerDots4 SpinnerName = "dots4"
	SpinnerDots5 SpinnerName = "dots5"
	SpinnerDots6 SpinnerName = "dots6"
	SpinnerDots7 SpinnerName = "dots7"

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
	GroupDefault SpinnerGroup = "default"

	GroupArrows   SpinnerGroup = "arrows"
	GroupBlocks   SpinnerGroup = "blocks"
	GroupDots     SpinnerGroup = "dots"
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
	Name      SpinnerName
	Frames    []string
	Group     SpinnerGroup
	Frequency time.Duration
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
		if spinner.Group == GroupDefault {
			continue
		}

		if _, ok := seen[spinner.Group]; ok {
			continue
		}

		seen[spinner.Group] = struct{}{}
		groups = append(groups, spinner.Group)
	}

	return groups
}()

var (
	defaultSymbols = dots

	// dots.
	dots  = []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
	dots3 = []string{"⠄", "⠆", "⠇", "⠋", "⠙", "⠸", "⠰", "⠠", "⠰", "⠸", "⠙", "⠋", "⠇", "⠆"}
	dots4 = []string{"⠁", "⠃", "⠇", "⠧", "⠷", "⠿", "⠷", "⠧", "⠇", "⠃"}
	dots5 = []string{"⠁", "⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈", "⠈"}
	dots6 = []string{
		"⢀⠀", "⡀⠀", "⠄⠀", "⢂⠀", "⡂⠀", "⠅⠀", "⢃⠀", "⡃⠀", "⠍⠀", "⢋⠀", "⡋⠀", "⠍⠁", "⢋⠁", "⡋⠁", "⠍⠉", "⠋⠉", "⠋⠉", "⠉⠙", "⠉⠙", "⠉⠩",
		"⠈⢙", "⠈⡙", "⢈⠩", "⡀⢙", "⠄⡙", "⢂⠩", "⡂⢘", "⠅⡘", "⢃⠨", "⡃⢐", "⠍⡐", "⢋⠠", "⡋⢀", "⠍⡁", "⢋⠁", "⡋⠁", "⠍⠉", "⠋⠉", "⠋⠉", "⠉⠙",
		"⠉⠙", "⠉⠩", "⠈⢙", "⠈⡙", "⠈⠩", "⠀⢙", "⠀⡙", "⠀⠩", "⠀⢘", "⠀⡘", "⠀⠨", "⠀⢐", "⠀⡐", "⠀⠠", "⠀⢀", "⠀⡀",
	}
	dots7 = []string{
		"⠀", "⠁", "⠂", "⠃", "⠄", "⠅", "⠆", "⠇", "⡀", "⡁", "⡂", "⡃", "⡄", "⡅", "⡆", "⡇", "⠈", "⠉", "⠊", "⠋", "⠌", "⠍", "⠎", "⠏",
		"⡈", "⡉", "⡊", "⡋", "⡌", "⡍", "⡎", "⡏", "⠐", "⠑", "⠒", "⠓", "⠔", "⠕", "⠖", "⠗", "⡐", "⡑", "⡒", "⡓", "⡔", "⡕", "⡖", "⡗",
		"⠘", "⠙", "⠚", "⠛", "⠜", "⠝", "⠞", "⠟", "⡘", "⡙", "⡚", "⡛", "⡜", "⡝", "⡞", "⡟", "⠠", "⠡", "⠢", "⠣", "⠤", "⠥", "⠦", "⠧",
		"⡠", "⡡", "⡢", "⡣", "⡤", "⡥", "⡦", "⡧", "⠨", "⠩", "⠪", "⠫", "⠬", "⠭", "⠮", "⠯", "⡨", "⡩", "⡪", "⡫", "⡬", "⡭", "⡮", "⡯",
		"⠰", "⠱", "⠲", "⠳", "⠴", "⠵", "⠶", "⠷", "⡰", "⡱", "⡲", "⡳", "⡴", "⡵", "⡶", "⡷", "⠸", "⠹", "⠺", "⠻", "⠼", "⠽", "⠾", "⠿",
		"⡸", "⡹", "⡺", "⡻", "⡼", "⡽", "⡾", "⡿", "⢀", "⢁", "⢂", "⢃", "⢄", "⢅", "⢆", "⢇", "⣀", "⣁", "⣂", "⣃", "⣄", "⣅", "⣆", "⣇",
		"⢈", "⢉", "⢊", "⢋", "⢌", "⢍", "⢎", "⢏", "⣈", "⣉", "⣊", "⣋", "⣌", "⣍", "⣎", "⣏", "⢐", "⢑", "⢒", "⢓", "⢔", "⢕", "⢖", "⢗",
		"⣐", "⣑", "⣒", "⣓", "⣔", "⣕", "⣖", "⣗", "⢘", "⢙", "⢚", "⢛", "⢜", "⢝", "⢞", "⢟", "⣘", "⣙", "⣚", "⣛", "⣜", "⣝", "⣞", "⣟",
		"⢠", "⢡", "⢢", "⢣", "⢤", "⢥", "⢦", "⢧", "⣠", "⣡", "⣢", "⣣", "⣤", "⣥", "⣦", "⣧", "⢨", "⢩", "⢪", "⢫", "⢬", "⢭", "⢮", "⢯",
		"⣨", "⣩", "⣪", "⣫", "⣬", "⣭", "⣮", "⣯", "⢰", "⢱", "⢲", "⢳", "⢴", "⢵", "⢶", "⢷", "⣰", "⣱", "⣲", "⣳", "⣴", "⣵", "⣶", "⣷",
		"⢸", "⢹", "⢺", "⢻", "⢼", "⢽", "⢾", "⢿", "⣸", "⣹", "⣺", "⣻", "⣼", "⣽", "⣾", "⣿",
	}

	// braille.
	braillePulse   = []string{"⣀", "⣄", "⣤", "⣦", "⣶", "⣿", "⣷", "⣯", "⣟", "⣻", "⣽", "⣾", "⣿", "⣾", "⣽", "⣻", "⣟", "⣯", "⣷", "⣶", "⣦", "⣤", "⣄", "⣀"}
	brailleSpin    = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏", "⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	brailleOrbit   = []string{"⠋", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋"}
	brailleBounce  = []string{"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈", "⠐", "⠠", "⢀", "⡀", "⠄", "⠂"}
	brailleWave    = []string{"⡀", "⡄", "⡆", "⡇", "⡏", "⡟", "⡿", "⣿", "⡿", "⡟", "⡏", "⡇", "⡆", "⡄"}
	brailleScanner = []string{"⠁", "⠃", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈"}
	brailleFire    = []string{"⠁", "⠉", "⠋", "⠛", "⠫", "⠭", "⠮", "⠷", "⠿", "⡿", "⢿", "⣷", "⣶", "⣤", "⣀", "⣄", "⣤", "⣶", "⣷", "⢿", "⡿", "⠿", "⠷", "⠮", "⠭", "⠫", "⠛", "⠋", "⠉", "⠁"}
	brailleSpark   = []string{"⠀", "⠁", "⠈", "⠐", "⠠", "⡀", "⢀", "⣀", "⣤", "⣶", "⣿", "⣷", "⣦", "⣄", "⢀", "⡀", "⠠", "⠐", "⠈", "⠁", "⠀"}

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
		"█▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁", "██▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁", "███▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁", "████▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁", "██████▁▁▁▁▁▁▁▁▁▁▁▁▁▁",
		"██████▁▁▁▁▁▁▁▁▁▁▁▁▁▁", "███████▁▁▁▁▁▁▁▁▁▁▁▁▁", "████████▁▁▁▁▁▁▁▁▁▁▁▁", "█████████▁▁▁▁▁▁▁▁▁▁▁", "█████████▁▁▁▁▁▁▁▁▁▁▁",
		"██████████▁▁▁▁▁▁▁▁▁▁", "███████████▁▁▁▁▁▁▁▁▁", "█████████████▁▁▁▁▁▁▁", "██████████████▁▁▁▁▁▁", "██████████████▁▁▁▁▁▁",
		"▁██████████████▁▁▁▁▁", "▁██████████████▁▁▁▁▁", "▁██████████████▁▁▁▁▁", "▁▁██████████████▁▁▁▁", "▁▁▁██████████████▁▁▁",
		"▁▁▁▁█████████████▁▁▁", "▁▁▁▁██████████████▁▁", "▁▁▁▁██████████████▁▁", "▁▁▁▁▁██████████████▁", "▁▁▁▁▁██████████████▁",
		"▁▁▁▁▁██████████████▁", "▁▁▁▁▁▁██████████████", "▁▁▁▁▁▁██████████████", "▁▁▁▁▁▁▁█████████████", "▁▁▁▁▁▁▁█████████████",
		"▁▁▁▁▁▁▁▁████████████", "▁▁▁▁▁▁▁▁████████████", "▁▁▁▁▁▁▁▁▁███████████", "▁▁▁▁▁▁▁▁▁███████████", "▁▁▁▁▁▁▁▁▁▁██████████",
		"▁▁▁▁▁▁▁▁▁▁██████████", "▁▁▁▁▁▁▁▁▁▁▁▁████████", "▁▁▁▁▁▁▁▁▁▁▁▁▁███████", "▁▁▁▁▁▁▁▁▁▁▁▁▁▁██████", "▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁█████",
		"▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁█████", "▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁████", "▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁███", "▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁███", "▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁███",
		"▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁██", "▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁█", "▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁█", "▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁▁█",
	}

	shark = []string{
		"▐|\\____________▌", "▐_|\\___________▌", "▐__|\\__________▌", "▐___|\\_________▌", "▐____|\\________▌", "▐_____|\\_______▌",
		"▐______|\\______▌", "▐_______|\\_____▌", "▐________|\\____▌", "▐_________|\\___▌", "▐__________|\\__▌", "▐___________|\\_▌",
		"▐____________|\\▌", "▐____________/|▌", "▐___________/|_▌", "▐__________/|__▌", "▐_________/|___▌", "▐________/|____▌",
		"▐_______/|_____▌", "▐______/|______▌", "▐_____/|_______▌", "▐____/|________▌", "▐___/|_________▌", "▐__/|__________▌",
		"▐_/|___________▌", "▐/|____________▌",
	}

	betawave = []string{"ρββββββ", "βρβββββ", "ββρββββ", "βββρβββ", "ββββρββ", "βββββρβ", "ββββββρ"}
	fistbump = []string{
		"🤜\u3000\u3000\u3000\u3000🤛 ", "🤜\u3000\u3000\u3000\u3000🤛 ", "🤜\u3000\u3000\u3000\u3000🤛 ",
		"\u3000🤜\u3000\u3000🤛\u3000 ", "\u3000\u3000🤜🤛\u3000\u3000 ", "\u3000🤜✨🤛\u3000\u3000 ",
		"🤜\u3000✨\u3000🤛\u3000 ",
	}
	futbolHead = []string{
		" 🧑⚽️       🧑 ", "🧑  ⚽️      🧑 ", "🧑   ⚽️     🧑 ", "🧑    ⚽️    🧑 ",
		"🧑     ⚽️   🧑 ", "🧑      ⚽️  🧑 ", "🧑       ⚽️🧑  ", "🧑      ⚽️  🧑 ",
		"🧑     ⚽️   🧑 ", "🧑    ⚽️    🧑 ", "🧑   ⚽️     🧑 ", "🧑  ⚽️      🧑 ",
	}
	mindblown = []string{
		"😐 ", "😐 ", "😮 ", "😮 ", "😦 ", "😦 ", "😧 ", "😧 ",
		"🤯 ", "💥 ", "✨ ", "\u3000 ", "\u3000 ", "\u3000 ",
	}
	speaker = []string{"🔈 ", "🔉 ", "🔊 ", "🔉 "}
	star    = []string{"✶", "✸", "✹", "✺", "✹", "✷"}
)

const defaultFreq = 100 * time.Millisecond

var registry = []SpinnerStyle{
	// dots
	{SpinnerDefault, defaultSymbols, GroupDots, 80 * time.Millisecond},
	{SpinnerDots, dots, GroupDots, 80 * time.Millisecond},
	{SpinnerDots3, dots3, GroupDots, 80 * time.Millisecond},
	{SpinnerDots4, dots4, GroupDots, 80 * time.Millisecond},
	{SpinnerDots5, dots5, GroupDots, 80 * time.Millisecond},
	{SpinnerDots6, dots6, GroupDots, 80 * time.Millisecond},
	{SpinnerDots7, dots7, GroupDots, 80 * time.Millisecond},

	// braille
	{SpinnerBrailleWave, brailleWave, GroupBraille, 80 * time.Millisecond},
	{SpinnerBraillePulse, braillePulse, GroupBraille, 80 * time.Millisecond},
	{SpinnerBrailleSpin, brailleSpin, GroupBraille, 80 * time.Millisecond},
	{SpinnerBrailleOrbit, brailleOrbit, GroupBraille, 90 * time.Millisecond},
	{SpinnerBrailleBounce, brailleBounce, GroupBraille, defaultFreq},
	{SpinnerBrailleScanner, brailleScanner, GroupBraille, 80 * time.Millisecond},
	{SpinnerBrailleFire, brailleFire, GroupBraille, 70 * time.Millisecond},
	{SpinnerBrailleSpark, brailleSpark, GroupBraille, 90 * time.Millisecond},

	// arrows
	{SpinnerArrow, arrow, GroupArrows, defaultFreq},
	{SpinnerArrow2, arrow2, GroupArrows, defaultFreq},
	{SpinnerArrow3, arrow3, GroupArrows, defaultFreq},
	{SpinnerArrow4, arrow4, GroupArrows, 120 * time.Millisecond},
	{SpinnerSweep, sweep, GroupArrows, defaultFreq},

	// lines
	{SpinnerSlash, slash, GroupLines, defaultFreq},
	{SpinnerBackslash, backslash, GroupLines, defaultFreq},
	{SpinnerPipe, pipe, GroupLines, 120 * time.Millisecond},
	{SpinnerPipe2, pipe2, GroupLines, defaultFreq},
	{SpinnerLines, lines, GroupLines, 120 * time.Millisecond},

	// blocks
	{SpinnerBlock, block, GroupBlocks, 180 * time.Millisecond},
	{SpinnerBlockbar, blockbar, GroupBlocks, 120 * time.Millisecond},
	{SpinnerBlockbar2, blockbar2, GroupBlocks, 120 * time.Millisecond},
	{SpinnerBlockbar3, blockbar3, GroupBlocks, 120 * time.Millisecond},
	{SpinnerBlockbar4, blockbar4, GroupBlocks, 120 * time.Millisecond},
	{SpinnerBlockbar5, blockbar5, GroupBlocks, defaultFreq},
	{SpinnerBlockbar6, blockbar6, GroupBlocks, defaultFreq},
	{SpinnerBlockbar7, blockbar7, GroupBlocks, 120 * time.Millisecond},
	{SpinnerBlockbarPretty, blockbarpretty, GroupBlocks, 120 * time.Millisecond},
	{SpinnerBoxFill, boxfill, GroupBlocks, 120 * time.Millisecond},
	{SpinnerBoxFillshort, boxfillshort, GroupBlocks, 120 * time.Millisecond},
	{SpinnerBoxBounce, boxBounce, GroupBlocks, 150 * time.Millisecond},

	// motion
	{SpinnerBounce, bounce, GroupMotion, defaultFreq},
	{SpinnerBounceball, bounceball, GroupMotion, defaultFreq},
	{SpinnerPingpong, pingpong, GroupMotion, 120 * time.Millisecond},
	{SpinnerPingpong2, pingpong2, GroupMotion, 80 * time.Millisecond},
	{SpinnerRunner, runner, GroupMotion, 80 * time.Millisecond},

	// circular
	{SpinnerCircle, circle, GroupCircular, 150 * time.Millisecond},
	{SpinnerCircle2, circle2, GroupCircular, 120 * time.Millisecond},
	{SpinnerCircle3, circle3, GroupCircular, 120 * time.Millisecond},
	{SpinnerCircle4, circle4, GroupCircular, 120 * time.Millisecond},
	{SpinnerCircle5, circle5, GroupCircular, 120 * time.Millisecond},
	{SpinnerCircle6, circle6, GroupCircular, 120 * time.Millisecond},
	{SpinnerCircle7, circle7, GroupCircular, 120 * time.Millisecond},
	{SpinnerOrbit, orbit, GroupCircular, 150 * time.Millisecond},
	{SpinnerMoon, moon, GroupCircular, 150 * time.Millisecond},
	{SpinnerClock, clock, GroupCircular, 200 * time.Millisecond},

	// shapes
	{SpinnerSquare, square, GroupShapes, 150 * time.Millisecond},
	{SpinnerSquare2, square2, GroupShapes, 120 * time.Millisecond},
	{SpinnerCubes, cubes, GroupShapes, 150 * time.Millisecond},
	{SpinnerTriangles, triangles, GroupShapes, 150 * time.Millisecond},
	{SpinnerDiamond, diamond, GroupShapes, 150 * time.Millisecond},
	{SpinnerDiamond2, diamond2, GroupShapes, 120 * time.Millisecond},
	{SpinnerGeometric, geometric, GroupShapes, 120 * time.Millisecond},

	// text
	{SpinnerLoading, loading, GroupText, 180 * time.Millisecond},
	{SpinnerEllipsis, ellipsis, GroupText, 250 * time.Millisecond},
	{SpinnerQuestion, question, GroupText, 200 * time.Millisecond},
	{SpinnerHexsymbols, hexsymbols, GroupText, 120 * time.Millisecond},

	// symbols
	{SpinnerCurrency, currency, GroupSymbols, 200 * time.Millisecond},
	{SpinnerMathops, mathops, GroupSymbols, 180 * time.Millisecond},
	{SpinnerLogicsymbols, logicsymbols, GroupSymbols, 180 * time.Millisecond},
	{SpinnerGreek, greek, GroupSymbols, 180 * time.Millisecond},

	// fun
	{SpinnerFistbump, fistbump, GroupFun, 150 * time.Millisecond},
	{SpinnerPacman, pacman, GroupFun, 150 * time.Millisecond},
	{SpinnerMindblown, mindblown, GroupFun, 200 * time.Millisecond},
	{SpinnerFutbolHead, futbolHead, GroupFun, 150 * time.Millisecond},
	{SpinnerSpeaker, speaker, GroupFun, 200 * time.Millisecond},

	// minimal
	{SpinnerToggle, toggle, GroupMinimal, 200 * time.Millisecond},
	{SpinnerToggle2, toggle2, GroupMinimal, 200 * time.Millisecond},
	{SpinnerToggle3, toggle3, GroupMinimal, 250 * time.Millisecond},
	{SpinnerCursorBlink, cursorBlink, GroupMinimal, 300 * time.Millisecond},
	{SpinnerPluscross, pluscross, GroupMinimal, 200 * time.Millisecond},

	// effects
	{SpinnerFade, fade, GroupEffects, 120 * time.Millisecond},
	{SpinnerPulse, pulse, GroupEffects, 120 * time.Millisecond},
	{SpinnerGrow, grow, GroupEffects, 120 * time.Millisecond},
	{SpinnerGrowvert, growvert, GroupEffects, defaultFreq},
	{SpinnerWave, wave, GroupEffects, defaultFreq},

	// framed
	{SpinnerMarquee, marquee, GroupFramed, 120 * time.Millisecond},
	{SpinnerMatrix, matrix, GroupFramed, 150 * time.Millisecond},
	{SpinnerCorners, corners, GroupFramed, 200 * time.Millisecond},

	// misc
	{SpinnerFlip, flip, GroupMisc, 120 * time.Millisecond},
	{SpinnerMaterial, material, GroupMisc, 80 * time.Millisecond},
	{SpinnerShark, shark, GroupMisc, 80 * time.Millisecond},
	{SpinnerBetawave, betawave, GroupMisc, defaultFreq},
	{SpinnerStar, star, GroupMisc, 150 * time.Millisecond},
	{SpinnerSnail, snail, GroupMisc, 180 * time.Millisecond},
	{SpinnerWorm, worm, GroupMisc, 150 * time.Millisecond},
	{SpinnerWorm2, worm2, GroupMisc, 150 * time.Millisecond},
}
