// Code generated by smithy-go-codegen DO NOT EDIT.

package servicediscovery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the current health status ( Healthy , Unhealthy , or Unknown ) of one or
// more instances that are associated with a specified service. There's a brief
// delay between when you register an instance and when the health status for the
// instance is available.
func (c *Client) GetInstancesHealthStatus(ctx context.Context, params *GetInstancesHealthStatusInput, optFns ...func(*Options)) (*GetInstancesHealthStatusOutput, error) {
	if params == nil {
		params = &GetInstancesHealthStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetInstancesHealthStatus", params, optFns, c.addOperationGetInstancesHealthStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetInstancesHealthStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetInstancesHealthStatusInput struct {

	// The ID of the service that the instance is associated with.
	//
	// This member is required.
	ServiceId *string

	// An array that contains the IDs of all the instances that you want to get the
	// health status for. If you omit Instances , Cloud Map returns the health status
	// for all the instances that are associated with the specified service. To get the
	// IDs for the instances that you've registered by using a specified service,
	// submit a ListInstances (https://docs.aws.amazon.com/cloud-map/latest/api/API_ListInstances.html)
	// request.
	Instances []string

	// The maximum number of instances that you want Cloud Map to return in the
	// response to a GetInstancesHealthStatus request. If you don't specify a value
	// for MaxResults , Cloud Map returns up to 100 instances.
	MaxResults *int32

	// For the first GetInstancesHealthStatus request, omit this value. If more than
	// MaxResults instances match the specified criteria, you can submit another
	// GetInstancesHealthStatus request to get the next group of results. Specify the
	// value of NextToken from the previous response in the next request.
	NextToken *string

	noSmithyDocumentSerde
}

type GetInstancesHealthStatusOutput struct {

	// If more than MaxResults instances match the specified criteria, you can submit
	// another GetInstancesHealthStatus request to get the next group of results.
	// Specify the value of NextToken from the previous response in the next request.
	NextToken *string

	// A complex type that contains the IDs and the health status of the instances
	// that you specified in the GetInstancesHealthStatus request.
	Status map[string]types.HealthStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetInstancesHealthStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetInstancesHealthStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetInstancesHealthStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetInstancesHealthStatus"); err != nil {
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
	if err = addOpGetInstancesHealthStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetInstancesHealthStatus(options.Region), middleware.Before); err != nil {
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

// GetInstancesHealthStatusAPIClient is a client that implements the
// GetInstancesHealthStatus operation.
type GetInstancesHealthStatusAPIClient interface {
	GetInstancesHealthStatus(context.Context, *GetInstancesHealthStatusInput, ...func(*Options)) (*GetInstancesHealthStatusOutput, error)
}

var _ GetInstancesHealthStatusAPIClient = (*Client)(nil)

// GetInstancesHealthStatusPaginatorOptions is the paginator options for
// GetInstancesHealthStatus
type GetInstancesHealthStatusPaginatorOptions struct {
	// The maximum number of instances that you want Cloud Map to return in the
	// response to a GetInstancesHealthStatus request. If you don't specify a value
	// for MaxResults , Cloud Map returns up to 100 instances.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetInstancesHealthStatusPaginator is a paginator for GetInstancesHealthStatus
type GetInstancesHealthStatusPaginator struct {
	options   GetInstancesHealthStatusPaginatorOptions
	client    GetInstancesHealthStatusAPIClient
	params    *GetInstancesHealthStatusInput
	nextToken *string
	firstPage bool
}

// NewGetInstancesHealthStatusPaginator returns a new
// GetInstancesHealthStatusPaginator
func NewGetInstancesHealthStatusPaginator(client GetInstancesHealthStatusAPIClient, params *GetInstancesHealthStatusInput, optFns ...func(*GetInstancesHealthStatusPaginatorOptions)) *GetInstancesHealthStatusPaginator {
	if params == nil {
		params = &GetInstancesHealthStatusInput{}
	}

	options := GetInstancesHealthStatusPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetInstancesHealthStatusPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetInstancesHealthStatusPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetInstancesHealthStatus page.
func (p *GetInstancesHealthStatusPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetInstancesHealthStatusOutput, error) {
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

	result, err := p.client.GetInstancesHealthStatus(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetInstancesHealthStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetInstancesHealthStatus",
	}
}
