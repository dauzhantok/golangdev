package http

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	lru "github.com/hashicorp/golang-lru"
	"gomidka/internal/message_broker"
	"gomidka/internal/models"
	"gomidka/internal/store"
	"net/http"
	"strconv"
)

type BreadResource struct {
	store store.Store
	broker message_broker.MessageBroker
	cache *lru.TwoQueueCache
}

func NewBreadResource(store store.Store, broker message_broker.MessageBroker, cache *lru.TwoQueueCache) *BreadResource {
	return &BreadResource{
		store: store,
		broker: broker,
		cache: cache,
	}
}

func (cr *BreadResource) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", cr.CreateBread)
	r.Get("/", cr.AllBread)
	r.Get("/{id}", cr.ByID)
	r.Put("/", cr.UpdateBread)
	r.Delete("/{id}", cr.DeleteBread)

	return r
}

func (cr *BreadResource) CreateBread(w http.ResponseWriter, r *http.Request) {
	bread := new(models.Bread)
	if err := json.NewDecoder(r.Body).Decode(bread); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	if err := cr.store.Breads().Create(r.Context(), bread); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}

	// Правильно пройтись по всем буквам и всем словам
	cr.cache.Purge() // в рамках учебного проекта полностью чистим кэш после создания новой категории

	w.WriteHeader(http.StatusCreated)
}

func (cr *BreadResource) AllBread(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	filter := &models.BreadsFilter{}

	searchQuery := queryValues.Get("query")
	if searchQuery != "" {
		breadsFromCache, ok := cr.cache.Get(searchQuery)
		if ok {
			render.JSON(w, r, breadsFromCache)
			return
		}

		filter.Query = &searchQuery
	}

	breads, err := cr.store.Breads().All(r.Context(), filter)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}

	if searchQuery != "" {
		cr.cache.Add(searchQuery, breads)
	}
	render.JSON(w, r, breads)
}

func (cr *BreadResource) ByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	breadFromCache, ok := cr.cache.Get(id)
	if ok {
		render.JSON(w, r, breadFromCache)
		return
	}

	bread, err := cr.store.Breads().ByID(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}

	cr.cache.Add(id, bread)
	render.JSON(w, r, bread)
}

func (cr *BreadResource) UpdateBread(w http.ResponseWriter, r *http.Request) {
	bread := new(models.Bread)
	if err := json.NewDecoder(r.Body).Decode(bread); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	err := validation.ValidateStruct(
		bread,
		validation.Field(&bread.ID, validation.Required),
		validation.Field(&bread.Name, validation.Required),
	)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	if err := cr.store.Breads().Update(r.Context(), bread); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}

	cr.broker.Cache().Remove(bread.ID)
}

func (cr *BreadResource) DeleteBread(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	if err := cr.store.Breads().Delete(r.Context(), id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}

	cr.broker.Cache().Remove(id)
}