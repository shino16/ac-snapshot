package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type Snapshot struct {
	Name    string    `json:"name"`
	Rating  int       `json:"rating"`
	ACCount int       `json:"ac_count"`
	RPS     int       `json:"rps"`
	Time    time.Time `json:"timestamp"`
}

func getRating(name string, retCh chan int, errCh chan error) {
	url := "https://atcoder.jp/users/" + url.PathEscape(name) + "/history/json"
	resp, err := http.Get(url)
	if err != nil {
		errCh <- err
		return
	}
	defer resp.Body.Close()
	body := make([]map[string]any, 0)
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		errCh <- err
		return
	}
	if len(body) == 0 {
		retCh <- 0
		return
	}
	last := body[len(body)-1]
	retCh <- int(last["NewRating"].(float64))
}

func getACCount(name string, retCh chan int, errCh chan error) {
	url := "https://kenkoooo.com/atcoder/atcoder-api/v3/user/ac_rank?user=" +
		url.QueryEscape(name)
	resp, err := http.Get(url)
	if err != nil {
		errCh <- err
		return
	}
	defer resp.Body.Close()
	body := make(map[string]float64)
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		errCh <- err
		return
	}
	retCh <- int(body["count"])
}

func getRPS(name string, retCh chan int, errCh chan error) {
	url := "https://kenkoooo.com/atcoder/atcoder-api/v3/user/" +
		"rated_point_sum_rank?user=" + url.QueryEscape(name)
	resp, err := http.Get(url)
	if err != nil {
		errCh <- err
		return
	}
	defer resp.Body.Close()
	body := make(map[string]float64)
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		errCh <- err
		return
	}
	retCh <- int(body["count"])
}

func takeSnapshot(name string) (Snapshot, error) {
	ratingCh := make(chan int)
	acCountCh := make(chan int)
	rpsCh := make(chan int)
	errCh := make(chan error)
	go getRating(name, ratingCh, errCh)
	go getACCount(name, acCountCh, errCh)
	go getRPS(name, rpsCh, errCh)
	shot := Snapshot{Name: name, Time: time.Now().UTC()}
	for i := 0; i < 3; i++ {
		select {
		case shot.Rating = <-ratingCh:
		case shot.ACCount = <-acCountCh:
		case shot.RPS = <-rpsCh:
		case err := <-errCh:
			return Snapshot{}, err
		}
	}
	return shot, nil
}
