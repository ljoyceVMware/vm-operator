// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package validation_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	vmopv1 "github.com/vmware-tanzu/vm-operator/api/v1alpha1"

	imgregv1a1 "github.com/vmware-tanzu/vm-operator/external/image-registry/api/v1alpha1"

	"github.com/vmware-tanzu/vm-operator/pkg/lib"
	"github.com/vmware-tanzu/vm-operator/test/builder"
)

func unitTests() {
	Describe("Invoking ValidateCreate", unitTestsValidateCreate)
	Describe("Invoking ValidateUpdate", unitTestsValidateUpdate)
	Describe("Invoking ValidateDelete", unitTestsValidateDelete)
}

type unitValidatingWebhookContext struct {
	builder.UnitTestContextForValidatingWebhook
	vmPub    *vmopv1.VirtualMachinePublishRequest
	oldVMPub *vmopv1.VirtualMachinePublishRequest
	vm       *vmopv1.VirtualMachine
	cl       *imgregv1a1.ContentLibrary
}

func newUnitTestContextForValidatingWebhook(isUpdate bool) *unitValidatingWebhookContext {
	vm := builder.DummyVirtualMachine()
	vm.Name = "dummy-vm"
	vm.Namespace = "dummy-ns"
	cl := builder.DummyContentLibrary("dummy-cl", "dummy-ns", "dummy-uuid")

	vmPub := builder.DummyVirtualMachinePublishRequest("dummy-vmpub", "dummy-ns", vm.Name,
		"dummy-item", cl.Name)
	obj, err := builder.ToUnstructured(vmPub)
	Expect(err).ToNot(HaveOccurred())

	var oldVMPub *vmopv1.VirtualMachinePublishRequest
	var oldObj *unstructured.Unstructured

	if isUpdate {
		oldVMPub = vmPub.DeepCopy()
		oldObj, err = builder.ToUnstructured(oldVMPub)
		Expect(err).ToNot(HaveOccurred())
	}

	return &unitValidatingWebhookContext{
		UnitTestContextForValidatingWebhook: *suite.NewUnitTestContextForValidatingWebhook(obj, oldObj, vm, cl),
		vmPub:                               vmPub,
		oldVMPub:                            oldVMPub,
		vm:                                  vm,
		cl:                                  cl,
	}
}

func unitTestsValidateCreate() {
	var (
		ctx *unitValidatingWebhookContext
		err error

		invalidAPIVersion = "vmoperator.vmware.com/v1"
	)

	type createArgs struct {
		invalidSourceAPIVersion         bool
		invalidSourceKind               bool
		invalidTargetLocationAPIVersion bool
		invalidTargetLocationKind       bool
		sourceNotFound                  bool
		defaultSourceNotFound           bool
		defaultSourceExists             bool
		targetLocationNotWritable       bool
		targetLocationNameEmpty         bool
		targetLocationNotFound          bool
	}

	validateCreate := func(args createArgs, expectedAllowed bool, expectedReason string, expectedErr error) {
		if args.invalidSourceAPIVersion {
			ctx.vmPub.Spec.Source.APIVersion = invalidAPIVersion
		}

		if args.invalidSourceKind {
			ctx.vmPub.Spec.Source.Kind = "Machine"
		}

		if args.invalidTargetLocationAPIVersion {
			ctx.vmPub.Spec.Target.Location.APIVersion = invalidAPIVersion
		}

		if args.invalidTargetLocationKind {
			ctx.vmPub.Spec.Target.Location.Kind = "ClusterContentLibrary"
		}

		if args.sourceNotFound {
			Expect(ctx.Client.Delete(ctx, ctx.vm)).To(Succeed())
		}

		if args.defaultSourceNotFound {
			ctx.vmPub.Spec.Source.Name = ""
		}

		if args.defaultSourceExists {
			defaultVM := builder.DummyVirtualMachine()
			defaultVM.Name = ctx.vmPub.Name
			defaultVM.Namespace = ctx.vmPub.Namespace
			Expect(ctx.Client.Create(ctx, defaultVM)).To(Succeed())
		}

		if args.targetLocationNotWritable {
			ctx.cl.Spec.Writable = false
			Expect(ctx.Client.Update(ctx, ctx.cl)).To(Succeed())
		}

		if args.targetLocationNameEmpty {
			ctx.vmPub.Spec.Target.Location.Name = ""
		}

		if args.targetLocationNotFound {
			Expect(ctx.Client.Delete(ctx, ctx.cl)).To(Succeed())
		}

		ctx.WebhookRequestContext.Obj, err = builder.ToUnstructured(ctx.vmPub)
		Expect(err).ToNot(HaveOccurred())

		response := ctx.ValidateCreate(&ctx.WebhookRequestContext)
		Expect(response.Allowed).To(Equal(expectedAllowed))
		if expectedReason != "" {
			Expect(string(response.Result.Reason)).To(ContainSubstring(expectedReason))
		}
		if expectedErr != nil {
			Expect(response.Result.Message).To(Equal(expectedErr.Error()))
		}
	}

	BeforeEach(func() {
		ctx = newUnitTestContextForValidatingWebhook(false)
		lib.IsWCPVMImageRegistryEnabled = func() bool {
			return true
		}
	})

	AfterEach(func() {
		ctx = nil
	})

	sourcePath := field.NewPath("spec").Child("source")
	targetLocationPath := field.NewPath("spec").Child("target", "location")
	DescribeTable("create table", validateCreate,
		Entry("should allow valid", createArgs{}, true, nil, nil),
		Entry("should allow if default source VM exists", createArgs{defaultSourceExists: true}, true, nil, nil),
		Entry("should deny invalid source API version", createArgs{invalidSourceAPIVersion: true}, false,
			field.NotSupported(sourcePath.Child("apiVersion"), invalidAPIVersion,
				[]string{"vmoperator.vmware.com/v1alpha1", ""}).Error(), nil),
		Entry("should deny invalid source kind", createArgs{invalidSourceKind: true}, false,
			field.NotSupported(sourcePath.Child("kind"), "Machine",
				[]string{"VirtualMachine", ""}).Error(), nil),
		Entry("should deny invalid target location API version", createArgs{invalidTargetLocationAPIVersion: true}, false,
			field.NotSupported(targetLocationPath.Child("apiVersion"), invalidAPIVersion,
				[]string{"imageregistry.vmware.com/v1alpha1", ""}).Error(), nil),
		Entry("should deny invalid target location kind", createArgs{invalidTargetLocationKind: true}, false,
			field.NotSupported(targetLocationPath.Child("kind"), "ClusterContentLibrary",
				[]string{"ContentLibrary", ""}).Error(), nil),
		Entry("should deny if source VM not found", createArgs{sourceNotFound: true}, false,
			field.NotFound(sourcePath.Child("name"), "dummy-vm").Error(), nil),
		Entry("should deny if default source VM not found", createArgs{defaultSourceNotFound: true}, false,
			field.Invalid(sourcePath.Child("name"), "", "failed to get the default source VM").Error(), nil),
		Entry("should deny if target location is not writable", createArgs{targetLocationNotWritable: true}, false,
			field.Invalid(targetLocationPath.Child("name"), "dummy-cl",
				"target location dummy-cl is not writable").Error(), nil),
		Entry("should deny if target location name is empty", createArgs{targetLocationNameEmpty: true}, false,
			field.Required(targetLocationPath.Child("name"), "").Error(), nil),
		Entry("should deny if target location not found", createArgs{targetLocationNotFound: true}, false,
			field.NotFound(targetLocationPath.Child("name"), "dummy-cl").Error(), nil),
	)
}

func unitTestsValidateUpdate() {
	var (
		ctx      *unitValidatingWebhookContext
		response admission.Response
	)

	BeforeEach(func() {
		ctx = newUnitTestContextForValidatingWebhook(true)
	})

	AfterEach(func() {
		ctx = nil
	})

	JustBeforeEach(func() {
		response = ctx.ValidateUpdate(&ctx.WebhookRequestContext)
	})

	Context("Source/Target is updated", func() {
		var err error

		BeforeEach(func() {
			ctx.vmPub.Spec.Source.Name = "updated-vm"
			ctx.vmPub.Spec.Target.Location.Name = "updated-cl"
			ctx.WebhookRequestContext.Obj, err = builder.ToUnstructured(ctx.vmPub)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should not allow the request", func() {
			Expect(response.Allowed).To(BeFalse())
			Expect(response.Result).ToNot(BeNil())
			Expect(string(response.Result.Reason)).To(ContainSubstring("field is immutable"))
		})
	})
}

func unitTestsValidateDelete() {
	var (
		ctx      *unitValidatingWebhookContext
		response admission.Response
	)

	BeforeEach(func() {
		ctx = newUnitTestContextForValidatingWebhook(false)
	})

	AfterEach(func() {
		ctx = nil
	})

	When("the delete is performed", func() {
		JustBeforeEach(func() {
			response = ctx.ValidateDelete(&ctx.WebhookRequestContext)
		})

		It("should allow the request", func() {
			Expect(response.Allowed).To(BeTrue())
			Expect(response.Result).ToNot(BeNil())
		})
	})
}
