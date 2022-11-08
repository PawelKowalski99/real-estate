package crawler

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/stealth"
	"math"
	"real-estate/core/entities"
	"strconv"
	"sync"
	"time"
)

const (
	//	Sprzedaz/Wynajem  Mieszkanie/Dom Miasto
	queryUrl        = `https://www.otodom.pl/pl/oferty/%s/%s/%s?page=%s&limit=200`
	SellMode        = "sprzedaz"
	RentMode        = "wynajem"
	CONCURRENT_JOBS = 15
)

var (
	expr = proto.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour).Unix())
)

type Crawler interface {
	GetEstates(mode, city, estateType string, buf chan []entities.Estate) error
}

type OtoDom struct {
	b        *rod.Browser
	PagePool chan *rod.Page

	domain string
}

func NewOtodom(br *rod.Browser) *OtoDom {
	return &OtoDom{
		b:        br,
		PagePool: make(chan *rod.Page, CONCURRENT_JOBS),
		domain:   "otodom.pl",
	}
}

func (s *OtoDom) GetEstates(mode, city, estateType string, buf chan []entities.Estate) error {
	p := stealth.MustPage(s.b)
	p.MustNavigate(fmt.Sprintf(queryUrl, mode, estateType, city, "1")).MustWaitLoad()

	defer func() {
		p.MustClose()
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic: ", r)
		}
	}()
	fmt.Println("1", p.MustInfo())
	p.MustElementX("//*[@id=\"__next\"]/div[2]/main/div/div[2]/div[1]/div[1]/div/div/div[1]/strong/span[2]").MustScrollIntoView()
	p.Mouse.Scroll(0, 500, 5)

	estCount, _ := strconv.Atoi(p.MustElementX("//*[@id=\"__next\"]/div[2]/main/div/div[2]/div[1]/div[1]/div/div/div[1]/strong/span[2]").MustText())
	pagesCount := math.Ceil(float64(estCount) / 200)

	var wg sync.WaitGroup

	for i := 1; i <= int(pagesCount); i++ {
		wg.Add(1)

		go func(pageNum int) {
			var estates []entities.Estate
			var pIn *rod.Page
			s.PagePool <- pIn

			for {
				var err error
				pIn = stealth.MustPage(s.b.MustIncognito())
				err = pIn.Navigate(fmt.Sprintf(queryUrl, mode, estateType, city, strconv.Itoa(pageNum)))
				if err != nil {
					time.Sleep(2 * time.Second)
					continue
				}

				pIn.MustWaitLoad()

				break
			}

			defer func() {
				wg.Done()

				<-s.PagePool
				pIn.MustClose()
				if r := recover(); r != nil {
					fmt.Println("Recovered from panic: ", r)
				}
			}()
			fmt.Println("running: ", pageNum)

			pIn.Mouse.Scroll(0, 1000, 20)

			fmt.Println(pageNum, pIn.MustInfo())

			var estatesElements rod.Elements
			for estatesElements.Empty() {
				var err error

				estatesElements, err = pIn.ElementsX(`//*[@data-cy='listing-item-link']`)

				if err != nil {
					fmt.Println("Could not search for estates elements: ", err)
				}
				el, _ := pIn.ElementX("//*[@id=\"__next\"]/div[2]/main/div/div[2]/div[1]/div[4]/div/nav")
				el.MustScrollIntoView()
				pIn.Mouse.Scroll(0, 1400, 20)

			}

			switch mode {
			case SellMode:
				for _, estateEl := range estatesElements {
					attrs := estateEl.MustElementsX("article/div[2]/span")
					estates = append(estates, entities.Estate{
						URL:        estateEl.MustProperty("href").String(),
						Address:    estateEl.MustElementX("article/p/span").MustText(),
						Price:      attrs[0].MustText(),
						PricePerM2: attrs[1].MustText(),
						RoomAmount: attrs[2].MustText(),
						Surface:    attrs[3].MustText(),
					})
				}
			case RentMode:
				for _, estateEl := range estatesElements {
					attrs := estateEl.MustElementsX("article/div[2]/span")
					estates = append(estates, entities.Estate{
						URL:        estateEl.MustProperty("href").String(),
						Address:    estateEl.MustElementX("article/p/span").MustText(),
						RentPrice:  attrs[0].MustText(),
						RoomAmount: attrs[1].MustText(),
						Surface:    attrs[2].MustText(),
					})
				}
			}
			buf <- estates
		}(i)
	}

	wg.Wait()
	return nil
}
