# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.31

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:

#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:

# Disable VCS-based implicit rules.
% : %,v

# Disable VCS-based implicit rules.
% : RCS/%

# Disable VCS-based implicit rules.
% : RCS/%,v

# Disable VCS-based implicit rules.
% : SCCS/s.%

# Disable VCS-based implicit rules.
% : s.%

.SUFFIXES: .hpux_make_needs_suffix_list

# Command-line flag to silence nested $(MAKE).
$(VERBOSE)MAKESILENT = -s

#Suppress display of executed commands.
$(VERBOSE).SILENT:

# A target that is always out of date.
cmake_force:
.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /usr/bin/cmake

# The command to remove a file.
RM = /usr/bin/cmake -E rm -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Additional_Implementations/Benchmarking

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Additional_Implementations/Benchmarking/build

# Include any dependencies generated for this target.
include CMakeFiles/cross_static.dir/depend.make
# Include any dependencies generated by the compiler for this target.
include CMakeFiles/cross_static.dir/compiler_depend.make

# Include the progress variables for this target.
include CMakeFiles/cross_static.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/cross_static.dir/flags.make

CMakeFiles/cross_static.dir/codegen:
.PHONY : CMakeFiles/cross_static.dir/codegen

CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/merkle.c.o: CMakeFiles/cross_static.dir/flags.make
CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/merkle.c.o: /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/merkle.c
CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/merkle.c.o: CMakeFiles/cross_static.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green --progress-dir=/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Additional_Implementations/Benchmarking/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building C object CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/merkle.c.o"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -MD -MT CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/merkle.c.o -MF CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/merkle.c.o.d -o CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/merkle.c.o -c /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/merkle.c

CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/merkle.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Preprocessing C source to CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/merkle.c.i"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/merkle.c > CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/merkle.c.i

CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/merkle.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Compiling C source to assembly CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/merkle.c.s"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/merkle.c -o CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/merkle.c.s

CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/seedtree.c.o: CMakeFiles/cross_static.dir/flags.make
CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/seedtree.c.o: /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/seedtree.c
CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/seedtree.c.o: CMakeFiles/cross_static.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green --progress-dir=/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Additional_Implementations/Benchmarking/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Building C object CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/seedtree.c.o"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -MD -MT CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/seedtree.c.o -MF CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/seedtree.c.o.d -o CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/seedtree.c.o -c /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/seedtree.c

CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/seedtree.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Preprocessing C source to CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/seedtree.c.i"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/seedtree.c > CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/seedtree.c.i

CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/seedtree.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Compiling C source to assembly CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/seedtree.c.s"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/seedtree.c -o CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/seedtree.c.s

CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/CROSS.c.o: CMakeFiles/cross_static.dir/flags.make
CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/CROSS.c.o: /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/CROSS.c
CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/CROSS.c.o: CMakeFiles/cross_static.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green --progress-dir=/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Additional_Implementations/Benchmarking/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_3) "Building C object CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/CROSS.c.o"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -MD -MT CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/CROSS.c.o -MF CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/CROSS.c.o.d -o CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/CROSS.c.o -c /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/CROSS.c

CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/CROSS.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Preprocessing C source to CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/CROSS.c.i"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/CROSS.c > CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/CROSS.c.i

CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/CROSS.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Compiling C source to assembly CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/CROSS.c.s"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/CROSS.c -o CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/CROSS.c.s

CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/keccakf1600.c.o: CMakeFiles/cross_static.dir/flags.make
CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/keccakf1600.c.o: /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/keccakf1600.c
CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/keccakf1600.c.o: CMakeFiles/cross_static.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green --progress-dir=/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Additional_Implementations/Benchmarking/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_4) "Building C object CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/keccakf1600.c.o"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -MD -MT CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/keccakf1600.c.o -MF CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/keccakf1600.c.o.d -o CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/keccakf1600.c.o -c /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/keccakf1600.c

CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/keccakf1600.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Preprocessing C source to CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/keccakf1600.c.i"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/keccakf1600.c > CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/keccakf1600.c.i

CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/keccakf1600.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Compiling C source to assembly CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/keccakf1600.c.s"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/keccakf1600.c -o CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/keccakf1600.c.s

CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/fips202.c.o: CMakeFiles/cross_static.dir/flags.make
CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/fips202.c.o: /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/fips202.c
CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/fips202.c.o: CMakeFiles/cross_static.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green --progress-dir=/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Additional_Implementations/Benchmarking/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_5) "Building C object CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/fips202.c.o"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -MD -MT CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/fips202.c.o -MF CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/fips202.c.o.d -o CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/fips202.c.o -c /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/fips202.c

CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/fips202.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Preprocessing C source to CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/fips202.c.i"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/fips202.c > CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/fips202.c.i

CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/fips202.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Compiling C source to assembly CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/fips202.c.s"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/fips202.c -o CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/fips202.c.s

CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/csprng_hash.c.o: CMakeFiles/cross_static.dir/flags.make
CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/csprng_hash.c.o: /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/csprng_hash.c
CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/csprng_hash.c.o: CMakeFiles/cross_static.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green --progress-dir=/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Additional_Implementations/Benchmarking/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_6) "Building C object CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/csprng_hash.c.o"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -MD -MT CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/csprng_hash.c.o -MF CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/csprng_hash.c.o.d -o CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/csprng_hash.c.o -c /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/csprng_hash.c

CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/csprng_hash.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Preprocessing C source to CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/csprng_hash.c.i"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/csprng_hash.c > CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/csprng_hash.c.i

CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/csprng_hash.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Compiling C source to assembly CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/csprng_hash.c.s"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/csprng_hash.c -o CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/csprng_hash.c.s

CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/pack_unpack.c.o: CMakeFiles/cross_static.dir/flags.make
CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/pack_unpack.c.o: /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/pack_unpack.c
CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/pack_unpack.c.o: CMakeFiles/cross_static.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green --progress-dir=/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Additional_Implementations/Benchmarking/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_7) "Building C object CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/pack_unpack.c.o"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -MD -MT CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/pack_unpack.c.o -MF CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/pack_unpack.c.o.d -o CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/pack_unpack.c.o -c /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/pack_unpack.c

CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/pack_unpack.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Preprocessing C source to CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/pack_unpack.c.i"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/pack_unpack.c > CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/pack_unpack.c.i

CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/pack_unpack.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Compiling C source to assembly CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/pack_unpack.c.s"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/pack_unpack.c -o CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/pack_unpack.c.s

CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/sign.c.o: CMakeFiles/cross_static.dir/flags.make
CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/sign.c.o: /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/sign.c
CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/sign.c.o: CMakeFiles/cross_static.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green --progress-dir=/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Additional_Implementations/Benchmarking/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_8) "Building C object CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/sign.c.o"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -MD -MT CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/sign.c.o -MF CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/sign.c.o.d -o CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/sign.c.o -c /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/sign.c

CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/sign.c.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Preprocessing C source to CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/sign.c.i"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -E /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/sign.c > CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/sign.c.i

CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/sign.c.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green "Compiling C source to assembly CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/sign.c.s"
	/usr/bin/cc $(C_DEFINES) $(C_INCLUDES) $(C_FLAGS) -S /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/sign.c -o CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/sign.c.s

# Object files for target cross_static
cross_static_OBJECTS = \
"CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/merkle.c.o" \
"CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/seedtree.c.o" \
"CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/CROSS.c.o" \
"CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/keccakf1600.c.o" \
"CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/fips202.c.o" \
"CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/csprng_hash.c.o" \
"CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/pack_unpack.c.o" \
"CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/sign.c.o"

# External object files for target cross_static
cross_static_EXTERNAL_OBJECTS =

libcross_static.a: CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/merkle.c.o
libcross_static.a: CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/seedtree.c.o
libcross_static.a: CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/CROSS.c.o
libcross_static.a: CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/keccakf1600.c.o
libcross_static.a: CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/fips202.c.o
libcross_static.a: CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/csprng_hash.c.o
libcross_static.a: CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/pack_unpack.c.o
libcross_static.a: CMakeFiles/cross_static.dir/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Reference_Implementation/lib/sign.c.o
libcross_static.a: CMakeFiles/cross_static.dir/build.make
libcross_static.a: CMakeFiles/cross_static.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color "--switch=$(COLOR)" --green --bold --progress-dir=/home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Additional_Implementations/Benchmarking/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_9) "Linking C static library libcross_static.a"
	$(CMAKE_COMMAND) -P CMakeFiles/cross_static.dir/cmake_clean_target.cmake
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/cross_static.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/cross_static.dir/build: libcross_static.a
.PHONY : CMakeFiles/cross_static.dir/build

CMakeFiles/cross_static.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/cross_static.dir/cmake_clean.cmake
.PHONY : CMakeFiles/cross_static.dir/clean

CMakeFiles/cross_static.dir/depend:
	cd /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Additional_Implementations/Benchmarking/build && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Additional_Implementations/Benchmarking /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Additional_Implementations/Benchmarking /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Additional_Implementations/Benchmarking/build /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Additional_Implementations/Benchmarking/build /home/rasmuso/Desktop/uni/PQC-Master-Thesis/debug_CROSS_submission/Additional_Implementations/Benchmarking/build/CMakeFiles/cross_static.dir/DependInfo.cmake "--color=$(COLOR)"
.PHONY : CMakeFiles/cross_static.dir/depend

