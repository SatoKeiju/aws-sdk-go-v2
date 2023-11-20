// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Exports the contents of a Amazon Lex resource in a specified format.
func (c *Client) GetExport(ctx context.Context, params *GetExportInput, optFns ...func(*Options)) (*GetExportOutput, error) {
	if params == nil {
		params = &GetExportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetExport", params, optFns, c.addOperationGetExportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetExportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetExportInput struct {

	// The format of the exported data.
	//
	// This member is required.
	ExportType types.ExportType

	// The name of the bot to export.
	//
	// This member is required.
	Name *string

	// The type of resource to export.
	//
	// This member is required.
	ResourceType types.ResourceType

	// The version of the bot to export.
	//
	// This member is required.
	Version *string

	noSmithyDocumentSerde
}

type GetExportOutput struct {

	// The status of the export.
	//   - IN_PROGRESS - The export is in progress.
	//   - READY - The export is complete.
	//   - FAILED - The export could not be completed.
	ExportStatus types.ExportStatus

	// The format of the exported data.
	ExportType types.ExportType

	// If status is FAILED , Amazon Lex provides the reason that it failed to export
	// the resource.
	FailureReason *string

	// The name of the bot being exported.
	Name *string

	// The type of the exported resource.
	ResourceType types.ResourceType

	// An S3 pre-signed URL that provides the location of the exported resource. The
	// exported resource is a ZIP archive that contains the exported resource in JSON
	// format. The structure of the archive may change. Your code should not rely on
	// the archive structure.
	Url *string

	// The version of the bot being exported.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetExportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetExport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetExport{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetExport"); err != nil {
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
	if err = addOpGetExportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetExport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetExport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetExport",
	}
}
