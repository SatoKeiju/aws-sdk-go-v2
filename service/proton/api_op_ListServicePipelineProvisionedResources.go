// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List provisioned resources for a service and pipeline with details.
func (c *Client) ListServicePipelineProvisionedResources(ctx context.Context, params *ListServicePipelineProvisionedResourcesInput, optFns ...func(*Options)) (*ListServicePipelineProvisionedResourcesOutput, error) {
	if params == nil {
		params = &ListServicePipelineProvisionedResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListServicePipelineProvisionedResources", params, optFns, c.addOperationListServicePipelineProvisionedResourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListServicePipelineProvisionedResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListServicePipelineProvisionedResourcesInput struct {

	// The name of the service whose pipeline's provisioned resources you want.
	//
	// This member is required.
	ServiceName *string

	// A token that indicates the location of the next provisioned resource in the
	// array of provisioned resources, after the list of provisioned resources that was
	// previously requested.
	NextToken *string

	noSmithyDocumentSerde
}

type ListServicePipelineProvisionedResourcesOutput struct {

	// An array of provisioned resources for a service and pipeline.
	//
	// This member is required.
	ProvisionedResources []types.ProvisionedResource

	// A token that indicates the location of the next provisioned resource in the
	// array of provisioned resources, after the current requested list of provisioned
	// resources.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListServicePipelineProvisionedResourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListServicePipelineProvisionedResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListServicePipelineProvisionedResources{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListServicePipelineProvisionedResources"); err != nil {
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
	if err = addOpListServicePipelineProvisionedResourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListServicePipelineProvisionedResources(options.Region), middleware.Before); err != nil {
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

// ListServicePipelineProvisionedResourcesAPIClient is a client that implements
// the ListServicePipelineProvisionedResources operation.
type ListServicePipelineProvisionedResourcesAPIClient interface {
	ListServicePipelineProvisionedResources(context.Context, *ListServicePipelineProvisionedResourcesInput, ...func(*Options)) (*ListServicePipelineProvisionedResourcesOutput, error)
}

var _ ListServicePipelineProvisionedResourcesAPIClient = (*Client)(nil)

// ListServicePipelineProvisionedResourcesPaginatorOptions is the paginator
// options for ListServicePipelineProvisionedResources
type ListServicePipelineProvisionedResourcesPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListServicePipelineProvisionedResourcesPaginator is a paginator for
// ListServicePipelineProvisionedResources
type ListServicePipelineProvisionedResourcesPaginator struct {
	options   ListServicePipelineProvisionedResourcesPaginatorOptions
	client    ListServicePipelineProvisionedResourcesAPIClient
	params    *ListServicePipelineProvisionedResourcesInput
	nextToken *string
	firstPage bool
}

// NewListServicePipelineProvisionedResourcesPaginator returns a new
// ListServicePipelineProvisionedResourcesPaginator
func NewListServicePipelineProvisionedResourcesPaginator(client ListServicePipelineProvisionedResourcesAPIClient, params *ListServicePipelineProvisionedResourcesInput, optFns ...func(*ListServicePipelineProvisionedResourcesPaginatorOptions)) *ListServicePipelineProvisionedResourcesPaginator {
	if params == nil {
		params = &ListServicePipelineProvisionedResourcesInput{}
	}

	options := ListServicePipelineProvisionedResourcesPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListServicePipelineProvisionedResourcesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListServicePipelineProvisionedResourcesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListServicePipelineProvisionedResources page.
func (p *ListServicePipelineProvisionedResourcesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListServicePipelineProvisionedResourcesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListServicePipelineProvisionedResources(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListServicePipelineProvisionedResources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListServicePipelineProvisionedResources",
	}
}
