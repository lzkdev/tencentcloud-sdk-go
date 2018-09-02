// Copyright 1999-2018 Tencent Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20180720

import (
	"fmt"

	"github.com/lzkdev/tencentcloud-sdk-go/tencentcloud/common"
	tchttp "github.com/lzkdev/tencentcloud-sdk-go/tencentcloud/common/http"
	"github.com/lzkdev/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-07-20"

type Client struct {
	common.Client
}

func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
	client = &Client{}
	client.Init(region).WithSecretId(secretId, secretKey)
	return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
	client = &Client{}
	client.Init(region).
		WithSecretId(credential.SecretId, credential.SecretKey).
		WithProfile(clientProfile)
	return
}

func NewReceiveMessageRequest() (request *ReceiveMessageRequest) {
	request = &ReceiveMessageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cmq", APIVersion, "ReceiveMessage")
	return
}

func NewReceiveMessageResponse() (response *ReceiveMessageResponse) {
	response = &ReceiveMessageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 消费消息
func (c *Client) ReceiveMessage(request *ReceiveMessageRequest) (response *ReceiveMessageResponse, err error) {
	if request == nil {
		request = NewReceiveMessageRequest()
	}
	response = NewReceiveMessageResponse()
	err = c.Send(request, response)
	fmt.Print(response)
	return
}

func NewDeleteMessageRequest() (request *DeleteMessageRequest) {
	request = &DeleteMessageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cmq", APIVersion, "DeleteMessage")
	return
}

func NewDeleteMessageResponse() (response *DeleteMessageResponse) {
	response = &DeleteMessageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 删除消息
func (c *Client) DeleteMessage(request *DeleteMessageRequest) (response *DeleteMessageResponse, err error) {
	if request == nil {
		request = NewDeleteMessageRequest()
	}
	response = NewDeleteMessageResponse()
	err = c.Send(request, response)
	fmt.Print(response)
	return
}

func NewSendMessageRequest() (request *SendMessageRequest) {
	request = &SendMessageRequest{
		BaseRequest: &tchttp.BaseRequest{},
	}
	request.Init().WithApiInfo("cmq", APIVersion, "ReceiveMessage")
	return
}

func NewSendMessageResponse() (response *SendMessageResponse) {
	response = &SendMessageResponse{
		BaseResponse: &tchttp.BaseResponse{},
	}
	return
}

// 发送消息
func (c *Client) SendMessage(request *SendMessageRequest) (response *SendMessageResponse, err error) {
	if request == nil {
		request = NewSendMessageRequest()
	}
	response = NewSendMessageResponse()
	err = c.Send(request, response)
	fmt.Print(response)
	return
}
