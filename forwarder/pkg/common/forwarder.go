// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

import (
	"context"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/networkservicemesh/networkservicemesh/controlplane/api/crossconnect"
	local "github.com/networkservicemesh/networkservicemesh/controlplane/api/local/connection"
	remote "github.com/networkservicemesh/networkservicemesh/controlplane/api/remote/connection"
	"github.com/networkservicemesh/networkservicemesh/forwarder/api/forwarder"
	"github.com/networkservicemesh/networkservicemesh/pkg/tools"
	monitor_crossconnect "github.com/networkservicemesh/networkservicemesh/sdk/monitor/crossconnect"
)

type NSMDataplane interface {
	forwarder.MechanismsMonitorServer
	Init(*DataplaneConfig) error
	CreateDataplaneServer(*DataplaneConfig) forwarder.DataplaneServer
}

// TODO Convert all the defaults to properly use NsmBaseDir
const (
	NSMBaseDirKey                        = "NSM_BASEDIR"
	NSMBaseDirDefault                    = "/var/lib/networkservicemesh/"
	DataplaneRegistrarSocketKey          = "FORWARDER_REGISTRAR_SOCKET"
	DataplaneRegistrarSocketDefault      = "/var/lib/networkservicemesh/nsm.forwarder-registrar.io.sock"
	DataplaneRegistrarSocketTypeKey      = "FORWARDER_REGISTRAR_SOCKET_TYPE"
	DataplaneRegistrarSocketTypeDefault  = "unix"
	DataplaneMetricsEnabledKey           = "METRICS_COLLECTOR_ENABLED"
	DataplaneMetricsEnabledDefault       = false
	DataplaneMetricsRequestPeriodKey     = "METRICS_COLLECTOR_REQUEST_PERIOD"
	DataplaneMetricsRequestPeriodDefault = time.Second * 2
	DataplaneNameKey                     = "FORWARDER_NAME"
	DataplaneNameDefault                 = "vppagent"
	DataplaneSocketKey                   = "FORWARDER_SOCKET"
	DataplaneSocketDefault               = "/var/lib/networkservicemesh/nsm-vppagent.forwarder.sock"
	DataplaneSocketTypeKey               = "FORWARDER_SOCKET_TYPE"
	DataplaneSocketTypeDefault           = "unix"
	DataplaneSrcIPKey                    = "NSM_FORWARDER_SRC_IP"
)

// DataplaneConfig keeps the common configuration for a forwarding plane
type DataplaneConfig struct {
	Name                    string
	NSMBaseDir              string
	RegistrarSocket         string
	RegistrarSocketType     string
	DataplaneSocket         string
	DataplaneSocketType     string
	MechanismsUpdateChannel chan *Mechanisms
	Mechanisms              *Mechanisms
	MetricsEnabled          bool
	MetricsPeriod           time.Duration
	SrcIP                   net.IP
	EgressInterface         EgressInterfaceType
	GRPCserver              *grpc.Server
	Monitor                 monitor_crossconnect.MonitorServer
	Listener                net.Listener
}

// Mechanisms is a message used to communicate any changes in operational parameters and constraints
type Mechanisms struct {
	RemoteMechanisms []*remote.Mechanism
	LocalMechanisms  []*local.Mechanism
}

func createDataplaneConfig(forwarderGoals *DataplaneProbeGoals) *DataplaneConfig {
	cfg := &DataplaneConfig{}
	var ok bool

	cfg.Name, ok = os.LookupEnv(DataplaneNameKey)
	if !ok {
		logrus.Debugf("%s not set, using default %s", DataplaneNameKey, DataplaneNameDefault)
		cfg.Name = DataplaneNameDefault
	}

	cfg.DataplaneSocket, ok = os.LookupEnv(DataplaneSocketKey)
	if !ok {
		logrus.Infof("%s not set, using default %s", DataplaneSocketKey, DataplaneSocketDefault)
		cfg.DataplaneSocket = DataplaneSocketDefault
	}
	logrus.Infof("DataplaneSocket: %s", cfg.DataplaneSocket)

	err := tools.SocketCleanup(cfg.DataplaneSocket)
	if err != nil {
		logrus.Fatalf("Error cleaning up socket %s: %s", cfg.DataplaneSocket, err)
	} else {
		forwarderGoals.SetSocketCleanReady()
	}

	cfg.DataplaneSocketType, ok = os.LookupEnv(DataplaneSocketTypeKey)
	if !ok {
		logrus.Infof("%s not set, using default %s", DataplaneSocketTypeKey, DataplaneSocketTypeDefault)
		cfg.DataplaneSocketType = DataplaneSocketTypeDefault
	}
	logrus.Infof("DataplaneSocketType: %s", cfg.DataplaneSocketType)

	cfg.NSMBaseDir, ok = os.LookupEnv(NSMBaseDirKey)
	if !ok {
		logrus.Infof("%s not set, using default %s", NSMBaseDirKey, NSMBaseDirDefault)
		cfg.NSMBaseDir = NSMBaseDirDefault
	}
	logrus.Infof("NSMBaseDir: %s", cfg.NSMBaseDir)

	cfg.RegistrarSocket, ok = os.LookupEnv(DataplaneRegistrarSocketKey)
	if !ok {
		logrus.Infof("%s not set, using default %s", DataplaneRegistrarSocketKey, DataplaneRegistrarSocketDefault)
		cfg.RegistrarSocket = DataplaneRegistrarSocketDefault
	}
	logrus.Infof("RegistrarSocket: %s", cfg.RegistrarSocket)

	cfg.RegistrarSocketType, ok = os.LookupEnv(DataplaneRegistrarSocketTypeKey)
	if !ok {
		logrus.Infof("%s not set, using default %s", DataplaneRegistrarSocketTypeKey, DataplaneRegistrarSocketTypeDefault)
		cfg.RegistrarSocketType = DataplaneRegistrarSocketTypeDefault
	}
	logrus.Infof("RegistrarSocketType: %s", cfg.RegistrarSocketType)

	cfg.GRPCserver = tools.NewServer(context.Background())

	cfg.Monitor = monitor_crossconnect.NewMonitorServer()
	crossconnect.RegisterMonitorCrossConnectServer(cfg.GRPCserver, cfg.Monitor)

	cfg.MetricsEnabled = DataplaneMetricsEnabledDefault
	val, ok := os.LookupEnv(DataplaneMetricsEnabledKey)
	if ok {
		res, err := strconv.ParseBool(val)
		if err == nil {
			cfg.MetricsEnabled = res
		}
	}
	logrus.Infof("MetricsEnabled: %v", cfg.MetricsEnabled)

	if cfg.MetricsEnabled {
		cfg.MetricsPeriod = DataplaneMetricsRequestPeriodDefault
		if val, ok = os.LookupEnv(DataplaneMetricsRequestPeriodKey); ok {
			parsedPeriod, err := time.ParseDuration(val)
			if err == nil {
				cfg.MetricsPeriod = parsedPeriod
			}
		}
		logrus.Infof("MetricsPeriod: %v ", cfg.MetricsPeriod)
	}

	srcIPStr, ok := os.LookupEnv(DataplaneSrcIPKey)
	if !ok {
		logrus.Fatalf("Env variable %s must be set to valid srcIP for use for tunnels from this Pod.  Consider using downward API to do so.", DataplaneSrcIPKey)
	} else {
		forwarderGoals.SetSrcIPReady()
	}
	cfg.SrcIP = net.ParseIP(srcIPStr)
	if cfg.SrcIP == nil {
		logrus.Fatalf("Env variable %s must be set to a valid IP address, was set to %s", DataplaneSrcIPKey, srcIPStr)
	} else {
		forwarderGoals.SetValidIPReady()
	}
	cfg.EgressInterface, err = NewEgressInterface(cfg.SrcIP)
	if err != nil {
		logrus.Fatalf("Unable to find egress Interface: %s", err)
	} else {
		forwarderGoals.SetNewEgressIFReady()
	}
	logrus.Infof("SrcIP: %s, IfaceName: %s, SrcIPNet: %s", cfg.SrcIP, cfg.EgressInterface.Name(), cfg.EgressInterface.SrcIPNet())

	return cfg
}

// CreateDataplane creates new Dataplane Registrar client
func CreateDataplane(dp NSMDataplane, forwarderGoals *DataplaneProbeGoals) *DataplaneRegistration {
	start := time.Now()
	// Populate common configuration
	config := createDataplaneConfig(forwarderGoals)

	// Initialize the forwarder
	err := dp.Init(config)
	if err != nil {
		logrus.Fatalf("Dataplane initialization failed: %s ", err)
	}

	// Verify the configuration is populated
	if !sanityCheckConfig(config) {
		logrus.Fatalf("Dataplane configuration sanity check failed: %s ", err)
	}

	// Prepare the gRPC server
	config.Listener, err = net.Listen(config.DataplaneSocketType, config.DataplaneSocket)
	if err != nil {
		logrus.Fatalf("Error listening on socket %s: %s ", config.DataplaneSocket, err)
	} else {
		forwarderGoals.SetSocketListenReady()
	}

	forwarder.RegisterDataplaneServer(config.GRPCserver, dp.CreateDataplaneServer(config))
	forwarder.RegisterMechanismsMonitorServer(config.GRPCserver, dp)

	// Start the server
	logrus.Infof("Creating %s server...", config.Name)
	go func() {
		_ = config.GRPCserver.Serve(config.Listener)
	}()
	logrus.Infof("%s server serving", config.Name)

	logrus.Debugf("Starting the %s forwarder server took: %s", config.Name, time.Since(start))

	logrus.Info("Creating Dataplane Registrar Client...")
	registrar := NewDataplaneRegistrarClient(config.RegistrarSocketType, config.RegistrarSocket)
	registration := registrar.Register(context.Background(), config.Name, config.DataplaneSocket, nil, nil)
	logrus.Info("Registered Dataplane Registrar Client")

	return registration
}

func sanityCheckConfig(forwarderConfig *DataplaneConfig) bool {
	return len(forwarderConfig.Name) > 0 &&
		len(forwarderConfig.NSMBaseDir) > 0 &&
		len(forwarderConfig.RegistrarSocket) > 0 &&
		len(forwarderConfig.RegistrarSocketType) > 0 &&
		len(forwarderConfig.DataplaneSocket) > 0 &&
		len(forwarderConfig.DataplaneSocketType) > 0
}

// SanityCheckConnectionType checks whether the forwarding plane supports the connection type in the request
func SanityCheckConnectionType(mechanisms *Mechanisms, crossConnect *crossconnect.CrossConnect) error {
	localFound, remoteFound := false, false
	/* Verify local mechanisms */
	for _, mech := range mechanisms.LocalMechanisms {
		if crossConnect.GetLocalSource().GetMechanism().GetType() == mech.GetType() || crossConnect.GetLocalDestination().GetMechanism().GetType() == mech.GetType() {
			localFound = true
			break
		}
	}
	/* Verify remote mechanisms */
	for _, mech := range mechanisms.RemoteMechanisms {
		if crossConnect.GetRemoteSource().GetMechanism().GetType() == mech.GetType() || crossConnect.GetRemoteDestination().GetMechanism().GetType() == mech.GetType() {
			remoteFound = true
			break
		}
	}
	/* If none of them matched, mechanism is not supported by the forwarding plane */
	if !localFound && !remoteFound {
		return errors.New("connection mechanism type not supported by the forwarding plane")
	}
	return nil
}