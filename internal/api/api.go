package api

import (
	"encoding/json"
	"location-history-server/internal/model"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (a *App) addLocationHistory(w http.ResponseWriter, r *http.Request) {
	location := new(model.Location)
	if err := json.NewDecoder(r.Body).Decode(location); err != nil {
		log.Println("addLocationHistory: json decoding failed:", err)
		fail(w, http.StatusBadRequest, err.Error())
		return
	}
	orderId := chi.URLParam(r, "order_id")
	val, ok := a.Cache.Get(orderId)
	var locations []model.Location
	if ok {
		locations = val.([]model.Location)
		locations = append(locations, *location)
	} else {
		locations = []model.Location{*location}
	}
	a.Cache.Set(orderId, locations)
	send(w, http.StatusOK, "success")
}

func (a *App) getLocationHistory(w http.ResponseWriter, r *http.Request) {
	orderId := chi.URLParam(r, "order_id")
	maxStr := r.URL.Query().Get("max")
	var out []model.Location
	var data []model.Location

	val, ok := a.Cache.Get(orderId)
	if ok {
		data = val.([]model.Location)
	} else {
		log.Println("getLocationHistory: not found")
		fail(w, http.StatusNotFound, "order id not found")
		return
	}
	if len(maxStr) == 0 {
		out = data
	} else {
		max, err := strconv.Atoi(maxStr)
		if err != nil {
			log.Println("getLocationHistory: invalid max value:", err)
			fail(w, http.StatusBadRequest, err.Error())
			return
		}
		locations := data
		if max > len(locations) {
			max = len(locations)
		}
		out = locations[len(locations)-max:]
	}

	send(w, http.StatusOK, map[string]interface{}{
		"order_id": orderId,
		"history":  out,
	})
}

func (a *App) deleteLocationHistory(w http.ResponseWriter, r *http.Request) {
	orderId := chi.URLParam(r, "order_id")
	a.Cache.Remove(orderId)
	send(w, http.StatusOK, "success")
}
