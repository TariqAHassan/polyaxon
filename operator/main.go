/*
Copyright 2019 Polyaxon, Inc.

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

package main

import (
	"flag"
	"os"

	corev1alpha1 "github.com/polyaxon/polyaxon/operator/api/v1alpha1"
	"github.com/polyaxon/polyaxon/operator/controllers"
	"github.com/polyaxon/polyaxon/operator/controllers/config"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	mpijobv1 "github.com/kubeflow/mpi-operator/pkg/apis/kubeflow/v1alpha2"
	pytorchjobv1 "github.com/kubeflow/pytorch-operator/pkg/apis/pytorch/v1"
	tfjobv1 "github.com/kubeflow/tf-operator/pkg/apis/tensorflow/v1"

	// mxnetjobv1 "github.com/kubeflow/mxnet-operator/pkg/apis/mxnet/v1"

	httpRuntime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	polyaxonSDK "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_client"
	// +kubebuilder:scaffold:imports
)

// PolyaxonAuth provides a custom auth info writer
// TODO: Move to sdk
func PolyaxonAuth(name, value string) httpRuntime.ClientAuthInfoWriter {
	return httpRuntime.ClientAuthInfoWriterFunc(func(r httpRuntime.ClientRequest, _ strfmt.Registry) error {
		return r.SetHeaderParam("Authorization", name+" "+value)
	})
}

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	_ = clientgoscheme.AddToScheme(scheme)

	_ = corev1.AddToScheme(scheme)
	_ = appsv1.AddToScheme(scheme)
	_ = batchv1.AddToScheme(scheme)
	_ = corev1alpha1.AddToScheme(scheme)

	if config.GetBoolEnv(config.TFfJobEnabled, false) {
		tfjobv1.AddToScheme(scheme)
	}
	if config.GetBoolEnv(config.PytorchJobEnabled, false) {
		pytorchjobv1.AddToScheme(scheme)
	}
	if config.GetBoolEnv(config.MpiJobEnabled, false) {
		mpijobv1.AddToScheme(scheme)
	}
	// +kubebuilder:scaffold:scheme
}

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	var host string
	var token string

	// Allow to pass by env and override by flag
	if config.GetStrEnv(config.AgentToken, "") != "" {
		token = config.GetStrEnv(config.AgentToken, "")
	}

	if config.GetStrEnv(config.AgentAPIHostgo, "") != "" {
		host = config.GetStrEnv(config.AgentAPIHostgo, "")
	}

	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. Enabling this will ensure there is only one active controller manager.")
	flag.StringVar(&host, "host", host,
		"Polyaxon host.")
	flag.StringVar(&token, "token", token,
		"Polyaxon token.")
	flag.Parse()

	ctrl.SetLogger(zap.Logger(true))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: metricsAddr,
		LeaderElection:     enableLeaderElection,
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	plxClient := polyaxonSDK.New(httptransport.New(host, "", []string{"http"}), strfmt.Default)
	plxToken := PolyaxonAuth("Token", token)

	if err = (&controllers.PolyaxonNotebookReconciler{
		Client:    mgr.GetClient(),
		Log:       ctrl.Log.WithName("controllers").WithName("PolyaxonNotebook"),
		Scheme:    mgr.GetScheme(),
		PlxClient: plxClient,
		PlxToken:  plxToken,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "PolyaxonNotebook")
		os.Exit(1)
	}
	if err = (&controllers.PolyaxonTensorboardReconciler{
		Client:    mgr.GetClient(),
		Log:       ctrl.Log.WithName("controllers").WithName("PolyaxonTensorboard"),
		Scheme:    mgr.GetScheme(),
		PlxClient: plxClient,
		PlxToken:  plxToken,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "PolyaxonTensorboard")
		os.Exit(1)
	}
	if err = (&controllers.PolyaxonJobReconciler{
		Client:    mgr.GetClient(),
		Log:       ctrl.Log.WithName("controllers").WithName("PolyaxonJob"),
		Scheme:    mgr.GetScheme(),
		PlxClient: plxClient,
		PlxToken:  plxToken,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "PolyaxonJob")
		os.Exit(1)
	}

	if config.KFEnabled() {

		if err = (&controllers.PolyaxonKFReconciler{
			Client:    mgr.GetClient(),
			Log:       ctrl.Log.WithName("controllers").WithName("PolyaxonKF"),
			Scheme:    mgr.GetScheme(),
			PlxClient: plxClient,
			PlxToken:  plxToken,
		}).SetupWithManager(mgr); err != nil {
			setupLog.Error(err, "unable to create controller", "controller", "PolyaxonKF")
			os.Exit(1)
		}
	}
	// +kubebuilder:scaffold:builder

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}
