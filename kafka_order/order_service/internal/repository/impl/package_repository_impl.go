package impl

import (
	"fmt"
	"package_service/internal/domain"
	"package_service/internal/repository"
	"package_service/internal/utils"
	"sync"

	"github.com/rs/zerolog/log"
)

type PackageRepository struct {
	mtx sync.Mutex
	packages map[int]domain.Package
}

func NewPackageRepository() repository.PackageRepositoryInterface {
	repo := &PackageRepository {
		packages: map[int]domain.Package{},
	}
	repo.initData()
	return repo
}

func (repo *PackageRepository) initData() {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()

	repo.packages[1] = domain.Package{
		PackageID: 1,
		Name: "Didi Ha",
		Price: 100,
	}

	repo.packages[2] = domain.Package{
		PackageID: 2,
		Name: "Bene Ha",
		Price: 200,
	}
}

func (repo *PackageRepository) FindById(id int) (domain.Package, error) {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()

	log.Trace().Msg("Inside package repository find by id")
	log.Trace().Msg("Attempting to fetch package")
	if foundPackage, exists := repo.packages[id]; exists {
		log.Trace().Msg("Fetching completed")
		return foundPackage, nil
	}
	log.Error().Msg(fmt.Sprintf("Package with ID %d not found", id))
	return domain.Package{}, utils.NewErrFindById(id)
}

func (repo *PackageRepository) GetAll() ([]domain.Package) {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()

	log.Trace().Msg("Inside user repository get all")
	log.Trace().Msg("Attempting to fetch packages")
	listOfPackages := make([]domain.Package, 0, len(repo.packages))

	for _, v := range repo.packages {
		temp := domain.Package {
			PackageID: v.PackageID,
			Name: v.Name,
			Price: v.Price,
		}
		listOfPackages = append(listOfPackages, temp)
	}

	log.Trace().Msg("Fetching completed")
	return listOfPackages
}

func (repo *PackageRepository) FindByName(name string) (domain.Package, error) {
	repo.mtx.Lock()
	defer repo.mtx.Unlock()

	log.Trace().Msg("Inside user repository find by username")
	for _, v := range repo.packages {
		if v.Name == name {
			temp := domain.Package {
				PackageID: v.PackageID,
				Name: v.Name,
				Price: v.Price,
			}
			return temp, nil
		}
	}
	return domain.Package{}, utils.NewErrFindByName(name)
}
