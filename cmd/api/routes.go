package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/suggestions", app.getSuggestionsHandler)
	router.HandlerFunc(http.MethodPost, "/suggestions", app.createSuggestionHandler)
	router.HandlerFunc(http.MethodDelete, "/suggestions/:id", app.deleteSuggestionHandler)
	router.HandlerFunc(http.MethodPatch, "/suggestions/:id", app.updateSuggestionHandler)

	/* router.HandlerFunc(http.MethodGet, "/comments", app.getCommentsHandler)
	router.HandlerFunc(http.MethodPost, "/comments", app.createCommentHandler)
	router.HandlerFunc(http.MethodDelete, "/comments/:id", app.deleteCommentHandler)
	router.HandlerFunc(http.MethodPatch, "/comments/:id", app.updateCommentHandler)

	router.HandlerFunc(http.MethodGet, "/replies", app.getRepliesHandler)
	router.HandlerFunc(http.MethodPost, "/replies", app.createReplyHandler)
	router.HandlerFunc(http.MethodDelete, "/replies/:id", app.deleteReplyHandler)
	router.HandlerFunc(http.MethodPatch, "/replies/:id", app.updateReplyHandler)

	router.HandlerFunc(http.MethodGet, "/users", app.getUsersHandler)
	router.HandlerFunc(http.MethodPost, "/users", app.createUserHandler)
	router.HandlerFunc(http.MethodDelete, "/users/:id", app.deleteUserHandler)
	router.HandlerFunc(http.MethodPatch, "/users/:id", app.updateUserHandler)
	*/
	return router
}
