package vanilla

import (
	"PQC-Master-Thesis/internal/common"
	"crypto/rand"
	"fmt"
	"math/big"

	"golang.org/x/crypto/sha3"
)

// TODO: This needs to Fisher-Yates shuffle
func (c *CROSSInstance) expand_digest_to_fixed_weight(digest_chall_2 []byte) []bool {
	chall_2 := make([]byte, c.ProtocolData.T)
	sha3.ShakeSum128(chall_2, append(digest_chall_2, byte(3*c.ProtocolData.T))) // 3*T = T+c+1

	bool_chall_2 := make([]bool, c.ProtocolData.T)
	for i := range chall_2 {
		bool_chall_2[i] = chall_2[i]%2 == 1
	}

	return bool_chall_2
}

func (c *CROSSInstance) expandSK(seed_sk []byte) ([]byte, [][]byte) {
	seed_e_pk := make([]byte, (4*c.ProtocolData.Lambda)/8)
	sha3.ShakeSum128(seed_e_pk, append(seed_sk, byte(3*c.ProtocolData.T+1)))
	seed_e := seed_e_pk[:2*c.ProtocolData.Lambda/8]
	seed_pk := seed_e_pk[2*c.ProtocolData.Lambda/8:]

	n_minus_k := c.ProtocolData.N - c.ProtocolData.K
	V := make([][]byte, n_minus_k)
	for i := range V {
		V[i] = make([]byte, c.ProtocolData.K)
	}
	buffer := make([]byte, n_minus_k*c.ProtocolData.K)

	// Security probably dies here since p=509 in RSDP-G, might be fine for RSDP
	sha3.ShakeSum128(buffer, append(seed_pk, byte(3*c.ProtocolData.T+2)))
	idx := 0
	for i := 0; i < n_minus_k; i++ {
		for j := 0; j < c.ProtocolData.K; j++ {
			// Ensure values are in Fp
			V[i][j] = buffer[idx]%byte(c.ProtocolData.P-1) + 1
			if V[i][j] > byte(c.ProtocolData.P) {
				panic("V[i][j] > P")
			}
			idx++
		}
	}
	// This will generate trailing zeros in each row, might be wrong?
	H := make([][]byte, n_minus_k)
	for i := range H {
		H[i] = make([]byte, c.ProtocolData.N)
		// Copy V part
		copy(H[i][:c.ProtocolData.K], V[i])
		// Add identity matrix part
		H[i][c.ProtocolData.K+i] = 1
	}
	// This needs to be rejection sampling
	e_bar := make([]byte, c.ProtocolData.N)
	sha3.ShakeSum128(e_bar, append(seed_e, byte(3*c.ProtocolData.T+3)))
	for i, v := range e_bar {
		e_bar[i] = v%byte(c.ProtocolData.Z-1) + 1
	}
	return e_bar, H
}

// Probably needs re-writing
func fz_vec_sub(e_bar, e_bar_prime []byte) []byte {
	result := make([]byte, len(e_bar))
	for i := range e_bar {
		result[i] = e_bar[i] - e_bar_prime[i]
	}
	return result
}

func element_wise_mul(v, u_prime []byte, Z int) []byte {
	result := make([]byte, len(v))
	for i := range v {
		result[i] = byte(v[i]*u_prime[i]) % byte(Z)
	}
	return result
}

func (c *CROSSInstance) Sign(sk, msg []byte) ([][]byte, error) {
	e_bar, H := c.expandSK(sk)
	C := 2*c.ProtocolData.T - 1
	seed := make([]byte, c.ProtocolData.Lambda/8)
	salt := make([]byte, (2*c.ProtocolData.Lambda)/8)
	rand.Read(seed)
	rand.Read(salt)
	commitments, err := c.SeedLeaves(seed, salt)
	if err != nil {
		return nil, fmt.Errorf("Error building seed leaves: %v", err)
	}
	e_bar_prime := make([][]byte, c.ProtocolData.T)
	u_prime := make([][]byte, c.ProtocolData.T)
	v_bar := make([][]byte, c.ProtocolData.T)
	v := make([][]byte, c.ProtocolData.T)
	u := make([][]byte, c.ProtocolData.T)
	s_prime := make([][]byte, c.ProtocolData.T)
	cmt_0 := make([][]byte, c.ProtocolData.T)
	cmt_1 := make([][]byte, c.ProtocolData.T)
	for i := 0; i < c.ProtocolData.T; i++ {
		// TODO: PROPER SAMPLING!
		e_bar_buffer := make([]byte, c.ProtocolData.N)
		sha3.ShakeSum128(e_bar_buffer, append(append(commitments[i], salt...), byte(i+C)))
		for i, v := range e_bar_buffer {
			e_bar_buffer[i] = v%byte(c.ProtocolData.Z-1) + 1
		}
		e_bar_prime[i] = e_bar_buffer

		u_prime_buffer := make([]byte, c.ProtocolData.N)
		sha3.ShakeSum128(u_prime_buffer, append(append(commitments[i], salt...), byte(i+C)))
		for i, v := range u_prime_buffer {
			e_bar_buffer[i] = v%byte(c.ProtocolData.P-1) + 1
		}
		u_prime[i] = u_prime_buffer
		v_bar[i] = fz_vec_sub(e_bar, e_bar_prime[i])
		v_buffer := make([]byte, c.ProtocolData.N)
		for j := 0; j < c.ProtocolData.N; j++ {
			v_buffer[j] = byte(new(big.Int).Exp(big.NewInt(int64(c.ProtocolData.G)), big.NewInt(int64(v_bar[i][j])), nil).Int64())
		}
		v[i] = v_buffer
		u[i] = element_wise_mul(v[i], u_prime[i], c.ProtocolData.Z)
		s_prime[i] = common.MultiplyVectorMatrix(u[i], common.TransposeByteMatrix(H))
		cmt_0_buffer := make([]byte, (2*c.ProtocolData.Lambda)/8)
		sha3.ShakeSum128(cmt_0_buffer, append(append(append(s_prime[i], v_bar[i]...), salt...), byte(i+C)))
		cmt_0[i] = cmt_0_buffer

		cmt_1_buffer := make([]byte, (2*c.ProtocolData.Lambda)/8)
		sha3.ShakeSum128(cmt_1_buffer, append(append(commitments[i], salt...), byte(i+C)))
		cmt_1[i] = cmt_1_buffer
	}
	digest_cmt_0, err := c.TreeRoot(cmt_0)
	digest_cmt_1 := make([]byte, (2*c.ProtocolData.Lambda)/8)
	flat_cmt_1 := make([]byte, 0)
	for _, b := range cmt_1 {
		flat_cmt_1 = append(flat_cmt_1, b...)
	}
	sha3.ShakeSum128(digest_cmt_1, flat_cmt_1)
	digest_cmt := make([]byte, (2*c.ProtocolData.Lambda)/8)
	sha3.ShakeSum128(digest_cmt, append(digest_cmt_0, digest_cmt_1...))

	digest_msg := make([]byte, (2*c.ProtocolData.Lambda)/8)
	sha3.ShakeSum128(digest_msg, msg)
	digest_chall_1 := make([]byte, (2*c.ProtocolData.Lambda)/8)
	sha3.ShakeSum128(digest_chall_1, append(append(digest_msg, digest_cmt...), salt...))
	//TODO: CSPRNG output needs to be in (F_p^*)^t, and fix value, gives us a problem with y[i] =
	chall_1 := make([]byte, c.ProtocolData.T)
	sha3.ShakeSum128(chall_1, append(digest_chall_1, byte(c.ProtocolData.T+C)))
	for i := range chall_1 {
		// -1, +1 to avoid 0
		chall_1[i] = chall_1[i]%byte(c.ProtocolData.P-1) + 1
	}
	var y []byte
	e_prime := make([][]byte, c.ProtocolData.T)
	for i := 0; i < c.ProtocolData.T; i++ {
		e_prime_i := make([]byte, c.ProtocolData.N)
		for j := 0; j < c.ProtocolData.N; j++ {
			//TODO: FIX THIS BULLSHIT MOST LIKELY QUITE WRONG!
			result := new(big.Int).Exp(big.NewInt(int64(c.ProtocolData.G)), big.NewInt(int64(e_bar_prime[i][j])), big.NewInt(int64(c.ProtocolData.P)))
			e_prime_i[j] = result.Bytes()[0]
			ctr := 0
			for _ = range e_prime_i[j] {
				ctr++
			}
			fmt.Println("Length of e_prime_i[j] = ", ctr, " Should be 1")
		}
		e_prime[i] = e_prime_i
		//TODO: Make sure this is correct
		y = common.ScalarVecMulByte(e_prime[i], chall_1[i])
		for j := 0; j < len(y); j++ {
			y[j] = (y[j] + u_prime[i][j]) % byte(255)
		}
	}
	digest_chall_2 := make([]byte, (2*c.ProtocolData.Lambda)/8)
	sha3.ShakeSum128(digest_chall_2, append(y[:c.ProtocolData.T], digest_chall_1...))
	chall_2 := c.expand_digest_to_fixed_weight(digest_chall_2)
	proof, err := c.TreeProof(cmt_0, chall_2)
	if err != nil {
		return nil, fmt.Errorf("Error generating proof: %v", err)
	}
	path, err := c.SeedPath(seed, salt, chall_2)
	if err != nil {
		return nil, fmt.Errorf("Error generating seed path: %v", err)
	}
	//TODO: Ensure compatibility with refernce code for this
	resp_0 := make([][]byte, c.ProtocolData.T)
	resp_1 := make([][]byte, c.ProtocolData.T)
	for i := 0; i < c.ProtocolData.T; i++ {
		if chall_2[i] == false {
			resp_0[i] = append([]byte{y[i]}, v_bar[i]...)
			resp_1[i] = cmt_1[i]
		}
	}
	//TODO: Temporary way of creating the signature, needs to pack properly
	sgn := make([][]byte, 7)
	sgn[0] = salt
	sgn[1] = digest_cmt
	sgn[2] = digest_chall_2
	sgn[3] = common.Flatten(path)
	sgn[4] = common.Flatten(proof)
	sgn[5] = common.Flatten(resp_0)
	sgn[6] = common.Flatten(resp_1)
	//sgn := append(append(append(append(append(append(salt, digest_cmt...), digest_chall_2...), common.Flatten(path)...),
	//		common.Flatten(proof)...), common.Flatten(resp_0)...), common.Flatten(resp_1)...)
	return sgn, nil
}

// DummySign is a dummy implementation of the Sign function, used for testing purposes ONLY
func (c *CROSSInstance) DummySign(sk, msg, seed, salt []byte) ([][]byte, error) {
	e_bar, H := c.expandSK(sk)
	C := 2*c.ProtocolData.T - 1
	commitments, err := c.SeedLeaves(seed, salt)
	if err != nil {
		return nil, fmt.Errorf("Error building seed leaves: %v", err)
	}
	e_bar_prime := make([][]byte, c.ProtocolData.T)
	u_prime := make([][]byte, c.ProtocolData.T)
	v_bar := make([][]byte, c.ProtocolData.T)
	v := make([][]byte, c.ProtocolData.T)
	u := make([][]byte, c.ProtocolData.T)
	s_prime := make([][]byte, c.ProtocolData.T)
	cmt_0 := make([][]byte, c.ProtocolData.T)
	cmt_1 := make([][]byte, c.ProtocolData.T)
	for i := 0; i < c.ProtocolData.T; i++ {
		// TODO: PROPER SAMPLING!
		e_bar_buffer := make([]byte, c.ProtocolData.N)
		sha3.ShakeSum128(e_bar_buffer, append(append(commitments[i], salt...), byte(i+C)))
		for i, v := range e_bar_buffer {
			e_bar_buffer[i] = v%byte(c.ProtocolData.Z-1) + 1
		}
		e_bar_prime[i] = e_bar_buffer

		u_prime_buffer := make([]byte, c.ProtocolData.N)
		sha3.ShakeSum128(u_prime_buffer, append(append(commitments[i], salt...), byte(i+C)))
		for i, v := range u_prime_buffer {
			e_bar_buffer[i] = v%byte(c.ProtocolData.P-1) + 1
		}
		u_prime[i] = u_prime_buffer
		v_bar[i] = fz_vec_sub(e_bar, e_bar_prime[i])
		v_buffer := make([]byte, c.ProtocolData.N)
		for j := 0; j < c.ProtocolData.N; j++ {
			v_buffer[j] = byte(new(big.Int).Exp(big.NewInt(int64(c.ProtocolData.G)), big.NewInt(int64(v_bar[i][j])), nil).Int64())
		}
		v[i] = v_buffer
		u[i] = element_wise_mul(v[i], u_prime[i], c.ProtocolData.Z)
		s_prime[i] = common.MultiplyVectorMatrix(u[i], common.TransposeByteMatrix(H))
		cmt_0_buffer := make([]byte, (2*c.ProtocolData.Lambda)/8)
		sha3.ShakeSum128(cmt_0_buffer, append(append(append(s_prime[i], v_bar[i]...), salt...), byte(i+C)))
		cmt_0[i] = cmt_0_buffer

		cmt_1_buffer := make([]byte, (2*c.ProtocolData.Lambda)/8)
		sha3.ShakeSum128(cmt_1_buffer, append(append(commitments[i], salt...), byte(i+C)))
		cmt_1[i] = cmt_1_buffer
	}
	digest_cmt_0, err := c.TreeRoot(cmt_0)
	digest_cmt_1 := make([]byte, (2*c.ProtocolData.Lambda)/8)
	flat_cmt_1 := make([]byte, 0)
	for _, b := range cmt_1 {
		flat_cmt_1 = append(flat_cmt_1, b...)
	}
	sha3.ShakeSum128(digest_cmt_1, flat_cmt_1)
	digest_cmt := make([]byte, (2*c.ProtocolData.Lambda)/8)
	sha3.ShakeSum128(digest_cmt, append(digest_cmt_0, digest_cmt_1...))

	digest_msg := make([]byte, (2*c.ProtocolData.Lambda)/8)
	sha3.ShakeSum128(digest_msg, msg)
	digest_chall_1 := make([]byte, (2*c.ProtocolData.Lambda)/8)
	sha3.ShakeSum128(digest_chall_1, append(append(digest_msg, digest_cmt...), salt...))
	//TODO: CSPRNG output needs to be in (F_p^*)^t, and fix value, gives us a problem with y[i] =
	chall_1 := make([]byte, c.ProtocolData.T)
	sha3.ShakeSum128(chall_1, append(digest_chall_1, byte(c.ProtocolData.T+C)))
	for i := range chall_1 {
		// -1, +1 to avoid 0
		chall_1[i] = chall_1[i]%byte(c.ProtocolData.P-1) + 1
	}
	var y []byte
	e_prime := make([][]byte, c.ProtocolData.T)
	for i := 0; i < c.ProtocolData.T; i++ {
		e_prime_i := make([]byte, c.ProtocolData.N)
		for j := 0; j < c.ProtocolData.N; j++ {
			//TODO: FIX THIS BULLSHIT MOST LIKELY QUITE WRONG!
			result := new(big.Int).Exp(big.NewInt(int64(c.ProtocolData.G)), big.NewInt(int64(e_bar_prime[i][j])), big.NewInt(int64(c.ProtocolData.P)))
			e_prime_i[j] = result.Bytes()[0]
			ctr := 0
			for _ = range e_prime_i[j] {
				ctr++
			}
			fmt.Println("Length of e_prime_i[j] = ", ctr, " Should be 1")
		}
		e_prime[i] = e_prime_i
		//TODO: Make sure this is correct
		y = common.ScalarVecMulByte(e_prime[i], chall_1[i])
		for j := 0; j < len(y); j++ {
			y[j] = (y[j] + u_prime[i][j]) % byte(255)
		}
	}
	digest_chall_2 := make([]byte, (2*c.ProtocolData.Lambda)/8)
	sha3.ShakeSum128(digest_chall_2, append(y[:c.ProtocolData.T], digest_chall_1...))
	chall_2 := c.expand_digest_to_fixed_weight(digest_chall_2)
	proof, err := c.TreeProof(cmt_0, chall_2)
	if err != nil {
		return nil, fmt.Errorf("Error generating proof: %v", err)
	}
	path, err := c.SeedPath(seed, salt, chall_2)
	if err != nil {
		return nil, fmt.Errorf("Error generating seed path: %v", err)
	}
	//TODO: Ensure compatibility with refernce code for this
	resp_0 := make([][]byte, c.ProtocolData.T)
	resp_1 := make([][]byte, c.ProtocolData.T)
	for i := 0; i < c.ProtocolData.T; i++ {
		if chall_2[i] == false {
			resp_0[i] = append([]byte{y[i]}, v_bar[i]...)
			resp_1[i] = cmt_1[i]
		}
	}
	//TODO: Temporary way of creating the signature, needs to pack properly
	sgn := make([][]byte, 7)
	sgn[0] = salt
	sgn[1] = digest_cmt
	sgn[2] = digest_chall_2
	sgn[3] = common.Flatten(path)
	sgn[4] = common.Flatten(proof)
	sgn[5] = common.Flatten(resp_0)
	sgn[6] = common.Flatten(resp_1)
	//sgn := append(append(append(append(append(append(salt, digest_cmt...), digest_chall_2...), common.Flatten(path)...),
	//		common.Flatten(proof)...), common.Flatten(resp_0)...), common.Flatten(resp_1)...)
	return sgn, nil

}
