// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// An activity that adds other attributes based on existing attributes in the
// message.
type AddAttributesActivity struct {
	// The name of the 'addAttributes' activity.
	Name *string
	// A list of 1-50 "AttributeNameMapping" objects that map an existing attribute to
	// a new attribute. The existing attributes remain in the message, so if you want
	// to remove the originals, use "RemoveAttributeActivity".
	Attributes map[string]*string
	// The next activity in the pipeline.
	Next *string
}

// Contains informations about errors.
type BatchPutMessageErrorEntry struct {
	// The message associated with the error.
	ErrorMessage *string
	// The ID of the message that caused the error. (See the value corresponding to the
	// "messageId" key in the message object.)
	MessageId *string
	// The code associated with the error.
	ErrorCode *string
}

// A collection of data from an MQTT topic. Channels archive the raw, unprocessed
// messages before publishing the data to a pipeline.
type Channel struct {
	// When the channel was last updated.
	LastUpdateTime *time.Time
	// Where channel data is stored. You may choose one of "serviceManagedS3" or
	// "customerManagedS3" storage. If not specified, the default is
	// "serviceManagedS3". This cannot be changed after creation of the channel.
	Storage *ChannelStorage
	// The name of the channel.
	Name *string
	// The ARN of the channel.
	Arn *string
	// The status of the channel.
	Status ChannelStatus
	// How long, in days, message data is kept for the channel.
	RetentionPeriod *RetentionPeriod
	// When the channel was created.
	CreationTime *time.Time
}

// The activity that determines the source of the messages to be processed.
type ChannelActivity struct {
	// The name of the 'channel' activity.
	Name *string
	// The name of the channel from which the messages are processed.
	ChannelName *string
	// The next activity in the pipeline.
	Next *string
}

// Statistics information about the channel.
type ChannelStatistics struct {
	// The estimated size of the channel.
	Size *EstimatedResourceSize
}

// Where channel data is stored. You may choose one of "serviceManagedS3" or
// "customerManagedS3" storage. If not specified, the default is
// "serviceManagedS3". This cannot be changed after creation of the channel.
type ChannelStorage struct {
	// Use this to store channel data in an S3 bucket managed by the AWS IoT Analytics
	// service. The choice of service-managed or customer-managed S3 storage cannot be
	// changed after creation of the channel.
	ServiceManagedS3 *ServiceManagedChannelS3Storage
	// Use this to store channel data in an S3 bucket that you manage. If customer
	// managed storage is selected, the "retentionPeriod" parameter is ignored. The
	// choice of service-managed or customer-managed S3 storage cannot be changed after
	// creation of the channel.
	CustomerManagedS3 *CustomerManagedChannelS3Storage
}

// Where channel data is stored.
type ChannelStorageSummary struct {
	// Used to store channel data in an S3 bucket managed by the AWS IoT Analytics
	// service.
	ServiceManagedS3 *ServiceManagedChannelS3StorageSummary
	// Used to store channel data in an S3 bucket that you manage.
	CustomerManagedS3 *CustomerManagedChannelS3StorageSummary
}

// A summary of information about a channel.
type ChannelSummary struct {
	// The status of the channel.
	Status ChannelStatus
	// When the channel was created.
	CreationTime *time.Time
	// Where channel data is stored.
	ChannelStorage *ChannelStorageSummary
	// The name of the channel.
	ChannelName *string
	// The last time the channel was updated.
	LastUpdateTime *time.Time
}

// Information needed to run the "containerAction" to produce data set contents.
type ContainerDatasetAction struct {
	// The ARN of the role which gives permission to the system to access needed
	// resources in order to run the "containerAction". This includes, at minimum,
	// permission to retrieve the data set contents which are the input to the
	// containerized application.
	ExecutionRoleArn *string
	// The values of variables used within the context of the execution of the
	// containerized application (basically, parameters passed to the application).
	// Each variable must have a name and a value given by one of "stringValue",
	// "datasetContentVersionValue", or "outputFileUriValue".
	Variables []*Variable
	// Configuration of the resource which executes the "containerAction".
	ResourceConfiguration *ResourceConfiguration
	// The ARN of the Docker container stored in your account. The Docker container
	// contains an application and needed support libraries and is used to generate
	// data set contents.
	Image *string
}

// Use this to store channel data in an S3 bucket that you manage. If customer
// managed storage is selected, the "retentionPeriod" parameter is ignored. The
// choice of service-managed or customer-managed S3 storage cannot be changed after
// creation of the channel.
type CustomerManagedChannelS3Storage struct {
	// The name of the Amazon S3 bucket in which channel data is stored.
	Bucket *string
	// [Optional] The prefix used to create the keys of the channel data objects. Each
	// object in an Amazon S3 bucket has a key that is its unique identifier within the
	// bucket (each object in a bucket has exactly one key). The prefix must end with a
	// '/'.
	KeyPrefix *string
	// The ARN of the role which grants AWS IoT Analytics permission to interact with
	// your Amazon S3 resources.
	RoleArn *string
}

// Used to store channel data in an S3 bucket that you manage.
type CustomerManagedChannelS3StorageSummary struct {
	// The ARN of the role which grants AWS IoT Analytics permission to interact with
	// your Amazon S3 resources.
	RoleArn *string
	// [Optional] The prefix used to create the keys of the channel data objects. Each
	// object in an Amazon S3 bucket has a key that is its unique identifier within the
	// bucket (each object in a bucket has exactly one key). The prefix must end with a
	// '/'.
	KeyPrefix *string
	// The name of the Amazon S3 bucket in which channel data is stored.
	Bucket *string
}

// Use this to store data store data in an S3 bucket that you manage. When customer
// managed storage is selected, the "retentionPeriod" parameter is ignored. The
// choice of service-managed or customer-managed S3 storage cannot be changed after
// creation of the data store.
type CustomerManagedDatastoreS3Storage struct {
	// The ARN of the role which grants AWS IoT Analytics permission to interact with
	// your Amazon S3 resources.
	RoleArn *string
	// [Optional] The prefix used to create the keys of the data store data objects.
	// Each object in an Amazon S3 bucket has a key that is its unique identifier
	// within the bucket (each object in a bucket has exactly one key). The prefix must
	// end with a '/'.
	KeyPrefix *string
	// The name of the Amazon S3 bucket in which data store data is stored.
	Bucket *string
}

// Used to store data store data in an S3 bucket that you manage.
type CustomerManagedDatastoreS3StorageSummary struct {
	// The name of the Amazon S3 bucket in which data store data is stored.
	Bucket *string
	// [Optional] The prefix used to create the keys of the data store data objects.
	// Each object in an Amazon S3 bucket has a key that is its unique identifier
	// within the bucket (each object in a bucket has exactly one key). The prefix must
	// end with a '/'.
	KeyPrefix *string
	// The ARN of the role which grants AWS IoT Analytics permission to interact with
	// your Amazon S3 resources.
	RoleArn *string
}

// Information about a data set.
type Dataset struct {
	// [Optional] How many versions of data set contents are kept. If not specified or
	// set to null, only the latest version plus the latest succeeded version (if they
	// are different) are kept for the time period specified by the "retentionPeriod"
	// parameter. (For more information, see
	// https://docs.aws.amazon.com/iotanalytics/latest/userguide/getting-started.html#aws-iot-analytics-dataset-versions)
	VersioningConfiguration *VersioningConfiguration
	// The status of the data set.
	Status DatasetStatus
	// The "DatasetTrigger" objects that specify when the data set is automatically
	// updated.
	Triggers []*DatasetTrigger
	// [Optional] How long, in days, message data is kept for the data set.
	RetentionPeriod *RetentionPeriod
	// The name of the data set.
	Name *string
	// The "DatasetAction" objects that automatically create the data set contents.
	Actions []*DatasetAction
	// The ARN of the data set.
	Arn *string
	// The last time the data set was updated.
	LastUpdateTime *time.Time
	// When data set contents are created they are delivered to destinations specified
	// here.
	ContentDeliveryRules []*DatasetContentDeliveryRule
	// When the data set was created.
	CreationTime *time.Time
}

// A "DatasetAction" object that specifies how data set contents are automatically
// created.
type DatasetAction struct {
	// Information which allows the system to run a containerized application in order
	// to create the data set contents. The application must be in a Docker container
	// along with any needed support libraries.
	ContainerAction *ContainerDatasetAction
	// An "SqlQueryDatasetAction" object that uses an SQL query to automatically create
	// data set contents.
	QueryAction *SqlQueryDatasetAction
	// The name of the data set action by which data set contents are automatically
	// created.
	ActionName *string
}

// Information about the action which automatically creates the data set's
// contents.
type DatasetActionSummary struct {
	// The type of action by which the data set's contents are automatically created.
	ActionType DatasetActionType
	// The name of the action which automatically creates the data set's contents.
	ActionName *string
}

// The destination to which data set contents are delivered.
type DatasetContentDeliveryDestination struct {
	// Configuration information for delivery of data set contents to Amazon S3.
	S3DestinationConfiguration *S3DestinationConfiguration
	// Configuration information for delivery of data set contents to AWS IoT Events.
	IotEventsDestinationConfiguration *IotEventsDestinationConfiguration
}

// When data set contents are created they are delivered to destination specified
// here.
type DatasetContentDeliveryRule struct {
	// The destination to which data set contents are delivered.
	Destination *DatasetContentDeliveryDestination
	// The name of the data set content delivery rules entry.
	EntryName *string
}

// The state of the data set contents and the reason they are in this state.
type DatasetContentStatus struct {
	// The reason the data set contents are in this state.
	Reason *string
	// The state of the data set contents. Can be one of "READY", "CREATING",
	// "SUCCEEDED" or "FAILED".
	State DatasetContentState
}

// Summary information about data set contents.
type DatasetContentSummary struct {
	// The status of the data set contents.
	Status *DatasetContentStatus
	// The version of the data set contents.
	Version *string
	// The time the dataset content status was updated to SUCCEEDED or FAILED.
	CompletionTime *time.Time
	// The time the creation of the data set contents was scheduled to start.
	ScheduleTime *time.Time
	// The actual time the creation of the data set contents was started.
	CreationTime *time.Time
}

// The data set whose latest contents are used as input to the notebook or
// application.
type DatasetContentVersionValue struct {
	// The name of the data set whose latest contents are used as input to the notebook
	// or application.
	DatasetName *string
}

// The reference to a data set entry.
type DatasetEntry struct {
	// The name of the data set item.
	EntryName *string
	// The pre-signed URI of the data set item.
	DataURI *string
}

// A summary of information about a data set.
type DatasetSummary struct {
	// The time the data set was created.
	CreationTime *time.Time
	// The status of the data set.
	Status DatasetStatus
	// A list of triggers. A trigger causes data set content to be populated at a
	// specified time interval or when another data set is populated. The list of
	// triggers can be empty or contain up to five DataSetTrigger objects
	Triggers []*DatasetTrigger
	// The last time the data set was updated.
	LastUpdateTime *time.Time
	// The name of the data set.
	DatasetName *string
	// A list of "DataActionSummary" objects.
	Actions []*DatasetActionSummary
}

// The "DatasetTrigger" that specifies when the data set is automatically updated.
type DatasetTrigger struct {
	// The "Schedule" when the trigger is initiated.
	Schedule *Schedule
	// The data set whose content creation triggers the creation of this data set's
	// contents.
	Dataset *TriggeringDataset
}

// Information about a data store.
type Datastore struct {
	// Where data store data is stored. You may choose one of "serviceManagedS3" or
	// "customerManagedS3" storage. If not specified, the default is
	// "serviceManagedS3". This cannot be changed after the data store is created.
	Storage *DatastoreStorage
	// When the data store was created.
	CreationTime *time.Time
	// The name of the data store.
	Name *string
	// The ARN of the data store.
	Arn *string
	// The last time the data store was updated.
	LastUpdateTime *time.Time
	// How long, in days, message data is kept for the data store. When
	// "customerManagedS3" storage is selected, this parameter is ignored.
	RetentionPeriod *RetentionPeriod
	// The status of a data store: CREATING The data store is being created. ACTIVE The
	// data store has been created and can be used. DELETING The data store is being
	// deleted.
	Status DatastoreStatus
}

// The 'datastore' activity that specifies where to store the processed data.
type DatastoreActivity struct {
	// The name of the data store where processed messages are stored.
	DatastoreName *string
	// The name of the 'datastore' activity.
	Name *string
}

// Statistical information about the data store.
type DatastoreStatistics struct {
	// The estimated size of the data store.
	Size *EstimatedResourceSize
}

// Where data store data is stored. You may choose one of "serviceManagedS3" or
// "customerManagedS3" storage. If not specified, the default is
// "serviceManagedS3". This cannot be changed after the data store is created.
type DatastoreStorage struct {
	// Use this to store data store data in an S3 bucket that you manage. When customer
	// managed storage is selected, the "retentionPeriod" parameter is ignored. The
	// choice of service-managed or customer-managed S3 storage cannot be changed after
	// creation of the data store.
	CustomerManagedS3 *CustomerManagedDatastoreS3Storage
	// Use this to store data store data in an S3 bucket managed by the AWS IoT
	// Analytics service. The choice of service-managed or customer-managed S3 storage
	// cannot be changed after creation of the data store.
	ServiceManagedS3 *ServiceManagedDatastoreS3Storage
}

// Where data store data is stored.
type DatastoreStorageSummary struct {
	// Used to store data store data in an S3 bucket that you manage.
	CustomerManagedS3 *CustomerManagedDatastoreS3StorageSummary
	// Used to store data store data in an S3 bucket managed by the AWS IoT Analytics
	// service.
	ServiceManagedS3 *ServiceManagedDatastoreS3StorageSummary
}

// A summary of information about a data store.
type DatastoreSummary struct {
	// Where data store data is stored.
	DatastoreStorage *DatastoreStorageSummary
	// When the data store was created.
	CreationTime *time.Time
	// The status of the data store.
	Status DatastoreStatus
	// The last time the data store was updated.
	LastUpdateTime *time.Time
	// The name of the data store.
	DatastoreName *string
}

// Used to limit data to that which has arrived since the last execution of the
// action.
type DeltaTime struct {
	// The number of seconds of estimated "in flight" lag time of message data. When
	// you create data set contents using message data from a specified time frame,
	// some message data may still be "in flight" when processing begins, and so will
	// not arrive in time to be processed. Use this field to make allowances for the
	// "in flight" time of your message data, so that data not processed from a
	// previous time frame will be included with the next time frame. Without this,
	// missed message data would be excluded from processing during the next time frame
	// as well, because its timestamp places it within the previous time frame.
	OffsetSeconds *int32
	// An expression by which the time of the message data may be determined. This may
	// be the name of a timestamp field, or a SQL expression which is used to derive
	// the time the message data was generated.
	TimeExpression *string
}

// An activity that adds data from the AWS IoT device registry to your message.
type DeviceRegistryEnrichActivity struct {
	// The name of the attribute that is added to the message.
	Attribute *string
	// The ARN of the role that allows access to the device's registry information.
	RoleArn *string
	// The name of the 'deviceRegistryEnrich' activity.
	Name *string
	// The name of the IoT device whose registry information is added to the message.
	ThingName *string
	// The next activity in the pipeline.
	Next *string
}

// An activity that adds information from the AWS IoT Device Shadows service to a
// message.
type DeviceShadowEnrichActivity struct {
	// The name of the 'deviceShadowEnrich' activity.
	Name *string
	// The name of the IoT device whose shadow information is added to the message.
	ThingName *string
	// The ARN of the role that allows access to the device's shadow.
	RoleArn *string
	// The name of the attribute that is added to the message.
	Attribute *string
	// The next activity in the pipeline.
	Next *string
}

// The estimated size of the resource.
type EstimatedResourceSize struct {
	// The time when the estimate of the size of the resource was made.
	EstimatedOn *time.Time
	// The estimated size of the resource in bytes.
	EstimatedSizeInBytes *float64
}

// An activity that filters a message based on its attributes.
type FilterActivity struct {
	// An expression that looks like a SQL WHERE clause that must return a Boolean
	// value.
	Filter *string
	// The next activity in the pipeline.
	Next *string
	// The name of the 'filter' activity.
	Name *string
}

// Configuration information for coordination with the AWS Glue ETL (extract,
// transform and load) service.
type GlueConfiguration struct {
	// The name of the database in your AWS Glue Data Catalog in which the table is
	// located. (An AWS Glue Data Catalog database contains Glue Data tables.)
	DatabaseName *string
	// The name of the table in your AWS Glue Data Catalog which is used to perform the
	// ETL (extract, transform and load) operations. (An AWS Glue Data Catalog table
	// contains partitioned data and descriptions of data sources and targets.)
	TableName *string
}

// Configuration information for delivery of data set contents to AWS IoT Events.
type IotEventsDestinationConfiguration struct {
	// The ARN of the role which grants AWS IoT Analytics permission to deliver data
	// set contents to an AWS IoT Events input.
	RoleArn *string
	// The name of the AWS IoT Events input to which data set contents are delivered.
	InputName *string
}

// An activity that runs a Lambda function to modify the message.
type LambdaActivity struct {
	// The next activity in the pipeline.
	Next *string
	// The number of messages passed to the Lambda function for processing. The AWS
	// Lambda function must be able to process all of these messages within five
	// minutes, which is the maximum timeout duration for Lambda functions.
	BatchSize *int32
	// The name of the 'lambda' activity.
	Name *string
	// The name of the Lambda function that is run on the message.
	LambdaName *string
}

// Information about logging options.
type LoggingOptions struct {
	// The logging level. Currently, only "ERROR" is supported.
	Level LoggingLevel
	// If true, logging is enabled for AWS IoT Analytics.
	Enabled *bool
	// The ARN of the role that grants permission to AWS IoT Analytics to perform
	// logging.
	RoleArn *string
}

// An activity that computes an arithmetic expression using the message's
// attributes.
type MathActivity struct {
	// The name of the attribute that contains the result of the math operation.
	Attribute *string
	// An expression that uses one or more existing attributes and must return an
	// integer value.
	Math *string
	// The name of the 'math' activity.
	Name *string
	// The next activity in the pipeline.
	Next *string
}

// Information about a message.
type Message struct {
	// The ID you wish to assign to the message. Each "messageId" must be unique within
	// each batch sent.
	MessageId *string
	// The payload of the message. This may be a JSON string or a Base64-encoded string
	// representing binary data (in which case you must decode it by means of a
	// pipeline activity).
	Payload []byte
}

// The value of the variable as a structure that specifies an output file URI.
type OutputFileUriValue struct {
	// The URI of the location where data set contents are stored, usually the URI of a
	// file in an S3 bucket.
	FileName *string
}

// Contains information about a pipeline.
type Pipeline struct {
	// When the pipeline was created.
	CreationTime *time.Time
	// The activities that perform transformations on the messages.
	Activities []*PipelineActivity
	// The ARN of the pipeline.
	Arn *string
	// The name of the pipeline.
	Name *string
	// The last time the pipeline was updated.
	LastUpdateTime *time.Time
	// A summary of information about the pipeline reprocessing.
	ReprocessingSummaries []*ReprocessingSummary
}

// An activity that performs a transformation on a message.
type PipelineActivity struct {
	// Adds information from the AWS IoT Device Shadows service to a message.
	DeviceShadowEnrich *DeviceShadowEnrichActivity
	// Filters a message based on its attributes.
	Filter *FilterActivity
	// Computes an arithmetic expression using the message's attributes and adds it to
	// the message.
	Math *MathActivity
	// Removes attributes from a message.
	RemoveAttributes *RemoveAttributesActivity
	// Runs a Lambda function to modify the message.
	Lambda *LambdaActivity
	// Determines the source of the messages to be processed.
	Channel *ChannelActivity
	// Adds other attributes based on existing attributes in the message.
	AddAttributes *AddAttributesActivity
	// Specifies where to store the processed message data.
	Datastore *DatastoreActivity
	// Creates a new message using only the specified attributes from the original
	// message.
	SelectAttributes *SelectAttributesActivity
	// Adds data from the AWS IoT device registry to your message.
	DeviceRegistryEnrich *DeviceRegistryEnrichActivity
}

// A summary of information about a pipeline.
type PipelineSummary struct {
	// When the pipeline was created.
	CreationTime *time.Time
	// When the pipeline was last updated.
	LastUpdateTime *time.Time
	// A summary of information about the pipeline reprocessing.
	ReprocessingSummaries []*ReprocessingSummary
	// The name of the pipeline.
	PipelineName *string
}

// Information which is used to filter message data, to segregate it according to
// the time frame in which it arrives.
type QueryFilter struct {
	// Used to limit data to that which has arrived since the last execution of the
	// action.
	DeltaTime *DeltaTime
}

// An activity that removes attributes from a message.
type RemoveAttributesActivity struct {
	// The next activity in the pipeline.
	Next *string
	// A list of 1-50 attributes to remove from the message.
	Attributes []*string
	// The name of the 'removeAttributes' activity.
	Name *string
}

// Information about pipeline reprocessing.
type ReprocessingSummary struct {
	// The status of the pipeline reprocessing.
	Status ReprocessingStatus
	// The 'reprocessingId' returned by "StartPipelineReprocessing".
	Id *string
	// The time the pipeline reprocessing was created.
	CreationTime *time.Time
}

// The configuration of the resource used to execute the "containerAction".
type ResourceConfiguration struct {
	// The size (in GB) of the persistent storage available to the resource instance
	// used to execute the "containerAction" (min: 1, max: 50).
	VolumeSizeInGB *int32
	// The type of the compute resource used to execute the "containerAction". Possible
	// values are: ACU_1 (vCPU=4, memory=16GiB) or ACU_2 (vCPU=8, memory=32GiB).
	ComputeType ComputeType
}

// How long, in days, message data is kept.
type RetentionPeriod struct {
	// If true, message data is kept indefinitely.
	Unlimited *bool
	// The number of days that message data is kept. The "unlimited" parameter must be
	// false.
	NumberOfDays *int32
}

// Configuration information for delivery of data set contents to Amazon S3.
type S3DestinationConfiguration struct {
	// Configuration information for coordination with the AWS Glue ETL (extract,
	// transform and load) service.
	GlueConfiguration *GlueConfiguration
	// The ARN of the role which grants AWS IoT Analytics permission to interact with
	// your Amazon S3 and AWS Glue resources.
	RoleArn *string
	// The key of the data set contents object. Each object in an Amazon S3 bucket has
	// a key that is its unique identifier within the bucket (each object in a bucket
	// has exactly one key). To produce a unique key, you can use
	// "!{iotanalytics:scheduledTime}" to insert the time of the scheduled SQL query
	// run, or "!{iotanalytics:versioned} to insert a unique hash identifying the data
	// set, for example:
	// "/DataSet/!{iotanalytics:scheduledTime}/!{iotanalytics:versioned}.csv".
	Key *string
	// The name of the Amazon S3 bucket to which data set contents are delivered.
	Bucket *string
}

// The schedule for when to trigger an update.
type Schedule struct {
	// The expression that defines when to trigger an update. For more information, see
	// Schedule Expressions for Rules
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/ScheduledEvents.html)
	// in the Amazon CloudWatch Events User Guide.
	Expression *string
}

// Creates a new message using only the specified attributes from the original
// message.
type SelectAttributesActivity struct {
	// The name of the 'selectAttributes' activity.
	Name *string
	// A list of the attributes to select from the message.
	Attributes []*string
	// The next activity in the pipeline.
	Next *string
}

// Use this to store channel data in an S3 bucket managed by the AWS IoT Analytics
// service. The choice of service-managed or customer-managed S3 storage cannot be
// changed after creation of the channel.
type ServiceManagedChannelS3Storage struct {
}

// Used to store channel data in an S3 bucket managed by the AWS IoT Analytics
// service.
type ServiceManagedChannelS3StorageSummary struct {
}

// Use this to store data store data in an S3 bucket managed by the AWS IoT
// Analytics service. The choice of service-managed or customer-managed S3 storage
// cannot be changed after creation of the data store.
type ServiceManagedDatastoreS3Storage struct {
}

// Used to store data store data in an S3 bucket managed by the AWS IoT Analytics
// service.
type ServiceManagedDatastoreS3StorageSummary struct {
}

// The SQL query to modify the message.
type SqlQueryDatasetAction struct {
	// Pre-filters applied to message data.
	Filters []*QueryFilter
	// A SQL query string.
	SqlQuery *string
}

// A set of key/value pairs which are used to manage the resource.
type Tag struct {
	// The tag's value.
	Value *string
	// The tag's key.
	Key *string
}

// Information about the data set whose content generation triggers the new data
// set content generation.
type TriggeringDataset struct {
	// The name of the data set whose content generation triggers the new data set
	// content generation.
	Name *string
}

// An instance of a variable to be passed to the "containerAction" execution. Each
// variable must have a name and a value given by one of "stringValue",
// "datasetContentVersionValue", or "outputFileUriValue".
type Variable struct {
	// The value of the variable as a string.
	StringValue *string
	// The value of the variable as a structure that specifies an output file URI.
	OutputFileUriValue *OutputFileUriValue
	// The value of the variable as a double (numeric).
	DoubleValue *float64
	// The value of the variable as a structure that specifies a data set content
	// version.
	DatasetContentVersionValue *DatasetContentVersionValue
	// The name of the variable.
	Name *string
}

// Information about the versioning of data set contents.
type VersioningConfiguration struct {
	// How many versions of data set contents will be kept. The "unlimited" parameter
	// must be false.
	MaxVersions *int32
	// If true, unlimited versions of data set contents will be kept.
	Unlimited *bool
}