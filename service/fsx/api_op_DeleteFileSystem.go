// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a file system. After deletion, the file system no longer exists, and
// its data is gone. Any existing automatic backups and snapshots are also deleted.
// To delete an Amazon FSx for NetApp ONTAP file system, first delete all the
// volumes and storage virtual machines (SVMs) on the file system. Then provide a
// FileSystemId value to the DeleFileSystem operation. By default, when you delete
// an Amazon FSx for Windows File Server file system, a final backup is created
// upon deletion. This final backup isn't subject to the file system's retention
// policy, and must be manually deleted. To delete an Amazon FSx for Lustre file
// system, first unmount (https://docs.aws.amazon.com/fsx/latest/LustreGuide/unmounting-fs.html)
// it from every connected Amazon EC2 instance, then provide a FileSystemId value
// to the DeleFileSystem operation. By default, Amazon FSx will not take a final
// backup when the DeleteFileSystem operation is invoked. On file systems not
// linked to an Amazon S3 bucket, set SkipFinalBackup to false to take a final
// backup of the file system you are deleting. Backups cannot be enabled on
// S3-linked file systems. To ensure all of your data is written back to S3 before
// deleting your file system, you can either monitor for the
// AgeOfOldestQueuedMessage (https://docs.aws.amazon.com/fsx/latest/LustreGuide/monitoring-cloudwatch.html#auto-import-export-metrics)
// metric to be zero (if using automatic export) or you can run an export data
// repository task (https://docs.aws.amazon.com/fsx/latest/LustreGuide/export-data-repo-task-dra.html)
// . If you have automatic export enabled and want to use an export data repository
// task, you have to disable automatic export before executing the export data
// repository task. The DeleteFileSystem operation returns while the file system
// has the DELETING status. You can check the file system deletion status by
// calling the DescribeFileSystems (https://docs.aws.amazon.com/fsx/latest/APIReference/API_DescribeFileSystems.html)
// operation, which returns a list of file systems in your account. If you pass the
// file system ID for a deleted file system, the DescribeFileSystems operation
// returns a FileSystemNotFound error. If a data repository task is in a PENDING
// or EXECUTING state, deleting an Amazon FSx for Lustre file system will fail
// with an HTTP status code 400 (Bad Request). The data in a deleted file system is
// also deleted and can't be recovered by any means.
func (c *Client) DeleteFileSystem(ctx context.Context, params *DeleteFileSystemInput, optFns ...func(*Options)) (*DeleteFileSystemOutput, error) {
	if params == nil {
		params = &DeleteFileSystemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteFileSystem", params, optFns, c.addOperationDeleteFileSystemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteFileSystemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request object for DeleteFileSystem operation.
type DeleteFileSystemInput struct {

	// The ID of the file system that you want to delete.
	//
	// This member is required.
	FileSystemId *string

	// A string of up to 63 ASCII characters that Amazon FSx uses to ensure idempotent
	// deletion. This token is automatically filled on your behalf when using the
	// Command Line Interface (CLI) or an Amazon Web Services SDK.
	ClientRequestToken *string

	// The configuration object for the Amazon FSx for Lustre file system being
	// deleted in the DeleteFileSystem operation.
	LustreConfiguration *types.DeleteFileSystemLustreConfiguration

	// The configuration object for the OpenZFS file system used in the
	// DeleteFileSystem operation.
	OpenZFSConfiguration *types.DeleteFileSystemOpenZFSConfiguration

	// The configuration object for the Microsoft Windows file system used in the
	// DeleteFileSystem operation.
	WindowsConfiguration *types.DeleteFileSystemWindowsConfiguration

	noSmithyDocumentSerde
}

// The response object for the DeleteFileSystem operation.
type DeleteFileSystemOutput struct {

	// The ID of the file system that's being deleted.
	FileSystemId *string

	// The file system lifecycle for the deletion request. If the DeleteFileSystem
	// operation is successful, this status is DELETING .
	Lifecycle types.FileSystemLifecycle

	// The response object for the Amazon FSx for Lustre file system being deleted in
	// the DeleteFileSystem operation.
	LustreResponse *types.DeleteFileSystemLustreResponse

	// The response object for the OpenZFS file system that's being deleted in the
	// DeleteFileSystem operation.
	OpenZFSResponse *types.DeleteFileSystemOpenZFSResponse

	// The response object for the Microsoft Windows file system used in the
	// DeleteFileSystem operation.
	WindowsResponse *types.DeleteFileSystemWindowsResponse

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteFileSystemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteFileSystem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteFileSystem{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteFileSystem"); err != nil {
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
	if err = addIdempotencyToken_opDeleteFileSystemMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDeleteFileSystemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFileSystem(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpDeleteFileSystem struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDeleteFileSystem) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDeleteFileSystem) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*DeleteFileSystemInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *DeleteFileSystemInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opDeleteFileSystemMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpDeleteFileSystem{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opDeleteFileSystem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteFileSystem",
	}
}
