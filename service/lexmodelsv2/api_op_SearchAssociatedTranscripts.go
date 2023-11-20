// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Search for associated transcripts that meet the specified criteria.
func (c *Client) SearchAssociatedTranscripts(ctx context.Context, params *SearchAssociatedTranscriptsInput, optFns ...func(*Options)) (*SearchAssociatedTranscriptsOutput, error) {
	if params == nil {
		params = &SearchAssociatedTranscriptsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchAssociatedTranscripts", params, optFns, c.addOperationSearchAssociatedTranscriptsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchAssociatedTranscriptsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchAssociatedTranscriptsInput struct {

	// The unique identifier of the bot associated with the transcripts that you are
	// searching.
	//
	// This member is required.
	BotId *string

	// The unique identifier of the bot recommendation associated with the transcripts
	// to search.
	//
	// This member is required.
	BotRecommendationId *string

	// The version of the bot containing the transcripts that you are searching.
	//
	// This member is required.
	BotVersion *string

	// A list of filter objects.
	//
	// This member is required.
	Filters []types.AssociatedTranscriptFilter

	// The identifier of the language and locale of the transcripts to search. The
	// string must match one of the supported locales. For more information, see
	// Supported languages (https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html)
	//
	// This member is required.
	LocaleId *string

	// The maximum number of bot recommendations to return in each page of results. If
	// there are fewer results than the max page size, only the actual number of
	// results are returned.
	MaxResults *int32

	// If the response from the SearchAssociatedTranscriptsRequest operation contains
	// more results than specified in the maxResults parameter, an index is returned in
	// the response. Use that index in the nextIndex parameter to return the next page
	// of results.
	NextIndex *int32

	// How SearchResults are ordered. Valid values are Ascending or Descending. The
	// default is Descending.
	SearchOrder types.SearchOrder

	noSmithyDocumentSerde
}

type SearchAssociatedTranscriptsOutput struct {

	// The object that contains the associated transcript that meet the criteria you
	// specified.
	AssociatedTranscripts []types.AssociatedTranscript

	// The unique identifier of the bot associated with the transcripts that you are
	// searching.
	BotId *string

	// The unique identifier of the bot recommendation associated with the transcripts
	// to search.
	BotRecommendationId *string

	// The version of the bot containing the transcripts that you are searching.
	BotVersion *string

	// The identifier of the language and locale of the transcripts to search. The
	// string must match one of the supported locales. For more information, see
	// Supported languages (https://docs.aws.amazon.com/lexv2/latest/dg/how-languages.html)
	LocaleId *string

	// A index that indicates whether there are more results to return in a response
	// to the SearchAssociatedTranscripts operation. If the nextIndex field is present,
	// you send the contents as the nextIndex parameter of a
	// SearchAssociatedTranscriptsRequest operation to get the next page of results.
	NextIndex *int32

	// The total number of transcripts returned by the search.
	TotalResults *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchAssociatedTranscriptsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchAssociatedTranscripts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchAssociatedTranscripts{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SearchAssociatedTranscripts"); err != nil {
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
	if err = addOpSearchAssociatedTranscriptsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchAssociatedTranscripts(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSearchAssociatedTranscripts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SearchAssociatedTranscripts",
	}
}
