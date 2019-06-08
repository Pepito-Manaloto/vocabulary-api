package database

import (
    "database/sql"
    j "github.com/Pepito-Manaloto/vocabulary-api/pkg/json"
    "github.com/pkg/errors"
    "github.com/rs/zerolog"
    "time"
)

type Repository struct {
    Db     *sql.DB
    Logger *zerolog.Logger
}

func (repository *Repository) GetVocabularies(lastUpdated time.Time) (*j.Vocabularies, error) {

    repository.Logger.Info().Msgf("getVocabularies. Start. lastUpdated=%s", lastUpdated)

    var recentlyAddedCount string
    rows, err := repository.Db.Query("CALL Get_Vocabularies(?, ?)", lastUpdated, sql.Out{Dest: &recentlyAddedCount})

    if err != nil {
        err = errors.Wrapf(err, "getVocabularies. Error calling Get_Vocabularies(%s, @recently_added_count)", lastUpdated)
        return nil, err
    }

    defer rows.Close()

    languages := &j.Languages{}
    vocabularies := &j.Vocabularies{}

    counter := 1
    for {
        var vocabularyArray []j.Vocabulary

        for rows.Next() {
            var englishWord string
            var foreignWord string

            err := rows.Scan(&englishWord, &foreignWord)
            if err != nil {
                err = errors.Wrapf(err, "getVocabularies. Error scanning row.")
                return nil, err
            } else {
                vocabulary := j.Vocabulary{englishWord, foreignWord}
                vocabularyArray = append(vocabularyArray, vocabulary)
            }
        }

        switch counter {
        case 1:
            languages.Hokkien = &vocabularyArray
        case 2:
            languages.Japanese = &vocabularyArray
        case 3:
            languages.Mandarin = &vocabularyArray
        }

        if !rows.NextResultSet() {
            break
        } else {
            counter++
        }
    }

    vocabularies.Languages = languages
    vocabularies.RecentlyAddedCount = recentlyAddedCount

    repository.Logger.Info().Msgf("getVocabularies. Done. lastUpdated=%s", lastUpdated)

    return vocabularies, err
}
