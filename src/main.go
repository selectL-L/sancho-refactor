package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

func main(){
	fmt.Printf("[%s] I shall pronounce the bot started.\n", time.Now().Format(time.TimeOnly))

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
	discord.Identify.Intents |= discordgo.IntentGuildMessageReactions
	discord.AddHandler(guildCreate)
	discord.AddHandler(ready)
	discord.AddHandler(messageCreate)
	defer discord.Close()

	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	fmt.Printf("[%s] The onus has fallen onto me.\n", time.Now().Format(time.TimeOnly))
	discord.MessageReactionAdd("1324527612042678393", "1324541333662204026", "🩸")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func guildCreate(s *discordgo.Session, m *discordgo.GuildCreate){
	if m.Guild.Unavailable {
		return
	}

	channels := m.Guild.Channels

	fmt.Println("Joined server "+m.Guild.Name)
	for i := 0; i<len(channels); i++ {
		perms, _ := s.State.UserChannelPermissions("1330935741018276022", channels[i].ID)
		if channels[i].Type == 0 && (perms & 2048 == 2048) && time.Now().Unix()-m.JoinedAt.Unix() < 30{
			s.ChannelMessageSend(channels[i].ID, "The Server will be well-cared for.\n...After all, the onus always fell on me to give roles that you abandoned.")
			return
		}
	}

	if m.Guild.ID == "1250579779837493278" {
		s.UpdateCustomStatus("Allow me to regale thee... that, in this... adventure of mine... Verily, I was blessed with a family of " + strconv.Itoa(m.Guild.MemberCount-2) + ".")
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate){
	if m.Author.ID == s.State.User.ID {
		return
	}
	var refid string
	if m.ReferencedMessage != nil {
		refid = m.ReferencedMessage.Author.ID
	}
	re := regexp.MustCompile("fuck|shit|ass|idiot|dumb|stupid|clanker|bitch")
	lowerMsg := strings.ToLower(m.Content)
	if strings.HasPrefix(m.Content, ".roll ") {
		roll(s, m)
	} else if m.Content == ".bod" {
		bod(s, m)
	} else if strings.Contains(lowerMsg,"kiss") && (refid == s.State.User.ID || strings.Contains(strings.ToLower(m.ContentWithMentionsReplaced()), "sancho")){
		s.ChannelMessageSendReply(m.ChannelID, "...Maybe.", m.Reference())
	} else if re.MatchString(lowerMsg) && (refid == s.State.User.ID || strings.Contains(strings.ToLower(m.ContentWithMentionsReplaced()), "sancho"))  && m.Author.ID != "530516460712361986"{
		fut(s,m)
	} else if strings.Contains(lowerMsg, "conceived") && m.Author.ID == "530516460712361986"{
		conceived(s,m)
	}
}

func ready(s *discordgo.Session, m *discordgo.Ready){
	server, err := s.State.Guild("1250579779837493278")
	var num int
	if err != nil{
		num = 13
	} else {
		num = server.MemberCount
	}
	s.UpdateCustomStatus("Allow me to regale thee... that, in this... adventure of mine... Verily, I was blessed with a family of " + strconv.Itoa(num-1) + ".")
}
