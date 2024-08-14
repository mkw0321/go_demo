// Package host_info
// 版本声明: Copyright 2022 烽台 Inc.
// Author: 王艺 <wangyi@fengtaisec.com>
// Desc: Windows 7 系统信息采集
// Date: 2022/7/28
package info

import (
	"DataGather/pkg/api"
	"DataGather/pkg/logger"
	"context"
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shopspring/decimal"
	"net"
	"os/user"
	"regexp"
	"strconv"
	"strings"
)

// 正则表达式模式
const ipv4Pattern = `(\b25[0-5]|\b2[0-4][0-9]|\b[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`
const ipv6Pattern = `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`

// AcquisitionProcessor 采集数据
func (m *MetricsHostInfo) AcquisitionProcessor() error {
	// TODO: 采集过程处理
	data := AcquisitionData{}
	if err := data.getHostInfo(); err != nil {
		logger.Error(context.Background(), "err", "getHostInfo err")
	}
	if err := data.getNetInfo(); err != nil {
		logger.Error(context.Background(), "err", "getNetInfo err")
	}
	if err := data.getCPUInfo(); err != nil {
		logger.Error(context.Background(), "err", "getCPUInfo err")
	}
	if err := data.getMemInfo(); err != nil {
		logger.Error(context.Background(), "err", "getMemInfo err")
	}
	if err := data.getDiskInfo(); err != nil {
		logger.Error(context.Background(), "err", "getDiskInfo err")
	}
	if err := data.getAccountInfo(); err != nil {
		logger.Error(context.Background(), "err", "getAccountInfo err")
	}
	// 补充数据
	data.LogType = m.logType

	m.collectInfo = data
	m.collectType = "object"
	return nil
}

// getAccountInfo 获取用户信息
func (data *AcquisitionData) getAccountInfo() error {
	user, err := user.Current()
	if err != nil {
		return err
	}
	Account := make(map[string]interface{})
	Account["用户名"] = user.Username
	if user.Uid == "0" {
		user.Uid = data.getAccountInfoByFile(2, user.Username)
	}
	Account["用户ID"] = user.Uid
	if user.Gid == "0" {
		user.Gid = data.getAccountInfoByFile(3, user.Username)
	}
	Account["用户组ID"] = user.Gid
	Account["家目录"] = user.HomeDir
	marshal, _ := json.Marshal(Account)
	data.Accounts = string(marshal)
	return nil
}

/*
/etc/passwd 是一个文本文件，其中包含了登录 Linux 系统所必需的每个用户的信息。它保存用户的有用信息，如
用户名：密码：用户 ID：群组 ID：用户 ID 信息：用户的家目录： Shell

cat /etc/passwd | grep root
root:x:0:0:root:/root:/bin/bash
root1:x:1000:1000:root,,,:/home/root1:/bin/bash
*/
func (data *AcquisitionData) getAccountInfoByFile(dataID int, username string) string {
	if dataID < 0 || dataID >= 7 {
		return "0"
	}
	fileData, err := api.ExecCommandLinuxGBK("cat /etc/passwd | grep " + username)
	if err != nil {
		return "0"
	}
	for _, fileInfo := range fileData {
		split := strings.Split(fileInfo, ":")
		if len(split) != 7 {
			return "0"
		} else {
			if split[0] != username {
				continue
			} else {
				return split[dataID]
			}
		}
	}
	return "0"
}

// getDiskInfo 获取硬盘信息
func (data *AcquisitionData) getDiskInfo() error {
	partitions, _ := disk.Partitions(false)
	var DiskList []map[string]interface{}
	for _, value := range partitions {
		usage, err := disk.Usage(value.Mountpoint)
		Disk := make(map[string]interface{})
		if err == nil {
			Disk["设备"] = value.Device
			Disk["操作权限"] = value.Opts
			Disk["挂载点"] = value.Mountpoint
			Disk["文件系统类型"] = value.Fstype
			f, _ := int2float64(usage.Total)
			Disk["容量"] = strconv.FormatFloat(f, 'f', 2, 64) + "MB"
		}
		DiskList = append(DiskList, Disk)
	}

	marshal, _ := json.Marshal(DiskList)
	data.Disk = string(marshal)

	return nil
}

// int2float64
func int2float64(a uint64) (float64, bool) {
	return decimal.NewFromFloat(float64(a)).Div(decimal.NewFromFloat(1024 * 1024)).Round(2).Float64()
}

// getMemInfo 获取内存信息
func (data *AcquisitionData) getMemInfo() error {
	//memory, err := mem.SwapMemory()
	//if err != nil {
	//	return err
	//}
	//
	//Memory := make(map[string]interface{})
	//capacity, _ := decimal.New(int64(memory.Total), 0).Div(decimal.NewFromFloat(1024 * 1024)).Round(2).Float64()
	//Memory["内存容量"] = strconv.FormatFloat(capacity, 'f', -1, 64) + "MB"
	//
	//marshal, _ := json.Marshal(Memory)
	//data.Memory = string(marshal)
	Memory := make(map[string]map[string]interface{})
	memoryInfo := make(map[string]interface{})
	command, err := api.ExecCommandLinux("free -m")
	if err != nil {
		return err
	}
	if len(command) >= 3 {
		fields := strings.Fields(command[1])
		if len(fields) >= 7 {
			memoryInfo["内存总容量"] = fields[1] + "MB"
			memoryInfo["内存使用空间"] = fields[2] + "MB"
			memoryInfo["内存空闲空间"] = fields[3] + "MB"
			memoryInfo["内存共享空间"] = fields[4] + "MB"
			memoryInfo["内存缓存空间"] = fields[5] + "MB"
			memoryInfo["内存可用空间"] = fields[6] + "MB"
		}
		fields2 := strings.Fields(command[2])
		if len(fields2) >= 4 {
			memoryInfo["内存交换总容量"] = fields2[1] + "MB"
			memoryInfo["内存交换使用空间"] = fields2[2] + "MB"
			memoryInfo["内存交换空闲空间"] = fields2[3] + "MB"
		}
	}
	Memory["内存信息"] = memoryInfo
	marshal, _ := json.Marshal(Memory)
	data.Memory = string(marshal)
	return nil
}

// getCPUInfo 获取 CPU 信息
func (data *AcquisitionData) getCPUInfo() error {
	info, _ := cpu.Info()
	logical, _ := cpu.Counts(true)
	physical, _ := cpu.Counts(false)

	Cpu := make(map[string]interface{})

	Cpu["名称"] = info[0].ModelName
	Cpu["物理核心"] = physical
	Cpu["逻辑核心"] = logical
	Cpu["基准频率"] = strconv.FormatFloat(info[0].Mhz, 'f', 2, 64) + "MHz"
	marshal, _ := json.Marshal(Cpu)
	data.Cpu = string(marshal)
	return nil
}

type Network struct {
	Name string `json:"名称"`
	Ipv4 []string
	Ipv6 []string
	Mac  string
}

// getNetInfo 获取网络信息
func (data *AcquisitionData) getNetInfo() error {
	var networks []Network
	var Ipv4, Ipv6, Mac []string

	interfaces, err := net.Interfaces()
	if err != nil {
		return err
	}

	for _, value := range interfaces {
		var network Network
		var mac = value.HardwareAddr.String()
		if mac == "" {
			continue
		}

		if addrs, err := value.Addrs(); err != nil || addrs == nil {
			continue
		} else if len(addrs) >= 1 {

			//2023-11-28 gaoxufan update
			for _, addr := range addrs {
				s := fmt.Sprint(addr)
				// 判断 IPv4 地址
				matchIPv4, err := regexp.MatchString(ipv4Pattern, s)
				if err != nil {
					logger.Error(context.Background(), "err", "getNetInfo err"+err.Error())
					continue
				}
				if matchIPv4 {
					network.Ipv4 = append(network.Ipv4, s)
					Ipv4 = append(Ipv4, s)
				} else {
					matchIPv6, err := regexp.MatchString(ipv6Pattern, s)
					if err != nil {
						logger.Error(context.Background(), "err", "getNetInfo err"+err.Error())
						continue
					}
					if matchIPv6 {
						Ipv6 = append(Ipv6, s)
						network.Ipv6 = append(network.Ipv6, s)
					}
				}
			}
			network.Mac = mac
			Mac = append(Mac, mac)
			network.Name = value.Name
			networks = append(networks, network)
		}
	}

	b, _ := json.Marshal(networks)

	data.Ipv4 = strings.Join(Ipv4, ",")
	data.Ipv6 = strings.Join(Ipv6, ",")
	data.Mac = strings.Join(Mac, ",")
	data.Network = string(b)
	return nil
}

// getHostInfo 获取系统信息
func (data *AcquisitionData) getHostInfo() error {
	info, _ := host.Info()

	// 数据克隆
	data.HostName = info.Hostname
	data.Os = info.Platform
	data.OsVersion = info.PlatformVersion
	data.OsArchitecture = info.KernelArch
	data.OsInstallDate = ""
	data.Kernel = info.KernelVersion
	data.Architecture = info.KernelArch
	data.LastBootupTime = info.BootTime
	return nil
}
