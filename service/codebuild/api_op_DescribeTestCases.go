// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of details about test cases for a report.
func (c *Client) DescribeTestCases(ctx context.Context, params *DescribeTestCasesInput, optFns ...func(*Options)) (*DescribeTestCasesOutput, error) {
	if params == nil {
		params = &DescribeTestCasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTestCases", params, optFns, c.addOperationDescribeTestCasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTestCasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTestCasesInput struct {

	// The ARN of the report for which test cases are returned.
	//
	// This member is required.
	ReportArn *string

	// A TestCaseFilter object used to filter the returned reports.
	Filter *types.TestCaseFilter

	// The maximum number of paginated test cases returned per response. Use nextToken
	// to iterate pages in the list of returned TestCase objects. The default value is
	// 100.
	MaxResults *int32

	// During a previous call, the maximum number of items that can be returned is the
	// value specified in maxResults . If there more items in the list, then a unique
	// string called a nextToken is returned. To get the next batch of items in the
	// list, call this operation again, adding the next token to the call. To get all
	// of the items in the list, keep calling this operation with each subsequent next
	// token that is returned, until no more next tokens are returned.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeTestCasesOutput struct {

	// During a previous call, the maximum number of items that can be returned is the
	// value specified in maxResults . If there more items in the list, then a unique
	// string called a nextToken is returned. To get the next batch of items in the
	// list, call this operation again, adding the next token to the call. To get all
	// of the items in the list, keep calling this operation with each subsequent next
	// token that is returned, until no more next tokens are returned.
	NextToken *string

	// The returned list of test cases.
	TestCases []types.TestCase

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTestCasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTestCases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTestCases{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeTestCases"); err != nil {
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
	if err = addOpDescribeTestCasesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTestCases(options.Region), middleware.Before); err != nil {
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

// DescribeTestCasesAPIClient is a client that implements the DescribeTestCases
// operation.
type DescribeTestCasesAPIClient interface {
	DescribeTestCases(context.Context, *DescribeTestCasesInput, ...func(*Options)) (*DescribeTestCasesOutput, error)
}

var _ DescribeTestCasesAPIClient = (*Client)(nil)

// DescribeTestCasesPaginatorOptions is the paginator options for DescribeTestCases
type DescribeTestCasesPaginatorOptions struct {
	// The maximum number of paginated test cases returned per response. Use nextToken
	// to iterate pages in the list of returned TestCase objects. The default value is
	// 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeTestCasesPaginator is a paginator for DescribeTestCases
type DescribeTestCasesPaginator struct {
	options   DescribeTestCasesPaginatorOptions
	client    DescribeTestCasesAPIClient
	params    *DescribeTestCasesInput
	nextToken *string
	firstPage bool
}

// NewDescribeTestCasesPaginator returns a new DescribeTestCasesPaginator
func NewDescribeTestCasesPaginator(client DescribeTestCasesAPIClient, params *DescribeTestCasesInput, optFns ...func(*DescribeTestCasesPaginatorOptions)) *DescribeTestCasesPaginator {
	if params == nil {
		params = &DescribeTestCasesInput{}
	}

	options := DescribeTestCasesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeTestCasesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeTestCasesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeTestCases page.
func (p *DescribeTestCasesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeTestCasesOutput, error) {
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

	result, err := p.client.DescribeTestCases(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeTestCases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeTestCases",
	}
}
