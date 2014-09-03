package main

import (
	"encoding/json"
	"fmt"
	"github.com/thoj/go-ircevent"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Config struct {
	Server   string
	Channels []string
	Realname string
	Nick     string
	Prefix   string
}

func main() {

	file, ferr := os.Open("config.json")
	if ferr != nil {
		fmt.Println(ferr)
	}
	decoder := json.NewDecoder(file)
	config := Config{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println(err)
	}

	ircobj := irc.IRC(config.Nick, config.Realname)

	error := ircobj.Connect(config.Server)
	if error != nil {
		fmt.Println(error)
	}
	ircobj.AddCallback("001", func(e *irc.Event) {
		for _, channel := range config.Channels {
			fmt.Println("Joining ", channel)
			ircobj.Join(channel)
		}
	})

	go ircobj.AddCallback("PRIVMSG", func(e *irc.Event) {
		message := e.Message()
		if strings.HasPrefix(message, config.Prefix) {
			message = strings.TrimPrefix(message, config.Prefix)
		}
		if message == "inflik" {
			ircobj.Privmsg(e.Arguments[0], "Jag skulle bara vilja inflika för ett ögonblick.")
			ircobj.Privmsg(e.Arguments[0], "Vad du kallar Linux är faktiskt GNU/Linux, eller som jag själv kallar det, GNU+Linux. Linux är inte ett operativsystem i sig själv, utan snarare ännu en del utav ett funktionellt GNU-system, som görs användbart av GNU corelibs, shell-utils, och andra nödvändiga delar, som tillsammans definerar ett OS enligt POSIX.")
			ircobj.Privmsg(e.Arguments[0], "Många datorer kör ett modifierat GNU-system varje dag, utan att inse det. Genom en lustig härva av händelser kallas det GNU som används ofta Linux, och många av dess användare inser inte att de använder GNU-systemet, som utvecklats utav GNU-projektet.")
			ircobj.Privmsg(e.Arguments[0], "Det finns ett Linux, och dessa människor använder det, men det är bara en del av systemet de använder. Linux är kerneln, programmet i systemet som allokerar maskinens resurser till de andra programmen du kör. Kerneln är en viktig del utav ett operativsystem, men helt oanvändbart i sig själv; den kan bara fungera i samband med ett helt operativsystem.")
			ircobj.Privmsg(e.Arguments[0], "Linux används oftast i samband med GNU-operativsystemet: hela systemet är bara GNU med Linux tillagt, eller GNU/Linux. Alla så kallade Linux-distrubitioner är egentligen distrubutioner utav GNU/Linux!")
		}
		if message == "blasera" {
			ircobj.Action(e.Arguments[0], "420blaserar det gröna")
		}
		if message == "kommandon" || message == "hjälp" {
			ircobj.Privmsg(e.Arguments[0], "Nuvarande kommandon: blasera, kattljud, pälsknulla <offer>, älska <offer>, fluffa <offer>")
		}
		if strings.Replace(message, " ", "", -1) == "pälsknulla" {
			ircobj.Privmsg(e.Arguments[0], "pälsknulla tar ett argument - offret")
		} else if strings.Fields(message)[0] == "pälsknulla" {
			target := strings.Fields(message)[1]
			var pälsknullastring = pälsknulla(target)
			ircobj.Action(e.Arguments[0], pälsknullastring)
		}

		if strings.Replace(message, " ", "", -1) == "fluffa" {
			ircobj.Privmsg(e.Arguments[0], "fluffa tar ett argument - offret")
		} else if strings.Fields(message)[0] == "fluffa" {
			target := strings.Fields(message)[1]
			var fluffastring = "fluffar " + target
			ircobj.Action(e.Arguments[0], fluffastring)
		}

		if strings.Replace(message, " ", "", -1) == "älska" {
			ircobj.Privmsg(e.Arguments[0], "älska tar ett argument - offret")
		} else if strings.Fields(message)[0] == "älska" {
			target := strings.Fields(message)[1]
			var älskastring = älska(target)
			ircobj.Action(e.Arguments[0], älskastring)
		}
		if message == "kattljud" {
			ircobj.Privmsg(e.Arguments[0], kattljud())
		}

	})
	ircobj.Loop()
}

func pälsknulla(person string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var action = []string{"biter", "slickar", "penetrerar", "leker med", "spräcker", "misshandlar", "kyssar", "våldtar", "glider in i", "drar ut sigsjälv från", "tänder eld på", "gnuggar sig själv mot"}
	var description = []string{"skrumpna", "illadoftande", "unga", "slitna", "stinkande", "enorma", "minimala", "söta", "lösa", "mjuka", "droppande", "brännheta", "iskalla", "regnbågsfärgade"}
	var object = []string{"murrhål", "armhåla", "ansikte", "navel", "kloak", "gentooinstallation", "tunga", "livmoder", "strumpa", "avgasrör", "plåttermos", ".... Skit samma, jag vill inte detta egentligen."}

	if !strings.HasSuffix(person, "s") {
		person = person + "s"
	}
	hack := []string{action[r.Intn(len(action))], person, description[r.Intn(len(description))], object[r.Intn(len(object))]}
	return strings.Join(hack, " ")

}

func älska(person string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var action = []string{"pussar", "slickar", "gosar", "smeker", "lindar armarna om", "kramar", "klappar", "myser", "sniffar", "nafsar", "gnuggar", "eskimåkysser"}
	var description = []string{"försiktigt", "mjukt", "hårt", "lugnt", "hjärtfyllt", "älskande", "varmt", "golligt", "lyckligt"}

	hack := []string{action[r.Intn(len(action))], person, description[r.Intn(len(description))], "♥"}
	return strings.Join(hack, " ")

}

func kattljud() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var noises = []string{"mew", "meow", "mrawr", "rawr", "rar", "mrowr", "rwr", "mrw", "mrwrw"}
	return noises[r.Intn(len(noises))]
}
