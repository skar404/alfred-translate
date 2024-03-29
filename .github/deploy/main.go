package main

import (
	"flag"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
	"time"
)

type Done struct {
	Url       string `json:"url"`
	AssetsUrl string `json:"assets_url"`
	UploadUrl string `json:"upload_url"`
	HtmlUrl   string `json:"html_url"`
	Id        int    `json:"id"`
	Author    struct {
		Login             string `json:"login"`
		Id                int    `json:"id"`
		NodeId            string `json:"node_id"`
		AvatarUrl         string `json:"avatar_url"`
		GravatarId        string `json:"gravatar_id"`
		Url               string `json:"url"`
		HtmlUrl           string `json:"html_url"`
		FollowersUrl      string `json:"followers_url"`
		FollowingUrl      string `json:"following_url"`
		GistsUrl          string `json:"gists_url"`
		StarredUrl        string `json:"starred_url"`
		SubscriptionsUrl  string `json:"subscriptions_url"`
		OrganizationsUrl  string `json:"organizations_url"`
		ReposUrl          string `json:"repos_url"`
		EventsUrl         string `json:"events_url"`
		ReceivedEventsUrl string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
	NodeId          string        `json:"node_id"`
	TagName         string        `json:"tag_name"`
	TargetCommitish string        `json:"target_commitish"`
	Name            interface{}   `json:"name"`
	Draft           bool          `json:"draft"`
	Prerelease      bool          `json:"prerelease"`
	CreatedAt       time.Time     `json:"created_at"`
	PublishedAt     time.Time     `json:"published_at"`
	Assets          []interface{} `json:"assets"`
	TarballUrl      string        `json:"tarball_url"`
	ZipballUrl      string        `json:"zipball_url"`
	Body            interface{}   `json:"body"`
}

type Error struct {
	Message string `json:"message"`
	Errors  []struct {
		Resource string `json:"resource"`
		Code     string `json:"code"`
		Field    string `json:"field"`
	} `json:"errors"`
	DocumentationUrl string `json:"documentation_url"`
}

func main() {
	var token, tagName string

	flag.StringVar(&token, "token", "", "")
	flag.StringVar(&tagName, "tag", "", "")
	flag.Parse()

	url := "https://api.github.com/repos/skar404/alfred-translate/releases"
	urlUpload := "https://uploads.github.com/repos/skar404/alfred-translate/releases/"

	client := resty.New()
	doneRes := Done{}
	errRes := Error{}

	releases, err := client.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]string{"tag_name": tagName}).
		SetResult(&doneRes).
		SetError(&errRes).
		Post(url)

	if err != nil {
		log.Fatalln("error create releases", err)
	}

	statusCode := releases.StatusCode()
	if statusCode != http.StatusCreated {
		log.Fatalf("create releases, not valid status, code: %d", statusCode)
	}

	urlUpload = fmt.Sprintf("%s%d/assets", urlUpload, doneRes.Id)

	for _, file := range []string{"translate-amd64.alfredworkflow", "translate-arm64.alfredworkflow"} {
		uploadRes, err := client.R().
			SetQueryString(fmt.Sprintf("name=%s", file)).
			SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
			SetHeader("Content-Type", "application/json").
			SetFiles(map[string]string{
				file: file,
			}).
			Post(urlUpload)

		if err != nil {
			log.Fatalln("error upload file ", err)
		}

		statusCode = uploadRes.StatusCode()
		if statusCode != http.StatusCreated {
			log.Fatalf("upload file, not valid status, code: %d", statusCode)
		}
	}

	log.Printf("done create relise")
}
