package dto

import (
	"errors"

	"github.com/Amierza/mc-kalak-backend/entity"
	"github.com/Amierza/mc-kalak-backend/response"
	"github.com/google/uuid"
)

const (
	// ====================================== Failed ======================================
	MESSAGE_INVALID_REQUEST_PAYLOAD = "invalid request payload"

	// Middleware
	MESSAGE_FAILED_PROSES_REQUEST      = "failed proses request"
	MESSAGE_FAILED_ACCESS_DENIED       = "failed access denied"
	MESSAGE_FAILED_TOKEN_NOT_FOUND     = "failed token not found"
	MESSAGE_FAILED_TOKEN_NOT_VALID     = "failed token not valid"
	MESSAGE_FAILED_TOKEN_DENIED_ACCESS = "failed token denied access"
	MESSAGE_FAILED_GET_CUSTOM_CLAIMS   = "failed get custom claims"

	// Query Params
	MESSAGE_INVALID_QUERY_PARAMS = "invalid query params"

	// UUID
	MESSAGE_FAILED_INVALID_UUID = "invalid UUID format"

	// File
	MESSAGE_FAILED_PARSE_MULTIPART_FORM = "failed to parse multipart form"
	MESSAGE_FAILED_NO_FILES_UPLOADED    = "failed no files uploaded"
	MESSAGE_FAILED_UPLOAD_FILES         = "failed upload files"

	// General Errors
	FAILED_LOGIN          = "failed login"
	FAILED_CREATE         = "failed to create"
	FAILED_UPDATE         = "failed to update"
	FAILED_DELETE         = "failed to delete"
	FAILED_GET_ALL        = "failed to get all"
	FAILED_GET_DETAIL     = "failed to get detail"
	FAILED_GET_PROFILE    = "failed get profile"
	NOT_FOUND             = "not found"
	INTERNAL_SERVER_ERROR = "internal server error"

	// ====================================== Success ======================================
	// File
	MESSAGE_SUCCESS_UPLOAD_FILES = "success upload files"
	MESSAGE_SUCCESS_UPLOAD_FILE  = "success upload file"

	// General Success
	SUCCESS_LOGIN       = "success login"
	SUCCESS_CREATE      = "success create"
	SUCCESS_UPDATE      = "success update"
	SUCCESS_DELETE      = "success delete"
	SUCCESS_GET_ALL     = "success get all"
	SUCCESS_GET_DETAIL  = "success get detail"
	SUCCESS_GET_PROFILE = "success get profile"
)

var (
	// Token
	ErrGenerateToken           = errors.New("failed to generate token")
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
	ErrDecryptToken            = errors.New("failed to decrypt token")
	ErrTokenInvalid            = errors.New("token invalid")
	ErrValidateToken           = errors.New("failed to validate token")

	// File
	ErrNoFilesUploaded    = errors.New("failed no files uploaded")
	ErrInvalidFileType    = errors.New("invalid file type")
	ErrSaveFile           = errors.New("failed save file")
	ErrCreateFolderAssets = errors.New("failed create folder assets")
	ErrDeleteOldImage     = errors.New("failed to delete old image")

	// General
	ErrNotFound         = errors.New("not found")
	ErrValidationFailed = errors.New("validation failed")
	ErrAlreadyExists    = errors.New("already exists")
	ErrInternal         = errors.New("error internal")
	ErrUnauthorized     = errors.New("unauthorized")

	// Input

	// Parse
)

// Timestamp
type (
	TimestampTemplate struct {
		CreatedAt string  `json:"created_at"`
		UpdatedAt string  `json:"updated_at"`
		DeletedAt *string `json:"deleted_at,omitempty"`
	}
)

// Auth
type (
	LoginRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	LoginResponse struct {
		Token string `json:"token"`
	}
)

// User
type (
	UserResponse struct {
		ID        uuid.UUID `json:"id"`
		Username  string    `json:"username"`
		Password  string    `json:"password"`
		AvatarURL string    `json:"avatar_url"`
		IsActive  bool      `json:"is_active"`
		TimestampTemplate
	}
	UpdateProfileRequest struct {
		AvatarURL string `binding:"required" json:"avatar_url"`
	}
	UserSimpleResponse struct {
		ID        uuid.UUID `json:"id"`
		Username  string    `json:"username"`
		AvatarURL string    `json:"avatar_url"`
	}
)

// Claim
type (
	CreateClaimRequest struct {
		Event           entity.ClaimEvent `binding:"required,oneof=KING KONG NGOK" json:"event"`
		MatchDate       string            `binding:"required" json:"match_date"`
		TotalPlayer     int               `binding:"required,min=2,max=8" json:"total_player"`
		ScreenshotURL   string            `binding:"required" json:"screenshot_url"`
		ClaimedPlayerID uuid.UUID         `binding:"required" json:"claimed_player_id"`
		ReporterID      uuid.UUID         `binding:"required" json:"reporter_id"`
	}
	ClaimResponse struct {
		ID            uuid.UUID          `json:"id"`
		Event         entity.ClaimEvent  `json:"event"`
		Status        entity.ClaimStatus `json:"status"`
		MatchDate     string             `json:"match_date"`
		TotalPlayer   int                `json:"total_player"`
		ScreenshotURL string             `json:"screenshot_url"`
		ApproveCount  int                `json:"approve_count"`
		RejectCount   int                `json:"reject_count"`
		ClaimedPlayer UserSimpleResponse `json:"claimed_player"`
		Reporter      UserSimpleResponse `json:"reporter"`
		TimestampTemplate
	}
	UpdateClaimRequest struct {
		ID              uuid.UUID         `json:"-"`
		Event           entity.ClaimEvent `binding:"required,oneof=KING KONG NGOK" json:"event"`
		MatchDate       string            `binding:"required" json:"match_date"`
		TotalPlayer     int               `binding:"required,min=2,max=8" json:"total_player"`
		ScreenshotURL   string            `binding:"required" json:"screenshot_url"`
		ClaimedPlayerID uuid.UUID         `binding:"required" json:"claimed_player_id"`
		ReporterID      uuid.UUID         `binding:"required" json:"reporter_id"`
	}
	ClaimVoteRequest struct {
		ID   uuid.UUID `json:"-"`
		Type string    `binding:"required,oneof=APPROVE REJECT" json:"type"`
	}
	ClaimVoteResponse struct {
		ID    uuid.UUID          `json:"id"`
		Voter UserSimpleResponse `json:"voter"`
		Type  entity.VoteType    `json:"type"`
	}
	ClaimPaginationResponse struct {
		response.PaginationResponse
		Data []*ClaimResponse `json:"data"`
	}
	ClaimPaginationRepositoryResponse struct {
		response.PaginationResponse
		Claims []*entity.Claim
	}
)
