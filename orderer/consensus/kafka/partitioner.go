
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:31</date>
//</624456101408018432>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package kafka

import "github.com/Shopify/sarama"

type staticPartitioner struct {
	partitionID int32
}

//NewStaticPartitioner返回的PartitionerConstructor
//返回始终选择指定分区的分区程序。
func newStaticPartitioner(partition int32) sarama.PartitionerConstructor {
	return func(topic string) sarama.Partitioner {
		return &staticPartitioner{partition}
	}
}

//分区接受消息和分区计数并选择一个分区。
func (prt *staticPartitioner) Partition(message *sarama.ProducerMessage, numPartitions int32) (int32, error) {
	return prt.partitionID, nil
}

//RequiresConsistency向分区用户指示
//key->partition的映射是否一致。
func (prt *staticPartitioner) RequiresConsistency() bool {
	return true
}

