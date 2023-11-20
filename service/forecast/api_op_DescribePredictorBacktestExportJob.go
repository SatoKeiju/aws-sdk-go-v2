// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes a predictor backtest export job created using the
// CreatePredictorBacktestExportJob operation. In addition to listing the
// properties provided by the user in the CreatePredictorBacktestExportJob
// request, this operation lists the following properties:
//   - CreationTime
//   - LastModificationTime
//   - Status
//   - Message (if an error occurred)
func (c *Client) DescribePredictorBacktestExportJob(ctx context.Context, params *DescribePredictorBacktestExportJobInput, optFns ...func(*Options)) (*DescribePredictorBacktestExportJobOutput, error) {
	if params == nil {
		params = &DescribePredictorBacktestExportJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePredictorBacktestExportJob", params, optFns, c.addOperationDescribePredictorBacktestExportJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePredictorBacktestExportJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePredictorBacktestExportJobInput struct {

	// The Amazon Resource Name (ARN) of the predictor backtest export job.
	//
	// This member is required.
	PredictorBacktestExportJobArn *string

	noSmithyDocumentSerde
}

type DescribePredictorBacktestExportJobOutput struct {

	// When the predictor backtest export job was created.
	CreationTime *time.Time

	// The destination for an export job. Provide an S3 path, an Identity and Access
	// Management (IAM) role that allows Amazon Forecast to access the location, and an
	// Key Management Service (KMS) key (optional).
	Destination *types.DataDestination

	// The format of the exported data, CSV or PARQUET.
	Format *string

	// The last time the resource was modified. The timestamp depends on the status of
	// the job:
	//   - CREATE_PENDING - The CreationTime .
	//   - CREATE_IN_PROGRESS - The current timestamp.
	//   - CREATE_STOPPING - The current timestamp.
	//   - CREATE_STOPPED - When the job stopped.
	//   - ACTIVE or CREATE_FAILED - When the job finished or failed.
	LastModificationTime *time.Time

	// Information about any errors that may have occurred during the backtest export.
	Message *string

	// The Amazon Resource Name (ARN) of the predictor.
	PredictorArn *string

	// The Amazon Resource Name (ARN) of the predictor backtest export job.
	PredictorBacktestExportJobArn *string

	// The name of the predictor backtest export job.
	PredictorBacktestExportJobName *string

	// The status of the predictor backtest export job. States include:
	//   - ACTIVE
	//   - CREATE_PENDING , CREATE_IN_PROGRESS , CREATE_FAILED
	//   - CREATE_STOPPING , CREATE_STOPPED
	//   - DELETE_PENDING , DELETE_IN_PROGRESS , DELETE_FAILED
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePredictorBacktestExportJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePredictorBacktestExportJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePredictorBacktestExportJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribePredictorBacktestExportJob"); err != nil {
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
	if err = addOpDescribePredictorBacktestExportJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePredictorBacktestExportJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribePredictorBacktestExportJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribePredictorBacktestExportJob",
	}
}
