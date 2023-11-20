// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieve inbox placement and engagement rates for the domains that you use to
// send email.
func (c *Client) GetDomainStatisticsReport(ctx context.Context, params *GetDomainStatisticsReportInput, optFns ...func(*Options)) (*GetDomainStatisticsReportOutput, error) {
	if params == nil {
		params = &GetDomainStatisticsReportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDomainStatisticsReport", params, optFns, c.addOperationGetDomainStatisticsReportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDomainStatisticsReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to obtain deliverability metrics for a domain.
type GetDomainStatisticsReportInput struct {

	// The domain that you want to obtain deliverability metrics for.
	//
	// This member is required.
	Domain *string

	// The last day (in Unix time) that you want to obtain domain deliverability
	// metrics for. The EndDate that you specify has to be less than or equal to 30
	// days after the StartDate .
	//
	// This member is required.
	EndDate *time.Time

	// The first day (in Unix time) that you want to obtain domain deliverability
	// metrics for.
	//
	// This member is required.
	StartDate *time.Time

	noSmithyDocumentSerde
}

// An object that includes statistics that are related to the domain that you
// specified.
type GetDomainStatisticsReportOutput struct {

	// An object that contains deliverability metrics for the domain that you
	// specified. This object contains data for each day, starting on the StartDate
	// and ending on the EndDate .
	//
	// This member is required.
	DailyVolumes []types.DailyVolume

	// An object that contains deliverability metrics for the domain that you
	// specified. The data in this object is a summary of all of the data that was
	// collected from the StartDate to the EndDate .
	//
	// This member is required.
	OverallVolume *types.OverallVolume

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDomainStatisticsReportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDomainStatisticsReport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDomainStatisticsReport{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetDomainStatisticsReport"); err != nil {
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
	if err = addOpGetDomainStatisticsReportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDomainStatisticsReport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDomainStatisticsReport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetDomainStatisticsReport",
	}
}
