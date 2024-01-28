package exceptions

import (
	"net/http"
	"os"
)

func MapToHttpStatusCode(err error) int {
	var statusCode = http.StatusInternalServerError
	switch err {
	case ErrSystem:
		statusCode = http.StatusInternalServerError
	case ErrNotFound:
		statusCode = http.StatusNotFound
	case ErrBadRequest:
		statusCode = http.StatusBadRequest
	case ErrPayload:
		statusCode = http.StatusBadRequest
	case ErrInvalidPayload:
		statusCode = http.StatusBadRequest
	case ErrEligibility:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}

	return statusCode
}

func MapToErrorMessage(err error) string {
	var errorMessage = os.Getenv("ERRSYSTEM_TEXT")

	switch err {
	case ErrSystem:
		errorMessage = os.Getenv("ERRSYSTEM_TEXT")
	case ErrNotFound:
		errorMessage = os.Getenv("ERRNOTFOUND_TEXT")
	case ErrBadRequest:
		errorMessage = os.Getenv("ERRNOTFOUND_TEXT")
	case ErrEligibility:
		errorMessage = os.Getenv("ERRELIGIBILITY_TEXT")
	case ErrInvalidPayload:
		errorMessage = os.Getenv("ERRPAYLOAD_TEXT")
	default:
		errorMessage = os.Getenv("ERRSYSTEM_TEXT")
	}

	return errorMessage
}

func MapToErrorType(reason string) error {
	var errorType = ErrSystem

	switch reason {
	case "Internal Application Error from RuleValidation - Could not get the Eligible Product Information":
		errorType = ErrEligibility
	case "Invalid Data":
		errorType = ErrInvalidPayload
	default:
		errorType = ErrSystem
	}

	return errorType
}

func MapToStatusMessage(statuscode int) string {
	var statusMessage = "Voucher invalid"

	switch statuscode {
	case 0:
		statusMessage = "Disable / not injected yet"
	case 1:
		statusMessage = "Enable / has been injected and ready to redeem"
	case 3:
		statusMessage = "Used / has been redeemed"
	case 4:
		statusMessage = "Blocked"
	case 5:
		statusMessage = "Expired"
	}

	return statusMessage
}

func MapRedeemDescription(statuscode string) string {
	var statusMessage = "Error"

	switch statuscode {
	case "00":
		statusMessage = "Success"
	case "14":
		statusMessage = "MsisdnNotExist"
	case "79":
		statusMessage = "MsisdnBlocked"
	case "81":
		statusMessage = "MsisdnExpire"
	case "70":
		statusMessage = "VoucherOutOfStock"
	case "05":
		statusMessage = "UndefinedError"
	case "13":
		statusMessage = "InvalidVoucherNominal"
	case "69":
		statusMessage = "LateResponse"
	case "91":
		statusMessage = "DatabaseProblem"
	case "92":
		statusMessage = "UnableToRouteTrx"
	case "12":
		statusMessage = "ReversalDenied"
	case "63":
		statusMessage = "InvalidReversal"
	case "94":
		statusMessage = "DuplicateReversal"
	case "15":
		statusMessage = "VoucherAlreadyUsed"
	case "21":
		statusMessage = "VoucherExpire"
	case "22":
		statusMessage = "VoucherBlocked"
	case "20":
		statusMessage = "VoucherDisable"
	case "78":
		statusMessage = "MsisdnSuspended"
	case "11":
		statusMessage = "VDSuccess"
	case "33":
		statusMessage = "VDRestricted"
	case "44":
		statusMessage = "VDTimeout"
	case "41":
		statusMessage = "InvalidVoucherControl"
	case "50":
		statusMessage = "AreaRestricted"
	case "95":
		statusMessage = "MsisdnMaxBalance"
	case "96":
		statusMessage = "MsisdnNotAllowed"
	}

	return statusMessage
}
