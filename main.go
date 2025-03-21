package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

// Bot represents our Discord bot instance
type Bot struct {
	Session      *discordgo.Session
	Commands     map[string]Command
	Reminders    *ReminderManager
	Config       *Config
	ShutdownChan chan struct{}
}

// Command interface for all bot commands
type Command interface {
	Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error
	Help() string
}

// Config stores bot configuration
type Config struct {
	Token           string
	CommandPrefix   string
	Owners          []string
	DefaultTimezone string
}

// ReminderManager handles all reminder-related functionality
type ReminderManager struct {
	Reminders []Reminder
	mu        sync.Mutex
}

// Reminder represents a single reminder
type Reminder struct {
	ID        string
	UserID    string
	ChannelID string
	AuthorID  string
	Message   string
	EndTime   time.Time
	StartTime time.Time
	Repeats   int
	Period    int
	timer     *time.Timer
}

// NewBot creates a new bot instance
func NewBot() (*Bot, error) {
	config, err := loadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	session, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to create Discord session: %w", err)
	}

	// Set required intents
	session.Identify.Intents = discordgo.IntentsGuildMessages |
		discordgo.IntentsGuildMembers |
		discordgo.IntentsGuilds |
		discordgo.IntentsGuildMessageTyping |
		discordgo.IntentsGuildMessageReactions |
		discordgo.IntentsGuildPresences

	bot := &Bot{
		Session:      session,
		Commands:     make(map[string]Command),
		Reminders:    NewReminderManager(),
		Config:       config,
		ShutdownChan: make(chan struct{}),
	}

	// Register handlers
	session.AddHandler(bot.handleReady)
	session.AddHandler(bot.handleGuildCreate)
	session.AddHandler(bot.handleMessageCreate)
	session.AddHandler(bot.handleMessageUpdate)

	return bot, nil
}

// loadConfig loads the bot configuration from environment or files
func loadConfig() (*Config, error) {
	// Read token from file
	tokensFile, err := os.Open("tokens.txt")
	if err != nil {
		return nil, fmt.Errorf("could not open tokens file: %w", err)
	}
	defer tokensFile.Close()

	scanner := bufio.NewScanner(tokensFile)
	if !scanner.Scan() {
		return nil, fmt.Errorf("tokens file is empty")
	}

	token := scanner.Text()

	return &Config{
		Token:           token,
		CommandPrefix:   ".",
		Owners:          []string{"479126092330827777"}, // Replace with actual owner IDs
		DefaultTimezone: "Etc/GMT+0",
	}, nil
}

// NewReminderManager creates a new reminder manager
func NewReminderManager() *ReminderManager {
	return &ReminderManager{
		Reminders: []Reminder{},
		mu:        sync.Mutex{},
	}
}

// Start begins the bot's operation
func (b *Bot) Start(ctx context.Context) error {
	// Register all commands
	b.registerCommands()

	// Load saved reminders
	if err := b.Reminders.LoadReminders(); err != nil {
		return fmt.Errorf("failed to load reminders: %w", err)
	}

	// Open connection to Discord
	if err := b.Session.Open(); err != nil {
		return fmt.Errorf("failed to open connection: %w", err)
	}
	defer b.Session.Close()

	fmt.Printf("[%s] Bot started successfully\n", time.Now().Format(time.TimeOnly))

	// Start processing reminders in background
	go b.Reminders.ProcessReminders(b.Session)

	// Set up CLI input handling
	go b.handleCLIInput()

	// Wait for shutdown signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	select {
	case <-sc:
		fmt.Println("Received shutdown signal")
	case <-b.ShutdownChan:
		fmt.Println("Received shutdown request")
	case <-ctx.Done():
		fmt.Println("Context cancelled")
	}

	// Perform cleanup before shutdown
	return b.shutdown()
}

// shutdown performs graceful shutdown operations
func (b *Bot) shutdown() error {
	fmt.Printf("[%s] Shutting down...\n", time.Now().Format(time.TimeOnly))

	// Save reminders
	if err := b.Reminders.SaveReminders(); err != nil {
		return fmt.Errorf("failed to save reminders: %w", err)
	}

	// Close Discord connection
	return b.Session.Close()
}

// registerCommands registers all bot commands
func (b *Bot) registerCommands() {
	// Register commands here
	// b.Commands["help"] = &HelpCommand{bot: b}
	// b.Commands["roll"] = &RollCommand{}
	// b.Commands["remind"] = &RemindCommand{manager: b.Reminders}
	// etc.
}

// handleCLIInput handles console input for bot management
func (b *Bot) handleCLIInput() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		parts := strings.SplitN(input, " ", 2)

		switch parts[0] {
		case "exit", "quit", "gn":
			// Send goodbye message if needed
			b.Session.ChannelMessageSend("1331332284372222074",
				"Good night, Family. Tomorrow we shall take part in the banquet... again. For now, however, I will rest.")

			close(b.ShutdownChan)
			return

		case "say":
			if len(parts) < 2 {
				fmt.Println("Usage: say <message>")
				continue
			}
			b.Session.ChannelMessageSend("1331332284372222074", parts[1])

		case "chan":
			if len(parts) < 2 {
				fmt.Println("Usage: chan <channel_id>")
				continue
			}
			fmt.Printf("Set default channel to %s\n", parts[1])

			// Add more CLI commands as needed
		}
	}
}

// handleReady handles the ready event
func (b *Bot) handleReady(s *discordgo.Session, r *discordgo.Ready) {
	fmt.Printf("[%s] Connected to Discord as %s#%s\n",
		time.Now().Format(time.TimeOnly), s.State.User.Username, s.State.User.Discriminator)

	// Update bot status
	s.UpdateStatusComplex(discordgo.UpdateStatusData{
		Status: "online",
		Activities: []*discordgo.Activity{
			{
				Name: "Allow me to regale thee...",
				Type: discordgo.ActivityTypeCustom,
			},
		},
	})
}

// handleGuildCreate handles the guild create event
func (b *Bot) handleGuildCreate(s *discordgo.Session, g *discordgo.GuildCreate) {
	if g.Guild.Unavailable {
		return
	}

	fmt.Printf("[%s] Connected to guild: %s (ID: %s)\n",
		time.Now().Format(time.TimeOnly), g.Name, g.ID)

	// Update custom status with member count if it's the main guild
	if g.ID == "1250579779837493278" {
		memberCount := g.MemberCount - 2 // Adjust for bots
		status := fmt.Sprintf("Allow me to regale thee... that, in this... adventure of mine... Verily, I was blessed with a family of %d.", memberCount)
		s.UpdateStatusComplex(discordgo.UpdateStatusData{
			Status: "online",
			Activities: []*discordgo.Activity{
				{
					Name: status,
					Type: discordgo.ActivityTypeCustom,
				},
			},
		})
	}
}

// handleMessageCreate handles incoming messages
func (b *Bot) handleMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Process commands
	if strings.HasPrefix(m.Content, b.Config.CommandPrefix) {
		parts := strings.SplitN(strings.TrimPrefix(m.Content, b.Config.CommandPrefix), " ", 2)
		cmdName := strings.ToLower(parts[0])

		if cmd, ok := b.Commands[cmdName]; ok {
			var args []string
			if len(parts) > 1 {
				args = strings.Split(parts[1], " ")
			}

			if err := cmd.Execute(s, m, args); err != nil {
				log.Printf("Error executing command %s: %v", cmdName, err)
				s.ChannelMessageSendReply(m.ChannelID,
					"Sorry, my creator must have fucked something up.\nPlease pierce him with a sanguine lance and drink his blood.",
					m.Reference())
			}
		}
	}

	// Handle other message patterns
	content := strings.ToLower(m.Content)

	// Add other message handlers here
	// These can be moved to separate functions or a message handler system
}

// handleMessageUpdate handles edited messages
func (b *Bot) handleMessageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate) {
	// Handle message edits, focusing on the .roll command update behavior
	if strings.HasPrefix(m.Content, ".roll") {
		// Implementation for updating roll results
	}
}

// LoadReminders loads saved reminders from file
func (rm *ReminderManager) LoadReminders() error {
	// Implementation for loading reminders from timers.txt
	return nil
}

// SaveReminders saves all reminders to file
func (rm *ReminderManager) SaveReminders() error {
	// Implementation for saving reminders to timers.txt
	return nil
}

// ProcessReminders processes active reminders
func (rm *ReminderManager) ProcessReminders(s *discordgo.Session) {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		rm.checkReminders(s)
	}
}

// checkReminders checks if any reminders have triggered
func (rm *ReminderManager) checkReminders(s *discordgo.Session) {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	// Check each reminder
	// Implementation for checking and triggering reminders
}

// main function
func main() {
	// Create a new bot instance
	bot, err := NewBot()
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	// Create a context that can be cancelled
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start the bot
	if err := bot.Start(ctx); err != nil {
		log.Fatalf("Error running bot: %v", err)
	}
}
