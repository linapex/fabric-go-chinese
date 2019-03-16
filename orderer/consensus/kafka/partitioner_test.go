
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:31</date>
//</624456101491904512>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package kafka

import (
	"testing"

	"github.com/Shopify/sarama"
	"github.com/stretchr/testify/assert"
)

func TestStaticPartitioner(t *testing.T) {
	var partition int32 = 3
	var numberOfPartitions int32 = 6

	partitionerConstructor := newStaticPartitioner(partition)
	partitioner := partitionerConstructor(channelNameForTest(t))

	for i := 0; i < 10; i++ {
		assignedPartition, err := partitioner.Partition(new(sarama.ProducerMessage), numberOfPartitions)
		assert.NoError(t, err, "Partitioner not functioning as expected:", err)
		assert.Equal(t, partition, assignedPartition, "Partitioner not returning the expected partition - expected %d, got %v", partition, assignedPartition)
	}
}

