package constant

const (
	Cosmos_TxTypeMultiSend                   = "MultiSend"
	Cosmos_TxTypeVerifyInvariant             = "VerifyInvariant"
	Cosmos_TxTypeTransfer                    = "Transfer"
	Cosmos_TxTypeStakeCreateValidator        = "CreateValidator"
	Cosmos_TxTypeStakeEditValidator          = "EditValidator"
	Cosmos_TxTypeStakeDelegate               = "Delegate"
	Cosmos_TxTypeStakeUnDelegate             = "UnDelegate"
	Cosmos_TxTypeBeginRedelegate             = "BeginRedelegate"
	Cosmos_TxTypeUnjail                      = "Unjail"
	Cosmos_TxTypeSetWithdrawAddress          = "SetWithdrawAddress"
	Cosmos_TxTypeWithdrawDelegatorReward     = "WithdrawDelegatorReward"
	Cosmos_TxTypeWithdrawValidatorCommission = "WithdrawValidatorCommission"
	Cosmos_TxTypeSubmitProposal              = "SubmitProposal"
	Cosmos_TxTypeDeposit                     = "Deposit"
	Cosmos_TxTypeVote                        = "Vote"

	Cosmos_TxEventWithdrawRewards = "withdraw_rewards"

	TxStatusSuccess = "success"
	TxStatusFail    = "fail"
)
