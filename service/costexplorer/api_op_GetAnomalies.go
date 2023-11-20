// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves all of the cost anomalies detected on your account during the time
// period that's specified by the DateInterval object. Anomalies are available for
// up to 90 days.
func (c *Client) GetAnomalies(ctx context.Context, params *GetAnomaliesInput, optFns ...func(*Options)) (*GetAnomaliesOutput, error) {
	if params == nil {
		params = &GetAnomaliesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAnomalies", params, optFns, c.addOperationGetAnomaliesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAnomaliesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAnomaliesInput struct {

	// Assigns the start and end dates for retrieving cost anomalies. The returned
	// anomaly object will have an AnomalyEndDate in the specified time range.
	//
	// This member is required.
	DateInterval *types.AnomalyDateInterval

	// Filters anomaly results by the feedback field on the anomaly object.
	Feedback types.AnomalyFeedbackType

	// The number of entries a paginated response contains.
	MaxResults *int32

	// Retrieves all of the cost anomalies detected for a specific cost anomaly
	// monitor Amazon Resource Name (ARN).
	MonitorArn *string

	// The token to retrieve the next set of results. Amazon Web Services provides the
	// token when the response from a previous call has more results than the maximum
	// page size.
	NextPageToken *string

	// Filters anomaly results by the total impact field on the anomaly object. For
	// example, you can filter anomalies GREATER_THAN 200.00 to retrieve anomalies,
	// with an estimated dollar impact greater than 200.
	TotalImpact *types.TotalImpactFilter

	noSmithyDocumentSerde
}

type GetAnomaliesOutput struct {

	// A list of cost anomalies.
	//
	// This member is required.
	Anomalies []types.Anomaly

	// The token to retrieve the next set of results. Amazon Web Services provides the
	// token when the response from a previous call has more results than the maximum
	// page size.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAnomaliesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetAnomalies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAnomalies{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetAnomalies"); err != nil {
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
	if err = addOpGetAnomaliesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAnomalies(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAnomalies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetAnomalies",
	}
}
