package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/regmi-bpn/movies-api/internal/controller"
)

func InitRoutes(movie controller.Movie) *gin.Engine {
	r := gin.Default()
	movieGroup := r.Group("/v1/movies")
	{
		movieGroup.GET("", movie.GetMovies)
		movieGroup.GET("/:id", movie.GetMovie)
		movieGroup.POST("", movie.SaveMovie)
		movieGroup.PUT("/:id", movie.UpdateMovie)
		movieGroup.DELETE("/:id", movie.DeleteMovie)
	}
	return r
}
