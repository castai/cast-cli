package client

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/castai/cast-cli/pkg/client/sdk"
)

func NewMock() Interface {
	cred1, cred2, cred3 := "cred1", "cred2", "cred3"
	c1 := "c1"
	n1 := "n1"
	c1Nodes := []sdk.Node{
		{
			Id:           stringPointer(n1),
			Cloud:        "aws",
			InstanceType: "t3a.large",
			Name:         stringPointer("node1"),
			Role:         "master",
			Shape:        "t3a.large",
			SpotConfig:   nil,
			State:        &sdk.NodeState{Phase: stringPointer("ready")},
			Network: &sdk.NodeNetwork{
				PrivateIp: "127.0.0.1",
				PublicIp:  "1.1.1.1",
			},
		},
	}
	return &mockClient{
		credentials: []sdk.CloudCredentials{
			{
				Cloud:  "aws",
				Id:     cred1,
				Name:   "aws",
				UsedBy: nil,
			},
			{
				Cloud:  "gcp",
				Id:     cred2,
				Name:   "gcp",
				UsedBy: nil,
			},
			{
				Cloud:  "azure",
				Id:     cred3,
				Name:   "azure",
				UsedBy: nil,
			},
		},
		clusters: map[string]sdk.KubernetesCluster{
			c1: {
				Addons:              nil,
				CloudCredentialsIDs: []string{cred1},
				Id:                  c1,
				Name:                "test-cluster-1",
				Region: sdk.ClusterRegion{
					Name:        "eu-central",
					DisplayName: " Europe Central (Frankfurt)",
				},
				Status: "ready",
				Nodes:  c1Nodes,
			},
		},
		nodes: map[string]map[string]sdk.Node{
			c1: {
				n1: c1Nodes[0],
			},
		},
		regions: []sdk.CastRegion{
			{
				Clouds: []sdk.Cloud{
					{
						Name: "aws",
					},
					{
						Name: "gcp",
					},
					{
						Name: "azure",
					},
					{
						Name: "do",
					},
				},
				DisplayName: "Europe Central (Frankfurt)",
				Name:        "eu-central",
			},
		},
		feedbackEvents: []sdk.KubernetesClusterFeedbackEvent{
			{
				CreatedAt: time.Date(2021, 1, 1, 12, 15, 5, 0, time.UTC),
				Id:        "",
				Message:   "[AWS] VPC Created",
				Severity:  "info",
			},
		},
	}
}

func stringPointer(v string) *string {
	return &v
}

type mockClient struct {
	credentials    []sdk.CloudCredentials
	clusters       map[string]sdk.KubernetesCluster
	nodes          map[string]map[string]sdk.Node
	regions        []sdk.CastRegion
	tokens         []sdk.AuthToken
	feedbackEvents []sdk.KubernetesClusterFeedbackEvent
}

func (m *mockClient) CloseNodeSSH(ctx context.Context, clusterID sdk.ClusterId, nodeID string) error {
	return nil
}

func (m *mockClient) GetClusterNode(ctx context.Context, clusterID sdk.ClusterId, nodeID string) (*sdk.Node, error) {
	nodes, ok := m.nodes[string(clusterID)]
	if !ok {
		return nil, fmt.Errorf("cluster %s not found", clusterID)
	}
	node, ok := nodes[nodeID]
	if !ok {
		return nil, fmt.Errorf("node %s not found", nodeID)
	}
	return &node, nil
}

func (m *mockClient) SetupNodeSSH(ctx context.Context, clusterID sdk.ClusterId, nodeID string, req sdk.SetupNodeSshJSONRequestBody) error {
	return nil
}

func (m *mockClient) ListClusterNodes(ctx context.Context, req sdk.ClusterId) ([]sdk.Node, error) {
	var res []sdk.Node
	for _, node := range m.nodes[string(req)] {
		res = append(res, node)
	}
	return res, nil
}

func (m *mockClient) GetCluster(ctx context.Context, req sdk.ClusterId) (*sdk.KubernetesCluster, error) {
	if c, ok := m.clusters[string(req)]; ok {
		return &c, nil
	}
	return nil, fmt.Errorf("cluster %s not found", req)
}

func (m *mockClient) DeleteCluster(ctx context.Context, req sdk.ClusterId) error {
	if _, ok := m.clusters[string(req)]; ok {
		delete(m.clusters, string(req))
		return nil
	}
	return fmt.Errorf("cluster %s not found", req)
}

func (m *mockClient) FeedbackEvents(ctx context.Context, req sdk.ClusterId) ([]sdk.KubernetesClusterFeedbackEvent, error) {
	return m.feedbackEvents, nil
}

func (m *mockClient) CreateNewCluster(ctx context.Context, req sdk.CreateNewClusterJSONRequestBody) (*sdk.KubernetesCluster, error) {
	newCluster := sdk.KubernetesCluster{
		Addons:              req.Addons,
		CloudCredentialsIDs: req.CloudCredentialsIDs,
		Id:                  uuid.New().String(),
		Name:                req.Name,
		Network:             req.Network,
		Nodes:               req.Nodes,
		ReconcileMode:       "ok",
		Region: sdk.ClusterRegion{
			Name: req.Region,
		},
		Status: "ready",
	}
	m.clusters[newCluster.Id] = newCluster
	return &newCluster, nil
}

func (m *mockClient) ListRegions(ctx context.Context) ([]sdk.CastRegion, error) {
	return m.regions, nil
}

func (m *mockClient) ListCloudCredentials(ctx context.Context) ([]sdk.CloudCredentials, error) {
	return m.credentials, nil
}

func (m *mockClient) GetClusterKubeconfig(ctx context.Context, req sdk.ClusterId) ([]byte, error) {
	config := `
apiVersion: v1
clusters:
- cluster:
    server: https://server.local.onmulti.cloud:6443
  name: test`
	return []byte(config), nil
}

func (m *mockClient) ListKubernetesClusters(ctx context.Context, req *sdk.ListKubernetesClustersParams) ([]sdk.KubernetesCluster, error) {
	list := make([]sdk.KubernetesCluster, 0, len(m.clusters))
	for _, cluster := range m.clusters {
		list = append(list, cluster)
	}
	return list, nil
}

func (m *mockClient) ListAuthTokens(ctx context.Context) ([]sdk.AuthToken, error) {
	return m.tokens, nil
}
