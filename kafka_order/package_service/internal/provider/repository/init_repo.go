package repository

import (
	"package_service/internal/repository"
	repoImplement "package_service/internal/repository/impl"
)

var (
	PackageRepository repository.PackageRepositoryInterface
)

func InitRepository() {
	PackageRepository = repoImplement.NewPackageRepository()
}
