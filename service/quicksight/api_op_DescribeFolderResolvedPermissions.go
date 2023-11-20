// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the folder resolved permissions. Permissions consists of both folder
// direct permissions and the inherited permissions from the ancestor folders.
func (c *Client) DescribeFolderResolvedPermissions(ctx context.Context, params *DescribeFolderResolvedPermissionsInput, optFns ...func(*Options)) (*DescribeFolderResolvedPermissionsOutput, error) {
	if params == nil {
		params = &DescribeFolderResolvedPermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFolderResolvedPermissions", params, optFns, c.addOperationDescribeFolderResolvedPermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFolderResolvedPermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFolderResolvedPermissionsInput struct {

	// The ID for the Amazon Web Services account that contains the folder.
	//
	// This member is required.
	AwsAccountId *string

	// The ID of the folder.
	//
	// This member is required.
	FolderId *string

	// The maximum number of results to be returned per request.
	MaxResults *int32

	// The namespace of the folder whose permissions you want described.
	Namespace *string

	// A pagination token for the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeFolderResolvedPermissionsOutput struct {

	// The Amazon Resource Name (ARN) of the folder.
	Arn *string

	// The ID of the folder.
	FolderId *string

	// A pagination token for the next set of results, or null if there are no more
	// results.
	NextToken *string

	// Information about the permissions for the folder.
	Permissions []types.ResourcePermission

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFolderResolvedPermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeFolderResolvedPermissions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeFolderResolvedPermissions{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeFolderResolvedPermissions"); err != nil {
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
	if err = addOpDescribeFolderResolvedPermissionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFolderResolvedPermissions(options.Region), middleware.Before); err != nil {
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

// DescribeFolderResolvedPermissionsAPIClient is a client that implements the
// DescribeFolderResolvedPermissions operation.
type DescribeFolderResolvedPermissionsAPIClient interface {
	DescribeFolderResolvedPermissions(context.Context, *DescribeFolderResolvedPermissionsInput, ...func(*Options)) (*DescribeFolderResolvedPermissionsOutput, error)
}

var _ DescribeFolderResolvedPermissionsAPIClient = (*Client)(nil)

// DescribeFolderResolvedPermissionsPaginatorOptions is the paginator options for
// DescribeFolderResolvedPermissions
type DescribeFolderResolvedPermissionsPaginatorOptions struct {
	// The maximum number of results to be returned per request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeFolderResolvedPermissionsPaginator is a paginator for
// DescribeFolderResolvedPermissions
type DescribeFolderResolvedPermissionsPaginator struct {
	options   DescribeFolderResolvedPermissionsPaginatorOptions
	client    DescribeFolderResolvedPermissionsAPIClient
	params    *DescribeFolderResolvedPermissionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeFolderResolvedPermissionsPaginator returns a new
// DescribeFolderResolvedPermissionsPaginator
func NewDescribeFolderResolvedPermissionsPaginator(client DescribeFolderResolvedPermissionsAPIClient, params *DescribeFolderResolvedPermissionsInput, optFns ...func(*DescribeFolderResolvedPermissionsPaginatorOptions)) *DescribeFolderResolvedPermissionsPaginator {
	if params == nil {
		params = &DescribeFolderResolvedPermissionsInput{}
	}

	options := DescribeFolderResolvedPermissionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeFolderResolvedPermissionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeFolderResolvedPermissionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeFolderResolvedPermissions page.
func (p *DescribeFolderResolvedPermissionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeFolderResolvedPermissionsOutput, error) {
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

	result, err := p.client.DescribeFolderResolvedPermissions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeFolderResolvedPermissions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeFolderResolvedPermissions",
	}
}
