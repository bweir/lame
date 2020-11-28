package lexer

func isSpace(ch rune) bool {
	return ch == ' ' || ch == '\t'
}

func isNewline(ch rune) bool {
	return ch == '\n'
}

func isIdentifier(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '_'
}

func isGroupSeparator(ch rune) bool {
	return ch == '_'
}

func isBinaryDigit(ch rune) bool {
	return (ch >= '0' && ch <= '1')
}

func isQuaternaryDigit(ch rune) bool {
	return (ch >= '0' && ch <= '3')
}

func isDecimalDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

func isHexadecimalDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'f') || (ch >= 'A' && ch <= 'F')
}

func isLineCommentStart(ch rune) bool {
	return ch == '\''
}

func isLineCommentEnd(ch rune) bool {
	return ch == '\n'
}

func isCommentStart(ch rune) bool {
	return ch == '{'
}

func isCommentEnd(ch rune) bool {
	return ch == '}'
}
