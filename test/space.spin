

PUB Load(source, start, w, h)
    font := source 
    startingchar := start

PUB Char(char_byte, x, y)
    gfx.Sprite(font, x, y, char_byte - startingchar)
