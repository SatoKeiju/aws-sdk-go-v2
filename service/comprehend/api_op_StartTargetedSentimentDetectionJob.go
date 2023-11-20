// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts an asynchronous targeted sentiment detection job for a collection of
// documents. Use the DescribeTargetedSentimentDetectionJob operation to track the
// status of a job.
func (c *Client) StartTargetedSentimentDetectionJob(ctx context.Context, params *StartTargetedSentimentDetectionJobInput, optFns ...func(*Options)) (*StartTargetedSentimentDetectionJobOutput, error) {
	if params == nil {
		params = &StartTargetedSentimentDetectionJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartTargetedSentimentDetectionJob", params, optFns, c.addOperationStartTargetedSentimentDetectionJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartTargetedSentimentDetectionJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartTargetedSentimentDetectionJobInput struct {

	// The Amazon Resource Name (ARN) of the IAM role that grants Amazon Comprehend
	// read access to your input data. For more information, see Role-based permissions (https://docs.aws.amazon.com/comprehend/latest/dg/security_iam_id-based-policy-examples.html#auth-role-permissions)
	// .
	//
	// This member is required.
	DataAccessRoleArn *string

	// The input properties for an inference job. The document reader config field
	// applies only to non-text inputs for custom analysis.
	//
	// This member is required.
	InputDataConfig *types.InputDataConfig

	// The language of the input documents. Currently, English is the only supported
	// language.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// Specifies where to send the output files.
	//
	// This member is required.
	OutputDataConfig *types.OutputDataConfig

	// A unique identifier for the request. If you don't set the client request token,
	// Amazon Comprehend generates one.
	ClientRequestToken *string

	// The identifier of the job.
	JobName *string

	// Tags to associate with the targeted sentiment detection job. A tag is a
	// key-value pair that adds metadata to a resource used by Amazon Comprehend. For
	// example, a tag with "Sales" as the key might be added to a resource to indicate
	// its use by the sales department.
	Tags []types.Tag

	// ID for the KMS key that Amazon Comprehend uses to encrypt data on the storage
	// volume attached to the ML compute instance(s) that process the analysis job. The
	// VolumeKmsKeyId can be either of the following formats:
	//   - KMS Key ID: "1234abcd-12ab-34cd-56ef-1234567890ab"
	//   - Amazon Resource Name (ARN) of a KMS Key:
	//   "arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"
	VolumeKmsKeyId *string

	// Configuration parameters for an optional private Virtual Private Cloud (VPC)
	// containing the resources you are using for the job. For more information, see
	// Amazon VPC (https://docs.aws.amazon.com/vpc/latest/userguide/what-is-amazon-vpc.html)
	// .
	VpcConfig *types.VpcConfig

	noSmithyDocumentSerde
}

type StartTargetedSentimentDetectionJobOutput struct {

	// The Amazon Resource Name (ARN) of the targeted sentiment detection job. It is a
	// unique, fully qualified identifier for the job. It includes the Amazon Web
	// Services account, Amazon Web Services Region, and the job ID. The format of the
	// ARN is as follows: arn::comprehend:::targeted-sentiment-detection-job/ The
	// following is an example job ARN:
	// arn:aws:comprehend:us-west-2:111122223333:targeted-sentiment-detection-job/1234abcd12ab34cd56ef1234567890ab
	JobArn *string

	// The identifier generated for the job. To get the status of a job, use this
	// identifier with the DescribeTargetedSentimentDetectionJob operation.
	JobId *string

	// The status of the job.
	//   - SUBMITTED - The job has been received and is queued for processing.
	//   - IN_PROGRESS - Amazon Comprehend is processing the job.
	//   - COMPLETED - The job was successfully completed and the output is available.
	//   - FAILED - The job did not complete. To get details, use the
	//   DescribeTargetedSentimentDetectionJob operation.
	JobStatus types.JobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartTargetedSentimentDetectionJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartTargetedSentimentDetectionJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartTargetedSentimentDetectionJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartTargetedSentimentDetectionJob"); err != nil {
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
	if err = addIdempotencyToken_opStartTargetedSentimentDetectionJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartTargetedSentimentDetectionJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartTargetedSentimentDetectionJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartTargetedSentimentDetectionJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartTargetedSentimentDetectionJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartTargetedSentimentDetectionJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartTargetedSentimentDetectionJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartTargetedSentimentDetectionJobInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opStartTargetedSentimentDetectionJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartTargetedSentimentDetectionJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartTargetedSentimentDetectionJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartTargetedSentimentDetectionJob",
	}
}
