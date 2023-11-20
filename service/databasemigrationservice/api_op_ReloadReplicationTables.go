// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Reloads the target database table with the source data for a given DMS
// Serverless replication configuration. You can only use this operation with a
// task in the RUNNING state, otherwise the service will throw an
// InvalidResourceStateFault exception.
func (c *Client) ReloadReplicationTables(ctx context.Context, params *ReloadReplicationTablesInput, optFns ...func(*Options)) (*ReloadReplicationTablesOutput, error) {
	if params == nil {
		params = &ReloadReplicationTablesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReloadReplicationTables", params, optFns, c.addOperationReloadReplicationTablesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ReloadReplicationTablesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReloadReplicationTablesInput struct {

	// The Amazon Resource Name of the replication config for which to reload tables.
	//
	// This member is required.
	ReplicationConfigArn *string

	// The list of tables to reload.
	//
	// This member is required.
	TablesToReload []types.TableToReload

	// Options for reload. Specify data-reload to reload the data and re-validate it
	// if validation is enabled. Specify validate-only to re-validate the table. This
	// option applies only when validation is enabled for the replication.
	ReloadOption types.ReloadOptionValue

	noSmithyDocumentSerde
}

type ReloadReplicationTablesOutput struct {

	// The Amazon Resource Name of the replication config for which to reload tables.
	ReplicationConfigArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationReloadReplicationTablesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpReloadReplicationTables{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpReloadReplicationTables{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ReloadReplicationTables"); err != nil {
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
	if err = addOpReloadReplicationTablesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opReloadReplicationTables(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opReloadReplicationTables(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ReloadReplicationTables",
	}
}
