// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoverycontrolconfig

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53recoverycontrolconfig/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the safety rules (the assertion rules and gating rules) that you've
// defined for the routing controls in a control panel.
func (c *Client) ListSafetyRules(ctx context.Context, params *ListSafetyRulesInput, optFns ...func(*Options)) (*ListSafetyRulesOutput, error) {
	if params == nil {
		params = &ListSafetyRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSafetyRules", params, optFns, c.addOperationListSafetyRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSafetyRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSafetyRulesInput struct {

	// The Amazon Resource Name (ARN) of the control panel.
	//
	// This member is required.
	ControlPanelArn *string

	// The number of objects that you want to return with this call.
	MaxResults *int32

	// The token that identifies which batch of results you want to see.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSafetyRulesOutput struct {

	// The token that identifies which batch of results you want to see.
	NextToken *string

	// The list of safety rules in a control panel.
	SafetyRules []types.Rule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSafetyRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSafetyRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSafetyRules{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSafetyRules"); err != nil {
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
	if err = addOpListSafetyRulesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSafetyRules(options.Region), middleware.Before); err != nil {
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

// ListSafetyRulesAPIClient is a client that implements the ListSafetyRules
// operation.
type ListSafetyRulesAPIClient interface {
	ListSafetyRules(context.Context, *ListSafetyRulesInput, ...func(*Options)) (*ListSafetyRulesOutput, error)
}

var _ ListSafetyRulesAPIClient = (*Client)(nil)

// ListSafetyRulesPaginatorOptions is the paginator options for ListSafetyRules
type ListSafetyRulesPaginatorOptions struct {
	// The number of objects that you want to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSafetyRulesPaginator is a paginator for ListSafetyRules
type ListSafetyRulesPaginator struct {
	options   ListSafetyRulesPaginatorOptions
	client    ListSafetyRulesAPIClient
	params    *ListSafetyRulesInput
	nextToken *string
	firstPage bool
}

// NewListSafetyRulesPaginator returns a new ListSafetyRulesPaginator
func NewListSafetyRulesPaginator(client ListSafetyRulesAPIClient, params *ListSafetyRulesInput, optFns ...func(*ListSafetyRulesPaginatorOptions)) *ListSafetyRulesPaginator {
	if params == nil {
		params = &ListSafetyRulesInput{}
	}

	options := ListSafetyRulesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSafetyRulesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSafetyRulesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSafetyRules page.
func (p *ListSafetyRulesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSafetyRulesOutput, error) {
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

	result, err := p.client.ListSafetyRules(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSafetyRules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSafetyRules",
	}
}
