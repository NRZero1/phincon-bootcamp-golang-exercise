package usecase

import (
	providerRepo "package_service/internal/provider/repository"
	"package_service/internal/usecase"
	useCaseImpl "package_service/internal/usecase/impl"
)

var (
	PackageUseCase usecase.PackageUseCaseInterface
)

func InitUseCase() {
	PackageUseCase = useCaseImpl.NewPackageUseCase(providerRepo.PackageRepository)
}
