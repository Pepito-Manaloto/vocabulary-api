package main

import (
    "encoding/json"
    "github.com/Pepito-Manaloto/vocabulary-api/pkg/database"
    "github.com/Pepito-Manaloto/vocabulary-api/pkg/file"
    "time"
)

const (
    configFilePath = "conf/vocabulary.json"
)

func main() {

    configuration := file.LoadConfiguration(configFilePath)
    logger := file.NewLogger(configuration)

    logger.Info().Msg("Started!")

    db, err := database.Connect(&configuration)
    if err != nil {
        logger.Error().Err(err).Msg("Failed getting database connection")
    }

    repository := database.Repository{
        Db: db,
        Logger: &logger,
    }

    vocabularies, err := repository.GetVocabularies(time.Now().AddDate(-10, 0,0))

    result, _ := json.Marshal(vocabularies)
    logger.Info().Msg(string(result))
}
