// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the map style descriptor from a map resource. The style descriptor
// contains speciﬁcations on how features render on a map. For example, what data
// to display, what order to display the data in, and the style for the data. Style
// descriptors follow the Mapbox Style Specification.
func (c *Client) GetMapStyleDescriptor(ctx context.Context, params *GetMapStyleDescriptorInput, optFns ...func(*Options)) (*GetMapStyleDescriptorOutput, error) {
	if params == nil {
		params = &GetMapStyleDescriptorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMapStyleDescriptor", params, optFns, c.addOperationGetMapStyleDescriptorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMapStyleDescriptorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMapStyleDescriptorInput struct {

	// The map resource to retrieve the style descriptor from.
	//
	// This member is required.
	MapName *string

	// The optional API key (https://docs.aws.amazon.com/location/latest/developerguide/using-apikeys.html)
	// to authorize the request.
	Key *string

	noSmithyDocumentSerde
}

type GetMapStyleDescriptorOutput struct {

	// Contains the body of the style descriptor.
	Blob []byte

	// The HTTP Cache-Control directive for the value.
	CacheControl *string

	// The style descriptor's content type. For example, application/json .
	ContentType *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMapStyleDescriptorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMapStyleDescriptor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMapStyleDescriptor{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMapStyleDescriptor"); err != nil {
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
	if err = addEndpointPrefix_opGetMapStyleDescriptorMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetMapStyleDescriptorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMapStyleDescriptor(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetMapStyleDescriptorMiddleware struct {
}

func (*endpointPrefix_opGetMapStyleDescriptorMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetMapStyleDescriptorMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "maps." + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opGetMapStyleDescriptorMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetMapStyleDescriptorMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opGetMapStyleDescriptor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMapStyleDescriptor",
	}
}
