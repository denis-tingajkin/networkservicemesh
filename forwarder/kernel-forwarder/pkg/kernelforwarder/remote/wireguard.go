package remote

import (
	"fmt"
	"net"

	"github.com/pkg/errors"
	"golang.zx2c4.com/wireguard/device"
	"golang.zx2c4.com/wireguard/ipc"
	"golang.zx2c4.com/wireguard/tun"
	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"

	"github.com/networkservicemesh/networkservicemesh/controlplane/api/connection"
	"github.com/networkservicemesh/networkservicemesh/controlplane/api/connection/mechanisms/wireguard"
)

const (
	wireguardPort = 51820
)

// CreateVXLANInterface creates a VXLAN interface
func (c *Connect) createWireguardInterface(ifaceName string, remoteConnection *connection.Connection, direction uint8) error {
	/* Create interface - host namespace */
	var localPrivateKey wgtypes.Key
	var remotePublicKey wgtypes.Key
	var dstIP net.IP
	var err error
	if direction == INCOMING {
		if localPrivateKey, err = wgtypes.ParseKey(remoteConnection.GetMechanism().GetParameters()[wireguard.DstPrivateKey]); err != nil {
			return errors.Errorf("failed to parse local private key: %v", err)
		}
		if remotePublicKey, err = wgtypes.ParseKey(remoteConnection.GetMechanism().GetParameters()[wireguard.SrcPublicKey]); err != nil {
			return errors.Errorf("failed to parse local private key: %v", err)
		}
		dstIP = net.ParseIP(remoteConnection.GetMechanism().GetParameters()[wireguard.SrcIP])
	} else {
		if localPrivateKey, err = wgtypes.ParseKey(remoteConnection.GetMechanism().GetParameters()[wireguard.SrcPrivateKey]); err != nil {
			return errors.Errorf("failed to parse local private key: %v", err)
		}
		if remotePublicKey, err = wgtypes.ParseKey(remoteConnection.GetMechanism().GetParameters()[wireguard.DstPublicKey]); err != nil {
			return errors.Errorf("failed to parse local private key: %v", err)
		}
		dstIP = net.ParseIP(remoteConnection.GetMechanism().GetParameters()[wireguard.DstIP])
	}

	wgDevice, err := createWireguardDevice(ifaceName)
	if err != nil {
		return errors.Errorf("Wireguard error: %v", err)
	}
	//defer wgDevice.Close()

	uapi, err := startWireguardAPI(ifaceName, wgDevice)
	if err != nil {
		wgDevice.Close()
		return errors.Errorf("Wireguard error: %v", err)
	}
	defer uapi.Close()

	err = configureWireguardDevice(ifaceName, localPrivateKey, remotePublicKey, dstIP)
	if err != nil {
		wgDevice.Close()
		return errors.Errorf("Wireguard error: %v", err)
	}

	return nil
}

func (c *Connect) deleteWireguardInterface(ifaceName string) error {
	if wgDevice, ok := c.wireguardDevices[ifaceName]; ok {
		wgDevice.Close()
		delete(c.wireguardDevices, ifaceName)
	}

	return nil
}

func createWireguardDevice(ifaceName string) (*device.Device, error) {
	tunIface, err := tun.CreateTUN(ifaceName, device.DefaultMTU)
	if err != nil {
		return nil, errors.Errorf("failed to create tun: %v", err)
	}

	logger := device.NewLogger(device.LogLevelDebug, fmt.Sprintf("Wireguard Error (%s): ", ifaceName))
	return device.NewDevice(tunIface, logger), nil
}

func startWireguardAPI(ifaceName string, wgDevice *device.Device) (net.Listener, error) {
	fileUAPI, err := ipc.UAPIOpen(ifaceName)
	if err != nil {
		return nil, err
	}

	uapi, err := ipc.UAPIListen(ifaceName, fileUAPI)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			conn, err := uapi.Accept()
			if err != nil {
				return
			}
			go wgDevice.IpcHandle(conn)
		}
	}()

	return uapi, nil
}

func configureWireguardDevice(ifaceName string, localPrivateKey wgtypes.Key, remotePublicKey wgtypes.Key, dstIP net.IP) error {
	client, err := wgctrl.New()
	if err != nil {
		return errors.Errorf("failed to create configuration client:", err)
	}
	defer client.Close()

	_, ipnet, err := net.ParseCIDR("0.0.0.0/0")
	if err != nil {
		return errors.Errorf("failed to configure device: %v", err)
	}
	err = client.ConfigureDevice(ifaceName, wgtypes.Config{
		ListenPort: intPtr(wireguardPort),
		PrivateKey: &localPrivateKey,
		Peers: []wgtypes.PeerConfig{
			{
				PublicKey: remotePublicKey,
				AllowedIPs: []net.IPNet{
					*ipnet,
				},
				Endpoint: &net.UDPAddr{
					IP:   dstIP,
					Port: wireguardPort,
				},
			},
		},
	})

	return errors.Wrapf(err, "failed to configure device: %v", err)
}

func intPtr(v int) *int {
	return &v
}