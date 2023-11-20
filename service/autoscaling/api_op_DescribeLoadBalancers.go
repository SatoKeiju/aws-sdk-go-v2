// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API operation is superseded by DescribeTrafficSources , which can describe
// multiple traffic sources types. We recommend using DescribeTrafficSources to
// simplify how you manage traffic sources. However, we continue to support
// DescribeLoadBalancers . You can use both the original DescribeLoadBalancers API
// operation and DescribeTrafficSources on the same Auto Scaling group. Gets
// information about the load balancers for the specified Auto Scaling group. This
// operation describes only Classic Load Balancers. If you have Application Load
// Balancers, Network Load Balancers, or Gateway Load Balancers, use the
// DescribeLoadBalancerTargetGroups API instead. To determine the attachment status
// of the load balancer, use the State element in the response. When you attach a
// load balancer to an Auto Scaling group, the initial State value is Adding . The
// state transitions to Added after all Auto Scaling instances are registered with
// the load balancer. If Elastic Load Balancing health checks are enabled for the
// Auto Scaling group, the state transitions to InService after at least one Auto
// Scaling instance passes the health check. When the load balancer is in the
// InService state, Amazon EC2 Auto Scaling can terminate and replace any instances
// that are reported as unhealthy. If no registered instances pass the health
// checks, the load balancer doesn't enter the InService state. Load balancers
// also have an InService state if you attach them in the CreateAutoScalingGroup
// API call. If your load balancer state is InService , but it is not working
// properly, check the scaling activities by calling DescribeScalingActivities and
// take any corrective actions necessary. For help with failed health checks, see
// Troubleshooting Amazon EC2 Auto Scaling: Health checks (https://docs.aws.amazon.com/autoscaling/ec2/userguide/ts-as-healthchecks.html)
// in the Amazon EC2 Auto Scaling User Guide. For more information, see Use
// Elastic Load Balancing to distribute traffic across the instances in your Auto
// Scaling group (https://docs.aws.amazon.com/autoscaling/ec2/userguide/autoscaling-load-balancer.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) DescribeLoadBalancers(ctx context.Context, params *DescribeLoadBalancersInput, optFns ...func(*Options)) (*DescribeLoadBalancersOutput, error) {
	if params == nil {
		params = &DescribeLoadBalancersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLoadBalancers", params, optFns, c.addOperationDescribeLoadBalancersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLoadBalancersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLoadBalancersInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The maximum number of items to return with this call. The default value is 100
	// and the maximum value is 100 .
	MaxRecords *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeLoadBalancersOutput struct {

	// The load balancers.
	LoadBalancers []types.LoadBalancerState

	// A string that indicates that the response contains more items than can be
	// returned in a single response. To receive additional items, specify this string
	// for the NextToken value when requesting the next set of items. This value is
	// null when there are no more items to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLoadBalancersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeLoadBalancers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeLoadBalancers{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeLoadBalancers"); err != nil {
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
	if err = addOpDescribeLoadBalancersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLoadBalancers(options.Region), middleware.Before); err != nil {
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

// DescribeLoadBalancersAPIClient is a client that implements the
// DescribeLoadBalancers operation.
type DescribeLoadBalancersAPIClient interface {
	DescribeLoadBalancers(context.Context, *DescribeLoadBalancersInput, ...func(*Options)) (*DescribeLoadBalancersOutput, error)
}

var _ DescribeLoadBalancersAPIClient = (*Client)(nil)

// DescribeLoadBalancersPaginatorOptions is the paginator options for
// DescribeLoadBalancers
type DescribeLoadBalancersPaginatorOptions struct {
	// The maximum number of items to return with this call. The default value is 100
	// and the maximum value is 100 .
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeLoadBalancersPaginator is a paginator for DescribeLoadBalancers
type DescribeLoadBalancersPaginator struct {
	options   DescribeLoadBalancersPaginatorOptions
	client    DescribeLoadBalancersAPIClient
	params    *DescribeLoadBalancersInput
	nextToken *string
	firstPage bool
}

// NewDescribeLoadBalancersPaginator returns a new DescribeLoadBalancersPaginator
func NewDescribeLoadBalancersPaginator(client DescribeLoadBalancersAPIClient, params *DescribeLoadBalancersInput, optFns ...func(*DescribeLoadBalancersPaginatorOptions)) *DescribeLoadBalancersPaginator {
	if params == nil {
		params = &DescribeLoadBalancersInput{}
	}

	options := DescribeLoadBalancersPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeLoadBalancersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeLoadBalancersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeLoadBalancers page.
func (p *DescribeLoadBalancersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeLoadBalancersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeLoadBalancers(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeLoadBalancers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeLoadBalancers",
	}
}
