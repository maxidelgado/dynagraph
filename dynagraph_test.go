package dynagraph

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
	"github.com/stretchr/testify/assert"
)

var db *dynamo.DB
var g graph

func init() {
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-2")}))
	db = dynamo.New(sess)
	g = graph{table: db.Table("torst-graph"), db: db}
}

type TestNode struct {
	Id     string
	Field1 string
	Field2 string
}

func createTestNode() string {
	n := TestNode{
		Field1: "testnode",
		Field2: "testnode",
	}

	n.Id, _ = g.Node(n.Id).Put(n)

	return n.Id
}

func TestGraphNode(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		t.Run("create node", func(t *testing.T) {
			var err error
			n := TestNode{
				Field1: "testnode",
				Field2: "testnode",
			}

			n.Id, err = g.Node(n.Id).Put(n)
			_ = g.Node().Delete(Filter{Id: n.Id, Type: "node-testnode"})

			assert.NoError(t, err)
		})
		t.Run("update node", func(t *testing.T) {
			n := TestNode{
				Id:     "testnode-someuuid",
				Field1: "updated",
			}

			err := g.Node(n.Id).Update(n)
			_ = g.Node().Delete(Filter{Id: n.Id, Type: "node-testnode"})

			assert.NoError(t, err)
		})
		t.Run("add property", func(t *testing.T) {
			n1 := TestNode{
				Id:     "testnode-someuuid",
				Field1: "testnode",
				Field2: "testnode",
			}

			err := g.Node(n1.Id).Prop(n1)
			_ = g.Node().Delete(Filter{Id: n1.Id, Type: "prop-testnode"})

			assert.NoError(t, err)
		})
		t.Run("add edge", func(t *testing.T) {
			n1 := TestNode{
				Id:     "testnode-someuuid",
				Field1: "testnode",
				Field2: "testnode",
			}

			err := g.Node(n1.Id).Edge(n1)
			_ = g.Node().Delete(Filter{Id: n1.Id, Type: "edge-testnode-someuuid"})

			assert.NoError(t, err)
		})
	})
}

func TestFilter(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		t.Run("get node", func(t *testing.T) {
			id := createTestNode()
			var actual TestNode
			expected := TestNode{
				Id:     id,
				Field1: "testnode",
				Field2: "testnode",
			}

			err := g.Query(Filter{Id: id, Type: "node:testnode"}).One(&actual)
			_ = g.Node().Delete(Filter{Id: id, Type: "node:testnode"})

			assert.NoError(t, err)
			assert.Equal(t, expected, actual)
		})
	})
}

func TestTransaction(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		t.Run("tx put", func(t *testing.T) {
			n1 := TestNode{
				Id: "testnode-someuuid",
			}
			n2 := TestNode{
				Id:     "testnode-otheruuid",
				Field1: "testnode",
				Field2: "testnode",
			}

			inputs := WriteItemsInput{}
			inputs, _ = inputs.AppendEdge(n1.Id, n2)

			err := g.Transaction().Put(inputs).Run()
			_ = g.Node().Delete(Filter{Id: n1.Id, Type: "edge-testnode-otheruuid"})

			assert.NoError(t, err)
		})
		t.Run("tx update", func(t *testing.T) {
			n1 := TestNode{
				Id: "testnode-someuuid",
			}
			n2 := TestNode{
				Id:     "testnode-otheruuid",
				Field2: "updated",
			}

			inputs := WriteItemsInput{}
			inputs, _ = inputs.AppendEdge(n1.Id, n2)

			err := g.Transaction().Update(inputs).Run()
			_ = g.Node().Delete(Filter{Id: n1.Id, Type: "edge-testnode-otheruuid"})

			assert.NoError(t, err)
		})
	})
}

func TestBatch(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		t.Run("add batch edges", func(t *testing.T) {
			n := TestNode{
				Field1: "testnode",
				Field2: "testnode",
			}
			n.Id = NewId(n)

			inputs := WriteItemsInput{}
			inputs, _ = inputs.AppendEdge(n.Id, n)
			count, err := g.Batch().Put(inputs)
			_ = g.Node().Delete(Filter{Id: n.Id, Type: "edge-" + n.Id})

			assert.NoError(t, err)
			assert.Equal(t, count, 1)
		})
		t.Run("get batch items", func(t *testing.T) {
			uid := createTestNode()
			var out []TestNode
			inputs := KeysInput{}
			inputs = inputs.AppendKeys(Filter{Id: uid, Type: "node:testnode"})

			err := g.Batch().Get(&out, inputs)
			_ = g.Node().Delete(Filter{Id: uid, Type: "node:testnode"})

			assert.NoError(t, err)
			assert.Len(t, out, 1)
		})
	})
}
