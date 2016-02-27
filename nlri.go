package main

//----------------------------------------------------------------------------------------------------------------------
// NLRI
//----------------------------------------------------------------------------------------------------------------------

// NLRI is the interface for all network layer reachability information
type NLRI interface {
	Encode(buf *MsgBuffer)
}
