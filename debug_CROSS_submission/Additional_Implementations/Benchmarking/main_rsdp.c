#include "csprng_hash.h"
#include <stdio.h>
#include <stdint.h>
/*
void print_csprng_state(const CSPRNG_STATE_T *csprng_state, size_t size) {
    for (size_t i = 0; i < size; i++) {
        printf("%u ", ((unsigned char*)csprng_state)[i]);
    }
    printf("\n");
}

void print_seed_array(uint8_t seed_e_seed_pk[2][KEYPAIR_SEED_LENGTH_BYTES]) {
    printf("seed_e_seed_pk:\n");
    for (int i = 0; i < 2; i++) {
        printf("Row %d: ", i);
        for (int j = 0; j < KEYPAIR_SEED_LENGTH_BYTES; j++) {
            printf("%u, ", seed_e_seed_pk[i][j]);
        }
        printf("\n");
    }
}

void print_hash(const uint8_t *digest, size_t length) {
    printf("Digest message: ");
    for (size_t i = 0; i < length; i++) {
        printf("%u, ", digest[i]); 
    }
    printf("\n");
}

void test_csprng(){
    printf("Testing csprng\n");
    const uint16_t dsc_csprng_seed_pk = CSPRNG_DOMAIN_SEP_CONST + (3*T+2);
    printf("T: %u\n", T);
    printf("dsc_csprng_seed_pk: %u\n", dsc_csprng_seed_pk);
    printf("KEYPAIR_SEED_LENGTH_BYTES: %u\n", KEYPAIR_SEED_LENGTH_BYTES);
    CSPRNG_STATE_T csprng_state_mat;
    //32 A's
    //const char * restrict seed_sk = "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA";
    const char * restrict seed_sk = "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA";
    csprng_initialize(&csprng_state_mat, seed_sk, KEYPAIR_SEED_LENGTH_BYTES, dsc_csprng_seed_pk);
    //printf("Hash generated\n");
    //print_csprng_state(&csprng_state_mat, sizeof(csprng_state_mat));
    //print_csprng_state(&csprng_state_mat, 32);
    uint8_t seed_e_seed_pk[2][KEYPAIR_SEED_LENGTH_BYTES];
    print_seed_array(seed_e_seed_pk);
    csprng_randombytes((uint8_t *)seed_e_seed_pk, 2*KEYPAIR_SEED_LENGTH_BYTES, &csprng_state_mat);
    print_seed_array(seed_e_seed_pk);
    return;
}

void test_hash(){
    printf("Testing hash\n");
    uint8_t digest_msg_cmt_salt[HASH_DIGEST_LENGTH] = {0};
    printf("HASH DOMAIN SEP CONST: %u\n", HASH_DOMAIN_SEP_CONST);
    printf("SALT LENGTH BYTES: %u\n", SALT_LENGTH_BYTES);
    const char *const m = "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA";
    const uint64_t mlen = 32;
    /* place digest_msg at the beginning of the input of the hash generating digest_chall_1 */
    print_hash(digest_msg_cmt_salt, HASH_DIGEST_LENGTH);
    hash(digest_msg_cmt_salt, (uint8_t*) m, mlen, HASH_DOMAIN_SEP_CONST);
    print_hash(digest_msg_cmt_salt, HASH_DIGEST_LENGTH);
    //print_csprng_state(digest_msg_cmt_salt, 2*HASH_DIGEST_LENGTH+SALT_LENGTH_BYTES);
    return;
}

void test_keygen(){
    printf("Testing keygen\n");
    
}

void print_V_tr(FP_ELEM V_tr[K][N-K]){
    printf("V_tr:\n");
    for (int i = 0; i < K; i++) {
        for (int j = 0; j < N-K; j++) {
            printf("%u, ", V_tr[i][j]);
        }
        printf("\n");
    }
    return;
}

void test_csprng_fp_mat(){
    FP_ELEM V_tr[K][N-K];
    printf("Testing CSPRNG - fp_mat\n");
    const uint16_t dsc_csprng_seed_pk = CSPRNG_DOMAIN_SEP_CONST + (3*T+2);
    CSPRNG_STATE_T csprng_state_mat;
    const char * restrict seed_pk = "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA";
    csprng_initialize(&csprng_state_mat, seed_pk, KEYPAIR_SEED_LENGTH_BYTES, dsc_csprng_seed_pk);
    //print_csprng_state(&csprng_state_mat, sizeof(csprng_state_mat));
    csprng_fp_mat(V_tr,&csprng_state_mat);
    print_V_tr(V_tr);
    return;
}

void print_z_vec(FP_ELEM z[N]){
    printf("z:\n");
    for (int i = 0; i < N; i++) {
        printf("%u, ", z[i]);
    }
    printf("\n");
    return;
}

void test_csprng_fz_vec(){
    printf("Testing CSPRNG - fz_vec\n");
    FZ_ELEM z[N];
    const uint16_t dsc_csprng_seed_pk = CSPRNG_DOMAIN_SEP_CONST + (3*T+3);
    CSPRNG_STATE_T csprng_state_mat;
    const char * restrict seed_pk = "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA";
    csprng_initialize(&csprng_state_mat, seed_pk, KEYPAIR_SEED_LENGTH_BYTES, dsc_csprng_seed_pk);
    //print_csprng_state(&csprng_state_mat, sizeof(csprng_state_mat));
    csprng_fz_vec(z,&csprng_state_mat);
    print_z_vec(z);
    return;
}

void test_csprng_fp_vec(){
    printf("Testing CSPRNG - fp_vec\n");
    FP_ELEM z[N];
    const uint16_t dsc_csprng_seed_pk = CSPRNG_DOMAIN_SEP_CONST + (2*T-1);
    CSPRNG_STATE_T csprng_state_mat;
    const char * restrict seed_pk = "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA";
    csprng_initialize(&csprng_state_mat, seed_pk, KEYPAIR_SEED_LENGTH_BYTES, dsc_csprng_seed_pk);
    //print_csprng_state(&csprng_state_mat, sizeof(csprng_state_mat));
    csprng_fp_vec(z,&csprng_state_mat);
    print_z_vec(z);
    return;
}

void print_chall_1(FP_ELEM chall_1[T]){
    printf("chall_1:\n");
    for (int i = 0; i < T; i++) {
        printf("%u, ", chall_1[i]);
    }
    printf("\n");
    return;
}

void test_csprng_fp_vec_chall_1(){
    printf("Testing CSPRNG - fp_vec_chall_1\n");
    FP_ELEM chall_1[T];
    const uint16_t dsc_csprng_seed_pk = CSPRNG_DOMAIN_SEP_CONST + (3*T-1);
    CSPRNG_STATE_T csprng_state_mat;
    const char * restrict seed_pk = "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA";
    csprng_initialize(&csprng_state_mat, seed_pk, KEYPAIR_SEED_LENGTH_BYTES, dsc_csprng_seed_pk);
    //print_csprng_state(&csprng_state_mat, sizeof(csprng_state_mat));
    csprng_fp_vec_chall_1(chall_1,&csprng_state_mat);
    print_chall_1(chall_1);
    return;
}

void print_e_G_bar(FZ_ELEM e_G_bar[M]){
    printf("e_G_bar:\n");
    for (int i = 0; i < M; i++) {
        printf("%u, ", e_G_bar[i]);
    }
    printf("\n");
    return;
}

//Reminder, this test only works in RSDP-G
void test_csprng_fz_inf_w(){
    printf("Testing CSPRNG - fz_inf_w\n");
    FZ_ELEM e_G_bar[M];
    const uint16_t dsc_csprng_seed_pk = CSPRNG_DOMAIN_SEP_CONST + (3*T+3);
    CSPRNG_STATE_T csprng_state_mat;
    const char * restrict seed_pk = "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA";
    csprng_initialize(&csprng_state_mat, seed_pk, KEYPAIR_SEED_LENGTH_BYTES, dsc_csprng_seed_pk);
    //print_csprng_state(&csprng_state_mat, sizeof(csprng_state_mat));
    csprng_fz_inf_w(e_G_bar,&csprng_state_mat);
    print_e_G_bar(e_G_bar);
    return;
}*/


int main() {
    //test_hash();
    //test_csprng();
    //test_csprng_fp_mat();
    //test_csprng_fz_vec();
    //test_csprng_fp_vec();
    //test_csprng_fp_vec_chall_1();
    return 0;
}

//Row 0: 15 121 106 185 65 60 38 57 192 11 100 5 36 234 50 253 115 61 99 71 54 20 106 223 64 83 75 131 107 171 179 163 
//Row 1: 197 184 200 221 6 37 92 70 124 127 54 125 11 163 142 207 26 21 208 178 226 28 152 49 104 87 51 136 32 87 109 243 


