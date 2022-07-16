/*
Copyright 2021 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package shared

import (
	"bytes"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"time"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type ChainConnection struct {
	clientConnection *grpc.ClientConn
	gateway          *client.Gateway
	network          *client.Network
	contract         *client.Contract
}

func CreateChainConnection() *ChainConnection {

	c := new(ChainConnection)
	var err error

	// The gRPC client connection should be shared by all Gateway connections to this endpoint
	c.clientConnection = newGrpcConnection()

	id := newIdentity()
	sign := newSign()

	// Create a Gateway connection for a specific client identity
	c.gateway, err = client.Connect(
		id,
		client.WithSign(sign),
		client.WithClientConnection(c.clientConnection),
		// Default timeouts for different gRPC calls
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		panic(err)
	}

	c.network = c.gateway.GetNetwork(Hyperconfig.ChannelName)
	c.contract = c.network.GetContract(Hyperconfig.ChaincodeName)

	return c
}

func (c *ChainConnection) CloseConnection() {
	c.clientConnection.Close()
	c.gateway.Close()
}

// newGrpcConnection creates a gRPC connection to the Gateway server.
func newGrpcConnection() *grpc.ClientConn {
	certificate, err := loadCertificate(Hyperconfig.TlsCertPath)
	if err != nil {
		panic(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(certificate)
	transportCredentials := credentials.NewClientTLSFromCert(certPool, Hyperconfig.GatewayPeer)

	connection, err := grpc.Dial(Hyperconfig.PeerEndpoint, grpc.WithTransportCredentials(transportCredentials))
	if err != nil {
		panic(fmt.Errorf("failed to create gRPC connection: %w", err))
	}

	return connection
}

// newIdentity creates a client identity for this Gateway connection using an X.509 certificate.
func newIdentity() *identity.X509Identity {
	certificate, err := loadCertificate(Hyperconfig.CertPath)
	if err != nil {
		panic(err)
	}

	id, err := identity.NewX509Identity(Hyperconfig.MspID, certificate)
	if err != nil {
		panic(err)
	}

	return id
}

func loadCertificate(filename string) (*x509.Certificate, error) {
	certificatePEM, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read certificate file: %w", err)
	}
	return identity.CertificateFromPEM(certificatePEM)
}

// newSign creates a function that generates a digital signature from a message digest using a private key.
func newSign() identity.Sign {
	files, err := ioutil.ReadDir(Hyperconfig.KeyPath)
	if err != nil {
		panic(fmt.Errorf("failed to read private key directory: %w", err))
	}
	privateKeyPEM, err := ioutil.ReadFile(path.Join(Hyperconfig.KeyPath, files[0].Name()))

	if err != nil {
		panic(fmt.Errorf("failed to read private key file: %w", err))
	}

	privateKey, err := identity.PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		panic(err)
	}

	sign, err := identity.NewPrivateKeySign(privateKey)
	if err != nil {
		panic(err)
	}

	return sign
}

// InitLedger runs the InitLedger function on the chaincode, which currently creates some dummy assets
func (c *ChainConnection) InitLedger() {
	fmt.Printf("Submit Transaction: InitLedger, function creates the initial set of assets on the ledger \n")

	_, err := c.contract.SubmitTransaction("InitLedger")
	if err != nil {
		log.Println(fmt.Errorf("failed to submit transaction: %w", err))
	}

	log.Println("*** Transaction committed successfully")
}

// GetAllAssets Returns all Assets on the ledger
func (c *ChainConnection) GetAllAssets() ([]Asset, error) {
	fmt.Println("Evaluate Transaction: GetAllAssets, function returns all the current assets on the ledger")

	evaluateResult, err := c.contract.EvaluateTransaction("GetAllAssets")
	if err != nil {
		log.Println(fmt.Errorf("failed to evaluate transaction: %w", err))
		return make([]Asset, 0), err
	}
	return toAssets(evaluateResult), nil
}

// CreateAsset Creates an asset synchronously
func (c *ChainConnection) CreateAsset(a *Asset) error {
	fmt.Printf("Submit Transaction: CreateAsset, creates new asset on the ledger \n")

	jsonData, err := json.Marshal(a.Data)
	jsonRefs, err := json.Marshal(a.Refs)
	_, err = c.contract.SubmitTransaction("CreateAsset", a.ID, a.AssetType.String(), string(jsonData), string(jsonRefs))
	if err != nil {
		log.Println(fmt.Errorf("failed to submit transaction: %w", err))
		return err
	}

	log.Println("*** Transaction committed successfully")
	return nil
}

// UpdateAsset Creates an asset synchronously
func (c *ChainConnection) UpdateAsset(a *Asset) error {
	fmt.Printf("Submit Transaction: UpdateAsset, updates an asset on the ledger \n")

	jsonData, err := json.Marshal(a.Data)
	jsonRefs, err := json.Marshal(a.Refs)
	_, err = c.contract.SubmitTransaction("UpdateAsset", a.ID, a.AssetType.String(), string(jsonData), string(jsonRefs))
	if err != nil {
		log.Println(fmt.Errorf("failed to submit transaction: %w", err))
		return err
	}

	log.Println("*** Transaction committed successfully")
	return nil
}

// ReadAssetByID queries a single Asset
func (c *ChainConnection) ReadAssetByID(id string) (Asset, error) {
	fmt.Printf("Evaluate Transaction: ReadAsset, function returns asset attributes\n")

	evaluateResult, err := c.contract.EvaluateTransaction("ReadAsset", id)
	if err != nil {
		log.Println(fmt.Errorf("failed to evaluate transaction: %w", err))
		return Asset{}, err
	}
	return toAsset(evaluateResult), nil
}

// GetMetadata returns all the metadata of the chaincode, e.g. available functions
func (c *ChainConnection) GetMetadata() (string, error) {
	fmt.Printf("Evaluate Transaction: GetMetadata\n")

	evaluateResult, err := c.contract.EvaluateTransaction("org.hyperledger.fabric:GetMetadata")
	if err != nil {
		log.Println(fmt.Errorf("failed to evaluate transaction: %w", err))
		return "", err
	}
	result := formatJSON(evaluateResult)
	return result, nil
}

// DeleteAsset deletes a single Asset synchronously
func (c *ChainConnection) DeleteAsset(id string) (string, error) {
	fmt.Printf("Evaluate Transaction: DeleteAsset, function deletes asset\n")

	evaluateResult, err := c.contract.SubmitTransaction("DeleteAsset", id)
	if err != nil {
		log.Println(fmt.Errorf("failed to evaluate transaction: %w", err))
		return "", err
	}
	result := formatJSON(evaluateResult)
	return result, nil
}

// Format JSON data
func formatJSON(data []byte) string {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, data, " ", "  "); err != nil {
		panic(fmt.Errorf("failed to parse JSON: %w", err))
	}
	return prettyJSON.String()
}

// convert bytes to asset
func toAsset(data []byte) Asset {
	var a Asset
	err := json.Unmarshal(data, &a)
	if err != nil {
		panic(fmt.Errorf("Error unmarshaling result: %w", err))
	}
	return a
}

// convert even more bytes to even more assets
func toAssets(data []byte) []Asset {
	a := make([]Asset, 0)
	err := json.Unmarshal(data, &a)
	if err != nil {
		panic(fmt.Errorf("Error unmarshaling result: %w", err))
	}
	return a
}
