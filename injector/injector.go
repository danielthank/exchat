package injector

import (
	"github.com/danielthank/exchat-server/domain/repository"
	"github.com/danielthank/exchat-server/handler"
	"github.com/danielthank/exchat-server/infra"
	"github.com/danielthank/exchat-server/usecase"
)

func InjectDB() *infra.SqlHandler {
	sqlhandler := infra.NewSqlHandler()
	return sqlhandler
}

func InjectRedis() *infra.RedisHandler {
	redisHandler := infra.NewRedisHandler()
	return redisHandler
}

func InjectProfileRepository() repository.ProfileRepository {
	sqlHandler := InjectDB()
	return infra.NewProfileRepository(sqlHandler)
}

func InjectWSHandler() *handler.WSHandler {
	redisHandler := InjectRedis()
	return handler.NewWSHandler(redisHandler)
}

func InjectAuthUsecase() usecase.AuthUsecase {
	profileRepository := InjectProfileRepository()
	authUsecase := usecase.NewAuthUsecase(&profileRepository)
	return authUsecase
}

func InjectAuthHandler() *handler.AuthHandler {
	authUsecase := InjectAuthUsecase()
	authHandler := handler.NewAuthHandler(&authUsecase)
	return authHandler
}
