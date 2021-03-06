package retrievalmarket

// DealStatus is the status of a retrieval deal returned by a provider
// in a DealResponse
type DealStatus uint64

const (
	// DealStatusNew is a deal that nothing has happened with yet
	DealStatusNew DealStatus = iota

	// DealStatusUnsealing means the provider is unsealing data
	DealStatusUnsealing

	// DealStatusUnsealed means the provider has finished unsealing data
	DealStatusUnsealed

	// DealStatusWaitForAcceptance means we're waiting to hear back if the provider accepted our deal
	DealStatusWaitForAcceptance

	// DealStatusPaymentChannelCreating is the status set while waiting for the
	// payment channel creation to complete
	DealStatusPaymentChannelCreating

	// DealStatusPaymentChannelAddingFunds is the status when we are waiting for funds
	// to finish being sent to the payment channel
	DealStatusPaymentChannelAddingFunds

	// DealStatusAccepted means a deal has been accepted by a provider
	// and its is ready to proceed with retrieval
	DealStatusAccepted

	// DealStatusFailing indicates something went wrong during a retrieval,
	// and we are cleaning up before termianting with an error
	DealStatusFailing

	// DealStatusRejected indicates the provider rejected a client's deal proposal
	// for some reason
	DealStatusRejected

	// DealStatusFundsNeeded indicates the provider needs a payment voucher to
	// continue processing the deal
	DealStatusFundsNeeded

	// DealStatusSendFunds indicats the client is now going to send funds because we reached the threshold of the last payment
	DealStatusSendFunds

	// DealStatusSendFundsLastPayment indicats the client is now going to send final funds because
	// we reached the threshold of the final payment
	DealStatusSendFundsLastPayment

	// DealStatusOngoing indicates the provider is continuing to process a deal
	DealStatusOngoing

	// DealStatusFundsNeededLastPayment indicates the provider needs a payment voucher
	// in order to complete a deal
	DealStatusFundsNeededLastPayment

	// DealStatusCompleted indicates a deal is complete
	DealStatusCompleted

	// DealStatusDealNotFound indicates an update was received for a deal that could
	// not be identified
	DealStatusDealNotFound

	// DealStatusErrored indicates a deal has terminated in an error
	DealStatusErrored

	// DealStatusBlocksComplete indicates that all blocks have been processed for the piece
	DealStatusBlocksComplete

	// DealStatusFinalizing means the last payment has been received and
	// we are just confirming the deal is complete
	DealStatusFinalizing

	// DealStatusCompleting is just an inbetween state to perform final cleanup of
	// complete deals
	DealStatusCompleting
)

// DealStatuses maps deal status to a human readable representation
var DealStatuses = map[DealStatus]string{
	DealStatusNew:                       "DealStatusNew",
	DealStatusUnsealing:                 "DealStatusUnsealing",
	DealStatusUnsealed:                  "DealStatusUnsealed",
	DealStatusWaitForAcceptance:         "DealStatusWaitForAcceptance",
	DealStatusPaymentChannelCreating:    "DealStatusPaymentChannelCreating",
	DealStatusPaymentChannelAddingFunds: "DealStatusPaymentChannelAddingFunds",
	DealStatusAccepted:                  "DealStatusAccepted",
	DealStatusFailing:                   "DealStatusFailed",
	DealStatusRejected:                  "DealStatusRejected",
	DealStatusFundsNeeded:               "DealStatusFundsNeeded",
	DealStatusSendFunds:                 "DealStatusSendFunds",
	DealStatusSendFundsLastPayment:      "DealStatusSendFundsLastPayment",
	DealStatusOngoing:                   "DealStatusOngoing",
	DealStatusFundsNeededLastPayment:    "DealStatusFundsNeededLastPayment",
	DealStatusCompleted:                 "DealStatusCompleted",
	DealStatusDealNotFound:              "DealStatusDealNotFound",
	DealStatusErrored:                   "DealStatusErrored",
	DealStatusBlocksComplete:            "DealStatusBlocksComplete",
	DealStatusFinalizing:                "DealStatusFinalizing",
	DealStatusCompleting:                "DealStatusCompleting",
}
