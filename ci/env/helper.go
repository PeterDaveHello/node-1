/*
 * Copyright (C) 2019 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package env

import "github.com/cihub/seelog"

// IsRelease indicates if it is any kind of release (snapshot/tag)
func IsRelease() bool {
	isPullRequest, _ := RequiredEnvBool(PrBuild)
	return isPullRequest
}

// IfRelease performs func passed as an arg if `IsRelease() == true`
func IfRelease(do func() error) error {
	if IsRelease() {
		return do()
	}
	seelog.Info("Not a release build, skipping action")
	return nil
}
