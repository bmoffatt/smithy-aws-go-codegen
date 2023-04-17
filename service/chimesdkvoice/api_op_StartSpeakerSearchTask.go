// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkvoice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chimesdkvoice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts a speaker search task. Before starting any speaker search tasks, you
// must provide all notices and obtain all consents from the speaker as required
// under applicable privacy and biometrics laws, and as required under the AWS
// service terms (https://aws.amazon.com/service-terms/) for the Amazon Chime SDK.
func (c *Client) StartSpeakerSearchTask(ctx context.Context, params *StartSpeakerSearchTaskInput, optFns ...func(*Options)) (*StartSpeakerSearchTaskOutput, error) {
	if params == nil {
		params = &StartSpeakerSearchTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartSpeakerSearchTask", params, optFns, c.addOperationStartSpeakerSearchTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartSpeakerSearchTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartSpeakerSearchTaskInput struct {

	// The transaction ID of the call being analyzed.
	//
	// This member is required.
	TransactionId *string

	// The Voice Connector ID.
	//
	// This member is required.
	VoiceConnectorId *string

	// The ID of the voice profile domain that will store the voice profile.
	//
	// This member is required.
	VoiceProfileDomainId *string

	// The unique identifier for the client request. Use a different token for
	// different speaker search tasks.
	ClientRequestToken *string

	noSmithyDocumentSerde
}

type StartSpeakerSearchTaskOutput struct {

	// The details of the speaker search task.
	SpeakerSearchTask *types.SpeakerSearchTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartSpeakerSearchTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartSpeakerSearchTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartSpeakerSearchTask{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpStartSpeakerSearchTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartSpeakerSearchTask(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opStartSpeakerSearchTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "StartSpeakerSearchTask",
	}
}