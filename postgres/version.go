/* Copyright (c) 2016 Jason Ish
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 *
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED ``AS IS'' AND ANY EXPRESS OR IMPLIED
 * WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT,
 * INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 * (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
 * SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT,
 * STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING
 * IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 */

package postgres

import (
	"fmt"
	"github.com/pkg/errors"
	"regexp"
	"strconv"
	"strings"
)

type PostgresVersion struct {
	Raw        string
	Full       string
	MajorMinor string
	Major      int64
	Minor      int64
}

func (v *PostgresVersion) String() string {
	return v.Raw
}

func ParseVersion(versionString string) (*PostgresVersion, error) {
	re := regexp.MustCompile("(\\d+)\\.(\\d+)\\.\\d+")
	parts := re.FindStringSubmatch(versionString)
	if parts == nil || len(parts) != 3 {
		return nil, errors.Errorf("failed to parse PostgreSQL version: %s", versionString)
	}

	major, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return nil, errors.Errorf("failed to parse PostgreSQL version: %s", versionString)
	}

	minor, err := strconv.ParseInt(parts[2], 10, 64)
	if err != nil {
		return nil, errors.Errorf("failed to parse PostgreSQL version: %s", versionString)
	}

	return &PostgresVersion{
		Raw:        strings.TrimSpace(versionString),
		Full:       parts[0],
		MajorMinor: fmt.Sprintf("%d.%d", major, minor),
		Major:      major,
		Minor:      minor,
	}, nil
}