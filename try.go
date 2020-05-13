package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"bytes"
	"encoding/json"
	"html/template"
	"strings"
)

var frontTemplate *template.Template
var articleTemplate *template.Template
var articleBodyPart *template.Template
var articleImagePart *template.Template
var articleEmbedPart *template.Template
var articleQuotePart *template.Template
var menu *template.Template
var err error

func ArticleHandler(w http.ResponseWriter, r *http.Request) {

	var body bytes.Buffer
	vars := mux.Vars(r)
	DataJSON, statusCode := dataGeter(getArticleQuery(vars["ArticleId"]), NEWS)
	Data := Article{}
	json.Unmarshal([]byte(DataJSON), &Data)
	if statusCode != 200 {
		w.Write([]byte("Napaka"))

	} else {

		imgURL := Data.Data.Article.FrontImage.Src
		imgURL = strings.Replace(imgURL, "PLACEHOLDER", "900x200", 1)
		Data.Data.Article.FrontImage.Src = imgURL
		for _, bp := range Data.Data.Article.BodyItems {
			switch bp.Type {
			case "html":
				articleBodyPart.Execute(&body, template.HTML(bp.Body))

			case "image":
				bodyStruct := Data.Data.Article.Images[bp.Index]
				bodyStruct.Src = strings.Replace(bodyStruct.Src, "PLACEHOLDER", "500x300", 1)
				articleImagePart.Execute(&body, bodyStruct)

			case "embed":
				articleEmbedPart.Execute(&body, template.HTML(Data.Data.Article.Embeds[bp.Index].Body))

			case "quote":
				articleQuotePart.Execute(&body, template.HTML(Data.Data.Article.Quotes[bp.Index].Body))

			}

		}

		Data.Data.Article.BodyHtml = template.HTML(body.String())
		articleTemplate.Execute(w, Data.Data.Article)

	}

}

func FrontHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	sectionId := vars["sectionId"]
	if sectionId == "" {
		sectionId = "1"
	}
	DataJSON, statusCode := dataGeter(getFrontQuery(sectionId), NEWS)
	Data := Front{}
	json.Unmarshal([]byte(DataJSON), &Data)

	if statusCode != 200 {
		w.Write([]byte("Napaka"))

	} else {
		for i, a := range Data.Data.Front.Articles {
			imgURL := a.FrontImage.Src
			imgURL = strings.Replace(imgURL, "PLACEHOLDER", "500x300", 1)
			Data.Data.Front.Articles[i].FrontImage.Src = imgURL
		}
		frontTemplate.Execute(w, Data.Data.Front)
		menuGeter()
	}

}

func init() {
	frontTemplate, err = template.ParseFiles("templates/front.html")
	if err != nil {
		log.Panic(err.Error())

	}

	articleTemplate, err = template.ParseFiles("templates/article.html")
	if err != nil {
		log.Panic(err.Error())

	}

	articleBodyPart, err = template.ParseFiles("templates/sub/bodyPart.html")
	if err != nil {
		log.Panic(err.Error())

	}
	articleImagePart, err = template.ParseFiles("templates/sub/bodyImagePart.html")
	if err != nil {
		log.Panic(err.Error())

	}

	articleEmbedPart, err = template.ParseFiles("templates/sub/bodyEmbedPart.html")
	if err != nil {
		log.Panic(err.Error())

	}

	articleQuotePart, err = template.ParseFiles("templates/sub/bodyQuotePart.html")
	if err != nil {
		log.Panic(err.Error())

	}
	menu, err = template.ParseFiles("templates/sub/menu.html")
	if err != nil {
		log.Panic(err.Error())

	}

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/article/{ArticleId}", ArticleHandler)
	r.HandleFunc("/{sectionId}", FrontHandler)
	r.HandleFunc("/", FrontHandler)

	log.Fatal(http.ListenAndServe(":8080", r))

}
