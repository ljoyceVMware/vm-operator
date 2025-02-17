// Copyright (c) 2019 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package builder

import (
	goctx "context"
	"fmt"
	"net/http"
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	"github.com/pkg/errors"
	admissionv1 "k8s.io/api/admission/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	ctrlmgr "sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"github.com/vmware-tanzu/vm-operator/pkg/context"
	"github.com/vmware-tanzu/vm-operator/pkg/record"
)

// MutatingWebhook is an admissions webhook that mutates VM Operator
// resources.
type MutatingWebhook struct {
	admission.Webhook

	// Name is the name of the webhook.
	Name string

	// Path is the path of the webhook.
	Path string
}

// Mutator is used to create a new admissions webhook for mutating requests.
type Mutator interface {
	// For returns the GroupVersionKind for which this webhook mutates requests.
	For() schema.GroupVersionKind

	// Mutate will try modify invalid value.
	Mutate(*context.WebhookRequestContext) admission.Response
}

type MutatorFunc func(client.Client) Mutator

// NewMutatingWebhook returns a new admissions webhook for mutating requests.
func NewMutatingWebhook(
	ctx *context.ControllerManagerContext,
	mgr ctrlmgr.Manager,
	webhookName string,
	mutator Mutator) (*MutatingWebhook, error) {

	if webhookName == "" {
		return nil, errors.New("webhookName arg is empty")
	}
	if mutator == nil {
		return nil, errors.New("mutator arg is nil")
	}

	webhookNameShort := generateMutateName(webhookName, mutator.For())
	webhookPath := "/" + webhookNameShort
	webhookNameLong := fmt.Sprintf("%s/%s/%s", ctx.Namespace, ctx.Name, webhookNameShort)

	// Build the WebhookContext.
	webhookContext := &context.WebhookContext{
		Context:            ctx,
		Name:               webhookNameShort,
		Namespace:          ctx.Namespace,
		ServiceAccountName: ctx.ServiceAccountName,
		Recorder:           record.New(mgr.GetEventRecorderFor(webhookNameLong)),
		Logger:             ctx.Logger.WithName(webhookNameShort),
	}

	// Initialize the webhook's decoder.
	decoder, err := admission.NewDecoder(mgr.GetScheme())
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize decoder")
	}

	// Create the webhook.
	return &MutatingWebhook{
		Name: webhookNameShort,
		Path: webhookPath,
		Webhook: webhook.Admission{
			Handler: &mutatingWebhookHandler{
				Decoder:        decoder,
				WebhookContext: webhookContext,
				Mutator:        mutator,
			},
		},
	}, nil
}

var _ admission.Handler = &mutatingWebhookHandler{}

type mutatingWebhookHandler struct {
	*context.WebhookContext
	*admission.Decoder
	Mutator
}

func (h *mutatingWebhookHandler) Handle(_ goctx.Context, req admission.Request) admission.Response {
	if h.Mutator == nil {
		panic("mutator should never be nil")
	}

	var (
		obj, oldObj *unstructured.Unstructured
	)

	// Get the Object for Create and Update operations.
	//
	// For Delete operations the OldObject contains the Object being
	// deleted. Please see https://github.com/kubernetes/kubernetes/pull/76346
	// for more information.

	if req.Operation == admissionv1.Create || req.Operation == admissionv1.Update {
		obj = &unstructured.Unstructured{}
		if err := h.DecodeRaw(req.Object, obj); err != nil {
			return webhook.Errored(http.StatusBadRequest, err)
		}
	} else if req.Operation == admissionv1.Delete {
		oldObj = &unstructured.Unstructured{}
		if err := h.DecodeRaw(req.OldObject, oldObj); err != nil {
			return webhook.Errored(http.StatusBadRequest, err)
		}
	}

	if req.Operation == admissionv1.Update && !obj.GetDeletionTimestamp().IsZero() {
		return admission.Allowed(AdmitMesgUpdateOnDeleting)
	}

	// Get the old Object for Update operations.
	if req.Operation == admissionv1.Update {
		oldObj = &unstructured.Unstructured{}
		if err := h.DecodeRaw(req.OldObject, oldObj); err != nil {
			return webhook.Errored(http.StatusBadRequest, err)
		}
	}

	webhookRequestContext := &context.WebhookRequestContext{
		WebhookContext:      h.WebhookContext,
		Op:                  req.Operation,
		Obj:                 obj,
		OldObj:              oldObj,
		UserInfo:            req.UserInfo,
		IsPrivilegedAccount: isPrivilegedAccount(h.WebhookContext, req.UserInfo),
		Logger:              h.WebhookContext.Logger.WithName(obj.GetNamespace()).WithName(obj.GetName()),
	}

	return h.Mutate(webhookRequestContext)
}

func generateMutateName(webhookName string, gvk schema.GroupVersionKind) string {
	return fmt.Sprintf("%s-mutate-", webhookName) +
		strings.ReplaceAll(gvk.Group, ".", "-") + "-" +
		gvk.Version + "-" + strings.ToLower(gvk.Kind)
}
