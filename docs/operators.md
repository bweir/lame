# Operators

Lame is a high-level interpreted language that makes hardware
interaction obvious. The operators that are available reflect that.

## Arithmetic

- `+` - Add

- `-` - Subtract / Negate

- `*` - Multiply

- `/` - Divide

- `%` - Modulus

## Comparison

- `==` - Equals

- `<>` - Not equals

- `>` - Greater than

- `>=` - Greater than or equal to

- `<` - Less than

- `<=` - Less than or equal to

## Location

Inside an object `obj`, anything can be referenced with `.`.

```
obj.CONSTANT

obj.variable

obj.function()
```

## Address

The address of a symbol can be returned with `@`.

```
addr = @x
```

The value can be retrieved from an address with `*`.

```
y = *addr
```

## Math Built-Ins

- `sqrt(x)` - Square root

- `abs(x)` - Absolute value

- `max(x, y)` - Maximum

- `min(x, y)` - Minimum

- `rand(x)` - Random

## Precedence

From highest to lowest:

1. Unary: `--`, `++`, `~`, `~~`, `@`

1. Unary: `-`, `!`

1. Shift: `<<`, `>>`, `~>`, `<-`, `->`, `><`

1. Bitwise and: `&`

1. Bitwise or: `|`, `^`

1. Multiplicative: `*`, `/`, `//`

1. Additive: `+`, `-`

1. Relational: `==`, `<>`, `<`, `>`, `<=`, `=>`

1. Boolean not: `not`

1. Boolean and: `and`

1. Boolean or: `or`

1. Assignment: `:=`, `+=`, `-=`, `*=`, `/=`, `//=`, `<<=`, `>>=`, `~>=`, `->=`, `<-=`, `><=`, `&=`, `|=`, `^=`
