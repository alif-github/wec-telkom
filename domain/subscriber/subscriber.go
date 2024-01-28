package domain

type RequestSubscriber struct {
	Initialize bool   `json:"initialize"`
	Msisdn     string `json:"msisdn"`
}

type ResponseSubscriber struct {
	Msisdn         string `json:"Msisdn"`
	Status         int    `json:"Status"`
	SubscriberType string `json:"SubscriberType"`
}
