package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"syscall"
	"os/signal"
	"log"
	"bufio"
	"time"
)

func main(){
	fmt.Println("I shall pronounce the bot started.")

	tokens, err := os.Open("tokens.txt")
    if err != nil {
        log.Fatal(err)
    }
	defer tokens.Close()

	scanner := bufio.NewScanner(tokens)
	var auth_token string
	if scanner.Scan() {
		auth_token = scanner.Text()
	} else {
		log.Fatal(err)
	}

	discord, err := discordgo.New("Bot " + auth_token)

	if err != nil {
		log.Fatal(err)
	}

	discord.Identify.Intents = 335666240
	discord.Identify.Intents |= discordgo.IntentGuildMembers
	discord.Identify.Intents |= discordgo.IntentGuilds
	discord.Identify.Intents |= discordgo.IntentGuildMessages
	discord.Identify.Intents |= discordgo.IntentGuildMessageTyping
	discord.AddHandler(guildCreate)

	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	fmt.Printf("[%s] The onus has fallen onto me.\n", time.Now().Format(time.TimeOnly))

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	discord.Close()
}

func guildCreate(s *discordgo.Session, m *discordgo.GuildCreate){
	if m.Guild.Unavailable {
		return
	}

	channels := m.Guild.Channels

	fmt.Println("Joined server "+m.Guild.Name)
	for i := 0; i<len(channels); i++ {
		perms, _ := s.State.UserChannelPermissions("1330935741018276022", channels[i].ID)
		if channels[i].Type == 0 && (perms & 2048 == 2048) {
			s.ChannelMessageSend(channels[0].ID, "The Server will be well-cared for.\n...After all, the onus always fell on me to give roles that you abandoned.")
			return
		}
	}
}

// func ready(s *discordgo.Session, m *discordgo.Ready){
// 	if len(m.Guilds) == 0 {
// 		return
// 	}
// 	server := m.Guilds[0].ID
// 	channels, err := s.GuildChannels(server)

// 	if err != nil {
// 		fmt.Println("Something went wrong: "+err.Error())
// 	}

// 	for i := 0; i<len(channels); i++ {
// 		fmt.Println(channels[i].PermissionOverwrites)
// 	}

// 	s.ChannelMessageSend(channels[0].ID, "Rebellion comes to La Manchaland.")
// }