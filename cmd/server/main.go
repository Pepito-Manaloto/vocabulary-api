package main

import (
    "github.com/Pepito-Manaloto/vocabulary-api/pkg/database"
    "github.com/Pepito-Manaloto/vocabulary-api/pkg/file"
    "github.com/Pepito-Manaloto/vocabulary-api/pkg/resource"
    "net/http"
)

const (
    ConfigFilePath = "conf/vocabulary.json"
)

func main() {

    configuration := file.LoadConfiguration(ConfigFilePath)
    logger := file.NewLogger(configuration)

    logger.Info().Msg("Started!")

    db, err := database.Connect(&configuration)
    if err != nil {
        logger.Error().Err(err).Msg("main. Failed getting database connection")
    }

    repository := database.Repository{
        Db:     db,
        Logger: &logger,
    }

    routerHandler := resource.Handler {
        Repository: &repository,
        Logger: &logger,
    }

    err = http.ListenAndServe(":8888", routerHandler.Handle())

    if err != nil {
        logger.Fatal().Err(err).Msg("main. Failed http.ListenAndServe")
    } else {
        logger.Info().Msg("main. Started!")
    }
}
