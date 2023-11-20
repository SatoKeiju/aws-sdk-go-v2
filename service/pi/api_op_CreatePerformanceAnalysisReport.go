// Code generated by smithy-go-codegen DO NOT EDIT.

package pi

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pi/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a new performance analysis report for a specific time period for the DB
// instance.
func (c *Client) CreatePerformanceAnalysisReport(ctx context.Context, params *CreatePerformanceAnalysisReportInput, optFns ...func(*Options)) (*CreatePerformanceAnalysisReportOutput, error) {
	if params == nil {
		params = &CreatePerformanceAnalysisReportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePerformanceAnalysisReport", params, optFns, c.addOperationCreatePerformanceAnalysisReportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePerformanceAnalysisReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePerformanceAnalysisReportInput struct {

	// The end time defined for the analysis report.
	//
	// This member is required.
	EndTime *time.Time

	// An immutable, Amazon Web Services Region-unique identifier for a data source.
	// Performance Insights gathers metrics from this data source. To use an Amazon RDS
	// instance as a data source, you specify its DbiResourceId value. For example,
	// specify db-ADECBTYHKTSAUMUZQYPDS2GW4A .
	//
	// This member is required.
	Identifier *string

	// The Amazon Web Services service for which Performance Insights will return
	// metrics. Valid value is RDS .
	//
	// This member is required.
	ServiceType types.ServiceType

	// The start time defined for the analysis report.
	//
	// This member is required.
	StartTime *time.Time

	// The metadata assigned to the analysis report consisting of a key-value pair.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreatePerformanceAnalysisReportOutput struct {

	// A unique identifier for the created analysis report.
	AnalysisReportId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePerformanceAnalysisReportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePerformanceAnalysisReport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePerformanceAnalysisReport{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePerformanceAnalysisReport"); err != nil {
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
	if err = addOpCreatePerformanceAnalysisReportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePerformanceAnalysisReport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePerformanceAnalysisReport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePerformanceAnalysisReport",
	}
}
