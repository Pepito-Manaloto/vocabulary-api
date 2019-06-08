package resource

import (
    "encoding/json"
    "github.com/Pepito-Manaloto/vocabulary-api/pkg/database"
    j "github.com/Pepito-Manaloto/vocabulary-api/pkg/json"
    "github.com/julienschmidt/httprouter"
    "github.com/rs/zerolog"
    "net/http"
    "time"
)

const (
    DefaultLastUpdatedDate = "1950-01-01"
)

type Handler struct {
    Repository *database.Repository
    Logger     *zerolog.Logger
}

func (handler *Handler) Handle() http.Handler {
    router := httprouter.New()

    router.GET("/Vocabulary/vocabularies", getVocabularies(handler))

    return router
}

func getVocabularies(h *Handler) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

        lastUpdatedQueryParam := r.URL.Query().Get("last_updated")

        lastUpdated, err := parseLastUpdatedQueryParam(lastUpdatedQueryParam)
        if err != nil {
            h.Logger.Err(err).Msgf("getVocabularies. Failed parsing lastUpdated query parameter. lastUpdatedQueryParam=%s lastUpdated=%s",
                lastUpdatedQueryParam, lastUpdated)

            jsonErrorResponse(w, err, http.StatusBadRequest)
            return
        }

        h.Logger.Info().Msgf("getVocabularies. Retrieving vocabularies from database. lastUpdated=%s", lastUpdated)
        vocabularies, err := h.Repository.GetVocabularies(lastUpdated)
        if err != nil {
            h.Logger.Err(err).Msgf("getVocabularies. Failed retrieving from database. lastUpdated=%s", lastUpdated)

            jsonErrorResponse(w, err, http.StatusInternalServerError)
            return
        }

        err = jsonSuccessResponse(w, vocabularies)
        if err != nil {
            h.Logger.Err(err).Msgf("getVocabularies. Failed encoding json to writer. lastUpdated=%s", lastUpdated)

            jsonErrorResponse(w, err, http.StatusInternalServerError)
            return
        }
    }
}

func parseLastUpdatedQueryParam(lastUpdatedQuery string) (time.Time, error) {

    if len(lastUpdatedQuery) <= 0 {
        lastUpdatedQuery = DefaultLastUpdatedDate
    }

    return time.Parse("2006-01-02", lastUpdatedQuery)
}

func jsonSuccessResponse(w http.ResponseWriter, response interface{}) error {
    return json.NewEncoder(w).Encode(response)
}

func jsonErrorResponse(w http.ResponseWriter, message error, code uint) {
    jsonErr :=  j.Error {
        Message: message.Error(),
        Code: code,
    }

    w.Header().Set("Content-Type", "application/json")
    _ = json.NewEncoder(w).Encode(jsonErr)
}