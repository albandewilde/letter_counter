package score

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/bwmarrin/discordgo"

	"github.com/albandewilde/letter_counter/discord_helpers"
)

func CalculateLevel(score int) string {
	switch {
	case score >= 1000000000:
		return "max"
	case score >= 900000000:
		return "29"
	case score >= 800000000:
		return "28"
	case score >= 700000000:
		return "27"
	case score >= 600000000:
		return "26"
	case score >= 500000000:
		return "25"
	case score >= 400000000:
		return "24"
	case score >= 300000000:
		return "23"
	case score >= 200000000:
		return "22"
	case score >= 100000000:
		return "21"
	case score >= 10000000:
		return "20"
	case score >= 9000000:
		return "19"
	case score >= 7000000:
		return "18"
	case score >= 5000000:
		return "17"
	case score >= 4000000:
		return "16"
	case score >= 3000000:
		return "15"
	case score >= 2000000:
		return "14"
	case score >= 1000000:
		return "13"
	case score >= 100000:
		return "12"
	case score >= 50000:
		return "11"
	case score >= 10000:
		return "10"
	case score >= 7750:
		return "9"
	case score >= 5500:
		return "8"
	case score >= 3250:
		return "7"
	case score >= 1000:
		return "6"
	case score >= 500:
		return "5"
	case score >= 100:
		return "4"
	case score >= 50:
		return "3"
	case score >= 10:
		return "2"
	case score >= 1:
		return "1"
	}
	return "0"
}

func userRank(username string) int {
	return 0
}

func readScores() (map[string]map[string]int, error) {
	scoresFile, err := ioutil.ReadFile("scores.json")

	// Check if the file already exist
	// if not, we create it with a empty json object in it
	if os.IsNotExist(err) {
		// Try to write the file
		err = ioutil.WriteFile("scores.json", []byte("{}"), 0665)
		if err != nil {
			return nil, err
		} else {
			// Re-read file with the content created
			scoresFile, err = ioutil.ReadFile("scores.json")
			if err != nil {
				return nil, err
			}
		}
	}

	// Load the json with score
	var scores map[string]map[string]int
	err = json.Unmarshal(scoresFile, &scores)
	if err != nil {
		return nil, err
	}

	return scores, nil
}

func writeScores(scores map[string]map[string]int) error {
	// Convert scores to json
	scoresBytes, err := json.Marshal(scores)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("scores.json", scoresBytes, 0665)
	if err != nil {
		return err
	}

	return nil
}

func SaveMessageScore(author *discordgo.User, messageLength int) error {
	// Read scores
	scores, err := readScores()
	if err != nil {
		return err
	}

	user := discord_helpers.DiscordUserCompleteName(author)

	// Create the user if he doesn't exist
	if scores[user] == nil {
		scores[user] = map[string]int{
			"char": 0,
			"msg":  0,
		}
	}

	// Upate the score of user
	scores[user]["char"] += messageLength
	scores[user]["msg"] += 1

	// Write the new score
	err = writeScores(scores)
	if err != nil {
		return err
	}

	return nil
}

func ReadUserScore(author *discordgo.User) (map[string]int, error) {
	var scores map[string]map[string]int
	var err error
	scores, err = readScores()
	return scores[discord_helpers.DiscordUserCompleteName(author)], err
}
