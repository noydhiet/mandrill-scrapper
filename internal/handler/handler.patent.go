package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	modelPatent "github.com/noydhiet/mandrill-scrapper/internal/repository/patent"
	"github.com/noydhiet/mandrill-scrapper/internal/scraper"
	"github.com/rs/zerolog/log"
)

func (h *Handler) RunWorkerPatent() {
	ctx := context.Background()
	patent := scraper.NewPatentScraper(h.collector)

	patent.GetEvent(func(data map[string]interface{}) {

		patentData := modelPatent.PatentDB{}

		if val, ok := data["company_name"]; ok {
			patentData.CompanyName = val.(string)
		}

		if val, ok := data["patent_expiry_date"]; ok {
			patentData.PatentExpiryDate = val.(string)
		}

		if val, ok := data["patent_number"]; ok {
			patentData.PatentNumber = val.(string)
		}

		if val, ok := data["title"]; ok {
			patentData.Title = val.(string)
		}

		if err := h.repoPatent.StorePatent(ctx, patentData); err != nil {
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

func (h *Handler) HandleGetPatentData(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	data, err := h.repoPatent.FindPatent(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to get patent data")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "{\"error\": \"failed to get patent data\"}", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
