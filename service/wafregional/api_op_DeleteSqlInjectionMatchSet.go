// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html)
// . With the latest version, AWS WAF has a single set of endpoints for regional
// and global use. Permanently deletes a SqlInjectionMatchSet . You can't delete a
// SqlInjectionMatchSet if it's still used in any Rules or if it still contains
// any SqlInjectionMatchTuple objects. If you just want to remove a
// SqlInjectionMatchSet from a Rule , use UpdateRule . To permanently delete a
// SqlInjectionMatchSet from AWS WAF, perform the following steps:
//   - Update the SqlInjectionMatchSet to remove filters, if any. For more
//     information, see UpdateSqlInjectionMatchSet .
//   - Use GetChangeToken to get the change token that you provide in the
//     ChangeToken parameter of a DeleteSqlInjectionMatchSet request.
//   - Submit a DeleteSqlInjectionMatchSet request.
func (c *Client) DeleteSqlInjectionMatchSet(ctx context.Context, params *DeleteSqlInjectionMatchSetInput, optFns ...func(*Options)) (*DeleteSqlInjectionMatchSetOutput, error) {
	if params == nil {
		params = &DeleteSqlInjectionMatchSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteSqlInjectionMatchSet", params, optFns, c.addOperationDeleteSqlInjectionMatchSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteSqlInjectionMatchSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to delete a SqlInjectionMatchSet from AWS WAF.
type DeleteSqlInjectionMatchSetInput struct {

	// The value returned by the most recent call to GetChangeToken .
	//
	// This member is required.
	ChangeToken *string

	// The SqlInjectionMatchSetId of the SqlInjectionMatchSet that you want to delete.
	// SqlInjectionMatchSetId is returned by CreateSqlInjectionMatchSet and by
	// ListSqlInjectionMatchSets .
	//
	// This member is required.
	SqlInjectionMatchSetId *string

	noSmithyDocumentSerde
}

// The response to a request to delete a SqlInjectionMatchSet from AWS WAF.
type DeleteSqlInjectionMatchSetOutput struct {

	// The ChangeToken that you used to submit the DeleteSqlInjectionMatchSet request.
	// You can also use this value to query the status of the request. For more
	// information, see GetChangeTokenStatus .
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteSqlInjectionMatchSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteSqlInjectionMatchSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteSqlInjectionMatchSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "DeleteSqlInjectionMatchSet"); err != nil {
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
	if err = addOpDeleteSqlInjectionMatchSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSqlInjectionMatchSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteSqlInjectionMatchSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "DeleteSqlInjectionMatchSet",
	}
}
