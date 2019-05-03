In the following documentation the explanation of the conversion process of particular function of ZILLIQA code is presented.

-----------------------------------------------------------

## The C++ code snippet form libNode/MicroBlockPostProcessing.cpp (133-143)

-----------------------------------------------------------

```
void Node::CommitMicroBlockConsensusBuffer() {
 lock_guard<mutex> g(m_mutexMicroBlockConsensusBuffer);

 for (const auto i : m_microBlockConsensusBuffer[m_mediator.m_consensusID]) {
   auto runconsensus = [this, i]() {
     ProcessMicroBlockConsensusCore(std::get<NODE_MSG>(i), MessageOffset::BODY,
                                    std::get<NODE_PEER>(i));
   };
   DetachedFunction(1, runconsensus);
 }
}
```


## ---------------------CPP FUNCTION EXPLANATION----------------

**Line 1:** Declaration of c++ func with no arguments, Particular function is declared in the class node in the node.h file.

**Line 2:** Mutex is being created to seclude the process in order to maintain the data continuity.

**Line 4:** For loop gets triggered by the using  auto variable for incrementation \
for range traversal on m_microBlockConsensusBuffer[] array.

**Line 5:** Using anonymous function concept and passing values of i in the anonymous func ]
and declaring the runconsensus variable as auto in order to make the variable generic \
with datatype and inorder to pass it as an argument in the function on line 20.

**Line 6:** Internal Process is triggered ProcessMicroBlockConsensusCore()
with three arguments accepting the messagebody, node message, and node peer values.

**Line7:** Internal Process is triggered with two arguments as message off set value ad anonymous func for achieving\
threading.


-----------------------------------------------------------------------------------------------------
##                                      FUNCTION IN **C++** DETAILED DESCRIPTION

This function in the c++ code is used to access the micro block buffer and \
pass on each vector of data from a map which holds consensusID as key and VectorOfNodeMsg as the value,\
on the function that processes the micro block based on the consensus core logic.

This whole process in the above is done by using mutexing and multi-threading.\
Each micro block is made to fucntion on a single thread intiated by the utility function \
declared using template in DetachedFunction.h of libUtil file.

As the function block opens mutex is presented which locks the state of the buffer for the time it is being processed.

DetachedFunction() is presenting every block on individual threads for parallel processing of the blocks.


-------------------------------------------------------------------------------------------------------
##                                      FUNCTION IN **GoLang** DETAILED DESCRIPTION

The above presented description of the function is aceived by replicating first level variable value sin the function in the same file.

```
std::mutex m_mutexMicroBlockConsensusBuffer;
std::unordered_map<uint32_t, VectorOfNodeMsg> m_microBlockConsensusBuffer;
```
_The above snippet is replicated as follows: LINE (41-73) test.go_

Here _map of consensusID_ which is a uin32 type value and vectorOfNodeMsg is being replcated by slice\
(sliceOfNodeMsg) this slice holds the data in the form of structure varibales which are: **struct Key, struct Peer, slice byteMsg**.

These ***three instances*** hold data that together makes up the nodeMsg.

Each struct in the NodeMsg struct holds multiple values like:\
  struct Key : private and public key of type Kyber.scalar and kyber.point respectively.


**Two functions generatePrivKey() and generatePubKey()** as the name explains to generate keys based on ecdsa. using _kyber/edis package_.


_struct Node_ holds values for mux that is mutex in go and struct mediator for presenting the value of consensusID. In order to attain data\
independence in a single file multiple structs are being made to use them individually according to the needs in the complete program.


```
  func (n *Node) commitMicroBlockConsensusBuffer() {}
```


The above line shows the function that is being replicated from c++ and also takes in a receiver an instance of Node type to replicate the\
object feature and passing of state to the function.

As in the ZILLIQA there a file that holds majority used instantiated variables which is mediator.h and here struct node holds the same status\
for as single file based program.


```
n.mux.Lock()

...code...

defer n.mux.Unlock()

```

The above present code block gives the functions the power of mutexing and defer provides unlocking of function as the last job to be done.


```
go func() {
			print(v.Key.pubKey)

			// C++ Function-------------------------------------------------
			// ProcessMicroBlockConsensusCore(a.bytesMsg, msgOffSet, a.Peer)
			// -------------------------------------------------------------
		}()
```

The above snippet uses the go routines in replacement of threading in the c++ code, letting individual block to be processed on seperate go \
routines.

```
func main() {

...code...
(LINE 110-144)test.go

}
```
Main function act as the starting point for the program and intializes all the strutures and pass on the values to the\
commitMicroBlockConsensusBuffer() function for further execution.




