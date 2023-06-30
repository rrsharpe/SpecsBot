package ssd

import (
	"fmt"
	"github.com/deckarep/golang-set"
	"github.com/turnage/graw/reddit"
	"strings"
)

func conditionallyComment(description string, bot reddit.Bot, p reddit.Post, modelMap map[modelKey][]string) {
	fmt.Printf("[SSD] Tagged post: [%s] posted [%s]\n", p.Author, p.Title)
	// Don't do anything if the bot has already commented
	commentPost, err := bot.Thread(p.Permalink)
	if err != nil {
		fmt.Println("Failed to fetch comments", err)
		return
	}
	for _, comment := range commentPost.Replies {
		if comment.Author == "SpecsBot" {
			fmt.Println("Found existing comment... Skipping!")
			return
		}
	}

	maxMatching, matchingRows := getMatching(description, modelMap)

	if len(matchingRows) == 1 && maxMatching >= 1000 {
		title := matchingRows[0][0] + " " + matchingRows[0][1]
		altNames := []string{matchingRows[0][15], matchingRows[0][16]}
		filteredAltNames := []string{}
		for _, altName := range altNames {
			if altName != "" && altName != title {
				filteredAltNames = append(filteredAltNames, altName)
			}
		}

		// Actually leave the comment
		comment := genComment(Comment{
			Confident:     maxMatching >= 10000,
			Title:         title,
			Interface:     matchingRows[0][2],
			FormFactor:    matchingRows[0][3],
			Capacities:    matchingRows[0][4],
			Controller:    matchingRows[0][5],
			Configuration: matchingRows[0][6],
			DRAM:          matchingRows[0][7],
			HMB:           matchingRows[0][8],
			NANDBrand:     matchingRows[0][9],
			NANDType:      matchingRows[0][10],
			Layers:        matchingRows[0][11],
			ReadWrite:     matchingRows[0][12],
			Categories:    matchingRows[0][13],
			Notes:         matchingRows[0][14],
			AltNames:      filteredAltNames,
		})
		fmt.Println("\t* Matched to", title)
		bot.Reply(p.Name, comment)
	} else {
		fmt.Println("Not confident enough to post!")
	}
}

func getMatching(description string, modelMap map[modelKey][]string) (int, [][]string) {
	words := strings.FieldsFunc(description, argMatchesDelimiters)
	descriptionSet := mapset.NewSet()
	for _, word := range words {
		descriptionSet.Add(strings.ToUpper(word))
	}

	maxMatching := 0
	var matchingRows [][]string = [][]string{}

	for keySet, value := range modelMap {
		numBrandMatching := len(keySet.brand.Intersect(descriptionSet).ToSlice()) * 10000 // brand > model > product pages
		numModelMatching := len(keySet.model.Intersect(descriptionSet).ToSlice()) * 1000
		numProductPagesMatching := len(keySet.productPages.Intersect(descriptionSet).ToSlice()) * 100
		numOthersMatching := len(keySet.others.Intersect(descriptionSet).ToSlice()) * 10  // Then others
		modelLenReduction := 10 - len(strings.FieldsFunc(value[1], argMatchesDelimiters)) // A model [A B] and [A] might tie if given just A. More model matches still wins
		numMatching := numBrandMatching + numModelMatching + numOthersMatching + numProductPagesMatching + modelLenReduction
		if numMatching > maxMatching {
			// fmt.Println("A", keySet.model, numBrandMatching, numModelMatching, numOthersMatching, numProductPagesMatching, modelLenReduction)
			matchingRows = [][]string{value}
			maxMatching = numMatching
		} else if numMatching == maxMatching {
			// fmt.Println("B", keySet.model, numBrandMatching, numModelMatching, numOthersMatching, numProductPagesMatching, modelLenReduction)
			matchingRows = append(matchingRows, value)
		}
	}
	fmt.Printf("\t%d matches with score: %d: %s\n", len(matchingRows), maxMatching, matchingRows)
	return maxMatching, matchingRows
}

func (b *SSDPostHandler) Post(p *reddit.Post) error {
	tagCloseIndex := strings.Index(p.Title, "]")
	priceStart := strings.Index(p.Title, "(") // Technically may be earlier than the price
	if tagCloseIndex <= 1 || priceStart < tagCloseIndex {
		return nil
	}

	postTag := strings.ToUpper(p.Title[1:tagCloseIndex])
	description := p.Title[tagCloseIndex+1 : priceStart]

	if strings.Contains(postTag, "SSD") || strings.Contains(postTag, "NVME") {
		conditionallyComment(description, b.bot, *p, b.ssdModelMap)
	}
	return nil
}
