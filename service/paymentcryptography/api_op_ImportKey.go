// Code generated by smithy-go-codegen DO NOT EDIT.

package paymentcryptography

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/paymentcryptography/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Imports keys and public key certificates into Amazon Web Services Payment
// Cryptography. Amazon Web Services Payment Cryptography simplifies main or root
// key exchange process by eliminating the need of a paper-based key exchange
// process. It takes a modern and secure approach based of the ANSI X9 TR-34 key
// exchange standard. You can use ImportKey to import main or root keys such as
// KEK (Key Encryption Key) using asymmetric key exchange technique following the
// ANSI X9 TR-34 standard. The ANSI X9 TR-34 standard uses asymmetric keys to
// establishes bi-directional trust between the two parties exchanging keys. After
// you have imported a main or root key, you can import working keys to perform
// various cryptographic operations within Amazon Web Services Payment Cryptography
// using the ANSI X9 TR-31 symmetric key exchange standard as mandated by PCI PIN.
// You can also import a root public key certificate, a self-signed certificate
// used to sign other public key certificates, or a trusted public key certificate
// under an already established root public key certificate. To import a public
// root key certificate Using this operation, you can import the public component
// (in PEM cerificate format) of your private root key. You can use the imported
// public root key certificate for digital signatures, for example signing wrapping
// key or signing key in TR-34, within your Amazon Web Services Payment
// Cryptography account. Set the following parameters:
//   - KeyMaterial : RootCertificatePublicKey
//   - KeyClass : PUBLIC_KEY
//   - KeyModesOfUse : Verify
//   - KeyUsage : TR31_S0_ASYMMETRIC_KEY_FOR_DIGITAL_SIGNATURE
//   - PublicKeyCertificate : The certificate authority used to sign the root
//     public key certificate.
//
// To import a trusted public key certificate The root public key certificate must
// be in place and operational before you import a trusted public key certificate.
// Set the following parameters:
//   - KeyMaterial : TrustedCertificatePublicKey
//   - CertificateAuthorityPublicKeyIdentifier : KeyArn of the
//     RootCertificatePublicKey .
//   - KeyModesOfUse and KeyUsage : Corresponding to the cryptographic operations
//     such as wrap, sign, or encrypt that you will allow the trusted public key
//     certificate to perform.
//   - PublicKeyCertificate : The certificate authority used to sign the trusted
//     public key certificate.
//
// Import main keys Amazon Web Services Payment Cryptography uses TR-34 asymmetric
// key exchange standard to import main keys such as KEK. In TR-34 terminology, the
// sending party of the key is called Key Distribution Host (KDH) and the receiving
// party of the key is called Key Receiving Host (KRH). During the key import
// process, KDH is the user who initiates the key import and KRH is Amazon Web
// Services Payment Cryptography who receives the key. Before initiating TR-34 key
// import, you must obtain an import token by calling GetParametersForImport . This
// operation also returns the wrapping key certificate that KDH uses wrap key under
// import to generate a TR-34 wrapped key block. The import token expires after 7
// days. Set the following parameters:
//   - CertificateAuthorityPublicKeyIdentifier : The KeyArn of the certificate
//     chain that will sign the signing key certificate and should exist within Amazon
//     Web Services Payment Cryptography before initiating TR-34 key import. If it does
//     not exist, you can import it by calling by calling ImportKey for
//     RootCertificatePublicKey .
//   - ImportToken : Obtained from KRH by calling GetParametersForImport .
//   - WrappedKeyBlock : The TR-34 wrapped key block from KDH. It contains the KDH
//     key under import, wrapped with KRH provided wrapping key certificate and signed
//     by the KDH private signing key. This TR-34 key block is generated by the KDH
//     Hardware Security Module (HSM) outside of Amazon Web Services Payment
//     Cryptography.
//   - SigningKeyCertificate : The public component of the private key that signed
//     the KDH TR-34 wrapped key block. In PEM certificate format.
//
// TR-34 is intended primarily to exchange 3DES keys. Your ability to export
// AES-128 and larger AES keys may be dependent on your source system. Import
// working keys Amazon Web Services Payment Cryptography uses TR-31 symmetric key
// exchange standard to import working keys. A KEK must be established within
// Amazon Web Services Payment Cryptography by using TR-34 key import. To initiate
// a TR-31 key import, set the following parameters:
//   - WrappedKeyBlock : The key under import and encrypted using KEK. The TR-31
//     key block generated by your HSM outside of Amazon Web Services Payment
//     Cryptography.
//   - WrappingKeyIdentifier : The KeyArn of the KEK that Amazon Web Services
//     Payment Cryptography uses to decrypt or unwrap the key under import.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts. Related operations:
//   - ExportKey
//   - GetParametersForImport
func (c *Client) ImportKey(ctx context.Context, params *ImportKeyInput, optFns ...func(*Options)) (*ImportKeyOutput, error) {
	if params == nil {
		params = &ImportKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ImportKey", params, optFns, c.addOperationImportKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ImportKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ImportKeyInput struct {

	// The key or public key certificate type to use during key material import, for
	// example TR-34 or RootCertificatePublicKey.
	//
	// This member is required.
	KeyMaterial types.ImportKeyMaterial

	// Specifies whether import key is enabled.
	Enabled *bool

	// The algorithm that Amazon Web Services Payment Cryptography uses to calculate
	// the key check value (KCV) for DES and AES keys. For DES key, the KCV is computed
	// by encrypting 8 bytes, each with value '00', with the key to be checked and
	// retaining the 3 highest order bytes of the encrypted result. For AES key, the
	// KCV is computed by encrypting 8 bytes, each with value '01', with the key to be
	// checked and retaining the 3 highest order bytes of the encrypted result.
	KeyCheckValueAlgorithm types.KeyCheckValueAlgorithm

	// The tags to attach to the key. Each tag consists of a tag key and a tag value.
	// Both the tag key and the tag value are required, but the tag value can be an
	// empty (null) string. You can't have more than one tag on an Amazon Web Services
	// Payment Cryptography key with the same tag key. You can't have more than one tag
	// on an Amazon Web Services Payment Cryptography key with the same tag key. If you
	// specify an existing tag key with a different tag value, Amazon Web Services
	// Payment Cryptography replaces the current tag value with the specified one. To
	// use this parameter, you must have TagResource permission. Don't include
	// confidential or sensitive information in this field. This field may be displayed
	// in plaintext in CloudTrail logs and other output. Tagging or untagging an Amazon
	// Web Services Payment Cryptography key can allow or deny permission to the key.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type ImportKeyOutput struct {

	// The KeyARN of the key material imported within Amazon Web Services Payment
	// Cryptography.
	//
	// This member is required.
	Key *types.Key

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationImportKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpImportKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpImportKey{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ImportKey"); err != nil {
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
	if err = addOpImportKeyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opImportKey(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opImportKey(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ImportKey",
	}
}
