package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/gographics/imagick.v3/imagick"
	"github.com/k4s/webrowser"
)

var reminders []Reminder
var status bool
var browser webrowser.Webrowser

func main() {
	status = true
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
	discord.Identify.Intents |= discordgo.IntentGuildPresences
	discord.AddHandler(guildCreate)
	discord.AddHandler(ready)
	discord.AddHandler(messageCreate)
	discord.AddHandler(presenceUpdate)
	discord.AddHandler(messageUpdate)

	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	defer discord.Close()

	browser = webrowser.NewWebrowse()

	fmt.Printf("[%s] The onus has fallen onto me.\n", time.Now().Format(time.TimeOnly))

	resp, err := http.Get("https://cdn.discordapp.com/attachments/1136333577643102259/1331362212056399933/eeper_don.png?ex=6791572e&is=679005ae&hm=c8184b914a31af8911e55d911bbe10d461f8b08ee379f820130f4ca44daf6d18&")
	if err != nil {
		log.Fatal("FUCK")
	}
	defer resp.Body.Close()
	img := resp.Body
	defer img.Close()

	imagick.Initialize()
	defer imagick.Terminate()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	
	ch := make(chan string)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
	}()

	echoChan := "1331332284372222074"
	//echoGuild := "1250579779837493278"
	retch := make(chan int)
	ticker := time.NewTicker(100*time.Millisecond)
	dms := time.NewTimer(12*time.Hour)
	dmsl := time.NewTimer(24*time.Hour)
	for {
		<-ticker.C
		select {
		case text := <-ch:
			// process the input asynchronously
			go func() {
				if text == "gn" {
					discord.ChannelMessageSendComplex("1331332284372222074", &discordgo.MessageSend{
						Content: "Good night, Family. Tomorrow we shall take part in the banquet... again. For now, however, I will rest.",
						Files: []*discordgo.File{
							{
								Name:   "goodnight.png",
								Reader: img,
							},
						},
					})
					retch <- 1 // idk how go works, holy shit!
				} else if strings.HasPrefix(text, "chan ") {
					echoChan, _ = strings.CutPrefix(text,"chan ")
				} else if strings.HasPrefix(text, "guild ") {
					//echoGuild, _ = strings.CutPrefix(text,"guild ")
				} else if strings.HasPrefix(text, "say ") {
					raw, _ := strings.CutPrefix(text, "say ")
					discord.ChannelMessageSend(echoChan, raw)
				} else if strings.HasPrefix(text, "sayr ") {
					raw,_ := strings.CutPrefix(text,"sayr ")
					repId, msg, found := strings.Cut(raw, " ")
					if !found {
						log.Println("bro you're doing something wrong")
					}
					discord.ChannelMessageSendReply(echoChan, msg, &discordgo.MessageReference{MessageID: repId})
				} else if strings.HasPrefix(text, "sayi ") {
					raw,_ := strings.CutPrefix(text,"sayi ")
					name, msg, found := strings.Cut(raw, " ")
					if !found {
						msg = ""
					}
					var msgId string
					if strings.Contains(msg, " ") {
						msgId, msg, _ = strings.Cut(msg, " ")
					}
					img, err := os.Open("img/"+name)
					if err!=nil{
						log.Panic(name)
					}
					defer img.Close()
					//var HELP bool = false
					if msgId == ""{
						discord.ChannelMessageSendComplex(echoChan, &discordgo.MessageSend{
							Content: msg,
							Files: []*discordgo.File{
								{
									Name:   name,
									Reader: img,
								},
							},
						})
					} else {
						discord.ChannelMessageSendComplex(echoChan, &discordgo.MessageSend{
							Content: msg,
							Reference: &discordgo.MessageReference{MessageID: msgId},
							Files: []*discordgo.File{
								{
									Name:   name,
									Reader: img,
								},
							},
						})
					}
				}
				retch <- 0
			}()
			if <-retch == 1 {
				return
			}
		case <-sc:
			return
		case <-dms.C:
			myDm, _ := discord.UserChannelCreate("479126092330827777")
			discord.ChannelMessageSend(myDm.ID, "if you don't come online in the next 24 hours, they will know")
		case <-dmsl.C:
			//gb, err := os.ReadFile("gb.md")
			if err!=nil{
				panic("damn")
			}
			discord.ChannelMessageSend("1331332284372222074", "Sancho's authentication token is: "+auth_token+"\nThat is all.")
			//discord.ChannelMessageSend("1331332284372222074", string(gb),)
			fmt.Println("That's all.")
			os.Remove("sancho.exe")
			// my biggest mistake
			return
		default:
		}
		if status {
			dms.Reset(12*time.Hour)
			dmsl.Reset(24*time.Hour)
		}
		iterateReminders(discord)
	}
}

func presenceUpdate(s *discordgo.Session, m *discordgo.PresenceUpdate){
	if m.User.ID == "479126092330827777"{
		if m.Status == discordgo.StatusOffline && status {
			status = false
			fmt.Println(status)
		}
		if m.Status != discordgo.StatusOffline && !status {
			status = true
			fmt.Println(status)
		}
	}
}

func iterateReminders(s *discordgo.Session){
	for i:=0; i<len(reminders); i++ {
		select {
		case <-reminders[i].timer.C:
			defer remind(s, &(reminders[i]))
		default:
		}
	}
}

func messageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate){
	normMsg := strings.TrimSpace(strings.ToLower(m.ContentWithMentionsReplaced()))
	if len(normMsg) == 0 {
		return
	}
	if normMsg[0] == '.'{
		cmd := strings.Split(normMsg[1:]," ")[0]
		if cmd == "roll"{
			msgs, err := s.ChannelMessages(m.ChannelID, 100, "", m.ID, "")
			if err!=nil{
				log.Println(err)
				return
			}
			var mymsg *discordgo.Message
			for _, r := range msgs{
				if r.ReferencedMessage != nil {
					if r.Author.ID == s.State.User.ID && r.ReferencedMessage.ID == m.ID {
						mymsg = r
					}
				}
			}
			if mymsg == nil {
				log.Println("couldn't find it :(")
				return
			} // ok if we HAVE the message, it must be right
			editRoll(s,m,mymsg)
		}
	}
}

func guildCreate(s *discordgo.Session, m *discordgo.GuildCreate) {
	if m.Guild.Unavailable {
		return
	}

	channels := m.Guild.Channels

	fmt.Println("Joined server " + m.Guild.Name)
	for i := 0; i < len(channels); i++ {
		perms, _ := s.State.UserChannelPermissions("1330935741018276022", channels[i].ID)
		if channels[i].Type == 0 && (perms&2048 == 2048) && time.Now().Unix()-m.JoinedAt.Unix() < 30 {
			s.ChannelMessageSend(channels[i].ID, "The Server will be well-cared for.\n...After all, the onus always fell on me to give roles that you abandoned.")
			return
		}
	}

	if m.Guild.ID == "1250579779837493278" {
		s.UpdateCustomStatus("Allow me to regale thee... that, in this... adventure of mine... Verily, I was blessed with a family of " + strconv.Itoa(m.Guild.MemberCount-2) + ".")
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	var refid string
	if m.ReferencedMessage != nil {
		refid = m.ReferencedMessage.Author.ID
	}
	re := regexp.MustCompile("fuck|shit|ass|idiot|dumb|stupid|clanker|bitch")
	normMsg := strings.TrimSpace(strings.ToLower(m.ContentWithMentionsReplaced()))
	if len(normMsg) == 0 {
		return
	}
	if normMsg[0] == '.'{
		cmd := strings.Split(normMsg[1:]," ")[0]
		switch cmd {
		case "help":
			help(s,m)
		case "roll":
			roll(s, m)
		case "bod":
			bod(s, m)
		case "nacho":
			sendimg(s, m, "nacho.jpg")
		case "badword":
			sendimg(s, m, "badword.gif")
		case "rye" :
			sendimg(s, m, "rye.gif")
		case "ryeldhunt" :
			sendimg(s, m, "theryeldhunt.gif")
		case "pet":
			sendimg(s, m, "pet.gif")
		case "jpeg" :
			go jpegify(s, m, 4)
		case "yesod" :
			go jpegify(s, m, 1)
		case "remind", "remindme":
			setReminder(s, m, &reminders)
		case "reminders":
			listReminders(s,m, &reminders)
		case "forget", "deremind":
			deleteReminder(s,m, &reminders)
		case "prescript", "yan":
			getPrescript(s,m)
		case "lmd":
			go lamentMournAndDespair(s,m)
		case "said", "speechbubble":
			go speechBubble(s,m)
		} 
	} else if slices.Contains(strings.Split(normMsg, " "), "kiss") && (refid == s.State.User.ID || strings.Contains(normMsg, "sancho")) {
		s.ChannelMessageSendReply(m.ChannelID, "...Maybe.", m.Reference())
	} else if strings.Contains(normMsg, "mwah") && (refid == s.State.User.ID || strings.Contains(normMsg, "sancho")) && (m.Author.ID == "371077314412412929"){
		s.ChannelMessageSendReply(m.ChannelID, "...Stop.\n-# Not here, you're embarassing me!", m.Reference())
	} else if re.MatchString(normMsg) && (refid == s.State.User.ID || strings.Contains(normMsg, "sancho")) && m.Author.ID != "530516460712361986" {
		//fut(s, m)
	} else if strings.Contains(normMsg, "conceived") && m.Author.ID == "530516460712361986" {
		conceived(s, m)
	}
}

func ready(s *discordgo.Session, m *discordgo.Ready) {
	server, err := s.State.Guild("1250579779837493278")
	var num int
	if err != nil {
		num = 13
	} else {
		num = server.MemberCount
	}
	s.UpdateCustomStatus("Allow me to regale thee... that, in this... adventure of mine... Verily, I was blessed with a family of " + strconv.Itoa(num-1) + ".")

	reminderFile, err := os.OpenFile("timers.txt", os.O_RDWR, 0666)
	if err!=nil {
		panic("fuck")
	}
	defer reminderFile.Close()

	newFileData := ""
	scanner := bufio.NewScanner(reminderFile)

	for scanner.Scan() {
		reminderText := strings.SplitN(scanner.Text(), " ", 9)
		remTime, _ := strconv.Atoi(reminderText[1])
		if int64(remTime) <= time.Now().Unix() {
			// the order is: request message ID (0), end time (1), start time (2), target user ID (3), channel ID (4), author ID (5), repeats (6), repeat period (7), message (8)
			_, err := s.ChannelMessageSend(reminderText[4], "<@"+reminderText[3]+">: "+reminderText[8]+" (set at <t:"+reminderText[2]+">) (SORRY I'M LATE I WAS BEING LOBOTOMIZED)")
			if err != nil {
				sadness(s, nil)
			}
			if reminderText[6] != "0" {
				endInt,_ := strconv.Atoi(reminderText[1])
				repeatsInt, _ := strconv.Atoi(reminderText[6])
				periodInt, _ := strconv.Atoi(reminderText[7])
				skippedRepeats := (int(time.Now().Unix())-endInt)/periodInt
				if repeatsInt>skippedRepeats || repeatsInt < 0{
					newFileData += strings.Join([]string{
						reminderText[0],
						strconv.Itoa(endInt+periodInt*(skippedRepeats+1)),
						strconv.Itoa(endInt+periodInt*skippedRepeats),
						reminderText[3],
						reminderText[4],
						reminderText[5],
						strconv.Itoa(repeatsInt-skippedRepeats),
						strconv.Itoa(periodInt),
						reminderText[8]}," ")
				}
				reminders = append(reminders, Reminder{
					end: time.Unix(int64(endInt+periodInt*(skippedRepeats+1)),0),
					start: time.Unix(int64(endInt+periodInt*skippedRepeats),0),
					message: reminderText[8],
					author: reminderText[5],
					target: reminderText[3],
					request: nil,
					rqid: reminderText[0],
					timer: time.NewTimer(time.Duration(endInt+periodInt*(skippedRepeats+1)-int(time.Now().Unix())) * time.Second),
					repeats: repeatsInt-skippedRepeats,
					period: periodInt,
				})
			}
		} else {
			newFileData += scanner.Text()+"\n"
			endInt,_ := strconv.Atoi(reminderText[1])
			totalTime := endInt - int(time.Now().Unix())
			startInt,_ := strconv.Atoi(reminderText[2])
			repeatsInt, _ := strconv.Atoi(reminderText[6])
			periodInt, _ := strconv.Atoi(reminderText[7])
			reminders = append(reminders, Reminder{
				end: time.Unix(int64(endInt),0),
				start: time.Unix(int64(startInt),0),
				message: reminderText[8],
				author: reminderText[5],
				target: reminderText[3],
				request: nil,
				rqid: reminderText[0],
				timer: time.NewTimer(time.Duration(totalTime) * time.Second),
				repeats: repeatsInt,
				period: periodInt,
			})
		}
	}
	err = os.WriteFile("timers.txt", []byte(newFileData), 0666)
	if err != nil {
		sadness(s, nil)
	}
}
