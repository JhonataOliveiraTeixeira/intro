package api

import (
	"encoding/json"
	application "intro/Application"
	"intro/db"
	"intro/domain"
	"intro/utils"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler() http.Handler{
	r:= chi.NewMux()
	db := db.Init()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Post("/user", handleCreate(*db))
	r.Put("/user/{id}", handleUpdate(*db))
	r.Get("/user/{id}", handleFindUnique(*db))
	r.Delete("/user/{id}", handleDelete(*db))
	r.Get("/users/{page}", handleFindAll(*db))
	return  r
}


func handleCreate(db db.DB) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request)  {
		var body domain.User
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil{
			slog.Error("Invalid body type", "recieved", r.Body)
			utils.SendJSON(w, utils.Response{Error: "Invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		err := application.Create(body.Name,body.Email,body.Password, db)
		if err != nil{
			slog.Error("Error in create user", "error", err)
			utils.SendJSON(w, utils.Response{Error: err.Error()}, http.StatusBadRequest)
			return 
		}
		utils.SendJSON(w, utils.Response{Data: body}, http.StatusCreated)
	}

}

func handleUpdate(db db.DB) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request)  {
		var body domain.User
		id := chi.URLParam(r, "id")
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil{
			slog.Error("Invalid body type", "recieved", r.Body)
			utils.SendJSON(w, utils.Response{Error: "Invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil{
			slog.Error("parse id to int failed", "recieved", id)
			utils.SendJSON(w, utils.Response{Error: "Invalid id"}, http.StatusBadRequest)
		}

		err = application.Update(int(idInt), body.Name,body.Email, db)
		if err != nil{
			slog.Error("error in update user", "id", id,"error", err)
			utils.SendJSON(w, utils.Response{Error: err.Error()}, http.StatusBadRequest)
			return 
		}
		slog.Info("Updating user", "id", id)

		utils.SendJSON(w, utils.Response{Data: body}, http.StatusCreated)
	}

}

func handleFindUnique(db db.DB) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request)  {
		id := chi.URLParam(r, "id")

		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil{
			slog.Error("parse id to int failed", "recieved", id)
			utils.SendJSON(w, utils.Response{Error: "Invalid id"}, http.StatusBadRequest)
			return 
		}

		user, err := application.FindUnique(int(idInt), db)
		if err != nil{
			slog.Error("error in find unique", "error", err)
			utils.SendJSON(w, utils.Response{Error: err.Error()}, http.StatusBadRequest)
			return 
		}

		slog.Info("Return to user", "id", id, "user", user)
		utils.SendJSON(w, utils.Response{Data: user}, http.StatusCreated)
	}

}

func handleFindAll(db db.DB) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request)  {
		id := chi.URLParam(r, "page")

		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil{
			slog.Error("parse page to int failed", "recieved", id)
			utils.SendJSON(w, utils.Response{Error: "Invalid page "+err.Error()}, http.StatusBadRequest)
			return 
		}

		user := application.FindAll(int(idInt), db)
		slog.Info("Return to users", "take", idInt)
		utils.SendJSON(w, utils.Response{Data: user}, http.StatusCreated)
	}

}

func handleDelete(db db.DB) http.HandlerFunc{
	return func (w http.ResponseWriter, r *http.Request)  {
		id := chi.URLParam(r, "id")

		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil{
			slog.Error("parse id to int failed", "recieved", id)
			utils.SendJSON(w, utils.Response{Error: "Invalid id"}, http.StatusBadRequest)
			return 
		}

		err = application.Delete(int(idInt), db)
		if err != nil{
			slog.Error("Error in delete user", "error", err)
		}
		utils.SendJSON(w, utils.Response{}, http.StatusOK)
	}

}