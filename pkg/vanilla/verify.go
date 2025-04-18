package vanilla

import (
	"PQC-Master-Thesis/internal/common"
	"bytes"
	"fmt"
)

func (c *CROSSInstance[T, P]) int16ToT(arr []int) []T {
	res := make([]T, len(arr))
	for i := range arr {
		res[i] = T(arr[i])
	}
	return res
}

// TODO: Check if we are allowed to bail out early, maybe should wait till final check
func (c *CROSSInstance[T, P]) Verify(pk Pub, m []byte, sig Signature) (bool, error) {
	// Length checks for all attributes of the signature
	valid_signature := true
	if len(sig.Salt) != 2*c.ProtocolData.Lambda/8 {
		valid_signature = false
	}
	if len(sig.Digest_cmt) != 2*c.ProtocolData.Lambda/8 {
		valid_signature = false
	}
	if len(sig.Digest_chall_2) != 2*c.ProtocolData.Lambda/8 {
		valid_signature = false
	}

	if c.ProtocolData.IsType(common.TYPE_BALANCED, common.TYPE_SMALL) {
		if len(sig.Path) != c.ProtocolData.TREE_NODES_TO_STORE {
			valid_signature = false
		}
		if len(sig.Proof) != c.ProtocolData.TREE_NODES_TO_STORE {
			valid_signature = false
		}
	} else {
		if len(sig.Path) != c.ProtocolData.W {
			valid_signature = false
		}
		if len(sig.Proof) != c.ProtocolData.W {
			valid_signature = false
		}
	}
	if len(sig.Resp_0) != c.ProtocolData.T-c.ProtocolData.W {
		for i := 0; i < c.ProtocolData.T-c.ProtocolData.W; i++ {
			if len(sig.Resp_0[i].Y) != c.DenselyPackedFpVecSize() {
				valid_signature = false
			}
			if c.ProtocolData.Variant() == common.VARIANT_RSDP {
				if len(sig.Resp_0[i].V_bar) != c.DenselyPackedFzVecSize() {
					valid_signature = false
				}
			} else {
				if len(sig.Resp_0[i].V_G_bar) != c.DenselyPackedFzRSDPGVecSize() {
					valid_signature = false
				}
			}
		}
	}
	if len(sig.Resp_1) != (c.ProtocolData.T-c.ProtocolData.W)*((2*c.ProtocolData.Lambda)/8) {
		valid_signature = false
	}
	if !valid_signature {
		return false, fmt.Errorf("invalid signature")
	}
	V_tr, W_mat, err := c.Expand_pk(pk.SeedPK)
	if err != nil {
		return false, err
	}
	s, is_padd_key_ok := c.Unpack_fp_syn(pk.S)
	digest_msg_cmt_salt := make([]byte, 3*(2*c.ProtocolData.Lambda/8))
	hash_val, err := c.CSPRNG(m, 2*c.ProtocolData.Lambda/8, uint16(32768))
	if err != nil {
		return false, err
	}
	copy(digest_msg_cmt_salt, hash_val)
	copy(digest_msg_cmt_salt[(2*c.ProtocolData.Lambda/8):], sig.Digest_cmt)
	copy(digest_msg_cmt_salt[2*(2*c.ProtocolData.Lambda/8):], sig.Salt)
	digest_chall_1, err := c.CSPRNG(digest_msg_cmt_salt, 2*c.ProtocolData.Lambda/8, uint16(32768))
	if err != nil {
		return false, err
	}
	chall_1, err := c.CSPRNG_fp_vec_chall_1(digest_chall_1)
	if err != nil {
		return false, err
	}
	chall_2, err := c.Expand_digest_to_fixed_weight(sig.Digest_chall_2)
	if err != nil {
		return false, err
	}
	round_seeds, is_stree_padding_ok, err := c.RebuildLeaves(sig.Path, sig.Salt, chall_2)
	if err != nil {
		return false, err
	}
	var cmt_0_i_input []byte
	if c.ProtocolData.Variant() == common.VARIANT_RSDP {
		cmt_0_i_input = make([]byte, int(c.DenselyPackedFpSynSize())+c.DenselyPackedFzVecSize()+(2*c.ProtocolData.Lambda/8))
		copy(cmt_0_i_input[int(c.DenselyPackedFpSynSize())+c.DenselyPackedFzVecSize():], sig.Salt)
	} else {
		cmt_0_i_input = make([]byte, int(c.DenselyPackedFpSynSize())+c.DenselyPackedFzRSDPGVecSize()+(2*c.ProtocolData.Lambda/8))
		copy(cmt_0_i_input[int(c.DenselyPackedFpSynSize())+c.DenselyPackedFzRSDPGVecSize():], sig.Salt)

	}
	//remember to add salt, makes every dish more tasty
	cmt_1_i_input := make([]byte, 3*c.ProtocolData.Lambda/8)
	copy(cmt_1_i_input[c.ProtocolData.Lambda/8:], sig.Salt)
	cmt_0 := make([][]byte, c.ProtocolData.T)
	for i := 0; i < c.ProtocolData.T; i++ {
		cmt_0[i] = make([]byte, 2*c.ProtocolData.Lambda/8)
	}
	cmt_1 := make([]byte, c.ProtocolData.T*(2*c.ProtocolData.Lambda/8))
	e_bar_prime := make([]byte, c.ProtocolData.N)
	u_prime := make([]T, c.ProtocolData.N)
	y_prime := make([]T, c.ProtocolData.N)
	y_prime_H := make([]T, c.ProtocolData.N-c.ProtocolData.K)
	s_prime := make([]T, c.ProtocolData.N-c.ProtocolData.K)
	y := make([]T, c.ProtocolData.T*c.ProtocolData.N)
	v_bar := make([]byte, c.ProtocolData.N)
	used_rsps := 0
	is_signature_ok := true
	is_packed_padd_ok := true
	for i := 0; i < c.ProtocolData.T; i++ {
		domain_sep_csprng := uint16(0 + i + (2*c.ProtocolData.T - 1))
		domain_sep_hash := uint16(32768 + i + (2*c.ProtocolData.T - 1))
		if chall_2[i] {
			copy(cmt_1_i_input, round_seeds[i])
			cmt_1_val, err := c.CSPRNG(cmt_1_i_input, 2*c.ProtocolData.Lambda/8, domain_sep_hash)
			if err != nil {
				return false, err
			}
			copy(cmt_1[i*(2*c.ProtocolData.Lambda/8):], cmt_1_val)
			csprng_input := make([]byte, 3*c.ProtocolData.Lambda/8)
			copy(csprng_input, round_seeds[i])
			copy(csprng_input[c.ProtocolData.Lambda/8:], sig.Salt)
			state, err := c.CSPRNG_init(csprng_input, domain_sep_csprng)
			if err != nil {
				return false, err
			}
			if c.ProtocolData.Variant() == common.VARIANT_RSDP {
				e_bar_prime, _, err = c.CSPRNG_fz_vec_prime(state)
				if err != nil {
					return false, err
				}
			} else {
				e_G_bar_prime, _, err := c.CSPRNG_fz_inf_w_prime(state)
				if err != nil {
					return false, err
				}
				e_bar_prime = c.Fz_inf_w_by_fz_matrix(e_G_bar_prime, W_mat)
				e_bar_prime = c.Fz_dz_norm_n(e_bar_prime)
			}
			u_prime, err = c.CSPRNG_fp_vec_prime(state)
			if err != nil {
				return false, err
			}
			copy(y[i*c.ProtocolData.N:], c.Fp_vec_by_restr_vec_scaled(c.byteToT(e_bar_prime), u_prime, chall_1[i]))
			copy(y[i*c.ProtocolData.N:(i+1)*c.ProtocolData.N], c.Fp_dz_norm(y[i*c.ProtocolData.N:(i+1)*c.ProtocolData.N]))
		} else {
			temp_val, bool_res := c.Unpack_fp_vec(sig.Resp_0[used_rsps].Y)
			copy(y[i*c.ProtocolData.N:], temp_val)
			is_packed_padd_ok = is_packed_padd_ok && bool_res
			if c.ProtocolData.Variant() == common.VARIANT_RSDP {
				v_bar, bool_res = c.Unpack_fz_vec(sig.Resp_0[used_rsps].V_bar)
				is_packed_padd_ok = is_packed_padd_ok && bool_res
				copy(cmt_0_i_input[c.DenselyPackedFpSynSize():], sig.Resp_0[used_rsps].V_bar)
				is_signature_ok = is_signature_ok && c.Is_fz_vec_in_restr_group_n(v_bar)
			} else {
				copy(cmt_0_i_input[c.DenselyPackedFpSynSize():], sig.Resp_0[used_rsps].V_G_bar)
				v_G_bar, padd_bool := c.Unpack_fz_rsdp_g_vec(sig.Resp_0[used_rsps].V_G_bar)
				is_packed_padd_ok = is_packed_padd_ok && padd_bool
				is_signature_ok = is_signature_ok && c.Is_fz_vec_in_restr_group_m(v_G_bar)
				v_bar = c.Fz_inf_w_by_fz_matrix(v_G_bar, W_mat)
			}
			copy(cmt_1[i*(2*c.ProtocolData.Lambda/8):], sig.Resp_1[used_rsps*(2*c.ProtocolData.Lambda/8):(used_rsps+1)*(2*c.ProtocolData.Lambda/8)])
			used_rsps++
			v := c.Convert_restr_vec_to_fp(v_bar)
			y_prime = c.Fp_vec_by_fp_vec_pointwise(v, y[i*c.ProtocolData.N:(i+1)*c.ProtocolData.N])
			y_prime_H = c.Fp_vec_by_fp_matrix(y_prime, c.int16ToT(V_tr))
			y_prime_H = c.Fp_dz_norm_synd(y_prime_H)
			s_prime = c.Fp_synd_minus_fp_vec_scaled(y_prime_H, chall_1[i], s)
			s_prime = c.Fp_dz_norm_synd(s_prime)
			copy(cmt_0_i_input, c.Pack_fp_syn(s_prime))
			hash_val, err = c.CSPRNG(cmt_0_i_input, 2*c.ProtocolData.Lambda/8, domain_sep_hash)
			if err != nil {
				return false, err
			}
			cmt_0[i] = hash_val
		}
	}
	digest_cmt_0_cmt_1 := make([]byte, 2*(2*c.ProtocolData.Lambda/8))
	digest_val, is_mtree_padding_ok, err := c.RecomputeRoot(cmt_0, sig.Proof, chall_2)
	// TODO: set this to recomputeroot's result
	if err != nil {
		return false, err
	}
	copy(digest_cmt_0_cmt_1, digest_val)
	digest_hash_val, err := c.CSPRNG(cmt_1, 2*c.ProtocolData.Lambda/8, uint16(32768))
	if err != nil {
		return false, err
	}
	copy(digest_cmt_0_cmt_1[2*c.ProtocolData.Lambda/8:], digest_hash_val)
	digest_cmt_prime, err := c.CSPRNG(digest_cmt_0_cmt_1, 2*c.ProtocolData.Lambda/8, uint16(32768))
	if err != nil {
		return false, err
	}
	y_digest_chall_1 := make([]byte, c.ProtocolData.T*c.DenselyPackedFpVecSize()+(2*c.ProtocolData.Lambda/8))
	for x := 0; x < c.ProtocolData.T; x++ {
		copy(y_digest_chall_1[x*c.DenselyPackedFpVecSize():(x+1)*c.DenselyPackedFpVecSize()], c.Pack_fp_vec(y[x*c.ProtocolData.N:(x+1)*c.ProtocolData.N]))
	}
	copy(y_digest_chall_1[c.ProtocolData.T*c.DenselyPackedFpVecSize():], digest_chall_1)
	digest_chall_2_prime, err := c.CSPRNG(y_digest_chall_1, 2*c.ProtocolData.Lambda/8, uint16(32768))
	if err != nil {
		return false, err
	}
	does_digest_cmt_match := bytes.Equal(digest_cmt_prime, sig.Digest_cmt)
	does_digest_chall_2_match := bytes.Equal(digest_chall_2_prime, sig.Digest_chall_2)
	fmt.Println("is_signature_ok: ", is_signature_ok)
	fmt.Println("does_digest_cmt_match: ", does_digest_cmt_match)
	fmt.Println("does_digest_chall_2_match: ", does_digest_chall_2_match)
	fmt.Println("is_mtree_padding_ok: ", is_mtree_padding_ok)
	fmt.Println("is_stree_padding_ok: ", is_stree_padding_ok)
	fmt.Println("is_padd_key_ok: ", is_padd_key_ok)
	fmt.Println("is_packed_padd_ok: ", is_packed_padd_ok)

	is_signature_ok = is_signature_ok &&
		does_digest_cmt_match &&
		does_digest_chall_2_match &&
		is_mtree_padding_ok &&
		is_stree_padding_ok &&
		is_padd_key_ok &&
		is_packed_padd_ok
	return is_signature_ok, nil
}
