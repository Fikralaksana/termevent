package terminal

// https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797#general-ascii-codes
const (
	BEL = 7
	BS  = 8
	HT  = 9
	LF  = 10
	VT  = 11
	FF  = 12
	CR  = 13
	ESC = 27
	DEL = 127
)

// https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797#cursor-controls
const (
	ESC_HOME                   = "\x1B[H"
	ESC_CURSOR                 = "\x1B[%d;%dH"
	ESC_END                    = "\x1B[%d;%df"
	ESC_CURSOR_UP              = "\x1B[A"
	ESC_CURSOR_DOWN            = "\x1B[B"
	ESC_CURSOR_RIGHT           = "\x1B[C"
	ESC_CURSOR_LEFT            = "\x1B[D"
	ESC_CURSOR_NEXT            = "\x1B[E"
	ESC_CURSOR_PREV            = "\x1B[F"
	ESC_CURSOR_BEGIN           = "\x1B[G"
	ESC_CURSOR_REPORT          = "\x1B[6n"
	ESC_CURSOR_ONE_UP          = "\x1B[M"
	ESC_SAVE_CURSOR_POS_DEC    = "\x1B[s"
	ESC_RESTORE_CURSOR_POS_DEC = "\x1B[u"
	ESC_SAVE_CURSOR_POS_SCO    = "\x1B[s"
	ESC_RESTORE_CURSOR_POS_SCO = "\x1B[u"
)

// https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797#erase-functions
const (
	ESC_ERASE_SCREEN_DOWN                 = "\x1B[J"
	ESC_ERASE_FROM_CURSOR_TO_SCREEN_END   = "\x1B[0J"
	ESC_ERASE_FROM_CURSOR_TO_SCREEN_START = "\x1B[1J"
	ESC_ERASE_SCREEN                      = "\x1B[2J"
	ESC_ERASE_SAVED_LINES                 = "\x1B[3J"
	ESC_ERASE_IN_LINE                     = "\x1B[K"
	ESC_ERASE_FROM_CURSOR_TO_LINE_END     = "\x1B[0K"
	ESC_ERASE_FROM_CURSOR_TO_LINE_START   = "\x1B[1K"
	ESC_ERASE_LINE                        = "\x1B[2K"
)

// https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797#insert-delete-functions
const (
	ESC_SET_GRAPHICS_MODE                 = `\x1B\[[0-9;]*m`
	ESC_RESET_GRAPHICS_MODE               = "\x1B[0m"
	ESC_SET_GRAPHICS_MODE_BOLD            = "\x1B[1m"
	ESC_SET_GRAPHICS_MODE_DIM             = "\x1B[2m"
	ESC_SET_GRAPHICS_MODE_ITALIC          = "\x1B[3m"
	ESC_SET_GRAPHICS_MODE_UNDERLINE       = "\x1B[4m"
	ESC_SET_GRAPHICS_MODE_BLINK           = "\x1B[5m"
	ESC_SET_GRAPHICS_MODE_REVERSE         = "\x1B[7m"
	ESC_SET_GRAPHICS_MODE_INVISIBLE       = "\x1B[8m"
	ESC_SET_GRAPHICS_MODE_STRIKETHROUGH   = "\x1B[9m"
	ESC_RESET_GRAPHICS_MODE_BOLD          = "\x1B[22m"
	ESC_RESET_GRAPHICS_MODE_DIM           = "\x1B[22m"
	ESC_RESET_GRAPHICS_MODE_ITALIC        = "\x1B[23m"
	ESC_RESET_GRAPHICS_MODE_UNDERLINE     = "\x1B[24m"
	ESC_RESET_GRAPHICS_MODE_BLINK         = "\x1B[25m"
	ESC_RESET_GRAPHICS_MODE_REVERSE       = "\x1B[27m"
	ESC_RESET_GRAPHICS_MODE_INVISIBLE     = "\x1B[28m"
	ESC_RESET_GRAPHICS_MODE_STRIKETHROUGH = "\x1B[29m"
)

// https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797#screen-modes
const (
	ESC_SCREEN_WIDTH_OR_TYPE       = "\x1B[=%dh"
	ESC_RESET_SCREEN_WIDTH_OR_TYPE = "\x1B[=%dl"
)

// https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797#common-private-modes
const (
	ESC_MAKE_CURSOR_INVISIBLE      = "\x1B[?25l"
	ESC_MAKE_CURSOR_VISIBLE        = "\x1B[?25h"
	ESC_RESTORE_SCREEN             = "\x1B[?47l"
	ESC_SAVE_SCREEN                = "\x1B[?47h"
	ESC_ENABLE_ALTERNATIVE_BUFFER  = "\x1B[?1049h"
	ESC_DISABLE_ALTERNATIVE_BUFFER = "\x1B[?1049l"
)

// https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797#keyboard-strings
const (
	ESC_REDEFINE_KEY = "\x1B[%s;%s;%sp"
)

const ENTER = 13
