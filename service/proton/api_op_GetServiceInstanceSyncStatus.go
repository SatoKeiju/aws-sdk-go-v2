// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get the status of the synced service instance.
func (c *Client) GetServiceInstanceSyncStatus(ctx context.Context, params *GetServiceInstanceSyncStatusInput, optFns ...func(*Options)) (*GetServiceInstanceSyncStatusOutput, error) {
	if params == nil {
		params = &GetServiceInstanceSyncStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetServiceInstanceSyncStatus", params, optFns, c.addOperationGetServiceInstanceSyncStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetServiceInstanceSyncStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetServiceInstanceSyncStatusInput struct {

	// The name of the service instance that you want the sync status input for.
	//
	// This member is required.
	ServiceInstanceName *string

	// The name of the service that the service instance belongs to.
	//
	// This member is required.
	ServiceName *string

	noSmithyDocumentSerde
}

type GetServiceInstanceSyncStatusOutput struct {

	// The service instance sync desired state that's returned by Proton
	DesiredState *types.Revision

	// The detailed data of the latest successful sync with the service instance.
	LatestSuccessfulSync *types.ResourceSyncAttempt

	// The detailed data of the latest sync with the service instance.
	LatestSync *types.ResourceSyncAttempt

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetServiceInstanceSyncStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetServiceInstanceSyncStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetServiceInstanceSyncStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetServiceInstanceSyncStatus"); err != nil {
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
	if err = addOpGetServiceInstanceSyncStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetServiceInstanceSyncStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetServiceInstanceSyncStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetServiceInstanceSyncStatus",
	}
}
