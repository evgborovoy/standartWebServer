package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/evgborovoy/StandardWebServer/internal/app/middleware"
	"github.com/evgborovoy/StandardWebServer/internal/app/models"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gorilla/mux"
)

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func InitHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func (api *API) GetAllArticles(w http.ResponseWriter, r *http.Request) {
	InitHeaders(w)
	api.logger.Info("Get all articles GET api/v1/articles")
	articles, err := api.storage.Article().SelectAll()
	if err != nil {
		api.logger.Info("Error while Articles.SelectAll: ", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some trouble to accessing to database. Please,try again later",
			IsError:    true,
		}
		w.WriteHeader(501)
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(articles)
}

func (api *API) GetArticleById(w http.ResponseWriter, r *http.Request) {
	InitHeaders(w)
	api.logger.Info("Get article by id GET /api/v1/articles/{id}")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		api.logger.Info("Trooubles while parsing {id} param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Unsupport id value. Use ID casting to int value",
			IsError:    true,
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}
	article, ok, err := api.storage.Article().FindByID(id)
	if err != nil {
		api.logger.Info("Trouble with connect to DB", err)
		msg := Message{
			StatusCode: 500,
			Message:    "Trouble on our side. Please, try again",
			IsError:    true,
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("Article with {id} not exist")
		msg := Message{
			StatusCode: 404,
			Message:    "Arrticle not exist, try another id",
			IsError:    false,
		}
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(article)
}

func (api *API) DeleteArticleById(w http.ResponseWriter, r *http.Request) {
	InitHeaders(w)
	api.logger.Info("Delete article by Id DELETE /api/v1/articles/{id}")
	api.logger.Info("Get article by id GET /api/v1/articles/{id}")
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		api.logger.Info("Trooubles while parsing {id} param:", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Unsupport id value. Use ID casting to int value",
			IsError:    true,
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}
	_, ok, err := api.storage.Article().FindByID(id)
	if err != nil {
		api.logger.Info("Trouble with connect to DB", err)
		msg := Message{
			StatusCode: 500,
			Message:    "Trouble on our side. Please, try again",
			IsError:    true,
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("Can not find article with that id")
		msg := Message{
			StatusCode: 404,
			Message:    "Arrticle not exist, try another id",
			IsError:    false,
		}
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(msg)
		return
	}
	_, err = api.storage.Article().DeleteByID(id)
	if err != nil {
		api.logger.Info("Trouble with deleting article", err)
		msg := Message{
			StatusCode: 501,
			Message:    "Trouble on our side. Please, try again",
			IsError:    true,
		}
		w.WriteHeader(501)
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(202)
	msg := Message{
		StatusCode: 202,
		Message:    "successfuly deleted",
		IsError:    false,
	}
	json.NewEncoder(w).Encode(msg)
}

func (api *API) PostArticle(w http.ResponseWriter, r *http.Request) {
	InitHeaders(w)
	api.logger.Info("Post article POST /api/v1/articles")
	var article models.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		api.logger.Info("Invalid info recived from client", err)
		msg := Message{
			StatusCode: 400,
			Message:    "JSON is invalid",
			IsError:    true,
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}
	a, err := api.storage.Article().Create(&article)
	if err != nil {
		api.logger.Info("Something wrong while creating new article", err)
		msg := Message{
			StatusCode: 501,
			Message:    "We have some problem on our side. Please, try later",
			IsError:    true,
		}
		w.WriteHeader(501)
		json.NewEncoder(w).Encode(msg)
		return
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(a)
}

func (api *API) PostUserRegister(w http.ResponseWriter, r *http.Request) {
	InitHeaders(w)
	api.logger.Info("Post user register POST /api/v1/users/register")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		api.logger.Info("Invalid json from client", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Invalid data",
			IsError:    true,
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}
	_, ok, err := api.storage.User().FindByLogin(user.Login)
	if err != nil {
		api.logger.Info("Trouble with connect to DB", err)
		msg := Message{
			StatusCode: 500,
			Message:    "Trouble on our side. Please, try again",
			IsError:    true,
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(msg)
		return
	}
	if ok {
		api.logger.Info("User with that login already exist", user.Login)
		msg := Message{
			StatusCode: 400,
			Message:    "User with that login already exist",
			IsError:    true,
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}
	userAdded, err := api.storage.User().Create(&user)
	if err != nil {
		api.logger.Info("Trouble with connect to DB while crreating user", err)
		msg := Message{
			StatusCode: 500,
			Message:    "Trouble on our side. Please, try again",
			IsError:    true,
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(msg)
		return
	}
	msg := Message{
		StatusCode: 201,
		Message:    fmt.Sprintf("Successfuly created {login:%s}", userAdded.Login),
		IsError:    false,
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(msg)
}

func (api *API) PostUserAuth(w http.ResponseWriter, r *http.Request) {
	InitHeaders(w)
	api.logger.Info("Post to auth POST /api/v1/users/auth")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		api.logger.Info("Invalid json from client", err)
		msg := Message{
			StatusCode: 400,
			Message:    "Invalid data",
			IsError:    true,
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}
	userInDB, ok, err := api.storage.User().FindByLogin(user.Login)
	if err != nil {
		api.logger.Info("Trouble with connect to DB", err)
		msg := Message{
			StatusCode: 500,
			Message:    "Trouble on our side. Please, try again",
			IsError:    true,
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(msg)
		return
	}
	if !ok {
		api.logger.Info("User with that login does not exist", user.Login)
		msg := Message{
			StatusCode: 400,
			Message:    "User with that login does not exist. Try register first",
			IsError:    true,
		}
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(msg)
		return
	}
	if userInDB.Password != user.Password {
		api.logger.Info("Invalid credetials to auth")
		msg := Message{
			StatusCode: 404,
			Message:    "Invalid password to auth",
			IsError:    true,
		}
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(msg)
		return
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims) // Дополнительнык действия (в формате мапы) для шифрования
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()
	claims["admin"] = true
	claims["name"] = userInDB.Login
	tokenString, err := token.SignedString(middleware.SecretKey)
	if err != nil {
		api.logger.Info("Can not claimed jwt-token")
		msg := Message{
			StatusCode: 500,
			Message:    "We have some troubles :(. Try later",
			IsError:    true,
		}
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(msg)
		return
	}
	msg := Message{
		StatusCode: 201,
		Message:    tokenString,
		IsError:    false,
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(msg)
}
