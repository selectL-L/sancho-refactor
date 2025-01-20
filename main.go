package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
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
	discord.AddHandler(ready)
	discord.AddHandler(messageCreate)
	defer discord.Close()

	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	fmt.Printf("[%s] The onus has fallen onto me.\n", time.Now().Format(time.TimeOnly))

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
		if channels[i].Type == 0 && (perms & 2048 == 2048) {
			s.ChannelMessageSend(channels[i].ID, "The Server will be well-cared for.\n...After all, the onus always fell on me to give roles that you abandoned.")
			return
		}
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate){
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.HasPrefix(m.Content, ".roll ") {
		roll(s, m)
	}
	if strings.HasPrefix(m.Content, ".bod ") {
		bod(s, m)
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




func roll(s *discordgo.Session, m *discordgo.MessageCreate){
	c, err := s.State.Channel(m.ChannelID)
		if err != nil {
			return
		}
		// _, err := s.State.Guild(c.GuildID)
		// if err != nil {
		// 	return
		// }
		
		var count, max, mod int
		r, _ := strings.CutPrefix(m.Content, ".roll ")
		if idk, err := strconv.Atoi(r); err == nil && idk>0{
			count = 1
			max, _ = strconv.Atoi(r)
			mod = 0
			s.ChannelMessageSendReply(c.ID, "Your roll is "+strconv.Itoa(int(rand.Int63n(int64(max))+1))+".", m.Reference())
		} else {
			countS, rest, found := strings.Cut(r, "d")
			if !found {
				iKnowWhatYouAre(s, c, m)
				return
			}
			
			if countS == ""{
				count = 1
			} else {
				count, err = strconv.Atoi(countS)
				if err != nil {
					iKnowWhatYouAre(s, c, m)
					return
				}
			}

			sieveSimple := regexp.MustCompile("[0-9]+")
			sieve := regexp.MustCompile("[0-9]+[+*-][0-9]+")
			var modRune rune
			if !sieve.MatchString(rest) {
				if !sieveSimple.MatchString(rest){
					iKnowWhatYouAre(s, c, m)
					return
				}
				max,_ = strconv.Atoi(rest)
			} else {
				modRune := strings.IndexAny(rest, "+*-")
				max,_ = strconv.Atoi(rest[:modRune])
				mod,_ = strconv.Atoi(rest[modRune+1:])
			}

			if max<1 || count<1 {
				iKnowWhatYouAre(s,c,m)
				return
			}

			rawStr := ""
			sum := 0
			for i := 0; i<count; i++ {
				v := rand.Intn(max)+1
				rawStr += strconv.Itoa(v) + " "
				sum += v
			}

			if sieve.MatchString(rest){
				switch rest[modRune]{
					case '+':
						sum += mod
					case '*':
						sum *= mod
					case '-':
						sum -= mod
				}
			}

			s.ChannelMessageSendReply(c.ID, "Your roll is "+strconv.Itoa(sum)+" ("+rawStr[:len(rawStr)-1]+").", m.Reference())
		}
}

func iKnowWhatYouAre(s *discordgo.Session, c *discordgo.Channel, m *discordgo.MessageCreate){
	s.ChannelMessageSendReply(c.ID, "I know what you are.", m.Reference())
}