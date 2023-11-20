// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the results for a given media analysis job. Takes a JobId returned by
// StartMediaAnalysisJob.
func (c *Client) GetMediaAnalysisJob(ctx context.Context, params *GetMediaAnalysisJobInput, optFns ...func(*Options)) (*GetMediaAnalysisJobOutput, error) {
	if params == nil {
		params = &GetMediaAnalysisJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMediaAnalysisJob", params, optFns, c.addOperationGetMediaAnalysisJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMediaAnalysisJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMediaAnalysisJobInput struct {

	// Unique identifier for the media analysis job for which you want to retrieve
	// results.
	//
	// This member is required.
	JobId *string

	noSmithyDocumentSerde
}

type GetMediaAnalysisJobOutput struct {

	// The Unix date and time when the job was started.
	//
	// This member is required.
	CreationTimestamp *time.Time

	// Reference to the input manifest that was provided in the job creation request.
	//
	// This member is required.
	Input *types.MediaAnalysisInput

	// The identifier for the media analysis job.
	//
	// This member is required.
	JobId *string

	// Operation configurations that were provided during job creation.
	//
	// This member is required.
	OperationsConfig *types.MediaAnalysisOperationsConfig

	// Output configuration that was provided in the creation request.
	//
	// This member is required.
	OutputConfig *types.MediaAnalysisOutputConfig

	// The current status of the media analysis job.
	//
	// This member is required.
	Status types.MediaAnalysisJobStatus

	// The Unix date and time when the job finished.
	CompletionTimestamp *time.Time

	// Details about the error that resulted in failure of the job.
	FailureDetails *types.MediaAnalysisJobFailureDetails

	// The name of the media analysis job.
	JobName *string

	// KMS Key that was provided in the creation request.
	KmsKeyId *string

	// The summary manifest provides statistics on input manifest and errors
	// identified in the input manifest.
	ManifestSummary *types.MediaAnalysisManifestSummary

	// Output manifest that contains prediction results.
	Results *types.MediaAnalysisResults

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMediaAnalysisJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetMediaAnalysisJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMediaAnalysisJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMediaAnalysisJob"); err != nil {
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
	if err = addOpGetMediaAnalysisJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMediaAnalysisJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMediaAnalysisJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMediaAnalysisJob",
	}
}
