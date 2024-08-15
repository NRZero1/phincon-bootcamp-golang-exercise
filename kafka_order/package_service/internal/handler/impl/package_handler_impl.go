package impl

import (
	"net/http"
	"package_service/internal/domain/dto/response"
	"package_service/internal/handler"
	"package_service/internal/usecase"
	"package_service/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type PackageHandler struct {
	usecase usecase.PackageUseCaseInterface
}

func NewPackageHandler(usecase usecase.PackageUseCaseInterface) (handler.PackageHandlerInterface) {
	return PackageHandler {
		usecase: usecase,
	}
}

func (h PackageHandler) FindById(c *gin.Context) {
	log.Trace().Msg("Entering package handler find by id")

	idString := c.Param("id")
	log.Debug().Str("Received Id is: ", idString)

	log.Trace().Msg("Trying to convert id in string to int")
	id, errConv := strconv.Atoi(idString)

	if errConv != nil {
		log.Trace().Msg("Error happens when converting id string to int")
		log.Error().Str("Error message: ", errConv.Error())
		response := response.GlobalResponse {
			Message: utils.ErrPathVar.Error(),
			StatusCode: http.StatusBadRequest,
			Data: nil,
		}

		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")
        c.Writer.WriteHeader(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	foundPackage, errFound := h.usecase.FindById(id)

	if errFound != nil {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")

		var resp response.GlobalResponse

		log.Trace().Msg("Fetch error")
		log.Error().Str("Error message: ", errFound.Error())
		c.Writer.WriteHeader(http.StatusBadRequest)

		resp = response.GlobalResponse {
			Message: errFound.Error(),
			StatusCode: http.StatusBadRequest,
			Data: "",
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	response := response.GlobalResponse {
		Message: "OK",
		StatusCode: http.StatusOK,
		Data: foundPackage,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	log.Info().Msg("Package fetched and returning json")
	c.JSON(http.StatusOK, response)
}

func (h PackageHandler) GetAll(c *gin.Context) {
	log.Trace().Msg("Entering package get all handler")

	allPackages := h.usecase.GetAll()

	response := response.GlobalResponse {
		Message: "OK",
		StatusCode: http.StatusOK,
		Data: allPackages,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	log.Info().Msg("Package fetched and returning json")
	c.JSON(http.StatusOK, response)
}

func (h PackageHandler) FindByName(c *gin.Context) {
	log.Trace().Msg("Entering package handler find by name")

	name := c.Query("name")
	log.Debug().Str("Received name is: ", name)

	foundPackage, errFound := h.usecase.FindByName(name)

	if errFound != nil {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("X-Content-Type-Options", "nosniff")

		var resp response.GlobalResponse

		log.Trace().Msg("Fetch error")
		log.Error().Str("Error message: ", errFound.Error())
		c.Writer.WriteHeader(http.StatusBadRequest)

		resp = response.GlobalResponse {
			Message: errFound.Error(),
			StatusCode: http.StatusBadRequest,
			Data: "",
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	response := response.GlobalResponse {
		Message: "OK",
		StatusCode: http.StatusOK,
		Data: foundPackage,
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	log.Info().Msg("Package fetched and returning json")
	c.JSON(http.StatusOK, response)
}

func (h PackageHandler) GetAllOrFindByName(c *gin.Context) {
	name := c.Query("name")

	if name != "" {
		h.FindByName(c)
	} else {
		h.GetAll(c)
	}
}
