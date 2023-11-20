// Code generated by smithy-go-codegen DO NOT EDIT.

package medicalimaging

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medicalimaging/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Get image set properties.
func (c *Client) GetImageSet(ctx context.Context, params *GetImageSetInput, optFns ...func(*Options)) (*GetImageSetOutput, error) {
	if params == nil {
		params = &GetImageSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetImageSet", params, optFns, c.addOperationGetImageSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetImageSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetImageSetInput struct {

	// The data store identifier.
	//
	// This member is required.
	DatastoreId *string

	// The image set identifier.
	//
	// This member is required.
	ImageSetId *string

	// The image set version identifier.
	VersionId *string

	noSmithyDocumentSerde
}

type GetImageSetOutput struct {

	// The data store identifier.
	//
	// This member is required.
	DatastoreId *string

	// The image set identifier.
	//
	// This member is required.
	ImageSetId *string

	// The image set state.
	//
	// This member is required.
	ImageSetState types.ImageSetState

	// The image set version identifier.
	//
	// This member is required.
	VersionId *string

	// The timestamp when image set properties were created.
	CreatedAt *time.Time

	// The timestamp when the image set properties were deleted.
	DeletedAt *time.Time

	// The Amazon Resource Name (ARN) assigned to the image set.
	ImageSetArn *string

	// The image set workflow status.
	ImageSetWorkflowStatus types.ImageSetWorkflowStatus

	// The error message thrown if an image set action fails.
	Message *string

	// The timestamp when image set properties were updated.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetImageSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetImageSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetImageSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetImageSet"); err != nil {
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
	if err = addEndpointPrefix_opGetImageSetMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetImageSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetImageSet(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetImageSetMiddleware struct {
}

func (*endpointPrefix_opGetImageSetMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetImageSetMiddleware) HandleFinalize(ctx context.Context, in middleware.FinalizeInput, next middleware.FinalizeHandler) (
	out middleware.FinalizeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleFinalize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "runtime-" + req.URL.Host

	return next.HandleFinalize(ctx, in)
}
func addEndpointPrefix_opGetImageSetMiddleware(stack *middleware.Stack) error {
	return stack.Finalize.Insert(&endpointPrefix_opGetImageSetMiddleware{}, "ResolveEndpointV2", middleware.After)
}

func newServiceMetadataMiddleware_opGetImageSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetImageSet",
	}
}
