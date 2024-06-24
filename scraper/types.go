package scraper

import (
	"fmt"
	"strings"
	"strconv"
)

type Map struct {
	Image    string
	Name     string
	Brawlers []Brawler
}

type Brawler struct {
	Name      string
	WinRate   string
	UseRate   string
	StarpRate string
}

func (b *Brawler) IntoInt() (float64, int64, float64) {
	wr := strings.Replace(b.WinRate, "%", "", 1)
	ur := strings.Replace(b.UseRate, "#", "", 1)
	spr := strings.Replace(b.StarpRate, "%", "", 1)

	wri, wrerr := strconv.ParseFloat(wr, 32)
	uri, urerr := strconv.ParseInt(ur, 10, 32)
	spri, sprerr := strconv.ParseFloat(spr, 32)

	if wrerr != nil || urerr != nil || sprerr != nil {
		return 0.0, 0, 0.0
	} else {
		return wri, uri, spri
	}
}

func (b *Brawler) UseRateInt() int64 {
	_, ur, _ :=b.IntoInt()
	return ur
}

func (m *Map) Display() (bool, string) {
	message := make([]string, 0)
	message = append(message, fmt.Sprintf("Information for map %s", m.Name))

	if len(m.Brawlers) == 0 {
		return false, "Map wasn`t found"
	} else {
		for _, brawler := range m.Brawlers {
			wr, ur, spr := brawler.IntoInt()
			message = append(message, fmt.Sprintf("%d. %s (WR: %.2f, SPR: %.2f)", ur, brawler.Name, wr, spr))
		}
		return true, strings.Join(message, "\n")
	}
}