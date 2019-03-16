
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:09</date>
//</624456008722288640>

/*
版权所有State Street Corp.保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package endorser

import "github.com/hyperledger/fabric/common/metrics"

var (
	proposalDurationHistogramOpts = metrics.HistogramOpts{
		Namespace:    "endorser",
		Name:         "propsal_duration",
		Help:         "The time to complete a proposal.",
		LabelNames:   []string{"channel", "chaincode", "success"},
		StatsdFormat: "%{#fqname}.%{channel}.%{chaincode}.%{success}",
	}

	receivedProposalsCounterOpts = metrics.CounterOpts{
		Namespace: "endorser",
		Name:      "proposals_received",
		Help:      "The number of proposals received.",
	}

	successfulProposalsCounterOpts = metrics.CounterOpts{
		Namespace: "endorser",
		Name:      "successful_proposals",
		Help:      "The number of successful proposals.",
	}

	proposalValidationFailureCounterOpts = metrics.CounterOpts{
		Namespace: "endorser",
		Name:      "proposal_validation_failures",
		Help:      "The number of proposals that have failed initial validation.",
	}

	proposalChannelACLFailureOpts = metrics.CounterOpts{
		Namespace:    "endorser",
		Name:         "proposal_acl_failures",
		Help:         "The number of proposals that failed ACL checks.",
		LabelNames:   []string{"channel", "chaincode"},
		StatsdFormat: "%{#fqname}.%{channel}.%{chaincode}",
	}

	initFailureCounterOpts = metrics.CounterOpts{
		Namespace:    "endorser",
		Name:         "chaincode_instantiation_failures",
		Help:         "The number of chaincode instantiations or upgrade that have failed.",
		LabelNames:   []string{"channel", "chaincode"},
		StatsdFormat: "%{#fqname}.%{channel}.%{chaincode}",
	}

	endorsementFailureCounterOpts = metrics.CounterOpts{
		Namespace:    "endorser",
		Name:         "endorsement_failures",
		Help:         "The number of failed endorsements.",
		LabelNames:   []string{"channel", "chaincode", "chaincodeerror"},
		StatsdFormat: "%{#fqname}.%{channel}.%{chaincode}.%{chaincodeerror}",
	}

	duplicateTxsFailureCounterOpts = metrics.CounterOpts{
		Namespace:    "endorser",
		Name:         "duplicate_transaction_failures",
		Help:         "The number of failed proposals due to duplicate transaction ID.",
		LabelNames:   []string{"channel", "chaincode"},
		StatsdFormat: "%{#fqname}.%{channel}.%{chaincode}",
	}
)

type EndorserMetrics struct {
	ProposalDuration         metrics.Histogram
	ProposalsReceived        metrics.Counter
	SuccessfulProposals      metrics.Counter
	ProposalValidationFailed metrics.Counter
	ProposalACLCheckFailed   metrics.Counter
	InitFailed               metrics.Counter
	EndorsementsFailed       metrics.Counter
	DuplicateTxsFailure      metrics.Counter
}

func NewEndorserMetrics(p metrics.Provider) *EndorserMetrics {
	return &EndorserMetrics{
		ProposalDuration:         p.NewHistogram(proposalDurationHistogramOpts),
		ProposalsReceived:        p.NewCounter(receivedProposalsCounterOpts),
		SuccessfulProposals:      p.NewCounter(successfulProposalsCounterOpts),
		ProposalValidationFailed: p.NewCounter(proposalValidationFailureCounterOpts),
		ProposalACLCheckFailed:   p.NewCounter(proposalChannelACLFailureOpts),
		InitFailed:               p.NewCounter(initFailureCounterOpts),
		EndorsementsFailed:       p.NewCounter(endorsementFailureCounterOpts),
		DuplicateTxsFailure:      p.NewCounter(duplicateTxsFailureCounterOpts),
	}
}

