package netflix

// Feature 6 As an alternative to feature 5, we want to implement a strategy of replacing the least frequently watched titles in the cache.

// Solution
// compare to feature 5, we will need another hashmap to store the frequency of each node, because search the frequency of node is vary slow. Then we can use the similar strategy of feature 5.
// 1) if an element is accessed and is present in our data set, we will:
//    %% increase its frequency count
//    %% move it to the end of its respective list
// 2) if an element is added and there is space for it in our data set, we create a node with the specified key and value, assign the node a frequency count of 1, and increase the size of the structure
// 3) if an element is added and eviction is needed, we will delete the keys and references of the least frequent node from both hashmaps, the we simply repeat step 2
// 4) if a key already exists for a certain key-value pair, we update the value of the node for the respective key

type LFUStructure struct {
	capacity int
	size     int
	minFreq  int
	freqDict map[int]*LFUList
	keyDict  map[int]*LFUNode
}

func (lf *LFUStructure) Get(key int) *LFUNode {
	if _, ok := lf.keyDict[key]; !ok {
		return nil
	}
	temp := lf.keyDict[key]

	lf.freqDict[temp.freq].DeleteNode(temp)
	if lf.freqDict[lf.keyDict[key].freq].head == nil {
		delete(lf.freqDict, lf.keyDict[key].freq)
		if lf.minFreq == lf.keyDict[key].freq {
			lf.minFreq++
		}
	}
	lf.keyDict[key].freq++

	if _, ok := lf.freqDict[lf.keyDict[key].freq]; !ok {
		lf.freqDict[lf.keyDict[key].freq] = &LFUList{}
	}
	lf.freqDict[lf.keyDict[key].freq].Append(lf.keyDict[key])
	return lf.keyDict[key]
}

func (lf *LFUStructure) Set(key, value int) {
	if lf.Get(key) != nil {
		lf.keyDict[key].val = value
		return
	}
	if lf.size == lf.capacity {
		delete(lf.keyDict, lf.freqDict[lf.minFreq].head.key)
		lf.freqDict[lf.minFreq].DeleteNode(lf.freqDict[lf.minFreq].head)
		if lf.freqDict[lf.minFreq].head == nil {
			delete(lf.freqDict, lf.minFreq)
		}
		lf.size--
	}

	lf.minFreq = 1
	lf.keyDict[key] = &LFUNode{key: key, val: value, freq: lf.minFreq}
	if _, ok := lf.freqDict[lf.keyDict[key].freq]; !ok {
		lf.freqDict[lf.keyDict[key].freq] = &LFUList{}
	}
	lf.freqDict[lf.keyDict[key].freq].Append(lf.keyDict[key])
	lf.size++
}
