// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all of the smart home appliances associated with a room.
//
// Deprecated: Alexa For Business is no longer supported
func (c *Client) ListSmartHomeAppliances(ctx context.Context, params *ListSmartHomeAppliancesInput, optFns ...func(*Options)) (*ListSmartHomeAppliancesOutput, error) {
	if params == nil {
		params = &ListSmartHomeAppliancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSmartHomeAppliances", params, optFns, c.addOperationListSmartHomeAppliancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSmartHomeAppliancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSmartHomeAppliancesInput struct {

	// The room that the appliances are associated with.
	//
	// This member is required.
	RoomArn *string

	// The maximum number of appliances to be returned, per paginated calls.
	MaxResults *int32

	// The tokens used for pagination.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSmartHomeAppliancesOutput struct {

	// The tokens used for pagination.
	NextToken *string

	// The smart home appliances.
	SmartHomeAppliances []types.SmartHomeAppliance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSmartHomeAppliancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSmartHomeAppliances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSmartHomeAppliances{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ListSmartHomeAppliances"); err != nil {
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
	if err = addOpListSmartHomeAppliancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSmartHomeAppliances(options.Region), middleware.Before); err != nil {
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

// ListSmartHomeAppliancesAPIClient is a client that implements the
// ListSmartHomeAppliances operation.
type ListSmartHomeAppliancesAPIClient interface {
	ListSmartHomeAppliances(context.Context, *ListSmartHomeAppliancesInput, ...func(*Options)) (*ListSmartHomeAppliancesOutput, error)
}

var _ ListSmartHomeAppliancesAPIClient = (*Client)(nil)

// ListSmartHomeAppliancesPaginatorOptions is the paginator options for
// ListSmartHomeAppliances
type ListSmartHomeAppliancesPaginatorOptions struct {
	// The maximum number of appliances to be returned, per paginated calls.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSmartHomeAppliancesPaginator is a paginator for ListSmartHomeAppliances
type ListSmartHomeAppliancesPaginator struct {
	options   ListSmartHomeAppliancesPaginatorOptions
	client    ListSmartHomeAppliancesAPIClient
	params    *ListSmartHomeAppliancesInput
	nextToken *string
	firstPage bool
}

// NewListSmartHomeAppliancesPaginator returns a new
// ListSmartHomeAppliancesPaginator
func NewListSmartHomeAppliancesPaginator(client ListSmartHomeAppliancesAPIClient, params *ListSmartHomeAppliancesInput, optFns ...func(*ListSmartHomeAppliancesPaginatorOptions)) *ListSmartHomeAppliancesPaginator {
	if params == nil {
		params = &ListSmartHomeAppliancesInput{}
	}

	options := ListSmartHomeAppliancesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSmartHomeAppliancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSmartHomeAppliancesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSmartHomeAppliances page.
func (p *ListSmartHomeAppliancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSmartHomeAppliancesOutput, error) {
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

	result, err := p.client.ListSmartHomeAppliances(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSmartHomeAppliances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ListSmartHomeAppliances",
	}
}
