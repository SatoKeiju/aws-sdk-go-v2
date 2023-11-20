// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the accounts in an organization that are contained by the specified
// target root or organizational unit (OU). If you specify the root, you get a list
// of all the accounts that aren't in any OU. If you specify an OU, you get a list
// of all the accounts in only that OU and not in any child OUs. To get a list of
// all accounts in the organization, use the ListAccounts operation. Always check
// the NextToken response parameter for a null value when calling a List*
// operation. These operations can occasionally return an empty set of results even
// when there are more results available. The NextToken response parameter value
// is null only when there are no more results to display. This operation can be
// called only from the organization's management account or by a member account
// that is a delegated administrator for an Amazon Web Services service.
func (c *Client) ListAccountsForParent(ctx context.Context, params *ListAccountsForParentInput, optFns ...func(*Options)) (*ListAccountsForParentOutput, error) {
	if params == nil {
		params = &ListAccountsForParentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAccountsForParent", params, optFns, c.addOperationListAccountsForParentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAccountsForParentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAccountsForParentInput struct {

	// The unique identifier (ID) for the parent root or organization unit (OU) whose
	// accounts you want to list.
	//
	// This member is required.
	ParentId *string

	// The total number of results that you want included on each page of the
	// response. If you do not include this parameter, it defaults to a value that is
	// specific to the operation. If additional items exist beyond the maximum you
	// specify, the NextToken response element is present and has a value (is not
	// null). Include that value as the NextToken request parameter in the next call
	// to the operation to get the next part of the results. Note that Organizations
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	MaxResults *int32

	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more output
	// is available. Set this parameter to the value of the previous call's NextToken
	// response to indicate where the output should continue from.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAccountsForParentOutput struct {

	// A list of the accounts in the specified root or OU.
	Accounts []types.Account

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null .
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAccountsForParentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAccountsForParent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAccountsForParent{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListAccountsForParent"); err != nil {
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
	if err = addOpListAccountsForParentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAccountsForParent(options.Region), middleware.Before); err != nil {
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

// ListAccountsForParentAPIClient is a client that implements the
// ListAccountsForParent operation.
type ListAccountsForParentAPIClient interface {
	ListAccountsForParent(context.Context, *ListAccountsForParentInput, ...func(*Options)) (*ListAccountsForParentOutput, error)
}

var _ ListAccountsForParentAPIClient = (*Client)(nil)

// ListAccountsForParentPaginatorOptions is the paginator options for
// ListAccountsForParent
type ListAccountsForParentPaginatorOptions struct {
	// The total number of results that you want included on each page of the
	// response. If you do not include this parameter, it defaults to a value that is
	// specific to the operation. If additional items exist beyond the maximum you
	// specify, the NextToken response element is present and has a value (is not
	// null). Include that value as the NextToken request parameter in the next call
	// to the operation to get the next part of the results. Note that Organizations
	// might return fewer results than the maximum even when there are more results
	// available. You should check NextToken after every operation to ensure that you
	// receive all of the results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAccountsForParentPaginator is a paginator for ListAccountsForParent
type ListAccountsForParentPaginator struct {
	options   ListAccountsForParentPaginatorOptions
	client    ListAccountsForParentAPIClient
	params    *ListAccountsForParentInput
	nextToken *string
	firstPage bool
}

// NewListAccountsForParentPaginator returns a new ListAccountsForParentPaginator
func NewListAccountsForParentPaginator(client ListAccountsForParentAPIClient, params *ListAccountsForParentInput, optFns ...func(*ListAccountsForParentPaginatorOptions)) *ListAccountsForParentPaginator {
	if params == nil {
		params = &ListAccountsForParentInput{}
	}

	options := ListAccountsForParentPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAccountsForParentPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAccountsForParentPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAccountsForParent page.
func (p *ListAccountsForParentPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAccountsForParentOutput, error) {
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

	result, err := p.client.ListAccountsForParent(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAccountsForParent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListAccountsForParent",
	}
}
