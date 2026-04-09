package pet

import "math/rand/v2"

var adjectives = []string{
	"Tiny", "Fuzzy", "Sleepy", "Happy", "Brave",
	"Lucky", "Cosmic", "Pixel", "Turbo", "Fluffy",
	"Mighty", "Sneaky", "Bubbly", "Wiggly", "Zippy",
	"Sparkly", "Cozy", "Bouncy", "Crispy", "Cloudy",
}

var nouns = []string{
	"Paw", "Bean", "Boop", "Noodle", "Mochi",
	"Tofu", "Dumpling", "Nugget", "Pickle", "Waffle",
	"Sprout", "Peach", "Cookie", "Pudding", "Bubble",
	"Taco", "Pretzel", "Muffin", "Biscuit", "Pebble",
}

func RandomName() string {
	adj := adjectives[rand.IntN(len(adjectives))]
	noun := nouns[rand.IntN(len(nouns))]
	return adj + noun
}
