// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides information about a stream processor created by CreateStreamProcessor .
// You can get information about the input and output streams, the input parameters
// for the face recognition being performed, and the current status of the stream
// processor.
func (c *Client) DescribeStreamProcessor(ctx context.Context, params *DescribeStreamProcessorInput, optFns ...func(*Options)) (*DescribeStreamProcessorOutput, error) {
	if params == nil {
		params = &DescribeStreamProcessorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStreamProcessor", params, optFns, c.addOperationDescribeStreamProcessorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStreamProcessorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStreamProcessorInput struct {

	// Name of the stream processor for which you want information.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type DescribeStreamProcessorOutput struct {

	// Date and time the stream processor was created
	CreationTimestamp *time.Time

	// Shows whether you are sharing data with Rekognition to improve model
	// performance. You can choose this option at the account level or on a per-stream
	// basis. Note that if you opt out at the account level this setting is ignored on
	// individual streams.
	DataSharingPreference *types.StreamProcessorDataSharingPreference

	// Kinesis video stream that provides the source streaming video.
	Input *types.StreamProcessorInput

	// The identifier for your AWS Key Management Service key (AWS KMS key). This is
	// an optional parameter for label detection stream processors.
	KmsKeyId *string

	// The time, in Unix format, the stream processor was last updated. For example,
	// when the stream processor moves from a running state to a failed state, or when
	// the user starts or stops the stream processor.
	LastUpdateTimestamp *time.Time

	// Name of the stream processor.
	Name *string

	// The Amazon Simple Notification Service topic to which Amazon Rekognition
	// publishes the object detection results and completion status of a video analysis
	// operation. Amazon Rekognition publishes a notification the first time an object
	// of interest or a person is detected in the video stream. For example, if Amazon
	// Rekognition detects a person at second 2, a pet at second 4, and a person again
	// at second 5, Amazon Rekognition sends 2 object class detected notifications, one
	// for a person at second 2 and one for a pet at second 4. Amazon Rekognition also
	// publishes an an end-of-session notification with a summary when the stream
	// processing session is complete.
	NotificationChannel *types.StreamProcessorNotificationChannel

	// Kinesis data stream to which Amazon Rekognition Video puts the analysis results.
	Output *types.StreamProcessorOutput

	// Specifies locations in the frames where Amazon Rekognition checks for objects
	// or people. This is an optional parameter for label detection stream processors.
	RegionsOfInterest []types.RegionOfInterest

	// ARN of the IAM role that allows access to the stream processor.
	RoleArn *string

	// Input parameters used in a streaming video analyzed by a stream processor. You
	// can use FaceSearch to recognize faces in a streaming video, or you can use
	// ConnectedHome to detect labels.
	Settings *types.StreamProcessorSettings

	// Current status of the stream processor.
	Status types.StreamProcessorStatus

	// Detailed status message about the stream processor.
	StatusMessage *string

	// ARN of the stream processor.
	StreamProcessorArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeStreamProcessorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeStreamProcessor{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeStreamProcessor{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DescribeStreamProcessor"); err != nil {
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
	if err = addOpDescribeStreamProcessorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStreamProcessor(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeStreamProcessor(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DescribeStreamProcessor",
	}
}
