package exceptions

import (
	"errors"
)

var ErrSystem = errors.New("system error")
var ErrNotFound = errors.New("error not found")
var ErrBadRequest = errors.New("bad request")
var ErrEligibility = errors.New("error not eligible")
var ErrInvalidPayload = errors.New("invalid payload")

// GLOBAL
var ErrPayload = errors.New("Invalid request payload")
var ErrInternalServerError = errors.New("Error - 500 Internal Server Error - An error occurred in the system. Please try again later.")
var ErrUnAuthorized = errors.New("401 Unauthorized")
var ErrTimeOut = errors.New("Error - timeout to call backend service, please try again later.")
var ErrDeviceID = errors.New("sorry, your device id was not found")
var ErrPaymentCredit = errors.New("orderid can't be used anymore")
var ErrOrderCredit = errors.New("order not found")