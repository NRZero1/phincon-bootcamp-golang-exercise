package handler

import (
	"package_service/internal/handler"
	handlerImpl "package_service/internal/handler/impl"
	providerUseCase "package_service/internal/provider/usecase"
)

var (
	PackageHandler handler.PackageHandlerInterface
)

func InitHandler() {
	PackageHandler = handlerImpl.NewPackageHandler(providerUseCase.PackageUseCase)
}
