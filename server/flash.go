package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// **********************************
// Models

// Box is a collection of flash cards
type Box struct {
}

// A FlashCard contains markup content
// on two sides.
type FlashCard struct {
	Question string
	Answer   string
}

// **********************************
// Handlers

// FlashCardCtrl is the FlashCard CRUD controller.
type FlashCardCtrl struct {
	db *gorm.DB
}

// NewFlashCardCtrl makes a new `FlashCardCtrl`.
func NewFlashCardCtrl(route *mux.Router, db *gorm.DB) FlashCardCtrl {
	ctrl := FlashCardCtrl{db: db}
	RegisterController(&ctrl, route)
	return ctrl
}

// ReadMany retrieves all of the current user's flash cards.
func (ctrl *FlashCardCtrl) ReadMany(w http.ResponseWriter, r *http.Request) {
	logger := r.Context().Value(loggerKey).(*log.Entry)
	logger.Debug("ReadMany(): Hello, World!")
	fmt.Fprint(w, "Hello, World!")
}
