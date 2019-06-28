package main

type Songs struct {
	Result struct {
		Songs []struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
			Ar   []struct {
				ID    int           `json:"id"`
				Name  string        `json:"name"`
				Tns   []interface{} `json:"tns"`
				Alias []interface{} `json:"alias"`
			} `json:"ar"`
			Al struct {
				ID     int           `json:"id"`
				Name   string        `json:"name"`
				PicURL string        `json:"picUrl"`
				Tns    []interface{} `json:"tns"`
				PicStr string        `json:"pic_str"`
				Pic    int64         `json:"pic"`
			} `json:"al"`
		} `json:"songs"`
		SongCount int `json:"songCount"`
	} `json:"result"`
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type Lyric struct {
	SongStatus   int    `json:"songStatus"`
	LyricVersion int    `json:"lyricVersion"`
	Lyric        string `json:"lyric"`
	Code         int    `json:"code"`
	Msg          string `json:"msg"`
}
