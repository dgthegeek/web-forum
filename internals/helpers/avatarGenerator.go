package helpers

import (
	"fmt"
	"math/rand"
)

func AvatarGenerator() string {
	names := []string{"Willow", "Abby", "Angel", "Annie",
		"Baby", "Bella", "Bandit", "Whiskers",
		"Zoe", "Rocky", "Bailey", "Zoey",
		"Bear", "Peanut", "Samantha", "Shadow",
		"Sam", "Buster", "Boo", "Buddy",
	}
	randomSeed := rand.Intn(len(names)-1) + 1
	avatarURL := fmt.Sprintf("https://api.dicebear.com/6.x/bottts/svg?seed=%s", names[randomSeed])

	return avatarURL
}
