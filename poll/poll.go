package poll

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
    // "github.com/discord_bot"

)
var (
	botID     string
	commandPrefix = "!"
)

type Poll struct {
    Title   string
    Options []string
    Votes   map[string]int // Map to store user votes for each option
}

var polls = make(map[string]Poll) // Map to store active polls

// Command to create a poll
func CreatePoll(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
    title := args[0]
    options := args[1:]

    // Create the poll with initial counters set to 0
    poll := Poll{
        Title:   title,
        Options: options,
        Votes:   make(map[string]int),
    }

    // Store the poll
    polls[m.ChannelID] = poll

    // Display the poll message with counters
    pollMessage := generatePollMessage(&poll)

    // Send the poll message
    s.ChannelMessageSend(m.ChannelID, pollMessage)
}
func generatePollMessage(poll *Poll) string {
    pollMessage := "Poll created!\n\n" + poll.Title + "\n"
    for i, option := range poll.Options {
        // Get the count for this option
        count := 0
        for _, vote := range poll.Votes {
            if vote == i {
                count++
            }
        }
        pollMessage += fmt.Sprintf("%d️⃣ %s - %d\n", i+1, option, count)
    }
    return pollMessage
}

// Function to handle reactions to the poll message
func handleReactionAdd(s *discordgo.Session, r *discordgo.MessageReactionAdd) {
    if r.UserID == botID {
        return
    }

    if poll, ok := polls[r.ChannelID]; ok {
        for i := range poll.Options {
            if r.Emoji.Name == fmt.Sprintf("%d️⃣", i+1) {
                poll.Votes[r.UserID] = i // Record the user's vote
                break
            }
        }
    }
}

