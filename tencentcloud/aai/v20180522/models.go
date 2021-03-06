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

package v20180522

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type SentenceRecognitionRequest struct {
	*tchttp.BaseRequest
	// 腾讯云项目 ID，可填 0，总长度不超过 1024 字节。
	ProjectId *uint64 `json:"ProjectId" name:"ProjectId"`
	// 子服务类型。0：离线语音识别。1：实时流式识别，2，一句话识别。
	SubServiceType *uint64 `json:"SubServiceType" name:"SubServiceType"`
	// 引擎类型。8k：电话 8k 通用模型；16k：16k 通用模型。
	EngSerViceType *string `json:"EngSerViceType" name:"EngSerViceType"`
	// 语音数据来源。0：语音 URL；1：语音数据（post body）。
	SourceType *uint64 `json:"SourceType" name:"SourceType"`
	// 识别音频的音频格式（支持mp3,wav）。
	VoiceFormat *string `json:"VoiceFormat" name:"VoiceFormat"`
	// 用户端对此任务的唯一标识，用户自助生成，用于用户查找识别结果。
	UsrAudioKey *string `json:"UsrAudioKey" name:"UsrAudioKey"`
	// 语音 URL，公网可下载。当 SourceType 值为 0 时须填写该字段，为 1 时不填；URL 的长度大于 0，小于 2048，需进行urlencode编码。音频时间长度要小于60s。
	Url *string `json:"Url" name:"Url"`
	// 语音数据，当SourceType 值为1时必须填写，为0可不写。要base64编码。音频数据要小于900k。
	Data *string `json:"Data" name:"Data"`
	// 数据长度，当 SourceType 值为1时必须填写，为0可不写。
	DataLen *int64 `json:"DataLen" name:"DataLen"`
}

func (r *SentenceRecognitionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SentenceRecognitionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SentenceRecognitionResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// 识别结果。
		Result *string `json:"Result" name:"Result"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *SentenceRecognitionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SentenceRecognitionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TextToVoiceRequest struct {
	*tchttp.BaseRequest
	// 合成语音的源文本
	Text *string `json:"Text" name:"Text"`
	// 一次请求对应一个SessionId，会原样返回，建议传入类似于uuid的字符串防止重复
	SessionId *string `json:"SessionId" name:"SessionId"`
	// 模型类型，1-默认模型
	ModelType *int64 `json:"ModelType" name:"ModelType"`
	// 音量大小，暂仅支持默认值1
	Volume *float64 `json:"Volume" name:"Volume"`
	// 语速，暂仅支持默认值1
	Speed *float64 `json:"Speed" name:"Speed"`
	// 用户自定义项目id，默认为0
	ProjectId *int64 `json:"ProjectId" name:"ProjectId"`
	// 音色，1-默认音色
	VoiceType *int64 `json:"VoiceType" name:"VoiceType"`
	// 主语言类型<li>1-中文(包括粤语)，最大300字符</li><li>2-英文，最大支持600字符</li>
	PrimaryLanguage *uint64 `json:"PrimaryLanguage" name:"PrimaryLanguage"`
	// 音频采样率：暂仅支持16k
	SampleRate *uint64 `json:"SampleRate" name:"SampleRate"`
}

func (r *TextToVoiceRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TextToVoiceRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TextToVoiceResponse struct {
	*tchttp.BaseResponse
	Response *struct {
		// base编码的wav音频
		Audio *string `json:"Audio" name:"Audio"`
		// 一次请求对应一个SessionId
		SessionId *string `json:"SessionId" name:"SessionId"`
		// 唯一请求ID，每次请求都会返回。定位问题时需要提供该次请求的RequestId。
		RequestId *string `json:"RequestId" name:"RequestId"`
	} `json:"Response"`
}

func (r *TextToVoiceResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *TextToVoiceResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
