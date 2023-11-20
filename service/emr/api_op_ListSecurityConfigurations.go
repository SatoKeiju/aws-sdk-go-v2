// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the security configurations visible to this account, providing their
// creation dates and times, and their names. This call returns a maximum of 50
// clusters per call, but returns a marker to track the paging of the cluster list
// across multiple ListSecurityConfigurations calls.
func (c *Client) ListSecurityConfigurations(ctx context.Context, params *ListSecurityConfigurationsInput, optFns ...func(*Options)) (*ListSecurityConfigurationsOutput, error) {
	if params == nil {
		params = &ListSecurityConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSecurityConfigurations", params, optFns, c.addOperationListSecurityConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSecurityConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSecurityConfigurationsInput struct {

	// The pagination token that indicates the set of results to retrieve.
	Marker *string

	noSmithyDocumentSerde
}

type ListSecurityConfigurationsOutput struct {

	// A pagination token that indicates the next set of results to retrieve. Include
	// the marker in the next ListSecurityConfiguration call to retrieve the next page
	// of results, if required.
	Marker *string

	// The creation date and time, and name, of each security configuration.
	SecurityConfigurations []types.SecurityConfigurationSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSecurityConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSecurityConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSecurityConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSecurityConfigurations"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSecurityConfigurations(options.Region), middleware.Before); err != nil {
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

// ListSecurityConfigurationsAPIClient is a client that implements the
// ListSecurityConfigurations operation.
type ListSecurityConfigurationsAPIClient interface {
	ListSecurityConfigurations(context.Context, *ListSecurityConfigurationsInput, ...func(*Options)) (*ListSecurityConfigurationsOutput, error)
}

var _ ListSecurityConfigurationsAPIClient = (*Client)(nil)

// ListSecurityConfigurationsPaginatorOptions is the paginator options for
// ListSecurityConfigurations
type ListSecurityConfigurationsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSecurityConfigurationsPaginator is a paginator for
// ListSecurityConfigurations
type ListSecurityConfigurationsPaginator struct {
	options   ListSecurityConfigurationsPaginatorOptions
	client    ListSecurityConfigurationsAPIClient
	params    *ListSecurityConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewListSecurityConfigurationsPaginator returns a new
// ListSecurityConfigurationsPaginator
func NewListSecurityConfigurationsPaginator(client ListSecurityConfigurationsAPIClient, params *ListSecurityConfigurationsInput, optFns ...func(*ListSecurityConfigurationsPaginatorOptions)) *ListSecurityConfigurationsPaginator {
	if params == nil {
		params = &ListSecurityConfigurationsInput{}
	}

	options := ListSecurityConfigurationsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSecurityConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSecurityConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSecurityConfigurations page.
func (p *ListSecurityConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSecurityConfigurationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	result, err := p.client.ListSecurityConfigurations(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListSecurityConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSecurityConfigurations",
	}
}
