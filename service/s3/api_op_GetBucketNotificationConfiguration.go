// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the notification configuration of a bucket. If notifications are not
// enabled on the bucket, the action returns an empty NotificationConfiguration
// element. By default, you must be the bucket owner to read the notification
// configuration of a bucket. However, the bucket owner can use a bucket policy to
// grant permission to other users to read this configuration with the
// s3:GetBucketNotification permission. For more information about setting and
// reading the notification configuration on a bucket, see Setting Up Notification
// of Bucket Events
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html). For
// more information about bucket policies, see Using Bucket Policies
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html). The
// following action is related to GetBucketNotification:
//
// * PutBucketNotification
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketNotification.html)
func (c *Client) GetBucketNotificationConfiguration(ctx context.Context, params *GetBucketNotificationConfigurationInput, optFns ...func(*Options)) (*GetBucketNotificationConfigurationOutput, error) {
	if params == nil {
		params = &GetBucketNotificationConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketNotificationConfiguration", params, optFns, c.addOperationGetBucketNotificationConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketNotificationConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketNotificationConfigurationInput struct {

	// The name of the bucket for which to get the notification configuration.
	//
	// This member is required.
	Bucket *string

	// The account ID of the expected bucket owner. If the bucket is owned by a
	// different account, the request fails with the HTTP status code 403 Forbidden
	// (access denied).
	ExpectedBucketOwner *string

	noSmithyDocumentSerde
}

// A container for specifying the notification configuration of the bucket. If this
// element is empty, notifications are turned off for the bucket.
type GetBucketNotificationConfigurationOutput struct {

	// Enables delivery of events to Amazon EventBridge.
	EventBridgeConfiguration *types.EventBridgeConfiguration

	// Describes the Lambda functions to invoke and the events for which to invoke
	// them.
	LambdaFunctionConfigurations []types.LambdaFunctionConfiguration

	// The Amazon Simple Queue Service queues to publish messages to and the events for
	// which to publish messages.
	QueueConfigurations []types.QueueConfiguration

	// The topic to which notifications are sent and the events for which notifications
	// are generated.
	TopicConfigurations []types.TopicConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBucketNotificationConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetBucketNotificationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketNotificationConfiguration{}, middleware.After)
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
	if err = swapWithCustomHTTPSignerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpGetBucketNotificationConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketNotificationConfiguration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = add100Continue(stack); err != nil {
		return err
	}
	if err = addGetBucketNotificationConfigurationUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetBucketNotificationConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketNotificationConfiguration",
	}
}

// getGetBucketNotificationConfigurationBucketMember returns a pointer to string
// denoting a provided bucket member valueand a boolean indicating if the input has
// a modeled bucket name,
func getGetBucketNotificationConfigurationBucketMember(input interface{}) (*string, bool) {
	in := input.(*GetBucketNotificationConfigurationInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addGetBucketNotificationConfigurationUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getGetBucketNotificationConfigurationBucketMember,
		},
		UsePathStyle:                   options.UsePathStyle,
		UseAccelerate:                  options.UseAccelerate,
		SupportsAccelerate:             true,
		TargetS3ObjectLambda:           false,
		EndpointResolver:               options.EndpointResolver,
		EndpointResolverOptions:        options.EndpointOptions,
		UseARNRegion:                   options.UseARNRegion,
		DisableMultiRegionAccessPoints: options.DisableMultiRegionAccessPoints,
	})
}
