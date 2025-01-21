package main

import (
	"log"
	"math"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func roll(s *discordgo.Session, m *discordgo.MessageCreate) {
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
	if idk, err := strconv.Atoi(r); err == nil && idk > 0 {
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

		if countS == "" {
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
			if !sieveSimple.MatchString(rest) {
				iKnowWhatYouAre(s, c, m)
				return
			}
			max, _ = strconv.Atoi(rest)
		} else {
			modRune = strings.IndexAny(rest, "+*-^_")
			max, _ = strconv.Atoi(rest[:modRune])
			mod, _ = strconv.Atoi(rest[modRune+1:])
		}

		if max < 1 || count < 1 {
			iKnowWhatYouAre(s, c, m)
			return
		}

		rawStr := ""
		sum := 0
		for i := 0; i < count; i++ {
			v := rand.Intn(max) + 1
			rawStr += strconv.Itoa(v) + " "
			sum += v
		}

		if sieve.MatchString(rest) {
			// fmt.Println(string(rest[modRune]))
			switch rest[modRune] {
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

func iKnowWhatYouAre(s *discordgo.Session, c *discordgo.Channel, m *discordgo.MessageCreate) {
	s.ChannelMessageSendReply(c.ID, "I know what you are.", m.Reference())
}

func bod(s *discordgo.Session, m *discordgo.MessageCreate) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		return
	}
	resp, err := http.Get("https://cdn.discordapp.com/attachments/927117412770394135/1331006346925178932/Shi1.png?ex=67900bc2&is=678eba42&hm=ee9cae73f48a5b64a9499f7b927313f56131115d103a7367832044058e42576b&")
	if err != nil {
		log.Fatal("FUCK")
	}
	defer resp.Body.Close()
	yujin := resp.Body

	resp, err = http.Get("https://tiphereth.zasz.su/static/assets/cards_thumb/Roland4Phase_Yujin.jpg")
	if err != nil {
		log.Fatal("FUCK")
	}
	defer resp.Body.Close()
	yujinDead := resp.Body

	roll := rand.Intn(4) + 1
	if roll == 4 {
		s.ChannelMessageSendComplex(c.ID, &discordgo.MessageSend{
			Content: "**4**",
			Files: []*discordgo.File{
				{
					Name:   "yujin.png",
					Reader: yujin,
				},
			},
		})
	} else {
		s.ChannelMessageSendComplex(c.ID, &discordgo.MessageSend{
			Content: strconv.Itoa(roll),
			Files: []*discordgo.File{
				{
					Name:   "yujinDead.jpg",
					Reader: yujinDead,
				},
			},
		})
	}
}

func conceived(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendReply(m.ChannelID, "...What is it this time?", m.Reference())
}

func fut(s *discordgo.Session, m *discordgo.MessageCreate) {
	messages := []string{"...I won't even have to call Father for this.",
		"...Hold your tongue; I will no longer tolerate any more \"ingenious ideas\".",
		"You're nothing before a Second Kindred, let alone Father.",
		"...Did you learn that from that Knight? How humorous.",
	}
	s.ChannelMessageSendReply(m.ChannelID, messages[rand.Intn(len(messages))], m.Reference())
}