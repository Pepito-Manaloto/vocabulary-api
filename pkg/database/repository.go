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

    rows, err := repository.Db.Query("SELECT v.english_word, v.foreign_word, f.language FROM vocabulary v, foreign_language f WHERE v.foreign_id = f.id order by f.language ASC, v.english_word ASC")
    if err != nil {
        err = errors.Wrapf(err, "getVocabularies. Error in SELECT query. lastUpdated=%s", lastUpdated)
        return nil, err
    }
    defer rows.Close()

    languages, err := processSelectVocabularies(rows)
    if err != nil {
        err = errors.Wrapf(err, "getVocabularies. Error processing database result. lastUpdated=%s", lastUpdated)
        return nil, err
    }

    vocabularies := &j.Vocabularies{}
    vocabularies.Languages = languages

    var recentlyAddedCount string
    err = repository.Db.QueryRow("SELECT COUNT(*) FROM vocabulary WHERE last_updated > (?  - INTERVAL 30 MINUTE)", lastUpdated).Scan(&recentlyAddedCount)
    if err != nil {
        err = errors.Wrapf(err, "getVocabularies. Error getting recently added count. lastUpdated=%s", lastUpdated)
        return nil, err
    }

    vocabularies.RecentlyAddedCount = string(recentlyAddedCount)

    repository.Logger.Info().Msgf("getVocabularies. Done. lastUpdated=%s", lastUpdated)

    return vocabularies, err
}

func processSelectVocabularies(rows *sql.Rows) (*j.Languages, error) {

    languages := &j.Languages{}

    var hokkienVocabularies []j.Vocabulary
    var japaneseVocabularies []j.Vocabulary
    var mandarinVocabularies []j.Vocabulary

    for rows.Next() {
        var englishWord string
        var foreignWord string
        var language string

        err := rows.Scan(&englishWord, &foreignWord, &language)
        if err != nil {
            err = errors.Wrapf(err, "processSelectVocabularies. Error scanning row.")
            return nil, err
        } else {
            vocabulary := j.Vocabulary{englishWord, foreignWord}

            switch language {
                case "Hokkien":
                    hokkienVocabularies = append(hokkienVocabularies, vocabulary)
                case "Japanese":
                    japaneseVocabularies = append(japaneseVocabularies, vocabulary)
                case "Mandarin":
                    mandarinVocabularies = append(mandarinVocabularies, vocabulary)
            }

        }
    }

    languages.Hokkien = &hokkienVocabularies
    languages.Japanese = &japaneseVocabularies
    languages.Mandarin = &mandarinVocabularies

    return languages, nil
}
