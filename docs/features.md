# Feature Roadmap

## Blocks

- [ ] Syntax
  - [x] `ASM`
  - [x] `CON`
  - [x] `DAT`
  - [x] `OBJ`
  - [x] `PRI`
  - [x] `PUB`
  - [x] `VAR`

## Comments

- [x] Syntax
  - [x] Single-line comment
  - [x] Single-line doc comment
  - [x] Multi-line comment
  - [x] Multi-line doc comment
- [ ] Tooling
  - [ ] Generate Markdown from doc comments

## Numbers

- [ ] Syntax
  - [ ] Binary numbers
  - [ ] Quaternary numbers
  - [ ] Decimal numbers
  - [ ] Hexadecimal numbers
  - [ ] Group separators
- [ ] Functions

## Operators

- [ ] Syntax
  - [ ] Basic
    - [x] `+` - add
    - [x] `-` - subtract
    - [x] `*` - multiply
    - [x] `/` - divide
    - [x] `%` - modulo
    - [x] `=` - assignment
    - [ ] `+=` - assignment add
    - [ ] `-=` - assignment subtract
    - [ ] `*=` - assignment multiply
    - [ ] `/=` - assignment divide
    - [ ] `%=` - assignment modulo
  - [ ] Comparison
    - [ ] `==` - equal to
    - [ ] `!=` - not equal to
    - [ ] `>=` - greater than or equal to
    - [ ] `>` - greater than
    - [ ] `>=` - greater than or equal to
    - [ ] `<` - less than
    - [ ] `<=` - less than or equal to
  - [ ] Bitwise
    - [ ] `&` - bitwise and
    - [ ] `!` - bitwise not
    - [ ] `|` - bitwise or
    - [ ] `<<` - bitwise shift left
    - [ ] `>>` - bitwise shift right
    - [ ] `~>` - bitwise signed shift right
  - [ ] Misc.
    - [ ] `@` - address

## Strings

- [ ] Syntax
  - [x] Quoted strings as string literals
  - [x] Escape common characters
    - [x] `\n`
    - [x] `\t`
  - [x] Escape any ASCII character code: `\0`, `\1`, ...

## Whitespace

- [ ] Syntax
  - [x] Ignore whitespace after line start
  - [x] Detect initial starting indent per block
  - [x] Generate `INDENT` and `DEDENT` tokens on subsequent indents and dedents
  - [ ] Detect inconsistent indents
  - [ ] Detect tabs
