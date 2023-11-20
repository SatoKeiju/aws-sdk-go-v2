// Code generated by smithy-go-codegen DO NOT EDIT.

package omics

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/omics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates an annotation store.
func (c *Client) UpdateAnnotationStore(ctx context.Context, params *UpdateAnnotationStoreInput, optFns ...func(*Options)) (*UpdateAnnotationStoreOutput, error) {
	if params == nil {
		params = &UpdateAnnotationStoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAnnotationStore", params, optFns, c.addOperationUpdateAnnotationStoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAnnotationStoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAnnotationStoreInput struct {

	// A name for the store.
	//
	// This member is required.
	Name *string

	// A description for the store.
	Description *string

	noSmithyDocumentSerde
}

type UpdateAnnotationStoreOutput struct {

	// When the store was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The store's description.
	//
	// This member is required.
	Description *string

	// The store's ID.
	//
	// This member is required.
	Id *string

	// The store's name.
	//
	// This member is required.
	Name *string

	// The store's genome reference.
	//
	// This member is required.
	Reference types.ReferenceItem

	// The store's status.
	//
	// This member is required.
	Status types.StoreStatus

	// When the store was updated.
	//
	// This member is required.
	UpdateTime *time.Time

	// The annotation file format of the store.
	StoreFormat types.StoreFormat

	// Parsing options for the store.
	StoreOptions types.StoreOptions

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateAnnotationStoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAnnotationStore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAnnotationStore{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateAnnotationStore"); err != nil {
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
	if err = addEndpointPrefix_opUpdateAnnotationStoreMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateAnnotationStoreValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAnnotationStore(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opUpdateAnnotationStoreMiddleware struct {
}

func (*endpointPrefix_opUpdateAnnotationStoreMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opUpdateAnnotationStoreMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "analytics-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opUpdateAnnotationStoreMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opUpdateAnnotationStoreMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opUpdateAnnotationStore(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateAnnotationStore",
	}
}
