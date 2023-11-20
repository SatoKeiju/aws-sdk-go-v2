// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of distributions that have a cache behavior that's associated with
// the specified real-time log configuration. You can specify the real-time log
// configuration by its name or its Amazon Resource Name (ARN). You must provide at
// least one. If you provide both, CloudFront uses the name to identify the
// real-time log configuration to list distributions for. You can optionally
// specify the maximum number of items to receive in the response. If the total
// number of items in the list exceeds the maximum that you specify, or the default
// maximum, the response is paginated. To get the next page of items, send a
// subsequent request that specifies the NextMarker value from the current
// response as the Marker value in the subsequent request.
func (c *Client) ListDistributionsByRealtimeLogConfig(ctx context.Context, params *ListDistributionsByRealtimeLogConfigInput, optFns ...func(*Options)) (*ListDistributionsByRealtimeLogConfigOutput, error) {
	if params == nil {
		params = &ListDistributionsByRealtimeLogConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDistributionsByRealtimeLogConfig", params, optFns, c.addOperationListDistributionsByRealtimeLogConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDistributionsByRealtimeLogConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDistributionsByRealtimeLogConfigInput struct {

	// Use this field when paginating results to indicate where to begin in your list
	// of distributions. The response includes distributions in the list that occur
	// after the marker. To get the next page of the list, set this field's value to
	// the value of NextMarker from the current page's response.
	Marker *string

	// The maximum number of distributions that you want in the response.
	MaxItems *int32

	// The Amazon Resource Name (ARN) of the real-time log configuration whose
	// associated distributions you want to list.
	RealtimeLogConfigArn *string

	// The name of the real-time log configuration whose associated distributions you
	// want to list.
	RealtimeLogConfigName *string

	noSmithyDocumentSerde
}

type ListDistributionsByRealtimeLogConfigOutput struct {

	// A distribution list.
	DistributionList *types.DistributionList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDistributionsByRealtimeLogConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpListDistributionsByRealtimeLogConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListDistributionsByRealtimeLogConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListDistributionsByRealtimeLogConfig"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDistributionsByRealtimeLogConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListDistributionsByRealtimeLogConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListDistributionsByRealtimeLogConfig",
	}
}
