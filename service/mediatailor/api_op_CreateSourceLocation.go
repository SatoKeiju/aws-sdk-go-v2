// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a source location. A source location is a container for sources. For
// more information about source locations, see Working with source locations (https://docs.aws.amazon.com/mediatailor/latest/ug/channel-assembly-source-locations.html)
// in the MediaTailor User Guide.
func (c *Client) CreateSourceLocation(ctx context.Context, params *CreateSourceLocationInput, optFns ...func(*Options)) (*CreateSourceLocationOutput, error) {
	if params == nil {
		params = &CreateSourceLocationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSourceLocation", params, optFns, c.addOperationCreateSourceLocationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSourceLocationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSourceLocationInput struct {

	// The source's HTTP package configurations.
	//
	// This member is required.
	HttpConfiguration *types.HttpConfiguration

	// The name associated with the source location.
	//
	// This member is required.
	SourceLocationName *string

	// Access configuration parameters. Configures the type of authentication used to
	// access content from your source location.
	AccessConfiguration *types.AccessConfiguration

	// The optional configuration for the server that serves segments.
	DefaultSegmentDeliveryConfiguration *types.DefaultSegmentDeliveryConfiguration

	// A list of the segment delivery configurations associated with this resource.
	SegmentDeliveryConfigurations []types.SegmentDeliveryConfiguration

	// The tags to assign to the source location. Tags are key-value pairs that you
	// can associate with Amazon resources to help with organization, access control,
	// and cost tracking. For more information, see Tagging AWS Elemental MediaTailor
	// Resources (https://docs.aws.amazon.com/mediatailor/latest/ug/tagging.html) .
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateSourceLocationOutput struct {

	// Access configuration parameters. Configures the type of authentication used to
	// access content from your source location.
	AccessConfiguration *types.AccessConfiguration

	// The ARN to assign to the source location.
	Arn *string

	// The time the source location was created.
	CreationTime *time.Time

	// The optional configuration for the server that serves segments.
	DefaultSegmentDeliveryConfiguration *types.DefaultSegmentDeliveryConfiguration

	// The source's HTTP package configurations.
	HttpConfiguration *types.HttpConfiguration

	// The time the source location was last modified.
	LastModifiedTime *time.Time

	// The segment delivery configurations for the source location. For information
	// about MediaTailor configurations, see Working with configurations in AWS
	// Elemental MediaTailor (https://docs.aws.amazon.com/mediatailor/latest/ug/configurations.html)
	// .
	SegmentDeliveryConfigurations []types.SegmentDeliveryConfiguration

	// The name to assign to the source location.
	SourceLocationName *string

	// The tags to assign to the source location. Tags are key-value pairs that you
	// can associate with Amazon resources to help with organization, access control,
	// and cost tracking. For more information, see Tagging AWS Elemental MediaTailor
	// Resources (https://docs.aws.amazon.com/mediatailor/latest/ug/tagging.html) .
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSourceLocationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSourceLocation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSourceLocation{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateSourceLocation"); err != nil {
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
	if err = addOpCreateSourceLocationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSourceLocation(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSourceLocation(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateSourceLocation",
	}
}
