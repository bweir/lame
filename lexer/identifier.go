package lexer

import (
	"bytes"
	"strings"

	"github.com/bweir/lame/token"
	"github.com/bweir/lame/token/state"
)

func (s *Scanner) scanIdentifier() (tok token.Token) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isIdentifier(ch) && !isDigit(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	// If the string matches a keyword then return that keyword.
	switch strings.ToUpper(buf.String()) {

	// Blocks
	case "CON":
		s.state = state.CONSTANT
		s.blockStart = true
		return s.makeToken(token.CON, buf.String())
	case "DAT":
		s.state = state.DATA
		s.blockStart = true
		return s.makeToken(token.DAT, buf.String())
	case "OBJ":
		s.state = state.OBJECT
		s.blockStart = true
		return s.makeToken(token.OBJ, buf.String())
	case "PRI":
		s.state = state.FUNCTION
		s.blockStart = true
		return s.makeToken(token.PRI, buf.String())
	case "PUB":
		s.state = state.FUNCTION
		s.blockStart = true
		return s.makeToken(token.PUB, buf.String())
	case "VAR":
		s.state = state.VARIABLE
		s.blockStart = true
		return s.makeToken(token.VAR, buf.String())

	// Constants
	case "TRUE":
		return s.makeToken(token.TRUE, buf.String())
	case "FALSE":
		return s.makeToken(token.FALSE, buf.String())
	case "POSX":
		return s.makeToken(token.POSX, buf.String())
	case "NEGX":
		return s.makeToken(token.NEGX, buf.String())
	case "PI":
		return s.makeToken(token.PI, buf.String())
	case "RCFAST":
		return s.makeToken(token.RCFAST, buf.String())
	case "RCSLOW":
		return s.makeToken(token.RCSLOW, buf.String())
	case "XINPUT":
		return s.makeToken(token.XINPUT, buf.String())
	case "XTAL1":
		return s.makeToken(token.XTAL1, buf.String())
	case "XTAL2":
		return s.makeToken(token.XTAL2, buf.String())
	case "XTAL3":
		return s.makeToken(token.XTAL3, buf.String())
	case "PLL1X":
		return s.makeToken(token.PLL1X, buf.String())
	case "PLL2X":
		return s.makeToken(token.PLL2X, buf.String())
	case "PLL4X":
		return s.makeToken(token.PLL4X, buf.String())
	case "PLL8X":
		return s.makeToken(token.PLL8X, buf.String())
	case "PLL16X":
		return s.makeToken(token.PLL16X, buf.String())

	// Variables
	case "RESULT":
		return s.makeToken(token.RESULT, buf.String())

	// Flow Control
	case "ABORT":
		return s.makeToken(token.ABORT, buf.String())
	case "CASE":
		return s.makeToken(token.CASE, buf.String())
	case "IF":
		return s.makeToken(token.IF, buf.String())
	case "IFNOT":
		return s.makeToken(token.IFNOT, buf.String())
	case "NEXT":
		return s.makeToken(token.NEXT, buf.String())
	case "QUIT":
		return s.makeToken(token.QUIT, buf.String())
	case "REPEAT":
		return s.makeToken(token.REPEAT, buf.String())
	case "RETURN":
		return s.makeToken(token.RETURN, buf.String())

	// Memory
	case "BYTE":
		return s.makeToken(token.BYTE, buf.String())
	case "WORD":
		return s.makeToken(token.WORD, buf.String())
	case "LONG":
		return s.makeToken(token.LONG, buf.String())
	case "BYTEFILL":
		return s.makeToken(token.BYTEFILL, buf.String())
	case "WORDFILL":
		return s.makeToken(token.WORDFILL, buf.String())
	case "LONGFILL":
		return s.makeToken(token.LONGFILL, buf.String())
	case "BYTEMOVE":
		return s.makeToken(token.BYTEMOVE, buf.String())
	case "WORDMOVE":
		return s.makeToken(token.WORDMOVE, buf.String())
	case "LONGMOVE":
		return s.makeToken(token.LONGMOVE, buf.String())
	case "LOOKUP":
		return s.makeToken(token.LOOKUP, buf.String())
	case "LOOKUPZ":
		return s.makeToken(token.LOOKUPZ, buf.String())
	case "LOOKDOWN":
		return s.makeToken(token.LOOKDOWN, buf.String())
	case "LOOKDOWNZ":
		return s.makeToken(token.LOOKDOWNZ, buf.String())
	case "STRSIZE":
		return s.makeToken(token.STRSIZE, buf.String())
	case "STRCOMP":
		return s.makeToken(token.STRCOMP, buf.String())

	// Directives
	case "STRING":
		return s.makeToken(token.STRING, buf.String())
	case "CONSTANT":
		return s.makeToken(token.CONSTANT, buf.String())
	case "FLOAT":
		return s.makeToken(token.FLOAT, buf.String())
	case "ROUND":
		return s.makeToken(token.ROUND, buf.String())
	case "TRUNC":
		return s.makeToken(token.TRUNC, buf.String())
	case "FILE":
		return s.makeToken(token.FILE, buf.String())

	// Locks
	case "LOCKNEW":
		return s.makeToken(token.LOCKNEW, buf.String())
	case "LOCKRET":
		return s.makeToken(token.LOCKRET, buf.String())
	case "LOCKCLR":
		return s.makeToken(token.LOCKCLR, buf.String())
	case "LOCKSET":
		return s.makeToken(token.LOCKSET, buf.String())

	// Chip Config
	case "CHIPVER":
		return s.makeToken(token.CHIPVER, buf.String())
	case "CLKMODE":
		return s.makeToken(token.CLKMODE, buf.String())
	case "CLKFREQ":
		return s.makeToken(token.CLKFREQ, buf.String())
	case "CLKSET":
		return s.makeToken(token.CLKSET, buf.String())
	case "_CLKMODE":
		return s.makeToken(token.P_CLKMODE, buf.String())
	case "_CLKFREQ":
		return s.makeToken(token.P_CLKFREQ, buf.String())
	case "_XINFREQ":
		return s.makeToken(token.P_XINFREQ, buf.String())
	case "_STACK":
		return s.makeToken(token.P_STACK, buf.String())
	case "_FREE":
		return s.makeToken(token.P_FREE, buf.String())

	// Registers
	case "CNT":
		return s.makeToken(token.CNT, buf.String())
	case "CTRA":
		return s.makeToken(token.CTRA, buf.String())
	case "CTRB":
		return s.makeToken(token.CTRB, buf.String())
	case "DIRA":
		return s.makeToken(token.DIRA, buf.String())
	case "DIRB":
		return s.makeToken(token.DIRB, buf.String())
	case "INA":
		return s.makeToken(token.INA, buf.String())
	case "INB":
		return s.makeToken(token.INB, buf.String())
	case "OUTA":
		return s.makeToken(token.OUTA, buf.String())
	case "OUTB":
		return s.makeToken(token.OUTB, buf.String())
	case "FRQA":
		return s.makeToken(token.FRQA, buf.String())
	case "FRQB":
		return s.makeToken(token.FRQB, buf.String())
	case "PHSA":
		return s.makeToken(token.PHSA, buf.String())
	case "PHSB":
		return s.makeToken(token.PHSB, buf.String())
	case "VCFG":
		return s.makeToken(token.VCFG, buf.String())
	case "VSCL":
		return s.makeToken(token.VSCL, buf.String())
	case "PAR":
		return s.makeToken(token.PAR, buf.String())
	case "SPR":
		return s.makeToken(token.SPR, buf.String())

	// Process Control
	case "WAITCNT":
		return s.makeToken(token.WAITCNT, buf.String())
	case "WAITPEQ":
		return s.makeToken(token.WAITPEQ, buf.String())
	case "WAITPNE":
		return s.makeToken(token.WAITPNE, buf.String())
	case "WAITVID":
		return s.makeToken(token.WAITVID, buf.String())

	// Cog Control
	case "COGID":
		return s.makeToken(token.COGID, buf.String())
	case "COGNEW":
		return s.makeToken(token.COGNEW, buf.String())
	case "COGINIT":
		return s.makeToken(token.COGINIT, buf.String())
	case "COGSTOP":
		return s.makeToken(token.COGSTOP, buf.String())
	case "REBOOT":
		return s.makeToken(token.REBOOT, buf.String())

	// logical
	case "NOT":
		return s.makeToken(token.NOT, buf.String())
	case "AND":
		return s.makeToken(token.AND, buf.String())
	case "OR":
		return s.makeToken(token.OR, buf.String())
	}

	return s.makeToken(token.IDENTIFIER, buf.String())
}
