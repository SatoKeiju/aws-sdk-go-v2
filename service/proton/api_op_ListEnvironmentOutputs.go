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

// List the infrastructure as code outputs for your environment.
func (c *Client) ListEnvironmentOutputs(ctx context.Context, params *ListEnvironmentOutputsInput, optFns ...func(*Options)) (*ListEnvironmentOutputsOutput, error) {
	if params == nil {
		params = &ListEnvironmentOutputsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEnvironmentOutputs", params, optFns, c.addOperationListEnvironmentOutputsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEnvironmentOutputsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEnvironmentOutputsInput struct {

	// The environment name.
	//
	// This member is required.
	EnvironmentName *string

	// The ID of the deployment whose outputs you want.
	DeploymentId *string

	// A token that indicates the location of the next environment output in the array
	// of environment outputs, after the list of environment outputs that was
	// previously requested.
	NextToken *string

	noSmithyDocumentSerde
}

type ListEnvironmentOutputsOutput struct {

	// An array of environment outputs with detail data.
	//
	// This member is required.
	Outputs []types.Output

	// A token that indicates the location of the next environment output in the array
	// of environment outputs, after the current requested list of environment outputs.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEnvironmentOutputsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListEnvironmentOutputs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListEnvironmentOutputs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListEnvironmentOutputs"); err != nil {
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
	if err = addOpListEnvironmentOutputsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEnvironmentOutputs(options.Region), middleware.Before); err != nil {
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

// ListEnvironmentOutputsAPIClient is a client that implements the
// ListEnvironmentOutputs operation.
type ListEnvironmentOutputsAPIClient interface {
	ListEnvironmentOutputs(context.Context, *ListEnvironmentOutputsInput, ...func(*Options)) (*ListEnvironmentOutputsOutput, error)
}

var _ ListEnvironmentOutputsAPIClient = (*Client)(nil)

// ListEnvironmentOutputsPaginatorOptions is the paginator options for
// ListEnvironmentOutputs
type ListEnvironmentOutputsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEnvironmentOutputsPaginator is a paginator for ListEnvironmentOutputs
type ListEnvironmentOutputsPaginator struct {
	options   ListEnvironmentOutputsPaginatorOptions
	client    ListEnvironmentOutputsAPIClient
	params    *ListEnvironmentOutputsInput
	nextToken *string
	firstPage bool
}

// NewListEnvironmentOutputsPaginator returns a new ListEnvironmentOutputsPaginator
func NewListEnvironmentOutputsPaginator(client ListEnvironmentOutputsAPIClient, params *ListEnvironmentOutputsInput, optFns ...func(*ListEnvironmentOutputsPaginatorOptions)) *ListEnvironmentOutputsPaginator {
	if params == nil {
		params = &ListEnvironmentOutputsInput{}
	}

	options := ListEnvironmentOutputsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEnvironmentOutputsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEnvironmentOutputsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEnvironmentOutputs page.
func (p *ListEnvironmentOutputsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEnvironmentOutputsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListEnvironmentOutputs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEnvironmentOutputs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListEnvironmentOutputs",
	}
}
