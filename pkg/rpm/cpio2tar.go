/*
 * Copyright (c) SAS Institute, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rpm

import (
	"archive/tar"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"github.com/sassoftware/go-rpmutils/cpio"
)

const (
	capabilities_header = "SCHILY.xattr.security.capability"
)

var cap_empty_bitmask = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
var supported_capabilities = map[string][]byte{
	"cap_net_bind_service": {1, 0, 0, 2, 0, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}

// Extract the contents of a cpio stream from and writes it as a tar file into the provided writer
func Tar(rs io.Reader, tarfile *tar.Writer, noSymlinksAndDirs bool, capabilities map[string][]string) error {
	hardLinks := map[int][]*tar.Header{}
	inodes := map[int]string{}

	stream := cpio.NewCpioStream(rs)

	for {
		entry, err := stream.ReadNextEntry()
		if err != nil {
			return err
		}

		if entry.Header.Filename() == cpio.TRAILER {
			break
		}

		pax := map[string]string{}
		if caps, exists := capabilities[entry.Header.Filename()]; exists {
			for _, cap := range caps {
				if _, supported := supported_capabilities[cap]; !supported {
					return fmt.Errorf("Requested capability '%s' for file '%s' is not supported", cap, entry.Header.Filename())
				}
				if _, exists := pax[capabilities_header]; !exists {
					pax[capabilities_header] = string(cap_empty_bitmask)
				}
				val := []byte(pax[capabilities_header])
				for i, b := range supported_capabilities[cap] {
					val[i] = val[i] | b
				}
				pax[capabilities_header] = string(val)
			}
		}

		tarHeader := &tar.Header{
			Name:       entry.Header.Filename(),
			Size:       entry.Header.Filesize64(),
			Mode:       int64(entry.Header.Mode()),
			Uid:        entry.Header.Uid(),
			Gid:        entry.Header.Gid(),
			ModTime:    time.Unix(int64(entry.Header.Mtime()), 0),
			Devmajor:   int64(entry.Header.Devmajor()),
			Devminor:   int64(entry.Header.Devminor()),
			PAXRecords: pax,
		}

		var payload io.Reader
		switch entry.Header.Mode() &^ 07777 {
		case cpio.S_ISCHR:
			tarHeader.Typeflag = tar.TypeChar
		case cpio.S_ISBLK:
			tarHeader.Typeflag = tar.TypeBlock
		case cpio.S_ISDIR:
			if noSymlinksAndDirs {
				continue
			}
			tarHeader.Typeflag = tar.TypeDir
		case cpio.S_ISFIFO:
			tarHeader.Typeflag = tar.TypeFifo
		case cpio.S_ISLNK:
			if noSymlinksAndDirs {
				continue
			}
			tarHeader.Typeflag = tar.TypeSymlink
			tarHeader.Size = 0
			buf, err := ioutil.ReadAll(entry.Payload)
			if err != nil {
				return err
			}
			tarHeader.Linkname = string(buf)
		case cpio.S_ISREG:
			if entry.Header.Nlink() > 1 && entry.Header.Filesize() == 0 {
				tarHeader.Typeflag = tar.TypeLink
				tarHeader.Size = 0
				hardLinks[entry.Header.Ino()] = append(hardLinks[entry.Header.Ino()], tarHeader)
				continue
			}
			tarHeader.Typeflag = tar.TypeReg
			payload = entry.Payload
			inodes[entry.Header.Ino()] = entry.Header.Filename()
		default:
			return fmt.Errorf("unknown file mode 0%o for %s",
				entry.Header.Mode(), entry.Header.Filename())
		}
		if err := tarfile.WriteHeader(tarHeader); err != nil {
			return fmt.Errorf("could not write tar header for %v: %v", tarHeader.Name, err)
		}
		if payload != nil {
			written, err := io.Copy(tarfile, entry.Payload)
			if err != nil {
				return fmt.Errorf("could not write body for %v: %v", tarHeader.Name, err)
			}
			if written != int64(entry.Header.Filesize()) {
				return fmt.Errorf("short write body for %v", tarHeader.Name)
			}
		}
	}
	// write hardlinks
	for node, links := range hardLinks {
		target := inodes[node]
		if target == "" {
			return fmt.Errorf("no target file for inode %v found", node)
		}
		for _, tarHeader := range links {
			tarHeader.Linkname = target
			if err := tarfile.WriteHeader(tarHeader); err != nil {
				return fmt.Errorf("could not write tar header for %v", tarHeader.Name)
			}
		}
	}

	return nil
}
