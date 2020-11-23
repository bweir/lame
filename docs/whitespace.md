# Whitespace

Lame is whitespace-sensitive.

Valid:

```
s = 4
repeat
	b = 4
```

Invalid:

```
	s = 4
s = 4
```

Valid:

```
s = 1
	s = 2
		s = 3
s = 4
```

Invalid:

```
	s = 1
		s = 2
s = 3
```
