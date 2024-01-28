package domain

type SubscriberService interface {
	GetSubscriberType(msisdn string) (string, error)
}
