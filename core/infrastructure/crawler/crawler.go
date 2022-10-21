package crawler

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/labstack/gommon/log"
	"real-estate/core/entities"
	"strconv"
	"time"
)

const (
	//	Sprzedaz/Wynajem  Mieszkanie/Dom Miasto
	queryUrl = `https://www.otodom.pl/pl/oferty/%s/%s/%s?page=%s`
)
var (
	expr = proto.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour).Unix())
)

type Crawler interface {
	GetEstates(mode, city, estateType string, pageNum int) ([]entities.Estate, error)
}

type OtoDom struct {
	count int

	b *rod.Browser

	domain string
}

func NewOtodom(br *rod.Browser) *OtoDom {
	return &OtoDom{
		b:          br,
		domain:     "otodom.pl",
	}
}

func (s *OtoDom) GetEstates(mode, city, estateType string, pageNum int) (estates []entities.Estate, err error){
	p := s.b.MustPage(fmt.Sprintf(queryUrl, mode, estateType, city, strconv.Itoa(pageNum))).MustWaitLoad()
	fmt.Println(pageNum, p.MustInfo())

	estatesElements := p.MustElementsX("//*[@id=\"__next\"]/div[2]/main/div/div[2]/div[1]/div[2]/div[1]/div/ul/li/a")

	estatesElements = append(estatesElements,
		p.MustElementsX("//*[@id=\"__next\"]/div[2]/main/div/div[2]/div[1]/div[2]/div/ul/li/a")...)

	for _, estateEl := range estatesElements {
		attrs := estateEl.MustElementsX("article/div[2]/span")
		estates = append(estates, entities.Estate{
			URL:         estateEl.MustProperty("href").String(),
			Address:     estateEl.MustElementX("article/p/span").MustText(),
			Price:       attrs[0].MustText(),
			PricePerM2:  attrs[1].MustText(),
			RoomAmount:  attrs[2].MustText(),
			Surface:     attrs[3].MustText(),
		})
	}


	// Recursive navigation to the last page
	if p.
		MustElementsX("//*[@id=\"__next\"]/div[2]/main/div/div[2]/div[1]/div[4]/div/nav/button").Last().
		MustAttribute("disabled") == nil {
			pageNum++
			estatesHrefTemp, err2 := s.GetEstates(mode, city, estateType,pageNum)
			if err2 != nil {
				log.Infof(err2.Error())
				return nil, err2
			}
		estates = append(estatesHrefTemp, estatesHrefTemp...)
	}


	return estates, nil
}




