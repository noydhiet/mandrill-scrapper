package handler

import (
	"github.com/gocolly/colly"
)

type Handler struct {
	collector        *colly.Collector
	repoPatent       repoPatentInterface
	repoLawsuit      repoLawsuitInterface
	repoManufacturer repoManufacturerInterface
	repoRecall       repoRecallInterface
	repoRegistration repoRegistrationInterface
}

func NewHandler(
	collector *colly.Collector,
	repoPatent repoPatentInterface,
	repoLawsuit repoLawsuitInterface,
	repoManufacturer repoManufacturerInterface,
	repoRecall repoRecallInterface,
	repoRegistration repoRegistrationInterface,
) *Handler {
	hdl := &Handler{
		collector:        collector,
		repoPatent:       repoPatent,
		repoLawsuit:      repoLawsuit,
		repoManufacturer: repoManufacturer,
		repoRecall:       repoRecall,
		repoRegistration: repoRegistration,
	}

	return hdl
}

func (h *Handler) Worker() {

}
