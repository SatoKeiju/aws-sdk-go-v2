// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the evidence folders from a specified assessment in Audit Manager.
func (c *Client) GetEvidenceFoldersByAssessment(ctx context.Context, params *GetEvidenceFoldersByAssessmentInput, optFns ...func(*Options)) (*GetEvidenceFoldersByAssessmentOutput, error) {
	if params == nil {
		params = &GetEvidenceFoldersByAssessmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEvidenceFoldersByAssessment", params, optFns, c.addOperationGetEvidenceFoldersByAssessmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEvidenceFoldersByAssessmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEvidenceFoldersByAssessmentInput struct {

	// The unique identifier for the assessment.
	//
	// This member is required.
	AssessmentId *string

	// Represents the maximum number of results on a page or for an API request call.
	MaxResults *int32

	// The pagination token that's used to fetch the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetEvidenceFoldersByAssessmentOutput struct {

	// The list of evidence folders that the GetEvidenceFoldersByAssessment API
	// returned.
	EvidenceFolders []types.AssessmentEvidenceFolder

	// The pagination token that's used to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEvidenceFoldersByAssessmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetEvidenceFoldersByAssessment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetEvidenceFoldersByAssessment{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetEvidenceFoldersByAssessment"); err != nil {
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
	if err = addOpGetEvidenceFoldersByAssessmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEvidenceFoldersByAssessment(options.Region), middleware.Before); err != nil {
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

// GetEvidenceFoldersByAssessmentAPIClient is a client that implements the
// GetEvidenceFoldersByAssessment operation.
type GetEvidenceFoldersByAssessmentAPIClient interface {
	GetEvidenceFoldersByAssessment(context.Context, *GetEvidenceFoldersByAssessmentInput, ...func(*Options)) (*GetEvidenceFoldersByAssessmentOutput, error)
}

var _ GetEvidenceFoldersByAssessmentAPIClient = (*Client)(nil)

// GetEvidenceFoldersByAssessmentPaginatorOptions is the paginator options for
// GetEvidenceFoldersByAssessment
type GetEvidenceFoldersByAssessmentPaginatorOptions struct {
	// Represents the maximum number of results on a page or for an API request call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetEvidenceFoldersByAssessmentPaginator is a paginator for
// GetEvidenceFoldersByAssessment
type GetEvidenceFoldersByAssessmentPaginator struct {
	options   GetEvidenceFoldersByAssessmentPaginatorOptions
	client    GetEvidenceFoldersByAssessmentAPIClient
	params    *GetEvidenceFoldersByAssessmentInput
	nextToken *string
	firstPage bool
}

// NewGetEvidenceFoldersByAssessmentPaginator returns a new
// GetEvidenceFoldersByAssessmentPaginator
func NewGetEvidenceFoldersByAssessmentPaginator(client GetEvidenceFoldersByAssessmentAPIClient, params *GetEvidenceFoldersByAssessmentInput, optFns ...func(*GetEvidenceFoldersByAssessmentPaginatorOptions)) *GetEvidenceFoldersByAssessmentPaginator {
	if params == nil {
		params = &GetEvidenceFoldersByAssessmentInput{}
	}

	options := GetEvidenceFoldersByAssessmentPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetEvidenceFoldersByAssessmentPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetEvidenceFoldersByAssessmentPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetEvidenceFoldersByAssessment page.
func (p *GetEvidenceFoldersByAssessmentPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetEvidenceFoldersByAssessmentOutput, error) {
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

	result, err := p.client.GetEvidenceFoldersByAssessment(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetEvidenceFoldersByAssessment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetEvidenceFoldersByAssessment",
	}
}
