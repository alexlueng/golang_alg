package facebook

import "fmt"

// feature 4 limit the request rate from users. The same request cannot be sent from the other platform
// until a specified amount of time has elapsed since the request from the first platform.
// The facebook status API queues requests using the requested Status ID and a timestamp.
// Solution
// 1) Initialize the hashtable
// 2) When a request arrives, check if it's a new request or a repeated request that came after the
// assigned time limit
// 3) If either of the above conditions is satisfied, accept the request and update the entry associated
// with that request in the hashtable
// 4) If not, reject the request

type RequestLimiter struct {
	requests map[string]int
}

func (r *RequestLimiter) RequestApprover(timestamp int, request string) bool {
	if _, ok := r.requests[request]; !ok {
		r.requests[request] = timestamp
		fmt.Println("Request Accepted")
		return true
	}
	oldTimestamp := r.requests[request]
	if (timestamp - oldTimestamp) >= 5 {
		r.requests[request] = timestamp
		fmt.Println("Request Accepted")
		return true
	}
	fmt.Println("Request rejected")
	return false
}
