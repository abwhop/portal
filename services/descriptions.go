package services

import (
	"encoding/json"
	"github.com/abwhop/portal_models/models"
	"github.com/antchfx/htmlquery"
	"github.com/lib/pq"
	"regexp"
	"strconv"
	"strings"
)

func ConvertDescriptions(text string) (pq.Int64Array, string, *models.ListOfDescriptionDB, error) {
	var descriptions []*models.DescriptionDB
	var list *models.ListOfDescriptionDB

	text = strings.ReplaceAll(text, "<img", "9bb92469479a4edbab8ea4fbb9e00e73 <img")
	text = strings.ReplaceAll(text, "/>", "/> 9bb92469479a4edbab8ea4fbb9e00e73")
	text = strings.ReplaceAll(text, "<video", "9bb92469479a4edbab8ea4fbb9e00e73 <video")
	text = strings.ReplaceAll(text, "</video>", "</video> 9bb92469479a4edbab8ea4fbb9e00e73")
	text = strings.ReplaceAll(text, "<file", "9bb92469479a4edbab8ea4fbb9e00e73 <file")
	text = strings.ReplaceAll(text, "</file>", "</file> 9bb92469479a4edbab8ea4fbb9e00e73")
	text = regexp.MustCompile(`\$vote_\d+\$`).ReplaceAllString(text, "9bb92469479a4edbab8ea4fbb9e00e73 ${0} 9bb92469479a4edbab8ea4fbb9e00e73")
	text = regexp.MustCompile(`#FORM_ID_\d+`).ReplaceAllString(text, "9bb92469479a4edbab8ea4fbb9e00e73 ${0} 9bb92469479a4edbab8ea4fbb9e00e73")
	newsBody := strings.Split(text, "9bb92469479a4edbab8ea4fbb9e00e73")
	for _, body := range newsBody {
		str := strings.TrimSpace(strings.Trim(body, "\n"))
		if str == "" || str == "<br />" {
			continue
		}
		descriptions = append(descriptions, parser(str))

	}
	/*doc, err := htmlquery.Parse(strings.NewReader(text))
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
	}*/
	text, formIds := getTags(text, `\#FORM_ID_\d+\#`)
	var forms pq.Int64Array
	for _, n := range formIds {
		/*descriptions = append(descriptions, &models.DescriptionDB{
			Type: "form",
			Id:   n,
		})*/
		forms = append(forms, int64(n))
	}
	/*descriptions = append(descriptions, &models.DescriptionDB{
		Type: "text",
		Html: text,
	})*/
	marshal, err := json.Marshal(descriptions)
	if err != nil {
		return pq.Int64Array{}, text, nil, err
	}

	if err := json.Unmarshal(marshal, &list); err != nil {
		return pq.Int64Array{}, text, nil, err
	}
	return forms, text, list, nil
}
func parser(content string) *models.DescriptionDB {

	doc, err := htmlquery.Parse(strings.NewReader(content))
	if err != nil {
		return nil
	}
	if n := htmlquery.FindOne(doc, `//img`); n != nil {
		height, err := strconv.Atoi(htmlquery.SelectAttr(n, "height"))
		if err != nil {
			height = 0
		}
		width, err := strconv.Atoi(htmlquery.SelectAttr(n, "width"))
		if err != nil {
			width = 0
		}

		if height > 16 && width > 16 {
			return &models.DescriptionDB{
				Type:   "img",
				Src:    htmlquery.SelectAttr(n, "src"),
				Title:  htmlquery.SelectAttr(n, "title"),
				Height: height,
				Width:  width,
			}
		}
	}

	if n := htmlquery.FindOne(doc, `//video`); n != nil {
		return &models.DescriptionDB{
			Type:     "video",
			Src:      htmlquery.SelectAttr(n, "src"),
			TypeFile: htmlquery.SelectAttr(n, "type"),
		}
	}

	if match, err := regexp.MatchString(`\$vote_\d+\$`, content); err != nil {
		return nil
	} else if match {
		first := regexp.MustCompile(`\d+`).FindStringSubmatch(content)
		if len(first) < 1 {
			return nil
		}
		id, err := strconv.Atoi(first[0])
		if err != nil {
			return nil
		}
		return &models.DescriptionDB{
			Type: "vote",
			Id:   id,
		}
	}
	if match, err := regexp.MatchString(`#FORM_ID_\d+`, content); err != nil {
		return nil
	} else if match {
		first := regexp.MustCompile(`\d+`).FindStringSubmatch(content)
		if len(first) < 1 {
			return nil
		}
		id, err := strconv.Atoi(first[0])
		if err != nil {
			return nil
		}
		return &models.DescriptionDB{
			Type: "form",
			Id:   id,
		}
	}

	return &models.DescriptionDB{
		Type: "text",
		Html: content,
	}
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
