package impl

import (
	"package_service/internal/domain"
	"package_service/internal/repository"
	"package_service/internal/usecase"

	"github.com/rs/zerolog/log"
)

type PackageUseCase struct {
	repo repository.PackageRepositoryInterface
}

func NewPackageUseCase(repo repository.PackageRepositoryInterface) (usecase.PackageUseCaseInterface) {
	return PackageUseCase{
		repo: repo,
	}
}

func (uc PackageUseCase) FindById(id int) (domain.Package, error) {
	log.Trace().Msg("Entering package usecase find by id")
	return uc.repo.FindById(id)
}

func (uc PackageUseCase) GetAll() ([]domain.Package) {
	log.Trace().Msg("Entering package usecase get all")
	return uc.repo.GetAll()
}

func (uc PackageUseCase) FindByName(name string) (domain.Package, error) {
	log.Trace().Msg("Enter package usecase FindByName")
	return uc.repo.FindByName(name)
}
