package v3

import "time"

//Reference ...
type Reference struct {
	Kind *string `json:"kind"`
	Name *string `json:"name,omitempty"`
	UUID *string `json:"uuid"`
}

//VMMetadata Metadata The vm kind metadata
type VMMetadata struct {

	// Categories for the vm
	Categories map[string]string `json:"categories,omitempty"`

	// UTC date and time in RFC-3339 format when vm was created
	CreationTime *time.Time `json:"creation_time,omitempty"`

	// The kind name
	Kind *string `json:"kind"`

	// UTC date and time in RFC-3339 format when vm was last updated
	LastUpdateTime *time.Time `json:"last_update_time,omitempty"`

	// vm name
	Name *string `json:"name,omitempty"`

	// project reference
	ProjectReference *Reference `json:"project_reference,omitempty"`

	OwnerReference *Reference `json:"owner_reference,omitempty"`

	// Hash of the spec. This will be returned from server.
	SpecHash *string `json:"spec_hash,omitempty"`

	// Version number of the latest spec.
	SpecVersion *int64 `json:"spec_version,omitempty"`

	// vm uuid
	UUID *string `json:"uuid,omitempty"`
}

//VMVnumaConfig Indicates how VM vNUMA should be configured
type VMVnumaConfig struct {

	// Number of vNUMA nodes. 0 means vNUMA is disabled.
	NumVnumaNodes *int64 `json:"num_vnuma_nodes,omitempty"`
}

//IPAddress An IP address.
type IPAddress struct {

	// Address *string.
	IP *string `json:"ip,omitempty"`

	// Address type. It can only be \"ASSIGNED\" in the spec. If no type is specified in the spec, the default type is set to \"ASSIGNED\".
	Type *string `json:"type,omitempty"`
}

//VMNic Virtual Machine NIC.
type VMNic struct {

	// IP endpoints for the adapter. Currently, IPv4 addresses are supported.
	IPEndpointList []*IPAddress `json:"ip_endpoint_list,omitempty"`

	// The MAC address for the adapter.
	MacAddress *string `json:"mac_address,omitempty"`

	// The model of this NIC.
	Model *string `json:"model,omitempty"`

	NetworkFunctionChainReference *Reference `json:"network_function_chain_reference,omitempty"`

	// The type of this Network function NIC. Defaults to INGRESS.
	NetworkFunctionNicType *string `json:"network_function_nic_type,omitempty"`

	// The type of this NIC. Defaults to NORMAL_NIC.
	NicType *string `json:"nic_type,omitempty"`

	SubnetReference *Reference `json:"subnet_reference,omitempty"`

	// The NIC's UUID, which is used to uniquely identify this particular NIC. This UUID may be used to refer to the NIC outside the context of the particular VM it is attached to.
	UUID *string `json:"uuid,omitempty"`
}

//DiskAddress Disk Address.
type DiskAddress struct {
	AdapterType *string `json:"adapter_type"`

	DeviceIndex *int64 `json:"device_index"`
}

//VMBootDevice Indicates which device a VM should boot from. One of disk_address or mac_address should be provided.
type VMBootDevice struct {

	// Address of disk to boot from.
	DiskAddress *DiskAddress `json:"disk_address,omitempty"`

	// MAC address of nic to boot from.
	MacAddress *string `json:"mac_address,omitempty"`
}

//VMBootConfig Indicates which device a VM should boot from.
type VMBootConfig struct {

	// Indicates which device a VM should boot from. Boot device takes precdence over boot device order. If both are given then specified boot device will be primary boot device and remaining devices will be assigned boot order according to boot device order field.
	BootDevice *VMBootDevice `json:"boot_device,omitempty"`

	// Indicates the order of device types in which VM should try to boot from. If boot device order is not provided the system will decide appropriate boot device order.
	BootDeviceOrderList []*string `json:"boot_device_order_list,omitempty"`
}

//NutanixGuestToolsSpec Information regarding Nutanix Guest Tools.
type NutanixGuestToolsSpec struct {

	// Application names that are enabled.
	EnabledCapabilityList []*string `json:"enabled_capability_list,omitempty"`

	// Desired mount state of Nutanix Guest Tools ISO.
	IsoMountState *string `json:"iso_mount_state,omitempty"`

	// Nutanix Guest Tools is enabled or not.
	State *string `json:"state,omitempty"`
}

//GuestToolsSpec Information regarding guest tools.
type GuestToolsSpec struct {

	// Nutanix Guest Tools information
	NutanixGuestTools *NutanixGuestToolsSpec `json:"nutanix_guest_tools,omitempty"`
}

//VMGpu Graphics resource information for the Virtual Machine.
type VMGpu struct {

	// The device ID of the GPU.
	DeviceID *int64 `json:"device_id,omitempty"`

	// The mode of this GPU.
	Mode *string `json:"mode,omitempty"`

	// The vendor of the GPU.
	Vendor *string `json:"vendor,omitempty"`
}

//GuestCustomizationCloudInit If this field is set, the guest will be customized using cloud-init. Either user_data or custom_key_values should be provided. If custom_key_ves are provided then the user data will be generated using these key-value pairs.
type GuestCustomizationCloudInit struct {

	// Generic key value pair used for custom attributes
	CustomKeyValues map[string]string `json:"custom_key_values,omitempty"`

	// The contents of the meta_data configuration for cloud-init. This can be formatted as YAML or JSON. The value must be base64 encoded.
	MetaData *string `json:"meta_data,omitempty"`

	// The contents of the user_data configuration for cloud-init. This can be formatted as YAML, JSON, or could be a shell script. The value must be base64 encoded.
	UserData *string `json:"user_data,omitempty"`
}

//GuestCustomizationSysprep If this field is set, the guest will be customized using Sysprep. Either unattend_xml or custom_key_values should be provided. If custom_key_values are provided then the unattended answer file will be generated using these key-value pairs.
type GuestCustomizationSysprep struct {

	// Generic key value pair used for custom attributes
	CustomKeyValues map[string]string `json:"custom_key_values,omitempty"`

	// Whether the guest will be freshly installed using this unattend configuration, or whether this unattend configuration will be applied to a pre-prepared image. Default is \"PREPARED\".
	InstallType *string `json:"install_type,omitempty"`

	// This field contains a Sysprep unattend xml definition, as a *string. The value must be base64 encoded.
	UnattendXML *string `json:"unattend_xml,omitempty"`
}

//GuestCustomization VM guests may be customized at boot time using one of several different methods. Currently, cloud-init w/ ConfigDriveV2 (for Linux VMs) and Sysprep (for Windows VMs) are supported. Only ONE OF sysprep or cloud_init should be provided. Note that guest customization can currently only be set during VM creation. Attempting to change it after creation will result in an error. Additional properties can be specified. For example - in the context of VM template creation if \"override_script\" is set to \"True\" then the deployer can upload their own custom script.
type GuestCustomization struct {
	CloudInit *GuestCustomizationCloudInit `json:"cloud_init,omitempty"`

	// Flag to allow override of customization by deployer.
	IsOverridable *bool `json:"is_overridable,omitempty"`

	Sysprep map[string]interface{} `json:"sysprep,omitempty"`
}

//VMGuestPowerStateTransitionConfig Extra configs related to power state transition.
type VMGuestPowerStateTransitionConfig struct {

	// Indicates whether to execute set script before ngt shutdown/reboot.
	EnableScriptExec *bool `json:"enable_script_exec,omitempty"`

	// Indicates whether to abort ngt shutdown/reboot if script fails.
	ShouldFailOnScriptFailure *bool `json:"should_fail_on_script_failure,omitempty"`
}

//VMPowerStateMechanism Indicates the mechanism guiding the VM power state transition. Currently used for the transition to \"OFF\" state.
type VMPowerStateMechanism struct {
	GuestTransitionConfig *VMGuestPowerStateTransitionConfig `json:"guest_transition_config,omitempty"`

	// Power state mechanism (ACPI/GUEST/HARD).
	Mechanism *string `json:"mechanism,omitempty"`
}

//VMDiskDeviceProperties ...
type VMDiskDeviceProperties struct {
	DeviceType *string `json:"device_type,omitempty"`

	DiskAddress DiskAddress `json:"disk_address,omitempty"`
}

//VMDisk VirtualMachine Disk (VM Disk).
type VMDisk struct {
	DataSourceReference *Reference `json:"data_source_reference,omitempty"`

	DeviceProperties *VMDiskDeviceProperties `json:"device_properties,omitempty"`

	// Size of the disk in Bytes.
	DiskSizeBytes *int64 `json:"disk_size_bytes,omitempty"`

	// Size of the disk in MiB. Must match the size specified in 'disk_size_bytes' - rounded up to the nearest MiB -  when that field is present.
	DiskSizeMib *int64 `json:"disk_size_mib,omitempty"`

	// The device ID which is used to uniquely identify this particular disk.
	UUID *string `json:"uuid,omitempty"`

	VolumeGroupReference *Reference `json:"volume_group_reference,omitempty"`
}

//VMResources VM Resources Definition.
type VMResources struct {

	// Indicates which device the VM should boot from.
	BootConfig *VMBootConfig `json:"boot_config,omitempty"`

	// Disks attached to the VM.
	DiskList []*VMDisk `json:"disk_list,omitempty"`

	// GPUs attached to the VM.
	GpuList []*VMGpu `json:"gpu_list,omitempty"`

	GuestCustomization *GuestCustomization `json:"guest_customization,omitempty"`

	// Guest OS Identifier. For ESX, refer to VMware documentation link (https://www.vmware.com/support/developer/converter-sdk/conv43_apireference/vim.vm.GuestOsDescriptor.GuestOsIdentifier.html) for the list of guest OS identifiers.
	GuestOsID *string `json:"guest_os_id,omitempty"`

	// Information regarding guest tools.
	GuestTools *GuestToolsSpec `json:"guest_tools,omitempty"`

	// VM's hardware clock timezone in IANA TZDB format (America/Los_Angeles).
	HardwareClockTimezone *string `json:"hardware_clock_timezone,omitempty"`

	// Memory size in MiB.
	MemorySizeMib *int64 `json:"memory_size_mib,omitempty"`

	// NICs attached to the VM.
	NicList []VMNic `json:"nic_list,omitempty"`

	// Number of vCPU sockets.
	NumSockets *int64 `json:"num_sockets,omitempty"`

	// Number of vCPUs per socket.
	NumVcpusPerSocket *int64 `json:"num_vcpus_per_socket,omitempty"`

	// *Reference to an entity that the VM should be cloned from.
	ParentReference *Reference `json:"parent_reference,omitempty"`

	// The current or desired power state of the VM.
	PowerState *string `json:"power_state,omitempty"`

	PowerStateMechanism *VMPowerStateMechanism `json:"power_state_mechanism,omitempty"`

	// Indicates whether VGA console should be enabled or not.
	VgaConsoleEnabled *bool `json:"vga_console_enabled,omitempty"`

	// Information regarding vNUMA configuration.
	VMVnumaConfig *VMVnumaConfig `json:"vnuma_config,omitempty"`
}

//VM An intentful representation of a vm spec
type VM struct {
	AvailabilityZoneReference *Reference `json:"availability_zone_reference,omitempty"`

	ClusterReference *Reference `json:"cluster_reference,omitempty"`

	// A description for vm.
	Description *string `json:"description,omitempty"`

	// vm Name.
	Name *string `json:"name"`

	Resources *VMResources `json:"resources,omitempty"`
}

//VMIntentInput ...
type VMIntentInput struct {
	APIVersion *string `json:"api_version,omitempty"`

	Metadata *VMMetadata `json:"metadata"`

	Spec *VM `json:"spec"`
}

//MessageResource ...
type MessageResource struct {

	// Custom key-value details relevant to the status.
	Details map[string]string `json:"details,omitempty"`

	// If state is ERROR, a message describing the error.
	Message *string `json:"message"`

	// If state is ERROR, a machine-readable snake-cased *string.
	Reason *string `json:"reason"`
}

//VMStatus The status of a REST API call. Only used when there is a failure to report.
type VMStatus struct {
	APIVersion *string `json:"api_version,omitempty"`

	// The HTTP error code.
	Code *int64 `json:"code,omitempty"`

	// The kind name
	Kind *string `json:"kind,omitempty"`

	MessageList []*MessageResource `json:"message_list,omitempty"`

	State *string `json:"state,omitempty"`
}

//VMNicOutputStatus Virtual Machine NIC Status.
type VMNicOutputStatus struct {

	// The Floating IP associated with the vnic.
	FloatingIP *string `json:"floating_ip,omitempty"`

	// IP endpoints for the adapter. Currently, IPv4 addresses are supported.
	IPEndpointList []*IPAddress `json:"ip_endpoint_list,omitempty"`

	// The MAC address for the adapter.
	MacAddress *string `json:"mac_address,omitempty"`

	// The model of this NIC.
	Model *string `json:"model,omitempty"`

	NetworkFunctionChainReference *Reference `json:"network_function_chain_reference,omitempty"`

	// The type of this Network function NIC. Defaults to INGRESS.
	NetworkFunctionNicType *string `json:"network_function_nic_type,omitempty"`

	// The type of this NIC. Defaults to NORMAL_NIC.
	NicType *string `json:"nic_type,omitempty"`

	SubnetReference *Reference `json:"subnet_reference,omitempty"`

	// The NIC's UUID, which is used to uniquely identify this particular NIC. This UUID may be used to refer to the NIC outside the context of the particular VM it is attached to.
	UUID *string `json:"uuid,omitempty"`
}

//NutanixGuestToolsStatus Information regarding Nutanix Guest Tools.
type NutanixGuestToolsStatus struct {

	// Version of Nutanix Guest Tools available on the cluster.
	AvailableVersion *string `json:"available_version,omitempty"`

	// Application names that are enabled.
	EnabledCapabilityList []*string `json:"enabled_capability_list,omitempty"`

	// Version of the operating system on the VM.
	GuestOsVersion *string `json:"guest_os_version,omitempty"`

	// Communication from VM to CVM is active or not.
	IsReachable *bool `json:"is_reachable,omitempty"`

	// Desired mount state of Nutanix Guest Tools ISO.
	IsoMountState *string `json:"iso_mount_state,omitempty"`

	// Nutanix Guest Tools is enabled or not.
	State *string `json:"state,omitempty"`

	// Version of Nutanix Guest Tools installed on the VM.
	Version *string `json:"version,omitempty"`

	// Whether VM mobility drivers are installed in the VM.
	VMMobilityDriversInstalled *bool `json:"vm_mobility_drivers_installed,omitempty"`

	// Whether the VM is configured to take VSS snapshots through NGT.
	VSSSnapshotCapable *bool `json:"vss_snapshot_capable,omitempty"`
}

//GuestToolsStatus Information regarding guest tools.
type GuestToolsStatus struct {

	// Nutanix Guest Tools information
	NutanixGuestTools *NutanixGuestToolsStatus `json:"nutanix_guest_tools,omitempty"`
}

//VMGpuOutputStatus Graphics resource status information for the Virtual Machine.
type VMGpuOutputStatus struct {

	// The device ID of the GPU.
	DeviceID *int64 `json:"device_id,omitempty"`

	// Fraction of the physical GPU assigned.
	Fraction *int64 `json:"fraction,omitempty"`

	// GPU frame buffer size in MiB.
	FrameBufferSizeMib *int64 `json:"frame_buffer_size_mib,omitempty"`

	// Last determined guest driver version.
	GuestDriverVersion *string `json:"guest_driver_version,omitempty"`

	// The mode of this GPU
	Mode *string `json:"mode,omitempty"`

	// Name of the GPU resource.
	Name *string `json:"name,omitempty"`

	// Number of supported virtual display heads.
	NumVirtualDisplayHeads *int64 `json:"num_virtual_display_heads,omitempty"`

	// GPU {segment:bus:device:function} (sbdf) address if assigned.
	PCIAddress *string `json:"pci_address,omitempty"`

	// UUID of the GPU.
	UUID *string `json:"uuid,omitempty"`

	// The vendor of the GPU.
	Vendor *string `json:"vendor,omitempty"`
}

//GuestCustomizationStatus VM guests may be customized at boot time using one of several different methods. Currently, cloud-init w/ ConfigDriveV2 (for Linux VMs) and Sysprep (for Windows VMs) are supported. Only ONE OF sysprep or cloud_init should be provided. Note that guest customization can currently only be set during VM creation. Attempting to change it after creation will result in an error. Additional properties can be specified. For example - in the context of VM template creation if \"override_script\" is set to \"True\" then the deployer can upload their own custom script.
type GuestCustomizationStatus struct {
	CloudInit *GuestCustomizationCloudInit `json:"cloud_init,omitempty"`

	// Flag to allow override of customization by deployer.
	IsOverridable *bool `json:"is_overridable,omitempty"`

	Sysprep *GuestCustomizationSysprep `json:"sysprep,omitempty"`
}

//VMResourcesDefStatus VM Resources Status Definition.
type VMResourcesDefStatus struct {

	// Indicates which device the VM should boot from.
	BootConfig *VMBootConfig `json:"boot_config,omitempty"`

	// Disks attached to the VM.
	DiskList []*VMDisk `json:"disk_list,omitempty"`

	// GPUs attached to the VM.
	GpuList []*VMGpuOutputStatus `json:"gpu_list,omitempty"`

	GuestCustomization *GuestCustomizationStatus `json:"guest_customization,omitempty"`

	// Guest OS Identifier. For ESX, refer to VMware documentation link (https://www.vmware.com/support/developer/converter-sdk/conv43_apireference/vim.vm.GuestOsDescriptor.GuestOsIdentifier.html) for the list of guest OS identifiers.
	GuestOsID *string `json:"guest_os_id,omitempty"`

	// Information regarding guest tools.
	GuestTools *GuestToolsStatus `json:"guest_tools,omitempty"`

	// VM's hardware clock timezone in IANA TZDB format (America/Los_Angeles).
	HardwareClockTimezone *string `json:"hardware_clock_timezone,omitempty"`

	HostReference *Reference `json:"host_reference,omitempty"`

	// The hypervisor type for the hypervisor the VM is hosted on.
	HypervisorType *string `json:"hypervisor_type,omitempty"`

	// Memory size in MiB.
	MemorySizeMib *int64 `json:"memory_size_mib,omitempty"`

	// NICs attached to the VM.
	NicList []*VMNicOutputStatus `json:"nic_list,omitempty"`

	// Number of vCPU sockets.
	NumSockets *int64 `json:"num_sockets,omitempty"`

	// Number of vCPUs per socket.
	NumVcpusPerSocket *int64 `json:"num_vcpus_per_socket,omitempty"`

	// *Reference to an entity that the VM cloned from.
	ParentReference *Reference `json:"parent_reference,omitempty"`

	// Current power state of the VM.
	PowerState *string `json:"power_state,omitempty"`

	PowerStateMechanism *VMPowerStateMechanism `json:"power_state_mechanism,omitempty"`

	// Indicates whether VGA console has been enabled or not.
	VgaConsoleEnabled *bool `json:"vga_console_enabled,omitempty"`

	// Information regarding vNUMA configuration.
	VnumaConfig *VMVnumaConfig `json:"vnuma_config,omitempty"`
}

//VMDefStatus An intentful representation of a vm status
type VMDefStatus struct {
	AvailabilityZoneReference *Reference `json:"availability_zone_reference,omitempty"`

	ClusterReference *Reference `json:"cluster_reference,omitempty"`

	// A description for vm.
	Description *string `json:"description,omitempty"`

	// Any error messages for the vm, if in an error state.
	MessageList []*MessageResource `json:"message_list,omitempty"`

	// vm Name.
	Name *string `json:"name,omitempty"`

	Resources *VMResourcesDefStatus `json:"resources,omitempty"`

	// The state of the vm.
	State *string `json:"state,omitempty"`
}

//VMIntentResponse Response object for intentful operations on a vm
type VMIntentResponse struct {
	APIVersion *string `json:"api_version"`

	Metadata *VMMetadata `json:"metadata,omitempty"`

	Spec *VM `json:"spec,omitempty"`

	Status *VMDefStatus `json:"status,omitempty"`
}

//VMListMetadata All api calls that return a list will have this metadata block as input
type VMListMetadata struct {

	// The filter in FIQL syntax used for the results.
	Filter *string `json:"filter,omitempty"`

	// The kind name
	Kind *string `json:"kind,omitempty"`

	// The number of records to retrieve relative to the offset
	Length *int64 `json:"length,omitempty"`

	// Offset from the start of the entity list
	Offset *int64 `json:"offset,omitempty"`

	// The attribute to perform sort on
	SortAttribute *string `json:"sort_attribute,omitempty"`

	// The sort order in which results are returned
	SortOrder *string `json:"sort_order,omitempty"`
}

//VMIntentResource Response object for intentful operations on a vm
type VMIntentResource struct {
	APIVersion *string `json:"api_version,omitempty"`

	Metadata *VMMetadata `json:"metadata"`

	Spec *VM `json:"spec,omitempty"`

	Status *VMDefStatus `json:"status,omitempty"`
}

//VMListMetadataOutput All api calls that return a list will have this metadata block
type VMListMetadataOutput struct {

	// The filter used for the results
	Filter *string `json:"filter,omitempty"`

	// The kind name
	Kind *string `json:"kind,omitempty"`

	// The number of records retrieved relative to the offset
	Length *int64 `json:"length,omitempty"`

	// Offset from the start of the entity list
	Offset *int64 `json:"offset,omitempty"`

	// The attribute to perform sort on
	SortAttribute *string `json:"sort_attribute,omitempty"`

	// The sort order in which results are returned
	SortOrder *string `json:"sort_order,omitempty"`

	// Total matches found
	TotalMatches *int64 `json:"total_matches,omitempty"`
}

//VMListIntentResponse Response object for intentful operation of vms
type VMListIntentResponse struct {
	APIVersion *string `json:"api_version"`

	Entities []*VMIntentResource `json:"entities,omitempty"`

	Metadata *VMListMetadataOutput `json:"metadata"`
}

//SubnetMetadata The subnet kind metadata
type SubnetMetadata struct {

	// Categories for the subnet
	Categories map[string]string `json:"categories,omitempty"`

	// UTC date and time in RFC-3339 format when subnet was created
	CreationTime *time.Time `json:"creation_time,omitempty"`

	// The kind name
	Kind *string `json:"kind"`

	// UTC date and time in RFC-3339 format when subnet was last updated
	LastUpdateTime *time.Time `json:"last_update_time,omitempty"`

	// subnet name
	Name *string `json:"name,omitempty"`

	OwnerReference *Reference `json:"owner_reference,omitempty"`

	// project reference
	ProjectReference *Reference `json:"project_reference,omitempty"`

	// Hash of the spec. This will be returned from server.
	SpecHash *string `json:"spec_hash,omitempty"`

	// Version number of the latest spec.
	SpecVersion *int64 `json:"spec_version,omitempty"`

	// subnet uuid
	UUID *string `json:"uuid,omitempty"`
}

//Address represents the Host address.
type Address struct {

	// Fully qualified domain name.
	FQDN *string `json:"fqdn,omitempty"`

	// IPV4 address.
	IP *string `json:"ip,omitempty"`

	// IPV6 address.
	IPV6 *string `json:"ipv6,omitempty"`

	// Port Number
	Port *int64 `json:"port,omitempty"`
}

//IPPool represents IP pool.
type IPPool struct {

	// Range of IPs (example: 10.0.0.9 10.0.0.19).
	Range *string `json:"range,omitempty"`
}

//DHCPOptions Spec for defining DHCP options.
type DHCPOptions struct {
	BootFileName *string `json:"boot_file_name,omitempty"`

	DomainName *string `json:"domain_name,omitempty"`

	DomainNameServerList []*string `json:"domain_name_server_list,omitempty"`

	DomainSearchList []*string `json:"domain_search_list,omitempty"`

	TFTPServerName *string `json:"tftp_server_name,omitempty"`
}

//IPConfig represents the configurtion of IP.
type IPConfig struct {

	// Default gateway IP address.
	DefaultGatewayIP *string `json:"default_gateway_ip,omitempty"`

	DHCPOptions *DHCPOptions `json:"dhcp_options,omitempty"`

	DHCPServerAddress *Address `json:"dhcp_server_address,omitempty"`

	PoolList []*IPPool `json:"pool_list,omitempty"`

	PrefixLength *int64 `json:"prefix_length,omitempty"`

	// Subnet IP address.
	SubnetIP *string `json:"subnet_ip,omitempty"`
}

//SubnetResources represents Subnet creation/modification spec.
type SubnetResources struct {
	IPConfig *IPConfig `json:"ip_config,omitempty"`

	NetworkFunctionChainReference *Reference `json:"network_function_chain_reference,omitempty"`

	SubnetType *string `json:"subnet_type"`

	VlanID *int64 `json:"vlan_id,omitempty"`

	VswitchName *string `json:"vswitch_name,omitempty"`
}

//Subnet An intentful representation of a subnet spec
type Subnet struct {
	AvailabilityZoneReference *Reference `json:"availability_zone_reference,omitempty"`

	ClusterReference *Reference `json:"cluster_reference,omitempty"`

	// A description for subnet.
	Description *string `json:"description,omitempty"`

	// subnet Name.
	Name *string `json:"name"`

	Resources *SubnetResources `json:"resources,omitempty"`
}

//SubnetIntentInput An intentful representation of a subnet
type SubnetIntentInput struct {
	APIVersion *string `json:"api_version,omitempty"`

	Metadata *SubnetMetadata `json:"metadata"`

	Spec *Subnet `json:"spec"`
}

//SubnetStatus represents The status of a REST API call. Only used when there is a failure to report.
type SubnetStatus struct {
	APIVersion *string `json:"api_version,omitempty"`

	// The HTTP error code.
	Code *int64 `json:"code,omitempty"`

	// The kind name
	Kind *string `json:"kind,omitempty"`

	MessageList []*MessageResource `json:"message_list,omitempty"`

	State *string `json:"state,omitempty"`
}

//SubnetResourcesDefStatus represents a Subnet creation/modification status.
type SubnetResourcesDefStatus struct {
	IPConfig *IPConfig `json:"ip_config,omitempty"`

	NetworkFunctionChainReference *Reference `json:"network_function_chain_reference,omitempty"`

	SubnetType *string `json:"subnet_type"`

	VlanID *int64 `json:"vlan_id,omitempty"`

	VswitchName *string `json:"vswitch_name,omitempty"`
}

//SubnetDefStatus An intentful representation of a subnet status
type SubnetDefStatus struct {
	AvailabilityZoneReference *Reference `json:"availability_zone_reference,omitempty"`

	ClusterReference *Reference `json:"cluster_reference,omitempty"`

	// A description for subnet.
	Description *string `json:"description,omitempty"`

	// Any error messages for the subnet, if in an error state.
	MessageList []*MessageResource `json:"message_list,omitempty"`

	// subnet Name.
	Name *string `json:"name"`

	Resources *SubnetResourcesDefStatus `json:"resources,omitempty"`

	// The state of the subnet.
	State *string `json:"state,omitempty"`
}

// SubnetIntentResponse represents the response object for intentful operations on a subnet
type SubnetIntentResponse struct {
	APIVersion *string `json:"api_version"`

	Metadata *SubnetMetadata `json:"metadata,omitempty"`

	Spec *Subnet `json:"spec,omitempty"`

	Status *SubnetDefStatus `json:"status,omitempty"`
}

//SubnetIntentResource represents Response object for intentful operations on a subnet
type SubnetIntentResource struct {
	APIVersion *string `json:"api_version,omitempty"`

	Metadata *SubnetMetadata `json:"metadata"`

	Spec *Subnet `json:"spec,omitempty"`

	Status *SubnetDefStatus `json:"status,omitempty"`
}

//SubnetListMetadataOutput All api calls that return a list will have this metadata block
type SubnetListMetadataOutput struct {

	// The filter used for the results
	Filter *string `json:"filter,omitempty"`

	// The kind name
	Kind *string `json:"kind,omitempty"`

	// The number of records retrieved relative to the offset
	Length *int64 `json:"length,omitempty"`

	// Offset from the start of the entity list
	Offset *int64 `json:"offset,omitempty"`

	// The attribute to perform sort on
	SortAttribute *string `json:"sort_attribute,omitempty"`

	// The sort order in which results are returned
	SortOrder *string `json:"sort_order,omitempty"`

	// Total matches found
	TotalMatches *int64 `json:"total_matches,omitempty"`
}

//SubnetListIntentResponse represents the response object for intentful operation of subnets
type SubnetListIntentResponse struct {
	APIVersion *string `json:"api_version"`

	Entities []*SubnetIntentResource `json:"entities,omitempty"`

	Metadata *SubnetListMetadataOutput `json:"metadata"`
}

//SubnetListMetadata ...
type SubnetListMetadata struct {

	// The filter in FIQL syntax used for the results.
	Filter *string `json:"filter,omitempty"`

	// The kind name
	Kind *string `json:"kind,omitempty"`

	// The number of records to retrieve relative to the offset
	Length *int64 `json:"length,omitempty"`

	// Offset from the start of the entity list
	Offset *int64 `json:"offset,omitempty"`

	// The attribute to perform sort on
	SortAttribute *string `json:"sort_attribute,omitempty"`

	// The sort order in which results are returned
	SortOrder *string `json:"sort_order,omitempty"`
}

//Checksum represents the image checksum
type Checksum struct {
	ChecksumAlgorithm *string `json:"checksum_algorithm"`
	ChecksumValue     *string `json:"checksum_value"`
}

//ImageVersionResources The image version, which is composed of a product name and product version.
type ImageVersionResources struct {

	// Name of the producer/distribution of the image. For example windows or red hat.
	ProductName *string `json:"product_name"`

	// Version *string for the disk image.
	ProductVersion *string `json:"product_version"`
}

//ImageResources describes the image spec resources object.
type ImageResources struct {

	// The supported CPU architecture for a disk image.
	Architecture *string `json:"architecture,omitempty"`

	// Checksum of the image. The checksum is used for image validation if the image has a source specified. For images that do not have their source specified the checksum is generated by the image service.
	Checksum *Checksum `json:"checksum,omitempty"`

	// The type of image.
	ImageType *string `json:"image_type,omitempty"`

	// The source URI points at the location of a the source image which is used to create/update image.
	SourceURI *string `json:"source_uri,omitempty"`

	// The image version
	Version *ImageVersionResources `json:"version,omitempty"`
}

//Image An intentful representation of a image spec
type Image struct {

	// A description for image.
	Description *string `json:"description,omitempty"`

	// image Name.
	Name *string `json:"name,omitempty"`

	Resources *ImageResources `json:"resources"`
}

//ImageMetadata Metadata The image kind metadata
type ImageMetadata struct {

	// Categories for the image
	Categories map[string]string `json:"categories,omitempty"`

	// UTC date and time in RFC-3339 format when vm was created
	CreationTime *time.Time `json:"creation_time,omitempty"`

	// The kind name
	Kind *string `json:"kind"`

	// UTC date and time in RFC-3339 format when image was last updated
	LastUpdateTime *time.Time `json:"last_update_time,omitempty"`

	// image name
	Name *string `json:"name,omitempty"`

	// project reference
	ProjectReference *Reference `json:"project_reference,omitempty"`

	OwnerReference *Reference `json:"owner_reference,omitempty"`

	// Hash of the spec. This will be returned from server.
	SpecHash *string `json:"spec_hash,omitempty"`

	// Version number of the latest spec.
	SpecVersion *int64 `json:"spec_version,omitempty"`

	// image uuid
	UUID *string `json:"uuid,omitempty"`
}

//ImageIntentInput An intentful representation of a image
type ImageIntentInput struct {
	APIVersion *string `json:"api_version,omitempty"`

	Metadata *ImageMetadata `json:"metadata,omitempty"`

	Spec *Image `json:"spec,omitempty"`
}

//ImageStatus represents the status of a REST API call. Only used when there is a failure to report.
type ImageStatus struct {
	APIVersion *string `json:"api_version,omitempty"`

	// The HTTP error code.
	Code *int64 `json:"code,omitempty"`

	// The kind name
	Kind *string `json:"kind,omitempty"`

	MessageList []*MessageResource `json:"message_list,omitempty"`

	State *string `json:"state,omitempty"`
}

//ImageVersionStatus represents the image version, which is composed of a product name and product version.
type ImageVersionStatus struct {

	// Name of the producer/distribution of the image. For example windows or red hat.
	ProductName *string `json:"product_name"`

	// Version *string for the disk image.
	ProductVersion *string `json:"product_version"`
}

//ImageResourcesDefStatus describes the image status resources object.
type ImageResourcesDefStatus struct {

	// The supported CPU architecture for a disk image.
	Architecture *string `json:"architecture,omitempty"`

	// Checksum of the image. The checksum is used for image validation if the image has a source specified. For images that do not have their source specified the checksum is generated by the image service.
	Checksum *Checksum `json:"checksum,omitempty"`

	// The type of image.
	ImageType *string `json:"image_type,omitempty"`

	// List of URIs where the raw image data can be accessed.
	RetrievalURIList []*string `json:"retrieval_uri_list,omitempty"`

	// The size of the image in bytes.
	SizeBytes *int64 `json:"size_bytes,omitempty"`

	// The source URI points at the location of a the source image which is used to create/update image.
	SourceURI *string `json:"source_uri,omitempty"`

	// The image version
	Version *ImageVersionStatus `json:"version,omitempty"`
}

//ImageDefStatus represents an intentful representation of a image status
type ImageDefStatus struct {
	AvailabilityZoneReference *Reference `json:"availability_zone_reference,omitempty"`

	ClusterReference *Reference `json:"cluster_reference,omitempty"`

	// A description for image.
	Description *string `json:"description,omitempty"`

	// Any error messages for the image, if in an error state.
	MessageList []*MessageResource `json:"message_list,omitempty"`

	// image Name.
	Name *string `json:"name"`

	Resources ImageResourcesDefStatus `json:"resources"`

	// The state of the image.
	State *string `json:"state,omitempty"`
}

//ImageIntentResponse represents the response object for intentful operations on a image
type ImageIntentResponse struct {
	APIVersion *string `json:"api_version"`

	Metadata *ImageMetadata `json:"metadata"`

	Spec *Image `json:"spec,omitempty"`

	Status *ImageDefStatus `json:"status,omitempty"`
}

//ImageListMetadata represents metadata input
type ImageListMetadata struct {

	// The filter in FIQL syntax used for the results.
	Filter *string `json:"filter,omitempty"`

	// The kind name
	Kind *string `json:"kind,omitempty"`

	// The number of records to retrieve relative to the offset
	Length *int64 `json:"length,omitempty"`

	// Offset from the start of the entity list
	Offset *int64 `json:"offset,omitempty"`

	// The attribute to perform sort on
	SortAttribute *string `json:"sort_attribute,omitempty"`

	// The sort order in which results are returned
	SortOrder *string `json:"sort_order,omitempty"`
}

//ImageIntentResource represents the response object for intentful operations on a image
type ImageIntentResource struct {
	APIVersion *string `json:"api_version,omitempty"`

	Metadata *ImageMetadata `json:"metadata"`

	Spec *Image `json:"spec,omitempty"`

	Status *ImageDefStatus `json:"status,omitempty"`
}

//ImageListMetadataOutput represents metadata block in image list
type ImageListMetadataOutput struct {

	// The filter used for the results
	Filter *string `json:"filter,omitempty"`

	// The kind name
	Kind *string `json:"kind,omitempty"`

	// The number of records retrieved relative to the offset
	Length *int64 `json:"length,omitempty"`

	// Offset from the start of the entity list
	Offset *int64 `json:"offset,omitempty"`

	// The attribute to perform sort on
	SortAttribute *string `json:"sort_attribute,omitempty"`

	// The sort order in which results are returned
	SortOrder *string `json:"sort_order,omitempty"`

	// Total matches found
	TotalMatches *int64 `json:"total_matches,omitempty"`
}

//ImageListIntentResponse represents the response object for intentful operation of images
type ImageListIntentResponse struct {
	APIVersion *string `json:"api_version"`

	Entities []*ImageIntentResource `json:"entities,omitempty"`

	Metadata *ImageListMetadataOutput `json:"metadata"`
}
