package pet

import "time"

type Stage int

const (
	Egg      Stage = iota // 알
	Baby                  // 아기
	Teen                  // 청소년
	Adult                 // 성체
	Legend                // 전설
)

func (s Stage) String() string {
	return [...]string{"Egg", "Baby", "Teen", "Adult", "Legend"}[s]
}

func (s Stage) Korean() string {
	return [...]string{"🥚 알", "🐣 아기", "🐥 청소년", "🐔 성체", "🌟 전설"}[s]
}

// XP required to evolve to next stage
func (s Stage) XPToEvolve() int {
	return [...]int{10, 30, 70, 150, 0}[s]
}

type Mood int

const (
	Happy   Mood = iota
	Normal
	Hungry
	Sad
	Dead
)

func (m Mood) String() string {
	return [...]string{"Happy", "Normal", "Hungry", "Sad", "Dead"}[m]
}

func (m Mood) Emoji() string {
	return [...]string{"😊", "😐", "😿", "😢", "💀"}[m]
}

type Pet struct {
	Name      string    `json:"name"`
	Stage     Stage     `json:"stage"`
	XP        int       `json:"xp"`
	Hunger    int       `json:"hunger"`    // 0-100, 100 = starving
	Happiness int       `json:"happiness"` // 0-100, 100 = max happy
	Commits   int       `json:"commits"`
	Streak    int       `json:"streak"`    // consecutive days with commits
	BornAt    time.Time `json:"born_at"`
	LastFed   time.Time `json:"last_fed"`
	LastCheck time.Time `json:"last_check"`
	Alive     bool      `json:"alive"`
}

type GravePet struct {
	Name    string    `json:"name"`
	Stage   Stage     `json:"stage"`
	Commits int       `json:"commits"`
	BornAt  time.Time `json:"born_at"`
	DiedAt  time.Time `json:"died_at"`
	Cause   string    `json:"cause"`
}

type State struct {
	CurrentPet *Pet       `json:"current_pet"`
	Graveyard  []GravePet `json:"graveyard"`
}
