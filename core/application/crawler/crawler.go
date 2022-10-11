package crawler

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"net/http"
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

type EstateCrawler interface {
	GetEstates(pageNum int) (int, []entities.Estate)
}

type OtodomCrawler struct {
	count int

	b *rod.Browser

	domain string

	mode string

	estateType string

	city string
}

func NewBrowser(br *rod.Browser, d, mode, city, estateType string) *OtodomCrawler {
	return &OtodomCrawler{b: br, domain: d, mode: mode, city: city, estateType: estateType}
}

func (s *OtodomCrawler) GetEstates(pageNum int) (int, []entities.Estate){
	p := s.b.MustPage(fmt.Sprintf(queryUrl, s.mode, s.estateType, s.city, strconv.Itoa(pageNum))).MustWaitLoad()
	fmt.Println(pageNum, p.MustInfo())

	estatesElements := p.MustElementsX("//*[@id=\"__next\"]/div[2]/main/div/div[2]/div[1]/div[2]/div[1]/div/ul/li")
	//fmt.Println(estatesElements)

	estatesElements = append(estatesElements, p.
		MustElementsX("//*[@id=\"__next\"]/div[2]/main/div/div[2]/div[1]/div[2]/div/ul/li")...)

	var estatesHref []entities.Estate
	for _, estateEl := range estatesElements {
		attrs := estateEl.MustElementsX("a/article/div[1]/span")
		fmt.Println(estateEl.MustElementX("/a").MustHTML())
		estatesHref = append(estatesHref, entities.Estate{
			URL: "",
			Address:              estateEl.MustElementX("article/p/span").MustText(),
			Price:                attrs[0].MustText(),
			PricePerM2:           attrs[1].MustText(),
			RoomAmount:           attrs[2].MustText(),
			Surface:              attrs[3].MustText(),
		})
	}

	// Recursive navigation to the last page
	if p.
		MustElementsX("//*[@id=\"__next\"]/div[2]/main/div/div[2]/div[1]/div[4]/div/nav/button").Last().
		MustAttribute("disabled") == nil {
			pageNum++
			_, estatesHrefTemp := s.GetEstates(pageNum)
			estatesHref = append(estatesHrefTemp, estatesHrefTemp...)
	}


	return http.StatusOK, estatesHref
}

func (s *OtodomCrawler) GetEstatesLinks() []string {
	return []string{}
}


func(s *OtodomCrawler) crawlEstate() entities.Estate {


	return entities.Estate{}
}





