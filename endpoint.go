package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Endpoint struct {
	s *Storage
}

func (e *Endpoint) getUser(ctx *gin.Context) {
	name := ctx.Param("name")
	history, err := e.s.getHistory(name)
	if err != nil {
		ctx.IndentedJSON(
			http.StatusInternalServerError,
			"Could not load the history from the database: "+err.Error())
		return
	}
	ctx.IndentedJSON(http.StatusOK, history)
}

func (e *Endpoint) postUser(ctx *gin.Context) {
	name := ctx.Param("name")
	snapshot, err := takeSnapshot(name)
	if err != nil {
		ctx.IndentedJSON(
			http.StatusInternalServerError,
			gin.H{"message": "taking a snapshot failed"})
		return
	}
	if err := e.s.addSnapshot(snapshot); err != nil {
		ctx.IndentedJSON(
			http.StatusInternalServerError,
			"Could not register a snapshot: "+err.Error())
		return
	}
	ctx.IndentedJSON(http.StatusCreated, snapshot)
}
