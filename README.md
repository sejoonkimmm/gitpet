# gitpet

A Tamagotchi that lives in your git history.

It hatches when you install it, eats when you commit, evolves the more you code. Stop coding for a week and it dies.

```
    ,,,          /\_/\        /\_____/\           ★  ✦
   /   \        ( ^.^ )      /  o   o  \     ✧  /\_/\  ★
  | ^_^ |        > ~ <      ( ==  ^  == )   / ★   ★ \
  |     |       /|   |\      )         (    (  =▽=  )
   \___/       (_|   |_)    (__(__)___(__)    )     (

   Egg          Baby          Teen            Legend
```

## Install

```sh
go install github.com/sejoonkimmm/gitpet@latest
```

## Usage

```sh
gitpet init              # adopt a pet
gitpet                   # check on it
gitpet feed              # feed manually
gitpet hook install      # auto-feed on every commit
gitpet hook remove       # remove the hook
gitpet rip               # visit the graveyard
```

The hook is the whole point — install it in your repo and forget about it. Every commit feeds your pet.

## How it works

Hunger and happiness tick down over time. Commits reset hunger and give XP. Enough XP and your pet evolves.

Stages: Egg (10 XP) → Baby (30 XP) → Teen (70 XP) → Adult (150 XP) → Legend

If you stop coding:
- 1 day — hungry
- 3 days — sad, streak gone
- 7 days — dead

Dead pets go to the graveyard (`gitpet rip`). You can always adopt again.

## Example

```
$ gitpet

   /\_/\
  ( ^.^ )
   > ~ <
  /|   |\
 (_|   |_)

  Name:      FuzzyMochi 😊
  Stage:     🐣 Baby
  Age:       3 days
  Commits:   12
  Streak:    3 days 🔥

  Fullness:  🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟩🟥🟥🟥 80%
  Happiness: 💛💛💛💛💛💛💛💛💛💛💛💛🖤🖤🖤 80%

  Next evolution: 22/30 XP

  "Life is great! Keep those commits coming! ✨"
```

## Why

I just wanted something dumb and cute in my terminal. That's the whole reason.

## License

MIT
