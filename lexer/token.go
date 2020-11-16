package lexer

type Token int

var eof = rune(0)

const (
	// Special tokens
	ILLEGAL Token = iota

	EOF
	WS
	NEWLINE // \n

	// Literals
	IDENTIFIER // fields, table_name
	NUMBER
	COMMENT
	DOC_COMMENT

	INDENT
	DEDENT

	// Operators
	EQUAL            // =
	EXCLAMATION_MARK // !
	AT               // @
	POUND            // #
	DOLLAR           // $
	PERCENT          // %
	DOT              // .
	PIPE             // |

	LESS_THAN    // <
	GREATER_THAN // >
	TILDE        // ~
	AMPERSAND    // &

	PLUS     // +
	MINUS    // -
	ASTERISK // *
	SLASH    // /

	// Misc characters
	BRACKET_OPEN  // [
	BRACKET_CLOSE // ]
	COMMA         // ,
	PAREN_OPEN    // (
	PAREN_CLOSE   // )
	COLON         // :
	QUOTE_DOUBLE  // "
	QUOTE_SINGLE  // '
	BRACE_OPEN    // {
	BRACE_CLOSE   // }

	// Keywords
	// Blocks
	CON
	DAT
	OBJ
	PRI
	PUB
	VAR

	// Constants
	TRUE
	FALSE
	POSX
	NEGX
	PI
	RCFAST
	RCSLOW
	XINPUT
	XTAL1
	XTAL2
	XTAL3
	PLL1X
	PLL2X
	PLL4X
	PLL8X
	PLL16X

	// Variables
	RESULT

	// Flow Control
	ABORT
	CASE
	IF
	IFNOT
	NEXT
	QUIT
	REPEAT
	RETURN

	// Memory
	BYTE
	WORD
	LONG
	BYTEFILL
	WORDFILL
	LONGFILL
	BYTEMOVE
	WORDMOVE
	LONGMOVE
	LOOKUP
	LOOKUPZ
	LOOKDOWN
	LOOKDOWNZ
	STRSIZE
	STRCOMP

	// Directives
	STRING
	CONSTANT
	FLOAT
	ROUND
	TRUNC
	FILE

	// Locks
	LOCKNEW
	LOCKRET
	LOCKCLR
	LOCKSET

	// Chip Config
	CHIPVER
	CLKMODE
	_CLKMODE
	CLKFREQ
	_CLKFREQ
	CLKSET
	_XINFREQ
	_STACK
	_FREE

	// Registers
	CNT
	CTRA
	CTRB
	DIRA
	DIRB
	INA
	INB
	OUTA
	OUTB
	FRQA
	FRQB
	PHSA
	PHSB
	VCFG
	VSCL
	PAR
	SPR

	// Process Control
	WAITCNT
	WAITPEQ
	WAITPNE
	WAITVID

	// Cog Control
	COGID
	COGNEW
	COGINIT
	COGSTOP
	REBOOT
)
