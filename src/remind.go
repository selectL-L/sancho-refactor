package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

// THIS WAS HELL!!!!!!!!!!!

type Reminder struct {
	end, start time.Time
	message string
	author string
	target string
	request *discordgo.Message
	rqid string
	timer *time.Timer
}

func setReminder(s *discordgo.Session, m *discordgo.MessageCreate, t *[]Reminder) {
	// first we need to parse time, then we set a timer and record it (in case the bot goes out)

	rawCmd, _ := strings.CutPrefix(m.Content, ".remind")
	if !slices.Contains(strings.Split(rawCmd, " "), "to"){
		s.ChannelMessageSendReply(m.ChannelID, "I know what you are. (please use \"to\" at the beginning of your reminder)", m.Reference())
		return
	}
	cmd := strings.Split(strings.TrimSpace(rawCmd), " ")
	msg := strings.Join(cmd[slices.Index(cmd,"to")+1:], " ")
	cmd = cmd[:slices.Index(cmd,"to")]
	targetUser := m.Author.ID

	if cmd[0] == "me" {
		cmd = cmd[1:]
	}
	if cmd[0][0] == '<' {
		if rawCmd[:2] != "me" {
			targetUser = cmd[0][2 : len(cmd[0])-1]
			cmd = cmd[1:]
		} else {
			iKnowWhatYouAre(s,m)
			return
		}
	}

	timeInUnix := 0
	eot := false
	if slices.Contains(cmd, "at") || slices.Contains(cmd, "on"){
		var sec, min, hour, day, month, year int
		if slices.Contains(cmd, "at"){
			rawt := strings.Split(cmd[slices.Index(cmd,"at")+1], ":")
			var cl []int
			for _, r := range rawt {
				val, err := strconv.Atoi(r)
				if err != nil {
					iKnowWhatYouAre(s,m)
					return
				}
				cl = append(cl, val)
			}
			if len(cl) == 2 {
				sec = 0
			} else if len(cl) == 3 {sec = cl[2]} else {iKnowWhatYouAre(s,m); return}
			min = cl[1]
			hour = cl[0]
		} else {
			sec = time.Now().Second()
			min = time.Now().Minute()
			hour = time.Now().Hour()
		}
		if slices.Contains(cmd, "on"){
			rawt := strings.Split(cmd[slices.Index(cmd,"on")+1], ".")
			var cl []int
			for _, r := range rawt {
				val, err := strconv.Atoi(r)
				if err != nil {
					iKnowWhatYouAre(s,m)
					return
				}
				cl = append(cl, val)
			}
			if len(cl) == 2 {
				year = time.Now().Year()
			} else if len(cl) == 3 {year = cl[2]} else {iKnowWhatYouAre(s,m); return}
			if year%100 == year {
				year += (time.Now().Year()/100)*100
			}
			month = cl[1]
			day = cl[0]
		} else {
			year = time.Now().Year()
			month = int(time.Now().Month())
			day = time.Now().Day()
		}
		destTime := time.Date(year, time.Month(month), day, hour, min, sec, 0, time.Now().Location())
		if destTime.Unix() <= time.Now().Unix() {
			iKnowWhatYouAre(s,m)
			return
		}
		timeInUnix = int(destTime.Unix())
	} else {
		totalTime := 0
		if cmd[0] == "in" {
			cmd = cmd[1:]
		}
		for len(cmd) > 0 || eot {
			timeIncr, err := strconv.Atoi(cmd[0])
			if err != nil {
				if cmd[0] == "a" || cmd[0] == "an" {
					timeIncr = 1
				} else if strings.Contains("smhdwyc", string(cmd[0][len(cmd[0])-1])){
					time, thing := cmd[0][:len(cmd[0])-1], cmd[0][len(cmd[0])-1]
					together := append([]string {time}, string(thing))
					cmd = append(together, cmd[1:]...)
					timeIncr, err = strconv.Atoi(cmd[0])
					if err != nil {
						iKnowWhatYouAre(s,m)
						return
					}
				} else {
					break
				}
			}
			switch strings.ToLower(cmd[1]) {
			case "s", "seconds", "second", "sec":
				totalTime += timeIncr * 1
			case "m", "minutes", "minute", "min":
				totalTime += timeIncr * 60
			case "h", "hours", "hour":
				totalTime += timeIncr * 60 * 60
			case "d", "days", "day":
				totalTime += int(time.Now().AddDate(0,0,1).Unix())-int(time.Now().Unix())
			case "w", "weeks", "week":
				totalTime += int(time.Now().AddDate(0,0,7).Unix())-int(time.Now().Unix())
			case "months", "month":
				totalTime += int(time.Now().AddDate(0,1,0).Unix())-int(time.Now().Unix())
			case "y", "years", "year":
				totalTime += int(time.Now().AddDate(1,0,0).Unix())-int(time.Now().Unix())
			case "c", "centuries", "century":
				totalTime += int(time.Now().AddDate(100,0,0).Unix())-int(time.Now().Unix())
			default:
				eot = true
			}
			if !eot {
				cmd = cmd[2:]
			}
		}
		if totalTime == 0 {
			iKnowWhatYouAre(s, m)
			return
		}
		timeInUnix = int(time.Now().Unix()) + totalTime
	}

	timerFile, err := os.OpenFile("timers.txt", os.O_APPEND, 0666)
	if err != nil {
		sadness(s, m)
		return
	}
	defer timerFile.Close()

	timerFile.WriteString(strings.Join([]string{m.Message.ID, strconv.Itoa(timeInUnix), strconv.Itoa(int(time.Now().Unix())), targetUser, m.Message.ChannelID, m.Author.ID, msg}, " ") + "\n")
	*t = append(*t, Reminder{
		end: time.Unix(int64(timeInUnix),0),
		start: time.Now(),
		message: msg,
		author: m.Author.ID,
		target: targetUser,
		request: m.Message,
		rqid: m.Message.ID,
		timer: time.NewTimer(time.Duration(timeInUnix-int(time.Now().Unix())) * time.Second),
	})
	s.ChannelMessageSendReply(m.ChannelID, "...As you wish. I shall send a reminder at <t:"+strconv.Itoa(timeInUnix)+">.", m.Reference())
}


func remind(s *discordgo.Session, r *Reminder){
	timerFile, err := os.OpenFile("timers.txt", os.O_RDONLY, 0666)
	if err != nil {
		sadness(s, nil)
		return
	}
	defer timerFile.Close()

	scanner := bufio.NewScanner(timerFile)

	request := r.request
	var reqChan string
	if request == nil { // if the original message was deleted and we have restored at least once...
		for scanner.Scan() {
			reminderText := strings.SplitN(scanner.Text(), " ", 6)
			if reminderText[0] == r.rqid {
				reqChan = reminderText[4]
			}
		}
		if reqChan != "" {
			if len("<@"+r.target+">: "+r.message+" (set at <t:"+strconv.Itoa(int(r.start.Unix()))+">)") >2000 {
				_, err = s.ChannelMessageSend(reqChan, "<@"+r.target+">: YOUR MESSAGE DIDN'T FIT")
			} else {
				_, err = s.ChannelMessageSend(reqChan, "<@"+r.target+">: "+r.message+" (set at <t:"+strconv.Itoa(int(r.start.Unix()))+">)")
			}
			if err != nil {
				sadness(s,nil)
				return
			}
		} else{
			sadness(s, nil)
			print(r.message)
			//SHIT WE DIDN'T FIND IT ABORT ABORT
			return
		}
	} else {
		if len("<@"+r.target+">: "+r.message+" (set at <t:"+strconv.Itoa(int(r.start.Unix()))+">)") >2000 {
			_, err = s.ChannelMessageSendReply(r.request.ChannelID, "<@"+r.target+">: YOUR MESSAGE DIDN'T FIT", r.request.SoftReference())
		} else {
			_, err = s.ChannelMessageSendReply(r.request.ChannelID, "<@"+r.target+">: "+r.message+" (set at <t:"+strconv.Itoa(int(r.start.Unix()))+">)", r.request.SoftReference())
		}
		if err != nil {
			sadness(s,nil)
			return
		}
	}
	
	if err!=nil {
		sadness(s, nil)
		return
	}

	// timerFile.Close()
	// timerFile, err = os.OpenFile("timers.txt", os.O_RDWR, 0666)
	// if err != nil {
	// 	sadness(s, nil)
	// 	return
	// }
	defer timerFile.Close()
	head, err := timerFile.Seek(0,0)
	if err!=nil {
		log.Fatal(head, err)
		return
	}

	newFileData := ""
	timerFile.Seek(0,0)
	scanner = bufio.NewScanner(timerFile)

	for scanner.Scan() {
		if strings.Split(scanner.Text(), " ")[0] != r.rqid {
			newFileData += scanner.Text() +"\n"
		}
	}

	i := slices.Index(reminders, *r)
	reminders = slices.Delete(reminders, i, i+1)
	// help!!!!

	err = os.WriteFile("timers.txt", []byte(newFileData), 0666)
	if err != nil {
		sadness(s, nil)
		return
	}
}

func listReminders(s *discordgo.Session, m *discordgo.MessageCreate, r *[]Reminder) {
	fullResponse := ""
	for i, rem := range *r{
		if rem.target == m.Author.ID {
			fullResponse += strconv.Itoa(i+1)+": "+rem.message+" @ <t:"+strconv.Itoa(int(rem.end.Unix()))+">\n"
		}
	}
	_, err := s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
		Content: fullResponse,
		Reference: m.Reference(),
		Flags: discordgo.MessageFlagsEphemeral,
	})
	if err != nil {
		sadness(s,m)
		return
	}
}

func deleteReminder(s *discordgo.Session, m *discordgo.MessageCreate, r *[]Reminder) {
	_, rawind, found := strings.Cut(m.Content, " ")
	if !found {
		iKnowWhatYouAre(s,m)
		return
	}
	ind, err := strconv.Atoi(rawind)
	if err!=nil {
		iKnowWhatYouAre(s,m)
		return
	}

	timerFile, err := os.OpenFile("timers.txt", os.O_RDWR, 0666)
	if err != nil {
		sadness(s, nil)
		return
	}
	defer timerFile.Close()

	counter := 0
	scanner := bufio.NewScanner(timerFile)
	newFileData := ""
	for scanner.Scan() {
		if strings.SplitN(scanner.Text()," ",7)[5] != m.Author.ID {
			newFileData += scanner.Text()+"\n"
		} else {
			counter++
			if counter != ind {
				newFileData += scanner.Text()+"\n"
			}
		}
	}
	err = os.WriteFile("timers.txt", []byte(newFileData), 0666)
	if err != nil {
		sadness(s, nil)
		return
	}

	counter = 0
	for i,rem := range *r {
		if rem.author == m.Author.ID {
			counter++
		}
		if counter == ind {
			s.ChannelMessageSendReply(m.ChannelID, "...Reminder to "+rem.message+" successfully deleted.", m.Reference())
			*r = slices.Delete(*r, i, i+1)
			return
		}
	}
}