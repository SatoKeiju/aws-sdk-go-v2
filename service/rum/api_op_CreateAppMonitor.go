// Code generated by smithy-go-codegen DO NOT EDIT.

package rum

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rum/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a Amazon CloudWatch RUM app monitor, which collects telemetry data from
// your application and sends that data to RUM. The data includes performance and
// reliability information such as page load time, client-side errors, and user
// behavior. You use this operation only to create a new app monitor. To update an
// existing app monitor, use UpdateAppMonitor (https://docs.aws.amazon.com/cloudwatchrum/latest/APIReference/API_UpdateAppMonitor.html)
// instead. After you create an app monitor, sign in to the CloudWatch RUM console
// to get the JavaScript code snippet to add to your web application. For more
// information, see How do I find a code snippet that I've already generated? (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-find-code-snippet.html)
func (c *Client) CreateAppMonitor(ctx context.Context, params *CreateAppMonitorInput, optFns ...func(*Options)) (*CreateAppMonitorOutput, error) {
	if params == nil {
		params = &CreateAppMonitorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAppMonitor", params, optFns, c.addOperationCreateAppMonitorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAppMonitorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAppMonitorInput struct {

	// The top-level internet domain name for which your application has
	// administrative authority.
	//
	// This member is required.
	Domain *string

	// A name for the app monitor.
	//
	// This member is required.
	Name *string

	// A structure that contains much of the configuration data for the app monitor.
	// If you are using Amazon Cognito for authorization, you must include this
	// structure in your request, and it must include the ID of the Amazon Cognito
	// identity pool to use for authorization. If you don't include
	// AppMonitorConfiguration , you must set up your own authorization method. For
	// more information, see Authorize your application to send data to Amazon Web
	// Services (https://docs.aws.amazon.com/monitoring/CloudWatch-RUM-get-started-authorization.html)
	// . If you omit this argument, the sample rate used for RUM is set to 10% of the
	// user sessions.
	AppMonitorConfiguration *types.AppMonitorConfiguration

	// Specifies whether this app monitor allows the web client to define and send
	// custom events. If you omit this parameter, custom events are DISABLED . For more
	// information about custom events, see Send custom events (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch-RUM-custom-events.html)
	// .
	CustomEvents *types.CustomEvents

	// Data collected by RUM is kept by RUM for 30 days and then deleted. This
	// parameter specifies whether RUM sends a copy of this telemetry data to Amazon
	// CloudWatch Logs in your account. This enables you to keep the telemetry data for
	// more than 30 days, but it does incur Amazon CloudWatch Logs charges. If you omit
	// this parameter, the default is false .
	CwLogEnabled *bool

	// Assigns one or more tags (key-value pairs) to the app monitor. Tags can help
	// you organize and categorize your resources. You can also use them to scope user
	// permissions by granting a user permission to access or change only resources
	// with certain tag values. Tags don't have any semantic meaning to Amazon Web
	// Services and are interpreted strictly as strings of characters. You can
	// associate as many as 50 tags with an app monitor. For more information, see
	// Tagging Amazon Web Services resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
	// .
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateAppMonitorOutput struct {

	// The unique ID of the new app monitor.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAppMonitorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAppMonitor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAppMonitor{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateAppMonitor"); err != nil {
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
	if err = addOpCreateAppMonitorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAppMonitor(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAppMonitor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateAppMonitor",
	}
}
