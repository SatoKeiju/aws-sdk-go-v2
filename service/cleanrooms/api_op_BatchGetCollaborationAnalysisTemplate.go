// Code generated by smithy-go-codegen DO NOT EDIT.

package cleanrooms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cleanrooms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves multiple analysis templates within a collaboration by their Amazon
// Resource Names (ARNs).
func (c *Client) BatchGetCollaborationAnalysisTemplate(ctx context.Context, params *BatchGetCollaborationAnalysisTemplateInput, optFns ...func(*Options)) (*BatchGetCollaborationAnalysisTemplateOutput, error) {
	if params == nil {
		params = &BatchGetCollaborationAnalysisTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetCollaborationAnalysisTemplate", params, optFns, c.addOperationBatchGetCollaborationAnalysisTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetCollaborationAnalysisTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetCollaborationAnalysisTemplateInput struct {

	// The Amazon Resource Name (ARN) associated with the analysis template within a
	// collaboration.
	//
	// This member is required.
	AnalysisTemplateArns []string

	// A unique identifier for the collaboration that the analysis templates belong
	// to. Currently accepts collaboration ID.
	//
	// This member is required.
	CollaborationIdentifier *string

	noSmithyDocumentSerde
}

type BatchGetCollaborationAnalysisTemplateOutput struct {

	// The retrieved list of analysis templates within a collaboration.
	//
	// This member is required.
	CollaborationAnalysisTemplates []types.CollaborationAnalysisTemplate

	// Error reasons for collaboration analysis templates that could not be retrieved.
	// One error is returned for every collaboration analysis template that could not
	// be retrieved.
	//
	// This member is required.
	Errors []types.BatchGetCollaborationAnalysisTemplateError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetCollaborationAnalysisTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchGetCollaborationAnalysisTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchGetCollaborationAnalysisTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchGetCollaborationAnalysisTemplate"); err != nil {
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
	if err = addOpBatchGetCollaborationAnalysisTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetCollaborationAnalysisTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetCollaborationAnalysisTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchGetCollaborationAnalysisTemplate",
	}
}
