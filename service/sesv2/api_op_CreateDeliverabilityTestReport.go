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
)

// Create a new predictive inbox placement test. Predictive inbox placement tests
// can help you predict how your messages will be handled by various email
// providers around the world. When you perform a predictive inbox placement test,
// you provide a sample message that contains the content that you plan to send to
// your customers. Amazon SES then sends that message to special email addresses
// spread across several major email providers. After about 24 hours, the test is
// complete, and you can use the GetDeliverabilityTestReport operation to view the
// results of the test.
func (c *Client) CreateDeliverabilityTestReport(ctx context.Context, params *CreateDeliverabilityTestReportInput, optFns ...func(*Options)) (*CreateDeliverabilityTestReportOutput, error) {
	if params == nil {
		params = &CreateDeliverabilityTestReportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDeliverabilityTestReport", params, optFns, c.addOperationCreateDeliverabilityTestReportMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDeliverabilityTestReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to perform a predictive inbox placement test. Predictive inbox
// placement tests can help you predict how your messages will be handled by
// various email providers around the world. When you perform a predictive inbox
// placement test, you provide a sample message that contains the content that you
// plan to send to your customers. We send that message to special email addresses
// spread across several major email providers around the world. The test takes
// about 24 hours to complete. When the test is complete, you can use the
// GetDeliverabilityTestReport operation to view the results of the test.
type CreateDeliverabilityTestReportInput struct {

	// The HTML body of the message that you sent when you performed the predictive
	// inbox placement test.
	//
	// This member is required.
	Content *types.EmailContent

	// The email address that the predictive inbox placement test email was sent from.
	//
	// This member is required.
	FromEmailAddress *string

	// A unique name that helps you to identify the predictive inbox placement test
	// when you retrieve the results.
	ReportName *string

	// An array of objects that define the tags (keys and values) that you want to
	// associate with the predictive inbox placement test.
	Tags []types.Tag

	noSmithyDocumentSerde
}

// Information about the predictive inbox placement test that you created.
type CreateDeliverabilityTestReportOutput struct {

	// The status of the predictive inbox placement test. If the status is IN_PROGRESS
	// , then the predictive inbox placement test is currently running. Predictive
	// inbox placement tests are usually complete within 24 hours of creating the test.
	// If the status is COMPLETE , then the test is finished, and you can use the
	// GetDeliverabilityTestReport to view the results of the test.
	//
	// This member is required.
	DeliverabilityTestStatus types.DeliverabilityTestStatus

	// A unique string that identifies the predictive inbox placement test.
	//
	// This member is required.
	ReportId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDeliverabilityTestReportMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDeliverabilityTestReport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDeliverabilityTestReport{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateDeliverabilityTestReport"); err != nil {
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
	if err = addOpCreateDeliverabilityTestReportValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDeliverabilityTestReport(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDeliverabilityTestReport(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateDeliverabilityTestReport",
	}
}
