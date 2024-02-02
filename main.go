package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	botID         string
	commandPrefix = "!"
	reminders     = make(map[string]time.Time)
)

const (
	Token = "YOUR_DISCORD_BOT_TOKEN_HERE"
)

func main() {
	discord, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	discord.AddHandler(messageCreate)

	err = discord.Open()
	if err != nil {
		fmt.Println("Error opening Discord session:", err)
		return
	}

	defer discord.Close()

	// Start a goroutine to continuously check for reminders
	go func() {
		for {
			checkReminders(discord)
			time.Sleep(1 * time.Minute) // Check reminders every minute
		}
	}()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	select {}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Spawn a goroutine to handle each incoming message asynchronously
	go func() {
		if m.Author.ID == botID {
			return
		}

		if strings.HasPrefix(m.Content, commandPrefix) {
			fields := strings.Fields(m.Content[len(commandPrefix):])
			command := fields[0]
			args := fields[1:]

			switch command {
			case "ping":
				s.ChannelMessageSend(m.ChannelID, "Pong!")
			case "repeat":
				if len(args) < 1 {
					s.ChannelMessageSend(m.ChannelID, "Please provide a message to repeat.")
					return
				}
				s.ChannelMessageSend(m.ChannelID, strings.Join(args, " "))
			case "remindme":
				if len(args) < 2 {
					s.ChannelMessageSend(m.ChannelID, "Please provide a reminder message and duration.")
					return
				}

				reminderMessage := strings.Join(args[:len(args)-1], " ")
				durationString := args[len(args)-1]
				duration, err := time.ParseDuration(durationString)
				if err != nil {
					s.ChannelMessageSend(m.ChannelID, "Invalid duration format. Please provide a valid duration (e.g., 1h30m).")
					return
				}

				reminderTime := time.Now().Add(duration)
				reminders[reminderMessage] = reminderTime

				s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Reminder set: '%s' at %s", reminderMessage, reminderTime.Format("15:04:05 MST")))
			case "help":
				helpMessage := "Available commands:\n" +
					"!ping - Responds with 'Pong!'\n" +
					"!repeat <message> - Repeats the provided message\n" +
					"!remindme <message> <duration> - Sets a reminder\n" +
					"!help - Displays this help message"
				s.ChannelMessageSend(m.ChannelID, helpMessage)
			default:
				// Handle unknown command
				s.ChannelMessageSend(m.ChannelID, "Unknown command. Type !help to see available commands.")
			}
		}
	}()
}

func checkReminders(s *discordgo.Session) {
	currentTime := time.Now()
	for reminderMessage, reminderTime := range reminders {
		if currentTime.After(reminderTime) {
			// Send reminder message
			user, err := s.User(botID)
			if err != nil {
				fmt.Println("Error getting user information:", err)
				return
			}
			channel, err := s.UserChannelCreate(user.ID)
			if err != nil {
				fmt.Println("Error creating user channel:", err)
				return
			}
			s.ChannelMessageSend(channel.ID, reminderMessage)
			// Remove reminder from map
			delete(reminders, reminderMessage)
		}
	}
}
