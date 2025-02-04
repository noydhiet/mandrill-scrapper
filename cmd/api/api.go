/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package api

import (
	"net/http"
	"os"

	"github.com/gocolly/colly"
	"github.com/noydhiet/mandrill-scrapper/internal/handler"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	pkgStorage "github.com/noydhiet/mandrill-scrapper/internal/pkg/storage"
	repoLawsuit "github.com/noydhiet/mandrill-scrapper/internal/repository/lawsuit"
	repoManufacture "github.com/noydhiet/mandrill-scrapper/internal/repository/manufacture"
	repoPatent "github.com/noydhiet/mandrill-scrapper/internal/repository/patent"
	repoRecall "github.com/noydhiet/mandrill-scrapper/internal/repository/recall"
	repoRegistration "github.com/noydhiet/mandrill-scrapper/internal/repository/registration"
)

// searchCmd represents the search command
var ApiCmd = &cobra.Command{
	Use:   "api",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runSearchAPI,
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ApiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ApiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runSearchAPI(cmd *cobra.Command, args []string) {
	log.Info().Msg("worker initiated")
	collector := colly.NewCollector()

	mongoDsn := os.Getenv("MONGO_DSN")

	storage, err := pkgStorage.NewStorageMongo(mongoDsn)
	if err != nil {
		log.Error().Err(err).Msg("failed to connect to database")
		return
	}
	patentDb := repoPatent.NewRepository(storage)
	manufactureDb := repoManufacture.NewRepository(storage)
	lawsuitDb := repoLawsuit.NewRepository(storage)
	registrationDb := repoRegistration.NewRepository(storage)
	recallDb := repoRecall.NewRepository(storage)

	hdl := handler.NewHandler(collector, patentDb, lawsuitDb, manufactureDb, recallDb, registrationDb)

	http.HandleFunc("/v1/search", hdl.HandleGetPatentData)

	log.Info().Msg("api server started at :8080")
	http.ListenAndServe(":8080", nil)
}
