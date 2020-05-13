package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"encoding/json"
)

const NEWS="http://gql.24ur.si/news"
const NEWSOTHER= "http://gql.24ur.si/news-other"



func getFrontQuery(sectionId string) string {

	Q := `
{
front(siteId: 1 sectionId: ##sectionId##) {
 articles {
   id
   title
   frontImage {
     src
   }
 comments {
     total
   }
 }
}
}
`
	Q = strings.Replace(Q, "##sectionId##", sectionId, 1)
	return Q
}

func getMenuQuery() string {
	return `{
		menu(version:"mobile" siteId:1) {
	menuItems {
	sectionId
	title

	}
	}
	}`

}

func getArticleQuery(artId string) string {
	Q := `
{
  article(id: ##articleId##) {
    title
    summary
    body
    subtitle
    frontImage {
      src
    }
    bodyItems
    embeds {
      id
      body
    }
    images {
      id
      src
      caption
      type
    }
    quotes {
      order
      body
      author
      type
    }
    videos {
      id
      title
      source
    }
    comments {
      comments {
        id
        body
        owner {
          id
          nickname
          avatarUrl
        }
        replies {
          id
          body
          owner {
            id
            nickname
            avatarUrl
          }
        }
      }
    }
  }
}
`
	Q = strings.Replace(Q, "##articleId##", artId, 1)
	return Q

}

func menuGeter(menuItems string)(string) {

	DataJSON, menuItems:= getMenuQuery()
	Data := Menu{}
	json.Unmarshal([]byte(DataJSON), &Data)

	return string(menuItems)



}


func dataGeter(query string, url string) (string, int) {


	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(query)))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/graphql")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return "", 500
	}

	if resp.StatusCode != 200 {
		return "", resp.StatusCode
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), resp.StatusCode

}
