package v1alpha2

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SeldonDeployment is a specificaiton for serving deployment.
type SeldonDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
	Spec              *DeploymentSpec   `protobuf:"bytes,4,req,name=spec" json:"spec,omitempty"`
	Status            *DeploymentStatus `protobuf:"bytes,5,opt,name=status" json:"status,omitempty"`
}

type DeploymentSpec struct {
	Name        *string           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Predictors  []*PredictorSpec  `protobuf:"bytes,2,rep,name=predictors" json:"predictors,omitempty"`
	OauthKey    *string           `protobuf:"bytes,3,opt,name=oauth_key,json=oauthKey" json:"oauth_key,omitempty"`
	OauthSecret *string           `protobuf:"bytes,4,opt,name=oauth_secret,json=oauthSecret" json:"oauth_secret,omitempty"`
	Annotations map[string]string `protobuf:"bytes,5,rep,name=annotations" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

// Status for seldon deployment
type DeploymentStatus struct {
	State           *string            `protobuf:"bytes,1,opt,name=state" json:"state,omitempty"`
	Description     *string            `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	PredictorStatus []*PredictorStatus `protobuf:"bytes,3,rep,name=predictorStatus" json:"predictorStatus,omitempty"`
}

type PredictorSpec struct {
	Name            *string                      `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Graph           *PredictiveUnit              `protobuf:"bytes,2,req,name=graph" json:"graph,omitempty"`
	ComponentSpecs  []*corev1.PodTemplateSpec    `protobuf:"bytes,3,rep,name=componentSpecs" json:"componentSpecs,omitempty"`
	Replicas        *int32                       `protobuf:"varint,4,opt,name=replicas" json:"replicas,omitempty"`
	Annotations     map[string]string            `protobuf:"bytes,5,rep,name=annotations" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	EngineResources *corev1.ResourceRequirements `protobuf:"bytes,6,opt,name=engineResources" json:"engineResources,omitempty"`
	Labels          map[string]string            `protobuf:"bytes,7,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

type PredictorStatus struct {
	Name              *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Status            *string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	Description       *string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	Replicas          *int32  `protobuf:"varint,4,opt,name=replicas" json:"replicas,omitempty"`
	ReplicasAvailable *int32  `protobuf:"varint,5,opt,name=replicasAvailable" json:"replicasAvailable,omitempty"`
}

// *
// The main type of the predictive unit. Routers decide where requests are sent, e.g. AB Tests and Multi-Armed Bandits. Combiners ensemble responses from their children. Models are leaft nodes in the predictive tree and provide request/reponse functionality encapsulating a machine learning model. Transformers alter the request features.
type PredictiveUnit_PredictiveUnitType int32

const (
	// Each one of these defines a default combination of Predictive Unit Methods
	PredictiveUnit_UNKNOWN_TYPE       PredictiveUnit_PredictiveUnitType = 0
	PredictiveUnit_ROUTER             PredictiveUnit_PredictiveUnitType = 1
	PredictiveUnit_COMBINER           PredictiveUnit_PredictiveUnitType = 2
	PredictiveUnit_MODEL              PredictiveUnit_PredictiveUnitType = 3
	PredictiveUnit_TRANSFORMER        PredictiveUnit_PredictiveUnitType = 4
	PredictiveUnit_OUTPUT_TRANSFORMER PredictiveUnit_PredictiveUnitType = 5
)

type PredictiveUnit_PredictiveUnitImplementation int32

const (
	// Each one of these are hardcoded in the engine, no microservice is used
	PredictiveUnit_UNKNOWN_IMPLEMENTATION PredictiveUnit_PredictiveUnitImplementation = 0
	PredictiveUnit_SIMPLE_MODEL           PredictiveUnit_PredictiveUnitImplementation = 1
	PredictiveUnit_SIMPLE_ROUTER          PredictiveUnit_PredictiveUnitImplementation = 2
	PredictiveUnit_RANDOM_ABTEST          PredictiveUnit_PredictiveUnitImplementation = 3
	PredictiveUnit_AVERAGE_COMBINER       PredictiveUnit_PredictiveUnitImplementation = 4
)

type PredictiveUnit_PredictiveUnitMethod int32

const (
	PredictiveUnit_TRANSFORM_INPUT  PredictiveUnit_PredictiveUnitMethod = 0
	PredictiveUnit_TRANSFORM_OUTPUT PredictiveUnit_PredictiveUnitMethod = 1
	PredictiveUnit_ROUTE            PredictiveUnit_PredictiveUnitMethod = 2
	PredictiveUnit_AGGREGATE        PredictiveUnit_PredictiveUnitMethod = 3
	PredictiveUnit_SEND_FEEDBACK    PredictiveUnit_PredictiveUnitMethod = 4
)

type PredictiveUnit struct {
	Name           *string                                      `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Children       []*PredictiveUnit                            `protobuf:"bytes,2,rep,name=children" json:"children,omitempty"`
	Type           *PredictiveUnit_PredictiveUnitType           `protobuf:"varint,3,opt,name=type,enum=seldon.protos.PredictiveUnit_PredictiveUnitType" json:"type,omitempty"`
	Implementation *PredictiveUnit_PredictiveUnitImplementation `protobuf:"varint,4,opt,name=implementation,enum=seldon.protos.PredictiveUnit_PredictiveUnitImplementation" json:"implementation,omitempty"`
	Methods        []PredictiveUnit_PredictiveUnitMethod        `protobuf:"varint,5,rep,name=methods,enum=seldon.protos.PredictiveUnit_PredictiveUnitMethod" json:"methods,omitempty"`
	Endpoint       *Endpoint                                    `protobuf:"bytes,6,opt,name=endpoint" json:"endpoint,omitempty"`
	Parameters     []*Parameter                                 `protobuf:"bytes,7,rep,name=parameters" json:"parameters,omitempty"`
}

type Parameter_ParameterType int32

const (
	Parameter_INT    Parameter_ParameterType = 0
	Parameter_FLOAT  Parameter_ParameterType = 1
	Parameter_DOUBLE Parameter_ParameterType = 2
	Parameter_STRING Parameter_ParameterType = 3
	Parameter_BOOL   Parameter_ParameterType = 4
)

type Parameter struct {
	Name  *string                  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value *string                  `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	Type  *Parameter_ParameterType `protobuf:"varint,3,req,name=type,enum=seldon.protos.Parameter_ParameterType" json:"type,omitempty"`
}

type Endpoint_EndpointType int32

const (
	Endpoint_REST Endpoint_EndpointType = 0
	Endpoint_GRPC Endpoint_EndpointType = 1
)

var Endpoint_EndpointType_name = map[int32]string{
	0: "REST",
	1: "GRPC",
}
var Endpoint_EndpointType_value = map[string]int32{
	"REST": 0,
	"GRPC": 1,
}

type Endpoint struct {
	ServiceHost *string                `protobuf:"bytes,1,opt,name=service_host,json=serviceHost" json:"service_host,omitempty"`
	ServicePort *int32                 `protobuf:"varint,2,opt,name=service_port,json=servicePort" json:"service_port,omitempty"`
	Type        *Endpoint_EndpointType `protobuf:"varint,3,opt,name=type,enum=seldon.protos.Endpoint_EndpointType" json:"type,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SeldonDeploymentList is a list of SeldonDeployment
type SeldonDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []SeldonDeployment `json:"items"`
}
