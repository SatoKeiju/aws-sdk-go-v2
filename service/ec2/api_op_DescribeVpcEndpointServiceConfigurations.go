// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the VPC endpoint service configurations in your account (your
// services).
func (c *Client) DescribeVpcEndpointServiceConfigurations(ctx context.Context, params *DescribeVpcEndpointServiceConfigurationsInput, optFns ...func(*Options)) (*DescribeVpcEndpointServiceConfigurationsOutput, error) {
	if params == nil {
		params = &DescribeVpcEndpointServiceConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVpcEndpointServiceConfigurations", params, optFns, c.addOperationDescribeVpcEndpointServiceConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVpcEndpointServiceConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVpcEndpointServiceConfigurationsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	// The filters.
	//   - service-name - The name of the service.
	//   - service-id - The ID of the service.
	//   - service-state - The state of the service ( Pending | Available | Deleting |
	//   Deleted | Failed ).
	//   - supported-ip-address-types - The IP address type ( ipv4 | ipv6 ).
	//   - tag : - The key/value combination of a tag assigned to the resource. Use the
	//   tag key in the filter name and the tag value as the filter value. For example,
	//   to find all resources that have a tag with the key Owner and the value TeamA ,
	//   specify tag:Owner for the filter name and TeamA for the filter value.
	//   - tag-key - The key of a tag assigned to the resource. Use this filter to find
	//   all resources assigned a tag with a specific key, regardless of the tag value.
	Filters []types.Filter

	// The maximum number of results to return for the request in a single page. The
	// remaining results of the initial request can be seen by sending another request
	// with the returned NextToken value. This value can be between 5 and 1,000; if
	// MaxResults is given a value larger than 1,000, only 1,000 results are returned.
	MaxResults *int32

	// The token to retrieve the next page of results.
	NextToken *string

	// The IDs of the endpoint services.
	ServiceIds []string

	noSmithyDocumentSerde
}

type DescribeVpcEndpointServiceConfigurationsOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the services.
	ServiceConfigurations []types.ServiceConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeVpcEndpointServiceConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeVpcEndpointServiceConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVpcEndpointServiceConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeVpcEndpointServiceConfigurations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpcEndpointServiceConfigurations(options.Region), middleware.Before); err != nil {
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

// DescribeVpcEndpointServiceConfigurationsAPIClient is a client that implements
// the DescribeVpcEndpointServiceConfigurations operation.
type DescribeVpcEndpointServiceConfigurationsAPIClient interface {
	DescribeVpcEndpointServiceConfigurations(context.Context, *DescribeVpcEndpointServiceConfigurationsInput, ...func(*Options)) (*DescribeVpcEndpointServiceConfigurationsOutput, error)
}

var _ DescribeVpcEndpointServiceConfigurationsAPIClient = (*Client)(nil)

// DescribeVpcEndpointServiceConfigurationsPaginatorOptions is the paginator
// options for DescribeVpcEndpointServiceConfigurations
type DescribeVpcEndpointServiceConfigurationsPaginatorOptions struct {
	// The maximum number of results to return for the request in a single page. The
	// remaining results of the initial request can be seen by sending another request
	// with the returned NextToken value. This value can be between 5 and 1,000; if
	// MaxResults is given a value larger than 1,000, only 1,000 results are returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeVpcEndpointServiceConfigurationsPaginator is a paginator for
// DescribeVpcEndpointServiceConfigurations
type DescribeVpcEndpointServiceConfigurationsPaginator struct {
	options   DescribeVpcEndpointServiceConfigurationsPaginatorOptions
	client    DescribeVpcEndpointServiceConfigurationsAPIClient
	params    *DescribeVpcEndpointServiceConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeVpcEndpointServiceConfigurationsPaginator returns a new
// DescribeVpcEndpointServiceConfigurationsPaginator
func NewDescribeVpcEndpointServiceConfigurationsPaginator(client DescribeVpcEndpointServiceConfigurationsAPIClient, params *DescribeVpcEndpointServiceConfigurationsInput, optFns ...func(*DescribeVpcEndpointServiceConfigurationsPaginatorOptions)) *DescribeVpcEndpointServiceConfigurationsPaginator {
	if params == nil {
		params = &DescribeVpcEndpointServiceConfigurationsInput{}
	}

	options := DescribeVpcEndpointServiceConfigurationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeVpcEndpointServiceConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeVpcEndpointServiceConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeVpcEndpointServiceConfigurations page.
func (p *DescribeVpcEndpointServiceConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeVpcEndpointServiceConfigurationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.DescribeVpcEndpointServiceConfigurations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeVpcEndpointServiceConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeVpcEndpointServiceConfigurations",
	}
}
