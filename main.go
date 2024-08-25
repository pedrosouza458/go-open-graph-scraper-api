package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pedrosouza458/go-open-graph-scraper/pkg/scraper"
)

func main() {
	http.HandleFunc("/websites", websiteHandler)
	http.ListenAndServe(":8080", nil)
}

type Response struct {
	Logo       string `json:"logo"`
	StatucCode int    `json:"statusCode"`
}

func websiteHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*") // Adjust the origin as needed
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	url := r.URL.Query().Get("url")
	if url != "" {
		logo, err := scraper.GetWebsiteLogo(url)
		if err != nil {
			fmt.Fprintf(w, "Error fetching logo: %v\n", err)
			return
		}
		name, err := scraper.GetWebsiteName(url)
		if err != nil {
			fmt.Fprintf(w, "Error fetching name: %v\n", err)
			return
		}
		img, err := scraper.GetWebsiteImg(url)
		if err != nil {
			fmt.Fprintf(w, "Error fetching img: %v\n", err)
			return
		}
		title, err := scraper.GetWebsitePageTitle(url)
		if err != nil {
			fmt.Fprintf(w, "Error fetching website title: %v\n", err)
			return
		}
		description, err := scraper.GetWebsiteDescription(url)
		if err != nil {
			fmt.Fprintf(w, "Error fetching description: %v\n", err)
			return
		}
		type_name, err := scraper.GetWebsiteType(url)
		if err != nil {
			fmt.Fprintf(w, "Error fetching type: %v\n", err)
			return
		}
		video, err := scraper.GetWebsiteVideo(url)
		if err != nil {
			fmt.Fprintf(w, "Error fetching vudei: %v\n", err)
			return
		}
		locale, err := scraper.GetWebsiteLocale(url)
		if err != nil {
			fmt.Fprintf(w, "Error fetching type: %v\n", err)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{
			"logo":        logo,
			"name":        name,
			"img":         img,
			"video":       video,
			"title":       title,
			"description": description,
			"type":        type_name,
			"locale":      locale,
		})
	}
}
