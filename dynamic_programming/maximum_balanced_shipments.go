package dynamicprogramming

// You are given an integer array weight of length n, representing the weights of n parcels arranged in a straight line. A shipment is defined as a contiguous subarray of parcels. A shipment is considered balanced if the weight of the last parcel is strictly less than the maximum weight among all parcels in that shipment.
//
// Select a set of non-overlapping, contiguous, balanced shipments such that each parcel appears in at most one shipment (parcels may remain unshipped).
//
// Return the maximum possible number of balanced shipments that can be formed.
func maxBalancedShipments(weight []int) int {

	numShipments := 0
	maxInShipment := 0
	shipmentLen := 0

	// EZ greedy algo, we just try to make a shipment as soon as possible
	// make sure we don't do it before it's long enough
	for _, w := range weight {
		shipmentLen += 1
		if w > maxInShipment {
			maxInShipment = w
		}
		if shipmentLen < 2 {
			continue
		}
		if w < maxInShipment {
			numShipments += 1
			shipmentLen = 0
			maxInShipment = 0
		}

	}

	return numShipments
}
