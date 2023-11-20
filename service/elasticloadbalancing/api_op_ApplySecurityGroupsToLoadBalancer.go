// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates one or more security groups with your load balancer in a virtual
// private cloud (VPC). The specified security groups override the previously
// associated security groups. For more information, see Security Groups for Load
// Balancers in a VPC (https://docs.aws.amazon.com/elasticloadbalancing/latest/classic/elb-security-groups.html#elb-vpc-security-groups)
// in the Classic Load Balancers Guide.
func (c *Client) ApplySecurityGroupsToLoadBalancer(ctx context.Context, params *ApplySecurityGroupsToLoadBalancerInput, optFns ...func(*Options)) (*ApplySecurityGroupsToLoadBalancerOutput, error) {
	if params == nil {
		params = &ApplySecurityGroupsToLoadBalancerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ApplySecurityGroupsToLoadBalancer", params, optFns, c.addOperationApplySecurityGroupsToLoadBalancerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ApplySecurityGroupsToLoadBalancerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for ApplySecurityGroupsToLoadBalancer.
type ApplySecurityGroupsToLoadBalancerInput struct {

	// The name of the load balancer.
	//
	// This member is required.
	LoadBalancerName *string

	// The IDs of the security groups to associate with the load balancer. Note that
	// you cannot specify the name of the security group.
	//
	// This member is required.
	SecurityGroups []string

	noSmithyDocumentSerde
}

// Contains the output of ApplySecurityGroupsToLoadBalancer.
type ApplySecurityGroupsToLoadBalancerOutput struct {

	// The IDs of the security groups associated with the load balancer.
	SecurityGroups []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationApplySecurityGroupsToLoadBalancerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpApplySecurityGroupsToLoadBalancer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpApplySecurityGroupsToLoadBalancer{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ApplySecurityGroupsToLoadBalancer"); err != nil {
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
	if err = addOpApplySecurityGroupsToLoadBalancerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opApplySecurityGroupsToLoadBalancer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opApplySecurityGroupsToLoadBalancer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ApplySecurityGroupsToLoadBalancer",
	}
}
