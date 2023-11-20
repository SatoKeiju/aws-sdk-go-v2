// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Checks whether occurrences of sensitive data can be retrieved for a finding.
func (c *Client) GetSensitiveDataOccurrencesAvailability(ctx context.Context, params *GetSensitiveDataOccurrencesAvailabilityInput, optFns ...func(*Options)) (*GetSensitiveDataOccurrencesAvailabilityOutput, error) {
	if params == nil {
		params = &GetSensitiveDataOccurrencesAvailabilityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSensitiveDataOccurrencesAvailability", params, optFns, c.addOperationGetSensitiveDataOccurrencesAvailabilityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSensitiveDataOccurrencesAvailabilityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSensitiveDataOccurrencesAvailabilityInput struct {

	// The unique identifier for the finding.
	//
	// This member is required.
	FindingId *string

	noSmithyDocumentSerde
}

type GetSensitiveDataOccurrencesAvailabilityOutput struct {

	// Specifies whether occurrences of sensitive data can be retrieved for the
	// finding. Possible values are: AVAILABLE, the sensitive data can be retrieved;
	// and, UNAVAILABLE, the sensitive data can't be retrieved. If this value is
	// UNAVAILABLE, the reasons array indicates why the data can't be retrieved.
	Code types.AvailabilityCode

	// Specifies why occurrences of sensitive data can't be retrieved for the finding.
	// Possible values are:
	//   - ACCOUNT_NOT_IN_ORGANIZATION - The affected account isn't currently part of
	//   your organization. Or the account is part of your organization but Macie isn't
	//   currently enabled for the account. You're not allowed to access the affected S3
	//   object by using Macie.
	//   - INVALID_CLASSIFICATION_RESULT - There isn't a corresponding sensitive data
	//   discovery result for the finding. Or the corresponding sensitive data discovery
	//   result isn't available, is malformed or corrupted, or uses an unsupported
	//   storage format. Macie can't verify the location of the sensitive data to
	//   retrieve.
	//   - INVALID_RESULT_SIGNATURE - The corresponding sensitive data discovery
	//   result is stored in an S3 object that wasn't signed by Macie. Macie can't verify
	//   the integrity and authenticity of the sensitive data discovery result.
	//   Therefore, Macie can't verify the location of the sensitive data to retrieve.
	//   - MEMBER_ROLE_TOO_PERMISSIVE - The affected member account is configured to
	//   retrieve occurrences of sensitive data by using an IAM role whose trust or
	//   permissions policy doesn't meet Macie requirements for restricting access to the
	//   role. Or the role's trust policy doesn't specify the correct external ID. Macie
	//   can't assume the role to retrieve the sensitive data.
	//   - MISSING_GET_MEMBER_PERMISSION - You're not allowed to retrieve information
	//   about the association between your account and the affected account. Macie can't
	//   determine whether you’re allowed to access the affected S3 object as the
	//   delegated Macie administrator for the affected account.
	//   - OBJECT_EXCEEDS_SIZE_QUOTA - The storage size of the affected S3 object
	//   exceeds the size quota for retrieving occurrences of sensitive data from this
	//   type of file.
	//   - OBJECT_UNAVAILABLE - The affected S3 object isn't available. The object was
	//   renamed, moved, or deleted. Or the object was changed after Macie created the
	//   finding.
	//   - RESULT_NOT_SIGNED - The corresponding sensitive data discovery result is
	//   stored in an S3 object that hasn't been signed. Macie can't verify the integrity
	//   and authenticity of the sensitive data discovery result. Therefore, Macie can't
	//   verify the location of the sensitive data to retrieve.
	//   - ROLE_TOO_PERMISSIVE - Your account is configured to retrieve occurrences of
	//   sensitive data by using an IAM role whose trust or permissions policy doesn't
	//   meet Macie requirements for restricting access to the role. Macie can’t assume
	//   the role to retrieve the sensitive data.
	//   - UNSUPPORTED_FINDING_TYPE - The specified finding isn't a sensitive data
	//   finding.
	//   - UNSUPPORTED_OBJECT_TYPE - The affected S3 object uses a file or storage
	//   format that Macie doesn't support for retrieving occurrences of sensitive data.
	// This value is null if sensitive data can be retrieved for the finding.
	Reasons []types.UnavailabilityReasonCode

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSensitiveDataOccurrencesAvailabilityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSensitiveDataOccurrencesAvailability{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSensitiveDataOccurrencesAvailability{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetSensitiveDataOccurrencesAvailability"); err != nil {
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
	if err = addOpGetSensitiveDataOccurrencesAvailabilityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSensitiveDataOccurrencesAvailability(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSensitiveDataOccurrencesAvailability(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetSensitiveDataOccurrencesAvailability",
	}
}
