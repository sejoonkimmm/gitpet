package pet

import (
	"fmt"
	"time"
)

func NewPet() *Pet {
	now := time.Now()
	return &Pet{
		Name:      RandomName(),
		Stage:     Egg,
		XP:        0,
		Hunger:    20,
		Happiness: 80,
		Commits:   0,
		Streak:    0,
		BornAt:    now,
		LastFed:   now,
		LastCheck: now,
		Alive:     true,
	}
}

func Feed(p *Pet) string {
	if !p.Alive {
		return "💀 Your pet is no longer with us... Use 'gitpet init' to adopt a new friend."
	}

	UpdateState(p)

	p.Commits++
	p.XP += 5
	p.LastFed = time.Now()

	// Reduce hunger
	p.Hunger -= 30
	if p.Hunger < 0 {
		p.Hunger = 0
	}

	// Increase happiness
	p.Happiness += 15
	if p.Happiness > 100 {
		p.Happiness = 100
	}

	// Streak bonus
	p.XP += streakBonus(p.Streak)

	// Check evolution
	evolved := tryEvolve(p)

	mood := CalculateMood(p)
	art := GetArt(p.Stage, mood)

	msg := fmt.Sprintf("%s\n  %s says: \"Yummy! Thank you!\" 🍖\n", art, p.Name)

	if evolved {
		msg += fmt.Sprintf("\n  🎉 %s evolved to %s! 🎉\n", p.Name, p.Stage.Korean())
	}

	msg += fmt.Sprintf("  [XP: %d | Hunger: %d | Happy: %d | Commits: %d]\n",
		p.XP, p.Hunger, p.Happiness, p.Commits)

	return msg
}

func UpdateState(p *Pet) {
	if !p.Alive {
		return
	}

	now := time.Now()
	hoursSinceLastFed := now.Sub(p.LastFed).Hours()
	hoursSinceLastCheck := now.Sub(p.LastCheck).Hours()

	if hoursSinceLastCheck < 0.01 {
		return
	}

	// Increase hunger over time (10 per 6 hours)
	hungerIncrease := int(hoursSinceLastCheck / 6 * 10)
	p.Hunger += hungerIncrease
	if p.Hunger > 100 {
		p.Hunger = 100
	}

	// Decrease happiness over time
	happinessDecrease := int(hoursSinceLastCheck / 12 * 10)
	p.Happiness -= happinessDecrease
	if p.Happiness < 0 {
		p.Happiness = 0
	}

	// Update streak
	daysSinceLastFed := hoursSinceLastFed / 24
	if daysSinceLastFed > 2 {
		p.Streak = 0
	}

	// Pet dies after 7 days without food
	if hoursSinceLastFed > 24*7 {
		p.Alive = false
	}

	p.LastCheck = now
}

func CalculateMood(p *Pet) Mood {
	if !p.Alive {
		return Dead
	}
	if p.Hunger >= 80 {
		return Sad
	}
	if p.Hunger >= 50 {
		return Hungry
	}
	if p.Happiness >= 70 {
		return Happy
	}
	if p.Happiness <= 30 {
		return Sad
	}
	return Normal
}

func StatusText(p *Pet) string {
	if !p.Alive {
		art := DeadArt()
		return fmt.Sprintf("%s\n  %s has passed away... 😢\n  They lived for %d commits.\n  Use 'gitpet init' to adopt a new friend.\n",
			art, p.Name, p.Commits)
	}

	UpdateState(p)

	mood := CalculateMood(p)
	art := GetArt(p.Stage, mood)

	ageHours := time.Since(p.BornAt).Hours()
	ageDays := int(ageHours / 24)

	hungerBar := progressBar(100-p.Hunger, 100, 15, "🟩", "🟥")
	happyBar := progressBar(p.Happiness, 100, 15, "💛", "🖤")

	nextEvol := ""
	if p.Stage < Legend {
		needed := p.Stage.XPToEvolve()
		nextEvol = fmt.Sprintf("  Next evolution: %d/%d XP\n", p.XP, needed)
	} else {
		nextEvol = "  ✨ MAX EVOLUTION REACHED ✨\n"
	}

	sayings := moodSayings(mood)

	return fmt.Sprintf(`%s
  Name:      %s %s
  Stage:     %s
  Age:       %d days
  Commits:   %d
  Streak:    %d days 🔥

  Fullness:  %s
  Happiness: %s

%s
  "%s"
`,
		art,
		p.Name, mood.Emoji(),
		p.Stage.Korean(),
		ageDays,
		p.Commits,
		p.Streak,
		hungerBar,
		happyBar,
		nextEvol,
		sayings,
	)
}

func tryEvolve(p *Pet) bool {
	if p.Stage >= Legend {
		return false
	}
	needed := p.Stage.XPToEvolve()
	if p.XP >= needed {
		p.Stage++
		return true
	}
	return false
}

func streakBonus(streak int) int {
	switch {
	case streak >= 30:
		return 10
	case streak >= 7:
		return 5
	case streak >= 3:
		return 2
	default:
		return 0
	}
}

func progressBar(value, max, width int, filled, empty string) string {
	if value < 0 {
		value = 0
	}
	if value > max {
		value = max
	}
	filledCount := value * width / max
	emptyCount := width - filledCount
	bar := ""
	for i := 0; i < filledCount; i++ {
		bar += filled
	}
	for i := 0; i < emptyCount; i++ {
		bar += empty
	}
	return fmt.Sprintf("%s %d%%", bar, value*100/max)
}

func moodSayings(mood Mood) string {
	switch mood {
	case Happy:
		return "Life is great! Keep those commits coming! ✨"
	case Normal:
		return "Doing okay... a commit would be nice though."
	case Hungry:
		return "I'm getting hungry... please commit something!"
	case Sad:
		return "I miss you... please come back and code with me..."
	case Dead:
		return "..."
	default:
		return "meow~"
	}
}

func KillPet(p *Pet, cause string) GravePet {
	return GravePet{
		Name:    p.Name,
		Stage:   p.Stage,
		Commits: p.Commits,
		BornAt:  p.BornAt,
		DiedAt:  time.Now(),
		Cause:   cause,
	}
}

func GraveyardText(graves []GravePet) string {
	if len(graves) == 0 {
		return "\n  🌿 The graveyard is empty. All pets are alive and well! 🌿\n"
	}

	text := "\n  ⚰️  GitPet Graveyard  ⚰️\n"
	text += "  ========================\n\n"

	for i, g := range graves {
		ageDays := int(g.DiedAt.Sub(g.BornAt).Hours() / 24)
		text += fmt.Sprintf("  %d. %s (%s)\n", i+1, g.Name, g.Stage.Korean())
		text += fmt.Sprintf("     Lived: %d days | Commits: %d\n", ageDays, g.Commits)
		text += fmt.Sprintf("     Cause: %s\n", g.Cause)
		text += fmt.Sprintf("     %s - %s\n\n",
			g.BornAt.Format("2006-01-02"),
			g.DiedAt.Format("2006-01-02"))
	}

	return text
}

func IncrementStreak(p *Pet) {
	now := time.Now()
	lastFedDay := p.LastFed.Truncate(24 * time.Hour)
	today := now.Truncate(24 * time.Hour)

	daysDiff := today.Sub(lastFedDay).Hours() / 24

	if daysDiff <= 1 && daysDiff > 0 {
		p.Streak++
	} else if daysDiff > 1 {
		p.Streak = 1
	}
}
