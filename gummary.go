package gummary

import (
	"strings"

	"github.com/DavidBelicza/TextRank"
	"github.com/DavidBelicza/TextRank/convert"
	"github.com/DavidBelicza/TextRank/parse"
	"github.com/PuerkitoBio/goquery"
)

// number of characters in p element to consider a content.
// remove stuffs like ads and attribution in p.
const paraLimit = 175

func Scrape(sel, url string) ([]string, error) {
	var items []string

	doc, err := goquery.NewDocument(url)
	if err != nil {
		return items, err
	}
	doc.Find(sel).Each(func(i int, s *goquery.Selection) {
		paragraph := strings.TrimSpace(s.Text())
		lastDot := strings.LastIndex(paragraph, ".")
		// Remove insufficient length paragraph and cut string after last fullstop
		// todo: fix getting tripped by decimal
		if lastDot >= paraLimit {
			item := string(paragraph[0 : lastDot+1])
			items = append(items, item)
		}
	})
	return items, err
}

func RankText(paragraphs []string) []string {
	ranked := []string{}

	tr := textrank.NewTextRank()
	language := textrank.NewDefaultLanguage()
	algorithmDef := textrank.NewDefaultAlgorithm()

	text := parse.Text{}
	for _, p := range paragraphs {
		// Get words from sentence
		for _, i := range strings.SplitAfter(p, ". ") {
			text.Append(i, strings.Fields(i))
		}
	}
	for _, sentence := range text.GetSentences() {
		convert.TextToRank(sentence, language, tr.GetRankData())
	}

	tr.Ranking(algorithmDef)
	sentences := textrank.FindSentencesByRelationWeight(tr, 4)

	// Put just the sentences in slice
	for _, s := range sentences {
		ranked = append(ranked, strings.TrimSpace(s.Value))
	}
	return ranked
}
