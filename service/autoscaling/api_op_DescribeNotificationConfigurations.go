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

// Gets information about the Amazon SNS notifications that are configured for one
// or more Auto Scaling groups.
func (c *Client) DescribeNotificationConfigurations(ctx context.Context, params *DescribeNotificationConfigurationsInput, optFns ...func(*Options)) (*DescribeNotificationConfigurationsOutput, error) {
	if params == nil {
		params = &DescribeNotificationConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeNotificationConfigurations", params, optFns, c.addOperationDescribeNotificationConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeNotificationConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeNotificationConfigurationsInput struct {

	// The name of the Auto Scaling group.
	AutoScalingGroupNames []string

	// The maximum number of items to return with this call. The default value is 50
	// and the maximum value is 100 .
	MaxRecords *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeNotificationConfigurationsOutput struct {

	// The notification configurations.
	//
	// This member is required.
	NotificationConfigurations []types.NotificationConfiguration

	// A string that indicates that the response contains more items than can be
	// returned in a single response. To receive additional items, specify this string
	// for the NextToken value when requesting the next set of items. This value is
	// null when there are no more items to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeNotificationConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeNotificationConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeNotificationConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeNotificationConfigurations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNotificationConfigurations(options.Region), middleware.Before); err != nil {
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

// DescribeNotificationConfigurationsAPIClient is a client that implements the
// DescribeNotificationConfigurations operation.
type DescribeNotificationConfigurationsAPIClient interface {
	DescribeNotificationConfigurations(context.Context, *DescribeNotificationConfigurationsInput, ...func(*Options)) (*DescribeNotificationConfigurationsOutput, error)
}

var _ DescribeNotificationConfigurationsAPIClient = (*Client)(nil)

// DescribeNotificationConfigurationsPaginatorOptions is the paginator options for
// DescribeNotificationConfigurations
type DescribeNotificationConfigurationsPaginatorOptions struct {
	// The maximum number of items to return with this call. The default value is 50
	// and the maximum value is 100 .
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeNotificationConfigurationsPaginator is a paginator for
// DescribeNotificationConfigurations
type DescribeNotificationConfigurationsPaginator struct {
	options   DescribeNotificationConfigurationsPaginatorOptions
	client    DescribeNotificationConfigurationsAPIClient
	params    *DescribeNotificationConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeNotificationConfigurationsPaginator returns a new
// DescribeNotificationConfigurationsPaginator
func NewDescribeNotificationConfigurationsPaginator(client DescribeNotificationConfigurationsAPIClient, params *DescribeNotificationConfigurationsInput, optFns ...func(*DescribeNotificationConfigurationsPaginatorOptions)) *DescribeNotificationConfigurationsPaginator {
	if params == nil {
		params = &DescribeNotificationConfigurationsInput{}
	}

	options := DescribeNotificationConfigurationsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeNotificationConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeNotificationConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeNotificationConfigurations page.
func (p *DescribeNotificationConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeNotificationConfigurationsOutput, error) {
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

	result, err := p.client.DescribeNotificationConfigurations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeNotificationConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeNotificationConfigurations",
	}
}
