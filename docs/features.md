# Feature Roadmap

## Blocks

- [x] `ASM`
- [x] `CON`
- [x] `DAT`
- [x] `OBJ`
- [x] `PRI`
- [x] `PUB`
- [x] `VAR`

## Comments

- [x] Single-line comment
- [x] Single-line doc comment
- [x] Multi-line comment
- [x] Multi-line doc comment
- [ ] Generate Markdown from doc comments

## Numbers

- [x] Binary numbers
- [x] Quaternary numbers
- [x] Decimal numbers
- [x] Hexadecimal numbers
- [x] Group separators

## Operators

- [x] `+` - add
- [x] `-` - subtract
- [x] `*` - multiply
- [x] `/` - divide
- [x] `%` - modulo
- [x] `=` - assign
- [x] `+=` - add assign
- [x] `-=` - subtract assign
- [x] `*=` - multiply assign
- [x] `/=` - divide assign
- [x] `%=` - modulo assign
- [x] `==` - equal to
- [x] `>` - greater than
- [x] `>=` - greater than or equal to
- [x] `<` - less than
- [x] `<=` - less than or equal to
- [x] `&` - bitwise and
- [x] `&=` - bitwise and assign
- [x] `|` - bitwise or
- [x] `|=` - bitwise or assign
- [x] `^` - bitwise xor
- [x] `^=` - bitwise xor assign
- [x] `!` - bitwise not
- [x] `<<` - bitwise shift left
- [x] `>>` - bitwise shift right
- [x] `~>` - bitwise signed shift right
- [x] `~` - bitwise sign extend 7
- [x] `~~` - bitwise sign extend 15
- [ ] `@` - address

## Strings

- [x] Quoted strings as string literals
- [x] Escape common characters: `\n`, `\t`
- [x] Escape any ASCII character code: `\0`, `\1`, ...

## Whitespace

- [x] Ignore whitespace after line start
- [x] Detect initial starting indent per block
- [x] Generate `INDENT` and `DEDENT` tokens on subsequent indents and dedents
- [ ] Detect inconsistent indents
- [ ] Detect tabs
