// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets an endpoint for a specified stream for either reading or writing. Use this
// endpoint in your application to read from the specified stream (using the
// GetMedia or GetMediaForFragmentList operations) or write to it (using the
// PutMedia operation). The returned endpoint does not have the API name appended.
// The client needs to add the API name to the returned endpoint. In the request,
// specify the stream either by StreamName or StreamARN .
func (c *Client) GetDataEndpoint(ctx context.Context, params *GetDataEndpointInput, optFns ...func(*Options)) (*GetDataEndpointOutput, error) {
	if params == nil {
		params = &GetDataEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDataEndpoint", params, optFns, c.addOperationGetDataEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDataEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDataEndpointInput struct {

	// The name of the API action for which to get an endpoint.
	//
	// This member is required.
	APIName types.APIName

	// The Amazon Resource Name (ARN) of the stream that you want to get the endpoint
	// for. You must specify either this parameter or a StreamName in the request.
	StreamARN *string

	// The name of the stream that you want to get the endpoint for. You must specify
	// either this parameter or a StreamARN in the request.
	StreamName *string

	noSmithyDocumentSerde
}

type GetDataEndpointOutput struct {

	// The endpoint value. To read data from the stream or to write data to it,
	// specify this endpoint in your application.
	DataEndpoint *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDataEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDataEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDataEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetDataEndpoint"); err != nil {
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
	if err = addOpGetDataEndpointValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDataEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDataEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDataEndpoint",
	}
}
