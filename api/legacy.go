package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
	dphandlers "github.com/ONSdigital/dp-net/handlers"
	"github.com/ONSdigital/dp-release-calendar-api/mapper"
	"github.com/ONSdigital/log.go/v2/log"
)

func LegacyHandler(ctx context.Context, zc ZebedeeClient) http.HandlerFunc {
	return dphandlers.ControllerHandler(func(w http.ResponseWriter, r *http.Request, lang, collectionID, accessToken string) {
		handleLegacy(w, r, lang, collectionID, accessToken, zc)
	})
}

func handleLegacy(w http.ResponseWriter, req *http.Request, lang, collectionID, accessToken string, zc ZebedeeClient) {
	ctx := req.Context()

	urlParam := req.URL.Query().Get("url")
	lang = func(secondary string) string {
		if primary := req.URL.Query().Get("lang"); primary != "" {
			return primary
		}
		return secondary
	}(lang)

	if urlParam == "" {
		err := errors.New("url parameter not found")
		log.Error(ctx, "url parameter not found", err)
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	zebedeeRelease, err := zc.GetRelease(ctx, accessToken, collectionID, lang, urlParam)
	if err != nil {
		setStatusCode(ctx, w, "retrieving release from Zebedee", err)
		return
	}

	response := mapper.CreateReleaseFromZebedee(zebedeeRelease)

	w.Header().Set("Content-Type", "application/json")
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		setStatusCode(ctx, w, "marshalling response failed", err)
		return
	}

	_, err = w.Write(jsonResponse)
	if err != nil {
		setStatusCode(ctx, w, "writing response failed", err)
		return
	}
}

func setStatusCode(ctx context.Context, w http.ResponseWriter, msg string, err error) {
	statusCode := http.StatusInternalServerError
	var e zebedee.ErrInvalidZebedeeResponse
	if errors.As(err, &e) {
		statusCode = e.ActualCode
	}
	log.Error(ctx, msg, err)
	w.WriteHeader(statusCode)
}
