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

// The VOD source configuration parameters.
func (c *Client) CreateVodSource(ctx context.Context, params *CreateVodSourceInput, optFns ...func(*Options)) (*CreateVodSourceOutput, error) {
	if params == nil {
		params = &CreateVodSourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVodSource", params, optFns, c.addOperationCreateVodSourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVodSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVodSourceInput struct {

	// A list of HTTP package configuration parameters for this VOD source.
	//
	// This member is required.
	HttpPackageConfigurations []types.HttpPackageConfiguration

	// The name of the source location for this VOD source.
	//
	// This member is required.
	SourceLocationName *string

	// The name associated with the VOD source.>
	//
	// This member is required.
	VodSourceName *string

	// The tags to assign to the VOD source. Tags are key-value pairs that you can
	// associate with Amazon resources to help with organization, access control, and
	// cost tracking. For more information, see Tagging AWS Elemental MediaTailor
	// Resources (https://docs.aws.amazon.com/mediatailor/latest/ug/tagging.html) .
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateVodSourceOutput struct {

	// The ARN to assign to this VOD source.
	Arn *string

	// The time the VOD source was created.
	CreationTime *time.Time

	// A list of HTTP package configuration parameters for this VOD source.
	HttpPackageConfigurations []types.HttpPackageConfiguration

	// The time the VOD source was last modified.
	LastModifiedTime *time.Time

	// The name to assign to the source location for this VOD source.
	SourceLocationName *string

	// The tags to assign to the VOD source. Tags are key-value pairs that you can
	// associate with Amazon resources to help with organization, access control, and
	// cost tracking. For more information, see Tagging AWS Elemental MediaTailor
	// Resources (https://docs.aws.amazon.com/mediatailor/latest/ug/tagging.html) .
	Tags map[string]string

	// The name to assign to the VOD source.
	VodSourceName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateVodSourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateVodSource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateVodSource{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateVodSource"); err != nil {
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
	if err = addOpCreateVodSourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVodSource(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateVodSource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateVodSource",
	}
}
