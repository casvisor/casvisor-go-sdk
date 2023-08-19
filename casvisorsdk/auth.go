// Copyright 2023 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package casvisorsdk

type Config struct {
	Endpoint         string
	ClientId         string
	ClientSecret     string
	OrganizationName string
	ApplicationName  string
}

type Client struct {
	Config
}

var globalClient *Client

func InitConfig(endpoint string, clientId string, clientSecret string, organizationName string, applicationName string) {
	globalClient = NewClient(endpoint, clientId, clientSecret, organizationName, applicationName)
}

func NewClient(endpoint string, clientId string, clientSecret string, organizationName string, applicationName string) *Client {
	return NewClientWithConf(
		&Config{
			Endpoint:         endpoint,
			ClientId:         clientId,
			ClientSecret:     clientSecret,
			OrganizationName: organizationName,
			ApplicationName:  applicationName,
		})
}

func NewClientWithConf(config *Config) *Client {
	return &Client{
		*config,
	}
}

func GetClient() *Client {
	return globalClient
}
