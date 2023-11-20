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
)

// This API operation initiates a Face Liveness session. It returns a SessionId ,
// which you can use to start streaming Face Liveness video and get the results for
// a Face Liveness session. You can use the OutputConfig option in the Settings
// parameter to provide an Amazon S3 bucket location. The Amazon S3 bucket stores
// reference images and audit images. If no Amazon S3 bucket is defined, raw bytes
// are sent instead. You can use AuditImagesLimit to limit the number of audit
// images returned when GetFaceLivenessSessionResults is called. This number is
// between 0 and 4. By default, it is set to 0. The limit is best effort and based
// on the duration of the selfie-video.
func (c *Client) CreateFaceLivenessSession(ctx context.Context, params *CreateFaceLivenessSessionInput, optFns ...func(*Options)) (*CreateFaceLivenessSessionOutput, error) {
	if params == nil {
		params = &CreateFaceLivenessSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFaceLivenessSession", params, optFns, c.addOperationCreateFaceLivenessSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFaceLivenessSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFaceLivenessSessionInput struct {

	// Idempotent token is used to recognize the Face Liveness request. If the same
	// token is used with multiple CreateFaceLivenessSession requests, the same
	// session is returned. This token is employed to avoid unintentionally creating
	// the same session multiple times.
	ClientRequestToken *string

	// The identifier for your AWS Key Management Service key (AWS KMS key). Used to
	// encrypt audit images and reference images.
	KmsKeyId *string

	// A session settings object. It contains settings for the operation to be
	// performed. For Face Liveness, it accepts OutputConfig and AuditImagesLimit .
	Settings *types.CreateFaceLivenessSessionRequestSettings

	noSmithyDocumentSerde
}

type CreateFaceLivenessSessionOutput struct {

	// A unique 128-bit UUID identifying a Face Liveness session.
	//
	// This member is required.
	SessionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateFaceLivenessSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateFaceLivenessSession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateFaceLivenessSession{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateFaceLivenessSession"); err != nil {
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
	if err = addOpCreateFaceLivenessSessionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFaceLivenessSession(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateFaceLivenessSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateFaceLivenessSession",
	}
}
