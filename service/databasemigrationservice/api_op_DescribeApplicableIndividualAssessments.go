// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides a list of individual assessments that you can specify for a new
// premigration assessment run, given one or more parameters. If you specify an
// existing migration task, this operation provides the default individual
// assessments you can specify for that task. Otherwise, the specified parameters
// model elements of a possible migration task on which to base a premigration
// assessment run. To use these migration task modeling parameters, you must
// specify an existing replication instance, a source database engine, a target
// database engine, and a migration type. This combination of parameters
// potentially limits the default individual assessments available for an
// assessment run created for a corresponding migration task. If you specify no
// parameters, this operation provides a list of all possible individual
// assessments that you can specify for an assessment run. If you specify any one
// of the task modeling parameters, you must specify all of them or the operation
// cannot provide a list of individual assessments. The only parameter that you can
// specify alone is for an existing migration task. The specified task definition
// then determines the default list of individual assessments that you can specify
// in an assessment run for the task.
func (c *Client) DescribeApplicableIndividualAssessments(ctx context.Context, params *DescribeApplicableIndividualAssessmentsInput, optFns ...func(*Options)) (*DescribeApplicableIndividualAssessmentsOutput, error) {
	if params == nil {
		params = &DescribeApplicableIndividualAssessmentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeApplicableIndividualAssessments", params, optFns, c.addOperationDescribeApplicableIndividualAssessmentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeApplicableIndividualAssessmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeApplicableIndividualAssessmentsInput struct {

	// Optional pagination token provided by a previous request. If this parameter is
	// specified, the response includes only records beyond the marker, up to the value
	// specified by MaxRecords .
	Marker *string

	// Maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	MaxRecords *int32

	// Name of the migration type that each provided individual assessment must
	// support.
	MigrationType types.MigrationTypeValue

	// ARN of a replication instance on which you want to base the default list of
	// individual assessments.
	ReplicationInstanceArn *string

	// Amazon Resource Name (ARN) of a migration task on which you want to base the
	// default list of individual assessments.
	ReplicationTaskArn *string

	// Name of a database engine that the specified replication instance supports as a
	// source.
	SourceEngineName *string

	// Name of a database engine that the specified replication instance supports as a
	// target.
	TargetEngineName *string

	noSmithyDocumentSerde
}

type DescribeApplicableIndividualAssessmentsOutput struct {

	// List of names for the individual assessments supported by the premigration
	// assessment run that you start based on the specified request parameters. For
	// more information on the available individual assessments, including
	// compatibility with different migration task configurations, see Working with
	// premigration assessment runs (https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Tasks.AssessmentReport.html)
	// in the Database Migration Service User Guide.
	IndividualAssessmentNames []string

	// Pagination token returned for you to pass to a subsequent request. If you pass
	// this token as the Marker value in a subsequent request, the response includes
	// only records beyond the marker, up to the value specified in the request by
	// MaxRecords .
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeApplicableIndividualAssessmentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeApplicableIndividualAssessments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeApplicableIndividualAssessments{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeApplicableIndividualAssessments"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeApplicableIndividualAssessments(options.Region), middleware.Before); err != nil {
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

// DescribeApplicableIndividualAssessmentsAPIClient is a client that implements
// the DescribeApplicableIndividualAssessments operation.
type DescribeApplicableIndividualAssessmentsAPIClient interface {
	DescribeApplicableIndividualAssessments(context.Context, *DescribeApplicableIndividualAssessmentsInput, ...func(*Options)) (*DescribeApplicableIndividualAssessmentsOutput, error)
}

var _ DescribeApplicableIndividualAssessmentsAPIClient = (*Client)(nil)

// DescribeApplicableIndividualAssessmentsPaginatorOptions is the paginator
// options for DescribeApplicableIndividualAssessments
type DescribeApplicableIndividualAssessmentsPaginatorOptions struct {
	// Maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeApplicableIndividualAssessmentsPaginator is a paginator for
// DescribeApplicableIndividualAssessments
type DescribeApplicableIndividualAssessmentsPaginator struct {
	options   DescribeApplicableIndividualAssessmentsPaginatorOptions
	client    DescribeApplicableIndividualAssessmentsAPIClient
	params    *DescribeApplicableIndividualAssessmentsInput
	nextToken *string
	firstPage bool
}

// NewDescribeApplicableIndividualAssessmentsPaginator returns a new
// DescribeApplicableIndividualAssessmentsPaginator
func NewDescribeApplicableIndividualAssessmentsPaginator(client DescribeApplicableIndividualAssessmentsAPIClient, params *DescribeApplicableIndividualAssessmentsInput, optFns ...func(*DescribeApplicableIndividualAssessmentsPaginatorOptions)) *DescribeApplicableIndividualAssessmentsPaginator {
	if params == nil {
		params = &DescribeApplicableIndividualAssessmentsInput{}
	}

	options := DescribeApplicableIndividualAssessmentsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeApplicableIndividualAssessmentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeApplicableIndividualAssessmentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeApplicableIndividualAssessments page.
func (p *DescribeApplicableIndividualAssessmentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeApplicableIndividualAssessmentsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeApplicableIndividualAssessments(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeApplicableIndividualAssessments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeApplicableIndividualAssessments",
	}
}
