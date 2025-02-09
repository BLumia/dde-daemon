/*
 * Copyright (C) 2014 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package appinfo

// #cgo pkg-config: glib-2.0
// #cgo CFLAGS: -W -Wall -fstack-protector-all -fPIC
// #include <glib.h>
import "C"
import (
	"os"
	"path"

	"github.com/linuxdeepin/go-gir/glib-2.0"
	"github.com/linuxdeepin/go-lib/utils"
)

const (
	_DirDefaultPerm os.FileMode = 0755
)

// ConfigFilePath returns path in user's config dir.
func ConfigFilePath(name string) string {
	return path.Join(glib.GetUserConfigDir(), name)
}

// ConfigFile open the given keyfile, this file will be created if not existed.
func ConfigFile(name string) (*glib.KeyFile, error) {
	file := glib.NewKeyFile()
	conf := ConfigFilePath(name)
	if !utils.IsFileExist(conf) {
		_ = os.MkdirAll(path.Dir(conf), _DirDefaultPerm)
		f, err := os.Create(conf)
		if err != nil {
			return nil, err
		}
		defer func() {
			_ = f.Close()
		}()
	}

	if ok, err := file.LoadFromFile(conf, glib.KeyFileFlagsNone); !ok {
		file.Free()
		return nil, err
	}
	return file, nil
}
