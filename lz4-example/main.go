/*
 * Copyright 2011 Branimir Karadzic. All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without modification,
 * are permitted provided that the following conditions are met:
 *
 *    1. Redistributions of source code must retain the above copyright notice, this
 *       list of conditions and the following disclaimer.
 *
 *    2. Redistributions in binary form must reproduce the above copyright notice,
 *       this list of conditions and the following disclaimer in the documentation
 *       and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY COPYRIGHT HOLDER ``AS IS'' AND ANY EXPRESS OR
 * IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT
 * SHALL COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
 * INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 * LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
 * PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
 * WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE
 * OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF
 * THE POSSIBILITY OF SUCH DAMAGE.
 */

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bkaradzic/go-lz4"
)

var (
	decompress = flag.Bool("d", false, "decompress")
)

func main() {

	flag.Parse()

	args := flag.Args()

	var data []byte

	if len(args) < 2 {
		fmt.Print("Usage: lz4 [-d] <input> <output>\n")
		os.Exit(1)
	}

	input, err := os.OpenFile(args[0], os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open input file %s\n", args[0])
		os.Exit(1)
	}
	defer input.Close()

	if *decompress {
		lz4r := lz4.NewReader(input)
		defer lz4r.Close()
		data, _ = ioutil.ReadAll(lz4r)
	} else {
		lz4w := lz4.NewWriter(input)
		defer lz4w.Close()
		data, _ = ioutil.ReadAll(lz4w)
	}

	err = ioutil.WriteFile(args[1], data, 0644)
	if err != nil {
		fmt.Printf("Failed to open output file %s\n", args[1])
		os.Exit(1)
	}
}
