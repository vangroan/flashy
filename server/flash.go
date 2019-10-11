package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// **********************************
// Payloads

// CreateFlashCardRequest is the request payload for creating a new flash card.
type CreateFlashCardRequest struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

// **********************************
// Models

// Box is a collection of flash cards
type Box struct {
	gorm.Model
	Cards []BoxCard
}

// BoxCard is the junction between boxes and cards.
type BoxCard struct {
	ID   uint `gorm:"primary_key"`
	Box  Box
	Card FlashCard
}

// A FlashCard contains markup content
// on two sides.
type FlashCard struct {
	gorm.Model
	Question string
	Answer   string
}

// **********************************
// Repository

// FlashRepository is a repository that offers CRUD operations
// for flash cards and boxes.
type FlashRepository interface {
	CreateFlashCard(card *FlashCard) (uint, error)
}

// FlashDB is an implementation of `FlashRepository` that is
// backed by a database.
type FlashDB struct {
	db *gorm.DB
}

// CreateFlashCard creates a new flash card in the backing storage.
//
// Returns error if the given flashcard is `nil`.
func (r *FlashDB) CreateFlashCard(card *FlashCard) (uint, error) {
	if card == nil {
		return 0, fmt.Errorf("Argument `card` is nil")
	}

	if !r.db.NewRecord(card) {
		return 0, fmt.Errorf("Argument `card` is not a new record")
	}

	r.db.Create(card)

	return card.ID, nil
}

// **********************************
// Handlers

// FlashCardCtrl is the FlashCard CRUD controller.
type FlashCardCtrl struct {
	repo FlashRepository
}

// NewFlashCardCtrl makes a new `FlashCardCtrl`.
func NewFlashCardCtrl(route *mux.Router, db *gorm.DB) FlashCardCtrl {
	ctrl := FlashCardCtrl{repo: &FlashDB{db: db}}
	RegisterController(&ctrl, route)
	return ctrl
}

// ReadMany retrieves all of the current user's flash cards.
func (ctrl *FlashCardCtrl) ReadMany(w http.ResponseWriter, r *http.Request) {
	logger := r.Context().Value(loggerKey).(*log.Entry)
	logger.Debug("ReadMany(): Hello, World!")
	fmt.Fprint(w, "Hello, World!")
}

// Create creates a new record in the backing store.
func (ctrl *FlashCardCtrl) Create(w http.ResponseWriter, r *http.Request) {
	logger := r.Context().Value(loggerKey).(*log.Entry)
	logger.Debug("Create(): Creating new card")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		err = WriteJSONError(w, logger, http.StatusInternalServerError, err)
		logger.Errorf("Create(): Failed to respond with error: %s", err)
		return
	}

	var payload CreateFlashCardRequest
	err = json.Unmarshal(data, &payload)
	if err != nil {
		err = WriteJSONError(w, logger, http.StatusInternalServerError, err)
		logger.Errorf("Create(): Failed to deserialize payload with error: %s", err)
		return
	}

	id, err := ctrl.repo.CreateFlashCard(&FlashCard{
		Question: payload.Question,
		Answer:   payload.Answer,
	})
	if err != nil {
		err = WriteJSONError(w, logger, http.StatusInternalServerError, err)
		logger.Errorf("Create(): Failed to store new flash card: %s", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Location", strconv.FormatUint(uint64(id), 10))
}
