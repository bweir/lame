# Comments

Comments are used to describe what your code does.

Lame supports four kinds of comments.

- Regular comments, or "Doc" (documentation) comments

- Single line or multi-line

## Single-Line Comment

Everything after a single quote (`'`)is commented out.

```
' my comment is here
```

## Single-Line Doc Comment

A second single quote marks a "doc" comment.

```
'' My doc comment is here
```

## Multi-Line Comment

Anything between curly braces (`{}`) is a comment.

```
{
    My comment is here
}
```

## Multi-Line Doc Comment

A second set of curly braces marks a "doc" comment.

```
{{
    My doc comment is here
}}
```

## Documentation

Doc comments written in Markdown can be rendered by the `lame` tool.

```
lame doc MyObject.lame
```
