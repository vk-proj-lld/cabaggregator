package signals

const (
	AckAccept AckSignal = iota + 1
	AckReject
	AckAcceptFinishingIn10
)

type AckSignal int

func (a *AckSignal) String() string {
	if a == nil {
		return "invalid signal"
	}
	switch *a {
	case AckAccept:
		return "Accepted"
	case AckReject:
		return "Rejected"
	case AckAcceptFinishingIn10:
		return "ConditionAccepting-10mn"
	default:
		return "invalid signal"
	}
}
