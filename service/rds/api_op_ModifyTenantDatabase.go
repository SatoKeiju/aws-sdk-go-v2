// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies an existing tenant database in a DB instance. You can change the
// tenant database name or the master user password. This operation is supported
// only for RDS for Oracle CDB instances using the multi-tenant configuration.
func (c *Client) ModifyTenantDatabase(ctx context.Context, params *ModifyTenantDatabaseInput, optFns ...func(*Options)) (*ModifyTenantDatabaseOutput, error) {
	if params == nil {
		params = &ModifyTenantDatabaseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyTenantDatabase", params, optFns, c.addOperationModifyTenantDatabaseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyTenantDatabaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyTenantDatabaseInput struct {

	// The identifier of the DB instance that contains the tenant database that you
	// are modifying. This parameter isn't case-sensitive. Constraints:
	//   - Must match the identifier of an existing DB instance.
	//
	// This member is required.
	DBInstanceIdentifier *string

	// The user-supplied name of the tenant database that you want to modify. This
	// parameter isn’t case-sensitive. Constraints:
	//   - Must match the identifier of an existing tenant database.
	//
	// This member is required.
	TenantDBName *string

	// The new password for the master user of the specified tenant database in your
	// DB instance. Amazon RDS operations never return the password, so this action
	// provides a way to regain access to a tenant database user if the password is
	// lost. This includes restoring privileges that might have been accidentally
	// revoked. Constraints:
	//   - Can include any printable ASCII character except / , " (double quote), @ , &
	//   (ampersand), and ' (single quote).
	// Length constraints:
	//   - Must contain between 8 and 30 characters.
	MasterUserPassword *string

	// The new name of the tenant database when renaming a tenant database. This
	// parameter isn’t case-sensitive. Constraints:
	//   - Can't be the string null or any other reserved word.
	//   - Can't be longer than 8 characters.
	NewTenantDBName *string

	noSmithyDocumentSerde
}

type ModifyTenantDatabaseOutput struct {

	// A tenant database in the DB instance. This data type is an element in the
	// response to the DescribeTenantDatabases action.
	TenantDatabase *types.TenantDatabase

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyTenantDatabaseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyTenantDatabase{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyTenantDatabase{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyTenantDatabase"); err != nil {
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
	if err = addOpModifyTenantDatabaseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyTenantDatabase(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyTenantDatabase(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyTenantDatabase",
	}
}
