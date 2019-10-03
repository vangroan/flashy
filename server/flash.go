package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
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

func (ctrl *FlashCardCtrl) Read(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
