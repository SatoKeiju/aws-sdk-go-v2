// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a new custom vocabulary filter. You can use custom vocabulary filters
// to mask, delete, or flag specific words from your transcript. Custom vocabulary
// filters are commonly used to mask profanity in transcripts. Each language has a
// character set that contains all allowed characters for that specific language.
// If you use unsupported characters, your custom vocabulary filter request fails.
// Refer to Character Sets for Custom Vocabularies (https://docs.aws.amazon.com/transcribe/latest/dg/charsets.html)
// to get the character set for your language. For more information, see
// Vocabulary filtering (https://docs.aws.amazon.com/transcribe/latest/dg/vocabulary-filtering.html)
// .
func (c *Client) CreateVocabularyFilter(ctx context.Context, params *CreateVocabularyFilterInput, optFns ...func(*Options)) (*CreateVocabularyFilterOutput, error) {
	if params == nil {
		params = &CreateVocabularyFilterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVocabularyFilter", params, optFns, c.addOperationCreateVocabularyFilterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVocabularyFilterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVocabularyFilterInput struct {

	// The language code that represents the language of the entries in your
	// vocabulary filter. Each custom vocabulary filter must contain terms in only one
	// language. A custom vocabulary filter can only be used to transcribe files in the
	// same language as the filter. For example, if you create a custom vocabulary
	// filter using US English ( en-US ), you can only apply this filter to files that
	// contain English audio. For a list of supported languages and their associated
	// language codes, refer to the Supported languages (https://docs.aws.amazon.com/transcribe/latest/dg/supported-languages.html)
	// table.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// A unique name, chosen by you, for your new custom vocabulary filter. This name
	// is case sensitive, cannot contain spaces, and must be unique within an Amazon
	// Web Services account. If you try to create a new custom vocabulary filter with
	// the same name as an existing custom vocabulary filter, you get a
	// ConflictException error.
	//
	// This member is required.
	VocabularyFilterName *string

	// The Amazon Resource Name (ARN) of an IAM role that has permissions to access
	// the Amazon S3 bucket that contains your input files (in this case, your custom
	// vocabulary filter). If the role that you specify doesn’t have the appropriate
	// permissions to access the specified Amazon S3 location, your request fails. IAM
	// role ARNs have the format arn:partition:iam::account:role/role-name-with-path .
	// For example: arn:aws:iam::111122223333:role/Admin . For more information, see
	// IAM ARNs (https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-arns)
	// .
	DataAccessRoleArn *string

	// Adds one or more custom tags, each in the form of a key:value pair, to a new
	// custom vocabulary filter at the time you create this new vocabulary filter. To
	// learn more about using tags with Amazon Transcribe, refer to Tagging resources (https://docs.aws.amazon.com/transcribe/latest/dg/tagging.html)
	// .
	Tags []types.Tag

	// The Amazon S3 location of the text file that contains your custom vocabulary
	// filter terms. The URI must be located in the same Amazon Web Services Region as
	// the resource you're calling. Here's an example URI path:
	// s3://DOC-EXAMPLE-BUCKET/my-vocab-filter-file.txt Note that if you include
	// VocabularyFilterFileUri in your request, you cannot use Words ; you must choose
	// one or the other.
	VocabularyFilterFileUri *string

	// Use this parameter if you want to create your custom vocabulary filter by
	// including all desired terms, as comma-separated values, within your request. The
	// other option for creating your vocabulary filter is to save your entries in a
	// text file and upload them to an Amazon S3 bucket, then specify the location of
	// your file using the VocabularyFilterFileUri parameter. Note that if you include
	// Words in your request, you cannot use VocabularyFilterFileUri ; you must choose
	// one or the other. Each language has a character set that contains all allowed
	// characters for that specific language. If you use unsupported characters, your
	// custom vocabulary filter request fails. Refer to Character Sets for Custom
	// Vocabularies (https://docs.aws.amazon.com/transcribe/latest/dg/charsets.html) to
	// get the character set for your language.
	Words []string

	noSmithyDocumentSerde
}

type CreateVocabularyFilterOutput struct {

	// The language code you selected for your custom vocabulary filter.
	LanguageCode types.LanguageCode

	// The date and time you created your custom vocabulary filter. Timestamps are in
	// the format YYYY-MM-DD'T'HH:MM:SS.SSSSSS-UTC . For example,
	// 2022-05-04T12:32:58.761000-07:00 represents 12:32 PM UTC-7 on May 4, 2022.
	LastModifiedTime *time.Time

	// The name you chose for your custom vocabulary filter.
	VocabularyFilterName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateVocabularyFilterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateVocabularyFilter{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateVocabularyFilter{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateVocabularyFilter"); err != nil {
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
	if err = addOpCreateVocabularyFilterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVocabularyFilter(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateVocabularyFilter(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateVocabularyFilter",
	}
}
