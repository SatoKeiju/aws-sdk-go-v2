// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Specifies an Amazon Web Services rule that you want to share with another
// account, the account that you want to share the rule with, and the operations
// that you want the account to be able to perform on the rule.
func (c *Client) PutResolverRulePolicy(ctx context.Context, params *PutResolverRulePolicyInput, optFns ...func(*Options)) (*PutResolverRulePolicyOutput, error) {
	if params == nil {
		params = &PutResolverRulePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutResolverRulePolicy", params, optFns, c.addOperationPutResolverRulePolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutResolverRulePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutResolverRulePolicyInput struct {

	// The Amazon Resource Name (ARN) of the rule that you want to share with another
	// account.
	//
	// This member is required.
	Arn *string

	// An Identity and Access Management policy statement that lists the rules that
	// you want to share with another Amazon Web Services account and the operations
	// that you want the account to be able to perform. You can specify the following
	// operations in the Action section of the statement:
	//   - route53resolver:GetResolverRule
	//   - route53resolver:AssociateResolverRule
	//   - route53resolver:DisassociateResolverRule
	//   - route53resolver:ListResolverRules
	//   - route53resolver:ListResolverRuleAssociations
	// In the Resource section of the statement, specify the ARN for the rule that you
	// want to share with another account. Specify the same ARN that you specified in
	// Arn .
	//
	// This member is required.
	ResolverRulePolicy *string

	noSmithyDocumentSerde
}

// The response to a PutResolverRulePolicy request.
type PutResolverRulePolicyOutput struct {

	// Whether the PutResolverRulePolicy request was successful.
	ReturnValue bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutResolverRulePolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutResolverRulePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutResolverRulePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutResolverRulePolicy"); err != nil {
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
	if err = addOpPutResolverRulePolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutResolverRulePolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutResolverRulePolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutResolverRulePolicy",
	}
}
