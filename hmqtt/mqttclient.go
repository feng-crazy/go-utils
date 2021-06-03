package hmqtt

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"

	"github.com/feng-crazy/go-utils/uuid"
)

// MqttClient is parameters for Mqtt client.
type MqttClient struct {
	Qos        byte
	Retained   bool
	IP         string
	User       string
	Passwd     string
	CA         string
	Cert       string
	PrivateKey string
	Client     mqtt.Client
}

// newTLSConfig new TLS configuration.
// Only one side check. Mqtt broker check the cert from client.
func (mc *MqttClient) newTLSConfig() (*tls.Config, error) {
	certpool := x509.NewCertPool()
	pemCerts, err := ioutil.ReadFile(mc.CA)
	if err == nil {
		certpool.AppendCertsFromPEM(pemCerts)
	}

	// Import client certificate/key pair
	cert, err := tls.LoadX509KeyPair(mc.Cert, mc.PrivateKey)
	if err != nil {
		return nil, err
	}

	// Just to print out the client certificate..
	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])
	if err != nil {
		panic(err)
	}
	// fmt.Println(cert.Leaf)

	// Create tls.Config with desired tls properties
	return &tls.Config{
		// RootCAs = certs used to verify server cert.
		RootCAs: certpool,
		// ClientAuth = whether to request cert from server.
		// Since the server is set up for SSL, this happens
		// anyways.
		ClientAuth: tls.NoClientCert,
		// ClientCAs = certs used to validate client cert.
		ClientCAs: nil,
		// InsecureSkipVerify = verify that cert contents
		// match server. IP matches what is in cert etc.
		InsecureSkipVerify: true,
		// Certificates = list of certs client sends to server.
		Certificates: []tls.Certificate{cert},
	}, nil
}

func (mc *MqttClient) DisConnect() error {
	mc.Client.Disconnect(100)
	return nil
}

func SetReconnectingCb(mqtt.Client, *mqtt.ClientOptions) {
	log.Error("#################MQTT RECONNECTING!!")
}

// Connect connect to the Mqtt server.
func (mc *MqttClient) Connect() error {
	opts := mqtt.NewClientOptions()
	if mc.Cert != "" {
		tlsConfig, err := mc.newTLSConfig()
		if err != nil {
			return err
		}
		opts.SetTLSConfig(tlsConfig)
		if !strings.Contains(mc.IP, "://") {
			mc.IP = "ssl://" + mc.IP
		} else {
			mc.IP = strings.Replace(mc.IP, "tcp://", "ssl://", 1)
		}
	} else {
		opts.SetUsername(mc.User)
		opts.SetPassword(mc.Passwd)
	}

	opts.AddBroker(mc.IP)
	opts.SetClientID(uuid.NewUUIDv1())
	opts.SetCleanSession(false)
	opts.SetAutoReconnect(true)
	opts.SetConnectRetryInterval(10 * time.Second)
	opts.SetMaxReconnectInterval(1 * time.Minute)
	opts.SetConnectTimeout(2 * time.Second)
	opts.SetReconnectingHandler(SetReconnectingCb)

	var reconn = false
	opts.SetConnectionLostHandler(func(client mqtt.Client, e error) {
		log.Errorf("The connection %s is disconnected due to error %s, will try to re-connect later.", mc.IP, e)
		mc.Client = client
		reconn = true
	})

	opts.SetOnConnectHandler(func(client mqtt.Client) {
		if reconn {
			log.Infof("The connection is %s re-established successfully.", mc.IP)
		}
	})

	mc.Client = mqtt.NewClient(opts)
	mc.Qos = 1          // At most 1 time
	mc.Retained = false // Not retained
	// The token is used to indicate when actions have completed.
	if tc := mc.Client.Connect(); tc.Wait() && tc.Error() != nil {
		return tc.Error()
	}

	return nil
}

// Publish publish Mqtt message.
func (mc *MqttClient) Publish(topic string, payload interface{}) error {
	if tc := mc.Client.Publish(topic, mc.Qos, mc.Retained, payload); tc.Wait() && tc.Error() != nil {
		return tc.Error()
	}
	return nil
}

// Subscribe subsribe a Mqtt topic.
func (mc *MqttClient) Subscribe(topic string, onMessage mqtt.MessageHandler) error {
	if tc := mc.Client.Subscribe(topic, mc.Qos, onMessage); tc.Wait() && tc.Error() != nil {
		return tc.Error()
	}
	return nil
}

// Subscribe subsribe a Mqtt topic.
func (mc *MqttClient) UnSubscribe(topic string) error {
	if tc := mc.Client.Unsubscribe(topic); tc.Wait() && tc.Error() != nil {
		return tc.Error()
	}
	return nil
}

func NewMqttClient() *MqttClient {
	return &MqttClient{}
}
