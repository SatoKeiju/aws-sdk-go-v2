// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Reserves an open player slot in a game session for a player. New player
// sessions can be created in any game session with an open slot that is in ACTIVE
// status and has a player creation policy of ACCEPT_ALL . You can add a group of
// players to a game session with CreatePlayerSessions (https://docs.aws.amazon.com/gamelift/latest/apireference/API_CreatePlayerSessions.html)
// . To create a player session, specify a game session ID, player ID, and
// optionally a set of player data. If successful, a slot is reserved in the game
// session for the player and a new PlayerSessions object is returned with a
// player session ID. The player references the player session ID when sending a
// connection request to the game session, and the game server can use it to
// validate the player reservation with the Amazon GameLift service. Player
// sessions cannot be updated. The maximum number of players per game session is
// 200. It is not adjustable. Related actions All APIs by task (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) CreatePlayerSession(ctx context.Context, params *CreatePlayerSessionInput, optFns ...func(*Options)) (*CreatePlayerSessionOutput, error) {
	if params == nil {
		params = &CreatePlayerSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePlayerSession", params, optFns, c.addOperationCreatePlayerSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePlayerSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePlayerSessionInput struct {

	// A unique identifier for the game session to add a player to.
	//
	// This member is required.
	GameSessionId *string

	// A unique identifier for a player. Player IDs are developer-defined.
	//
	// This member is required.
	PlayerId *string

	// Developer-defined information related to a player. Amazon GameLift does not use
	// this data, so it can be formatted as needed for use in the game.
	PlayerData *string

	noSmithyDocumentSerde
}

type CreatePlayerSessionOutput struct {

	// Object that describes the newly created player session record.
	PlayerSession *types.PlayerSession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreatePlayerSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePlayerSession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePlayerSession{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreatePlayerSession"); err != nil {
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
	if err = addOpCreatePlayerSessionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePlayerSession(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreatePlayerSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreatePlayerSession",
	}
}
