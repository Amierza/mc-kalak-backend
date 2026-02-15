package service

import (
	"context"
	"fmt"
	"time"

	"github.com/Amierza/mc-kalak-backend/dto"
	"github.com/Amierza/mc-kalak-backend/entity"
	"github.com/Amierza/mc-kalak-backend/helper"
	"github.com/Amierza/mc-kalak-backend/jwt"
	"github.com/Amierza/mc-kalak-backend/repository"
	"github.com/Amierza/mc-kalak-backend/response"
	"github.com/google/uuid"
)

type (
	IClaimService interface {
		Create(ctx context.Context, req *dto.CreateClaimRequest) (*dto.ClaimResponse, error)
		GetAllWithPagination(ctx context.Context, req response.PaginationRequest) (dto.ClaimPaginationResponse, error)
		GetDetailByID(ctx context.Context, id *uuid.UUID) (*dto.ClaimResponse, error)
		Update(ctx context.Context, req *dto.UpdateClaimRequest) (*dto.ClaimResponse, error)
		DeleteByID(ctx context.Context, id *uuid.UUID) (*dto.ClaimResponse, error)
		Vote(ctx context.Context, req *dto.ClaimVoteRequest) (*dto.ClaimResponse, error)
		GetAllVotesByClaimID(ctx context.Context, claimID *uuid.UUID) ([]dto.ClaimVoteResponse, error)
	}

	claimService struct {
		claimRepo repository.IClaimRepository
		userRepo  repository.IUserRepository
		voteRepo  repository.IVoteRepository
		jwt       jwt.IJWT
	}
)

func NewClaimService(claimRepo repository.IClaimRepository, userRepo repository.IUserRepository, voteRepo repository.IVoteRepository, jwt jwt.IJWT) *claimService {
	return &claimService{
		claimRepo: claimRepo,
		userRepo:  userRepo,
		voteRepo:  voteRepo,
		jwt:       jwt,
	}
}

func (cs *claimService) Create(ctx context.Context, req *dto.CreateClaimRequest) (*dto.ClaimResponse, error) {
	reporter, found, err := cs.userRepo.GetDetailByID(ctx, nil, &req.ReporterID)
	if err != nil {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed to get reporter by id: %v\n", err)
	}
	if !found {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed reporter not found: %v\n", err)
	}

	claimedPlayer, found, err := cs.userRepo.GetDetailByID(ctx, nil, &req.ClaimedPlayerID)
	if err != nil {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed to get claimed player by id: %v\n", err)
	}
	if !found {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed claimed player not found: %v\n", err)
	}

	date, err := helper.ParseDateTime(req.MatchDate)
	if err != nil {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed parse date: %v\n", err)
	}

	claim := &entity.Claim{
		ID:              uuid.New(),
		Event:           req.Event,
		Status:          entity.StatusPending,
		MatchDate:       date,
		TotalPlayer:     req.TotalPlayer,
		ScreenshotURL:   req.ScreenshotURL,
		ApproveCount:    0,
		RejectCount:     0,
		ClaimedPlayerID: claimedPlayer.ID,
		ClaimedPlayer:   *claimedPlayer,
		ReporterID:      reporter.ID,
		Reporter:        *reporter,
	}

	if err := cs.claimRepo.Create(ctx, nil, claim); err != nil {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed to create claim: %v\n", err)
	}

	res := &dto.ClaimResponse{
		ID:            claim.ID,
		Event:         claim.Event,
		Status:        claim.Status,
		MatchDate:     req.MatchDate,
		TotalPlayer:   claim.TotalPlayer,
		ScreenshotURL: claim.ScreenshotURL,
		ApproveCount:  claim.ApproveCount,
		RejectCount:   claim.RejectCount,
		ClaimedPlayer: dto.UserSimpleResponse{
			ID:        claim.ClaimedPlayer.ID,
			Username:  claim.ClaimedPlayer.Username,
			AvatarURL: claim.ClaimedPlayer.AvatarURL,
		},
		Reporter: dto.UserSimpleResponse{
			ID:        claim.Reporter.ID,
			Username:  claim.Reporter.Username,
			AvatarURL: claim.Reporter.AvatarURL,
		},
		TimestampTemplate: dto.TimestampTemplate{
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		},
	}

	return res, nil
}

func (cs *claimService) GetAllWithPagination(ctx context.Context, req response.PaginationRequest) (dto.ClaimPaginationResponse, error) {
	datas, err := cs.claimRepo.GetAllClaimsWithPagination(ctx, nil, req)
	if err != nil {
		return dto.ClaimPaginationResponse{}, fmt.Errorf("Failed to get all claims: %v\n", err)
	}

	claims := make([]*dto.ClaimResponse, 0, len(datas.Claims))
	for _, claim := range datas.Claims {
		claims = append(claims, &dto.ClaimResponse{
			ID:            claim.ID,
			Event:         claim.Event,
			Status:        claim.Status,
			MatchDate:     claim.MatchDate.Format("2006-01-02 15:04:05"),
			TotalPlayer:   claim.TotalPlayer,
			ScreenshotURL: claim.ScreenshotURL,
			ApproveCount:  claim.ApproveCount,
			RejectCount:   claim.RejectCount,
			ClaimedPlayer: dto.UserSimpleResponse{
				ID:        claim.ClaimedPlayer.ID,
				Username:  claim.ClaimedPlayer.Username,
				AvatarURL: claim.ClaimedPlayer.AvatarURL,
			},
			Reporter: dto.UserSimpleResponse{
				ID:        claim.Reporter.ID,
				Username:  claim.Reporter.Username,
				AvatarURL: claim.Reporter.AvatarURL,
			},
			TimestampTemplate: dto.TimestampTemplate{
				CreatedAt: claim.CreatedAt.Format("2006-01-02 15:04:05"),
				UpdatedAt: claim.UpdatedAt.Format("2006-01-02 15:04:05"),
			},
		})
	}

	return dto.ClaimPaginationResponse{
		Data: claims,
		PaginationResponse: response.PaginationResponse{
			Page:    datas.Page,
			PerPage: datas.PerPage,
			MaxPage: datas.MaxPage,
			Count:   datas.Count,
		},
	}, nil
}

func (cs *claimService) GetDetailByID(ctx context.Context, id *uuid.UUID) (*dto.ClaimResponse, error) {
	claim, found, err := cs.claimRepo.GetDetailByID(ctx, nil, id)
	if err != nil {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed to get claim by id: %v\n", err)
	}
	if !found {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed claim not found: %v\n", err)
	}

	res := &dto.ClaimResponse{
		ID:            claim.ID,
		Event:         claim.Event,
		Status:        claim.Status,
		MatchDate:     claim.MatchDate.Format("2006-01-02 15:04:05"),
		TotalPlayer:   claim.TotalPlayer,
		ScreenshotURL: claim.ScreenshotURL,
		ApproveCount:  claim.ApproveCount,
		RejectCount:   claim.RejectCount,
		ClaimedPlayer: dto.UserSimpleResponse{
			ID:        claim.ClaimedPlayer.ID,
			Username:  claim.ClaimedPlayer.Username,
			AvatarURL: claim.ClaimedPlayer.AvatarURL,
		},
		Reporter: dto.UserSimpleResponse{
			ID:        claim.Reporter.ID,
			Username:  claim.Reporter.Username,
			AvatarURL: claim.Reporter.AvatarURL,
		},
		TimestampTemplate: dto.TimestampTemplate{
			CreatedAt: claim.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: claim.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	}

	return res, nil
}

func (cs *claimService) Update(ctx context.Context, req *dto.UpdateClaimRequest) (*dto.ClaimResponse, error) {
	claim, found, err := cs.claimRepo.GetDetailByID(ctx, nil, &req.ID)
	if err != nil {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed to get claim by id: %v\n", err)
	}
	if !found {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed claim not found: %v\n", err)
	}

	_, found, err = cs.userRepo.GetDetailByID(ctx, nil, &req.ReporterID)
	if err != nil {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed to get reporter by id: %v\n", err)
	}
	if !found {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed reporter not found: %v\n", err)
	}

	_, found, err = cs.userRepo.GetDetailByID(ctx, nil, &req.ClaimedPlayerID)
	if err != nil {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed to get claimed player by id: %v\n", err)
	}
	if !found {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed claimed player not found: %v\n", err)
	}

	date, err := helper.ParseDateTime(req.MatchDate)
	if err != nil {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed parse date: %v\n", err)
	}

	claim.Event = req.Event
	claim.MatchDate = date
	claim.TotalPlayer = req.TotalPlayer
	claim.ScreenshotURL = req.ScreenshotURL
	claim.ClaimedPlayerID = req.ClaimedPlayerID
	claim.ReporterID = req.ReporterID

	if err := cs.claimRepo.Update(ctx, nil, claim); err != nil {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed to update claim: %v\n", err)
	}

	res := &dto.ClaimResponse{
		ID:            claim.ID,
		Event:         claim.Event,
		Status:        claim.Status,
		MatchDate:     req.MatchDate,
		TotalPlayer:   claim.TotalPlayer,
		ScreenshotURL: claim.ScreenshotURL,
		ApproveCount:  claim.ApproveCount,
		RejectCount:   claim.RejectCount,
		ClaimedPlayer: dto.UserSimpleResponse{
			ID:        claim.ClaimedPlayer.ID,
			Username:  claim.ClaimedPlayer.Username,
			AvatarURL: claim.ClaimedPlayer.AvatarURL,
		},
		Reporter: dto.UserSimpleResponse{
			ID:        claim.Reporter.ID,
			Username:  claim.Reporter.Username,
			AvatarURL: claim.Reporter.AvatarURL,
		},
		TimestampTemplate: dto.TimestampTemplate{
			CreatedAt: claim.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: claim.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	}

	return res, nil
}

func (cs *claimService) DeleteByID(ctx context.Context, id *uuid.UUID) (*dto.ClaimResponse, error) {
	deletedClaim, found, err := cs.claimRepo.GetDetailByID(ctx, nil, id)
	if err != nil {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed to get claim by id: %v\n", err)
	}
	if !found {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed claim not found: %v\n", err)
	}

	err = cs.claimRepo.DeleteByID(ctx, nil, id)
	if err != nil {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed to delete claim by id: %v\n", err)
	}

	res := &dto.ClaimResponse{
		ID:            deletedClaim.ID,
		Event:         deletedClaim.Event,
		Status:        deletedClaim.Status,
		MatchDate:     deletedClaim.MatchDate.Format("2006-01-02 15:04:05"),
		TotalPlayer:   deletedClaim.TotalPlayer,
		ScreenshotURL: deletedClaim.ScreenshotURL,
		ApproveCount:  deletedClaim.ApproveCount,
		RejectCount:   deletedClaim.RejectCount,
		ClaimedPlayer: dto.UserSimpleResponse{
			ID:        deletedClaim.ClaimedPlayer.ID,
			Username:  deletedClaim.ClaimedPlayer.Username,
			AvatarURL: deletedClaim.ClaimedPlayer.AvatarURL,
		},
		Reporter: dto.UserSimpleResponse{
			ID:        deletedClaim.Reporter.ID,
			Username:  deletedClaim.Reporter.Username,
			AvatarURL: deletedClaim.Reporter.AvatarURL,
		},
		TimestampTemplate: dto.TimestampTemplate{
			CreatedAt: deletedClaim.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: deletedClaim.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	}

	return res, nil
}

func (cs *claimService) Vote(ctx context.Context, req *dto.ClaimVoteRequest) (*dto.ClaimResponse, error) {
	claim, found, err := cs.claimRepo.GetDetailByID(ctx, nil, &req.ID)
	if err != nil {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed to get claim by id: %v\n", err)
	}
	if !found {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed claim not found: %v\n", err)
	}

	if claim.Status != entity.StatusPending {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed claim is over: %v\n", err)
	}

	token := ctx.Value("Authorization").(string)
	userIDString, err := cs.jwt.GetUserIDByToken(token)
	if err != nil {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed to get user ID by token: %v\n", err)
	}
	userID, err := uuid.Parse(userIDString)
	if err != nil {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed parse id from string to uuid: %v\n", err)
	}

	_, found, err = cs.voteRepo.GetByClaimIDAndVoterID(ctx, nil, &claim.ID, &userID)
	if err != nil {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed to get vote by claim id and voter ID: %v\n", err)
	}
	if found {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed already vote: %v\n", err)
	}

	if req.Type == string(entity.VoteApprove) {
		claim.ApproveCount++
	} else {
		claim.RejectCount++
	}

	vote := &entity.Vote{
		ID:      uuid.New(),
		ClaimID: claim.ID,
		VoterID: userID,
		Type:    entity.VoteType(req.Type),
	}
	if err := cs.voteRepo.Create(ctx, nil, vote); err != nil {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed to create vote: %v\n", err)
	}

	totalUser, err := cs.userRepo.Count(ctx, nil)
	if err != nil {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed to count total user: %v\n", err)
	}

	minApprove := (totalUser / 2) + 1
	if claim.ApproveCount >= int(minApprove) {
		claim.Status = entity.StatusFinalApproved
	}
	remainingVote := int(totalUser) - (claim.ApproveCount + claim.RejectCount)
	maxPossibleApprove := claim.ApproveCount + remainingVote
	if maxPossibleApprove < int(minApprove) {
		claim.Status = entity.StatusFinalRejected
	}

	if err := cs.claimRepo.Update(ctx, nil, claim); err != nil {
		return &dto.ClaimResponse{}, fmt.Errorf("Failed to update claim: %v\n", err)
	}

	res := &dto.ClaimResponse{
		ID:            claim.ID,
		Event:         claim.Event,
		Status:        claim.Status,
		MatchDate:     claim.MatchDate.Format("2006-01-02 15:04:05"),
		TotalPlayer:   claim.TotalPlayer,
		ScreenshotURL: claim.ScreenshotURL,
		ApproveCount:  claim.ApproveCount,
		RejectCount:   claim.RejectCount,
		ClaimedPlayer: dto.UserSimpleResponse{
			ID:        claim.ClaimedPlayer.ID,
			Username:  claim.ClaimedPlayer.Username,
			AvatarURL: claim.ClaimedPlayer.AvatarURL,
		},
		Reporter: dto.UserSimpleResponse{
			ID:        claim.Reporter.ID,
			Username:  claim.Reporter.Username,
			AvatarURL: claim.Reporter.AvatarURL,
		},
		TimestampTemplate: dto.TimestampTemplate{
			CreatedAt: claim.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: claim.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	}

	return res, nil
}

func (cs *claimService) GetAllVotesByClaimID(ctx context.Context, claimID *uuid.UUID) ([]dto.ClaimVoteResponse, error) {
	datas, err := cs.voteRepo.GetAllByClaimID(ctx, nil, claimID)
	if err != nil {
		return []dto.ClaimVoteResponse{}, fmt.Errorf("Failed to get all votes by claim ID: %v\n", err)
	}

	votes := make([]dto.ClaimVoteResponse, 0, len(datas))
	for _, vote := range datas {
		votes = append(votes, dto.ClaimVoteResponse{
			ID: vote.ID,
			Voter: dto.UserSimpleResponse{
				ID:        vote.Voter.ID,
				Username:  vote.Voter.Username,
				AvatarURL: vote.Voter.AvatarURL,
			},
			Type: vote.Type,
		})
	}

	return votes, nil
}
