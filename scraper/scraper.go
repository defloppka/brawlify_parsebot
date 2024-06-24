package scraper

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
	"github.com/gocolly/colly/v2"
)

const baseUrl string = "https://brawlify.com/maps/detail/"

func GetMapInfo(name string) Map {
	c := colly.NewCollector()
	mapInfo := Map{}
	brawlers := make([]Brawler, 0)

	c.OnHTML("div#content-container", func(h *colly.HTMLElement) {
		mapInfo.Image = h.ChildAttr("img#mapImg", "src")
		mapInfo.Name = h.ChildAttr("img#mapImg", "title")
		h.ForEach("div#brawlers div.d-flex.justify-content-center.p-1", func(i int, h *colly.HTMLElement) {
			brawler := ParseBrawler(h, i)
			brawlers = append(brawlers, brawler)
		})
		sort.Slice(brawlers[:], func(i, j int) bool {
			return brawlers[i].UseRateInt() < brawlers[j].UseRateInt()
		})
		mapInfo.Brawlers = brawlers
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	url := baseUrl + ParseName(name)
	c.Visit(url)

	return mapInfo
}

func ParseBrawler(el *colly.HTMLElement, i int) Brawler  {
	name := el.ChildAttrs("div.d-flex.flex-column.justify-content-center a.link.opacity", "title")[0]
	
	cls := fmt.Sprintf("div.font-rank%d.small", i + 1)
	wr := el.ChildTexts(cls)[0]
	ur := el.ChildTexts("div.text-primary.small")[0]
	spr := el.ChildTexts("div.text-orange.small")[0]

	return Brawler {
		Name: name,
		WinRate: wr,
		UseRate: ur,
		StarpRate: spr,
	}
}

func ParseName(name string) string {
	titled := MakeTitle(name);
	newName := strings.ReplaceAll(titled, " ", "-")
	fmt.Println("Name: ", newName)
	return newName
}

func MakeTitle(str string) string {
	var output []rune    
	isWord := true
	for _, val := range str {
		if isWord && unicode.IsLetter(val) {  
			output = append(output, unicode.ToUpper(val))
			isWord = false
		} else if !unicode.IsLetter(val) {
			isWord = true
			output = append(output, val)
		} else {
			output = append(output, val)
		}
	}
	return string(output)
}