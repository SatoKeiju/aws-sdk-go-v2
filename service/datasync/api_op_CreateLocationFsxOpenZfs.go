// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an endpoint for an Amazon FSx for OpenZFS file system that DataSync can
// access for a transfer. For more information, see Creating a location for FSx
// for OpenZFS (https://docs.aws.amazon.com/datasync/latest/userguide/create-openzfs-location.html)
// . Request parameters related to SMB aren't supported with the
// CreateLocationFsxOpenZfs operation.
func (c *Client) CreateLocationFsxOpenZfs(ctx context.Context, params *CreateLocationFsxOpenZfsInput, optFns ...func(*Options)) (*CreateLocationFsxOpenZfsOutput, error) {
	if params == nil {
		params = &CreateLocationFsxOpenZfsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLocationFsxOpenZfs", params, optFns, c.addOperationCreateLocationFsxOpenZfsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLocationFsxOpenZfsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLocationFsxOpenZfsInput struct {

	// The Amazon Resource Name (ARN) of the FSx for OpenZFS file system.
	//
	// This member is required.
	FsxFilesystemArn *string

	// The type of protocol that DataSync uses to access your file system.
	//
	// This member is required.
	Protocol *types.FsxProtocol

	// The ARNs of the security groups that are used to configure the FSx for OpenZFS
	// file system.
	//
	// This member is required.
	SecurityGroupArns []string

	// A subdirectory in the location's path that must begin with /fsx . DataSync uses
	// this subdirectory to read or write data (depending on whether the file system is
	// a source or destination location).
	Subdirectory *string

	// The key-value pair that represents a tag that you want to add to the resource.
	// The value can be an empty string. This value helps you manage, filter, and
	// search for your resources. We recommend that you create a name tag for your
	// location.
	Tags []types.TagListEntry

	noSmithyDocumentSerde
}

type CreateLocationFsxOpenZfsOutput struct {

	// The ARN of the FSx for OpenZFS file system location that you created.
	LocationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLocationFsxOpenZfsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateLocationFsxOpenZfs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateLocationFsxOpenZfs{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateLocationFsxOpenZfs"); err != nil {
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
	if err = addOpCreateLocationFsxOpenZfsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLocationFsxOpenZfs(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLocationFsxOpenZfs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateLocationFsxOpenZfs",
	}
}
