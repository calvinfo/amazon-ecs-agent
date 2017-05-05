// +build !linux

// Copyright 2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package udevwrapper

// All interfaces, structs and methods exposed here are placeholders
// for build purposes on unsupported platforms

import (
	"fmt"
	"runtime"
)

// uEventUnsupported is a placeholder for build purposes on unsupported platforms
type uEventUnsupported struct{}

// UdevMonitorUnsupported is a placeholder for build purposes on unsupported platforms
type UdevMonitorUnsupported struct{}

// Udev placeholer interface for build purposes on unsupported platforms
type Udev interface {
	Monitor(notify chan *uEventUnsupported) (shutdown chan bool)
	Close() error
}

// New is a placeholder that returns an error on unsupported platforms
func New() (*UdevMonitorUnsupported, error) {
	return nil, fmt.Errorf("udev monitor new: unsupported platform: %s/%s", runtime.GOOS, runtime.GOARCH)
}

// Monitor is a placeholder for unsupported platforms
func (UdevMonitorUnsupported) Monitor(notify chan *uEventUnsupported) (shutdown chan bool) {
	return make(chan bool)
}

// Close is a placeholder for unsupported platforms
func (UdevMonitorUnsupported) Close() error {
	return fmt.Errorf("udev monitor close: unsupported platform: %s/%s", runtime.GOOS, runtime.GOARCH)
}