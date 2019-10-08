/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by apiregister-gen. Do not edit it manually!

package vmoperator

import (
	"context"
	"fmt"

	"github.com/kubernetes-incubator/apiserver-builder-alpha/pkg/builders"
	corev1 "k8s.io/api/core/v1"
	apiresource "k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/registry/rest"
)

var (
	InternalVirtualMachine = builders.NewInternalResource(
		"virtualmachines",
		"VirtualMachine",
		func() runtime.Object { return &VirtualMachine{} },
		func() runtime.Object { return &VirtualMachineList{} },
	)
	InternalVirtualMachineStatus = builders.NewInternalResourceStatus(
		"virtualmachines",
		"VirtualMachineStatus",
		func() runtime.Object { return &VirtualMachine{} },
		func() runtime.Object { return &VirtualMachineList{} },
	)
	InternalVirtualMachineClass = builders.NewInternalResource(
		"virtualmachineclasses",
		"VirtualMachineClass",
		func() runtime.Object { return &VirtualMachineClass{} },
		func() runtime.Object { return &VirtualMachineClassList{} },
	)
	InternalVirtualMachineClassStatus = builders.NewInternalResourceStatus(
		"virtualmachineclasses",
		"VirtualMachineClassStatus",
		func() runtime.Object { return &VirtualMachineClass{} },
		func() runtime.Object { return &VirtualMachineClassList{} },
	)
	InternalVirtualMachineImage = builders.NewInternalResource(
		"virtualmachineimages",
		"VirtualMachineImage",
		func() runtime.Object { return &VirtualMachineImage{} },
		func() runtime.Object { return &VirtualMachineImageList{} },
	)
	InternalVirtualMachineImageStatus = builders.NewInternalResourceStatus(
		"virtualmachineimages",
		"VirtualMachineImageStatus",
		func() runtime.Object { return &VirtualMachineImage{} },
		func() runtime.Object { return &VirtualMachineImageList{} },
	)
	InternalVirtualMachineService = builders.NewInternalResource(
		"virtualmachineservices",
		"VirtualMachineService",
		func() runtime.Object { return &VirtualMachineService{} },
		func() runtime.Object { return &VirtualMachineServiceList{} },
	)
	InternalVirtualMachineServiceStatus = builders.NewInternalResourceStatus(
		"virtualmachineservices",
		"VirtualMachineServiceStatus",
		func() runtime.Object { return &VirtualMachineService{} },
		func() runtime.Object { return &VirtualMachineServiceList{} },
	)
	InternalVirtualMachineSetResourcePolicy = builders.NewInternalResource(
		"virtualmachinesetresourcepolicies",
		"VirtualMachineSetResourcePolicy",
		func() runtime.Object { return &VirtualMachineSetResourcePolicy{} },
		func() runtime.Object { return &VirtualMachineSetResourcePolicyList{} },
	)
	InternalVirtualMachineSetResourcePolicyStatus = builders.NewInternalResourceStatus(
		"virtualmachinesetresourcepolicies",
		"VirtualMachineSetResourcePolicyStatus",
		func() runtime.Object { return &VirtualMachineSetResourcePolicy{} },
		func() runtime.Object { return &VirtualMachineSetResourcePolicyList{} },
	)
	// Registered resources and subresources
	ApiVersion = builders.NewApiGroup("vmoperator.vmware.com").WithKinds(
		InternalVirtualMachine,
		InternalVirtualMachineStatus,
		InternalVirtualMachineClass,
		InternalVirtualMachineClassStatus,
		InternalVirtualMachineImage,
		InternalVirtualMachineImageStatus,
		InternalVirtualMachineService,
		InternalVirtualMachineServiceStatus,
		InternalVirtualMachineSetResourcePolicy,
		InternalVirtualMachineSetResourcePolicyStatus,
	)

	// Required by code generated by go2idl
	AddToScheme        = ApiVersion.SchemaBuilder.AddToScheme
	SchemeBuilder      = ApiVersion.SchemaBuilder
	localSchemeBuilder = &SchemeBuilder
	SchemeGroupVersion = ApiVersion.GroupVersion
)

// Required by code generated by go2idl
// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Required by code generated by go2idl
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

type Protocol string
type VMStatusPhase string
type VirtualMachineServiceType string

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VirtualMachine struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   VirtualMachineSpec
	Status VirtualMachineStatus
}

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VirtualMachineSetResourcePolicy struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   VirtualMachineSetResourcePolicySpec
	Status VirtualMachineSetResourcePolicyStatus
}

type VirtualMachineStatus struct {
	Conditions         []VirtualMachineCondition
	Host               string
	PowerState         string
	ResourcePolicyName string
	Phase              VMStatusPhase
	VmIp               string
	BiosUuid           string
	Volumes            []VirtualMachineVolumeStatus
}

type VirtualMachineSetResourcePolicyStatus struct {
}

type VirtualMachineVolumeStatus struct {
	Name     string
	Attached bool
	DiskUuid string
	Error    string
}

type VirtualMachineCondition struct {
	LastProbeTime      metav1.Time
	LastTransitionTime metav1.Time
	Message            string
	Reason             string
	Status             string
	Type               string
}

type VirtualMachineSetResourcePolicySpec struct {
	ResourcePool ResourcePoolSpec
	Folder       FolderSpec
}

type VirtualMachineSpec struct {
	ImageName          string
	ClassName          string
	PowerState         string
	Env                corev1.EnvVar
	Ports              []VirtualMachinePort
	VmMetadata         *VirtualMachineMetadata
	StorageClass       string
	NetworkInterfaces  []VirtualMachineNetworkInterface
	ResourcePolicyName string
	Volumes            []VirtualMachineVolumes
}

type FolderSpec struct {
	Name string
}

type VirtualMachineVolumes struct {
	Name                  string
	PersistentVolumeClaim *corev1.PersistentVolumeClaimVolumeSource
}

type VirtualMachineNetworkInterface struct {
	NetworkName      string
	EthernetCardType string
	NetworkType      string
}

type VirtualMachineMetadata struct {
	ConfigMapName string
	Transport     string
}

type VirtualMachinePort struct {
	Port     int
	Ip       string
	Name     string
	Protocol corev1.Protocol
}

type ResourcePoolSpec struct {
	Name         string
	Reservations VirtualMachineResourceSpec
	Limits       VirtualMachineResourceSpec
}

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VirtualMachineService struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   VirtualMachineServiceSpec
	Status VirtualMachineServiceStatus
}

type VirtualMachineResourceSpec struct {
	Cpu    apiresource.Quantity
	Memory apiresource.Quantity
}

type VirtualMachineServiceStatus struct {
	LoadBalancer LoadBalancerStatus
}

type VirtualMachineServiceSpec struct {
	Type         VirtualMachineServiceType
	Ports        []VirtualMachineServicePort
	Selector     map[string]string
	ClusterIP    string
	ExternalName string
}

type LoadBalancerStatus struct {
	Ingress []LoadBalancerIngress
}

type VirtualMachineServicePort struct {
	Name       string
	Protocol   string
	Port       int32
	TargetPort int32
}

type LoadBalancerIngress struct {
	IP       string
	Hostname string
}

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VirtualMachineImage struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   VirtualMachineImageSpec
	Status VirtualMachineImageStatus
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VirtualMachineClass struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   VirtualMachineClassSpec
	Status VirtualMachineClassStatus
}

type VirtualMachineImageStatus struct {
	Uuid       string
	InternalId string
	PowerState string
}

type VirtualMachineClassStatus struct {
}

type VirtualMachineClassSpec struct {
	Hardware VirtualMachineClassHardware
	Policies VirtualMachineClassPolicies
}

type VirtualMachineImageSpec struct {
	Type            string
	ImageSource     string
	ImageSourceType string
	ImagePath       string
}

type VirtualMachineClassPolicies struct {
	Resources VirtualMachineClassResources
}

type VirtualMachineClassHardware struct {
	Cpus   int64
	Memory apiresource.Quantity
}

type VirtualMachineClassResources struct {
	Requests VirtualMachineResourceSpec
	Limits   VirtualMachineResourceSpec
}

//
// VirtualMachine Functions and Structs
//
// +k8s:deepcopy-gen=false
type VirtualMachineStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type VirtualMachineStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VirtualMachineList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []VirtualMachine
}

func (VirtualMachine) NewStatus() interface{} {
	return VirtualMachineStatus{}
}

func (pc *VirtualMachine) GetStatus() interface{} {
	return pc.Status
}

func (pc *VirtualMachine) SetStatus(s interface{}) {
	pc.Status = s.(VirtualMachineStatus)
}

func (pc *VirtualMachine) GetSpec() interface{} {
	return pc.Spec
}

func (pc *VirtualMachine) SetSpec(s interface{}) {
	pc.Spec = s.(VirtualMachineSpec)
}

func (pc *VirtualMachine) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *VirtualMachine) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc VirtualMachine) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store VirtualMachine.
// +k8s:deepcopy-gen=false
type VirtualMachineRegistry interface {
	ListVirtualMachines(ctx context.Context, options *internalversion.ListOptions) (*VirtualMachineList, error)
	GetVirtualMachine(ctx context.Context, id string, options *metav1.GetOptions) (*VirtualMachine, error)
	CreateVirtualMachine(ctx context.Context, id *VirtualMachine) (*VirtualMachine, error)
	UpdateVirtualMachine(ctx context.Context, id *VirtualMachine) (*VirtualMachine, error)
	DeleteVirtualMachine(ctx context.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewVirtualMachineRegistry(sp builders.StandardStorageProvider) VirtualMachineRegistry {
	return &storageVirtualMachine{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageVirtualMachine struct {
	builders.StandardStorageProvider
}

func (s *storageVirtualMachine) ListVirtualMachines(ctx context.Context, options *internalversion.ListOptions) (*VirtualMachineList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachineList), err
}

func (s *storageVirtualMachine) GetVirtualMachine(ctx context.Context, id string, options *metav1.GetOptions) (*VirtualMachine, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachine), nil
}

func (s *storageVirtualMachine) CreateVirtualMachine(ctx context.Context, object *VirtualMachine) (*VirtualMachine, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, &metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachine), nil
}

func (s *storageVirtualMachine) UpdateVirtualMachine(ctx context.Context, object *VirtualMachine) (*VirtualMachine, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil, false, &metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachine), nil
}

func (s *storageVirtualMachine) DeleteVirtualMachine(ctx context.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, &metav1.DeleteOptions{})
	return sync, err
}

//
// VirtualMachineClass Functions and Structs
//
// +k8s:deepcopy-gen=false
type VirtualMachineClassStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type VirtualMachineClassStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VirtualMachineClassList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []VirtualMachineClass
}

func (VirtualMachineClass) NewStatus() interface{} {
	return VirtualMachineClassStatus{}
}

func (pc *VirtualMachineClass) GetStatus() interface{} {
	return pc.Status
}

func (pc *VirtualMachineClass) SetStatus(s interface{}) {
	pc.Status = s.(VirtualMachineClassStatus)
}

func (pc *VirtualMachineClass) GetSpec() interface{} {
	return pc.Spec
}

func (pc *VirtualMachineClass) SetSpec(s interface{}) {
	pc.Spec = s.(VirtualMachineClassSpec)
}

func (pc *VirtualMachineClass) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *VirtualMachineClass) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc VirtualMachineClass) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store VirtualMachineClass.
// +k8s:deepcopy-gen=false
type VirtualMachineClassRegistry interface {
	ListVirtualMachineClasss(ctx context.Context, options *internalversion.ListOptions) (*VirtualMachineClassList, error)
	GetVirtualMachineClass(ctx context.Context, id string, options *metav1.GetOptions) (*VirtualMachineClass, error)
	CreateVirtualMachineClass(ctx context.Context, id *VirtualMachineClass) (*VirtualMachineClass, error)
	UpdateVirtualMachineClass(ctx context.Context, id *VirtualMachineClass) (*VirtualMachineClass, error)
	DeleteVirtualMachineClass(ctx context.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewVirtualMachineClassRegistry(sp builders.StandardStorageProvider) VirtualMachineClassRegistry {
	return &storageVirtualMachineClass{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageVirtualMachineClass struct {
	builders.StandardStorageProvider
}

func (s *storageVirtualMachineClass) ListVirtualMachineClasss(ctx context.Context, options *internalversion.ListOptions) (*VirtualMachineClassList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachineClassList), err
}

func (s *storageVirtualMachineClass) GetVirtualMachineClass(ctx context.Context, id string, options *metav1.GetOptions) (*VirtualMachineClass, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachineClass), nil
}

func (s *storageVirtualMachineClass) CreateVirtualMachineClass(ctx context.Context, object *VirtualMachineClass) (*VirtualMachineClass, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, &metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachineClass), nil
}

func (s *storageVirtualMachineClass) UpdateVirtualMachineClass(ctx context.Context, object *VirtualMachineClass) (*VirtualMachineClass, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil, false, &metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachineClass), nil
}

func (s *storageVirtualMachineClass) DeleteVirtualMachineClass(ctx context.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, &metav1.DeleteOptions{})
	return sync, err
}

//
// VirtualMachineImage Functions and Structs
//
// +k8s:deepcopy-gen=false
type VirtualMachineImageStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type VirtualMachineImageStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VirtualMachineImageList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []VirtualMachineImage
}

func (VirtualMachineImage) NewStatus() interface{} {
	return VirtualMachineImageStatus{}
}

func (pc *VirtualMachineImage) GetStatus() interface{} {
	return pc.Status
}

func (pc *VirtualMachineImage) SetStatus(s interface{}) {
	pc.Status = s.(VirtualMachineImageStatus)
}

func (pc *VirtualMachineImage) GetSpec() interface{} {
	return pc.Spec
}

func (pc *VirtualMachineImage) SetSpec(s interface{}) {
	pc.Spec = s.(VirtualMachineImageSpec)
}

func (pc *VirtualMachineImage) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *VirtualMachineImage) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc VirtualMachineImage) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store VirtualMachineImage.
// +k8s:deepcopy-gen=false
type VirtualMachineImageRegistry interface {
	ListVirtualMachineImages(ctx context.Context, options *internalversion.ListOptions) (*VirtualMachineImageList, error)
	GetVirtualMachineImage(ctx context.Context, id string, options *metav1.GetOptions) (*VirtualMachineImage, error)
	CreateVirtualMachineImage(ctx context.Context, id *VirtualMachineImage) (*VirtualMachineImage, error)
	UpdateVirtualMachineImage(ctx context.Context, id *VirtualMachineImage) (*VirtualMachineImage, error)
	DeleteVirtualMachineImage(ctx context.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewVirtualMachineImageRegistry(sp builders.StandardStorageProvider) VirtualMachineImageRegistry {
	return &storageVirtualMachineImage{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageVirtualMachineImage struct {
	builders.StandardStorageProvider
}

func (s *storageVirtualMachineImage) ListVirtualMachineImages(ctx context.Context, options *internalversion.ListOptions) (*VirtualMachineImageList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachineImageList), err
}

func (s *storageVirtualMachineImage) GetVirtualMachineImage(ctx context.Context, id string, options *metav1.GetOptions) (*VirtualMachineImage, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachineImage), nil
}

func (s *storageVirtualMachineImage) CreateVirtualMachineImage(ctx context.Context, object *VirtualMachineImage) (*VirtualMachineImage, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, &metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachineImage), nil
}

func (s *storageVirtualMachineImage) UpdateVirtualMachineImage(ctx context.Context, object *VirtualMachineImage) (*VirtualMachineImage, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil, false, &metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachineImage), nil
}

func (s *storageVirtualMachineImage) DeleteVirtualMachineImage(ctx context.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, &metav1.DeleteOptions{})
	return sync, err
}

//
// VirtualMachineService Functions and Structs
//
// +k8s:deepcopy-gen=false
type VirtualMachineServiceStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type VirtualMachineServiceStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VirtualMachineServiceList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []VirtualMachineService
}

func (VirtualMachineService) NewStatus() interface{} {
	return VirtualMachineServiceStatus{}
}

func (pc *VirtualMachineService) GetStatus() interface{} {
	return pc.Status
}

func (pc *VirtualMachineService) SetStatus(s interface{}) {
	pc.Status = s.(VirtualMachineServiceStatus)
}

func (pc *VirtualMachineService) GetSpec() interface{} {
	return pc.Spec
}

func (pc *VirtualMachineService) SetSpec(s interface{}) {
	pc.Spec = s.(VirtualMachineServiceSpec)
}

func (pc *VirtualMachineService) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *VirtualMachineService) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc VirtualMachineService) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store VirtualMachineService.
// +k8s:deepcopy-gen=false
type VirtualMachineServiceRegistry interface {
	ListVirtualMachineServices(ctx context.Context, options *internalversion.ListOptions) (*VirtualMachineServiceList, error)
	GetVirtualMachineService(ctx context.Context, id string, options *metav1.GetOptions) (*VirtualMachineService, error)
	CreateVirtualMachineService(ctx context.Context, id *VirtualMachineService) (*VirtualMachineService, error)
	UpdateVirtualMachineService(ctx context.Context, id *VirtualMachineService) (*VirtualMachineService, error)
	DeleteVirtualMachineService(ctx context.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewVirtualMachineServiceRegistry(sp builders.StandardStorageProvider) VirtualMachineServiceRegistry {
	return &storageVirtualMachineService{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageVirtualMachineService struct {
	builders.StandardStorageProvider
}

func (s *storageVirtualMachineService) ListVirtualMachineServices(ctx context.Context, options *internalversion.ListOptions) (*VirtualMachineServiceList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachineServiceList), err
}

func (s *storageVirtualMachineService) GetVirtualMachineService(ctx context.Context, id string, options *metav1.GetOptions) (*VirtualMachineService, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachineService), nil
}

func (s *storageVirtualMachineService) CreateVirtualMachineService(ctx context.Context, object *VirtualMachineService) (*VirtualMachineService, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, &metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachineService), nil
}

func (s *storageVirtualMachineService) UpdateVirtualMachineService(ctx context.Context, object *VirtualMachineService) (*VirtualMachineService, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil, false, &metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachineService), nil
}

func (s *storageVirtualMachineService) DeleteVirtualMachineService(ctx context.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, &metav1.DeleteOptions{})
	return sync, err
}

//
// VirtualMachineSetResourcePolicy Functions and Structs
//
// +k8s:deepcopy-gen=false
type VirtualMachineSetResourcePolicyStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type VirtualMachineSetResourcePolicyStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type VirtualMachineSetResourcePolicyList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []VirtualMachineSetResourcePolicy
}

func (VirtualMachineSetResourcePolicy) NewStatus() interface{} {
	return VirtualMachineSetResourcePolicyStatus{}
}

func (pc *VirtualMachineSetResourcePolicy) GetStatus() interface{} {
	return pc.Status
}

func (pc *VirtualMachineSetResourcePolicy) SetStatus(s interface{}) {
	pc.Status = s.(VirtualMachineSetResourcePolicyStatus)
}

func (pc *VirtualMachineSetResourcePolicy) GetSpec() interface{} {
	return pc.Spec
}

func (pc *VirtualMachineSetResourcePolicy) SetSpec(s interface{}) {
	pc.Spec = s.(VirtualMachineSetResourcePolicySpec)
}

func (pc *VirtualMachineSetResourcePolicy) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *VirtualMachineSetResourcePolicy) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc VirtualMachineSetResourcePolicy) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store VirtualMachineSetResourcePolicy.
// +k8s:deepcopy-gen=false
type VirtualMachineSetResourcePolicyRegistry interface {
	ListVirtualMachineSetResourcePolicys(ctx context.Context, options *internalversion.ListOptions) (*VirtualMachineSetResourcePolicyList, error)
	GetVirtualMachineSetResourcePolicy(ctx context.Context, id string, options *metav1.GetOptions) (*VirtualMachineSetResourcePolicy, error)
	CreateVirtualMachineSetResourcePolicy(ctx context.Context, id *VirtualMachineSetResourcePolicy) (*VirtualMachineSetResourcePolicy, error)
	UpdateVirtualMachineSetResourcePolicy(ctx context.Context, id *VirtualMachineSetResourcePolicy) (*VirtualMachineSetResourcePolicy, error)
	DeleteVirtualMachineSetResourcePolicy(ctx context.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewVirtualMachineSetResourcePolicyRegistry(sp builders.StandardStorageProvider) VirtualMachineSetResourcePolicyRegistry {
	return &storageVirtualMachineSetResourcePolicy{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageVirtualMachineSetResourcePolicy struct {
	builders.StandardStorageProvider
}

func (s *storageVirtualMachineSetResourcePolicy) ListVirtualMachineSetResourcePolicys(ctx context.Context, options *internalversion.ListOptions) (*VirtualMachineSetResourcePolicyList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachineSetResourcePolicyList), err
}

func (s *storageVirtualMachineSetResourcePolicy) GetVirtualMachineSetResourcePolicy(ctx context.Context, id string, options *metav1.GetOptions) (*VirtualMachineSetResourcePolicy, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachineSetResourcePolicy), nil
}

func (s *storageVirtualMachineSetResourcePolicy) CreateVirtualMachineSetResourcePolicy(ctx context.Context, object *VirtualMachineSetResourcePolicy) (*VirtualMachineSetResourcePolicy, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, &metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachineSetResourcePolicy), nil
}

func (s *storageVirtualMachineSetResourcePolicy) UpdateVirtualMachineSetResourcePolicy(ctx context.Context, object *VirtualMachineSetResourcePolicy) (*VirtualMachineSetResourcePolicy, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil, false, &metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return obj.(*VirtualMachineSetResourcePolicy), nil
}

func (s *storageVirtualMachineSetResourcePolicy) DeleteVirtualMachineSetResourcePolicy(ctx context.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, &metav1.DeleteOptions{})
	return sync, err
}
