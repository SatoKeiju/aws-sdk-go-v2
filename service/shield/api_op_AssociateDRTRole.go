// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Authorizes the Shield Response Team (SRT) using the specified role, to access
// your Amazon Web Services account to assist with DDoS attack mitigation during
// potential attacks. This enables the SRT to inspect your WAF configuration and
// create or update WAF rules and web ACLs. You can associate only one RoleArn
// with your subscription. If you submit an AssociateDRTRole request for an
// account that already has an associated role, the new RoleArn will replace the
// existing RoleArn . Prior to making the AssociateDRTRole request, you must
// attach the AWSShieldDRTAccessPolicy managed policy to the role that you'll
// specify in the request. You can access this policy in the IAM console at
// AWSShieldDRTAccessPolicy (https://console.aws.amazon.com/iam/home?#/policies/arn:aws:iam::aws:policy/service-role/AWSShieldDRTAccessPolicy)
// . For more information see Adding and removing IAM identity permissions (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_manage-attach-detach.html)
// . The role must also trust the service principal drt.shield.amazonaws.com . For
// more information, see IAM JSON policy elements: Principal (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html)
// . The SRT will have access only to your WAF and Shield resources. By submitting
// this request, you authorize the SRT to inspect your WAF and Shield configuration
// and create and update WAF rules and web ACLs on your behalf. The SRT takes these
// actions only if explicitly authorized by you. You must have the iam:PassRole
// permission to make an AssociateDRTRole request. For more information, see
// Granting a user permissions to pass a role to an Amazon Web Services service (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_passrole.html)
// . To use the services of the SRT and make an AssociateDRTRole request, you must
// be subscribed to the Business Support plan (http://aws.amazon.com/premiumsupport/business-support/)
// or the Enterprise Support plan (http://aws.amazon.com/premiumsupport/enterprise-support/)
// .
func (c *Client) AssociateDRTRole(ctx context.Context, params *AssociateDRTRoleInput, optFns ...func(*Options)) (*AssociateDRTRoleOutput, error) {
	if params == nil {
		params = &AssociateDRTRoleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateDRTRole", params, optFns, c.addOperationAssociateDRTRoleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateDRTRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateDRTRoleInput struct {

	// The Amazon Resource Name (ARN) of the role the SRT will use to access your
	// Amazon Web Services account. Prior to making the AssociateDRTRole request, you
	// must attach the AWSShieldDRTAccessPolicy (https://console.aws.amazon.com/iam/home?#/policies/arn:aws:iam::aws:policy/service-role/AWSShieldDRTAccessPolicy)
	// managed policy to this role. For more information see Attaching and Detaching
	// IAM Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_manage-attach-detach.html)
	// .
	//
	// This member is required.
	RoleArn *string

	noSmithyDocumentSerde
}

type AssociateDRTRoleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateDRTRoleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateDRTRole{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateDRTRole{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociateDRTRole"); err != nil {
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
	if err = addOpAssociateDRTRoleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateDRTRole(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateDRTRole(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociateDRTRole",
	}
}
