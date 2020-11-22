# lame compiler suite

I'm making a Spin compiler. Why?

- Spin is an interesting language for its limitations

- LameStation has years of content built around a subset of Spin

- The ecosystem is fragmented and not fully specified

- A documented lexer/parser would enable many kinds of analyses

- This gives me an excuse to learn Go

Notes:

- https://blog.gopheracademy.com/advent-2014/parsers-lexers/

- https://github.com/benbjohnson/sql-parser/blob/master/scanner.go

## Usage

```
lame COMMAND [OPTIONS]
```

## Goals

- **To compile Spin objects**:

  ```
  lame build
  ```

- **To load Spin objects over serial**:

  ```
  lame load
  ```

- To enable a live build/load/debug cycle over serial

  ```
  lame dev
  ```

- To reformat Spin code accurately:

  ```
  lame fmt
  ```

- To generate documentation from any Spin object:

  ```
  lame doc
  ```

- To push and pull Spin objects to/from a remote repo:

  ```bash
  # upload an object to the repo
  lame up

  # pull an object
  lame down pst=1.0.0
  ```

- To support multiple code generator backends (Propeller VM?)

### Finer Goals

- To write a (semi-)formal specification for the Spin language

- To perform code generation for the Propeller 1

- To support an alternative VM target for desktop / browser use

- To support direct addressing via the `@` symbol

- To add a path system to Spin:

  ```
  OBJ
      obj : "path.to.module"
  ```

- To add automatic constant folding

- To reduce the number of operators / keywords

## Roadmap

- Identity - Translate the input to the AST and re-emit to the output

- Reverse engineer Propeller binary format
