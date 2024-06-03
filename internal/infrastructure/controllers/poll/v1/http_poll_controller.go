package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	application "github.com/yunusemre12500/poll-api/internal/application/poll/v1"
	domain "github.com/yunusemre12500/poll-api/internal/domain/poll/v1"
)

type HTTPPollController struct {
	service domain.PollService
}

func NewHTTPPollController(service domain.PollService) *HTTPPollController {
	return &HTTPPollController{
		service: service,
	}
}

func (controller *HTTPPollController) Create(ctx *gin.Context) {
	var createPollRequestBody *application.CreatePollRequestBody

	if err := ctx.ShouldBindJSON(&createPollRequestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, &gin.H{
			"code":    "DecodeError",
			"message": "Failed to decode request body.",
		})

		return
	}

	newPoll := domain.NewPollFromCreatePollRequestBody(createPollRequestBody)

	if err := controller.service.Create(ctx, newPoll); err != nil {
		if err == domain.ErrPollExists {
			ctx.JSON(http.StatusConflict, &gin.H{
				"code":    "AlreadyExists",
				"message": "Poll already exists.",
			})

			return
		}

		ctx.JSON(http.StatusInternalServerError, &gin.H{
			"code":    "InternalServerError",
			"message": "Failed to save created poll.",
		})

		return
	}

	ctx.JSON(http.StatusCreated, newPoll.IntoCreatePollResponseBody())
}

func (controller *HTTPPollController) GetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := uuid.Parse(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, &gin.H{
			"code":    "InvalidPathParameter",
			"message": "Failed to parse 'id' path parameter.",
		})

		return
	}

	poll, err := controller.service.GetByID(ctx, &id)

	if err != nil {
		if err == domain.ErrPollNotFound {
			ctx.JSON(http.StatusNotFound, &gin.H{
				"code":    "NotFound",
				"message": "Poll not found.",
			})

			return
		}

		ctx.JSON(http.StatusInternalServerError, &gin.H{
			"code":    "InternalServerError",
			"message": "Failed to get poll by id.",
		})

		return
	}

	ctx.JSON(http.StatusOK, poll.IntoGetPollByIdResponseBody())
}

func (controller *HTTPPollController) List(ctx *gin.Context) {
	var limit uint
	var offset int

	limitQueryParamValue, exists := ctx.GetQuery("limit")

	if exists {
		limitQueryParamParsedValue, err := strconv.ParseInt(limitQueryParamValue, 0, 64)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, &gin.H{
				"code":    "InvalidQueryParameter",
				"message": "Failed to parse 'limit' query parameter.",
			})

			return
		}

		limit = uint(limitQueryParamParsedValue)
	} else {
		limit = 100
	}

	offsetQueryParamValue, exists := ctx.GetQuery("offset")

	if exists {
		offsetQueryParamParsedValue, err := strconv.ParseInt(offsetQueryParamValue, 0, 64)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, &gin.H{
				"code":    "InvalidQueryParameter",
				"message": "Failed to parse 'offset' query parameter.",
			})

			return
		}

		offset = int(offsetQueryParamParsedValue)
	} else {
		offset = 0
	}

	if limit < 2 || limit > 100 {
		ctx.JSON(http.StatusBadRequest, &gin.H{
			"code":    "BadRequest",
			"message": "Query parameter 'limit' must be between 2 and 100.",
		})

		return
	}

	if offset < 0 {
		ctx.JSON(http.StatusBadRequest, &gin.H{
			"code":    "BadRequest",
			"message": "Query parameter 'offset' must be greater than or equal to 0.",
		})

		return
	}

	polls, err := controller.service.List(ctx, limit, uint(offset))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &gin.H{
			"code":    "InternalServerError",
			"message": "Failed to get polls.",
		})

		return
	}

	if len(polls) == 0 {
		ctx.JSON(http.StatusNoContent, &gin.H{
			"code":    "NotFound",
			"message": "No polls not found.",
		})

		return
	}

	var listedPolls []*application.ListPollsResponseBody

	for _, poll := range polls {
		listedPolls = append(listedPolls, poll.IntoListPollsResponseBody())
	}

	ctx.JSON(http.StatusOK, listedPolls)
}
