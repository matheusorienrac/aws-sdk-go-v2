// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchainquery

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchainquery/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the token balance for a batch of tokens by using the BatchGetTokenBalance
// action for every token in the request. Only the native tokens BTC and ETH, and
// the ERC-20, ERC-721, and ERC 1155 token standards are supported.
func (c *Client) BatchGetTokenBalance(ctx context.Context, params *BatchGetTokenBalanceInput, optFns ...func(*Options)) (*BatchGetTokenBalanceOutput, error) {
	if params == nil {
		params = &BatchGetTokenBalanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetTokenBalance", params, optFns, c.addOperationBatchGetTokenBalanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetTokenBalanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetTokenBalanceInput struct {

	// An array of BatchGetTokenBalanceInputItem objects whose balance is being
	// requested.
	GetTokenBalanceInputs []types.BatchGetTokenBalanceInputItem

	noSmithyDocumentSerde
}

type BatchGetTokenBalanceOutput struct {

	// An array of BatchGetTokenBalanceErrorItem objects returned from the request.
	//
	// This member is required.
	Errors []types.BatchGetTokenBalanceErrorItem

	// An array of BatchGetTokenBalanceOutputItem objects returned by the response.
	//
	// This member is required.
	TokenBalances []types.BatchGetTokenBalanceOutputItem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetTokenBalanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchGetTokenBalance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchGetTokenBalance{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "BatchGetTokenBalance"); err != nil {
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
	if err = addOpBatchGetTokenBalanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetTokenBalance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetTokenBalance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "BatchGetTokenBalance",
	}
}
