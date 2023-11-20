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

// The AcceptQualificationRequest operation approves a Worker's request for a
// Qualification. Only the owner of the Qualification type can grant a
// Qualification request for that type. A successful request for the
// AcceptQualificationRequest operation returns with no errors and an empty body.
func (c *Client) AcceptQualificationRequest(ctx context.Context, params *AcceptQualificationRequestInput, optFns ...func(*Options)) (*AcceptQualificationRequestOutput, error) {
	if params == nil {
		params = &AcceptQualificationRequestInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AcceptQualificationRequest", params, optFns, c.addOperationAcceptQualificationRequestMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AcceptQualificationRequestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AcceptQualificationRequestInput struct {

	// The ID of the Qualification request, as returned by the GetQualificationRequests
	// operation.
	//
	// This member is required.
	QualificationRequestId *string

	// The value of the Qualification. You can omit this value if you are using the
	// presence or absence of the Qualification as the basis for a HIT requirement.
	IntegerValue *int32

	noSmithyDocumentSerde
}

type AcceptQualificationRequestOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAcceptQualificationRequestMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAcceptQualificationRequest{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAcceptQualificationRequest{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AcceptQualificationRequest"); err != nil {
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
	if err = addOpAcceptQualificationRequestValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAcceptQualificationRequest(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAcceptQualificationRequest(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AcceptQualificationRequest",
	}
}
