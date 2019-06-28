package main

import (
	"encoding/json"
	"flag"
	"github.com/dhowden/tag"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

var dir = flag.String("dir", ".", "input dir")
var proxy = flag.String("proxy", "", "network proxy")
var second = flag.Int("timeout", 10, "timeout second")
var thread = flag.Int("thread", 20, "thread number")
var extensions = flag.String("extensions", "mp3,flac", "music extensions")
var transport *http.Transport

func init() {
	flag.Parse()
	err := setupNetwork()
	if err != nil {
		panic(err)
	}
}

func main() {
	var wg sync.WaitGroup
	files := LoopFile(*dir, strings.Split(*extensions, ","))
	for i := 0; i < *thread; i++ {
		wg.Add(1)
		go execute(files, &wg)
	}
	wg.Wait()
}

func execute(files chan string, wg *sync.WaitGroup) {
	timeout := time.Second * time.Duration(*second)
	for f := range files {
		file, _ := os.Open(f)
		m, err := tag.ReadFrom(file)
		if err != nil {
			log.Fatal(err)
		}
		searchText := m.Title() + " " + m.Artist()
		var songJson string
		for songJson, err = SearchSong(searchText, transport, timeout); songJson == "" && err == nil; {
			songJson, err = SearchSong(searchText, transport, timeout)
		}
		if err != nil {
			log.Printf("failed to search song: %s, %s", f, err)
			break
		}
		var songList Songs
		err = json.Unmarshal([]byte(songJson), &songList)
		if err != nil {
			panic(err)
		}

		if songList.Code != 200 {
			log.Printf("failed to fetch song: %s, %s", f, songList.Msg)
			break
		} else if len(songList.Result.Songs) <= 0 {
			log.Printf("failed to fetch song: %s, no song found", f)
			break
		}
		for j, song := range songList.Result.Songs {
			if strings.TrimSpace(song.Name) == strings.TrimSpace(m.Title()) &&
				strings.Contains(m.Artist(), song.Ar[0].Name) {
				err = DownloadLyric(f, song.ID, transport, timeout)
				if err != nil {
					log.Printf("download lyric occurs error: %s", err)
				}
				break
			}
			if j == len(songList.Result.Songs)-1 {
				log.Printf("failed to match song: %s, no matched song", f)
			}
		}
	}
	wg.Done()
}

func setupNetwork() error {
	if *proxy != "" {
		proxyUrl, err := url.Parse(*proxy)
		if err != nil {
			return err
		}

		transport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	} else {
		transport = &http.Transport{}
	}

	return nil
}
