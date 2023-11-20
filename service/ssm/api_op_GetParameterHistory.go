// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the history of all changes to a parameter. If you change the KMS key
// alias for the KMS key used to encrypt a parameter, then you must also update the
// key alias the parameter uses to reference KMS. Otherwise, GetParameterHistory
// retrieves whatever the original key alias was referencing.
func (c *Client) GetParameterHistory(ctx context.Context, params *GetParameterHistoryInput, optFns ...func(*Options)) (*GetParameterHistoryOutput, error) {
	if params == nil {
		params = &GetParameterHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetParameterHistory", params, optFns, c.addOperationGetParameterHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetParameterHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetParameterHistoryInput struct {

	// The name of the parameter for which you want to review history.
	//
	// This member is required.
	Name *string

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	// Return decrypted values for secure string parameters. This flag is ignored for
	// String and StringList parameter types.
	WithDecryption *bool

	noSmithyDocumentSerde
}

type GetParameterHistoryOutput struct {

	// The token to use when requesting the next set of items. If there are no
	// additional items to return, the string is empty.
	NextToken *string

	// A list of parameters returned by the request.
	Parameters []types.ParameterHistory

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetParameterHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetParameterHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetParameterHistory{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetParameterHistory"); err != nil {
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
	if err = addOpGetParameterHistoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetParameterHistory(options.Region), middleware.Before); err != nil {
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

// GetParameterHistoryAPIClient is a client that implements the
// GetParameterHistory operation.
type GetParameterHistoryAPIClient interface {
	GetParameterHistory(context.Context, *GetParameterHistoryInput, ...func(*Options)) (*GetParameterHistoryOutput, error)
}

var _ GetParameterHistoryAPIClient = (*Client)(nil)

// GetParameterHistoryPaginatorOptions is the paginator options for
// GetParameterHistory
type GetParameterHistoryPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetParameterHistoryPaginator is a paginator for GetParameterHistory
type GetParameterHistoryPaginator struct {
	options   GetParameterHistoryPaginatorOptions
	client    GetParameterHistoryAPIClient
	params    *GetParameterHistoryInput
	nextToken *string
	firstPage bool
}

// NewGetParameterHistoryPaginator returns a new GetParameterHistoryPaginator
func NewGetParameterHistoryPaginator(client GetParameterHistoryAPIClient, params *GetParameterHistoryInput, optFns ...func(*GetParameterHistoryPaginatorOptions)) *GetParameterHistoryPaginator {
	if params == nil {
		params = &GetParameterHistoryInput{}
	}

	options := GetParameterHistoryPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetParameterHistoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetParameterHistoryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetParameterHistory page.
func (p *GetParameterHistoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetParameterHistoryOutput, error) {
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

	result, err := p.client.GetParameterHistory(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetParameterHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetParameterHistory",
	}
}
