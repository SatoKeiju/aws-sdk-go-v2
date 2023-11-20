// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified constraint.
func (c *Client) UpdateConstraint(ctx context.Context, params *UpdateConstraintInput, optFns ...func(*Options)) (*UpdateConstraintOutput, error) {
	if params == nil {
		params = &UpdateConstraintInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateConstraint", params, optFns, c.addOperationUpdateConstraintMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateConstraintOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateConstraintInput struct {

	// The identifier of the constraint.
	//
	// This member is required.
	Id *string

	// The language code.
	//   - jp - Japanese
	//   - zh - Chinese
	AcceptLanguage *string

	// The updated description of the constraint.
	Description *string

	// The constraint parameters, in JSON format. The syntax depends on the constraint
	// type as follows: LAUNCH You are required to specify either the RoleArn or the
	// LocalRoleName but can't use both. Specify the RoleArn property as follows:
	// {"RoleArn" : "arn:aws:iam::123456789012:role/LaunchRole"} Specify the
	// LocalRoleName property as follows: {"LocalRoleName": "SCBasicLaunchRole"} If
	// you specify the LocalRoleName property, when an account uses the launch
	// constraint, the IAM role with that name in the account will be used. This allows
	// launch-role constraints to be account-agnostic so the administrator can create
	// fewer resources per shared account. The given role name must exist in the
	// account used to create the launch constraint and the account of the user who
	// launches a product with this launch constraint. You cannot have both a LAUNCH
	// and a STACKSET constraint. You also cannot have more than one LAUNCH constraint
	// on a product and portfolio. NOTIFICATION Specify the NotificationArns property
	// as follows: {"NotificationArns" : ["arn:aws:sns:us-east-1:123456789012:Topic"]}
	// RESOURCE_UPDATE Specify the TagUpdatesOnProvisionedProduct property as follows:
	// {"Version":"2.0","Properties":{"TagUpdateOnProvisionedProduct":"String"}} The
	// TagUpdatesOnProvisionedProduct property accepts a string value of ALLOWED or
	// NOT_ALLOWED . STACKSET Specify the Parameters property as follows: {"Version":
	// "String", "Properties": {"AccountList": [ "String" ], "RegionList": [ "String"
	// ], "AdminRole": "String", "ExecutionRole": "String"}} You cannot have both a
	// LAUNCH and a STACKSET constraint. You also cannot have more than one STACKSET
	// constraint on a product and portfolio. Products with a STACKSET constraint will
	// launch an CloudFormation stack set. TEMPLATE Specify the Rules property. For
	// more information, see Template Constraint Rules (http://docs.aws.amazon.com/servicecatalog/latest/adminguide/reference-template_constraint_rules.html)
	// .
	Parameters *string

	noSmithyDocumentSerde
}

type UpdateConstraintOutput struct {

	// Information about the constraint.
	ConstraintDetail *types.ConstraintDetail

	// The constraint parameters.
	ConstraintParameters *string

	// The status of the current request.
	Status types.Status

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateConstraintMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateConstraint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateConstraint{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "UpdateConstraint"); err != nil {
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
	if err = addOpUpdateConstraintValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateConstraint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateConstraint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "UpdateConstraint",
	}
}
