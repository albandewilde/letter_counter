package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Score struct {
	USER     string // username#discriminator
	CHAR     int64
	CHAR_LVL string
	MSG      int64
	MSG_LVL  string
	RATIO    float64
	RANK     int
}

func main() {
	var discordToken string = readDiscordToken()

	// Create a new Discord session using the provided bot token.
	bot, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		fmt.Println("Error while creating the Discord session,", err)
		return
	}

	// Register funcs callback.
	bot.AddHandler(getScore)
	bot.AddHandler(counting)

	// Open a websocket connection to Discord and begin listening.
	err = bot.Open()
	if err != nil {
		fmt.Println("Error while opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("I'm logged in ! (Press CTRL-C to exit.)\n")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	bot.Close()
}

func readDiscordToken() (token string) {
	// Read token in the `secrets.json` file
	secretFile, err := ioutil.ReadFile("./secrets.json")
	if err != nil {
		fmt.Println("Error while reading secrets:", err)
	}

	type Secrets struct {
		DISCORD string
	}

	var secrets Secrets

	// Parse json content
	err = json.Unmarshal(secretFile, &secrets)
	if err != nil {
		fmt.Println("Error while parsing secrets:", err)
	}

	token = secrets.DISCORD
	return
}

func getScore(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println(
		"\033[36m",
		time.Now().Format("[2006-01-02 15:04:05]"),
		"\033[0",
		":",
		"\033[35m",
		m.Author.Username,
		"#",
		m.Author.Discriminator,
		"\033[0m",
		"want to know his score.",
	)

	username := m.Author.Username + "#" + m.Author.Discriminator
	var score Score = calculateScoreOfUser(username)

	response := "```txt\n" +
		fmt.Sprintf("User: %s\n\n", score.USER) +
		fmt.Sprintf("Written characters: %d\nCaracter level: %s\n\n", score.CHAR, score.CHAR_LVL) +
		fmt.Sprintf("Messages sent: %d\nMessage level: %s\n\n", score.MSG, score.MSG_LVL) +
		fmt.Sprintf("Ratio (written caracters/messages sent): %.2f\n\n", score.RATIO) +
		fmt.Sprintf("Rank: #%d\n\n", score.RANK) +
		"```"

	s.ChannelMessageSend(m.ChannelID, response)

}

func counting(s *discordgo.Session, m *discordgo.MessageCreate) {
	messageLength := len(m.Content)

	fmt.Println(
		"\033[36m",
		time.Now().Format("[2006-01-02 15:04:05]"),
		"\033[0",
		":",
		"The user",
		"\033[35m",
		m.Author.Username,
		"#",
		m.Author.Discriminator,
		"\033[0m",
		"write a message with",
		messageLength,
		"characters.",
	)

	saveMessageScore(m.Author, messageLength)
}

func calculateScoreOfUser(username string) Score {
	// Read user characters and mesages send
	var userScore map[string]string = readUserScore(username)

	// Convert value into int
	char, err := strconv.ParseInt(userScore["char"], 10, 64)
	msg, err := strconv.ParseInt(userScore["msg"], 10, 64)

	// Calculate the user rank
	var rank int = userRank(username)

	return Score{
		USER:     username,
		CHAR:     char,
		CHAR_LVL: calculateLevel(char),
		MSG:      msg,
		MSG_LVL:  calculateLevel(msg),
		RATIO:    float64(char) / float64(msg),
		RANK:     rank,
	}
}

func calculateLevel(score int64) string {
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
