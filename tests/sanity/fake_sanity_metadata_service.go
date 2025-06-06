// Copyright 2024 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the 'License');
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an 'AS IS' BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sanity

import (
	"github.com/aws/aws-sdk-go-v2/aws/arn"
)

type fakeMetadataService struct {
	instanceID       string
	region           string
	availabilityZone string
	outpostArn       arn.ARN
}

func newFakeMetadataService(id string, r string, az string, oa arn.ARN) *fakeMetadataService {
	return &fakeMetadataService{
		instanceID:       id,
		region:           r,
		availabilityZone: az,
		outpostArn:       oa,
	}
}

func (m *fakeMetadataService) UpdateMetadata() error {
	return nil
}

func (m *fakeMetadataService) GetInstanceID() string {
	return m.instanceID
}

func (m *fakeMetadataService) GetInstanceType() string {
	return ""
}

func (m *fakeMetadataService) GetRegion() string {
	return m.region
}

func (m *fakeMetadataService) GetAvailabilityZone() string {
	return m.availabilityZone
}

func (m *fakeMetadataService) GetNumAttachedENIs() int {
	return 0
}

func (m *fakeMetadataService) GetNumBlockDeviceMappings() int {
	return 0
}

func (m *fakeMetadataService) GetOutpostArn() arn.ARN {
	return m.outpostArn
}
