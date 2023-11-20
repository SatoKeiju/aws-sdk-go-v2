// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The GetFileUploadURL operation generates and returns a temporary URL. You use
// the temporary URL to retrieve a file uploaded by a Worker as an answer to a
// FileUploadAnswer question for a HIT. The temporary URL is generated the instant
// the GetFileUploadURL operation is called, and is valid for 60 seconds. You can
// get a temporary file upload URL any time until the HIT is disposed. After the
// HIT is disposed, any uploaded files are deleted, and cannot be retrieved.
// Pending Deprecation on December 12, 2017. The Answer Specification
//
// structure will no longer support the FileUploadAnswer element to be used for
// the QuestionForm data structure. Instead, we recommend that Requesters who want
// to create HITs asking Workers to upload files to use Amazon S3.
func (c *Client) GetFileUploadURL(ctx context.Context, params *GetFileUploadURLInput, optFns ...func(*Options)) (*GetFileUploadURLOutput, error) {
	if params == nil {
		params = &GetFileUploadURLInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFileUploadURL", params, optFns, c.addOperationGetFileUploadURLMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFileUploadURLOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFileUploadURLInput struct {

	// The ID of the assignment that contains the question with a FileUploadAnswer.
	//
	// This member is required.
	AssignmentId *string

	// The identifier of the question with a FileUploadAnswer, as specified in the
	// QuestionForm of the HIT.
	//
	// This member is required.
	QuestionIdentifier *string

	noSmithyDocumentSerde
}

type GetFileUploadURLOutput struct {

	// A temporary URL for the file that the Worker uploaded for the answer.
	FileUploadURL *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetFileUploadURLMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetFileUploadURL{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetFileUploadURL{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetFileUploadURL"); err != nil {
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
	if err = addOpGetFileUploadURLValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFileUploadURL(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetFileUploadURL(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetFileUploadURL",
	}
}
