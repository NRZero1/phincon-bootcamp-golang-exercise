package usecase

import "package_service/internal/domain"

type PackageUseCaseInterface interface {
	PackageFindById
	PackageGetAll
	PackageFindByName
}

type PackageFindById interface {
	FindById(id int) (domain.Package, error)
}

type PackageGetAll interface {
	GetAll() ([]domain.Package)
}

type PackageFindByName interface {
	FindByName(name string) (domain.Package, error)
}
