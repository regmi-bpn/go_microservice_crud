package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/regmi-bpn/movies-api/internal/service"
	"github.com/regmi-bpn/movies-api/types"
)

type Movie struct {
	service service.MovieService
}

func NewMovieController(service service.MovieService) Movie {
	return Movie{service: service}
}

func (c Movie) SaveMovie(ctx *gin.Context) {
	var request types.MovieRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := c.service.SaveMovie(request)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(201, res)
}

func (c Movie) UpdateMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(400, gin.H{
			"error": "id is required!!",
		})
		return
	}
	var request types.MovieRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := c.service.UpdateMovie(id, request)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, res)
}

func (c Movie) DeleteMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(400, gin.H{
			"error": "id is required!!",
		})
		return
	}

	if err := c.service.DeleteMovie(id); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.Status(200)
}

func (c Movie) GetMovie(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(400, gin.H{
			"error": "id is required!!",
		})
		return
	}
	res, err := c.service.GetMovie(id)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, res)

}

func (c Movie) GetMovies(ctx *gin.Context) {

	res, err := c.service.GetMovies()
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, res)

}
