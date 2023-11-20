// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all of the authentication methods supported by the specified application.
func (c *Client) ListApplicationAuthenticationMethods(ctx context.Context, params *ListApplicationAuthenticationMethodsInput, optFns ...func(*Options)) (*ListApplicationAuthenticationMethodsOutput, error) {
	if params == nil {
		params = &ListApplicationAuthenticationMethodsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListApplicationAuthenticationMethods", params, optFns, c.addOperationListApplicationAuthenticationMethodsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListApplicationAuthenticationMethodsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListApplicationAuthenticationMethodsInput struct {

	// Specifies the ARN of the application with the authentication methods you want
	// to list.
	//
	// This member is required.
	ApplicationArn *string

	// Specifies that you want to receive the next page of results. Valid only if you
	// received a NextToken response in the previous request. If you did, it indicates
	// that more output is available. Set this parameter to the value provided by the
	// previous call's NextToken response to request the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListApplicationAuthenticationMethodsOutput struct {

	// An array list of authentication methods for the specified application.
	AuthenticationMethods []types.AuthenticationMethodItem

	// If present, this value indicates that more output is available than is included
	// in the current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null . This
	// indicates that this is the last page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListApplicationAuthenticationMethodsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListApplicationAuthenticationMethods{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListApplicationAuthenticationMethods{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListApplicationAuthenticationMethods"); err != nil {
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
	if err = addOpListApplicationAuthenticationMethodsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListApplicationAuthenticationMethods(options.Region), middleware.Before); err != nil {
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

// ListApplicationAuthenticationMethodsAPIClient is a client that implements the
// ListApplicationAuthenticationMethods operation.
type ListApplicationAuthenticationMethodsAPIClient interface {
	ListApplicationAuthenticationMethods(context.Context, *ListApplicationAuthenticationMethodsInput, ...func(*Options)) (*ListApplicationAuthenticationMethodsOutput, error)
}

var _ ListApplicationAuthenticationMethodsAPIClient = (*Client)(nil)

// ListApplicationAuthenticationMethodsPaginatorOptions is the paginator options
// for ListApplicationAuthenticationMethods
type ListApplicationAuthenticationMethodsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListApplicationAuthenticationMethodsPaginator is a paginator for
// ListApplicationAuthenticationMethods
type ListApplicationAuthenticationMethodsPaginator struct {
	options   ListApplicationAuthenticationMethodsPaginatorOptions
	client    ListApplicationAuthenticationMethodsAPIClient
	params    *ListApplicationAuthenticationMethodsInput
	nextToken *string
	firstPage bool
}

// NewListApplicationAuthenticationMethodsPaginator returns a new
// ListApplicationAuthenticationMethodsPaginator
func NewListApplicationAuthenticationMethodsPaginator(client ListApplicationAuthenticationMethodsAPIClient, params *ListApplicationAuthenticationMethodsInput, optFns ...func(*ListApplicationAuthenticationMethodsPaginatorOptions)) *ListApplicationAuthenticationMethodsPaginator {
	if params == nil {
		params = &ListApplicationAuthenticationMethodsInput{}
	}

	options := ListApplicationAuthenticationMethodsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListApplicationAuthenticationMethodsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListApplicationAuthenticationMethodsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListApplicationAuthenticationMethods page.
func (p *ListApplicationAuthenticationMethodsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListApplicationAuthenticationMethodsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListApplicationAuthenticationMethods(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListApplicationAuthenticationMethods(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListApplicationAuthenticationMethods",
	}
}
