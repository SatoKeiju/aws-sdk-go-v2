// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an unreferenced commit that represents the result of merging two
// branches using a specified merge strategy. This can help you determine the
// outcome of a potential merge. This API cannot be used with the fast-forward
// merge strategy because that strategy does not create a merge commit. This
// unreferenced merge commit can only be accessed using the GetCommit API or
// through git commands such as git fetch. To retrieve this commit, you must
// specify its commit ID or otherwise reference it.
func (c *Client) CreateUnreferencedMergeCommit(ctx context.Context, params *CreateUnreferencedMergeCommitInput, optFns ...func(*Options)) (*CreateUnreferencedMergeCommitOutput, error) {
	if params == nil {
		params = &CreateUnreferencedMergeCommitInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateUnreferencedMergeCommit", params, optFns, c.addOperationCreateUnreferencedMergeCommitMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateUnreferencedMergeCommitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateUnreferencedMergeCommitInput struct {

	// The branch, tag, HEAD, or other fully qualified reference used to identify a
	// commit (for example, a branch name or a full commit ID).
	//
	// This member is required.
	DestinationCommitSpecifier *string

	// The merge option or strategy you want to use to merge the code.
	//
	// This member is required.
	MergeOption types.MergeOptionTypeEnum

	// The name of the repository where you want to create the unreferenced merge
	// commit.
	//
	// This member is required.
	RepositoryName *string

	// The branch, tag, HEAD, or other fully qualified reference used to identify a
	// commit (for example, a branch name or a full commit ID).
	//
	// This member is required.
	SourceCommitSpecifier *string

	// The name of the author who created the unreferenced commit. This information is
	// used as both the author and committer for the commit.
	AuthorName *string

	// The commit message for the unreferenced commit.
	CommitMessage *string

	// The level of conflict detail to use. If unspecified, the default FILE_LEVEL is
	// used, which returns a not-mergeable result if the same file has differences in
	// both branches. If LINE_LEVEL is specified, a conflict is considered not
	// mergeable if the same file in both branches has differences on the same line.
	ConflictDetailLevel types.ConflictDetailLevelTypeEnum

	// If AUTOMERGE is the conflict resolution strategy, a list of inputs to use when
	// resolving conflicts during a merge.
	ConflictResolution *types.ConflictResolution

	// Specifies which branch to use when resolving conflicts, or whether to attempt
	// automatically merging two versions of a file. The default is NONE, which
	// requires any conflicts to be resolved manually before the merge operation is
	// successful.
	ConflictResolutionStrategy types.ConflictResolutionStrategyTypeEnum

	// The email address for the person who created the unreferenced commit.
	Email *string

	// If the commit contains deletions, whether to keep a folder or folder structure
	// if the changes leave the folders empty. If this is specified as true, a .gitkeep
	// file is created for empty folders. The default is false.
	KeepEmptyFolders bool

	noSmithyDocumentSerde
}

type CreateUnreferencedMergeCommitOutput struct {

	// The full commit ID of the commit that contains your merge results.
	CommitId *string

	// The full SHA-1 pointer of the tree information for the commit that contains the
	// merge results.
	TreeId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateUnreferencedMergeCommitMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateUnreferencedMergeCommit{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateUnreferencedMergeCommit{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateUnreferencedMergeCommit"); err != nil {
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
	if err = addOpCreateUnreferencedMergeCommitValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUnreferencedMergeCommit(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateUnreferencedMergeCommit(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateUnreferencedMergeCommit",
	}
}
