package cmd

import (
	"movielibrary/api"
	"movielibrary/client"
	"movielibrary/config"
	"movielibrary/internal/controller"
	"movielibrary/internal/repository"
	"movielibrary/internal/routes"
	"movielibrary/internal/service"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use: "serve",
	Run: serve,
}

func serve(cmd *cobra.Command, args []string) {
	conf := config.Get()

	db := client.ConnectMySQL(conf.Db)
	omdbc := client.NewOmdbClient(conf.Omdb)

	libraryRepo := repository.NewMovieLibraryRepository(db)
	librarySvc := service.NewMovieLibraryService(libraryRepo, omdbc)
	libraryController := controller.NewMovieLibraryController(librarySvc)
	libraryRouter := routes.NewLibraryRouter(libraryController)

	server := api.NewServer(libraryRouter)
	server.Start()
}
