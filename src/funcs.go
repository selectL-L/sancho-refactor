package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image/png"
	"io"
	"math"
	"math/big"
	cryptorand "crypto/rand"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/k4s/phantomgo"
	"gopkg.in/gographics/imagick.v3/imagick"
)

var err error

// call if something goes wrong
func sadness(s *discordgo.Session, m *discordgo.MessageCreate, errl error) {
	if m != nil{
		s.ChannelMessageSendReply(m.ChannelID, "Sorry, my creator must have fucked something up.\nPlease pierce him with a sanguine lance and drink his blood.", m.Reference())
	}
	fmt.Println(errl)
}

func help(s *discordgo.Session, m *discordgo.MessageCreate){
	ht, err := os.ReadFile("help.md")
	if err!=nil{
		sadness(s,m,err)
	}
	s.ChannelMessageSendReply(m.ChannelID, 
		string(ht),
		m.Reference())
}

func roll(s *discordgo.Session, m *discordgo.MessageCreate) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil || strings.Contains(m.Content, "bread"){
		return
	}
	if roll := composeRoll(m.Content); roll == ""{
		s.ChannelMessageSendReply(c.ID, roll, m.SoftReference())
	} else {
		s.ChannelMessageSendReply(c.ID, "Your roll is "+ roll+".", m.SoftReference())
	}
}


func editRoll(s *discordgo.Session, m *discordgo.MessageUpdate, mymsg *discordgo.Message){
	c, err := s.State.Channel(m.ChannelID)
	if err != nil || strings.Contains(m.Content, "bread"){
		return
	}
	if roll := composeRoll(m.Content); roll == ""{
		s.ChannelMessageEdit(c.ID, mymsg.ID, "I know what you are.")
	} else {
		s.ChannelMessageEdit(c.ID, mymsg.ID, "Your roll is "+ roll +".")
	}
}

func composeRoll(i string) string{
	var count, mod int
	r, _ := strings.CutPrefix(i, ".roll ")
	if idk, err := strconv.Atoi(r); err == nil && idk > 0 {
		count = 1
		mod = 0
		num, _ := cryptorand.Int(cryptorand.Reader, big.NewInt(int64(idk)))
		return strconv.Itoa(int(num.Int64()+1))
	} else {
		countS, rest, found := strings.Cut(r, "d")
		if !found {
			return ""
		}

		if countS == "" {
			count = 1
		} else {
			count, err = strconv.Atoi(countS)
			if err != nil {
				return ""
			}
		}

		modRune := strings.IndexAny(rest, "+-*^_")
		if modRune == -1 {
			modRune = len(rest)
		}

		max, err := strconv.Atoi(rest[:modRune])
		if err != nil {
			return ""
		}

		if max < 1 || count < 1 {
			fmt.Println(max, count)
			return ""
		}

		rawStr := ""
		sum := 0
		for i := 0; i < count; i++ {
			vbig, _ := cryptorand.Int(cryptorand.Reader, big.NewInt(int64((max))))
			v := int(vbig.Int64()) + 1
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
				return ""
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

		out := strconv.Itoa(sum) + " (" + rawStr[:len(rawStr)-1] + ")"
		outFull := "Your roll is " + out + "."
		if len(outFull) > 2000 {
			return strconv.Itoa(sum)
		} else {return out}
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
	resp, err := http.Get("https://tiphereth.zasz.su/static/assets/cards/Shi1.png")
	if err != nil {
		sadness(s,m,err)
	}
	defer resp.Body.Close()
	yujin := resp.Body

	resp, err = http.Get("https://tiphereth.zasz.su/static/assets/cards/Roland4Phase_Yujin.png")
	if err != nil {
		sadness(s,m,err)
	}
	defer resp.Body.Close()
	yujinDead := resp.Body

	pingId, _ := strings.CutPrefix(m.Content, ".bod ")
	//pingId = pingId[1:len(pingId)-1]
	ping := ""
	if _, err := strconv.Atoi(pingId[2 : len(pingId)-1]); pingId[1] == '@' && err == nil {
		ping = pingId
	}
	rollbig,_ := cryptorand.Int(cryptorand.Reader, big.NewInt(4))
	roll := int(rollbig.Int64())+1
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
	s.ChannelMessageSendReply(m.ChannelID, "What... is it this time?", m.Reference())
}

func fut(s *discordgo.Session, m *discordgo.MessageCreate) {
	messages := []string{"...I won't even have to call Father for this.",
		"...Hold your tongue; I will no longer tolerate any more \"ingenious ideas\".",
		"You're nothing before a Second Kindred, let alone Father.",
		"...Did you learn that from that Knight? How humorous.",
	}
	s.ChannelMessageSendReply(m.ChannelID, messages[rand.Intn(len(messages))], m.Reference())
}

func sendimg(s *discordgo.Session, m *discordgo.MessageCreate, name string){
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		return
	}
	img, err := os.Open("img/"+name)
	if err!=nil{
		sadness(s,m,err)
	}
	defer img.Close()
	s.ChannelMessageSendComplex(c.ID, &discordgo.MessageSend{
		Files: []*discordgo.File{
			{
				Name:   name,
				Reader: img,
			},
		},
	})
}

func jpegify(s *discordgo.Session, m *discordgo.MessageCreate, quality int) {
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		return
	}
	orb := imagick.NewMagickWand()
	defer orb.Destroy()
	var resp *http.Response
	var form string = "jpg"
	if len(m.Attachments) == 0 || !strings.Contains(m.Attachments[0].ContentType, "image") {
		if m.ReferencedMessage == nil {
			s.ChannelMessageSendReply(m.ChannelID, "Please send an actual image.", m.Reference())
			return
		}
		if len(m.ReferencedMessage.Attachments) == 0 || !strings.Contains(m.ReferencedMessage.Attachments[0].ContentType, "image") {
			if !strings.HasPrefix(m.ReferencedMessage.Content, "http"){
				s.ChannelMessageSendReply(m.ChannelID, "Please send an actual image.", m.Reference())
				return
			} else {
				// s.ChannelMessageSendReply(m.ChannelID, "Please select a non-Tenor gif.", m.Reference())
				// return
				resp, err = http.Get(m.ReferencedMessage.Content)
			}
		} else {
			resp, err = http.Get(m.ReferencedMessage.Attachments[0].URL)
		}
	} else {
		resp, err = http.Get(m.Attachments[0].URL)
	}
	

	//now we pipe the first (for now) image into imagemagick and wait for the result - how?

	//fn := m.Attachments[0].Filename
	//ext := fn[strings.LastIndex(fn, ".")+1:]
	if err != nil {
		fmt.Println("couldn't get image from internet")
		sadness(s,m,err)
		return
	}
	orig, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println("couldn't extract from html")
		sadness(s,m,err)
		return
	}

	if m.ReferencedMessage != nil {
		if strings.Contains(m.ReferencedMessage.Content, "tenor"){
			st := string(orig)
			url := strings.ReplaceAll(st[strings.Index(st, "contentUrl")+len("contentUrl\":\""):strings.Index(st, "thumbnailUrl")-3], "\\u002F", "/")
			resp, err = http.Get(url)
			if err != nil {
				fmt.Println(url)
				sadness(s,m,err)
				return
			}
			orig, err = io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				fmt.Println(url)
				sadness(s,m,err)
				return
			}
		}
	}
	
	err = orb.ReadImageBlob(orig)
	defer orb.Clear()
	if err != nil {
		fmt.Println("couldn't read img into orb", err)
		sadness(s,m,err)
		return
	}
	var out []byte
	if orb.GetNumberImages()>1{
		orb = orb.CoalesceImages()
		var disp imagick.DisposeType
		var w,h,del uint; var x,y int
		for i := 0; i<int(orb.GetNumberImages()); i++{
			orb.SetIteratorIndex(i)
			del = orb.GetImageDelay()
			w,h,x,y,_ = orb.GetImagePage()
			disp = orb.GetImageDispose()
			jpegifyImg(orb,quality)
			out, _ = orb.GetImageBlob()
			orb.ReadImageBlob(out)
			if i!=int(orb.GetIteratorIndex()){
				orb.PreviousImage()
			}
			orb.RemoveImage()
			orb.SetImageFormat("gif")
			orb.SetImageDispose(disp)
			orb.SetImageDelay(del)
			orb.SetImagePage(w,h,x,y)
		}
		out, err = orb.GetImagesBlob()
		form = "gif"
	} else {
		jpegifyImg(orb,quality)
		out, err = orb.GetImageBlob()
		form = "jpg"
	}
	if err != nil {
		fmt.Println("couldn't shove it back in")
		sadness(s,m,err)
		return
	}
	outReader := bytes.NewReader(out)
	s.ChannelMessageSendComplex(c.ID, &discordgo.MessageSend{
		Reference: m.Reference(),
		Files: []*discordgo.File{
			{
				Name:   "img."+form,
				Reader: outReader,
			},
		},
	})
}

func jpegifyImg(orb *imagick.MagickWand, q int){
	orb.SetImageFormat("JPEG")
	orb.SetImageCompressionQuality(uint(q))
	orb.SetCompressionQuality(uint(q))
	if q < 2 {
		x, y := orb.GetImageWidth(), orb.GetImageHeight()
		scalingFactor := math.Max(float64(x/160), float64(y/120)) // analogous to downscaling it to fit in a 240x180 box
		orb.ModulateImage(100, 150, 100)
		orb.ResizeImage(uint(float64(x)/scalingFactor), uint(float64(y)/scalingFactor), imagick.FILTER_BOX)
		out, _ := orb.GetImageBlob()
		orb.ReadImageBlob(out)
		if orb.GetIteratorIndex()==orb.GetNumberImages()-1{
			orb.PreviousImage()
		}
		orb.RemoveImage()
		orb.ResizeImage(x, y, imagick.FILTER_BOX)
		orb.SetImageFormat("JPEG")
		orb.PosterizeImage(16, imagick.DITHER_METHOD_FLOYD_STEINBERG)
		orb.SetImageCompressionQuality(uint(q))
		orb.SetCompressionQuality(uint(q))
	}
}

// IN DEVELOPMENT
func getPrescript(s *discordgo.Session, m *discordgo.MessageCreate){
	prescript, err := browser.Download(&phantomgo.Param{
		Method: "GET",
		Url: "https://nyos.dev/prescript",
		Header:       http.Header{"Cookie": []string{"your cookies"}},
		UsePhantomJS: true,
	})
	if err!=nil{
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(prescript.Body)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}


func lamentMournAndDespair(s *discordgo.Session, m *discordgo.MessageCreate){
	type frameData struct{
		f,s,x,y int
	}

	frames := []frameData{
		{16,105,409,153},
		{17,100,418,146},
		{18,100,378,147},
		{27,100,249,167},
		{28,95,249,167},
		{29,93,249,167},
		{30,90,249,167},
		{31,87,249,167},
		{32,85,248,154},
		{33,85,245,142},
		{34,85,281,119},
		{35,85,270,90},
		{38,85,238,66},
		{39,78,263,77},
		{40,70,291,88},
		{41,66,294,91},
		{42,63,298,91},
		{43,63,300,91},
		{51,85,292,49},
		{52,80,298,65},
		{53,80,307,70},
		{54,78,313,75},
		{55,77,315,73},
		{65,170,438,-11},
		{66,165,408,11},
		{67,160,404,10},
		{68,155,389,17},
		{91,70,393,148},
		{92,65,378,147},
		{93,57,375,147},
		{94,56,373,146},
		{95,55,371,146},
		{96,54,369,146},
		{97,53,367,146},
		{98,52,365,146},
		{99,52,364,146},
		{100,51,363,146},
		{101,51,362,146},
		{102,50,361,146},
		{103,50,360,146},
		{104,49,359,146},
		{105,49,358,146},
	}

	orb := imagick.NewMagickWand()
	defer orb.Destroy()
	lmdOrig, err := os.Open("img/lmd.gif")
	if err!=nil{
		sadness(s,m,err)
		return
	}
	defer lmdOrig.Close()
	orb.ReadImageFile(lmdOrig)
	target := m.Author
	if m.ReferencedMessage != nil{
		target = m.ReferencedMessage.Author
	}
	if len(m.Mentions)>0{
		target = m.Mentions[0]
	}
	avatar, err := s.UserAvatarDecode(target)
	if err!=nil{
		sadness(s,m,err)
		return
	}
	var buf bytes.Buffer
	err = png.Encode(&buf, avatar)
	if err!=nil{
		sadness(s,m,err)
		return
	}
	avatarOrb := imagick.NewMagickWand()
	defer avatarOrb.Destroy()
	err = avatarOrb.ReadImageBlob(buf.Bytes())
	if err!=nil{
		sadness(s,m,err)
		return
	}
	bwavatarOrb := avatarOrb.Clone()
	bwavatarOrb.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_DEACTIVATE)
	// bwavatarOrb.ModulateImage(50,0,100)
	// bwavatarOrb.SigmoidalContrastImage(true,20,100)
	// bwavatarOrb.BrightnessContrastImage(-25,80)
	avgOrb := bwavatarOrb.Clone()
	avgOrb.ResizeImage(1,1,imagick.FILTER_BOX)
	avgOrb.ModulateImage(100,0,0)
	a, _ := avgOrb.GetImagePixelColor(1,1)
	avg := a.GetBlueQuantum() //should be the same, hopefully
	bwavatarOrb.ThresholdImage(float64(avg))
	bwavatarOrb.ModulateImage(101,100,100)
	for _, fr := range frames {
		orb.SetIteratorIndex(fr.f)
		if fr.f>50 && fr.f<90 {
			bwavatarOrb.ScaleImage(uint(128*fr.s/100),uint(128*fr.s/100))
			orb.CompositeImage(bwavatarOrb, imagick.COMPOSITE_OP_OVER, true, fr.x-fr.s*64/100, fr.y-fr.s*64/100)
		} else {
			avatarOrb.ScaleImage(uint(128*fr.s/100),uint(128*fr.s/100))
			orb.CompositeImage(avatarOrb, imagick.COMPOSITE_OP_OVER, true, fr.x-fr.s*64/100, fr.y-fr.s*64/100)
		}
		orb.SetImageDispose(imagick.DISPOSE_NONE)
	}
	orb = orb.CoalesceImages()
	orb.SetImageFormat("GIF")
	out, err := orb.GetImagesBlob()
	if err!=nil{
		sadness(s,m,err)
		return
	}
	outReader := bytes.NewReader(out)
	s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
		Reference: m.Reference(),
		Files: []*discordgo.File{
			{
				Name:   "img.gif",
				Reader: outReader,
			},
		},
	})
}

func speechBubble(s *discordgo.Session, m *discordgo.MessageCreate){
	orb := imagick.NewMagickWand()
	defer orb.Destroy()
	alphaOrb := imagick.NewMagickWand()
	defer alphaOrb.Destroy()

	var resp *http.Response
	if len(m.Attachments) == 0 || !strings.Contains(m.Attachments[0].ContentType, "image") {
		if m.ReferencedMessage == nil {
			s.ChannelMessageSendReply(m.ChannelID, "Please send an actual image.", m.Reference())
			return
		}
		if len(m.ReferencedMessage.Attachments) == 0 || !strings.Contains(m.ReferencedMessage.Attachments[0].ContentType, "image") {
			if !strings.HasPrefix(m.ReferencedMessage.Content, "http"){
				s.ChannelMessageSendReply(m.ChannelID, "Please send an actual image.", m.Reference())
				return
			} else {
				resp, err = http.Get(m.ReferencedMessage.Content)
				fmt.Println(m.ReferencedMessage.Content)
			}
		} else {
			resp, err = http.Get(m.ReferencedMessage.Attachments[0].URL)
		}
	} else {
		resp, err = http.Get(m.Attachments[0].URL)
	}

	if err != nil {
		fmt.Println("couldn't get image from internet")
		sadness(s,m,err)
		return
	}
	orig, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Println("couldn't extract from html")
		sadness(s,m,err)
		return
	}

	_ = orb.ReadImageBlob(orig)
	defer orb.Clear()

	orb.SetFormat("PNGA")
	orb.SetImageFormat("PNGA")
	orb.AddImage(orb.GetImage())
	orb.SetLastIterator()
	orb.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_EXTRACT)
	
	maskFile, err := os.Open("img/mask.png")
	if err!=nil {
		sadness(s,m,err)
		return
	}
	alphaOrb.ReadImageFile(maskFile)
	alphaOrb.ResizeImage(orb.GetImageWidth(),orb.GetImageHeight(), imagick.FILTER_GAUSSIAN)
	alphaOrb.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_OFF)
	alphaOrb.CompositeImage(orb,imagick.COMPOSITE_OP_DARKEN,true,0,0)

	orb.SetFirstIterator()
	orb.SetImageAlphaChannel(imagick.ALPHA_CHANNEL_OFF)
	orb.SetFormat("PNG")
	orb.SetImageFormat("PNG")
	alphaOrb.SetFormat("PNG")
	alphaOrb.SetImageFormat("PNG")
	orb.CompositeImage(alphaOrb,imagick.COMPOSITE_OP_COPY_ALPHA, true,0,0)

	out, err := orb.GetImageBlob()
	if err!=nil{
		sadness(s,m,err)
		return
	}
	outReader := bytes.NewReader(out)
	s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
		Reference: m.Reference(),
		Files: []*discordgo.File{
			{
				Name:   "img.png",
				Reader: outReader,
			},
		},
	})
}


func sanchoball(s *discordgo.Session, m *discordgo.MessageCreate){
	outcomes := []string{"No.","Perhaps.","If Father wills it.","Most definitely.","Absolutely.","Maybe.","Clearly not.", "You'd be stupid to try.", "Is it not obvious?", "...", "Depends on you, and you alone.", "...You remind me of that arrogantly hopeful Fixer.", "For the Family, of course.", "V-verily, tis t-true-\n...You get the point. Yes.", "That is... simply impossible."}
	today := time.Now().YearDay()+time.Now().Year()*365

	_,msg,_ := strings.Cut(m.Content, " ")
	msg = strings.ToLower(msg)
	msgSalt := 0
	for i, l := range msg {
		msgSalt += int(l)*(i+1) * (((i+1)%2)*2-1)
	}
	var pseudo int
	id, _ := strconv.Atoi(m.Author.ID)
	if msgSalt != 0{
		saltedRand := rand.New(rand.NewSource(int64(msgSalt)))
		pseudo = int(math.Abs(float64((today + id + saltedRand.Int())%len(outcomes))))
	} else {
		pseudo = int(math.Abs(float64((today + id + rand.Int())%len(outcomes))))
	}

	s.ChannelMessageSendReply(m.ChannelID, outcomes[pseudo], m.SoftReference())
}