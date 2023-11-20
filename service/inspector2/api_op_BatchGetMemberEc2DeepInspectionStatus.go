// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves Amazon Inspector deep inspection activation status of multiple member
// accounts within your organization. You must be the delegated administrator of an
// organization in Amazon Inspector to use this API.
func (c *Client) BatchGetMemberEc2DeepInspectionStatus(ctx context.Context, params *BatchGetMemberEc2DeepInspectionStatusInput, optFns ...func(*Options)) (*BatchGetMemberEc2DeepInspectionStatusOutput, error) {
	if params == nil {
		params = &BatchGetMemberEc2DeepInspectionStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetMemberEc2DeepInspectionStatus", params, optFns, c.addOperationBatchGetMemberEc2DeepInspectionStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetMemberEc2DeepInspectionStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetMemberEc2DeepInspectionStatusInput struct {

	// The unique identifiers for the Amazon Web Services accounts to retrieve Amazon
	// Inspector deep inspection activation status for.
	AccountIds []string

	noSmithyDocumentSerde
}

type BatchGetMemberEc2DeepInspectionStatusOutput struct {

	// An array of objects that provide details on the activation status of Amazon
	// Inspector deep inspection for each of the requested accounts.
	AccountIds []types.MemberAccountEc2DeepInspectionStatusState

	// An array of objects that provide details on any accounts that failed to
	// activate Amazon Inspector deep inspection and why.
	FailedAccountIds []types.FailedMemberAccountEc2DeepInspectionStatusState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetMemberEc2DeepInspectionStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchGetMemberEc2DeepInspectionStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchGetMemberEc2DeepInspectionStatus{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchGetMemberEc2DeepInspectionStatus"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetMemberEc2DeepInspectionStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetMemberEc2DeepInspectionStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchGetMemberEc2DeepInspectionStatus",
	}
}
