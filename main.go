package main

import (
	"crypto/tls"
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
	ircobj.UseTLS = true
	ircobj.TLSConfig = &tls.Config{InsecureSkipVerify: true}
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
			switch {
			case message == "inflik":
				ircobj.Privmsg(e.Arguments[0], "Jag skulle bara vilja inflika för ett ögonblick.")
				ircobj.Privmsg(e.Arguments[0], "Vad du kallar Linux är faktiskt GNU/Linux, eller som jag själv kallar det, GNU+Linux. Linux är inte ett operativsystem i sig själv, utan snarare ännu en del utav ett funktionellt GNU-system, som görs användbart av GNU kärnbibliotek, skalverktyg, och andra nödvändiga delar, som tillsammans definerar ett OS enligt POSIX.")
				ircobj.Privmsg(e.Arguments[0], "Många datorer kör ett modifierat GNU-system varje dag, utan att inse det. Genom en lustig härva av händelser kallas det GNU som används ofta Linux, och många av dess användare inser inte att de använder GNU-systemet, som utvecklats utav GNU-projektet.")
				ircobj.Privmsg(e.Arguments[0], "Det finns ett Linux, och dessa människor använder det, men det är bara en del av systemet de använder. Linux är kärnan, programmet i systemet som allokerar maskinens resurser till de andra programmen du kör. Kärnan är en viktig del utav ett operativsystem, men helt oanvändbart i sig själv; den kan bara fungera i samband med ett helt operativsystem.")
				ircobj.Privmsg(e.Arguments[0], "Linux används oftast i samband med GNU-operativsystemet: hela systemet är bara GNU med Linux tillagt, eller GNU/Linux. Alla så kallade Linux-distrubitioner är egentligen distrubutioner utav GNU/Linux!")

			case message == "blasera":
				ircobj.Action(e.Arguments[0], "420blaserar det gröna")
			case message == "hellseger":
				ircobj.Privmsg(e.Arguments[0], hellseger())
			case message == "kommandon" || message == "hjälp":
				ircobj.Privmsg(e.Arguments[0], "Nuvarande kommandon: blasera, kattljud, hellseger, pälsknulla <offer>, älska <offer>, fluffa <offer>, mörda <offer>")
			case strings.HasPrefix(message, "pälsknulla"):
				if strings.Replace(message, " ", "", -1) == "pälsknulla" {
					ircobj.Privmsg(e.Arguments[0], "pälsknulla tar ett argument - offret")
				} else {
					target := strings.Fields(message)[1]
					var yiffstring = pälsknulla(target)
					ircobj.Action(e.Arguments[0], yiffstring)
				}

			case strings.HasPrefix(message, "fluffa"):
				if strings.Replace(message, " ", "", -1) == "fluffa" {
					ircobj.Privmsg(e.Arguments[0], "fluffa tar ett argument - offret")
				} else {
					target := strings.Fields(message)[1]
					ircobj.Action(e.Arguments[0], "fluffar "+target)
				}
			case strings.HasPrefix(message, "älska"):
				if strings.Replace(message, " ", "", -1) == "älska" {
					ircobj.Privmsg(e.Arguments[0], "älska tar ett argument - offret")
				} else {
					target := strings.Fields(message)[1]
					var lovestring = älska(target)
					ircobj.Action(e.Arguments[0], lovestring)
				}
			case message == "kattljud":
				ircobj.Privmsg(e.Arguments[0], kattljud())
			case strings.TrimSpace(message) == "ss" || strings.TrimSpace(message) == "per":
				ircobj.Privmsg(e.Arguments[0], sprutskit())
			case strings.HasPrefix(message, "mörda"):
				if strings.Replace(message, " ", "", -1) == "mörda" {
					ircobj.Privmsg(e.Arguments[0], "mörda tar ett argument - offret")
				} else {
					target := strings.Fields(message)[1]
					ircobj.Action(e.Arguments[0], mörda(target))
				}
			}
		}
	})

	ircobj.Loop()
}

func mörda(person string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	action := []string{"isar", "cuttar", "misshandlar", "spräcker", "dränker", "klyver", "maler", "kväver", "stryper", "åderlåter", "kokar", "pastöriserar"}
	other := []string{"hårt", "mjukt", "försiktigt", "med glimten i ögat", "med ett mordiskt vrål", "med en tår i ögat", "naken", "med ett maniskt skratt", "medans Klas Lund tittar på", "ute i skogen", "med enbart en halsduk på sig"}

	hack := []string{action[r.Intn(len(action))], person, other[r.Intn(len(action))]}
	return strings.Join(hack, " ")
}
func sprutskit() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := []string{"sprutskiter", "bajsar", "kastar skit", "kräks", "kanonbajsar", "spottar bajs", "kletar bajs", "kissar", "fiser", "hostar bajs", "spottar skit"}
	p := []string{"på", "mot", "under" }
	t := []string{"taket", "stolen", "skärmen", "golvet", "datorn", "bänken", "skrivbordet", "chloe", "grannen", "mig själv", "alla flash-XSSer som någonsin hittats", "alla som gillar family guy", "alla u-länder", "äldreboendets fönster", "kyrkporten", "alla som är dumma nog att skaffa barn", "pentagramet jag kissade i snön", "sql-injektionerna som butkus hittade", "hela jävla swehack", "goober", "butkus", "bordet", "musmattan", "chassit", "chloes dator" }
	hack := []string{s[r.Intn(len(s))], p[r.Intn(len(p))], t[r.Intn(len(t))]}
	return strings.Join(hack, " ")
}
func pälsknulla(person string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var action = []string{"biter", "slickar", "penetrerar", "leker med", "spräcker", "misshandlar", "kysser", "våldtar", "glider in i", "drar ut sigsjälv från", "tänder eld på", "gnuggar sig själv mot"}
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
	var description = []string{"försiktigt", "mjukt", "hårt", "lugnt", "hjärtfyllt", "älskande", "varmt", "gulligt", "lyckligt"}

	hack := []string{action[r.Intn(len(action))], person, description[r.Intn(len(description))], "♥"}
	return strings.Join(hack, " ")

}

func kattljud() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var noises = []string{"mew", "meow", "mrawr", "rawr", "rar", "mrowr", "rwr", "mrw", "mrwrw"}
	return noises[r.Intn(len(noises))]
}

func hellseger() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fraser := []string{"Klas \"Ajöken, ökenfröken!\" Lund", "Klas \"Bemöter hula-hula med dumdumkula\" Lund", "Klas \"Beväpna germanderna, nu intar vi skanderna\" Lund", "Klas \"Blixt & Dunder inte Banker & Kunder\" Lund", "Klas \"Blodad tand för nordiskt land\" Lund", "Klas \"Bottenlöst hat mot kalifat\" Lund", "Klas \"Bryt dig fri från skuld - sno en judes guld\" Lund", "Klas \"Demokrati är hyckleri\" Lund", "Klas \"Dränk en bankman i en ankdamm\" Lund", "Klas \"En nordisk renässans krossar judens storfinans\" Lund", "Klas \"En nordisk union är min vision\" Lumd", "Klas \"Fjorton åttioåtta, gasa en zigenarråtta\" Lund", "Klas \"Framgent blir Sverige homogent\" Lund", "Klas \"Genant att vara släkt med en migrant\" Lund", "Klas \"Gör profit på att vara antisemit\" Lund", "Klas \"Grov patron för semitisk religion\" Lund", "Klas \"Grov patron mot semitisk religion\" Lund", "Klas \"Gul, brun, svart eller röd, efter raskriget är du död\" Lund", "Klas \"Hell seger, häng en neger\" Lund", "Klas \"Hyperboreas kall blir västvärldens fall\" Lund", "Klas \"Ingen degeneration i vår nordiska federation\" Lund", "Klas \"Ingen frid för negroid individ\" Lund", "Klas \"Juden står böjd, för flaggan på höjd, nu är jag nöjd\" Lund", "Klas \"Kalabalik baserad på euganik\" Lund", "Klas \"Laddad AK-47 väntar på HBTQ\" Lund", "Klas \"Makt åt norden, lägg juden i jorden\" Lund", "Klas \"Man bestiger ingen klippa i kippa\" Lund", "Klas \"Med blodad klo mot NWO\" Lund", "Klas \"Med kniv mot Tel Aviv\" Lund", "Klas \"Migration med båt slutar i gråt\" Lund", "Klas \"Oden har orden i Norden\" Lund", "Klas \"Överflöd av bråd död\" Lund", "Klas \"Rasfrämling ger dålig stämning\" Lund", "Klas \"Rasren skurk dödar varenda turk\" Lund", "Klas \"Rätt hudfärg för att klättra i berg\" Lund", "Klas \"Rätt ras, annars gas\" Lund", "Klas \"Sätter ryssen i abyssen\" Lund", "Klas \"Schäferhund har kurd i mund\" Lund", "Klas \"Skinnflå en eskimå\" Lund", "Klas \"Skyddar förhuden från Bonnierjuden\" Lund", "Klas \"Slå ett slag för repets dag\" Lund", "Klas \"Slösa ingen sympati på papperslösa\" Lund", "Klas \"Stångarhornet mot de som kraschade in i WTC-tornet\" Lund", "Klas \"Ta och tig, nu är dax för rasligt krig\" Lund", "Klas \"Tor är min storebror\" Lund", "Klas \"Total banzai mot alla Thai\" Lund", "Klas \"Utländsk mat? Smaka vårt hat\" Lund", "Klas \"Vit man sätter bonnier i brand\" Lund", "Klas \"Zion får smällen, sen bestiger vi fjällen\" Lund"}
	return fraser[r.Intn(len(fraser))]
}
