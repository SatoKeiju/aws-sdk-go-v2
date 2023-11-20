// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of security standards controls. For each control, the results
// include information about whether it is currently enabled, the severity, and a
// link to remediation information.
func (c *Client) DescribeStandardsControls(ctx context.Context, params *DescribeStandardsControlsInput, optFns ...func(*Options)) (*DescribeStandardsControlsOutput, error) {
	if params == nil {
		params = &DescribeStandardsControlsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStandardsControls", params, optFns, c.addOperationDescribeStandardsControlsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStandardsControlsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStandardsControlsInput struct {

	// The ARN of a resource that represents your subscription to a supported
	// standard. To get the subscription ARNs of the standards you have enabled, use
	// the GetEnabledStandards operation.
	//
	// This member is required.
	StandardsSubscriptionArn *string

	// The maximum number of security standard controls to return.
	MaxResults *int32

	// The token that is required for pagination. On your first call to the
	// DescribeStandardsControls operation, set the value of this parameter to NULL .
	// For subsequent calls to the operation, to continue listing data, set the value
	// of this parameter to the value returned from the previous response.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeStandardsControlsOutput struct {

	// A list of security standards controls.
	Controls []types.StandardsControl

	// The pagination token to use to request the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeStandardsControlsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeStandardsControls{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeStandardsControls{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeStandardsControls"); err != nil {
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
	if err = addOpDescribeStandardsControlsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStandardsControls(options.Region), middleware.Before); err != nil {
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

// DescribeStandardsControlsAPIClient is a client that implements the
// DescribeStandardsControls operation.
type DescribeStandardsControlsAPIClient interface {
	DescribeStandardsControls(context.Context, *DescribeStandardsControlsInput, ...func(*Options)) (*DescribeStandardsControlsOutput, error)
}

var _ DescribeStandardsControlsAPIClient = (*Client)(nil)

// DescribeStandardsControlsPaginatorOptions is the paginator options for
// DescribeStandardsControls
type DescribeStandardsControlsPaginatorOptions struct {
	// The maximum number of security standard controls to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeStandardsControlsPaginator is a paginator for DescribeStandardsControls
type DescribeStandardsControlsPaginator struct {
	options   DescribeStandardsControlsPaginatorOptions
	client    DescribeStandardsControlsAPIClient
	params    *DescribeStandardsControlsInput
	nextToken *string
	firstPage bool
}

// NewDescribeStandardsControlsPaginator returns a new
// DescribeStandardsControlsPaginator
func NewDescribeStandardsControlsPaginator(client DescribeStandardsControlsAPIClient, params *DescribeStandardsControlsInput, optFns ...func(*DescribeStandardsControlsPaginatorOptions)) *DescribeStandardsControlsPaginator {
	if params == nil {
		params = &DescribeStandardsControlsInput{}
	}

	options := DescribeStandardsControlsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeStandardsControlsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeStandardsControlsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeStandardsControls page.
func (p *DescribeStandardsControlsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeStandardsControlsOutput, error) {
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

	result, err := p.client.DescribeStandardsControls(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeStandardsControls(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeStandardsControls",
	}
}
