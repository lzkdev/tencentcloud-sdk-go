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
	"encoding/json"

	tchttp "github.com/lzkdev/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ReceiveMessageRequest struct {
	*tchttp.BaseRequest
	// 队列名字
	QueueName *string `json:"queueName" name:"queueName"`
	// RequestResponse(同步) 和 Event(异步)，默认为同步。
	PollingWaitSeconds *int `json:"pollingWaitSeconds" name:"pollingWaitSeconds"`
}

func (r *ReceiveMessageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReceiveMessageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type ReceiveMessageResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 0：表示成功，others：错误
		Code *int `json:"code" name:"code"`
		// 错误提示信息
		Message *string `json:"message" name:"message"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"requestId" name:"requestId"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		ClientRequestId *int64 `json:"clientRequestId" name:"clientRequestId" omitempty`
		// 本次消费的消息唯一标识 Id
		MsgId *string `json:"msgId" name:"msgId" omitempty`
		// 每次消费返回唯一的消息句柄。用于删除该消息，仅上一次消费时产生的消息句柄能用于删除消息。
		ReceiptHandle *string `json:"receiptHandle" name:"receiptHandle" omitempty`
		// 本次消费的消息正文
		MsgBody *string `json:"msgBody" name:"msgBody"`
		// 消费被生产出来，进入队列的时间。返回 Unix 时间戳，精确到秒
		EnqueueTime *uint64 `json:"enqueueTime" name:"enqueueTime"`
		// 消息的下次可见（可再次被消费）时间。返回 Unix 时间戳，精确到秒。
		NextVisibleTime *uint64 `json:"nextVisibleTime" name:"nextVisibleTime"`
		// 第一次消费该消息的时间。返回 Unix 时间戳，精确到秒。
		FirstDequeueTime *uint64 `json:"firstDequeueTime" name:"firstDequeueTime"`
		// 消息被消费的次数。
		DequeueCount *int `json:"dequeueCount" name:"dequeueCount"`
	} `json:"Response"`
}

func (r *ReceiveMessageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *ReceiveMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMessageRequest struct {
	*tchttp.BaseRequest
	// 队列名字
	QueueName *string `json:"queueName" name:"queueName"`
	// 上次消费返回唯一的消息句柄，用于删除消息。
	ReceiptHandle *string `json:"receiptHandle" name:"receiptHandle"`
}

func (r *DeleteMessageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMessageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type DeleteMessageResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 0：表示成功，others：错误
		Code *int `json:"code" name:"code"`
		// 错误提示信息
		Message *string `json:"message" name:"message"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"requestId" name:"requestId"`
	} `json:"Response"`
}

func (r *DeleteMessageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *DeleteMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendMessageRequest struct {
	*tchttp.BaseRequest
	// 队列名字
	QueueName *string `json:"queueName" name:"queueName"`
	// 消息正文。至少 1 Byte，最大长度受限于设置的队列消息最大长度属性
	MsgBody *string `json:"msgBody" name:"msgBody"`
	// 单位为秒，表示该消息发送到队列后，需要延时多久用户才可见该消息。
	DelaySeconds int `json:"delaySeconds" name:"delaySeconds"`
}

func (r *SendMessageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendMessageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type SendMessageResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 0：表示成功，others：错误
		Code *int `json:"code" name:"code"`
		// 错误提示信息
		Message *string `json:"message" name:"message"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"requestId" name:"requestId"`
		// 本次消费的消息唯一标识 Id
		MsgId *string `json:"msgId" name:"msgId" omitempty`
	} `json:"Response"`
}

func (r *SendMessageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *SendMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublishMessageRequest struct {
	*tchttp.BaseRequest
	// 主题名字
	TopicName *string `json:"topicName" name:"topicName"`
	// 消息正文。至少 1 Byte，最大长度受限于设置的队列消息最大长度属性
	MsgBody *string `json:"msgBody" name:"msgBody"`
	// 消息过滤标签
	MsgTag *string `json:"msgTag.n" name:"msgTag.n"`
	// 长度<=64字节，该字段用于表示发送消息的路由路径
	RoutingKey *string `json:"routingKey" name:"routingKey"`
}

func (r *PublishMessageRequest) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishMessageRequest) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type PublishMessageResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 0：表示成功，others：错误
		Code *int `json:"code" name:"code"`
		// 错误提示信息
		Message *string `json:"message" name:"message"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"requestId" name:"requestId"`
		// 本次消费的消息唯一标识 Id
		MsgId *string `json:"msgId" name:"msgId" omitempty`
	} `json:"Response"`
}

func (r *PublishMessageResponse) ToJsonString() string {
	b, _ := json.Marshal(r)
	return string(b)
}

func (r *PublishMessageResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}
