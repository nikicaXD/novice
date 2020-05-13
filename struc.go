package main

import "html/template"

type Front struct {
	Data struct {
		Front struct {
			Articles []struct {
				Id         int    `json:"id"`
				Title      string `json:"title"`
				FrontImage struct {
					Src string `json:"src"`
				} `json:"frontImage"`
				Comments struct {
					Total int `json:"total"`
				} `json:"comments"`
			} `json:"articles"`
		} `json:"front"`
	} `json:"data"`
	Extensions struct {
		RunTime int `json:"runTime"`
	} `json:"extensions"`
}

type Article struct {
	Data struct {
		Article struct {
			Title      string        `json:"title"`
			Summary    string        `json:"summary"`
			Subtitle   string        `json:"subtitle"`
			Body       string        `json:"body"`
			BodyHtml   template.HTML `json:"bodyHTML"`
			FrontImage struct {
				Src string `json:"src"`
			} `json:"frontImage"`
			BodyItems []struct {
				Type  string `json:"type"`
				Body  string `json:"body,omitempty"`
				Index int    `json:"index,omitempty"`
			} `json:"bodyItems"`
			Embeds []struct {
				ID   int    `json:"id"`
				Body string `json:"body"`
			} `json:"embeds"`
			Images []struct {
				ID      int    `json:"id"`
				Src     string `json:"src"`
				Caption string `json:"caption"`
				Type    string `json:"type"`
			} `json:"images"`
			Quotes []struct {
				Order  int    `json:"order"`
				Body   string `json:"body"`
				Author string `json:"author"`
				Type   int    `json:"type"`
			} `json:"quotes"`
			Videos []struct {
				ID     int    `json:"id"`
				Title  string `json:"title"`
				Source string `json:"source"`
			} `json:"videos"`
			Comments struct {
				Comments []struct {
					ID    string `json:"id"`
					Body  string `json:"body"`
					Owner struct {
						ID        int    `json:"id"`
						Nickname  string `json:"nickname"`
						AvatarURL string `json:"avatarUrl"`
					} `json:"owner"`
					Replies []struct {
						ID    string `json:"id"`
						Body  string `json:"body"`
						Owner struct {
							ID        int    `json:"id"`
							Nickname  string `json:"nickname"`
							AvatarURL string `json:"avatarUrl"`
						} `json:"owner"`
					} `json:"replies"`
				} `json:"comments"`
			} `json:"comments"`
		} `json:"article"`
	} `json:"data"`
	Extensions struct {
		RunTime int `json:"runTime"`
	} `json:"extensions"`
}

type Menu struct {
	Data struct {
		Menu struct {
			MenuItems []struct {
				SectionID int    `json:"sectionId"`
				Title     string `json:"title"`
			} `json:"menuItems"`
		} `json:"menu"`
	} `json:"data"`
	Extensions struct {
		RunTime int `json:"runTime"`
	} `json:"extensions"`
}
