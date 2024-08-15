package repository

import (
	"package_service/internal/domain"
)

type PackageRepositoryInterface interface {
	PackageFindByName
	PackageFindById
	PackageGetAll
}

type PackageFindByName interface {
	FindByName(name string) (domain.Package, error)
}

type PackageFindById interface {
	FindById(id int) (domain.Package, error)
}

type PackageGetAll interface {
	GetAll() ([]domain.Package)
}
