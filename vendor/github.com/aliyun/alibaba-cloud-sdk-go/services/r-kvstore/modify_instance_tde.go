package r_kvstore

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifyInstanceTDE invokes the r_kvstore.ModifyInstanceTDE API synchronously
func (client *Client) ModifyInstanceTDE(request *ModifyInstanceTDERequest) (response *ModifyInstanceTDEResponse, err error) {
	response = CreateModifyInstanceTDEResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyInstanceTDEWithChan invokes the r_kvstore.ModifyInstanceTDE API asynchronously
func (client *Client) ModifyInstanceTDEWithChan(request *ModifyInstanceTDERequest) (<-chan *ModifyInstanceTDEResponse, <-chan error) {
	responseChan := make(chan *ModifyInstanceTDEResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyInstanceTDE(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ModifyInstanceTDEWithCallback invokes the r_kvstore.ModifyInstanceTDE API asynchronously
func (client *Client) ModifyInstanceTDEWithCallback(request *ModifyInstanceTDERequest, callback func(response *ModifyInstanceTDEResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyInstanceTDEResponse
		var err error
		defer close(result)
		response, err = client.ModifyInstanceTDE(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ModifyInstanceTDERequest is the request struct for api ModifyInstanceTDE
type ModifyInstanceTDERequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	Product              string           `position:"Query" name:"Product"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EncryptionKey        string           `position:"Query" name:"EncryptionKey"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	EncryptionName       string           `position:"Query" name:"EncryptionName"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	RoleArn              string           `position:"Query" name:"RoleArn"`
	Category             string           `position:"Query" name:"Category"`
	TDEStatus            string           `position:"Query" name:"TDEStatus"`
}

// ModifyInstanceTDEResponse is the response struct for api ModifyInstanceTDE
type ModifyInstanceTDEResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyInstanceTDERequest creates a request to invoke ModifyInstanceTDE API
func CreateModifyInstanceTDERequest() (request *ModifyInstanceTDERequest) {
	request = &ModifyInstanceTDERequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("R-kvstore", "2015-01-01", "ModifyInstanceTDE", "redisa", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyInstanceTDEResponse creates a response to parse from ModifyInstanceTDE response
func CreateModifyInstanceTDEResponse() (response *ModifyInstanceTDEResponse) {
	response = &ModifyInstanceTDEResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}