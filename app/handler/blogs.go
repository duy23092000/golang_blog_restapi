package handler

import (
	"encoding/json"
	"net/http"
	"tutorial-rest/app/model"

	"github.com/jinzhu/gorm"
)

func GetAllBlogs(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	blogs := []model.Blog{}
	db.Find(&blogs)
	responseJSON(w, http.StatusOK, blogs)
}

func CreateBlogs(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	blog := model.Blog{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&blog); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&blog).Error; err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseJSON(w, http.StatusCreated, blog)
}
