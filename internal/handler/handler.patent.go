package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	modelPatent "github.com/noydhiet/mandrill-scrapper/internal/repository/patent"
	"github.com/noydhiet/mandrill-scrapper/internal/scraper"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (h *Handler) RunWorkerPatent() {
	ctx := context.Background()
	patent := scraper.NewPatentScraper(h.collector)

	patent.GetEvent(func(data map[string]interface{}) {

		patentData := modelPatent.PatentDB{
			CreatedAt: time.Now(),
		}

		if val, ok := data["Company Name"]; ok {
			patentData.CompanyName = val.(string)
		}

		if val, ok := data["Patent Expiry"]; ok {
			patentData.PatentExpiryDate = val.(string)
		}

		if val, ok := data["Patent Number"]; ok {
			patentData.PatentNumber = val.(string)
		}

		if val, ok := data["Patent Title"]; ok {
			patentData.Title = val.(string)
		}

		if err := h.repoPatent.StorePatent(ctx, patentData); err != nil {
			log.Error().Err(err).Msg("failed to store patent data")
			return
		}
		log.Info().Msg(fmt.Sprintf("store patent data: %v", patentData))
	})

	if err := patent.Scrapping(); err != nil {
		log.Error().Err(err).Msg("failed to scrap patent data")
		return
	}

	log.Info().Msg("worker command called")
}

func (h *Handler) HandleGetPatentData(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	query := r.URL.Query().Get("q")
	filter := bson.M{}
	if query != "" {
		filter = bson.M{"CompanyName": bson.M{"$regex": query, "$options": "i"}}
	}

	data, err := h.repoPatent.FindPatent(ctx, filter)
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
