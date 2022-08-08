// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package storagegateway

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeInternalServerError for service response error code
	// "InternalServerError".
	//
	// An internal server error has occurred during the request. For more information,
	// see the error and message fields.
	ErrCodeInternalServerError = "InternalServerError"

	// ErrCodeInvalidGatewayRequestException for service response error code
	// "InvalidGatewayRequestException".
	//
	// An exception occurred because an invalid gateway request was issued to the
	// service. For more information, see the error and message fields.
	ErrCodeInvalidGatewayRequestException = "InvalidGatewayRequestException"

	// ErrCodeServiceUnavailableError for service response error code
	// "ServiceUnavailableError".
	//
	// An internal server error has occurred because the service is unavailable.
	// For more information, see the error and message fields.
	ErrCodeServiceUnavailableError = "ServiceUnavailableError"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"InternalServerError":            newErrorInternalServerError,
	"InvalidGatewayRequestException": newErrorInvalidGatewayRequestException,
	"ServiceUnavailableError":        newErrorServiceUnavailableError,
}