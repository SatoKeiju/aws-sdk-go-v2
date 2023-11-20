// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworkscm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes backups. The results are ordered by time, with newest backups first.
// If you do not specify a BackupId or ServerName, the command returns all backups.
// This operation is synchronous. A ResourceNotFoundException is thrown when the
// backup does not exist. A ValidationException is raised when parameters of the
// request are not valid.
func (c *Client) DescribeBackups(ctx context.Context, params *DescribeBackupsInput, optFns ...func(*Options)) (*DescribeBackupsOutput, error) {
	if params == nil {
		params = &DescribeBackupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBackups", params, optFns, c.addOperationDescribeBackupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBackupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBackupsInput struct {

	// Describes a single backup.
	BackupId *string

	// This is not currently implemented for DescribeBackups requests.
	MaxResults *int32

	// This is not currently implemented for DescribeBackups requests.
	NextToken *string

	// Returns backups for the server with the specified ServerName.
	ServerName *string

	noSmithyDocumentSerde
}

type DescribeBackupsOutput struct {

	// Contains the response to a DescribeBackups request.
	Backups []types.Backup

	// This is not currently implemented for DescribeBackups requests.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeBackupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeBackups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeBackups{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeBackups"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBackups(options.Region), middleware.Before); err != nil {
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

// DescribeBackupsAPIClient is a client that implements the DescribeBackups
// operation.
type DescribeBackupsAPIClient interface {
	DescribeBackups(context.Context, *DescribeBackupsInput, ...func(*Options)) (*DescribeBackupsOutput, error)
}

var _ DescribeBackupsAPIClient = (*Client)(nil)

// DescribeBackupsPaginatorOptions is the paginator options for DescribeBackups
type DescribeBackupsPaginatorOptions struct {
	// This is not currently implemented for DescribeBackups requests.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeBackupsPaginator is a paginator for DescribeBackups
type DescribeBackupsPaginator struct {
	options   DescribeBackupsPaginatorOptions
	client    DescribeBackupsAPIClient
	params    *DescribeBackupsInput
	nextToken *string
	firstPage bool
}

// NewDescribeBackupsPaginator returns a new DescribeBackupsPaginator
func NewDescribeBackupsPaginator(client DescribeBackupsAPIClient, params *DescribeBackupsInput, optFns ...func(*DescribeBackupsPaginatorOptions)) *DescribeBackupsPaginator {
	if params == nil {
		params = &DescribeBackupsInput{}
	}

	options := DescribeBackupsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeBackupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeBackupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeBackups page.
func (p *DescribeBackupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeBackupsOutput, error) {
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

	result, err := p.client.DescribeBackups(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeBackups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeBackups",
	}
}
