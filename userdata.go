// Copyright 2017 Daniel Salvadori. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/gob"
	"os"
)

// The filepath of the file used to store the UserData instance via Gob
const USER_DATA_FILEPATH string = "user.data"

// UserData stores all the information that persists between game sessions
type UserData struct {
	MusicOn           bool
	SfxOn             bool
	MusicVol          float32
	SfxVol            float32
	FullScreen        bool
	LastLevel         int
	LastUnlockedLevel int
}

// NewUserData loads user data from file or creates a new object with default values if no file exists
func NewUserData() *UserData {
	ud := new(UserData)

	// Try to read existing file
	file, err := os.Open(USER_DATA_FILEPATH)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(ud)
		if err != nil {
			log.Debug("Error decoding user.data: %v", err)
		}
	} else {
		log.Debug("Error opening user.data file: %v", err)
	}
	file.Close()

	if err != nil {
		ud.SfxOn = true
		ud.MusicOn = true
		ud.SfxVol = 0.5
		ud.MusicVol = 0.5
		ud.FullScreen = false
		ud.LastLevel = 0
		ud.LastUnlockedLevel = 0
		log.Debug("Creating new user data with default values: %+v", ud)
	} else {
		log.Debug("Loaded user data: %+v", ud)
	}

	return ud
}

// Save saves the current user data to the user data file, overwriting existing (old) data
func (ud *UserData) Save() error {
	log.Debug("Saving user data: %+v", ud)
	newFile, err := os.Create(USER_DATA_FILEPATH)
	if err == nil {
		encoder := gob.NewEncoder(newFile)
		err = encoder.Encode(&ud)
		if err != nil {
			log.Debug("Error encoding user.data: %v", err)
		}
	} else {
		log.Debug("Error creating user.data file: %v", err)
	}
	newFile.Close()
	return err
}
