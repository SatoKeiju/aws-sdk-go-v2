// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsmv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies attributes for AWS CloudHSM backup.
func (c *Client) ModifyBackupAttributes(ctx context.Context, params *ModifyBackupAttributesInput, optFns ...func(*Options)) (*ModifyBackupAttributesOutput, error) {
	if params == nil {
		params = &ModifyBackupAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyBackupAttributes", params, optFns, c.addOperationModifyBackupAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyBackupAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyBackupAttributesInput struct {

	// The identifier (ID) of the backup to modify. To find the ID of a backup, use
	// the DescribeBackups operation.
	//
	// This member is required.
	BackupId *string

	// Specifies whether the service should exempt a backup from the retention policy
	// for the cluster. True exempts a backup from the retention policy. False means
	// the service applies the backup retention policy defined at the cluster.
	//
	// This member is required.
	NeverExpires *bool

	noSmithyDocumentSerde
}

type ModifyBackupAttributesOutput struct {

	// Contains information about a backup of an AWS CloudHSM cluster. All backup
	// objects contain the BackupId , BackupState , ClusterId , and CreateTimestamp
	// parameters. Backups that were copied into a destination region additionally
	// contain the CopyTimestamp , SourceBackup , SourceCluster , and SourceRegion
	// parameters. A backup that is pending deletion will include the DeleteTimestamp
	// parameter.
	Backup *types.Backup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyBackupAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpModifyBackupAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyBackupAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyBackupAttributes"); err != nil {
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
	if err = addOpModifyBackupAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyBackupAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyBackupAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyBackupAttributes",
	}
}
