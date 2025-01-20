package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"
	"math"

	"github.com/bwmarrin/discordgo"
)

var yujin, yujinDead io.ReadCloser

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

	resp, err := http.Get("https://cdn.discordapp.com/attachments/927117412770394135/1331006346925178932/Shi1.png?ex=67900bc2&is=678eba42&hm=ee9cae73f48a5b64a9499f7b927313f56131115d103a7367832044058e42576b&")
	if err != nil{
		log.Fatal("FUCK")
	}
	defer resp.Body.Close()
	yujin = resp.Body

	resp, err = http.Get("https://tiphereth.zasz.su/static/assets/cards_thumb/Roland4Phase_Yujin.jpg")
	if err != nil{
		log.Fatal("FUCK")
	}
	defer resp.Body.Close()
	yujinDead = resp.Body

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
		if channels[i].Type == 0 && (perms & 2048 == 2048) && time.Now().Unix()-m.JoinedAt.Unix() < 30{
			s.ChannelMessageSend(channels[i].ID, "The Server will be well-cared for.\n...After all, the onus always fell on me to give roles that you abandoned.")
			return
		}
	}

	if m.Guild.ID == "1250579779837493278" {
		s.UpdateCustomStatus("Allow me to regale thee... that, in this... adventure of mine... Verily, I was blessed with a family of " + strconv.Itoa(m.Guild.MemberCount-1) + ".")
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
	re := regexp.MustCompile("fuck|shit|ass|idiot|dumb|stupid|clanker")
	if strings.HasPrefix(m.Content, ".roll ") {
		roll(s, m)
	} else if m.Content == ".bod" {
		bod(s, m)
	} else if re.MatchString(m.Content) && (refid == s.State.User.ID || strings.Contains(strings.ToLower(m.Content), "sancho")){
		fut(s,m)
	} else if strings.Contains(m.Content, "conceived") && m.Author.ID == "530516460712361986"{
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
		sieve := regexp.MustCompile(`[0-9]+[\+\*\-\^_][0-9]+`)
		var modRune int
		if !sieve.MatchString(rest) {
			if !sieveSimple.MatchString(rest){
				iKnowWhatYouAre(s, c, m)
				return
			}
			max,_ = strconv.Atoi(rest)
		} else {
			modRune = strings.IndexAny(rest, "+*-^_")
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
			// fmt.Println(string(rest[modRune]))
			switch rest[modRune]{
				case '+':
					sum += mod
				case '*':
					sum *= mod
				case '-':
					sum -= mod
				case '^':
					sum = int(math.Pow(float64(sum), float64(mod)))
				case '_':
					sum = int(math.Pow(float64(mod), float64(sum)))
			}
		}

		s.ChannelMessageSendReply(c.ID, "Your roll is "+strconv.Itoa(sum)+" ("+rawStr[:len(rawStr)-1]+").", m.Reference())
	}
}

func iKnowWhatYouAre(s *discordgo.Session, c *discordgo.Channel, m *discordgo.MessageCreate){
	s.ChannelMessageSendReply(c.ID, "I know what you are.", m.Reference())
}

func bod(s *discordgo.Session, m *discordgo.MessageCreate){
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		return
	}
	resp, err := http.Get("https://cdn.discordapp.com/attachments/927117412770394135/1331006346925178932/Shi1.png?ex=67900bc2&is=678eba42&hm=ee9cae73f48a5b64a9499f7b927313f56131115d103a7367832044058e42576b&")
	if err != nil{
		log.Fatal("FUCK")
	}
	defer resp.Body.Close()
	yujin = resp.Body

	resp, err = http.Get("https://tiphereth.zasz.su/static/assets/cards_thumb/Roland4Phase_Yujin.jpg")
	if err != nil{
		log.Fatal("FUCK")
	}
	defer resp.Body.Close()
	yujinDead = resp.Body
	
	roll := rand.Intn(4)+1
	if roll == 4{
		s.ChannelMessageSendComplex(c.ID, &discordgo.MessageSend{
			Content: "**4**",
			Files: []*discordgo.File{
				{
					Name: "yujin.png",
					Reader: yujin,
				},
			},
		},)
	} else {
		s.ChannelMessageSendComplex(c.ID, &discordgo.MessageSend{
			Content: strconv.Itoa(roll),
			Files: []*discordgo.File{
				{
					Name: "yujinDead.jpg",
					Reader: yujinDead,
				},
			},
		},)
	}
}

func conceived(s *discordgo.Session, m *discordgo.MessageCreate){
	s.ChannelMessageSendReply(m.ChannelID, "...What is it this time?", m.Reference())
}

func fut(s *discordgo.Session, m *discordgo.MessageCreate){
	messages := []string{"...I won't even have to call Father for this.",
	"...Hold your tongue; I will no longer tolerate any more \"ingenious ideas\".",
	"You're nothing before a Second Kindred, let alone Father.",
	"...Did you learn that from that Knight? How humorous.",
	}
	s.ChannelMessageSendReply(m.ChannelID, messages[rand.Intn(len(messages))], m.Reference())
}