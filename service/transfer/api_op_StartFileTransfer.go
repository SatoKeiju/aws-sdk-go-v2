// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Begins a file transfer between local Amazon Web Services storage and a remote
// AS2 or SFTP server.
//   - For an AS2 connector, you specify the ConnectorId and one or more
//     SendFilePaths to identify the files you want to transfer.
//   - For an SFTP connector, the file transfer can be either outbound or inbound.
//     In both cases, you specify the ConnectorId . Depending on the direction of the
//     transfer, you also specify the following items:
//   - If you are transferring file from a partner's SFTP server to Amazon Web
//     Services storage, you specify one or more RetreiveFilePaths to identify the
//     files you want to transfer, and a LocalDirectoryPath to specify the
//     destination folder.
//   - If you are transferring file to a partner's SFTP server from Amazon Web
//     Services storage, you specify one or more SendFilePaths to identify the files
//     you want to transfer, and a RemoteDirectoryPath to specify the destination
//     folder.
func (c *Client) StartFileTransfer(ctx context.Context, params *StartFileTransferInput, optFns ...func(*Options)) (*StartFileTransferOutput, error) {
	if params == nil {
		params = &StartFileTransferInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartFileTransfer", params, optFns, c.addOperationStartFileTransferMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartFileTransferOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartFileTransferInput struct {

	// The unique identifier for the connector.
	//
	// This member is required.
	ConnectorId *string

	// For an inbound transfer, the LocaDirectoryPath specifies the destination for
	// one or more files that are transferred from the partner's SFTP server.
	LocalDirectoryPath *string

	// For an outbound transfer, the RemoteDirectoryPath specifies the destination for
	// one or more files that are transferred to the partner's SFTP server. If you
	// don't specify a RemoteDirectoryPath , the destination for transferred files is
	// the SFTP user's home directory.
	RemoteDirectoryPath *string

	// One or more source paths for the partner's SFTP server. Each string represents
	// a source file path for one inbound file transfer.
	RetrieveFilePaths []string

	// One or more source paths for the Amazon S3 storage. Each string represents a
	// source file path for one outbound file transfer. For example,
	// DOC-EXAMPLE-BUCKET/myfile.txt . Replace  DOC-EXAMPLE-BUCKET  with one of your
	// actual buckets.
	SendFilePaths []string

	noSmithyDocumentSerde
}

type StartFileTransferOutput struct {

	// Returns the unique identifier for the file transfer.
	//
	// This member is required.
	TransferId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartFileTransferMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartFileTransfer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartFileTransfer{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "StartFileTransfer"); err != nil {
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
	if err = addOpStartFileTransferValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartFileTransfer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartFileTransfer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "StartFileTransfer",
	}
}
