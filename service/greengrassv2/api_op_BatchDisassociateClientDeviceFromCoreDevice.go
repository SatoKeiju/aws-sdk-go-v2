// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrassv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrassv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates a list of client devices from a core device. After you
// disassociate a client device from a core device, the client device won't be able
// to use cloud discovery to retrieve the core device's connectivity information
// and certificates.
func (c *Client) BatchDisassociateClientDeviceFromCoreDevice(ctx context.Context, params *BatchDisassociateClientDeviceFromCoreDeviceInput, optFns ...func(*Options)) (*BatchDisassociateClientDeviceFromCoreDeviceOutput, error) {
	if params == nil {
		params = &BatchDisassociateClientDeviceFromCoreDeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDisassociateClientDeviceFromCoreDevice", params, optFns, c.addOperationBatchDisassociateClientDeviceFromCoreDeviceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDisassociateClientDeviceFromCoreDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDisassociateClientDeviceFromCoreDeviceInput struct {

	// The name of the core device. This is also the name of the IoT thing.
	//
	// This member is required.
	CoreDeviceThingName *string

	// The list of client devices to disassociate.
	Entries []types.DisassociateClientDeviceFromCoreDeviceEntry

	noSmithyDocumentSerde
}

type BatchDisassociateClientDeviceFromCoreDeviceOutput struct {

	// The list of any errors for the entries in the request. Each error entry
	// contains the name of the IoT thing that failed to disassociate.
	ErrorEntries []types.DisassociateClientDeviceFromCoreDeviceErrorEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchDisassociateClientDeviceFromCoreDeviceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchDisassociateClientDeviceFromCoreDevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchDisassociateClientDeviceFromCoreDevice{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchDisassociateClientDeviceFromCoreDevice"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addV4DetectSkewMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpBatchDisassociateClientDeviceFromCoreDeviceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDisassociateClientDeviceFromCoreDevice(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opBatchDisassociateClientDeviceFromCoreDevice(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchDisassociateClientDeviceFromCoreDevice",
	}
}
