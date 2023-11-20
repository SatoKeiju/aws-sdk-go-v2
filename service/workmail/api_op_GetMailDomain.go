// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets details for a mail domain, including domain records required to configure
// your domain with recommended security.
func (c *Client) GetMailDomain(ctx context.Context, params *GetMailDomainInput, optFns ...func(*Options)) (*GetMailDomainOutput, error) {
	if params == nil {
		params = &GetMailDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMailDomain", params, optFns, c.addOperationGetMailDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMailDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMailDomainInput struct {

	// The domain from which you want to retrieve details.
	//
	// This member is required.
	DomainName *string

	// The WorkMail organization for which the domain is retrieved.
	//
	// This member is required.
	OrganizationId *string

	noSmithyDocumentSerde
}

type GetMailDomainOutput struct {

	// Indicates the status of a DKIM verification.
	DkimVerificationStatus types.DnsRecordVerificationStatus

	// Specifies whether the domain is the default domain for your organization.
	IsDefault bool

	// Specifies whether the domain is a test domain provided by WorkMail, or a custom
	// domain.
	IsTestDomain bool

	// Indicates the status of the domain ownership verification.
	OwnershipVerificationStatus types.DnsRecordVerificationStatus

	// A list of the DNS records that WorkMail recommends adding in your DNS provider
	// for the best user experience. The records configure your domain with DMARC, SPF,
	// DKIM, and direct incoming email traffic to SES. See admin guide for more
	// details.
	Records []types.DnsRecord

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMailDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetMailDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMailDomain{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "GetMailDomain"); err != nil {
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
	if err = addOpGetMailDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMailDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetMailDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "GetMailDomain",
	}
}
