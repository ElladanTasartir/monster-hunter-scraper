package http

import (
	"fmt"
	"github.com/ElladanTasartir/monster-hunter-scraper/internal/scraper"
	"github.com/ElladanTasartir/monster-hunter-scraper/internal/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ScraperEndpoint struct {
	scraper scraper.Scraper
	storage *storage.Storage
}

func (s *Server) NewScraperEndpoint(scraperURL string) error {
	monhunScraper, err := scraper.NewMonsterHunterScraper(scraperURL)
	if err != nil {
		return err
	}

	endpoint := &ScraperEndpoint{
		scraper: monhunScraper,
		storage: s.storage,
	}

	s.httpServer.GET("/monsters", endpoint.FindMonsters)
	s.httpServer.POST("/scrape", endpoint.ScrapeData)

	return nil
}

func (e *ScraperEndpoint) ScrapeData(ctx *gin.Context) {
	storedMonsters, err := e.storage.FindMonsters(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error has ocurred while querying monsters. err = %v", err),
		})
		return
	}

	if len(storedMonsters) > 0 {
		ctx.JSON(http.StatusOK, storedMonsters)
		return
	}

	response, err := e.scraper.Scrape()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error has ocurred while scraping data. err = %v", err),
		})
		return
	}

	monsters, ok := response.([]scraper.Monster)
	if !ok {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": fmt.Sprintf("error has ocurred while parsing data. existing = %v", ok),
		})
		return
	}

	storageMonsters, err := e.storage.CreateMonsters(ctx, monsters)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": fmt.Sprintf("error has ocurred while inserting data. err = %v", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, storageMonsters)
}

func (e *ScraperEndpoint) FindMonsters(ctx *gin.Context) {
	storedMonsters, err := e.storage.FindMonsters(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error has ocurred while querying monsters. err = %v", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, storedMonsters)
}
