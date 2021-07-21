package hmqtt

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"strings"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"

	"github.com/feng-crazy/go-utils/uuid"
)

// MqttClient is parameters for Mqtt client.
type MqttClient struct {
	Qos          byte
	Retained     bool
	IP           string
	User         string
	Passwd       string
	CA           string
	Cert         string
	PrivateKey   string
	CAIsNotFile  bool
	WillEnabled  bool
	WillTopic    string
	WillPayload  []byte
	WillQos      byte
	WillRetained bool
	Client       mqtt.Client
	subTopicMap  map[string]mqtt.MessageHandler // 存在数据竞争，需要上锁
	mutex        sync.Mutex
}

// newTLSConfig new TLS configuration.
// Only one side check. Mqtt broker check the cert from client.
func (mc *MqttClient) newTLSConfig() (*tls.Config, error) {
	var err error
	certpool := x509.NewCertPool()
	var pemCerts []byte
	var cert tls.Certificate
	if mc.CAIsNotFile {
		pemCerts = []byte(mc.CA)
		cert, err = tls.X509KeyPair([]byte(mc.Cert), []byte(mc.PrivateKey))
		if err != nil {
			return nil, err
		}
	} else {
		pemCerts, err = ioutil.ReadFile(mc.CA)
		// Import client certificate/key pair
		cert, err = tls.LoadX509KeyPair(mc.Cert, mc.PrivateKey)
		if err != nil {
			return nil, err
		}
	}

	certpool.AppendCertsFromPEM(pemCerts)

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

func (mc *MqttClient) SetReconnectingCb(client mqtt.Client, ops *mqtt.ClientOptions) {
	log.Info("#################MQTT RECONNECTING!! ip ", mc.IP)
}

func (mc *MqttClient) SetConnectedCb(client mqtt.Client) {
	log.Info("#################MQTT Connected!!")
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	for topic, handler := range mc.subTopicMap {
		if tc := mc.Client.Subscribe(topic, mc.Qos, handler); tc.Wait() && tc.Error() != nil {
			log.Error(tc.Error())
		}
	}
}

func (mc *MqttClient) SetConnectLostCb(client mqtt.Client, err error) {
	log.Error("#################MQTT Connect lost !! error: ", err, " ip: ", mc.IP)
}

// Connect connect to the Mqtt server.
func (mc *MqttClient) Connect() error {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()

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
	}

	opts.SetUsername(mc.User)
	opts.SetPassword(mc.Passwd)

	opts.AddBroker(mc.IP)
	opts.SetClientID(uuid.NewUUIDv1())
	opts.SetCleanSession(false)
	opts.SetAutoReconnect(true)
	opts.SetConnectRetryInterval(10 * time.Second)
	opts.SetMaxReconnectInterval(1 * time.Minute)
	opts.SetConnectTimeout(2 * time.Second)
	opts.SetReconnectingHandler(mc.SetReconnectingCb)
	opts.SetOnConnectHandler(mc.SetConnectedCb)
	opts.SetConnectionLostHandler(mc.SetConnectLostCb)
	opts.WillEnabled = mc.WillEnabled
	opts.WillTopic = mc.WillTopic
	opts.WillPayload = mc.WillPayload
	opts.WillQos = mc.WillQos
	opts.WillRetained = mc.WillRetained

	mc.Client = mqtt.NewClient(opts)
	mc.Qos = 0          // At most 1 time
	mc.Retained = false // Not retained
	// The token is used to indicate when actions have completed.
	if tc := mc.Client.Connect(); tc.WaitTimeout(2*time.Second) && tc.Error() != nil {
		return tc.Error()
	}

	mc.subTopicMap = make(map[string]mqtt.MessageHandler)

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
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	if tc := mc.Client.Subscribe(topic, mc.Qos, onMessage); tc.Wait() && tc.Error() != nil {
		return tc.Error()
	}
	_, ok := mc.subTopicMap[topic]
	if !ok {
		mc.subTopicMap[topic] = onMessage
	} else {
		log.Error("topic has subscribe")
	}

	return nil
}

// Subscribe subsribe a Mqtt topic.
func (mc *MqttClient) UnSubscribe(topic string) error {
	mc.mutex.Lock()
	defer mc.mutex.Unlock()
	if tc := mc.Client.Unsubscribe(topic); tc.Wait() && tc.Error() != nil {
		return tc.Error()
	}
	_, ok := mc.subTopicMap[topic]
	if ok {
		delete(mc.subTopicMap, topic)
	}

	return nil
}

func NewMqttClient() *MqttClient {
	return &MqttClient{}
}
