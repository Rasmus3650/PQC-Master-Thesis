package internal

import (
	"PQC-Master-Thesis/internal/common"
	"fmt"
	"math"
)

func (c *CROSS[T, P]) Leaves(tree [][]byte) [][]byte {
	result := [][]byte{}
	for i := 0; i < len(c.TreeParams.LSI); i++ {
		index := c.TreeParams.LSI[i]
		for j := 0; j < c.TreeParams.NCL[i]; j++ {
			result = append(result, tree[index+j])
		}
	}
	return result
}

func (c *CROSS[T, P]) BuildTree(seed, salt []byte) ([][]byte, error) {
	if c.ProtocolData.IsType(common.TYPE_BALANCED, common.TYPE_SMALL) {
		t := make([][]byte, c.TreeParams.Total_nodes)
		t[0] = seed
		start_node := 0
		for level := 0; level <= len(c.TreeParams.NPL)-1; level++ {
			for i := 0; i <= c.TreeParams.NPL[level]-c.TreeParams.LPL[level]-1; i++ {
				node := start_node + i
				left_child := c.LeftChild(node, level)
				right_child := left_child + 1
				// Expand parent seed, salt and parent index
				//fmt.Println("seed: ", t[node])
				//fmt.Println("salt: ", salt)
				hash, err := c.CSPRNG(append(t[node], salt...), (2*c.ProtocolData.Lambda)/8, uint16(0+node))
				if err != nil {
					return nil, fmt.Errorf("Error: %s", err)
				}
				t[left_child] = hash[:c.ProtocolData.Lambda/8]
				t[right_child] = hash[c.ProtocolData.Lambda/8:]
			}
			start_node += c.TreeParams.NPL[level]

		}
		return t, nil
	} else if c.ProtocolData.IsType(common.TYPE_FAST) {
		tree := make([][]byte, c.TreeParams.Total_nodes)
		tree[0] = seed
		quad_seeds, err := c.CSPRNG(append(tree[0], salt...), (4*c.ProtocolData.Lambda)/8, uint16(0))
		if err != nil {
			return nil, fmt.Errorf("Error: %s", err)
		}
		for i := 0; i < 4; i++ {
			tree[i+1] = quad_seeds[i*c.ProtocolData.Lambda/8 : (i+1)*c.ProtocolData.Lambda/8]
		}
		children := make([]int, 4)
		if c.ProtocolData.T%4 == 0 {
			for i := 0; i < len(children); i++ {
				children[i] = int(math.Floor(float64(c.ProtocolData.T) / 4))
			}
		} else if c.ProtocolData.T%4 == 1 {
			for i := 0; i < len(children); i++ {
				if i == 0 {
					children[i] = int(math.Floor(float64(c.ProtocolData.T)/4)) + 1
				} else {
					children[i] = int(math.Floor(float64(c.ProtocolData.T) / 4))
				}
			}
		} else if c.ProtocolData.T%4 == 2 {
			for i := 0; i < len(children); i++ {
				if i == 0 || i == 1 {
					children[i] = int(math.Floor(float64(c.ProtocolData.T)/4)) + 1
				} else {
					children[i] = int(math.Floor(float64(c.ProtocolData.T) / 4))
				}
			}
		} else if c.ProtocolData.T%4 == 3 {
			for i := 0; i < len(children); i++ {
				if i == 0 || i == 1 || i == 2 {
					children[i] = int(math.Floor(float64(c.ProtocolData.T)/4)) + 1
				} else {
					children[i] = int(math.Floor(float64(c.ProtocolData.T) / 4))
				}
			}
		}
		result := [][]byte{}
		dsc_counter := 0
		csprng_input := make([]byte, 3*c.ProtocolData.Lambda/8)
		copy(csprng_input[c.ProtocolData.Lambda/8:], salt)
		for i := 0; i < 4; i++ {
			dsc_counter += 1
			copy(csprng_input, tree[i+1])
			hash, err := c.CSPRNG(csprng_input, children[i]*(c.ProtocolData.Lambda/8), uint16(0+dsc_counter))
			if err != nil {
				return nil, fmt.Errorf("Error: %s", err)
			}
			for j := 0; j < children[i]; j++ {
				result = append(result, hash[j*c.ProtocolData.Lambda/8:(j+1)*c.ProtocolData.Lambda/8])
			}
		}
		return result, nil
	} else {
		return nil, fmt.Errorf("Scheme type not supported only balanced, small and fast are supported")
	}
}
func (c *CROSS[T, P]) SeedLeaves(seed, salt []byte) ([][]byte, error) {
	if c.ProtocolData.IsType(common.TYPE_BALANCED, common.TYPE_SMALL) {
		t, err := c.BuildTree(seed, salt)
		if err != nil {
			return nil, fmt.Errorf("Error: %s", err)
		}
		return c.Leaves(t), nil
	} else if c.ProtocolData.IsType(common.TYPE_FAST) {
		return c.BuildTree(seed, salt)
	} else {
		return nil, fmt.Errorf("Scheme type not supported only balanced, small and fast are supported")
	}
}

// List of all leaf indices for the tree
func (c *CROSS[T, P]) leafSet() []int {
	var result []int
	for i := 0; i < len(c.TreeParams.LSI); i++ {
		count := c.TreeParams.NCL[i]
		start := c.TreeParams.LSI[i]
		for j := 0; j < count; j++ {
			result = append(result, start+j)
		}
	}
	return result
}

func (c *CROSS[T, P]) computeNodesToPublish(chall_2 []bool) []bool {
	result := make([]bool, c.TreeParams.Total_nodes)
	ctr := 0
	for i := 0; i < len(c.TreeParams.LSI); i++ {
		for j := 0; j < c.TreeParams.NCL[i]; j++ {
			if chall_2[ctr] {
				result[c.TreeParams.LSI[i]+j] = chall_2[ctr]
			}
			ctr++
		}
	}
	leafset := c.leafSet()
	for i := len(result) - 1; i >= 0; i-- {
		level := c.GetLevelOfNode(i)
		if !Contains(leafset, i) && result[c.LeftChild(i, level)] && result[c.LeftChild(i, level)+1] {
			result[i] = true
			result[c.LeftChild(i, level)] = false
			result[c.LeftChild(i, level)+1] = false
		}
	}
	return result
}

func (c *CROSS[T, P]) SeedPath(seed, salt []byte, chall_2 []bool) ([][]byte, error) {
	if c.ProtocolData.IsType(common.TYPE_BALANCED, common.TYPE_SMALL) {
		t, err := c.BuildTree(seed, salt)
		if err != nil {
			return nil, err
		}
		path := c.computeNodesToPublish(chall_2)
		seed_path := make([][]byte, c.ProtocolData.TREE_NODES_TO_STORE)
		for i := 0; i < len(seed_path); i++ {
			seed_path[i] = make([]byte, c.ProtocolData.Lambda/8)
		}
		published := 0
		for i := 0; i < len(path); i++ {
			if path[i] {
				seed_path[published] = t[i]
				published++
			}
		}
		return seed_path, nil
	} else if c.ProtocolData.IsType(common.TYPE_FAST) {
		leaves, err := c.SeedLeaves(seed, salt)
		if err != nil {
			return nil, err
		}
		result := make([][]byte, c.ProtocolData.W)
		for i := 0; i < len(result); i++ {
			result[i] = make([]byte, c.ProtocolData.Lambda/8)
		}
		published := 0
		for i := 0; i < len(chall_2); i++ {
			if chall_2[i] {
				result[published] = leaves[i]
				published++
			}
		}
		return result, nil
	} else {
		return nil, fmt.Errorf("Scheme type not supported only balanced, small and fast are supported")
	}
}

// TODO important: Add the zero padding from line 413-419 of seedtree.c, else forgery might be possible
func (c *CROSS[T, P]) RebuildLeaves(path [][]byte, salt []byte, chall_2 []bool) ([][]byte, bool, error) {
	if c.ProtocolData.IsType(common.TYPE_BALANCED, common.TYPE_SMALL) {
		T_prime := c.computeNodesToPublish(chall_2)
		t := make([][]byte, c.TreeParams.Total_nodes)
		start_node := 1
		pub_nodes := 0
		for level := 1; level <= len(c.TreeParams.NPL)-1; level++ {
			for i := 0; i <= c.TreeParams.NPL[level]-1; i++ {
				node := start_node + i
				parent := c.Parent(node, level)
				left_child := c.LeftChild(node, level)
				right_child := left_child + 1
				if T_prime[node] && !T_prime[parent] {
					t[node] = path[pub_nodes]
					pub_nodes++
				}
				if T_prime[node] && i < c.TreeParams.NPL[level]-c.TreeParams.LPL[level] {
					hash, err := c.CSPRNG(append(t[node], salt...), (2*c.ProtocolData.Lambda)/8, uint16(0+node))
					if err != nil {
						return nil, false, fmt.Errorf("Error: %s", err)
					}
					t[left_child] = hash[:c.ProtocolData.Lambda/8]
					t[right_child] = hash[c.ProtocolData.Lambda/8:]
					T_prime[left_child] = true
					T_prime[right_child] = true
				}
			}
			start_node += c.TreeParams.NPL[level]
		}
		result := c.Leaves(t)

		error_rate := uint8(0)
		// Check each row of the remaining rows of the path
		for i := pub_nodes; i < c.ProtocolData.TREE_NODES_TO_STORE; i++ {
			// Check each byte in the row
			for j := 0; j < c.ProtocolData.Lambda/8; j++ {
				error_rate |= path[i][j]
			}
		}
		return result, error_rate == 0, nil
	} else if c.ProtocolData.IsType(common.TYPE_FAST) {
		round_seeds := make([][]byte, c.ProtocolData.T)
		published := 0
		for i := 0; i < c.ProtocolData.T; i++ {
			if chall_2[i] {
				round_seeds[i] = path[published]
				published++
			} else {
				round_seeds[i] = make([]byte, c.ProtocolData.Lambda/8)
			}
		}
		return round_seeds, true, nil
	} else {
		return nil, false, fmt.Errorf("Scheme type not supported only balanced, small and fast are supported")
	}
}
