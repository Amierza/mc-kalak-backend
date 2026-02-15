package entity

type (
	ClaimEvent  string
	ClaimStatus string
	VoteType    string
)

const (
	EventKing ClaimEvent = "KING"
	EventKong ClaimEvent = "KONG"
	EventNgok ClaimEvent = "NGOK"
)

const (
	StatusPending       ClaimStatus = "PENDING"
	StatusFinalApproved ClaimStatus = "FINAL_APPROVED"
	StatusFinalRejected ClaimStatus = "FINAL_REJECTED"
)

const (
	VoteApprove VoteType = "APPROVE"
	VoteReject  VoteType = "REJECT"
)
