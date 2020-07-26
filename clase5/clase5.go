package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Person structure
type Person struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Age      int    `json:"age"`
	Handle   string `json:"handle,omitempty"`
	Password string `json:"-"`
	private  string
}

func initJSON() {
	fmt.Println("> Struct to JSON")
	martin := Person{
		Name:     "Martin",
		LastName: "Gonzalez",
		Age:      26,
		Handle:   "@ictg",
		Password: "password",
		private:  "private",
	}

	bs, err := json.Marshal(martin)
	if err != nil {
		fmt.Println(fmt.Errorf("%v", err))
		os.Exit(-1)
	}

	fmt.Println(string(bs))
	fmt.Println()

	fmt.Println("< JSON to Struct")
	var newPerson Person
	jsonBs := []byte(`{"name":"Martin","last_name":"Gonzalez","age":26,"handle":"@ictg","private":"private","password":"password"}`)
	err2 := json.Unmarshal(jsonBs, &newPerson)
	if err2 != nil {
		fmt.Println(fmt.Errorf("%v", err2))
		os.Exit(-1)
	}
	fmt.Printf("%#v\n", newPerson)
}

func randw() {
	martin := Person{
		Name:     "Martin",
		LastName: "Gonzalez",
		Age:      26,
		Handle:   "@ictg",
		Password: "password",
	}

	err := json.NewEncoder(os.Stdout).Encode(martin)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

// UserGHDTO get response struct
type UserGHDTO struct {
	Login             string      `json:"login"`
	ID                int         `json:"id"`
	NodeID            string      `json:"node_id"`
	AvatarURL         string      `json:"avatar_url"`
	GravatarID        string      `json:"gravatar_id"`
	URL               string      `json:"url"`
	HTMLURL           string      `json:"html_url"`
	FollowersURL      string      `json:"followers_url"`
	FollowingURL      string      `json:"following_url"`
	GistsURL          string      `json:"gists_url"`
	StarredURL        string      `json:"starred_url"`
	SubscriptionsURL  string      `json:"subscriptions_url"`
	OrganizationsURL  string      `json:"organizations_url"`
	ReposURL          string      `json:"repos_url"`
	EventsURL         string      `json:"events_url"`
	ReceivedEventsURL string      `json:"received_events_url"`
	Type              string      `json:"type"`
	SiteAdmin         bool        `json:"site_admin"`
	Name              string      `json:"name"`
	Company           interface{} `json:"company"`
	Blog              string      `json:"blog"`
	Location          string      `json:"location"`
	Email             interface{} `json:"email"`
	Hireable          interface{} `json:"hireable"`
	Bio               interface{} `json:"bio"`
	TwitterUsername   interface{} `json:"twitter_username"`
	PublicRepos       int         `json:"public_repos"`
	PublicGists       int         `json:"public_gists"`
	Followers         int         `json:"followers"`
	Following         int         `json:"following"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}

func httpGet() {
	res, err := http.Get("https://api.github.com/users/xetorthio")
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		log.Fatalf("Hubo un error: %d", res.StatusCode)
	}
	defer res.Body.Close()
	var user UserGHDTO
	json.NewDecoder(res.Body).Decode(&user)
	fmt.Printf("%#v\n", user)
}

func httpGeneric() {
	req, errReq := http.NewRequest(http.MethodGet, "https://api.github.com/users/xetorthio", nil)
	if errReq != nil {
		log.Fatal(errReq)
	}
	// req.Header.Set()
	res, errRes := http.DefaultClient.Do(req)
	if errRes != nil {
		log.Fatal(errRes)
	}
	defer res.Body.Close()
	var user UserGHDTO
	json.NewDecoder(res.Body).Decode(&user)
	fmt.Printf("%#v\n", user)
}

func contexto() {
	ctx := context.Background()
	ctxTimeout, cancel := context.WithTimeout(ctx, time.Microsecond)
	defer cancel()
	req, errReq := http.NewRequestWithContext(
		ctxTimeout,
		http.MethodGet,
		"https://api.github.com/users/xetorthio",
		nil,
	)
	if errReq != nil {
		fmt.Println(errReq)
		return
	}
	res, errRes := http.DefaultClient.Do(req)
	if errRes != nil {
		fmt.Println(errRes)
		return
	}
	defer res.Body.Close()
	var user UserGHDTO
	json.NewDecoder(res.Body).Decode(&user)
	fmt.Printf("%#v\n", user)
}

func main() {
	fmt.Println("Clase 5")
	fmt.Println("--------------")
	fmt.Println()

	fmt.Println("JSON")
	fmt.Println("--------------")
	initJSON()
	fmt.Println()

	fmt.Println("Reader & Writer")
	fmt.Println("--------------")
	randw()
	fmt.Println()

	fmt.Println("HTTP GET")
	fmt.Println("--------------")
	httpGet()
	fmt.Println()

	fmt.Println("HTTP Generic")
	fmt.Println("--------------")
	httpGeneric()
	fmt.Println()

	fmt.Println("Context")
	fmt.Println("--------------")
	contexto()
	fmt.Println()

	fmt.Println("Server")
	fmt.Println("--------------")
	Server()
	fmt.Println()
}
