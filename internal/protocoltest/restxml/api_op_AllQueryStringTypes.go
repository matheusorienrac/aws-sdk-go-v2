// Code generated by smithy-go-codegen DO NOT EDIT.

package restxml

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxml/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// This example uses all query string types.
func (c *Client) AllQueryStringTypes(ctx context.Context, params *AllQueryStringTypesInput, optFns ...func(*Options)) (*AllQueryStringTypesOutput, error) {
	if params == nil {
		params = &AllQueryStringTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AllQueryStringTypes", params, optFns, c.addOperationAllQueryStringTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AllQueryStringTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AllQueryStringTypesInput struct {
	QueryBoolean *bool

	QueryBooleanList []bool

	QueryByte *int8

	QueryDouble *float64

	QueryDoubleList []float64

	QueryEnum types.FooEnum

	QueryEnumList []types.FooEnum

	QueryFloat *float32

	QueryInteger *int32

	QueryIntegerEnum types.IntegerEnum

	QueryIntegerEnumList []types.IntegerEnum

	QueryIntegerList []int32

	QueryIntegerSet []int32

	QueryLong *int64

	QueryParamsMapOfStrings map[string]string

	QueryShort *int16

	QueryString *string

	QueryStringList []string

	QueryStringSet []string

	QueryTimestamp *time.Time

	QueryTimestampList []time.Time

	noSmithyDocumentSerde
}

type AllQueryStringTypesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAllQueryStringTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestxml_serializeOpAllQueryStringTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpAllQueryStringTypes{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "AllQueryStringTypes"); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAllQueryStringTypes(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opAllQueryStringTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "AllQueryStringTypes",
	}
}
