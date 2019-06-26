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
package main

import (
	"flag"
	"sync"
	"umc-agent/pkg/config"
	"umc-agent/pkg/constant"
	"umc-agent/pkg/monitor/docker"
	"umc-agent/pkg/monitor/physical"
	"umc-agent/pkg/monitor/redis"
	"umc-agent/pkg/monitor/zookeeper"
	"umc-agent/pkg/transport"
)

func init() {
	var confPath = constant.DefaultConfigPath

	// Command config path
	flag.StringVar(&confPath, "c", constant.DefaultConfigPath, "Config must is required!")
	flag.Parse()
	//flag.Usage()

	//logger.Main.Info("Initialize config file", zap.String("confPath", confPath))

	// Init global config.
	config.InitGlobalConfig(confPath)

	// Init kafka launcher.(if necessary)
	transport.InitKafkaLauncherIfNecessary()
}

func main() {
	wg := &sync.WaitGroup{}
	startingIndicatorRunners(wg)
	wg.Wait()
}

// Starting indicator runners all
func startingIndicatorRunners(wg *sync.WaitGroup) {
	wg.Add(1)
	go physical.IndicatorRunner()
	go docker.DockerIndicatorRunner()
	go redis.IndicatorRunner()
	go zookeeper.IndicatorRunner()
}
