package adb

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

// DBClusterInDescribeDBClusters is a nested struct in adb response
type DBClusterInDescribeDBClusters struct {
	DBClusterId          string                   `json:"DBClusterId" xml:"DBClusterId"`
	DBClusterType        string                   `json:"DBClusterType" xml:"DBClusterType"`
	DBClusterDescription string                   `json:"DBClusterDescription" xml:"DBClusterDescription"`
	PayType              string                   `json:"PayType" xml:"PayType"`
	RegionId             string                   `json:"RegionId" xml:"RegionId"`
	ExpireTime           string                   `json:"ExpireTime" xml:"ExpireTime"`
	Expired              string                   `json:"Expired" xml:"Expired"`
	DBClusterStatus      string                   `json:"DBClusterStatus" xml:"DBClusterStatus"`
	DBVersion            string                   `json:"DBVersion" xml:"DBVersion"`
	LockMode             string                   `json:"LockMode" xml:"LockMode"`
	LockReason           string                   `json:"LockReason" xml:"LockReason"`
	CreateTime           string                   `json:"CreateTime" xml:"CreateTime"`
	DBNodeStorage        int64                    `json:"DBNodeStorage" xml:"DBNodeStorage"`
	DBNodeClass          string                   `json:"DBNodeClass" xml:"DBNodeClass"`
	DBNodeCount          int64                    `json:"DBNodeCount" xml:"DBNodeCount"`
	CommodityCode        string                   `json:"CommodityCode" xml:"CommodityCode"`
	Category             string                   `json:"Category" xml:"Category"`
	RdsInstanceId        string                   `json:"RdsInstanceId" xml:"RdsInstanceId"`
	DtsJobId             string                   `json:"DtsJobId" xml:"DtsJobId"`
	ExecutorCount        string                   `json:"ExecutorCount" xml:"ExecutorCount"`
	DiskType             string                   `json:"DiskType" xml:"DiskType"`
	VPCCloudInstanceId   string                   `json:"VPCCloudInstanceId" xml:"VPCCloudInstanceId"`
	Engine               string                   `json:"Engine" xml:"Engine"`
	DBClusterNetworkType string                   `json:"DBClusterNetworkType" xml:"DBClusterNetworkType"`
	VPCId                string                   `json:"VPCId" xml:"VPCId"`
	VSwitchId            string                   `json:"VSwitchId" xml:"VSwitchId"`
	ZoneId               string                   `json:"ZoneId" xml:"ZoneId"`
	ConnectionString     string                   `json:"ConnectionString" xml:"ConnectionString"`
	Port                 string                   `json:"Port" xml:"Port"`
	ComputeResource      string                   `json:"ComputeResource" xml:"ComputeResource"`
	StorageResource      string                   `json:"StorageResource" xml:"StorageResource"`
	Mode                 string                   `json:"Mode" xml:"Mode"`
	ResourceGroupId      string                   `json:"ResourceGroupId" xml:"ResourceGroupId"`
	Tags                 TagsInDescribeDBClusters `json:"Tags" xml:"Tags"`
}
