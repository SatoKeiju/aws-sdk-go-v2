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

// Creates a new managed policy for your Amazon Web Services account. This
// operation creates a policy version with a version identifier of v1 and sets v1
// as the policy's default version. For more information about policy versions, see
// Versioning for managed policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-versions.html)
// in the IAM User Guide. As a best practice, you can validate your IAM policies.
// To learn more, see Validating IAM policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_policy-validator.html)
// in the IAM User Guide. For more information about managed policies in general,
// see Managed policies and inline policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
func (c *Client) CreatePolicy(ctx context.Context, params *CreatePolicyInput, optFns ...func(*Options)) (*CreatePolicyOutput, error) {
	if params == nil {
		params = &CreatePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePolicy", params, optFns, c.addOperationCreatePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePolicyInput struct {

	// The JSON policy document that you want to use as the content for the new
	// policy. You must provide policies in JSON format in IAM. However, for
	// CloudFormation templates formatted in YAML, you can provide the policy in JSON
	// or YAML format. CloudFormation always converts a YAML policy to JSON format
	// before submitting it to IAM. The maximum length of the policy document that you
	// can pass in this operation, including whitespace, is listed below. To view the
	// maximum character counts of a managed policy with no whitespaces, see IAM and
	// STS character quotas (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html#reference_iam-quotas-entity-length)
	// . To learn more about JSON policy grammar, see Grammar of the IAM JSON policy
	// language (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_grammar.html)
	// in the IAM User Guide. The regex pattern (http://wikipedia.org/wiki/regex) used
	// to validate this parameter is a string of characters consisting of the
	// following:
	//   - Any printable ASCII character ranging from the space character ( \u0020 )
	//   through the end of the ASCII character range
	//   - The printable characters in the Basic Latin and Latin-1 Supplement
	//   character set (through \u00FF )
	//   - The special characters tab ( \u0009 ), line feed ( \u000A ), and carriage
	//   return ( \u000D )
	//
	// This member is required.
	PolicyDocument *string

	// The friendly name of the policy. IAM user, group, role, and policy names must
	// be unique within the account. Names are not distinguished by case. For example,
	// you cannot create resources named both "MyResource" and "myresource".
	//
	// This member is required.
	PolicyName *string

	// A friendly description of the policy. Typically used to store information about
	// the permissions defined in the policy. For example, "Grants access to production
	// DynamoDB tables." The policy description is immutable. After a value is
	// assigned, it cannot be changed.
	Description *string

	// The path for the policy. For more information about paths, see IAM identifiers (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html)
	// in the IAM User Guide. This parameter is optional. If it is not included, it
	// defaults to a slash (/). This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex)
	// ) a string of characters consisting of either a forward slash (/) by itself or a
	// string that must begin and end with forward slashes. In addition, it can contain
	// any ASCII character from the ! ( \u0021 ) through the DEL character ( \u007F ),
	// including most punctuation characters, digits, and upper and lowercased letters.
	// You cannot use an asterisk (*) in the path name.
	Path *string

	// A list of tags that you want to attach to the new IAM customer managed policy.
	// Each tag consists of a key name and an associated value. For more information
	// about tagging, see Tagging IAM resources (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html)
	// in the IAM User Guide. If any one of the tags is invalid or if you exceed the
	// allowed maximum number of tags, then the entire request fails and the resource
	// is not created.
	Tags []types.Tag

	noSmithyDocumentSerde
}

// Contains the response to a successful CreatePolicy request.
type CreatePolicyOutput struct {

	// A structure containing details about the new policy.
	Policy *types.Policy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpCreatePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpCreatePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePolicy"); err != nil {
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
	if err = addOpCreatePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePolicy",
	}
}
