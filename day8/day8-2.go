package main

func getNodeValue(n *node) int {
	value := 0
	if n.childCount == 0 {
		for _, v := range n.metadata {
			value += v
		}
	} else {
		for _, i := range n.metadata {
			if i > len(n.childNode) {
				value += 0
			} else {
				value += getNodeValue(n.childNode[i-1])
			}
		}
	}
	return value
}

func printNode(n *node) {
	println(n.childCount, n.metadataLen, n.metadata[0])
	for _, n_tmp := range n.childNode {
		printNode(n_tmp)
	}
}

func part2() {
	rootNode := part1()
	println("part2", getNodeValue(rootNode))
	//printNode(rootNode)
}
