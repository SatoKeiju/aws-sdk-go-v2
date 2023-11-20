// Code generated by smithy-go-codegen DO NOT EDIT.

package resiliencehub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Imports resources to Resilience Hub application draft version from different
// input sources. For more information about the input sources supported by
// Resilience Hub, see Discover the structure and describe your Resilience Hub
// application (https://docs.aws.amazon.com/resilience-hub/latest/userguide/discover-structure.html)
// .
func (c *Client) ImportResourcesToDraftAppVersion(ctx context.Context, params *ImportResourcesToDraftAppVersionInput, optFns ...func(*Options)) (*ImportResourcesToDraftAppVersionOutput, error) {
	if params == nil {
		params = &ImportResourcesToDraftAppVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportResourcesToDraftAppVersion", params, optFns, c.addOperationImportResourcesToDraftAppVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportResourcesToDraftAppVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportResourcesToDraftAppVersionInput struct {

	// Amazon Resource Name (ARN) of the Resilience Hub application. The format for
	// this ARN is: arn: partition :resiliencehub: region : account :app/ app-id . For
	// more information about ARNs, see Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference guide.
	//
	// This member is required.
	AppArn *string

	// The input sources of the Amazon Elastic Kubernetes Service resources you need
	// to import.
	EksSources []types.EksSource

	// The import strategy you would like to set to import resources into Resilience
	// Hub application.
	ImportStrategy types.ResourceImportStrategyType

	// The Amazon Resource Names (ARNs) for the resources.
	SourceArns []string

	// A list of terraform file s3 URLs you need to import.
	TerraformSources []types.TerraformSource

	noSmithyDocumentSerde
}

type ImportResourcesToDraftAppVersionOutput struct {

	// Amazon Resource Name (ARN) of the Resilience Hub application. The format for
	// this ARN is: arn: partition :resiliencehub: region : account :app/ app-id . For
	// more information about ARNs, see Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference guide.
	//
	// This member is required.
	AppArn *string

	// The version of the application.
	//
	// This member is required.
	AppVersion *string

	// Status of the action.
	//
	// This member is required.
	Status types.ResourceImportStatusType

	// The input sources of the Amazon Elastic Kubernetes Service resources you have
	// imported.
	EksSources []types.EksSource

	// The Amazon Resource Names (ARNs) for the resources you have imported.
	SourceArns []string

	// A list of terraform file s3 URLs you have imported.
	TerraformSources []types.TerraformSource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportResourcesToDraftAppVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpImportResourcesToDraftAppVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpImportResourcesToDraftAppVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ImportResourcesToDraftAppVersion"); err != nil {
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
	if err = addOpImportResourcesToDraftAppVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportResourcesToDraftAppVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opImportResourcesToDraftAppVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ImportResourcesToDraftAppVersion",
	}
}
