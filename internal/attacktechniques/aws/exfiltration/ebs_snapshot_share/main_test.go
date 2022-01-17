package aws

import (
	"github.com/datadog/stratus-red-team/internal/providers/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEbsSnapshotShare(t *testing.T) {
	mockProvider := new(mocks.StratusProvider)
	mockAwsProvider := new(mocks.AwsProvider)
	mockAwsProvider.On("S3Client").Return()
	mockProvider.On("GetAwsProvider").Return(mockAwsProvider)
	err := ebsSnapshotShare.Detonate(map[string]string{"snapshot_id": "snap-1234"}, mockProvider)
	assert.Nil(t, err)
}
