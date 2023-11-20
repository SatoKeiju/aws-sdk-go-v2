// Code generated by smithy-go-codegen DO NOT EDIT.

package neptunedata

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptunedata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about a specified model transform job. See Use a trained model
// to generate new model artifacts (https://docs.aws.amazon.com/neptune/latest/userguide/machine-learning-model-transform.html)
// . When invoking this operation in a Neptune cluster that has IAM authentication
// enabled, the IAM user or role making the request must have a policy attached
// that allows the neptune-db:GetMLModelTransformJobStatus (https://docs.aws.amazon.com/neptune/latest/userguide/iam-dp-actions.html#getmlmodeltransformjobstatus)
// IAM action in that cluster.
func (c *Client) GetMLModelTransformJob(ctx context.Context, params *GetMLModelTransformJobInput, optFns ...func(*Options)) (*GetMLModelTransformJobOutput, error) {
	if params == nil {
		params = &GetMLModelTransformJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMLModelTransformJob", params, optFns, c.addOperationGetMLModelTransformJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMLModelTransformJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMLModelTransformJobInput struct {

	// The unique identifier of the model-transform job to be reetrieved.
	//
	// This member is required.
	Id *string

	// The ARN of an IAM role that provides Neptune access to SageMaker and Amazon S3
	// resources. This must be listed in your DB cluster parameter group or an error
	// will occur.
	NeptuneIamRoleArn *string

	noSmithyDocumentSerde
}

type GetMLModelTransformJobOutput struct {

	// The base data processing job.
	BaseProcessingJob *types.MlResourceDefinition

	// The unique identifier of the model-transform job to be retrieved.
	Id *string

	// A list of the configuration information for the models being used.
	Models []types.MlConfigDefinition

	// The remote model transform job.
	RemoteModelTransformJob *types.MlResourceDefinition

	// The status of the model-transform job.
	Status *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMLModelTransformJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMLModelTransformJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMLModelTransformJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMLModelTransformJob"); err != nil {
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
	if err = addOpGetMLModelTransformJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMLModelTransformJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMLModelTransformJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMLModelTransformJob",
	}
}
