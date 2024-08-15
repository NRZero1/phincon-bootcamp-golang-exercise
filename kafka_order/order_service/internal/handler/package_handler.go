package handler

import "github.com/gin-gonic/gin"

type PackageHandlerInterface interface {
	PackageFindById
	PackageGetAll
	PackageFindByName
	PackageGetAllOrFindByName
}

type PackageFindById interface {
	FindById(context *gin.Context)
}

type PackageGetAll interface {
	GetAll(context *gin.Context)
}

type PackageFindByName interface {
	FindByName(context *gin.Context)
}

type PackageGetAllOrFindByName interface {
	GetAllOrFindByName(context *gin.Context)
}
