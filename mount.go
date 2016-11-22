// +build linux darwin

/*
 * btrfscue version 0.3
 * Copyright (c)2011-2016 Christian Blichmann
 *
 * Sub-command to provide and mount a "rescue fs"
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *     * Redistributions of source code must retain the above copyright
 *       notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above copyright
 *       notice, this list of conditions and the following disclaimer in the
 *       documentation and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
 * LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 * CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
 * SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 * INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
 * CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 * ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 */

package main

import (
	"flag"

	"blichmann.eu/code/btrfscue/btrfs/index"
	"blichmann.eu/code/btrfscue/rescuefs"
	"blichmann.eu/code/btrfscue/subcommand"
)

type mountCommand struct {
}

func (c *mountCommand) DefineFlags(fs *flag.FlagSet) {
}

func (c *mountCommand) Run(args []string) {
	if len(args) == 0 {
		fatalf("missing mount point\n")
	}
	if len(args) > 1 {
		fatalf("extra operand '%s'\n", args[1])
	}
	if len(*metadata) == 0 {
		fatalf("missing metadata option\n")
	}

	ix, err := index.OpenReadOnly(*metadata)
	reportError(err)
	defer ix.Close()

	fs := rescuefs.New(*metadata, ix, nil)
	reportError(fs.Mount(args[0]))
	fs.Serve()
}

func init() {
	subcommand.Register("mount",
		"provide a 'rescue' filesystem backed by metadata", &mountCommand{})
}
