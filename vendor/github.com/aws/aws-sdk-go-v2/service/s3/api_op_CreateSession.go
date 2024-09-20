// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a session that establishes temporary security credentials to support
// fast authentication and authorization for the Zonal endpoint API operations on
// directory buckets. For more information about Zonal endpoint API operations that
// include the Availability Zone in the request endpoint, see [S3 Express One Zone APIs]in the Amazon S3
// User Guide.
//
// To make Zonal endpoint API requests on a directory bucket, use the CreateSession
// API operation. Specifically, you grant s3express:CreateSession permission to a
// bucket in a bucket policy or an IAM identity-based policy. Then, you use IAM
// credentials to make the CreateSession API request on the bucket, which returns
// temporary security credentials that include the access key ID, secret access
// key, session token, and expiration. These credentials have associated
// permissions to access the Zonal endpoint API operations. After the session is
// created, you don’t need to use other policies to grant permissions to each Zonal
// endpoint API individually. Instead, in your Zonal endpoint API requests, you
// sign your requests by applying the temporary security credentials of the session
// to the request headers and following the SigV4 protocol for authentication. You
// also apply the session token to the x-amz-s3session-token request header for
// authorization. Temporary security credentials are scoped to the bucket and
// expire after 5 minutes. After the expiration time, any calls that you make with
// those credentials will fail. You must use IAM credentials again to make a
// CreateSession API request that generates a new set of temporary credentials for
// use. Temporary credentials cannot be extended or refreshed beyond the original
// specified interval.
//
// If you use Amazon Web Services SDKs, SDKs handle the session token refreshes
// automatically to avoid service interruptions when a session expires. We
// recommend that you use the Amazon Web Services SDKs to initiate and manage
// requests to the CreateSession API. For more information, see [Performance guidelines and design patterns]in the Amazon S3
// User Guide.
//
//   - You must make requests for this API operation to the Zonal endpoint. These
//     endpoints support virtual-hosted-style requests in the format
//     https://bucket_name.s3express-az_id.region.amazonaws.com . Path-style requests
//     are not supported. For more information, see [Regional and Zonal endpoints]in the Amazon S3 User Guide.
//
//   - CopyObject API operation - Unlike other Zonal endpoint API operations, the
//     CopyObject API operation doesn't use the temporary security credentials
//     returned from the CreateSession API operation for authentication and
//     authorization. For information about authentication and authorization of the
//     CopyObject API operation on directory buckets, see [CopyObject].
//
//   - HeadBucket API operation - Unlike other Zonal endpoint API operations, the
//     HeadBucket API operation doesn't use the temporary security credentials
//     returned from the CreateSession API operation for authentication and
//     authorization. For information about authentication and authorization of the
//     HeadBucket API operation on directory buckets, see [HeadBucket].
//
// Permissions To obtain temporary security credentials, you must create a bucket
// policy or an IAM identity-based policy that grants s3express:CreateSession
// permission to the bucket. In a policy, you can have the s3express:SessionMode
// condition key to control who can create a ReadWrite or ReadOnly session. For
// more information about ReadWrite or ReadOnly sessions, see [x-amz-create-session-mode]
// x-amz-create-session-mode . For example policies, see [Example bucket policies for S3 Express One Zone] and [Amazon Web Services Identity and Access Management (IAM) identity-based policies for S3 Express One Zone] in the Amazon S3
// User Guide.
//
// To grant cross-account access to Zonal endpoint API operations, the bucket
// policy should also grant both accounts the s3express:CreateSession permission.
//
// If you want to encrypt objects with SSE-KMS, you must also have the
// kms:GenerateDataKey and the kms:Decrypt permissions in IAM identity-based
// policies and KMS key policies for the target KMS key.
//
// Encryption For directory buckets, there are only two supported options for
// server-side encryption: server-side encryption with Amazon S3 managed keys
// (SSE-S3) ( AES256 ) and server-side encryption with KMS keys (SSE-KMS) ( aws:kms
// ). We recommend that the bucket's default encryption uses the desired encryption
// configuration and you don't override the bucket default encryption in your
// CreateSession requests or PUT object requests. Then, new objects are
// automatically encrypted with the desired encryption settings. For more
// information, see [Protecting data with server-side encryption]in the Amazon S3 User Guide. For more information about the
// encryption overriding behaviors in directory buckets, see [Specifying server-side encryption with KMS for new object uploads].
//
// For [Zonal endpoint (object-level) API operations] except [CopyObject] and [UploadPartCopy], you authenticate and authorize requests through [CreateSession] for low
// latency. To encrypt new objects in a directory bucket with SSE-KMS, you must
// specify SSE-KMS as the directory bucket's default encryption configuration with
// a KMS key (specifically, a [customer managed key]). Then, when a session is created for Zonal
// endpoint API operations, new objects are automatically encrypted and decrypted
// with SSE-KMS and S3 Bucket Keys during the session.
//
// Only 1 [customer managed key] is supported per directory bucket for the lifetime of the bucket. [Amazon Web Services managed key] (
// aws/s3 ) isn't supported. After you specify SSE-KMS as your bucket's default
// encryption configuration with a customer managed key, you can't change the
// customer managed key for the bucket's SSE-KMS configuration.
//
// In the Zonal endpoint API calls (except [CopyObject] and [UploadPartCopy]) using the REST API, you can't
// override the values of the encryption settings ( x-amz-server-side-encryption ,
// x-amz-server-side-encryption-aws-kms-key-id ,
// x-amz-server-side-encryption-context , and
// x-amz-server-side-encryption-bucket-key-enabled ) from the CreateSession
// request. You don't need to explicitly specify these encryption settings values
// in Zonal endpoint API calls, and Amazon S3 will use the encryption settings
// values from the CreateSession request to protect new objects in the directory
// bucket.
//
// When you use the CLI or the Amazon Web Services SDKs, for CreateSession , the
// session token refreshes automatically to avoid service interruptions when a
// session expires. The CLI or the Amazon Web Services SDKs use the bucket's
// default encryption configuration for the CreateSession request. It's not
// supported to override the encryption settings values in the CreateSession
// request. Also, in the Zonal endpoint API calls (except [CopyObject]and [UploadPartCopy]), it's not
// supported to override the values of the encryption settings from the
// CreateSession request.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket_name.s3express-az_id.region.amazonaws.com .
//
// [Specifying server-side encryption with KMS for new object uploads]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-specifying-kms-encryption.html
// [Performance guidelines and design patterns]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-optimizing-performance-guidelines-design-patterns.html#s3-express-optimizing-performance-session-authentication
// [Regional and Zonal endpoints]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-Regions-and-Zones.html
// [CopyObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CopyObject.html
// [CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html
// [S3 Express One Zone APIs]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-APIs.html
// [HeadBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_HeadBucket.html
// [UploadPartCopy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPartCopy.html
// [Amazon Web Services managed key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk
// [Amazon Web Services Identity and Access Management (IAM) identity-based policies for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam-identity-policies.html
// [Example bucket policies for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam-example-bucket-policies.html
// [customer managed key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk
// [Protecting data with server-side encryption]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-serv-side-encryption.html
// [x-amz-create-session-mode]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html#API_CreateSession_RequestParameters
// [Zonal endpoint (object-level) API operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-differences.html#s3-express-differences-api-operations
func (c *Client) CreateSession(ctx context.Context, params *CreateSessionInput, optFns ...func(*Options)) (*CreateSessionOutput, error) {
	if params == nil {
		params = &CreateSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSession", params, optFns, c.addOperationCreateSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSessionInput struct {

	// The name of the bucket that you create a session for.
	//
	// This member is required.
	Bucket *string

	// Specifies whether Amazon S3 should use an S3 Bucket Key for object encryption
	// with server-side encryption using KMS keys (SSE-KMS).
	//
	// S3 Bucket Keys are always enabled for GET and PUT operations in a directory
	// bucket and can’t be disabled. S3 Bucket Keys aren't supported, when you copy
	// SSE-KMS encrypted objects from general purpose buckets to directory buckets,
	// from directory buckets to general purpose buckets, or between directory buckets,
	// through [CopyObject], [UploadPartCopy], [the Copy operation in Batch Operations], or [the import jobs]. In this case, Amazon S3 makes a call to KMS every time a
	// copy request is made for a KMS-encrypted object.
	//
	// [CopyObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CopyObject.html
	// [the import jobs]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/create-import-job
	// [UploadPartCopy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPartCopy.html
	// [the Copy operation in Batch Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/directory-buckets-objects-Batch-Ops
	BucketKeyEnabled *bool

	// Specifies the Amazon Web Services KMS Encryption Context as an additional
	// encryption context to use for object encryption. The value of this header is a
	// Base64-encoded string of a UTF-8 encoded JSON, which contains the encryption
	// context as key-value pairs. This value is stored as object metadata and
	// automatically gets passed on to Amazon Web Services KMS for future GetObject
	// operations on this object.
	//
	// General purpose buckets - This value must be explicitly added during CopyObject
	// operations if you want an additional encryption context for your object. For
	// more information, see [Encryption context]in the Amazon S3 User Guide.
	//
	// Directory buckets - You can optionally provide an explicit encryption context
	// value. The value must match the default encryption context - the bucket Amazon
	// Resource Name (ARN). An additional encryption context value is not supported.
	//
	// [Encryption context]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/UsingKMSEncryption.html#encryption-context
	SSEKMSEncryptionContext *string

	// If you specify x-amz-server-side-encryption with aws:kms , you must specify the
	// x-amz-server-side-encryption-aws-kms-key-id header with the ID (Key ID or Key
	// ARN) of the KMS symmetric encryption customer managed key to use. Otherwise, you
	// get an HTTP 400 Bad Request error. Only use the key ID or key ARN. The key
	// alias format of the KMS key isn't supported. Also, if the KMS key doesn't exist
	// in the same account that't issuing the command, you must use the full Key ARN
	// not the Key ID.
	//
	// Your SSE-KMS configuration can only support 1 [customer managed key] per directory bucket for the
	// lifetime of the bucket. [Amazon Web Services managed key]( aws/s3 ) isn't supported.
	//
	// [customer managed key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk
	// [Amazon Web Services managed key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk
	SSEKMSKeyId *string

	// The server-side encryption algorithm to use when you store objects in the
	// directory bucket.
	//
	// For directory buckets, there are only two supported options for server-side
	// encryption: server-side encryption with Amazon S3 managed keys (SSE-S3) ( AES256
	// ) and server-side encryption with KMS keys (SSE-KMS) ( aws:kms ). By default,
	// Amazon S3 encrypts data with SSE-S3. For more information, see [Protecting data with server-side encryption]in the Amazon S3
	// User Guide.
	//
	// [Protecting data with server-side encryption]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-serv-side-encryption.html
	ServerSideEncryption types.ServerSideEncryption

	// Specifies the mode of the session that will be created, either ReadWrite or
	// ReadOnly . By default, a ReadWrite session is created. A ReadWrite session is
	// capable of executing all the Zonal endpoint API operations on a directory
	// bucket. A ReadOnly session is constrained to execute the following Zonal
	// endpoint API operations: GetObject , HeadObject , ListObjectsV2 ,
	// GetObjectAttributes , ListParts , and ListMultipartUploads .
	SessionMode types.SessionMode

	noSmithyDocumentSerde
}

func (in *CreateSessionInput) bindEndpointParams(p *EndpointParameters) {

	p.Bucket = in.Bucket
	p.DisableS3ExpressSessionAuth = ptr.Bool(true)
}

type CreateSessionOutput struct {

	// The established temporary security credentials for the created session.
	//
	// This member is required.
	Credentials *types.SessionCredentials

	// Indicates whether to use an S3 Bucket Key for server-side encryption with KMS
	// keys (SSE-KMS).
	BucketKeyEnabled *bool

	// If present, indicates the Amazon Web Services KMS Encryption Context to use for
	// object encryption. The value of this header is a Base64-encoded string of a
	// UTF-8 encoded JSON, which contains the encryption context as key-value pairs.
	// This value is stored as object metadata and automatically gets passed on to
	// Amazon Web Services KMS for future GetObject operations on this object.
	SSEKMSEncryptionContext *string

	// If you specify x-amz-server-side-encryption with aws:kms , this header indicates
	// the ID of the KMS symmetric encryption customer managed key that was used for
	// object encryption.
	SSEKMSKeyId *string

	// The server-side encryption algorithm used when you store objects in the
	// directory bucket.
	ServerSideEncryption types.ServerSideEncryption

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateSession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateSession{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateSession"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
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
	if err = addPutBucketContextMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addIsExpressUserAgent(stack); err != nil {
		return err
	}
	if err = addOpCreateSessionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSession(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addCreateSessionUpdateEndpoint(stack, options); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSerializeImmutableHostnameBucketMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func (v *CreateSessionInput) bucket() (string, bool) {
	if v.Bucket == nil {
		return "", false
	}
	return *v.Bucket, true
}

func newServiceMetadataMiddleware_opCreateSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateSession",
	}
}

// getCreateSessionBucketMember returns a pointer to string denoting a provided
// bucket member valueand a boolean indicating if the input has a modeled bucket
// name,
func getCreateSessionBucketMember(input interface{}) (*string, bool) {
	in := input.(*CreateSessionInput)
	if in.Bucket == nil {
		return nil, false
	}
	return in.Bucket, true
}
func addCreateSessionUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{
			GetBucketFromInput: getCreateSessionBucketMember,
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