// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Add or change the operations role used by an environment. After this call is
// made, Elastic Beanstalk uses the associated operations role for permissions to
// downstream services during subsequent calls acting on this environment. For more
// information, see Operations roles (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/iam-operationsrole.html)
// in the AWS Elastic Beanstalk Developer Guide.
func (c *Client) AssociateEnvironmentOperationsRole(ctx context.Context, params *AssociateEnvironmentOperationsRoleInput, optFns ...func(*Options)) (*AssociateEnvironmentOperationsRoleOutput, error) {
	if params == nil {
		params = &AssociateEnvironmentOperationsRoleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateEnvironmentOperationsRole", params, optFns, c.addOperationAssociateEnvironmentOperationsRoleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateEnvironmentOperationsRoleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to add or change the operations role used by an environment.
type AssociateEnvironmentOperationsRoleInput struct {

	// The name of the environment to which to set the operations role.
	//
	// This member is required.
	EnvironmentName *string

	// The Amazon Resource Name (ARN) of an existing IAM role to be used as the
	// environment's operations role.
	//
	// This member is required.
	OperationsRole *string

	noSmithyDocumentSerde
}

type AssociateEnvironmentOperationsRoleOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateEnvironmentOperationsRoleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAssociateEnvironmentOperationsRole{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAssociateEnvironmentOperationsRole{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AssociateEnvironmentOperationsRole"); err != nil {
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
	if err = addOpAssociateEnvironmentOperationsRoleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateEnvironmentOperationsRole(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateEnvironmentOperationsRole(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AssociateEnvironmentOperationsRole",
	}
}
