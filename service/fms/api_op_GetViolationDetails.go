// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves violations for a resource based on the specified Firewall Manager
// policy and Amazon Web Services account.
func (c *Client) GetViolationDetails(ctx context.Context, params *GetViolationDetailsInput, optFns ...func(*Options)) (*GetViolationDetailsOutput, error) {
	if params == nil {
		params = &GetViolationDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetViolationDetails", params, optFns, c.addOperationGetViolationDetailsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetViolationDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetViolationDetailsInput struct {

	// The Amazon Web Services account ID that you want the details for.
	//
	// This member is required.
	MemberAccount *string

	// The ID of the Firewall Manager policy that you want the details for. You can
	// get violation details for the following policy types:
	//   - DNS Firewall
	//   - Imported Network Firewall
	//   - Network Firewall
	//   - Security group content audit
	//   - Third-party firewall
	//
	// This member is required.
	PolicyId *string

	// The ID of the resource that has violations.
	//
	// This member is required.
	ResourceId *string

	// The resource type. This is in the format shown in the Amazon Web Services
	// Resource Types Reference (https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html)
	// . Supported resource types are: AWS::EC2::Instance , AWS::EC2::NetworkInterface
	// , AWS::EC2::SecurityGroup , AWS::NetworkFirewall::FirewallPolicy , and
	// AWS::EC2::Subnet .
	//
	// This member is required.
	ResourceType *string

	noSmithyDocumentSerde
}

type GetViolationDetailsOutput struct {

	// Violation detail for a resource.
	ViolationDetail *types.ViolationDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetViolationDetailsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetViolationDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetViolationDetails{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetViolationDetails"); err != nil {
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
	if err = addOpGetViolationDetailsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetViolationDetails(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetViolationDetails(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetViolationDetails",
	}
}
