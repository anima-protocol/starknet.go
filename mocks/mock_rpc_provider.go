// Code generated by MockGen. DO NOT EDIT.
// Source: provider.go
//
// Generated by this command:
//
//	mockgen -destination=../mocks/mock_rpc_provider.go -package=mocks -source=provider.go api
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	felt "github.com/anima-protocol/juno/core/felt"
	rpc "github.com/anima-protocol/starknet.go/rpc"
	gomock "github.com/golang/mock/gomock"
)

// MockRpcProvider is a mock of RpcProvider interface.
type MockRpcProvider struct {
	ctrl     *gomock.Controller
	recorder *MockRpcProviderMockRecorder
}

// MockRpcProviderMockRecorder is the mock recorder for MockRpcProvider.
type MockRpcProviderMockRecorder struct {
	mock *MockRpcProvider
}

// NewMockRpcProvider creates a new mock instance.
func NewMockRpcProvider(ctrl *gomock.Controller) *MockRpcProvider {
	mock := &MockRpcProvider{ctrl: ctrl}
	mock.recorder = &MockRpcProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRpcProvider) EXPECT() *MockRpcProviderMockRecorder {
	return m.recorder
}

// AddDeclareTransaction mocks base method.
func (m *MockRpcProvider) AddDeclareTransaction(ctx context.Context, declareTransaction rpc.BroadcastDeclareTxnType) (*rpc.AddDeclareTransactionResponse, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDeclareTransaction", ctx, declareTransaction)
	ret0, _ := ret[0].(*rpc.AddDeclareTransactionResponse)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// AddDeclareTransaction indicates an expected call of AddDeclareTransaction.
func (mr *MockRpcProviderMockRecorder) AddDeclareTransaction(ctx, declareTransaction any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDeclareTransaction", reflect.TypeOf((*MockRpcProvider)(nil).AddDeclareTransaction), ctx, declareTransaction)
}

// AddDeployAccountTransaction mocks base method.
func (m *MockRpcProvider) AddDeployAccountTransaction(ctx context.Context, deployAccountTransaction rpc.BroadcastAddDeployTxnType) (*rpc.AddDeployAccountTransactionResponse, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDeployAccountTransaction", ctx, deployAccountTransaction)
	ret0, _ := ret[0].(*rpc.AddDeployAccountTransactionResponse)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// AddDeployAccountTransaction indicates an expected call of AddDeployAccountTransaction.
func (mr *MockRpcProviderMockRecorder) AddDeployAccountTransaction(ctx, deployAccountTransaction any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDeployAccountTransaction", reflect.TypeOf((*MockRpcProvider)(nil).AddDeployAccountTransaction), ctx, deployAccountTransaction)
}

// AddInvokeTransaction mocks base method.
func (m *MockRpcProvider) AddInvokeTransaction(ctx context.Context, invokeTxn rpc.BroadcastInvokeTxnType) (*rpc.AddInvokeTransactionResponse, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddInvokeTransaction", ctx, invokeTxn)
	ret0, _ := ret[0].(*rpc.AddInvokeTransactionResponse)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// AddInvokeTransaction indicates an expected call of AddInvokeTransaction.
func (mr *MockRpcProviderMockRecorder) AddInvokeTransaction(ctx, invokeTxn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddInvokeTransaction", reflect.TypeOf((*MockRpcProvider)(nil).AddInvokeTransaction), ctx, invokeTxn)
}

// BlockHashAndNumber mocks base method.
func (m *MockRpcProvider) BlockHashAndNumber(ctx context.Context) (*rpc.BlockHashAndNumberOutput, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockHashAndNumber", ctx)
	ret0, _ := ret[0].(*rpc.BlockHashAndNumberOutput)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// BlockHashAndNumber indicates an expected call of BlockHashAndNumber.
func (mr *MockRpcProviderMockRecorder) BlockHashAndNumber(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockHashAndNumber", reflect.TypeOf((*MockRpcProvider)(nil).BlockHashAndNumber), ctx)
}

// BlockNumber mocks base method.
func (m *MockRpcProvider) BlockNumber(ctx context.Context) (uint64, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockNumber", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// BlockNumber indicates an expected call of BlockNumber.
func (mr *MockRpcProviderMockRecorder) BlockNumber(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockNumber", reflect.TypeOf((*MockRpcProvider)(nil).BlockNumber), ctx)
}

// BlockTransactionCount mocks base method.
func (m *MockRpcProvider) BlockTransactionCount(ctx context.Context, blockID rpc.BlockID) (uint64, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockTransactionCount", ctx, blockID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// BlockTransactionCount indicates an expected call of BlockTransactionCount.
func (mr *MockRpcProviderMockRecorder) BlockTransactionCount(ctx, blockID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockTransactionCount", reflect.TypeOf((*MockRpcProvider)(nil).BlockTransactionCount), ctx, blockID)
}

// BlockWithReceipts mocks base method.
func (m *MockRpcProvider) BlockWithReceipts(ctx context.Context, blockID rpc.BlockID) (any, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockWithReceipts", ctx, blockID)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// BlockWithReceipts indicates an expected call of BlockWithReceipts.
func (mr *MockRpcProviderMockRecorder) BlockWithReceipts(ctx, blockID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockWithReceipts", reflect.TypeOf((*MockRpcProvider)(nil).BlockWithReceipts), ctx, blockID)
}

// BlockWithTxHashes mocks base method.
func (m *MockRpcProvider) BlockWithTxHashes(ctx context.Context, blockID rpc.BlockID) (any, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockWithTxHashes", ctx, blockID)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// BlockWithTxHashes indicates an expected call of BlockWithTxHashes.
func (mr *MockRpcProviderMockRecorder) BlockWithTxHashes(ctx, blockID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockWithTxHashes", reflect.TypeOf((*MockRpcProvider)(nil).BlockWithTxHashes), ctx, blockID)
}

// BlockWithTxs mocks base method.
func (m *MockRpcProvider) BlockWithTxs(ctx context.Context, blockID rpc.BlockID) (any, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockWithTxs", ctx, blockID)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// BlockWithTxs indicates an expected call of BlockWithTxs.
func (mr *MockRpcProviderMockRecorder) BlockWithTxs(ctx, blockID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockWithTxs", reflect.TypeOf((*MockRpcProvider)(nil).BlockWithTxs), ctx, blockID)
}

// Call mocks base method.
func (m *MockRpcProvider) Call(ctx context.Context, call rpc.FunctionCall, block rpc.BlockID) ([]*felt.Felt, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call", ctx, call, block)
	ret0, _ := ret[0].([]*felt.Felt)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// Call indicates an expected call of Call.
func (mr *MockRpcProviderMockRecorder) Call(ctx, call, block any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockRpcProvider)(nil).Call), ctx, call, block)
}

// ChainID mocks base method.
func (m *MockRpcProvider) ChainID(ctx context.Context) (string, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainID", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// ChainID indicates an expected call of ChainID.
func (mr *MockRpcProviderMockRecorder) ChainID(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainID", reflect.TypeOf((*MockRpcProvider)(nil).ChainID), ctx)
}

// Class mocks base method.
func (m *MockRpcProvider) Class(ctx context.Context, blockID rpc.BlockID, classHash *felt.Felt) (rpc.ClassOutput, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Class", ctx, blockID, classHash)
	ret0, _ := ret[0].(rpc.ClassOutput)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// Class indicates an expected call of Class.
func (mr *MockRpcProviderMockRecorder) Class(ctx, blockID, classHash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Class", reflect.TypeOf((*MockRpcProvider)(nil).Class), ctx, blockID, classHash)
}

// ClassAt mocks base method.
func (m *MockRpcProvider) ClassAt(ctx context.Context, blockID rpc.BlockID, contractAddress *felt.Felt) (rpc.ClassOutput, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClassAt", ctx, blockID, contractAddress)
	ret0, _ := ret[0].(rpc.ClassOutput)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// ClassAt indicates an expected call of ClassAt.
func (mr *MockRpcProviderMockRecorder) ClassAt(ctx, blockID, contractAddress any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClassAt", reflect.TypeOf((*MockRpcProvider)(nil).ClassAt), ctx, blockID, contractAddress)
}

// ClassHashAt mocks base method.
func (m *MockRpcProvider) ClassHashAt(ctx context.Context, blockID rpc.BlockID, contractAddress *felt.Felt) (*felt.Felt, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClassHashAt", ctx, blockID, contractAddress)
	ret0, _ := ret[0].(*felt.Felt)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// ClassHashAt indicates an expected call of ClassHashAt.
func (mr *MockRpcProviderMockRecorder) ClassHashAt(ctx, blockID, contractAddress any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClassHashAt", reflect.TypeOf((*MockRpcProvider)(nil).ClassHashAt), ctx, blockID, contractAddress)
}

// EstimateFee mocks base method.
func (m *MockRpcProvider) EstimateFee(ctx context.Context, requests []rpc.BroadcastTxn, simulationFlags []rpc.SimulationFlag, blockID rpc.BlockID) ([]rpc.FeeEstimate, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EstimateFee", ctx, requests, simulationFlags, blockID)
	ret0, _ := ret[0].([]rpc.FeeEstimate)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// EstimateFee indicates an expected call of EstimateFee.
func (mr *MockRpcProviderMockRecorder) EstimateFee(ctx, requests, simulationFlags, blockID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateFee", reflect.TypeOf((*MockRpcProvider)(nil).EstimateFee), ctx, requests, simulationFlags, blockID)
}

// EstimateMessageFee mocks base method.
func (m *MockRpcProvider) EstimateMessageFee(ctx context.Context, msg rpc.MsgFromL1, blockID rpc.BlockID) (*rpc.FeeEstimate, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EstimateMessageFee", ctx, msg, blockID)
	ret0, _ := ret[0].(*rpc.FeeEstimate)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// EstimateMessageFee indicates an expected call of EstimateMessageFee.
func (mr *MockRpcProviderMockRecorder) EstimateMessageFee(ctx, msg, blockID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateMessageFee", reflect.TypeOf((*MockRpcProvider)(nil).EstimateMessageFee), ctx, msg, blockID)
}

// Events mocks base method.
func (m *MockRpcProvider) Events(ctx context.Context, input rpc.EventsInput) (*rpc.EventChunk, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Events", ctx, input)
	ret0, _ := ret[0].(*rpc.EventChunk)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// Events indicates an expected call of Events.
func (mr *MockRpcProviderMockRecorder) Events(ctx, input any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Events", reflect.TypeOf((*MockRpcProvider)(nil).Events), ctx, input)
}

// GetTransactionStatus mocks base method.
func (m *MockRpcProvider) GetTransactionStatus(ctx context.Context, transactionHash *felt.Felt) (*rpc.TxnStatusResp, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionStatus", ctx, transactionHash)
	ret0, _ := ret[0].(*rpc.TxnStatusResp)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// GetTransactionStatus indicates an expected call of GetTransactionStatus.
func (mr *MockRpcProviderMockRecorder) GetTransactionStatus(ctx, transactionHash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionStatus", reflect.TypeOf((*MockRpcProvider)(nil).GetTransactionStatus), ctx, transactionHash)
}

// Nonce mocks base method.
func (m *MockRpcProvider) Nonce(ctx context.Context, blockID rpc.BlockID, contractAddress *felt.Felt) (*felt.Felt, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Nonce", ctx, blockID, contractAddress)
	ret0, _ := ret[0].(*felt.Felt)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// Nonce indicates an expected call of Nonce.
func (mr *MockRpcProviderMockRecorder) Nonce(ctx, blockID, contractAddress any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nonce", reflect.TypeOf((*MockRpcProvider)(nil).Nonce), ctx, blockID, contractAddress)
}

// SimulateTransactions mocks base method.
func (m *MockRpcProvider) SimulateTransactions(ctx context.Context, blockID rpc.BlockID, txns []rpc.Transaction, simulationFlags []rpc.SimulationFlag) ([]rpc.SimulatedTransaction, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SimulateTransactions", ctx, blockID, txns, simulationFlags)
	ret0, _ := ret[0].([]rpc.SimulatedTransaction)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// SimulateTransactions indicates an expected call of SimulateTransactions.
func (mr *MockRpcProviderMockRecorder) SimulateTransactions(ctx, blockID, txns, simulationFlags any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SimulateTransactions", reflect.TypeOf((*MockRpcProvider)(nil).SimulateTransactions), ctx, blockID, txns, simulationFlags)
}

// SpecVersion mocks base method.
func (m *MockRpcProvider) SpecVersion(ctx context.Context) (string, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpecVersion", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// SpecVersion indicates an expected call of SpecVersion.
func (mr *MockRpcProviderMockRecorder) SpecVersion(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpecVersion", reflect.TypeOf((*MockRpcProvider)(nil).SpecVersion), ctx)
}

// StateUpdate mocks base method.
func (m *MockRpcProvider) StateUpdate(ctx context.Context, blockID rpc.BlockID) (*rpc.StateUpdateOutput, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateUpdate", ctx, blockID)
	ret0, _ := ret[0].(*rpc.StateUpdateOutput)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// StateUpdate indicates an expected call of StateUpdate.
func (mr *MockRpcProviderMockRecorder) StateUpdate(ctx, blockID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateUpdate", reflect.TypeOf((*MockRpcProvider)(nil).StateUpdate), ctx, blockID)
}

// StorageAt mocks base method.
func (m *MockRpcProvider) StorageAt(ctx context.Context, contractAddress *felt.Felt, key string, blockID rpc.BlockID) (string, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageAt", ctx, contractAddress, key, blockID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// StorageAt indicates an expected call of StorageAt.
func (mr *MockRpcProviderMockRecorder) StorageAt(ctx, contractAddress, key, blockID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageAt", reflect.TypeOf((*MockRpcProvider)(nil).StorageAt), ctx, contractAddress, key, blockID)
}

// Syncing mocks base method.
func (m *MockRpcProvider) Syncing(ctx context.Context) (*rpc.SyncStatus, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Syncing", ctx)
	ret0, _ := ret[0].(*rpc.SyncStatus)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// Syncing indicates an expected call of Syncing.
func (mr *MockRpcProviderMockRecorder) Syncing(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Syncing", reflect.TypeOf((*MockRpcProvider)(nil).Syncing), ctx)
}

// TraceBlockTransactions mocks base method.
func (m *MockRpcProvider) TraceBlockTransactions(ctx context.Context, blockID rpc.BlockID) ([]rpc.Trace, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TraceBlockTransactions", ctx, blockID)
	ret0, _ := ret[0].([]rpc.Trace)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// TraceBlockTransactions indicates an expected call of TraceBlockTransactions.
func (mr *MockRpcProviderMockRecorder) TraceBlockTransactions(ctx, blockID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TraceBlockTransactions", reflect.TypeOf((*MockRpcProvider)(nil).TraceBlockTransactions), ctx, blockID)
}

// TraceTransaction mocks base method.
func (m *MockRpcProvider) TraceTransaction(ctx context.Context, transactionHash *felt.Felt) (rpc.TxnTrace, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TraceTransaction", ctx, transactionHash)
	ret0, _ := ret[0].(rpc.TxnTrace)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// TraceTransaction indicates an expected call of TraceTransaction.
func (mr *MockRpcProviderMockRecorder) TraceTransaction(ctx, transactionHash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TraceTransaction", reflect.TypeOf((*MockRpcProvider)(nil).TraceTransaction), ctx, transactionHash)
}

// TransactionByBlockIdAndIndex mocks base method.
func (m *MockRpcProvider) TransactionByBlockIdAndIndex(ctx context.Context, blockID rpc.BlockID, index uint64) (rpc.Transaction, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionByBlockIdAndIndex", ctx, blockID, index)
	ret0, _ := ret[0].(rpc.Transaction)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// TransactionByBlockIdAndIndex indicates an expected call of TransactionByBlockIdAndIndex.
func (mr *MockRpcProviderMockRecorder) TransactionByBlockIdAndIndex(ctx, blockID, index any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionByBlockIdAndIndex", reflect.TypeOf((*MockRpcProvider)(nil).TransactionByBlockIdAndIndex), ctx, blockID, index)
}

// TransactionByHash mocks base method.
func (m *MockRpcProvider) TransactionByHash(ctx context.Context, hash *felt.Felt) (rpc.Transaction, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionByHash", ctx, hash)
	ret0, _ := ret[0].(rpc.Transaction)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// TransactionByHash indicates an expected call of TransactionByHash.
func (mr *MockRpcProviderMockRecorder) TransactionByHash(ctx, hash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionByHash", reflect.TypeOf((*MockRpcProvider)(nil).TransactionByHash), ctx, hash)
}

// TransactionReceipt mocks base method.
func (m *MockRpcProvider) TransactionReceipt(ctx context.Context, transactionHash *felt.Felt) (*rpc.TransactionReceiptWithBlockInfo, *rpc.RPCError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionReceipt", ctx, transactionHash)
	ret0, _ := ret[0].(*rpc.TransactionReceiptWithBlockInfo)
	ret1, _ := ret[1].(*rpc.RPCError)
	return ret0, ret1
}

// TransactionReceipt indicates an expected call of TransactionReceipt.
func (mr *MockRpcProviderMockRecorder) TransactionReceipt(ctx, transactionHash any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionReceipt", reflect.TypeOf((*MockRpcProvider)(nil).TransactionReceipt), ctx, transactionHash)
}
