/* SPDX-License-Identifier: GPL-3.0-or-later */
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type areaList struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Areas    []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func printAreaList(url string) (nextUrl string, previousUrl *string, err error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Couldn't get location-area:")
		return "", nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code %d:\n%s\n", res.StatusCode, body)
	}
	if err != nil {
		fmt.Println("Couldn't read response body:")
		return "", nil, err
	}

	areaList := areaList{}
	err = json.Unmarshal(body, &areaList)
	if err != nil {
		fmt.Println("Couldn't unmarshal JSON data")
		return "", nil, err
	}

	for _, area := range areaList.Areas {
		fmt.Println(area.Name)
	}

	return areaList.Next, areaList.Previous, nil
}

func derefOrEmpty(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func commandMap(cfg *config) error {
	nextUrl, previousUrl, err := printAreaList(cfg.Next)
	if err != nil {
		return err
	}

	cfg.Previous = derefOrEmpty(previousUrl)
	cfg.Next = nextUrl

	return nil
}

func commandMapBack(cfg *config) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	nextUrl, previousUrl, err := printAreaList(cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = nextUrl
	cfg.Previous = derefOrEmpty(previousUrl)

	return nil
}
