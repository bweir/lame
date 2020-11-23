# Blocks

Lame supports 7 "blocks".

## `asm` - Assembly

These blocks can be used to write assembly language directly.

```
asm
    mov     eins, rcnt
    or      eins, CMD_SetPage   ' chip select embedded
    call    #sendLCDcommand     ' set page
```

Available syntax will depend on the platform.

## `con` - Constants

These blocks are used to define constant values that will
never change.

```
con
    volume = 11
    pi = 3.14
    greet = "hello world"
```

## `dat` - Data

These blocks are used to store data directly in the program.

```
dat
    diamond
    byte 5,5
    byte 0,0,1,0,0
    byte 0,1,1,1,0
    byte 1,1,1,1,1
    byte 0,1,1,1,0
    byte 0,0,1,0,0
```

`dat` blocks are shared by all instances of an object.

## `obj` - Objects

```
obj
    ser = "com.serial"
```

## `pri` - Private Functions

Private functions are only available to the current object.

```
pri foo(a, b)
    return a + b
```

## `pub` - Public Functions

Public functions are exported so that other objects can call them.

```
pub bar(x)
    return x >> 2
```

Programs start in the first `pub` block
of the main object.

## `var` - Variables

These blocks define values that can change.

```
var
  long  controls
  long  shadow
  byte  stack[32]
```

Each instance of an object has its own `var` block. The values are not shared.
