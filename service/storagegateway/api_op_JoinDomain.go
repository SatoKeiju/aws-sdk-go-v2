// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds a file gateway to an Active Directory domain. This operation is only
// supported for file gateways that support the SMB file protocol. Joining a domain
// creates an Active Directory computer account in the default organizational unit,
// using the gateway's Gateway ID as the account name (for example, SGW-1234ADE).
// If your Active Directory environment requires that you pre-stage accounts to
// facilitate the join domain process, you will need to create this account ahead
// of time. To create the gateway's computer account in an organizational unit
// other than the default, you must specify the organizational unit when joining
// the domain.
func (c *Client) JoinDomain(ctx context.Context, params *JoinDomainInput, optFns ...func(*Options)) (*JoinDomainOutput, error) {
	if params == nil {
		params = &JoinDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "JoinDomain", params, optFns, c.addOperationJoinDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*JoinDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// JoinDomainInput
type JoinDomainInput struct {

	// The name of the domain that you want the gateway to join.
	//
	// This member is required.
	DomainName *string

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and Amazon Web Services Region.
	//
	// This member is required.
	GatewayARN *string

	// Sets the password of the user who has permission to add the gateway to the
	// Active Directory domain.
	//
	// This member is required.
	Password *string

	// Sets the user name of user who has permission to add the gateway to the Active
	// Directory domain. The domain user account should be enabled to join computers to
	// the domain. For example, you can use the domain administrator account or an
	// account with delegated permissions to join computers to the domain.
	//
	// This member is required.
	UserName *string

	// List of IPv4 addresses, NetBIOS names, or host names of your domain server. If
	// you need to specify the port number include it after the colon (“:”). For
	// example, mydc.mydomain.com:389 .
	DomainControllers []string

	// The organizational unit (OU) is a container in an Active Directory that can
	// hold users, groups, computers, and other OUs and this parameter specifies the OU
	// that the gateway will join within the AD domain.
	OrganizationalUnit *string

	// Specifies the time in seconds, in which the JoinDomain operation must complete.
	// The default is 20 seconds.
	TimeoutInSeconds *int32

	noSmithyDocumentSerde
}

// JoinDomainOutput
type JoinDomainOutput struct {

	// Indicates the status of the gateway as a member of the Active Directory domain.
	//   - ACCESS_DENIED : Indicates that the JoinDomain operation failed due to an
	//   authentication error.
	//   - DETACHED : Indicates that gateway is not joined to a domain.
	//   - JOINED : Indicates that the gateway has successfully joined a domain.
	//   - JOINING : Indicates that a JoinDomain operation is in progress.
	//   - NETWORK_ERROR : Indicates that JoinDomain operation failed due to a network
	//   or connectivity error.
	//   - TIMEOUT : Indicates that the JoinDomain operation failed because the
	//   operation didn't complete within the allotted time.
	//   - UNKNOWN_ERROR : Indicates that the JoinDomain operation failed due to
	//   another type of error.
	ActiveDirectoryStatus types.ActiveDirectoryStatus

	// The unique Amazon Resource Name (ARN) of the gateway that joined the domain.
	GatewayARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationJoinDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpJoinDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpJoinDomain{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "JoinDomain"); err != nil {
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
	if err = addOpJoinDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opJoinDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opJoinDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "JoinDomain",
	}
}
