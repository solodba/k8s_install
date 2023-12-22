package conf

import (
	"fmt"
)

// 全局配置
var (
	c *Config
)

// 通过函数返回初始化的结构体
func C() *Config {
	if c == nil {
		panic("请初始化全局配置!")
	}
	return c
}

// Config结构体
type Config struct {
	Docker     *Docker     `toml:"docker"`
	Master     *Master     `toml:"master"`
	Node1      *Node1      `toml:"node1"`
	Node2      *Node2      `toml:"node2"`
	Containerd *Containerd `toml:"containerd"`
	K8s        *K8s        `toml:"k8s"`
}

// Docker结构体
type Docker struct {
	BaseDir               string `toml:"base_dir" env:"DOCKER_BASE_DIR"`
	Dir                   string `toml:"dir" env:"DOCKER_DIR"`
	LimitsFile            string `toml:"limits_file" env:"DOCKER_LIMITS_FILE"`
	SysctlFile            string `toml:"sysctl_file" env:"DOCKER_SYSCTL_FILE"`
	DockerServiceFile     string `toml:"docker_service_file" env:"DOCKER_SERVICE_FILE"`
	DockerSocketFile      string `toml:"docker_socket_file" env:"DOCKER_SOCKET_FILE"`
	DockerComposeFile     string `toml:"docker_compose_file" env:"DOCKER_COMPOSE_FILE"`
	DockerDaemonFile      string `toml:"docker_daemon_file" env:"DOCKER_DAEMON_FILE"`
	DockerBinaryFile      string `toml:"docker_binary_file" env:"DOCKER_BINARY_FILE"`
	DockerUser            string `toml:"docker_user" env:"DOCKER_USER"`
	ContainerdServiceFile string `toml:"containerd_service_file" env:"CONTAINERD_SERVICE_FILE"`
	CompressInstallFile   string `toml:"compress_install_file" env:"DOCKER_COMPRESS_INSTALL_FILE"`
}

// Containerd结构体
type Containerd struct {
	BaseDir              string `toml:"base_dir" env:"CONTAINERD_BASE_DIR"`
	Dir                  string `toml:"dir" env:"CONTAINERD_DIR"`
	LimitsFile           string `toml:"limits_file" env:"CONTAINERD_LIMITS_FILE"`
	SysctlFile           string `toml:"sysctl_file" env:"CONTAINERD_SYSCTL_FILE"`
	ContainerdFile       string `toml:"containerd_file" env:"CONTAINERD_FILE"`
	ContainerdConfigFile string `toml:"containerd_config_file" env:"CONTAINERD_CONFIG_FILE"`
	NerdctlFile          string `toml:"nerdctl_file" env:"CONTAINERD_NERDCTL_FILE"`
	NerdctlConfigFile    string `toml:"nerdctl_config_file" env:"CONTAINERD_NERDCTL_CONFIG_FILE"`
	CniFile              string `toml:"cni_file" env:"CONTAINERD_CNI_FILE"`
	RuncFile             string `toml:"runc_file" env:"CONTAINERD_RUNC_FILE"`
	IpvsModuleFile       string `toml:"ipvs_module_file" env:"CONTAINERD_IPVS_MODULE_FILE"`
	CompressInstallFile  string `toml:"compress_install_file" env:"CONTAINERD_COMPRESS_INSTALL_FILE"`
	ServiceFile          string `toml:"service_file" env:"CONTAINERD_SERVICE_FILE"`
}

// k8s结构体
type K8s struct {
	Dir                 string `toml:"dir" env:"K8S_DIR"`
	KubeadmVersion      string `toml:"kubeadm_version" env:"K8S_KUBEADM_VERSION"`
	KubeletVersion      string `toml:"kubelet_version" env:"K8S_KUBELET_VERSION"`
	KubectlVersion      string `toml:"kubectl_version" env:"K8S_KUBECTL_VERSION"`
	DownloadImageScript string `toml:"download_image_script" env:"K8S_DOWNLOAD_IMAGE_SCRIPT"`
	ApiserverAddress    string `toml:"apiserver_address" env:"K8S_APISERVER_ADDRESS"`
	ApiserverPort       int64  `toml:"apiserver_port" env:"K8S_APISERVER_PORT"`
	Version             string `toml:"version" env:"K8S_VERSION"`
	PodNetwork          string `toml:"pod_network" env:"K8S_POD_NETWORK"`
	ServiceNetwork      string `toml:"service_network" env:"K8S_SERVICE_NETWORK"`
	ServiceDnsDomain    string `toml:"service_dns_domain" env:"K8S_SERVICE_DNS_DOAMIN"`
}

// Master结构体
type Master struct {
	SysUsername string `toml:"sys_username" env:"MASTER_SYS_USERNAME"`
	SysPassword string `toml:"sys_password" env:"MASTER_SYS_PASSWORD"`
	SysHost     string `toml:"sys_host" env:"MASTER_SYS_HOST"`
	SysPort     int64  `toml:"sys_port" env:"MASTER_SYS_PORT"`
}

// Node1结构体
type Node1 struct {
	SysUsername string `toml:"sys_username" env:"NODE1_SYS_USERNAME"`
	SysPassword string `toml:"sys_password" env:"NODE1_SYS_PASSWORD"`
	SysHost     string `toml:"sys_host" env:"NODE1_SYS_HOST"`
	SysPort     int64  `toml:"sys_port" env:"NODE1_SYS_PORT"`
}

// Node2结构体
type Node2 struct {
	SysUsername string `toml:"sys_username" env:"NODE2_SYS_USERNAME"`
	SysPassword string `toml:"sys_password" env:"NODE2_SYS_PASSWORD"`
	SysHost     string `toml:"sys_host" env:"NODE2_SYS_HOST"`
	SysPort     int64  `toml:"sys_port" env:"NODE2_SYS_PORT"`
}

// Config构造函数
func NewDefaultConfig() *Config {
	return &Config{
		Docker:     NewDefaultDocker(),
		Master:     NewDefaultMaster(),
		Node1:      NewDefaultNode1(),
		Node2:      NewDefaultNode2(),
		Containerd: NewDefaultContainerd(),
		K8s:        NewDefaultK8s(),
	}
}

// Docker结构体构造函数
func NewDefaultDocker() *Docker {
	return &Docker{}
}

// Master构造函数
func NewDefaultMaster() *Master {
	return &Master{}
}

// Node1构造函数
func NewDefaultNode1() *Node1 {
	return &Node1{}
}

// Node2构造函数
func NewDefaultNode2() *Node2 {
	return &Node2{}
}

// Containerd构造函数
func NewDefaultContainerd() *Containerd {
	return &Containerd{}
}

// K8s构造函数
func NewDefaultK8s() *K8s {
	return &K8s{}
}

// 获取Docker文件路径
func (m *Docker) LimitsFilePath() string {
	return fmt.Sprintf("%s/%s", m.Dir, m.LimitsFile)
}

func (m *Docker) SysctlFilePath() string {
	return fmt.Sprintf("%s/%s", m.Dir, m.SysctlFile)
}

func (m *Docker) DockerServiceFilePath() string {
	return fmt.Sprintf("%s/%s", m.Dir, m.DockerServiceFile)
}

func (m *Docker) DockerSocketFilePath() string {
	return fmt.Sprintf("%s/%s", m.Dir, m.DockerSocketFile)
}

func (m *Docker) DockerComposeFilePath() string {
	return fmt.Sprintf("%s/%s", m.Dir, m.DockerComposeFile)
}

func (m *Docker) DockerDaemonFilePath() string {
	return fmt.Sprintf("%s/%s", m.Dir, m.DockerDaemonFile)
}

func (m *Docker) DockerBinaryFilePath() string {
	return fmt.Sprintf("%s/%s", m.Dir, m.DockerBinaryFile)
}

func (m *Docker) ContainerdServiceFilePath() string {
	return fmt.Sprintf("%s/%s", m.Dir, m.ContainerdServiceFile)
}

func (m *Docker) DockerCompressInstallFilePath() string {
	return fmt.Sprintf("%s/%s", m.BaseDir, m.CompressInstallFile)
}

// 获取Containerd文件路径
func (m *Containerd) LimitsFilePath() string {
	return fmt.Sprintf("%s/%s", m.Dir, m.LimitsFile)
}

func (m *Containerd) SysctlFilePath() string {
	return fmt.Sprintf("%s/%s", m.Dir, m.SysctlFile)
}

func (m *Containerd) ContainerdFilePath() string {
	return fmt.Sprintf("%s/%s", m.Dir, m.ContainerdFile)
}

func (m *Containerd) NerdctlFilePath() string {
	return fmt.Sprintf("%s/%s", m.Dir, m.NerdctlFile)
}

func (m *Containerd) NerdctlConfigFilePath() string {
	return fmt.Sprintf("%s/%s", m.Dir, m.NerdctlConfigFile)
}

func (m *Containerd) CniFilePath() string {
	return fmt.Sprintf("%s/%s", m.Dir, m.CniFile)
}

func (m *Containerd) RuncFilePath() string {
	return fmt.Sprintf("%s/%s", m.Dir, m.RuncFile)
}

func (m *Containerd) IpvsModuleFilePath() string {
	return fmt.Sprintf("%s/%s", m.Dir, m.IpvsModuleFile)
}

func (m *Containerd) ContainerdConfigFilePath() string {
	return fmt.Sprintf("%s/%s", m.Dir, m.ContainerdConfigFile)
}

func (m *Containerd) ContainerdCompressInstallFilePath() string {
	return fmt.Sprintf("%s/%s", m.BaseDir, m.CompressInstallFile)
}

func (m *Containerd) ContainerdServiceFilePath() string {
	return fmt.Sprintf("%s/%s", m.Dir, m.ServiceFile)
}

// 获取k8s下载镜像脚本
func (m *K8s) DownloadImageScriptPath() string {
	return fmt.Sprintf("%s/%s", m.Dir, m.DownloadImageScript)
}
