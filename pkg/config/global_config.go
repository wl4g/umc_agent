/**
 * Copyright 2017 ~ 2025 the original author or authors[983708408@qq.com].
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"time"
	"umc-agent/pkg/constant"
)

// ---------------------
// Global properties
// ---------------------
type GlobalProperties struct {
	Logging    LoggingProperties    `yaml:"logging"`
	Launcher   LauncherProperties   `yaml:"launcher"`
	Indicators IndicatorsProperties `yaml:"indicators"`
}

// ---------------------
// Logging properties
// ---------------------
type LoggingProperties struct {
	FileName string           `yaml:"file"`
	Level    string           `yaml:"level"`
	Policy   PolicyProperties `yaml:"policy"`
}

// Logging archive policy
type PolicyProperties struct {
	RetentionDays int `yaml:"retention-days"`
	MaxBackups    int `yaml:"max-backups"`
	MaxSize       int `yaml:"max-size"`
}

// ---------------------
// Launcher properties.
// ---------------------
type LauncherProperties struct {
	Http  HttpLauncherProperties  `yaml:"http"`
	Kafka KafkaLauncherProperties `yaml:"kafka"`
}

type HttpLauncherProperties struct {
	ServerGateway string `yaml:"server-gateway"`
}

type KafkaLauncherProperties struct {
	Enabled          bool   `yaml:"enabled"`
	BootstrapServers string `yaml:"bootstrap.servers"`
	Topic            string `yaml:"topic"`
	Partitions       int32  `yaml:"partitions"`
}

// ----------------------
// Indicators properties.
// ----------------------

type IndicatorsProperties struct {
	Netcard   string                       `yaml:"netcard"`
	Physical  PhysicalIndicatorProperties  `yaml:"physical"`
	Virtual   VirtualIndicatorProperties   `yaml:"virtual"`
	Redis     RedisIndicatorProperties     `yaml:"redis"`
	Zookeeper ZookeeperIndicatorProperties `yaml:"zookeeper"`
	Kafka     KafkaIndicatorProperties     `yaml:"kafka"`
	Etcd      EtcdIndicatorProperties      `yaml:"etcd"`
	Emq       EmqIndicatorProperties       `yaml:"emq"`
	Consul    ConsulIndicatorProperties    `yaml:"consul"`
}

// Indicators physical properties.
type PhysicalIndicatorProperties struct {
	Delay     time.Duration `yaml:"delay"`
	RangePort string        `yaml:"range-port"`
}

// Indicators virtual properties.
type VirtualIndicatorProperties struct {
	Delay time.Duration `yaml:"delay"`
}

// Indicators redis properties.
type RedisIndicatorProperties struct {
	Delay time.Duration `yaml:"delay"`
	Ports string        `yaml:"ports"`
	Password string `yaml:"password"`
	Properties string `yaml:"properties"`

}

// Indicators zookeeper properties.
type ZookeeperIndicatorProperties struct {
	Delay time.Duration `yaml:"delay"`
	Host  string        `yaml:"host"`
	Command string `yaml:"command"`
	Properties string `yaml:"properties"`

}

// Indicators kafka properties.
type KafkaIndicatorProperties struct {
	Delay time.Duration `yaml:"delay"`
}

// Indicators etcd properties.
type EtcdIndicatorProperties struct {
	Delay time.Duration `yaml:"delay"`
}

// Indicators emq properties.
type EmqIndicatorProperties struct {
	Delay time.Duration `yaml:"delay"`
}

// Indicators consul properties.
type ConsulIndicatorProperties struct {
	Delay time.Duration `yaml:"delay"`
}

// Global configuration.
var GlobalConfig GlobalProperties

// Initialize global config properties.
func InitGlobalConfig(path string) {
	// Set defaults
	setDefaults()

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("Read config '%s' error! %s", path, err)
	}

	err = yaml.Unmarshal(yamlFile, &GlobalConfig)
	if err != nil {
		fmt.Printf("Unmarshal config '%s' error! %s", path, err)
	}
}

// Set defaults
func setDefaults() {
	globalConfig := &GlobalProperties{
		Logging: LoggingProperties{
			FileName: constant.DefaultLogFilename,
			Level:    constant.DefaultLogLevel,
			Policy: PolicyProperties{
				RetentionDays: constant.DefaultLogRetentionDays,
				MaxBackups:    constant.DefaultLogMaxBackups,
				MaxSize:       constant.DefaultLogMaxSize,
			},
		},
		Launcher: LauncherProperties{
			Http: HttpLauncherProperties{
				ServerGateway: constant.DefaultHttpServerGateway,
			},
			Kafka: KafkaLauncherProperties{
				Enabled:          false,
				BootstrapServers: constant.DefaultLauncherKafkaServers,
				Topic:            constant.DefaultLauncherKafkaTopic,
				Partitions:       constant.DefaultLauncherKafkaPartitions,
			},
		},
		Indicators: IndicatorsProperties{
			Netcard: constant.DefaultNetcard,
			Physical: PhysicalIndicatorProperties{
				Delay:     constant.DefaultIndicatorsDelay,
				RangePort: constant.DefaultNetIndicatorPortRange,
			},
			Virtual: VirtualIndicatorProperties{
				Delay: constant.DefaultIndicatorsDelay,
			},
			Redis: RedisIndicatorProperties{
				Delay: constant.DefaultIndicatorsDelay,
				Ports: constant.DefaultRedisPort,
				Properties: constant.DeaultRedisProperties,
			},
			Kafka: KafkaIndicatorProperties{
				Delay: constant.DefaultIndicatorsDelay,
			},
			Zookeeper: ZookeeperIndicatorProperties{
				Delay: constant.DefaultIndicatorsDelay,
				Host: constant.DefaultZookeeperHost,
				Command: constant.DefaultZookeeperHost,
				Properties: constant.DefaultZookeeperHost,
			},
			Etcd: EtcdIndicatorProperties{
				Delay: constant.DefaultIndicatorsDelay,
			},
			Emq: EmqIndicatorProperties{
				Delay: constant.DefaultIndicatorsDelay,
			},
			Consul: ConsulIndicatorProperties{
				Delay: constant.DefaultIndicatorsDelay,
			},
		},
	}
	GlobalConfig = *globalConfig
}
