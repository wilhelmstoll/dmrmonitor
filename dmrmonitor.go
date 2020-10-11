package dmrmonitor

import (
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// IpscDmrURL defines the dmr monitor url.
const IpscDmrURL = "http://ipsc.dmr-dl.net/ipsc/_monitor.html"

// DmrEntry represents a dmr log entry.
type DmrEntry struct {
	Nr       string
	Hw       string
	Time     string
	Sec      string
	Rptr     string
	Call     string
	ID       string
	Callsign string
	Alias    string
	Ts       string
	Tg       string
	Ber      string
	Rssi     string
	Floor    string
}

// DmrMonitor represents the content of the dmr website.
type DmrMonitor struct {
	ActiveDmrEntries   []DmrEntry
	FinishedDmrEntries []DmrEntry
}

// Get calls the dmr website and prepares the data.
func Get() DmrMonitor {
	resp, err := http.Get(IpscDmrURL)
	if err != nil {
		log.Fatalf("error fetching URL: %v\n", err)
	}

	defer resp.Body.Close()

	tokenizer := html.NewTokenizer(resp.Body)

	tableNr := 0

	counter := 0
	entry := DmrEntry{}

	activeDmrEntries := []DmrEntry{}
	finishedDmrEntries := []DmrEntry{}

	for tokenizer.Token().Data != "html" {
		tokenType := tokenizer.Next()
		if tokenType == html.StartTagToken {
			token := tokenizer.Token()

			if token.Data == "table" {
				tableNr++
			}

			if token.Data == "tr" {
				entry = DmrEntry{}
				counter = 0
			}

			if token.Data == "td" {
				inner := tokenizer.Next()
				value := ""
				if inner == html.TextToken {
					innerText := (string)(tokenizer.Text())
					value = strings.TrimSpace(innerText)
					counter++
				} else if inner == html.EndTagToken {
					counter++
				}

				switch counter {
				case 1:
					entry.Nr = value
					break
				case 2:
					entry.Hw = value
					break
				case 3:
					entry.Time = value
					break
				case 4:
					entry.Sec = value
					break
				case 5:
					entry.Rptr = value
					break
				case 6:
					entry.Call = value
					break
				case 7:
					entry.ID = value
					break
				case 8:
					entry.Callsign = value
					break
				case 9:
					entry.Alias = value
					break
				case 10:
					entry.Ts = value
					break
				case 11:
					entry.Tg = value
					break
				case 12:
					entry.Ber = value
					break
				case 13:
					entry.Rssi = value
					break
				case 14:
					entry.Floor = value
					break
				}

				if counter == 14 {
					if entry.Nr != "" {
						if tableNr == 1 {
							activeDmrEntries = append(activeDmrEntries, entry)
						} else if tableNr == 2 {
							finishedDmrEntries = append(finishedDmrEntries, entry)
						}
					}

				}
			}

		}
	}

	return DmrMonitor{activeDmrEntries, finishedDmrEntries}
}
