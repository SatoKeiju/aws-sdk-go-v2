// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an edge deployment plan, consisting of multiple stages. Each stage may
// have a different deployment configuration and devices.
func (c *Client) CreateEdgeDeploymentPlan(ctx context.Context, params *CreateEdgeDeploymentPlanInput, optFns ...func(*Options)) (*CreateEdgeDeploymentPlanOutput, error) {
	if params == nil {
		params = &CreateEdgeDeploymentPlanInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEdgeDeploymentPlan", params, optFns, c.addOperationCreateEdgeDeploymentPlanMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEdgeDeploymentPlanOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEdgeDeploymentPlanInput struct {

	// The device fleet used for this edge deployment plan.
	//
	// This member is required.
	DeviceFleetName *string

	// The name of the edge deployment plan.
	//
	// This member is required.
	EdgeDeploymentPlanName *string

	// List of models associated with the edge deployment plan.
	//
	// This member is required.
	ModelConfigs []types.EdgeDeploymentModelConfig

	// List of stages of the edge deployment plan. The number of stages is limited to
	// 10 per deployment.
	Stages []types.DeploymentStage

	// List of tags with which to tag the edge deployment plan.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateEdgeDeploymentPlanOutput struct {

	// The ARN of the edge deployment plan.
	//
	// This member is required.
	EdgeDeploymentPlanArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEdgeDeploymentPlanMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateEdgeDeploymentPlan{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateEdgeDeploymentPlan{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateEdgeDeploymentPlan"); err != nil {
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
	if err = addOpCreateEdgeDeploymentPlanValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEdgeDeploymentPlan(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateEdgeDeploymentPlan(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateEdgeDeploymentPlan",
	}
}
