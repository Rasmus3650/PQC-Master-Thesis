[ ] = Not done yet

[X] = done

Code:
[ ] aux.go: Faster/better random sampling of bits

[ ] key_gen.go: Switch on SHAKE128/SHAKE256 based on security level, switch on G-variant


[X] Remove "nodes to store" parameter


[ ] Seed & Merkle: ComputeNodesToPublish should give bool tree rather than index list to avoid the O(n) contains check

[ ] Seed & Merkle: Make file for auxiliary functions for trees

[ ] Make the implementation class based (signer, verifier), this will allow us to remove all the params that use proto_params and tree_params

[ ] Domain seperator for HASH and CSPRNG calls (look paramaters - HASH_DOMAIN_SEP_CONST)

[ ] Make sure each CSPRNG call is calling the correct shakesum depending on the variant/problem type currently being run

[ ] Refactor to ensure all variables have well-defined lengths

[ ] Cleanup code (some functions are declared twice), and refactor 


Util:
Make tests