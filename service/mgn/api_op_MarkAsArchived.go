// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Archives specific Source Servers by setting the SourceServer.isArchived
// property to true for specified SourceServers by ID. This command only works for
// SourceServers with a lifecycle. state which equals DISCONNECTED or CUTOVER.
func (c *Client) MarkAsArchived(ctx context.Context, params *MarkAsArchivedInput, optFns ...func(*Options)) (*MarkAsArchivedOutput, error) {
	if params == nil {
		params = &MarkAsArchivedInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "MarkAsArchived", params, optFns, c.addOperationMarkAsArchivedMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*MarkAsArchivedOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type MarkAsArchivedInput struct {

	// Mark as archived by Source Server ID.
	//
	// This member is required.
	SourceServerID *string

	// Mark as archived by Account ID.
	AccountID *string

	noSmithyDocumentSerde
}

type MarkAsArchivedOutput struct {

	// Source server application ID.
	ApplicationID *string

	// Source server ARN.
	Arn *string

	// Source Server connector action.
	ConnectorAction *types.SourceServerConnectorAction

	// Source server data replication info.
	DataReplicationInfo *types.DataReplicationInfo

	// Source server fqdn for action framework.
	FqdnForActionFramework *string

	// Source server archived status.
	IsArchived *bool

	// Source server launched instance.
	LaunchedInstance *types.LaunchedInstance

	// Source server lifecycle state.
	LifeCycle *types.LifeCycle

	// Source server replication type.
	ReplicationType types.ReplicationType

	// Source server properties.
	SourceProperties *types.SourceProperties

	// Source server ID.
	SourceServerID *string

	// Source server Tags.
	Tags map[string]string

	// Source server user provided ID.
	UserProvidedID *string

	// Source server vCenter client id.
	VcenterClientID *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationMarkAsArchivedMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpMarkAsArchived{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpMarkAsArchived{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "MarkAsArchived"); err != nil {
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
	if err = addOpMarkAsArchivedValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opMarkAsArchived(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opMarkAsArchived(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "MarkAsArchived",
	}
}
