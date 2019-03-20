// Copyright 2019 Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package result

import (
	"fmt"

	"github.com/gardener/gardener/pkg/client/kubernetes"
	tmv1beta1 "github.com/gardener/test-infra/pkg/apis/testmachinery/v1beta1"
	"github.com/gardener/test-infra/pkg/util"
	log "github.com/sirupsen/logrus"
)

// Collect collects results of all testruns and write them to a file.
// It returns wheter there are failed testruns or not.
func Collect(config *Config, tmClient kubernetes.Interface, namespace string, testruns []*tmv1beta1.Testrun, metadata *Metadata) (bool, error) {
	testrunsFailed := false
	for _, tr := range testruns {
		err := Output(config, tmClient, namespace, tr, metadata)
		if err != nil {
			return false, err
		}

		err = IngestFile(config.OutputFile, config.ESConfigName)
		if err != nil {
			log.Errorf("Cannot persist file %s: %s", config.OutputFile, err.Error())
		} else {
			err := MarkTestrunsAsIngested(tmClient, testruns)
			if err != nil {
				log.Warn(err.Error())
			}
		}

		if tr.Status.Phase == tmv1beta1.PhaseStatusSuccess {
			log.Infof("Testrun %s finished successfully", tr.Name)
		} else {
			testrunsFailed = true
			log.Errorf("Testrun %s failed with phase %s", tr.Name, tr.Status.Phase)
		}
		fmt.Print(util.PrettyPrintStruct(tr.Status))
	}

	return testrunsFailed, nil
}