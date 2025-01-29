package handler

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/suryatresna/srg-radar-project/internal/scraper"
)

func (h *Handler) RunWorkerPatent() {
	patent := scraper.NewPatentScraper(h.collector)

	patent.GetEvent(func(data map[string]interface{}) {
		if err := h.repoPatent.StorePatentDb(data); err != nil {
			log.Error().Err(err).Msg("failed to store patent data")
			return
		}
		log.Info().Msg(fmt.Sprintf("store patent data: %v", data))
	})

	if err := patent.Scrapping(); err != nil {
		log.Error().Err(err).Msg("failed to scrap patent data")
		return
	}

	log.Info().Msg("worker command called")
}
