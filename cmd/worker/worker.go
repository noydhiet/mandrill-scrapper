package worker

import (
	"os"

	"github.com/gocolly/colly"
	"github.com/noydhiet/mandrill-scrapper/internal/handler"
	pkgStorage "github.com/noydhiet/mandrill-scrapper/internal/pkg/storage"
	repoLawsuit "github.com/noydhiet/mandrill-scrapper/internal/repository/lawsuit"
	repoManufacture "github.com/noydhiet/mandrill-scrapper/internal/repository/manufacture"
	repoPatent "github.com/noydhiet/mandrill-scrapper/internal/repository/patent"
	repoRecall "github.com/noydhiet/mandrill-scrapper/internal/repository/recall"
	repoRegistration "github.com/noydhiet/mandrill-scrapper/internal/repository/registration"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/go-co-op/gocron/v2"
)

// searchCmd represents the search command
var WorkerCmd = &cobra.Command{
	Use:   "worker",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: runWorkerScraping,
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// WorkerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// WorkerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runWorkerScraping(cmd *cobra.Command, args []string) {
	log.Info().Msg("worker initiated")
	collector := colly.NewCollector()
	storage, err := pkgStorage.NewStorageDB("")
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

	// create a scheduler
	s, err := gocron.NewScheduler()
	if err != nil {
		log.Error().Err(err).Msg("failed to create scheduler")
		return
	}

	// every day at 00:00
	jPatent, err := s.NewJob(
		gocron.CronJob("0 0 * * *", false),
		gocron.NewTask(hdl.RunWorkerPatent),
	)
	if err != nil {
		log.Error().Err(err).Msg("failed to create job")
		return
	}

	log.Info().Msgf("worker command called patent %v", jPatent.ID())

	log.Info().Msg("worker started")
	s.Start()

	// block the main thread
	select {}

}
