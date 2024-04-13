package services

import (
	"encoding/json"
	"git.nlmk.com/mcs/micro/portal/portal_sync/models"
	"github.com/antchfx/htmlquery"
	"github.com/lib/pq"
	"regexp"
	"strconv"
	"strings"
)

func ConvertDescriptions(text string) (pq.Int64Array, string, *models.ListOfDescriptionDB, error) {
	var descriptions []*models.DescriptionDB
	var list *models.ListOfDescriptionDB

	doc, err := htmlquery.Parse(strings.NewReader(text))
	if err != nil {
		return pq.Int64Array{}, text, nil, err
	}

	for _, n := range htmlquery.Find(doc, `//img`) {
		height, err := strconv.Atoi(htmlquery.SelectAttr(n, "height"))
		if err != nil {
			height = 0
		}
		width, err := strconv.Atoi(htmlquery.SelectAttr(n, "width"))
		if err != nil {
			width = 0
		}
		item := &models.DescriptionDB{
			Type:   "img",
			Src:    htmlquery.SelectAttr(n, "src"),
			Title:  htmlquery.SelectAttr(n, "title"),
			Height: height,
			Width:  width,
		}
		if height > 16 && width > 16 {
			descriptions = append(descriptions, item)
		}
	}

	for _, n := range htmlquery.Find(doc, `//video`) {
		item := &models.DescriptionDB{
			Type:     "video",
			Src:      htmlquery.SelectAttr(n, "src"),
			TypeFile: htmlquery.SelectAttr(n, "type"),
		}
		descriptions = append(descriptions, item)
	}

	text, voteIds := getTags(text, `\$vote_\d+\$`)

	for _, n := range voteIds {
		descriptions = append(descriptions, &models.DescriptionDB{
			Type: "vote",
			Id:   n,
		})
	}
	text, formIds := getTags(text, `\#FORM_ID_\d+\#`)
	var forms pq.Int64Array
	for _, n := range formIds {
		descriptions = append(descriptions, &models.DescriptionDB{
			Type: "form",
			Id:   n,
		})
		forms = append(forms, int64(n))
	}
	descriptions = append(descriptions, &models.DescriptionDB{
		Type: "text",
		Html: text,
	})
	marshal, err := json.Marshal(descriptions)
	if err != nil {
		return pq.Int64Array{}, text, nil, err
	}

	if err := json.Unmarshal(marshal, &list); err != nil {
		return pq.Int64Array{}, text, nil, err
	}
	return forms, text, list, nil
}

func getTags(text string, mask string) (string, []int) {
	var items []int
	re, _ := regexp.Compile(mask)
	for _, value := range re.FindAllString(text, -1) {
		re, _ := regexp.Compile(`\d+`)
		for _, numStr := range re.FindAllString(value, -1) {
			num, err := strconv.Atoi(numStr)
			if err == nil {
				items = append(items, num)
			}
		}

	}
	return re.ReplaceAllString(text, ""), items
}
