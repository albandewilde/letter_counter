package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/albandewilde/letter_counter/discord_helpers"
	"github.com/albandewilde/letter_counter/score"
)

type Score struct {
	USER     string // username#discriminator
	CHAR     int
	CHAR_LVL string
	MSG      int
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
	bot.AddHandler(counting)
	bot.AddHandler(getScore)

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

	// To know his score, the message content must be `§score`
	if m.Content != "§score" {
		return
	}

	fmt.Print(
		"\033[36m",
		time.Now().Format("[2006-01-02 15:04:05]"),
		"\033[0m",
		": ",
		" The user ",
		"\033[35m",
		m.Author.Username+"#"+m.Author.Discriminator,
		"\033[0m",
		" want to know his score.\n",
	)

	var score Score
	var err error
	score, err = calculateScoreOfUser(m.Author)

	var response string
	if err != nil {
		displayError(err)
		response = "Sorry, something is wrong."
	} else {

		response = "```txt\n" +
			fmt.Sprintf("User: %s\n\n", score.USER) +
			fmt.Sprintf("Written characters: %d\nCaracter level: %s\n\n", score.CHAR, score.CHAR_LVL) +
			fmt.Sprintf("Messages sent: %d\nMessage level: %s\n\n", score.MSG, score.MSG_LVL) +
			fmt.Sprintf("Ratio (written caracters/messages sent): %.2f\n\n", score.RATIO) +
			fmt.Sprintf("Rank: #%d\n\n", score.RANK) +
			"```"
	}

	s.ChannelMessageSend(m.ChannelID, response)

}

func counting(s *discordgo.Session, m *discordgo.MessageCreate) {
	messageLength := len(m.Content)

	fmt.Print(
		"\033[36m",
		time.Now().Format("[2006-01-02 15:04:05]"),
		"\033[0m",
		": ",
		" The user ",
		"\033[35m",
		m.Author.Username+"#"+m.Author.Discriminator,
		"\033[0m",
		" write a message with ",
		"\033[33m",
		messageLength,
		"\033[0m",
		" characters.\n",
	)

	err := score.SaveMessageScore(m.Author, messageLength)
	if err != nil {
		displayError(err)
	}
}

func calculateScoreOfUser(username *discordgo.User) (Score, error) {
	// Read user characters and mesages send
	var userScore map[string]int
	var err error
	userScore, err = score.ReadUserScore(username)
	if err != nil {
		return Score{}, err
	}

	// Calculate the user rank
	rank, err := score.UserRank(discord_helpers.DiscordUserCompleteName(username))
	if err != nil {
		rank = 0
		displayError(err)
	}

	return Score{
		USER:     discord_helpers.DiscordUserCompleteName(username),
		CHAR:     userScore["char"],
		CHAR_LVL: score.CalculateLevel(userScore["char"]),
		MSG:      userScore["msg"],
		MSG_LVL:  score.CalculateLevel(userScore["msg"]),
		RATIO:    float64(userScore["char"]) / float64(userScore["msg"]),
		RANK:     rank,
	}, nil
}

func displayError(err error) {
	fmt.Print(
		"\033[31m",
		"Error: ",
		"\033[0m",
		err,
		"\n",
	)
}
