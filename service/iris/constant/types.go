package constant

const (
	Iris_TxTypeTransfer                    = "Transfer"
	Iris_TxTypeBurn                        = "Burn"
	Iris_TxTypeStakeCreateValidator        = "CreateValidator"
	Iris_TxTypeStakeEditValidator          = "EditValidator"
	Iris_TxTypeStakeDelegate               = "Delegate"
	Iris_TxTypeStakeBeginUnbonding         = "BeginUnbonding"
	Iris_TxTypeBeginRedelegate             = "BeginRedelegate"
	Iris_TxTypeUnjail                      = "Unjail"
	Iris_TxTypeSetWithdrawAddress          = "SetWithdrawAddress"
	Iris_TxTypeWithdrawDelegatorReward     = "WithdrawDelegatorReward"
	Iris_TxTypeWithdrawDelegatorRewardsAll = "WithdrawDelegatorRewardsAll"
	Iris_TxTypeWithdrawValidatorRewardsAll = "WithdrawValidatorRewardsAll"
	Iris_TxTypeSubmitProposal              = "SubmitProposal"
	Iris_TxTypeDeposit                     = "Deposit"
	Iris_TxTypeVote                        = "Vote"

	TxTypeAssetIssueToken           = "IssueToken"
	TxTypeAssetEditToken            = "EditToken"
	TxTypeAssetMintToken            = "MintToken"
	TxTypeAssetTransferTokenOwner   = "TransferTokenOwner"
	TxTypeAssetCreateGateway        = "CreateGateway"
	TxTypeAssetEditGateway          = "EditGateway"
	TxTypeAssetTransferGatewayOwner = "TransferGatewayOwner"

	TxMsgTypeAssetIssueToken           = "IssueToken"
	TxMsgTypeAssetEditToken            = "EditToken"
	TxMsgTypeAssetMintToken            = "MintToken"
	TxMsgTypeAssetTransferTokenOwner   = "TransferTokenOwner"
	TxMsgTypeAssetCreateGateway        = "CreateGateway"
	TxMsgTypeAssetEditGateway          = "EditGateway"
	TxMsgTypeAssetTransferGatewayOwner = "TransferGatewayOwner"

	TxStatusSuccess = "success"
	TxStatusFail    = "fail"
)
