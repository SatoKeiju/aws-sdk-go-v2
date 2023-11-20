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

// Updates a resiliency policy. Resilience Hub allows you to provide a value of
// zero for rtoInSecs and rpoInSecs of your resiliency policy. But, while
// assessing your application, the lowest possible assessment result is near zero.
// Hence, if you provide value zero for rtoInSecs and rpoInSecs , the estimated
// workload RTO and estimated workload RPO result will be near zero and the
// Compliance status for your application will be set to Policy breached.
func (c *Client) UpdateResiliencyPolicy(ctx context.Context, params *UpdateResiliencyPolicyInput, optFns ...func(*Options)) (*UpdateResiliencyPolicyOutput, error) {
	if params == nil {
		params = &UpdateResiliencyPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateResiliencyPolicy", params, optFns, c.addOperationUpdateResiliencyPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateResiliencyPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateResiliencyPolicyInput struct {

	// Amazon Resource Name (ARN) of the resiliency policy. The format for this ARN
	// is: arn: partition :resiliencehub: region : account :resiliency-policy/ policy-id
	// . For more information about ARNs, see Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference guide.
	//
	// This member is required.
	PolicyArn *string

	// Specifies a high-level geographical location constraint for where your
	// resilience policy data can be stored.
	DataLocationConstraint types.DataLocationConstraint

	// The type of resiliency policy to be created, including the recovery time
	// objective (RTO) and recovery point objective (RPO) in seconds.
	Policy map[string]types.FailurePolicy

	// The description for the policy.
	PolicyDescription *string

	// The name of the policy
	PolicyName *string

	// The tier for this resiliency policy, ranging from the highest severity (
	// MissionCritical ) to lowest ( NonCritical ).
	Tier types.ResiliencyPolicyTier

	noSmithyDocumentSerde
}

type UpdateResiliencyPolicyOutput struct {

	// The type of resiliency policy that was updated, including the recovery time
	// objective (RTO) and recovery point objective (RPO) in seconds.
	//
	// This member is required.
	Policy *types.ResiliencyPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateResiliencyPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateResiliencyPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateResiliencyPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateResiliencyPolicy"); err != nil {
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
	if err = addOpUpdateResiliencyPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateResiliencyPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateResiliencyPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateResiliencyPolicy",
	}
}
