package pet

import "fmt"

// GetArt returns ASCII art based on pet stage and mood
func GetArt(stage Stage, mood Mood) string {
	switch stage {
	case Egg:
		return eggArt(mood)
	case Baby:
		return babyArt(mood)
	case Teen:
		return teenArt(mood)
	case Adult:
		return adultArt(mood)
	case Legend:
		return legendArt(mood)
	default:
		return eggArt(Normal)
	}
}

func eggArt(mood Mood) string {
	switch mood {
	case Happy:
		return `
    ,,,
   /   \
  | ^_^ |
  |     |
   \___/
`
	case Hungry:
		return `
    ,,,
   /   \
  | ._. |
  |     |
   \___/
`
	case Sad:
		return `
    ,,,
   /   \
  | ;_; |
  |     |
   \___/
`
	case Dead:
		return `
    ,,,
   /   \
  | x_x |
  |     |
   \___/
`
	default:
		return `
    ,,,
   /   \
  | o_o |
  |     |
   \___/
`
	}
}

func babyArt(mood Mood) string {
	switch mood {
	case Happy:
		return `
   /\_/\
  ( ^.^ )
   > ~ <
  /|   |\
 (_|   |_)
`
	case Hungry:
		return `
   /\_/\
  ( o.o )
   > ~ <
  /|   |\
 (_|   |_)
`
	case Sad:
		return `
   /\_/\
  ( ;.; )
   > ~ <
  /|   |\
 (_|   |_)
`
	case Dead:
		return `
   /\_/\
  ( x.x )
   > ~ <
  /|   |\
 (_|   |_)
`
	default:
		return `
   /\_/\
  ( -.- )
   > ~ <
  /|   |\
 (_|   |_)
`
	}
}

func teenArt(mood Mood) string {
	switch mood {
	case Happy:
		return `
    /\_____/\
   /  o   o  \
  ( ==  ^  == )
   )         (
  (           )
 ( (  )   (  ) )
(__(__)___(__)__)
`
	case Hungry:
		return `
    /\_____/\
   /  .   .  \
  ( ==  o  == )
   )         (
  (           )
 ( (  )   (  ) )
(__(__)___(__)__)
`
	case Sad:
		return `
    /\_____/\
   /  T   T  \
  ( ==  n  == )
   )         (
  (           )
 ( (  )   (  ) )
(__(__)___(__)__)
`
	case Dead:
		return `
    /\_____/\
   /  x   x  \
  ( ==  n  == )
   )         (
  (           )
 ( (  )   (  ) )
(__(__)___(__)__)
`
	default:
		return `
    /\_____/\
   /  -   -  \
  ( ==  ^  == )
   )         (
  (           )
 ( (  )   (  ) )
(__(__)___(__)__)
`
	}
}

func adultArt(mood Mood) string {
	switch mood {
	case Happy:
		return `
      /\_/\
     / o o \
    (  =^=  )
     )     (
    (       )
   / |     | \
  /  |     |  \
 /   |     |   \
(    |     |    )
 \   |     |   /
  '''|     |'''
     |     |
     |     |
     (_   _)
      (   )
       \_/
`
	case Hungry:
		return `
      /\_/\
     / . . \
    (  =o=  )
     )     (
    (       )
   / |     | \
  /  |     |  \
 /   |     |   \
(    |     |    )
 \   |     |   /
  '''|     |'''
     |     |
     |     |
     (_   _)
      (   )
       \_/
`
	case Sad:
		return `
      /\_/\
     / T T \
    (  =n=  )
     )     (
    (       )
   / |     | \
  /  |     |  \
 /   |     |   \
(    |     |    )
 \   |     |   /
  '''|     |'''
     |     |
     |     |
     (_   _)
      (   )
       \_/
`
	case Dead:
		return `
      /\_/\
     / x x \
    (  =n=  )
     )     (
    (       )
   / |     | \
  /  |     |  \
 /   |     |   \
(    |     |    )
 \   |     |   /
  '''|     |'''
     |     |
     |     |
     (_   _)
      (   )
       \_/
`
	default:
		return `
      /\_/\
     / - - \
    (  =^=  )
     )     (
    (       )
   / |     | \
  /  |     |  \
 /   |     |   \
(    |     |    )
 \   |     |   /
  '''|     |'''
     |     |
     |     |
     (_   _)
      (   )
       \_/
`
	}
}

func legendArt(mood Mood) string {
	face := "^"
	eyes := "★"
	switch mood {
	case Happy:
		face = "▽"
		eyes = "★"
	case Hungry:
		face = "o"
		eyes = "◆"
	case Sad:
		face = "n"
		eyes = "◇"
	case Dead:
		face = "n"
		eyes = "✖"
	default:
		face = "^"
		eyes = "★"
	}
	return fmt.Sprintf(`
        ★  ✦
    ✧  /\_/\  ★
  ✦   / %s %s \   ✧
     (  =%s=  )
    ★ )     ( ✦
     (  ★ ★  )
    / |     | \
   /  | ✧   |  \
  / ★ |     | ✦ \
 (    |  ★  |    )
  \   |     |   /
   '''|  ✧  |'''
      |     |
    ★ |     | ✦
      (_   _)
   ✧  (   )  ★
       \_/
     ✦    ✧
`, eyes, eyes, face)
}

func DeadArt() string {
	return `
    .----.
   /      \
  |  R.I.P |
  |        |
  | Here   |
  | lies a |
  | beloved|
  | gitpet |
  |        |
  |________|
 /|  /\  /|\
/_|_/  \_|_\
  ~  ~~  ~
`
}

func GetFrame(p *Pet) string {
	mood := CalculateMood(p)
	art := GetArt(p.Stage, mood)
	return art
}
