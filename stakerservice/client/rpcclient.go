package client

import (
	"context"

	service "github.com/babylonchain/btc-staker/stakerservice"
	jsonrpcclient "github.com/cometbft/cometbft/rpc/jsonrpc/client"
)

type StakerServiceJsonRpcClient struct {
	client *jsonrpcclient.Client
}

// TODO Add some kind of timeout config
func NewStakerServiceJsonRpcClient(remoteAddress string) (*StakerServiceJsonRpcClient, error) {
	client, err := jsonrpcclient.New(remoteAddress)
	if err != nil {
		return nil, err
	}

	return &StakerServiceJsonRpcClient{
		client: client,
	}, nil
}

func (c *StakerServiceJsonRpcClient) Health(ctx context.Context) (*service.ResultHealth, error) {
	result := new(service.ResultHealth)
	_, err := c.client.Call(ctx, "health", map[string]interface{}{}, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *StakerServiceJsonRpcClient) ListOutputs(ctx context.Context) (*service.OutputsResponse, error) {
	result := new(service.OutputsResponse)
	_, err := c.client.Call(ctx, "list_outputs", map[string]interface{}{}, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *StakerServiceJsonRpcClient) BabylonValidators(ctx context.Context, offset *int, limit *int) (*service.ValidatorsResponse, error) {
	result := new(service.ValidatorsResponse)

	params := make(map[string]interface{})

	if limit != nil {
		params["limit"] = limit
	}

	if offset != nil {
		params["offset"] = offset
	}

	_, err := c.client.Call(ctx, "babylon_validators", params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *StakerServiceJsonRpcClient) Stake(ctx context.Context,
	stakerAddress string,
	stakingAmount int64,
	validatorPk string,
	stakingTimeBlocks int64,
) (*service.ResultStake, error) {
	result := new(service.ResultStake)

	params := make(map[string]interface{})
	params["stakerAddress"] = stakerAddress
	params["stakingAmount"] = stakingAmount
	params["validatorPk"] = validatorPk
	params["stakingTimeBlocks"] = stakingTimeBlocks

	_, err := c.client.Call(ctx, "stake", params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *StakerServiceJsonRpcClient) ListStakingTransactions(ctx context.Context, offset *int, limit *int) (*service.ListStakingTransactionsResponse, error) {
	result := new(service.ListStakingTransactionsResponse)

	params := make(map[string]interface{})

	if limit != nil {
		params["limit"] = limit
	}

	if offset != nil {
		params["offset"] = offset
	}

	_, err := c.client.Call(ctx, "list_staking_transactions", params, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
