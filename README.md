# ft_ality

## Usage

### Grammar file

How to format the grammar file :

```txt
(key single uppercase letter (ex: q) OR [Left|Right|Up|Down]):("action name")
...
(key single uppercase letter (ex: q) OR [Left|Right|Up|Down]):("action name")
(SINGLE EMPTY LINE TO SEPARATE ACTIONS THAN COMBOS)
([action name],[action name]):("Combo name")
...
([action name],[action name]):("Combo name")
```

Example :

```txt
q:Block
Down:Down
w:Flip Stance
Left:Left
Right:Right
e:Tag
a:Throw
up:Up
s:[BK]
d:[BP]
z:[FK]
x:[FP]

[BP]:Claw Slam (Freddy krueger)
[BP]:Knockdown (Sonya)
[BP]:Fist of Death (Liu-Kang)
[BP],[FP]:Saibot Blast (Noob Saibot)
[BP],[FP]:Active Duty (Jax)
```

## Requirements

### SDL

https://github.com/veandco/go-sdl2#installation

Install sdl in macos :

```bash
brew install sdl2{,_image,_mixer,_ttf,_gfx} pkg-config
```
