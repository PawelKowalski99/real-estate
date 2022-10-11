package app_todo

import (
	"context"
	"github.com/go-rod/rod"
	"github.com/labstack/echo/v4"
	"real-estate/core/application/crawler"
)

// The HTTP Handler for TODO
type CrawlerHttpService struct {
	gtw crawler.EstateCrawler
}

func (t *CrawlerHttpService) GetEstates(c echo.Context) error {
	status, res := t.gtw.GetEstates(1)

	return c.JSON(status, res)
}



// Constructor
func NewCrawlerHttpService(ctx context.Context, browser *rod.Browser) *CrawlerHttpService {
	return &CrawlerHttpService{crawler.NewBrowser(browser, "otodom.pl", "sprzedaz", "kobierzyce", "mieszkanie")}
}
