package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/gographics/imagick.v3/imagick"
)

var err error

// call if something goes wrong
func sadness(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m != nil{
		s.ChannelMessageSendReply(m.ChannelID, "Sorry, my creator must have fucked something up.\nPlease pierce him with a sanguine lance and drink his blood.", m.Reference())
	}
	fmt.Println(err)
}

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
			iKnowWhatYouAre(s, m)
			return
		}

		if countS == "" {
			count = 1
		} else {
			count, err = strconv.Atoi(countS)
			if err != nil {
				iKnowWhatYouAre(s, m)
				return
			}
		}

		modRune := strings.IndexAny(rest, "+-*^_")
		if modRune == -1 {
			modRune = len(rest)
		}

		max, err = strconv.Atoi(rest[:modRune])
		if err != nil {
			iKnowWhatYouAre(s, m)
			return
		}

		if max < 1 || count < 1 {
			fmt.Println(max, count)
			iKnowWhatYouAre(s, m)
			return
		}

		rawStr := ""
		sum := 0
		for i := 0; i < count; i++ {
			v := rand.Intn(max) + 1
			rawStr += strconv.Itoa(v) + " "
			sum += v
		}

		for modRune < len(rest) {
			// the following code is magic! don't touch
			rest = rest[modRune:]
			sign := rest[0]
			modRune = strings.IndexAny(rest[1:], "+-*^_") + 1
			if modRune == 0 {
				modRune = len(rest)
			}
			mod, err = strconv.Atoi(rest[1:modRune])
			if err != nil {
				iKnowWhatYouAre(s, m)
				return
			}

			switch sign {
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

		out := "Your roll is " + strconv.Itoa(sum) + " (" + rawStr[:len(rawStr)-1] + ")."
		if len(out) > 2000 {
			s.ChannelMessageSendReply(c.ID, "Your roll is "+strconv.Itoa(sum)+" .", m.Reference())
			return
		}
		s.ChannelMessageSendReply(c.ID, out, m.Reference())
	}
}

// call if user fucked up command use
func iKnowWhatYouAre(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSendReply(m.ChannelID, "I know what you are.", m.Reference())
}

func bod(s *discordgo.Session, m *discordgo.MessageCreate) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		return
	}
	resp, err := http.Get("https://tiphereth.zasz.su/static/assets/cards_thumb/Shi1.jpg")
	if err != nil {
		sadness(s, m)
	}
	defer resp.Body.Close()
	yujin := resp.Body

	resp, err = http.Get("https://tiphereth.zasz.su/static/assets/cards_thumb/Roland4Phase_Yujin.jpg")
	if err != nil {
		sadness(s, m)
	}
	defer resp.Body.Close()
	yujinDead := resp.Body

	pingId, _ := strings.CutPrefix(m.Content, ".bod ")
	//pingId = pingId[1:len(pingId)-1]
	ping := ""
	if _, err := strconv.Atoi(pingId[2 : len(pingId)-1]); pingId[1] == '@' && err == nil {
		ping = pingId
	}
	roll := rand.Intn(4) + 1
	if roll == 4 {
		s.ChannelMessageSendComplex(c.ID, &discordgo.MessageSend{
			Content: "**4**" + " " + ping,
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

func nacho(s *discordgo.Session, m *discordgo.MessageCreate) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		return
	}
	// resp, err := http.Get("https://cdn.discordapp.com/attachments/1136333577643102259/1331326711467475034/GfVZ6foXsAAXfdY.jpg?ex=6791361e&is=678fe49e&hm=d86cc63397b25dfb33ca7fec58287d8f23f177523aaf0fa0d0bcefa1a1145fb0&")
	// if err != nil {
	// 	sadness(s, m)
	// }
	// defer resp.Body.Close()
	// nacho := resp.Body
	nacho, err := os.Open("img/nacho.jpg")
	if err!=nil{
		sadness(s,m)
	}
	s.ChannelMessageSendComplex(c.ID, &discordgo.MessageSend{
		Files: []*discordgo.File{
			{
				Name:   "nacho.jpg",
				Reader: nacho,
			},
		},
	})
}

func badword(s *discordgo.Session, m *discordgo.MessageCreate) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		return
	}
	// resp, err := http.Get("https://cdn.discordapp.com/attachments/1136333577643102259/1331326712113270914/qxckfwvd8vzd1.gif?ex=6791361f&is=678fe49f&hm=c1cc119f1e91ea70e34a766992133b4fcc392957db3b4e25a7cfd44fb25df0e5&")
	// if err != nil {
	// 	sadness(s, m)
	// }
	// defer resp.Body.Close()
	// img := resp.Body
	img, err := os.Open("img/badword.jpg")
	if err!=nil{
		sadness(s,m)
	}
	s.ChannelMessageSendComplex(c.ID, &discordgo.MessageSend{
		Files: []*discordgo.File{
			{
				Name:   "badword.gif",
				Reader: img,
			},
		},
	})
}

func rye(s *discordgo.Session, m *discordgo.MessageCreate) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		return
	}
	resp, err := http.Get("https://cdn.discordapp.com/attachments/1192614939689484359/1331063478290616470/attachment.gif?ex=6790e9b7&is=678f9837&hm=6a31133ec49b8c54e3f91bc373e5db1ca94cb065450e8da9455a801395222d6a&")
	if err != nil {
		sadness(s, m)
	}
	defer resp.Body.Close()
	img := resp.Body
	s.ChannelMessageSendComplex(c.ID, &discordgo.MessageSend{
		Files: []*discordgo.File{
			{
				Name:   "rye.gif",
				Reader: img,
			},
		},
	})
}

func jpegify(s *discordgo.Session, m *discordgo.MessageCreate, orb *imagick.MagickWand, quality int) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		return
	}
	var resp *http.Response
	if len(m.Attachments) == 0 || !strings.Contains(m.Attachments[0].ContentType, "image") {
		if m.ReferencedMessage == nil {
			s.ChannelMessageSendReply(m.ChannelID, "Please send an actual image.", m.Reference())
			return
		}
		if len(m.ReferencedMessage.Attachments) == 0 || !strings.Contains(m.ReferencedMessage.Attachments[0].ContentType, "image") {
			s.ChannelMessageSendReply(m.ChannelID, "Please send an actual image.", m.Reference())
			return
		}
		resp, err = http.Get(m.ReferencedMessage.Attachments[0].URL)
	} else {
		resp, err = http.Get(m.Attachments[0].URL)
	}

	//now we pipe the first (for now) image into imagemagick and wait for the result - how?

	//fn := m.Attachments[0].Filename
	//ext := fn[strings.LastIndex(fn, ".")+1:]
	if err != nil {
		fmt.Println("couldn't get image from internet")
		sadness(s, m)
		return
	}
	defer resp.Body.Close()
	orig, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("couldn't extract from html")
		sadness(s, m)
		return
	}

	err = orb.ReadImageBlob(orig)
	defer orb.Clear()
	if err != nil {
		fmt.Println("couldn't read img into orb", err)
		sadness(s, m)
		return
	}
	orb.SetImageFormat("JPEG")
	if quality < 2 {
		x, y := orb.GetImageWidth(), orb.GetImageHeight()
		scalingFactor := math.Max(float64(x/160), float64(y/100))

		orb.ResizeImage(uint(float64(x)/scalingFactor), uint(float64(y)/scalingFactor), imagick.FILTER_BOX)
		orb.PosterizeImage(16, imagick.DITHER_METHOD_FLOYD_STEINBERG)
		orb.ResizeImage(x, y, imagick.FILTER_BOX)
		orb.ModulateImage(100, 135, 100)
	}
	orb.SetImageCompressionQuality(uint(quality))
	orb.SetCompressionQuality(uint(quality))
	out, err := orb.GetImageBlob()
	if err != nil {
		fmt.Println("couldn't shove it back in")
		sadness(s, m)
		return
	}
	outReader := bytes.NewReader(out)
	s.ChannelMessageSendComplex(c.ID, &discordgo.MessageSend{
		Reference: m.Reference(),
		Files: []*discordgo.File{
			{
				Name:   "img.jpg",
				Reader: outReader,
			},
		},
	})
}
