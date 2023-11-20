// Code generated by smithy-go-codegen DO NOT EDIT.

package budgets

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/budgets/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes all of the budget actions for a budget.
func (c *Client) DescribeBudgetActionsForBudget(ctx context.Context, params *DescribeBudgetActionsForBudgetInput, optFns ...func(*Options)) (*DescribeBudgetActionsForBudgetOutput, error) {
	if params == nil {
		params = &DescribeBudgetActionsForBudgetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBudgetActionsForBudget", params, optFns, c.addOperationDescribeBudgetActionsForBudgetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBudgetActionsForBudgetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBudgetActionsForBudgetInput struct {

	// The account ID of the user. It's a 12-digit number.
	//
	// This member is required.
	AccountId *string

	// A string that represents the budget name. The ":" and "\" characters, and the
	// "/action/" substring, aren't allowed.
	//
	// This member is required.
	BudgetName *string

	// An integer that represents how many entries a paginated response contains. The
	// maximum is 100.
	MaxResults *int32

	// A generic string.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeBudgetActionsForBudgetOutput struct {

	// A list of the budget action resources information.
	//
	// This member is required.
	Actions []types.Action

	// A generic string.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeBudgetActionsForBudgetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeBudgetActionsForBudget{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeBudgetActionsForBudget{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeBudgetActionsForBudget"); err != nil {
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
	if err = addOpDescribeBudgetActionsForBudgetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBudgetActionsForBudget(options.Region), middleware.Before); err != nil {
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

// DescribeBudgetActionsForBudgetAPIClient is a client that implements the
// DescribeBudgetActionsForBudget operation.
type DescribeBudgetActionsForBudgetAPIClient interface {
	DescribeBudgetActionsForBudget(context.Context, *DescribeBudgetActionsForBudgetInput, ...func(*Options)) (*DescribeBudgetActionsForBudgetOutput, error)
}

var _ DescribeBudgetActionsForBudgetAPIClient = (*Client)(nil)

// DescribeBudgetActionsForBudgetPaginatorOptions is the paginator options for
// DescribeBudgetActionsForBudget
type DescribeBudgetActionsForBudgetPaginatorOptions struct {
	// An integer that represents how many entries a paginated response contains. The
	// maximum is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeBudgetActionsForBudgetPaginator is a paginator for
// DescribeBudgetActionsForBudget
type DescribeBudgetActionsForBudgetPaginator struct {
	options   DescribeBudgetActionsForBudgetPaginatorOptions
	client    DescribeBudgetActionsForBudgetAPIClient
	params    *DescribeBudgetActionsForBudgetInput
	nextToken *string
	firstPage bool
}

// NewDescribeBudgetActionsForBudgetPaginator returns a new
// DescribeBudgetActionsForBudgetPaginator
func NewDescribeBudgetActionsForBudgetPaginator(client DescribeBudgetActionsForBudgetAPIClient, params *DescribeBudgetActionsForBudgetInput, optFns ...func(*DescribeBudgetActionsForBudgetPaginatorOptions)) *DescribeBudgetActionsForBudgetPaginator {
	if params == nil {
		params = &DescribeBudgetActionsForBudgetInput{}
	}

	options := DescribeBudgetActionsForBudgetPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeBudgetActionsForBudgetPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeBudgetActionsForBudgetPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeBudgetActionsForBudget page.
func (p *DescribeBudgetActionsForBudgetPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeBudgetActionsForBudgetOutput, error) {
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

	result, err := p.client.DescribeBudgetActionsForBudget(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeBudgetActionsForBudget(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeBudgetActionsForBudget",
	}
}
