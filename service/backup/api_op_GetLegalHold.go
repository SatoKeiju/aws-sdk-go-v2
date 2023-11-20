// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// This action returns details for a specified legal hold. The details are the
// body of a legal hold in JSON format, in addition to metadata.
func (c *Client) GetLegalHold(ctx context.Context, params *GetLegalHoldInput, optFns ...func(*Options)) (*GetLegalHoldOutput, error) {
	if params == nil {
		params = &GetLegalHoldInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLegalHold", params, optFns, c.addOperationGetLegalHoldMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLegalHoldOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLegalHoldInput struct {

	// This is the ID required to use GetLegalHold . This unique ID is associated with
	// a specific legal hold.
	//
	// This member is required.
	LegalHoldId *string

	noSmithyDocumentSerde
}

type GetLegalHoldOutput struct {

	// String describing the reason for removing the legal hold.
	CancelDescription *string

	// Time in number when legal hold was cancelled.
	CancellationDate *time.Time

	// Time in number format when legal hold was created.
	CreationDate *time.Time

	// This is the returned string description of the legal hold.
	Description *string

	// This is the returned framework ARN for the specified legal hold. An Amazon
	// Resource Name (ARN) uniquely identifies a resource. The format of the ARN
	// depends on the resource type.
	LegalHoldArn *string

	// This is the returned ID associated with a specified legal hold.
	LegalHoldId *string

	// This specifies criteria to assign a set of resources, such as resource types or
	// backup vaults.
	RecoveryPointSelection *types.RecoveryPointSelection

	// This is the date and time until which the legal hold record will be retained.
	RetainRecordUntil *time.Time

	// This is the status of the legal hold. Statuses can be ACTIVE , CREATING ,
	// CANCELED , and CANCELING .
	Status types.LegalHoldStatus

	// This is the string title of the legal hold.
	Title *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetLegalHoldMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetLegalHold{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetLegalHold{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetLegalHold"); err != nil {
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
	if err = addOpGetLegalHoldValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLegalHold(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetLegalHold(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetLegalHold",
	}
}
