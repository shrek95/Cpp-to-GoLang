// -------------------------------------------------------------

// The C++ code form MicroBlockPostProcessing.cpp (133-143)

// -------------------------------------------------------------

// void Node::CommitMicroBlockConsensusBuffer() {
//  lock_guard<mutex> g(m_mutexMicroBlockConsensusBuffer);

//  for (const auto i : m_microBlockConsensusBuffer[m_mediator.m_consensusID]) {
//    auto runconsensus = [this, i]() {
//      ProcessMicroBlockConsensusCore(std::get<NODE_MSG>(i), MessageOffset::BODY,
//                                     std::get<NODE_PEER>(i));
//    };
//    DetachedFunction(1, runconsensus);
//  }
// }

//----------------------------------------------------------------
//
//
// The GoLang code for the above:
//
//
//----------------------------------------------------------------

package main

import (
	"fmt"
	"sync"
	"time"

	"go.dedis.ch/kyber"
	"go.dedis.ch/kyber/group/edwards25519"
)

var curve = edwards25519.NewBlakeSHA256Ed25519()

// Mediator structure
type Mediator struct {
	consensusID uint32
}

// Node structure
type Node struct {
	mux sync.Mutex
	Mediator
}

// Key structure
type Key struct {
	privKey kyber.Scalar
	pubKey  kyber.Point
}

// Peer structure
type Peer struct {
	ipAddr        uint64
	listeningPort uint32
	hostname      string
}

// NodeMsg structure
type NodeMsg struct {
	Key
	Peer
	bytesMsg []byte // replacement of vector<uint8_t>
}

var sliceOfNodeMsg []NodeMsg

var microBlockConsensusBuffer = map[uint32]NodeMsg{}

func generatePrivKey() kyber.Scalar {
	privKey := curve.Scalar().Pick(curve.RandomStream())
	return privKey
}

func generatePubKey(privKey kyber.Scalar) kyber.Point {
	pubKey := curve.Point().Mul(privKey, curve.Point().Base())
	return pubKey
}

func print(key kyber.Point) {
	fmt.Println("Here is the pub key from the commitMicroBlockConsensusBuffer(): ", key)
}

func (n *Node) commitMicroBlockConsensusBuffer() {

	n.mux.Lock()

	for _, v := range microBlockConsensusBuffer {
		go func() {
			print(v.Key.pubKey)

			// C++ Function-------------------------------------------------
			// ProcessMicroBlockConsensusCore(a.bytesMsg, msgOffSet, a.Peer)
			// -------------------------------------------------------------
		}()
		// C++ Function-----------------------------------------------------
		// DetachedFunction(1, runconsensus)
		// ----------------------------------------------------------------
	}
	defer n.mux.Unlock()
}

func main() {

	privateKey := generatePrivKey()
	publicKey := generatePubKey(privateKey)

	fmt.Printf("Private Key : %s \n", privateKey)
	fmt.Printf("Public Key : %s\n\n", publicKey)

	k1 := Key{
		privKey: privateKey,
		pubKey:  publicKey,
	}

	p1 := Peer{
		ipAddr:        198,
		listeningPort: 3444,
		hostname:      "Herdius",
	}

	bMsg := []byte{1, 2, 3}

	n1 := NodeMsg{
		Key:      k1,
		Peer:     p1,
		bytesMsg: bMsg,
	}

	node := Node{}
	node.consensusID = 0

	sliceOfNodeMsg = append(sliceOfNodeMsg, n1)

	microBlockConsensusBuffer[node.consensusID] = sliceOfNodeMsg[0]

	node.commitMicroBlockConsensusBuffer()

	time.Sleep(1 * time.Second)
}
