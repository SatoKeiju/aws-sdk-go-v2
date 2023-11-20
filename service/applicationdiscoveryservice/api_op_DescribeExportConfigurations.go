// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// DescribeExportConfigurations is deprecated. Use DescribeExportTasks (https://docs.aws.amazon.com/application-discovery/latest/APIReference/API_DescribeExportTasks.html)
// , instead.
//
// Deprecated: This operation has been deprecated.
func (c *Client) DescribeExportConfigurations(ctx context.Context, params *DescribeExportConfigurationsInput, optFns ...func(*Options)) (*DescribeExportConfigurationsOutput, error) {
	if params == nil {
		params = &DescribeExportConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeExportConfigurations", params, optFns, c.addOperationDescribeExportConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeExportConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeExportConfigurationsInput struct {

	// A list of continuous export IDs to search for.
	ExportIds []string

	// A number between 1 and 100 specifying the maximum number of continuous export
	// descriptions returned.
	MaxResults int32

	// The token from the previous call to describe-export-tasks.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeExportConfigurationsOutput struct {

	//
	ExportsInfo []types.ExportInfo

	// The token from the previous call to describe-export-tasks.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeExportConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeExportConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeExportConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeExportConfigurations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeExportConfigurations(options.Region), middleware.Before); err != nil {
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

// DescribeExportConfigurationsAPIClient is a client that implements the
// DescribeExportConfigurations operation.
type DescribeExportConfigurationsAPIClient interface {
	DescribeExportConfigurations(context.Context, *DescribeExportConfigurationsInput, ...func(*Options)) (*DescribeExportConfigurationsOutput, error)
}

var _ DescribeExportConfigurationsAPIClient = (*Client)(nil)

// DescribeExportConfigurationsPaginatorOptions is the paginator options for
// DescribeExportConfigurations
type DescribeExportConfigurationsPaginatorOptions struct {
	// A number between 1 and 100 specifying the maximum number of continuous export
	// descriptions returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeExportConfigurationsPaginator is a paginator for
// DescribeExportConfigurations
type DescribeExportConfigurationsPaginator struct {
	options   DescribeExportConfigurationsPaginatorOptions
	client    DescribeExportConfigurationsAPIClient
	params    *DescribeExportConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewDescribeExportConfigurationsPaginator returns a new
// DescribeExportConfigurationsPaginator
func NewDescribeExportConfigurationsPaginator(client DescribeExportConfigurationsAPIClient, params *DescribeExportConfigurationsInput, optFns ...func(*DescribeExportConfigurationsPaginatorOptions)) *DescribeExportConfigurationsPaginator {
	if params == nil {
		params = &DescribeExportConfigurationsInput{}
	}

	options := DescribeExportConfigurationsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeExportConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeExportConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeExportConfigurations page.
func (p *DescribeExportConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeExportConfigurationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeExportConfigurations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeExportConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeExportConfigurations",
	}
}
