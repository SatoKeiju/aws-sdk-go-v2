// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new version of the specified managed policy. To update a managed
// policy, you create a new policy version. A managed policy can have up to five
// versions. If the policy has five versions, you must delete an existing version
// using DeletePolicyVersion before you create a new version. Optionally, you can
// set the new version as the policy's default version. The default version is the
// version that is in effect for the IAM users, groups, and roles to which the
// policy is attached. For more information about managed policy versions, see
// Versioning for managed policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-versions.html)
// in the IAM User Guide.
func (c *Client) CreatePolicyVersion(ctx context.Context, params *CreatePolicyVersionInput, optFns ...func(*Options)) (*CreatePolicyVersionOutput, error) {
	if params == nil {
		params = &CreatePolicyVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePolicyVersion", params, optFns, c.addOperationCreatePolicyVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePolicyVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePolicyVersionInput struct {

	// The Amazon Resource Name (ARN) of the IAM policy to which you want to add a new
	// version. For more information about ARNs, see Amazon Resource Names (ARNs) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the Amazon Web Services General Reference.
	//
	// This member is required.
	PolicyArn *string

	// The JSON policy document that you want to use as the content for this new
	// version of the policy. You must provide policies in JSON format in IAM. However,
	// for CloudFormation templates formatted in YAML, you can provide the policy in
	// JSON or YAML format. CloudFormation always converts a YAML policy to JSON format
	// before submitting it to IAM. The maximum length of the policy document that you
	// can pass in this operation, including whitespace, is listed below. To view the
	// maximum character counts of a managed policy with no whitespaces, see IAM and
	// STS character quotas (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html#reference_iam-quotas-entity-length)
	// . The regex pattern (http://wikipedia.org/wiki/regex) used to validate this
	// parameter is a string of characters consisting of the following:
	//   - Any printable ASCII character ranging from the space character ( \u0020 )
	//   through the end of the ASCII character range
	//   - The printable characters in the Basic Latin and Latin-1 Supplement
	//   character set (through \u00FF )
	//   - The special characters tab ( \u0009 ), line feed ( \u000A ), and carriage
	//   return ( \u000D )
	//
	// This member is required.
	PolicyDocument *string

	// Specifies whether to set this version as the policy's default version. When
	// this parameter is true , the new policy version becomes the operative version.
	// That is, it becomes the version that is in effect for the IAM users, groups, and
	// roles that the policy is attached to. For more information about managed policy
	// versions, see Versioning for managed policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-versions.html)
	// in the IAM User Guide.
	SetAsDefault bool

	noSmithyDocumentSerde
}

// Contains the response to a successful CreatePolicyVersion request.
type CreatePolicyVersionOutput struct {

	// A structure containing details about the new policy version.
	PolicyVersion *types.PolicyVersion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePolicyVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreatePolicyVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreatePolicyVersion{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePolicyVersion"); err != nil {
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
	if err = addOpCreatePolicyVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePolicyVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePolicyVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePolicyVersion",
	}
}
