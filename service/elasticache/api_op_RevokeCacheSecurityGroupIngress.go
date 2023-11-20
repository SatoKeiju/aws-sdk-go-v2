// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Revokes ingress from a cache security group. Use this operation to disallow
// access from an Amazon EC2 security group that had been previously authorized.
func (c *Client) RevokeCacheSecurityGroupIngress(ctx context.Context, params *RevokeCacheSecurityGroupIngressInput, optFns ...func(*Options)) (*RevokeCacheSecurityGroupIngressOutput, error) {
	if params == nil {
		params = &RevokeCacheSecurityGroupIngressInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RevokeCacheSecurityGroupIngress", params, optFns, c.addOperationRevokeCacheSecurityGroupIngressMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RevokeCacheSecurityGroupIngressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a RevokeCacheSecurityGroupIngress operation.
type RevokeCacheSecurityGroupIngressInput struct {

	// The name of the cache security group to revoke ingress from.
	//
	// This member is required.
	CacheSecurityGroupName *string

	// The name of the Amazon EC2 security group to revoke access from.
	//
	// This member is required.
	EC2SecurityGroupName *string

	// The Amazon account number of the Amazon EC2 security group owner. Note that
	// this is not the same thing as an Amazon access key ID - you must provide a valid
	// Amazon account number for this parameter.
	//
	// This member is required.
	EC2SecurityGroupOwnerId *string

	noSmithyDocumentSerde
}

type RevokeCacheSecurityGroupIngressOutput struct {

	// Represents the output of one of the following operations:
	//   - AuthorizeCacheSecurityGroupIngress
	//   - CreateCacheSecurityGroup
	//   - RevokeCacheSecurityGroupIngress
	CacheSecurityGroup *types.CacheSecurityGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRevokeCacheSecurityGroupIngressMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRevokeCacheSecurityGroupIngress{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRevokeCacheSecurityGroupIngress{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RevokeCacheSecurityGroupIngress"); err != nil {
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
	if err = addOpRevokeCacheSecurityGroupIngressValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRevokeCacheSecurityGroupIngress(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRevokeCacheSecurityGroupIngress(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RevokeCacheSecurityGroupIngress",
	}
}
