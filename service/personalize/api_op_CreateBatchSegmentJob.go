// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a batch segment job. The operation can handle up to 50 million records
// and the input file must be in JSON format. For more information, see Getting
// batch recommendations and user segments (https://docs.aws.amazon.com/personalize/latest/dg/recommendations-batch.html)
// .
func (c *Client) CreateBatchSegmentJob(ctx context.Context, params *CreateBatchSegmentJobInput, optFns ...func(*Options)) (*CreateBatchSegmentJobOutput, error) {
	if params == nil {
		params = &CreateBatchSegmentJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBatchSegmentJob", params, optFns, c.addOperationCreateBatchSegmentJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBatchSegmentJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBatchSegmentJobInput struct {

	// The Amazon S3 path for the input data used to generate the batch segment job.
	//
	// This member is required.
	JobInput *types.BatchSegmentJobInput

	// The name of the batch segment job to create.
	//
	// This member is required.
	JobName *string

	// The Amazon S3 path for the bucket where the job's output will be stored.
	//
	// This member is required.
	JobOutput *types.BatchSegmentJobOutput

	// The ARN of the Amazon Identity and Access Management role that has permissions
	// to read and write to your input and output Amazon S3 buckets respectively.
	//
	// This member is required.
	RoleArn *string

	// The Amazon Resource Name (ARN) of the solution version you want the batch
	// segment job to use to generate batch segments.
	//
	// This member is required.
	SolutionVersionArn *string

	// The ARN of the filter to apply to the batch segment job. For more information
	// on using filters, see Filtering batch recommendations (https://docs.aws.amazon.com/personalize/latest/dg/filter-batch.html)
	// .
	FilterArn *string

	// The number of predicted users generated by the batch segment job for each line
	// of input data. The maximum number of users per segment is 5 million.
	NumResults *int32

	// A list of tags (https://docs.aws.amazon.com/personalize/latest/dg/tagging-resources.html)
	// to apply to the batch segment job.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateBatchSegmentJobOutput struct {

	// The ARN of the batch segment job.
	BatchSegmentJobArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBatchSegmentJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateBatchSegmentJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateBatchSegmentJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateBatchSegmentJob"); err != nil {
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
	if err = addOpCreateBatchSegmentJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBatchSegmentJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateBatchSegmentJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateBatchSegmentJob",
	}
}
