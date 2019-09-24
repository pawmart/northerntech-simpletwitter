package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/imdario/mergo"
	"github.com/pawmart/northerntech-simpletwitter/internal/models"
	o "github.com/pawmart/northerntech-simpletwitter/internal/restapi/operations"
	"github.com/pawmart/northerntech-simpletwitter/internal/storage"
	"github.com/pawmart/northerntech-simpletwitter/pkg/utils"
)

// NewHandler function.
func NewHandler(storage storage.Storage) *TweetsHandler {
	return &TweetsHandler{storage: storage}
}
// Handler responsible for tweets endpoints.
type TweetsHandler struct {
	storage storage.Storage
}

// GetHealth handling.
func (h *TweetsHandler) GetHealth(params o.GetHealthParams) middleware.Responder {

	result := new(models.Health)
	result.Status = "up"

	err := h.storage.Ping()
	if err != nil {
		result.Status = "down"
	}

	return o.NewGetHealthOK().WithPayload(result)
}

// GetTweets handling.
func (h *TweetsHandler) GetTweets(params o.GetTweetsParams) middleware.Responder {

	var err error

	result := h.storage.FindTweets(params)
	if err != nil {
		log.Print("could not list resources, db query problems", params)
		return o.NewGetTweetsInternalServerError()
	}

	resp := new(models.TweetDetailsListResponse)
	resp.Data = result
	resp.Links = &models.Links{Self: new(o.GetTweetsURL).String()}

	return o.NewGetTweetsOK().WithPayload(resp)
}

// GetTweet handling.
func (h *TweetsHandler) GetTweet(params o.GetTweetsIDParams) middleware.Responder {

	p, err := h.storage.FindTweet(params.ID.String())
	if err != nil {
		return o.NewGetTweetsIDNotFound().WithPayload(&models.APIError{
			ErrorCode:    string(http.StatusNotFound),
			ErrorMessage: "not found",
		})
	}

	selfUrl := new(o.GetTweetsIDURL)
	selfUrl.ID = strfmt.UUID(p.ID)

	pd := new(models.TweetDetailsResponse)
	pd.Data = p
	pd.Links = &models.Links{Self: "/v1/tweets/" + p.ID}
	pd.Links = &models.Links{Self: selfUrl.String()}

	return o.NewGetTweetsIDOK().WithPayload(pd)
}

// CreateTweet handling.
func (h *TweetsHandler) CreateTweet(params o.PostTweetsParams) middleware.Responder {

	pd := params.TweetCreationRequest
	if pd.Data == nil {
		return o.NewPostTweetsBadRequest().WithPayload(&models.APIError{
			ErrorCode: string(http.StatusBadRequest), ErrorMessage: "bad request"})
	}

	id := utils.GenerateUUIDString()
	t := time.Now().Unix()

	pd.Data.ID = id
	pd.Data.CreatedOn = &t
	pd.Data.ModifiedOn = &t

	s := h.storage
	if err := s.InsertTweet(pd.Data); err != nil {
		log.Print("could not insert resource, db problems", err)
		return o.NewPostTweetsInternalServerError()
	}

	// Now get it...
	p, err := s.FindTweet(id)
	if err != nil {
		log.Print("freshly created entity could not be fetched", id)
		return o.NewPostTweetsInternalServerError()
	}

	selfUrl := new(o.GetTweetsIDURL)
	selfUrl.ID = strfmt.UUID(p.ID)

	pdr := new(models.TweetCreationResponse)
	pdr.Data = p
	pdr.Links = &models.Links{Self: selfUrl.String()}

	return o.NewPostTweetsCreated().WithPayload(pdr)
}

// UpdateTweet handling.
func (h *TweetsHandler) UpdateTweet(params o.PatchTweetsParams) middleware.Responder {

	src := params.TweetUpdateRequest.Data

	id := src.ID
	s := h.storage
	dst, err := s.FindTweet(id)
	if err != nil {
		return o.NewPatchTweetsNotFound().WithPayload(&models.APIError{
			ErrorCode: string(http.StatusNotFound), ErrorMessage: "not found"})
	}

	t := time.Now().Unix()
	dst.ModifiedOn = &t

	if err := mergo.Merge(dst, src, mergo.WithOverride); err != nil {
		log.Print("resource not merged", src, dst, err.Error())
		return o.NewPatchTweetsInternalServerError()
	}

	if err := s.UpdateTweet(id, dst); err != nil {
		log.Print("resource not updated", dst, err.Error())
		return o.NewPatchTweetsInternalServerError()
	}

	return o.NewPatchTweetsOK()
}

// DeleteTweet handling.
func (h *TweetsHandler) DeleteTweet(params o.DeleteTweetsIDParams) middleware.Responder {

	if err := h.storage.RemoveTweet(params.ID.String()); err != nil {
		return o.NewDeleteTweetsIDNotFound().WithPayload(&models.APIError{
			ErrorCode: string(http.StatusNotFound), ErrorMessage: "not found"})
	}

	return o.NewDeleteTweetsIDNoContent()
}
