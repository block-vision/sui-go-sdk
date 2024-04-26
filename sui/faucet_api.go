package sui

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/block-vision/sui-go-sdk/constant"
	"github.com/block-vision/sui-go-sdk/models"
	"io/ioutil"
	"net/http"
)

const (
	faucetUriGasV0 = "/gas"
	faucetUriGasV1 = "/v1/gas"
)

// RequestSuiFromFaucet requests sui from faucet.
func RequestSuiFromFaucet(faucetHost, recipientAddress string, header map[string]string) error {

	body := models.FaucetRequest{
		FixedAmountRequest: &models.FaucetFixedAmountRequest{
			Recipient: recipientAddress,
		},
	}

	err := faucetRequest(faucetHost+faucetUriGasV1, body, header)

	return err
}

// GetFaucetHost returns the faucet host for the given network.
func GetFaucetHost(network string) (string, error) {
	switch network {
	case constant.SuiTestnet:
		return constant.FaucetTestnetEndpoint, nil
	case constant.SuiDevnet:
		return constant.FaucetDevnetEndpoint, nil
	case constant.SuiLocalnet:
		return constant.FaucetLocalnetEndpoint, nil
	default:
		return "", errors.New(fmt.Sprintf("Unknown network: %s", network))
	}
}

func faucetRequest(faucetUrl string, body interface{}, headers map[string]string) error {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return errors.New(fmt.Sprintf("Marshal request body error: %s", err.Error()))
	}

	req, err := http.NewRequest(http.MethodPost, faucetUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		return errors.New(fmt.Sprintf("Create request error: %s", err.Error()))
	}

	req.Header.Set("Content-Type", "application/json")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return errors.New(fmt.Sprintf("Request faucet error: %s", err.Error()))
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("Read response body error: %s", err.Error()))
	}

	if resp.StatusCode < 200 || 300 <= resp.StatusCode {
		return errors.New(fmt.Sprintf("Request faucet failed, statusCode: %d, err: %+v", resp.StatusCode, string(bodyBytes)))
	}

	return nil
}
