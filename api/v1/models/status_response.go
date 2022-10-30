// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StatusResponse Health and status information of daemon
//
// +k8s:deepcopy-gen=true
//
// swagger:model StatusResponse
type StatusResponse struct {

	// Status of the SRv6 support
	Srv6 *Srv6 `json:"Srv6,omitempty"`

	// Status of bandwidth manager
	BandwidthManager *BandwidthManager `json:"bandwidth-manager,omitempty"`

	// Status of BPF maps
	BpfMaps *BPFMapStatus `json:"bpf-maps,omitempty"`

	// Status of Cilium daemon
	Cilium *Status `json:"cilium,omitempty"`

	// When supported by the API, this client ID should be used by the
	// client when making another request to the server.
	// See for example "/cluster/nodes".
	//
	ClientID int64 `json:"client-id,omitempty"`

	// Status of clock source
	ClockSource *ClockSource `json:"clock-source,omitempty"`

	// Status of cluster
	Cluster *ClusterStatus `json:"cluster,omitempty"`

	// Status of ClusterMesh
	ClusterMesh *ClusterMeshStatus `json:"cluster-mesh,omitempty"`

	// Status of CNI chaining
	CniChaining *CNIChainingStatus `json:"cni-chaining,omitempty"`

	// Status of the CNI configuration file
	CniFile *Status `json:"cni-file,omitempty"`

	// Status of local container runtime
	ContainerRuntime *Status `json:"container-runtime,omitempty"`

	// Status of all endpoint controllers
	Controllers ControllerStatuses `json:"controllers,omitempty"`

	// Status of transparent encryption
	Encryption *EncryptionStatus `json:"encryption,omitempty"`

	// Status of the host firewall
	HostFirewall *HostFirewall `json:"host-firewall,omitempty"`

	// Status of host routing
	HostRouting *HostRouting `json:"host-routing,omitempty"`

	// Status of Hubble server
	Hubble *HubbleStatus `json:"hubble,omitempty"`

	// Status of identity range of the cluster
	IdentityRange *IdentityRange `json:"identity-range,omitempty"`

	// Status of IP address management
	Ipam *IPAMStatus `json:"ipam,omitempty"`

	// Status of IPv6 BIG TCP
	IPV6BigTCP *IPV6BigTCP `json:"ipv6-big-tcp,omitempty"`

	// Status of kube-proxy replacement
	KubeProxyReplacement *KubeProxyReplacement `json:"kube-proxy-replacement,omitempty"`

	// Status of Kubernetes integration
	Kubernetes *K8sStatus `json:"kubernetes,omitempty"`

	// Status of key/value datastore
	Kvstore *Status `json:"kvstore,omitempty"`

	// Status of masquerading
	Masquerading *Masquerading `json:"masquerading,omitempty"`

	// Status of the node monitor
	NodeMonitor *MonitorStatus `json:"nodeMonitor,omitempty"`

	// Status of proxy
	Proxy *ProxyStatus `json:"proxy,omitempty"`

	// List of stale information in the status
	Stale map[string]strfmt.DateTime `json:"stale,omitempty"`
}

// Validate validates this status response
func (m *StatusResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSrv6(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBandwidthManager(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBpfMaps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCilium(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClockSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterMesh(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCniChaining(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCniFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainerRuntime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateControllers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncryption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostFirewall(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostRouting(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHubble(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentityRange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIpam(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPV6BigTCP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubeProxyReplacement(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubernetes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKvstore(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMasquerading(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeMonitor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStale(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatusResponse) validateSrv6(formats strfmt.Registry) error {

	if swag.IsZero(m.Srv6) { // not required
		return nil
	}

	if m.Srv6 != nil {
		if err := m.Srv6.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Srv6")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateBandwidthManager(formats strfmt.Registry) error {

	if swag.IsZero(m.BandwidthManager) { // not required
		return nil
	}

	if m.BandwidthManager != nil {
		if err := m.BandwidthManager.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bandwidth-manager")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateBpfMaps(formats strfmt.Registry) error {

	if swag.IsZero(m.BpfMaps) { // not required
		return nil
	}

	if m.BpfMaps != nil {
		if err := m.BpfMaps.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bpf-maps")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateCilium(formats strfmt.Registry) error {

	if swag.IsZero(m.Cilium) { // not required
		return nil
	}

	if m.Cilium != nil {
		if err := m.Cilium.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cilium")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateClockSource(formats strfmt.Registry) error {

	if swag.IsZero(m.ClockSource) { // not required
		return nil
	}

	if m.ClockSource != nil {
		if err := m.ClockSource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clock-source")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateCluster(formats strfmt.Registry) error {

	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateClusterMesh(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterMesh) { // not required
		return nil
	}

	if m.ClusterMesh != nil {
		if err := m.ClusterMesh.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster-mesh")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateCniChaining(formats strfmt.Registry) error {

	if swag.IsZero(m.CniChaining) { // not required
		return nil
	}

	if m.CniChaining != nil {
		if err := m.CniChaining.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cni-chaining")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateCniFile(formats strfmt.Registry) error {

	if swag.IsZero(m.CniFile) { // not required
		return nil
	}

	if m.CniFile != nil {
		if err := m.CniFile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cni-file")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateContainerRuntime(formats strfmt.Registry) error {

	if swag.IsZero(m.ContainerRuntime) { // not required
		return nil
	}

	if m.ContainerRuntime != nil {
		if err := m.ContainerRuntime.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("container-runtime")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateControllers(formats strfmt.Registry) error {

	if swag.IsZero(m.Controllers) { // not required
		return nil
	}

	if err := m.Controllers.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("controllers")
		}
		return err
	}

	return nil
}

func (m *StatusResponse) validateEncryption(formats strfmt.Registry) error {

	if swag.IsZero(m.Encryption) { // not required
		return nil
	}

	if m.Encryption != nil {
		if err := m.Encryption.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryption")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateHostFirewall(formats strfmt.Registry) error {

	if swag.IsZero(m.HostFirewall) { // not required
		return nil
	}

	if m.HostFirewall != nil {
		if err := m.HostFirewall.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host-firewall")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateHostRouting(formats strfmt.Registry) error {

	if swag.IsZero(m.HostRouting) { // not required
		return nil
	}

	if m.HostRouting != nil {
		if err := m.HostRouting.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host-routing")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateHubble(formats strfmt.Registry) error {

	if swag.IsZero(m.Hubble) { // not required
		return nil
	}

	if m.Hubble != nil {
		if err := m.Hubble.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hubble")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateIdentityRange(formats strfmt.Registry) error {

	if swag.IsZero(m.IdentityRange) { // not required
		return nil
	}

	if m.IdentityRange != nil {
		if err := m.IdentityRange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identity-range")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateIpam(formats strfmt.Registry) error {

	if swag.IsZero(m.Ipam) { // not required
		return nil
	}

	if m.Ipam != nil {
		if err := m.Ipam.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipam")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateIPV6BigTCP(formats strfmt.Registry) error {

	if swag.IsZero(m.IPV6BigTCP) { // not required
		return nil
	}

	if m.IPV6BigTCP != nil {
		if err := m.IPV6BigTCP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipv6-big-tcp")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateKubeProxyReplacement(formats strfmt.Registry) error {

	if swag.IsZero(m.KubeProxyReplacement) { // not required
		return nil
	}

	if m.KubeProxyReplacement != nil {
		if err := m.KubeProxyReplacement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kube-proxy-replacement")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateKubernetes(formats strfmt.Registry) error {

	if swag.IsZero(m.Kubernetes) { // not required
		return nil
	}

	if m.Kubernetes != nil {
		if err := m.Kubernetes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubernetes")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateKvstore(formats strfmt.Registry) error {

	if swag.IsZero(m.Kvstore) { // not required
		return nil
	}

	if m.Kvstore != nil {
		if err := m.Kvstore.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kvstore")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateMasquerading(formats strfmt.Registry) error {

	if swag.IsZero(m.Masquerading) { // not required
		return nil
	}

	if m.Masquerading != nil {
		if err := m.Masquerading.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("masquerading")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateNodeMonitor(formats strfmt.Registry) error {

	if swag.IsZero(m.NodeMonitor) { // not required
		return nil
	}

	if m.NodeMonitor != nil {
		if err := m.NodeMonitor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nodeMonitor")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateProxy(formats strfmt.Registry) error {

	if swag.IsZero(m.Proxy) { // not required
		return nil
	}

	if m.Proxy != nil {
		if err := m.Proxy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proxy")
			}
			return err
		}
	}

	return nil
}

func (m *StatusResponse) validateStale(formats strfmt.Registry) error {

	if swag.IsZero(m.Stale) { // not required
		return nil
	}

	for k := range m.Stale {

		if err := validate.FormatOf("stale"+"."+k, "body", "date-time", m.Stale[k].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StatusResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatusResponse) UnmarshalBinary(b []byte) error {
	var res StatusResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
