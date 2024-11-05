package person

import (
	"github.com/felipeversiane/go-otel/internal/infra/services/database"
	"github.com/gin-gonic/gin"
)

func PersonRouter(g *gin.RouterGroup, db database.DatabaseInterface) *gin.RouterGroup {
	controller := NewPersonController(NewPersonService(NewPersonRepository(db)))

	person := g.Group("/person")
	{
		person.POST("/", controller.InsertOneController)

	}

	return person
}
