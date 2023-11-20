// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This method works, but is deprecated. Use BatchGetDeploymentTargets instead.
// Returns an array of one or more instances associated with a deployment. This
// method works with EC2/On-premises and Lambda compute platforms. The newer
// BatchGetDeploymentTargets works with all compute platforms. The maximum number
// of instances that can be returned is 25.
//
// Deprecated: This operation is deprecated, use BatchGetDeploymentTargets instead.
func (c *Client) BatchGetDeploymentInstances(ctx context.Context, params *BatchGetDeploymentInstancesInput, optFns ...func(*Options)) (*BatchGetDeploymentInstancesOutput, error) {
	if params == nil {
		params = &BatchGetDeploymentInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetDeploymentInstances", params, optFns, c.addOperationBatchGetDeploymentInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetDeploymentInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a BatchGetDeploymentInstances operation.
type BatchGetDeploymentInstancesInput struct {

	// The unique ID of a deployment.
	//
	// This member is required.
	DeploymentId *string

	// The unique IDs of instances used in the deployment. The maximum number of
	// instance IDs you can specify is 25.
	//
	// This member is required.
	InstanceIds []string

	noSmithyDocumentSerde
}

// Represents the output of a BatchGetDeploymentInstances operation.
type BatchGetDeploymentInstancesOutput struct {

	// Information about errors that might have occurred during the API call.
	ErrorMessage *string

	// Information about the instance.
	InstancesSummary []types.InstanceSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetDeploymentInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetDeploymentInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetDeploymentInstances{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchGetDeploymentInstances"); err != nil {
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
	if err = addOpBatchGetDeploymentInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetDeploymentInstances(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetDeploymentInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchGetDeploymentInstances",
	}
}
