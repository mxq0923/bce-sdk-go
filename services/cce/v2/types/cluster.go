// Copyright 2019 Baidu Inc. All rights reserved
// Use of this source code is governed by a CCE
// license that can be found in the LICENSE file.
/*
modification history
--------------------
2020/07/28 16:26:00, by jichao04@baidu.com, create
*/

package types

const (
	// LatestSupportedK8SVersion latest K8S Version that we supported
	LatestSupportedK8SVersion = "1.16.8"

	// DefaultRuntime default runtime
	DefaultRuntime = "docker"

	// LatestSupportedDockerVersion default docker version
	LatestSupportedDockerVersion = "18.09.2"

	CCEPrefix = "cce-"

	// ClusterIDLabelKey 关联 ClusterCRD 和 InstanceCRD 或 InstanceGroupCRD
	ClusterIDLabelKey = "cluster-id"

	ClusterRoleLabelKey = "cluster-role"

	DoNotHandle = "not-handler-by-cce"
)

// 创建集群时使用的ClusterSpec
type ClusterSpec struct {
	// 创建集群时无需传入ClusterID
	ClusterID string `json:"clusterID,omitempty" validate:"readonly"`

	// ClusterName 由用户指定
	ClusterName string `json:"clusterName" valid:"Required" validate:"modifiable"`

	ClusterType ClusterType `json:"clusterType,omitempty" valid:"Required" validate:"readonly"`

	Description string `json:"description,omitempty" validate:"modifiable"`

	K8SVersion K8SVersion `json:"k8sVersion,omitempty"`

	RuntimeType    RuntimeType `json:"runtimeType,omitempty"`
	RuntimeVersion string      `json:"runtimeVersion,omitempty"`

	// VPCUUID && VPCCIDR 无需用户设置
	VPCID string `json:"vpcID,omitempty" valid:"Required" validate:"readonly"`

	VPCUUID     string `json:"vpcUUID,omitempty" validate:"readonly"`
	VPCCIDR     string `json:"vpcCIDR,omitempty" validate:"readonly"`
	VPCCIDRIPv6 string `json:"vpcCIDRIPv6,omitempty" validate:"readonly"`

	// PluginListType CCE 插件类型
	Plugins []string `json:"plugins,omitempty"`

	// PluginsConfig 插件 Helm 安装配置
	PluginsConfig map[string]PluginHelmConfig `json:"pluginsConfig,omitempty"`

	MasterConfig           MasterConfig           `json:"masterConfig,omitempty" valid:"Required" validate:"inline"`
	ContainerNetworkConfig ContainerNetworkConfig `json:"containerNetworkConfig,omitempty" valid:"Required" validate:"inline"`

	// 集群删除保护标识，true 表示开启删除保护不允许删除集群；false 表示关闭删除保护允许删除集群
	ForbidDelete bool `json:"forbidDelete"`

	// IaaS资源付费选项
	ResourceChargingOption ResourceChargingOption `json:"resourceChargingOption,omitempty" validate:"inline"`

	// K8S 自定义配置
	K8SCustomConfig K8SCustomConfig `json:"k8sCustomConfig,omitempty"`

	// APIServer 认证模式
	AuthenticateMode AuthenticateMode `json:"authenticateMode,omitempty" validate:"readonly"`

	Tags []Tag `json:"tags,omitempty" validate:"readonly"`

	// 资源分组 ID
	ResourceGroupID string `json:"resourceGroupID,omitempty"`
}

// ResourceChargingOption 定义IaaS资源付费配置
type ResourceChargingOption struct {
	ChargingType      PaymentTiming `json:"chargingType,omitempty"`      // 后付费或预付费
	PurchaseTime      int           `json:"purchaseTime,omitempty"`      // 预付费才生效：单位月，12 = 12 月
	PurchaseTimeUnit  string        `json:"purchaseTimeUnit,omitempty"`  // 预付费时间单位
	AutoRenew         bool          `json:"autoRenew,omitempty"`         // 是否自动续费
	AutoRenewTime     int           `json:"autoRenewTime,omitempty"`     // 12 = 12 个月
	AutoRenewTimeUnit string        `json:"autoRenewTimeUnit,omitempty"` // 续费单位：月
}

// PluginHelmConfig 使用 Helm 部署插件的插件的参数
type PluginHelmConfig struct {
	// 插件类型(插件名称) 非必要 用户要部署的是哪个插件,传空时和PluginName保持一致
	PluginType string `json:"pluginType,omitempty"`

	// 插件别名 非必要 有时用户是可以自定义部署的插件名称的 (如多个 NGINX Ingress Controller 场景) 所以不能用PluginName来判断用户部署的是什么插件
	PluginName string `json:"pluginName,omitempty"`

	// 插件在云端的ChartName是什么 用户无需传递这个值
	ChartName string `json:"chartName,omitempty"`

	// 使用的Chart版本 除非用户要指定版本否则无需传递此值
	ChartVersion string `json:"chartVersion,omitempty"`

	// 插件部署到哪个命名空间  非必要
	Namespaces string `json:"namespaces,omitempty"`

	// 非必要
	Description string `json:"description,omitempty"`

	// 取决于插件 系统插件传空值即可
	Values string `json:"values,omitempty"`
}

// K8SCustomConfig - K8S 自定义配置
type K8SCustomConfig struct {
	MasterFeatureGates  map[string]bool `json:"masterFeatureGates,omitempty"`  // 自定义 FeatureGates
	NodeFeatureGates    map[string]bool `json:"nodeFeatureGates,omitempty"`    // 自定义 FeatureGates
	AdmissionPlugins    []string        `json:"admissionPlugins,omitempty"`    // 自定义 AdmissionPlugins
	PauseImage          string          `json:"pauseImage,omitempty"`          // 自定义 PauseImage
	KubeAPIQPS          int             `json:"kubeAPIQPS,omitempty"`          // 自定义 KubeAPIQPS
	KubeAPIBurst        int             `json:"kubeAPIBurst,omitempty"`        // 自定义 KubeAPIBurst
	SchedulerPredicates []string        `json:"schedulerPredicates,omitempty"` // 自定义 SchedulerPredicates
	SchedulerPriorities map[string]int  `json:"schedulerPriorities,omitempty"` // 自定义 SchedulerPriorities
	ETCDDataPath        string          `json:"etcdDataPath,omitempty"`        // 自定义 etcd数据目录
}

// ClusterType usually used to init Provider
// and it represents the difference between IaaS
type ClusterType string

const (
	// ClusterTypeNormal = 普通类型集群
	ClusterTypeNormal ClusterType = "normal"
)

// K8SVersion defines the k8stypes version of cluster
type K8SVersion string

const (
	// 1.6和1.8不再支持，扩缩容需要联系CCE人员手动操作
	// K8S_1_6_2   K8SVersion = "1.6.2"
	// K8S_1_8_6   K8SVersion = "1.8.6"
	// K8S_1_8_12  K8SVersion = "1.8.12"
	// 1.11.1 1.11.5 1.13.4仅支持已有集群扩容节点，不支持新创建集群
	// K8S_1_11_1  K8SVersion = "1.11.1"
	// K8S_1_11_5  K8SVersion = "1.11.5"
	// K8S_1_13_4  K8SVersion = "1.13.4"
	// 支持在console创建集群
	// K8S_1_13_10 K8SVersion = "1.13.10"
	// K8S_1_16_3  K8SVersion = "1.16.3"
	// K8S_1_16_8  K8SVersion = "1.16.8"
	// K8S_1_17_17 K8SVersion = "1.17.17"
	K8S_1_18_9  K8SVersion = "1.18.9"
	K8S_1_20_8  K8SVersion = "1.20.8"
	K8S_1_21_14 K8SVersion = "1.21.14"
	K8S_1_22_5  K8SVersion = "1.22.5"
	K8S_1_24_4  K8SVersion = "1.24.4"
	K8S_1_26_9  K8SVersion = "1.26.9"
)

// MasterConfig Master 配置
type MasterConfig struct {
	// MasterTypes: 托管, 自定义, 已有 BCC, 已有 BBC
	MasterType MasterType `json:"masterType,omitempty"`

	// ClusterHA 对 3 种集群都有效: 对于 Custom 和 Existed 作为校验和展示作用
	ClusterHA ClusterHA `json:"clusterHA,omitempty"`

	ExposedPublic bool `json:"exposedPublic,omitempty"`

	ClusterBLBVPCSubnetID string `json:"clusterBLBVPCSubnetID,omitempty"`

	ManagedClusterMasterOption `json:"managedClusterMasterOption,omitempty"`
}

// ManagedClusterMasterOption 托管集群 Master 配置
type ManagedClusterMasterOption struct {
	MasterVPCSubnetZone     AvailableZone `json:"masterVPCSubnetZone,omitempty"`
	MasterVPCSubnetUUID     string        `json:"masterVPCSubnetUUID,omitempty"`
	MasterSecurityGroupUUID string        `json:"masterSecurityGroupUUID,omitempty"`

	MasterFlavor MasterFlavor `json:"masterFlavor,omitempty"`
}

type MasterFlavor string

const (
	MasterFlavorL50   MasterFlavor = "l50"
	MasterFlavorL200  MasterFlavor = "l200"
	MasterFlavorL500  MasterFlavor = "l500"
	MasterFlavorL1000 MasterFlavor = "l1000"
	MasterFlavorL3000 MasterFlavor = "l3000"
	MasterFlavorL5000 MasterFlavor = "l5000"
)

// RuntimeType defines the runtime on each node
type RuntimeType string

const (
	RuntimeTypeDocker     RuntimeType = "docker"
	RuntimeTypeContainerd RuntimeType = "containerd"
)

// ContainerNetworkConfig defines the network config
// Some attrs have default value
type ContainerNetworkConfig struct {
	// CCE 支持网络类型: kubenet 及 vpc-cni
	Mode ContainerNetworkMode `json:"mode,omitempty"` // If not set, set mode = kubenet

	// ENI 网络模式子网
	ENIVPCSubnetIDs    map[AvailableZone][]string `json:"eniVPCSubnetIDs,omitempty"`
	ENISecurityGroupID string                     `json:"eniSecurityGroupID,omitempty"`

	// CCE 支持集群 IP version: dual stack, ipv4 only, ipv6 only
	IPVersion ContainerNetworkIPType `json:"ipVersion,omitempty"` // if not set, set ipv4

	// LB Service 关联 BLB 所在子网, 目前只能为普通子网
	LBServiceVPCSubnetID string `json:"lbServiceVPCSubnetID,omitempty" valid:"Required"`

	// 指定 NodePort Service 的端口范围
	NodePortRangeMin int `json:"nodePortRangeMin,omitempty"`
	NodePortRangeMax int `json:"nodePortRangeMax,omitempty"`

	// 集群 PodIP CIDR, 在 kubenet 网络模式下有效
	ClusterPodCIDR     string `json:"clusterPodCIDR,omitempty"`
	ClusterPodCIDRIPv6 string `json:"clusterPodCIDRIPv6,omitempty"`

	// Service ClusterIP 的 CIDR
	ClusterIPServiceCIDR     string `json:"clusterIPServiceCIDR,omitempty"`
	ClusterIPServiceCIDRIPv6 string `json:"clusterIPServiceCIDRIPv6,omitempty"`

	// 每个 Node 上最大的 Pod 数, 限制 NodeCIDR 的分配
	MaxPodsPerNode int `json:"maxPodsPerNode,omitempty"` // If not set, MaxPodsPerNode = 128

	// KubeProxy 的模式: iptables 和 ipvs
	KubeProxyMode KubeProxyMode `json:"kubeProxyMode,omitempty"` // If not set, kubeProxyMode=ipvs
}

// ContainerNetworkIPType - 容器 IP 类型
type ContainerNetworkIPType string

const (
	// ContainerNetworkIPTypeIPv4 - 容器网段 IPv4
	ContainerNetworkIPTypeIPv4 ContainerNetworkIPType = "ipv4"
	// ContainerNetworkIPTypeIPv6 - 容器网段 IPv6
	ContainerNetworkIPTypeIPv6 ContainerNetworkIPType = "ipv6"
	// ContainerNetworkIPTypeDualStack - 容器网段双栈
	ContainerNetworkIPTypeDualStack ContainerNetworkIPType = "dualStack"
)

// ContainerNetworkMode defines container config
type ContainerNetworkMode string

const (
	// ContainerNetworkModeKubenet using kubenet
	ContainerNetworkModeKubenet ContainerNetworkMode = "kubenet"

	// ContainerNetworkModeVPCCNI using vpc-cni
	ContainerNetworkModeVPCCNI ContainerNetworkMode = "vpc-cni"

	// ContainerNetworkModeVPCRouteVeth using vpc route plus veth
	ContainerNetworkModeVPCRouteVeth ContainerNetworkMode = "vpc-route-veth"

	// ContainerNetworkModeVPCRouteIPVlan using vpc route plus ipvlan
	ContainerNetworkModeVPCRouteIPVlan ContainerNetworkMode = "vpc-route-ipvlan"

	// ContainerNetworkModeVPCRouteAutoDetect using vpc route and auto detects veth or ipvlan due to kernel version
	ContainerNetworkModeVPCRouteAutoDetect ContainerNetworkMode = "vpc-route-auto-detect"

	// ContainerNetworkModeVPCSecondaryIPVeth using vpc secondary ip plus veth
	ContainerNetworkModeVPCSecondaryIPVeth ContainerNetworkMode = "vpc-secondary-ip-veth"

	// ContainerNetworkModeVPCSecondaryIPIPVlan using vpc secondary ip plus ipvlan
	ContainerNetworkModeVPCSecondaryIPIPVlan ContainerNetworkMode = "vpc-secondary-ip-ipvlan"

	// ContainerNetworkModeVPCSecondaryIPAutoDetect using vpc secondary ip and auto detects veth or ipvlan due to kernel version
	ContainerNetworkModeVPCSecondaryIPAutoDetect ContainerNetworkMode = "vpc-secondary-ip-auto-detect"
)

// KubeProxyMode defines kube-proxy --proxy-mode
// If not set, using KubeProxyModeIPVS as default
type KubeProxyMode string

const (
	// KubeProxyModeIPVS --proxy-mode=ipvs
	KubeProxyModeIPVS KubeProxyMode = "ipvs"

	// KubeProxyModeIptables --proxy-mode=iptables
	KubeProxyModeIptables KubeProxyMode = "iptables"
)

// MasterType 定义 Master 机器来源
type MasterType string

const (
	// MasterTypeManaged 托管 Master
	MasterTypeManaged MasterType = "managed"

	// MasterTypeCustom 自定义集群, 包含:
	// 1. 新建 BCC;
	// 2. 已有 BCC;
	// 3. 已有 BBC.
	MasterTypeCustom MasterType = "custom"

	// MasterTypeServerless Serverless集群Master
	MasterTypeServerless MasterType = "serverless"
)

// ClusterHA Cluster Master 对应副本数
type ClusterHA int

const (
	// ClusterHALow 单 Master
	ClusterHALow ClusterHA = 1
	// ClusterHAMedium 三 Master
	ClusterHAMedium ClusterHA = 3
	// ClusterHAHigh 五 Master
	ClusterHAHigh ClusterHA = 5
	// ClusterHAServerless Cluster Master 副本数
	ClusterHAServerless ClusterHA = 2
)

// ClusterPhase for CCE K8S Cluster Phase
type ClusterPhase string

const (
	// ClusterPhasePending 创建 Cluster 时默认状态
	ClusterPhasePending ClusterPhase = "pending"

	// ClusterPhaseProvisioning IaaS 相关资源正在创建中
	ClusterPhaseProvisioning ClusterPhase = "provisioning"

	// ClusterPhaseProvisioned IaaS 相关资源已经 Ready
	ClusterPhaseProvisioned ClusterPhase = "provisioned"

	// ClusterPhaseRunning 集群运行正常
	ClusterPhaseRunning ClusterPhase = "running"

	// ClusterPhaseCreateFailed 集群创建失败
	ClusterPhaseCreateFailed ClusterPhase = "create_failed"

	// ClusterPhaseDeleting 集群正在删除
	ClusterPhaseDeleting ClusterPhase = "deleting"

	// ClusterPhaseDeleted 集群删除完成
	ClusterPhaseDeleted ClusterPhase = "deleted"

	// ClusterPhaseDeleteFailed 集群删除失败
	ClusterPhaseDeleteFailed ClusterPhase = "delete_failed"
)

// AuthenticateMode - 认证类型
type AuthenticateMode string

const (
	// AuthenticateModeX509 - X509
	AuthenticateModeX509 AuthenticateMode = "x509"

	// AuthenticateModeOIDC - OIDC
	AuthenticateModeOIDC AuthenticateMode = "oidc"
)
