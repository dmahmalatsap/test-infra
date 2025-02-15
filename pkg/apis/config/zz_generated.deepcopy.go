// +build !ignore_autogenerated

/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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
// Code generated by deepcopy-gen. DO NOT EDIT.

package config

import (
	json "encoding/json"
	time "time"

	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Argo) DeepCopyInto(out *Argo) {
	*out = *in
	out.ArgoUI = in.ArgoUI
	if in.ChartValues != nil {
		in, out := &in.ChartValues, &out.ChartValues
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Argo.
func (in *Argo) DeepCopy() *Argo {
	if in == nil {
		return nil
	}
	out := new(Argo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArgoUI) DeepCopyInto(out *ArgoUI) {
	*out = *in
	out.Ingress = in.Ingress
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArgoUI.
func (in *ArgoUI) DeepCopy() *ArgoUI {
	if in == nil {
		return nil
	}
	out := new(ArgoUI)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BotConfiguration) DeepCopyInto(out *BotConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.Webserver = in.Webserver
	in.Dashboard.DeepCopyInto(&out.Dashboard)
	out.GitHubBot = in.GitHubBot
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BotConfiguration.
func (in *BotConfiguration) DeepCopy() *BotConfiguration {
	if in == nil {
		return nil
	}
	out := new(BotConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BotConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Certificate) DeepCopyInto(out *Certificate) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Certificate.
func (in *Certificate) DeepCopy() *Certificate {
	if in == nil {
		return nil
	}
	out := new(Certificate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Configuration) DeepCopyInto(out *Configuration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.Controller = in.Controller
	in.TestMachinery.DeepCopyInto(&out.TestMachinery)
	in.Argo.DeepCopyInto(&out.Argo)
	in.GitHub.DeepCopyInto(&out.GitHub)
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(S3)
		(*in).DeepCopyInto(*out)
	}
	if in.ElasticSearch != nil {
		in, out := &in.ElasticSearch, &out.ElasticSearch
		*out = new(ElasticSearch)
		**out = **in
	}
	if in.ReservedExcessCapacity != nil {
		in, out := &in.ReservedExcessCapacity, &out.ReservedExcessCapacity
		*out = new(ReservedExcessCapacity)
		(*in).DeepCopyInto(*out)
	}
	in.Observability.DeepCopyInto(&out.Observability)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Configuration.
func (in *Configuration) DeepCopy() *Configuration {
	if in == nil {
		return nil
	}
	out := new(Configuration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Configuration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Controller) DeepCopyInto(out *Controller) {
	*out = *in
	out.TTLController = in.TTLController
	out.WebhookConfig = in.WebhookConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Controller.
func (in *Controller) DeepCopy() *Controller {
	if in == nil {
		return nil
	}
	out := new(Controller)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Dashboard) DeepCopyInto(out *Dashboard) {
	*out = *in
	in.Authentication.DeepCopyInto(&out.Authentication)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Dashboard.
func (in *Dashboard) DeepCopy() *Dashboard {
	if in == nil {
		return nil
	}
	out := new(Dashboard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DashboardAuthentication) DeepCopyInto(out *DashboardAuthentication) {
	*out = *in
	if in.GitHub != nil {
		in, out := &in.GitHub, &out.GitHub
		*out = new(GitHubAuthentication)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DashboardAuthentication.
func (in *DashboardAuthentication) DeepCopy() *DashboardAuthentication {
	if in == nil {
		return nil
	}
	out := new(DashboardAuthentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticSearch) DeepCopyInto(out *ElasticSearch) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticSearch.
func (in *ElasticSearch) DeepCopy() *ElasticSearch {
	if in == nil {
		return nil
	}
	out := new(ElasticSearch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitHub) DeepCopyInto(out *GitHub) {
	*out = *in
	if in.Cache != nil {
		in, out := &in.Cache, &out.Cache
		*out = new(GitHubCache)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitHub.
func (in *GitHub) DeepCopy() *GitHub {
	if in == nil {
		return nil
	}
	out := new(GitHub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitHubAuthentication) DeepCopyInto(out *GitHubAuthentication) {
	*out = *in
	if in.OAuth != nil {
		in, out := &in.OAuth, &out.OAuth
		*out = new(OAuth)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitHubAuthentication.
func (in *GitHubAuthentication) DeepCopy() *GitHubAuthentication {
	if in == nil {
		return nil
	}
	out := new(GitHubAuthentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitHubBot) DeepCopyInto(out *GitHubBot) {
	*out = *in
	out.GitHubCache = in.GitHubCache
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitHubBot.
func (in *GitHubBot) DeepCopy() *GitHubBot {
	if in == nil {
		return nil
	}
	out := new(GitHubBot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitHubCache) DeepCopyInto(out *GitHubCache) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitHubCache.
func (in *GitHubCache) DeepCopy() *GitHubCache {
	if in == nil {
		return nil
	}
	out := new(GitHubCache)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ingress) DeepCopyInto(out *Ingress) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ingress.
func (in *Ingress) DeepCopy() *Ingress {
	if in == nil {
		return nil
	}
	out := new(Ingress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Locations) DeepCopyInto(out *Locations) {
	*out = *in
	if in.ExcludeDomains != nil {
		in, out := &in.ExcludeDomains, &out.ExcludeDomains
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Locations.
func (in *Locations) DeepCopy() *Locations {
	if in == nil {
		return nil
	}
	out := new(Locations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Logging) DeepCopyInto(out *Logging) {
	*out = *in
	if in.ChartValues != nil {
		in, out := &in.ChartValues, &out.ChartValues
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Logging.
func (in *Logging) DeepCopy() *Logging {
	if in == nil {
		return nil
	}
	out := new(Logging)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MinioConfiguration) DeepCopyInto(out *MinioConfiguration) {
	*out = *in
	out.Ingress = in.Ingress
	if in.ChartValues != nil {
		in, out := &in.ChartValues, &out.ChartValues
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MinioConfiguration.
func (in *MinioConfiguration) DeepCopy() *MinioConfiguration {
	if in == nil {
		return nil
	}
	out := new(MinioConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OAuth) DeepCopyInto(out *OAuth) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OAuth.
func (in *OAuth) DeepCopy() *OAuth {
	if in == nil {
		return nil
	}
	out := new(OAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Observability) DeepCopyInto(out *Observability) {
	*out = *in
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(Logging)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Observability.
func (in *Observability) DeepCopy() *Observability {
	if in == nil {
		return nil
	}
	out := new(Observability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReservedExcessCapacity) DeepCopyInto(out *ReservedExcessCapacity) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReservedExcessCapacity.
func (in *ReservedExcessCapacity) DeepCopy() *ReservedExcessCapacity {
	if in == nil {
		return nil
	}
	out := new(ReservedExcessCapacity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3) DeepCopyInto(out *S3) {
	*out = *in
	in.Server.DeepCopyInto(&out.Server)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3.
func (in *S3) DeepCopy() *S3 {
	if in == nil {
		return nil
	}
	out := new(S3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Server) DeepCopyInto(out *S3Server) {
	*out = *in
	if in.Minio != nil {
		in, out := &in.Minio, &out.Minio
		*out = new(MinioConfiguration)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Server.
func (in *S3Server) DeepCopy() *S3Server {
	if in == nil {
		return nil
	}
	out := new(S3Server)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TTLController) DeepCopyInto(out *TTLController) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TTLController.
func (in *TTLController) DeepCopy() *TTLController {
	if in == nil {
		return nil
	}
	out := new(TTLController)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestMachinery) DeepCopyInto(out *TestMachinery) {
	*out = *in
	in.Locations.DeepCopyInto(&out.Locations)
	if in.RetryTimeoutDuration != nil {
		in, out := &in.RetryTimeoutDuration, &out.RetryTimeoutDuration
		*out = new(time.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestMachinery.
func (in *TestMachinery) DeepCopy() *TestMachinery {
	if in == nil {
		return nil
	}
	out := new(TestMachinery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookConfig) DeepCopyInto(out *WebhookConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookConfig.
func (in *WebhookConfig) DeepCopy() *WebhookConfig {
	if in == nil {
		return nil
	}
	out := new(WebhookConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Webserver) DeepCopyInto(out *Webserver) {
	*out = *in
	out.Certificate = in.Certificate
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Webserver.
func (in *Webserver) DeepCopy() *Webserver {
	if in == nil {
		return nil
	}
	out := new(Webserver)
	in.DeepCopyInto(out)
	return out
}
