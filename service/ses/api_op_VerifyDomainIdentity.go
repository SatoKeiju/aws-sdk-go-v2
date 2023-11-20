// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a domain to the list of identities for your Amazon SES account in the
// current Amazon Web Services Region and attempts to verify it. For more
// information about verifying domains, see Verifying Email Addresses and Domains (https://docs.aws.amazon.com/ses/latest/dg/verify-addresses-and-domains.html)
// in the Amazon SES Developer Guide. You can execute this operation no more than
// once per second.
func (c *Client) VerifyDomainIdentity(ctx context.Context, params *VerifyDomainIdentityInput, optFns ...func(*Options)) (*VerifyDomainIdentityOutput, error) {
	if params == nil {
		params = &VerifyDomainIdentityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "VerifyDomainIdentity", params, optFns, c.addOperationVerifyDomainIdentityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*VerifyDomainIdentityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to begin Amazon SES domain verification and to generate
// the TXT records that you must publish to the DNS server of your domain to
// complete the verification. For information about domain verification, see the
// Amazon SES Developer Guide (https://docs.aws.amazon.com/ses/latest/dg/creating-identities.html#verify-domain-procedure)
// .
type VerifyDomainIdentityInput struct {

	// The domain to be verified.
	//
	// This member is required.
	Domain *string

	noSmithyDocumentSerde
}

// Returns a TXT record that you must publish to the DNS server of your domain to
// complete domain verification with Amazon SES.
type VerifyDomainIdentityOutput struct {

	// A TXT record that you must place in the DNS settings of the domain to complete
	// domain verification with Amazon SES. As Amazon SES searches for the TXT record,
	// the domain's verification status is "Pending". When Amazon SES detects the
	// record, the domain's verification status changes to "Success". If Amazon SES is
	// unable to detect the record within 72 hours, the domain's verification status
	// changes to "Failed." In that case, to verify the domain, you must restart the
	// verification process from the beginning. The domain's verification status also
	// changes to "Success" when it is DKIM verified.
	//
	// This member is required.
	VerificationToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationVerifyDomainIdentityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpVerifyDomainIdentity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpVerifyDomainIdentity{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "VerifyDomainIdentity"); err != nil {
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
	if err = addOpVerifyDomainIdentityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opVerifyDomainIdentity(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opVerifyDomainIdentity(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "VerifyDomainIdentity",
	}
}
