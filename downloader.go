package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const lyricURL = "http://music.163.com/api/song/media?id="

func DownloadLyric(fileName string, songId int, transport *http.Transport, timeout time.Duration) error {
	var extension = filepath.Ext(fileName)
	var name = fileName[0 : len(fileName)-len(extension)]
	return downloadFile(name+".lrc", lyricURL+strconv.Itoa(songId), transport, time.Second*10)
}

func downloadFile(fileName string, url string, transport *http.Transport, timeout time.Duration) error {
	client := &http.Client{Transport: transport, Timeout: timeout}

	newReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	newReq.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) "+
		"Chrome/74.0.3729.131 Safari/537.36")
	newReq.Header.Set("Host", "music.163.com")
	newReq.Header.Set("Referer", "http://music.163.com/search/")
	newReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var resp *http.Response
	for resp, err = client.Do(newReq); err != nil; {
		resp, err = client.Do(newReq)
	}
	defer Close(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return errors.New(url + " " + resp.Status)
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var lyc Lyric
	err = json.Unmarshal(res, &lyc)
	if err != nil {
		panic(err)
	}

	if lyc.Code != 200 {
		return errors.New(lyc.Msg)
	}

	output, err := os.Create(fileName)
	defer Close(output)

	if err != nil {
		return err
	}

	_, err = output.Write([]byte(lyc.Lyric))
	if err != nil {
		return err
	}

	return nil
}
