package cs

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

// ModifyClusterTags invokes the cs.ModifyClusterTags API synchronously
func (client *Client) ModifyClusterTags(request *ModifyClusterTagsRequest) (response *ModifyClusterTagsResponse, err error) {
	response = CreateModifyClusterTagsResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyClusterTagsWithChan invokes the cs.ModifyClusterTags API asynchronously
func (client *Client) ModifyClusterTagsWithChan(request *ModifyClusterTagsRequest) (<-chan *ModifyClusterTagsResponse, <-chan error) {
	responseChan := make(chan *ModifyClusterTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyClusterTags(request)
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

// ModifyClusterTagsWithCallback invokes the cs.ModifyClusterTags API asynchronously
func (client *Client) ModifyClusterTagsWithCallback(request *ModifyClusterTagsRequest, callback func(response *ModifyClusterTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyClusterTagsResponse
		var err error
		defer close(result)
		response, err = client.ModifyClusterTags(request)
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

// ModifyClusterTagsRequest is the request struct for api ModifyClusterTags
type ModifyClusterTagsRequest struct {
	*requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
}

// ModifyClusterTagsResponse is the response struct for api ModifyClusterTags
type ModifyClusterTagsResponse struct {
	*responses.BaseResponse
}

// CreateModifyClusterTagsRequest creates a request to invoke ModifyClusterTags API
func CreateModifyClusterTagsRequest() (request *ModifyClusterTagsRequest) {
	request = &ModifyClusterTagsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "ModifyClusterTags", "/clusters/[ClusterId]/tags", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyClusterTagsResponse creates a response to parse from ModifyClusterTags response
func CreateModifyClusterTagsResponse() (response *ModifyClusterTagsResponse) {
	response = &ModifyClusterTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
