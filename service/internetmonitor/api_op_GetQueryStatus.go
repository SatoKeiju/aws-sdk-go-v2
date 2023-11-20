// Code generated by smithy-go-codegen DO NOT EDIT.

package internetmonitor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/internetmonitor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the current status of a query for the Amazon CloudWatch Internet
// Monitor query interface, for a specified query ID and monitor. When you run a
// query, check the status to make sure that the query has SUCCEEDED before you
// review the results.
//   - QUEUED : The query is scheduled to run.
//   - RUNNING : The query is in progress but not complete.
//   - SUCCEEDED : The query completed sucessfully.
//   - FAILED : The query failed due to an error.
//   - CANCELED : The query was canceled.
func (c *Client) GetQueryStatus(ctx context.Context, params *GetQueryStatusInput, optFns ...func(*Options)) (*GetQueryStatusOutput, error) {
	if params == nil {
		params = &GetQueryStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetQueryStatus", params, optFns, c.addOperationGetQueryStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetQueryStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetQueryStatusInput struct {

	// The name of the monitor.
	//
	// This member is required.
	MonitorName *string

	// The ID of the query that you want to return the status for. A QueryId is an
	// internally-generated dentifier for a specific query.
	//
	// This member is required.
	QueryId *string

	noSmithyDocumentSerde
}

type GetQueryStatusOutput struct {

	// The current status for a query.
	//
	// This member is required.
	Status types.QueryStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetQueryStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetQueryStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetQueryStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetQueryStatus"); err != nil {
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
	if err = addOpGetQueryStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetQueryStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetQueryStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetQueryStatus",
	}
}
