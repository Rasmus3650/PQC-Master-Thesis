package seed

import (
	"PQC-Master-Thesis/internal/common"
	"PQC-Master-Thesis/internal/trees"
	"fmt"
	"math"

	"golang.org/x/crypto/sha3"
)

func Leaves(tree [][]byte, tree_params common.TreeParams) [][]byte {
	result := [][]byte{}
	for i := 0; i < len(tree_params.LSI); i++ {
		index := tree_params.LSI[i]
		for j := 0; j < tree_params.NCL[i]; j++ {
			result = append(result, tree[index+j])
		}
	}
	return result
}

func BuildTree(seed, salt []byte, proto_params common.ProtocolData, tree_params common.TreeParams) ([][]byte, error) {
	if proto_params.SchemeType == "balanced" || proto_params.SchemeType == "small" {
		T := make([][]byte, tree_params.Total_nodes)
		T[0] = seed
		start_node := 0
		res := make([][]byte, tree_params.Total_nodes)
		ctr := 0
		for level := 0; level <= len(tree_params.NPL)-1; level++ {
			for i := 0; i <= tree_params.NPL[level]-tree_params.LPL[level]-1; i++ {
				node := start_node + i
				left_child := trees.LeftChild(node, level, tree_params)
				right_child := left_child + 1
				// Expand parent seed, salt and parent index
				res[ctr] = append(append(T[node], salt...), trees.ParentIndex(node)...)
				ctr++
				hash := make([]byte, (2*proto_params.Lambda)/8)
				sha3.ShakeSum128(hash, append(append(T[node], salt...), trees.ParentIndex(node)...))
				T[left_child] = hash[:proto_params.Lambda/8]
				T[right_child] = hash[proto_params.Lambda/8:]
			}
			start_node += tree_params.NPL[level]

		}
		return T, nil
	} else if proto_params.SchemeType == "fast" {
		T := make([][]byte, tree_params.Total_nodes)
		T[0] = seed
		hash := make([]byte, (4*proto_params.Lambda)/8)
		sha3.ShakeSum128(hash, append(append(T[0], salt...), trees.ParentIndex(0)...))
		for i := 1; i <= 4; i++ {
			T[i] = hash[(i-1)*proto_params.Lambda/8 : i*proto_params.Lambda/8]
		}
		children := make([]int, 4)
		if proto_params.T%4 == 0 {
			for i := 0; i < len(children); i++ {
				children[i] = int(math.Floor(float64(proto_params.T) / 4))
			}
		} else if proto_params.T%4 == 1 {
			for i := 0; i < len(children); i++ {
				if i == 0 {
					children[i] = int(math.Floor(float64(proto_params.T)/4)) + 1
				} else {
					children[i] = int(math.Floor(float64(proto_params.T) / 4))
				}
			}
		} else if proto_params.T%4 == 2 {
			for i := 0; i < len(children); i++ {
				if i == 0 || i == 1 {
					children[i] = int(math.Floor(float64(proto_params.T)/4)) + 1
				} else {
					children[i] = int(math.Floor(float64(proto_params.T) / 4))
				}
			}
		} else if proto_params.T%4 == 3 {
			for i := 0; i < len(children); i++ {
				if i == 0 || i == 1 || i == 2 {
					children[i] = int(math.Floor(float64(proto_params.T)/4)) + 1
				} else {
					children[i] = int(math.Floor(float64(proto_params.T) / 4))
				}
			}
		}
		result := [][]byte{}
		for i := 0; i <= 3; i++ {
			hash := make([]byte, (children[i]*proto_params.Lambda)/8)
			sha3.ShakeSum128(hash, append(append(T[i+1], salt...), trees.ParentIndex(i+1)...))
			for j := 0; j < children[i]; j++ {
				result = append(result, hash[j*proto_params.Lambda/8:(j+1)*proto_params.Lambda/8])
			}
		}
		return result, nil
	} else {
		return nil, fmt.Errorf("Scheme type not supported only balanced, small and fast are supported")
	}
}
func SeedLeaves(seed, salt []byte, proto_params common.ProtocolData, tree_params common.TreeParams) ([][]byte, error) {
	if proto_params.SchemeType == "balanced" || proto_params.SchemeType == "small" {
		T, err := BuildTree(seed, salt, proto_params, tree_params)
		if err != nil {
			return nil, fmt.Errorf("Error: %s", err)
		}
		return Leaves(T, tree_params), nil
	} else if proto_params.SchemeType == "fast" {
		return BuildTree(seed, salt, proto_params, tree_params)
	} else {
		return nil, fmt.Errorf("Scheme type not supported only balanced, small and fast are supported")
	}
}

// List of all leaf indices for the tree
func leafSet(tree_params common.TreeParams) []int {
	var result []int
	for i := 0; i < len(tree_params.LSI); i++ {
		count := tree_params.NCL[i]
		start := tree_params.LSI[i]
		for j := 0; j < count; j++ {
			result = append(result, start+j)
		}
	}
	return result
}

func computeNodesToPublish(chall_2 []bool, tree_params common.TreeParams) []bool {
	result := make([]bool, tree_params.Total_nodes)
	ctr := 0
	for i := 0; i < len(tree_params.LSI); i++ {
		for j := 0; j < tree_params.NCL[i]; j++ {
			if chall_2[ctr] {
				result[tree_params.LSI[i]+j] = chall_2[ctr]
			}
			ctr++
		}
	}
	leafset := leafSet(tree_params)
	for i := len(result) - 1; i >= 0; i-- {
		level := trees.GetLevelOfNode(i, tree_params)
		if !trees.Contains(leafset, i) && result[trees.LeftChild(i, level, tree_params)] && result[trees.LeftChild(i, level, tree_params)+1] {
			result[i] = true
			result[trees.LeftChild(i, level, tree_params)] = false
			result[trees.LeftChild(i, level, tree_params)+1] = false
		}
	}
	return result
}

func SeedPath(seed, salt []byte, chall_2 []bool, proto_params common.ProtocolData, tree_params common.TreeParams) ([][]byte, error) {
	if proto_params.SchemeType == "balanced" || proto_params.SchemeType == "small" {
		T, err := BuildTree(seed, salt, proto_params, tree_params)
		if err != nil {
			return nil, err
		}
		path := computeNodesToPublish(chall_2, tree_params)
		seedPath := [][]byte{}
		for i := 0; i < len(path); i++ {
			if path[i] {
				seedPath = append(seedPath, T[i])
			}
		}
		return seedPath, nil
	} else if proto_params.SchemeType == "fast" {
		leaves, err := SeedLeaves(seed, salt, proto_params, tree_params)
		if err != nil {
			return nil, err
		}
		var result [][]byte
		for i := 0; i < len(chall_2); i++ {
			if chall_2[i] {
				result = append(result, leaves[i])
			}
		}
		return result, nil
	} else {
		return nil, fmt.Errorf("Scheme type not supported only balanced, small and fast are supported")
	}
}

func RebuildLeaves(path [][]byte, salt []byte, chall_2 []bool, proto_params common.ProtocolData, tree_params common.TreeParams) ([][]byte, error) {
	if proto_params.SchemeType == "balanced" || proto_params.SchemeType == "small" {
		T_prime := computeNodesToPublish(chall_2, tree_params)
		T := make([][]byte, tree_params.Total_nodes)
		start_node := 1
		pub_nodes := 0
		res := make([][]byte, tree_params.Total_nodes)
		ctr := 0
		for level := 1; level <= len(tree_params.NPL)-1; level++ {
			for i := 0; i <= tree_params.NPL[level]-1; i++ {
				node := start_node + i
				parent := trees.Parent(node, level, tree_params)
				left_child := trees.LeftChild(node, level, tree_params)
				right_child := left_child + 1
				if T_prime[node] && !T_prime[parent] {
					T[node] = path[pub_nodes]
					pub_nodes++
				}
				if T_prime[node] && i < tree_params.NPL[level]-tree_params.LPL[level] {
					hash := make([]byte, (2*proto_params.Lambda)/8)
					res[ctr] = append(append(T[node], salt...), trees.ParentIndex(node)...)
					ctr++
					sha3.ShakeSum128(hash, append(append(T[node], salt...), trees.ParentIndex(node)...))
					T[left_child] = hash[:proto_params.Lambda/8]
					T[right_child] = hash[proto_params.Lambda/8:]
					T_prime[left_child] = true
					T_prime[right_child] = true
				}
			}
			start_node += tree_params.NPL[level]
		}
		res_prime := [][]byte{}
		result := [][]byte{}
		leaves := Leaves(T, tree_params)
		for i := 0; i < len(leaves); i++ {
			if chall_2[i] {
				result = append(result, leaves[i])
				res_prime = append(res_prime, res[i])
			}
		}
		return result, nil
	} else if proto_params.SchemeType == "fast" {
		return path, nil
	} else {
		return nil, fmt.Errorf("Scheme type not supported only balanced, small and fast are supported")
	}
}
