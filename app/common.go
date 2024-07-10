package app

import (
	"fmt"
	"strings"
)

/*** config ***/
type EnvSubsciption struct {
	Protocol     string
	EndPoint     string
	TopicArn     string
	QueueName    string
	Raw          bool
	FilterPolicy string
}

type EnvTopic struct {
	Name          string
	Subscriptions []EnvSubsciption
}

type EnvQueue struct {
	Name                          string
	ReceiveMessageWaitTimeSeconds int
	RedrivePolicy                 string
	MaximumMessageSize            int
	VisibilityTimeout             int
}

type EnvQueueAttributes struct {
	VisibilityTimeout             int
	ReceiveMessageWaitTimeSeconds int
	MaximumMessageSize            int
}

type ListenAddress struct {
	Network string
	Address string
}

func (a ListenAddress) IsUnixSocket() bool {
	return a.Network == "unix"
}

func (a ListenAddress) String() string {
	return fmt.Sprintf("%s://%s", a.Network, a.Address)
}

type Environment struct {
	Host                   string
	Port                   string
	SqsPort                string
	SnsPort                string
	Region                 string
	AccountID              string
	LogToFile              bool
	LogFile                string
	EnableDuplicates       bool
	Topics                 []EnvTopic
	Queues                 []EnvQueue
	QueueAttributeDefaults EnvQueueAttributes
	RandomLatency          RandomLatency
}

func (e Environment) GetListenAddresses() []ListenAddress {
	if strings.HasPrefix(e.Host, "/") {
		return []ListenAddress{{Network: "unix", Address: e.Host}}
	}
	var addresses []ListenAddress
	for _, port := range e.GetPorts() {
		addresses = append(addresses, ListenAddress{Network: "tcp", Address: fmt.Sprintf("0.0.0.0:%s", port)})
	}
	return addresses
}

func (e Environment) GetPorts() []string {
	if e.Port != "" {
		return []string{e.Port}
	}
	if e.SqsPort != "" && e.SnsPort != "" {
		return []string{e.SqsPort, e.SnsPort}
	}
	return []string{}
}

var CurrentEnvironment Environment

/*** Common ***/
type ResponseMetadata struct {
	RequestId string `xml:"RequestId"`
}

/*** Error Responses ***/
type ErrorResult struct {
	Type    string `xml:"Type,omitempty"`
	Code    string `xml:"Code,omitempty"`
	Message string `xml:"Message,omitempty"`
}

type ErrorResponse struct {
	Result    ErrorResult `xml:"Error"`
	RequestId string      `xml:"RequestId"`
}

type RandomLatency struct {
	Min int
	Max int
}
