/*
 * Copyright Greg Haskins All Rights Reserved
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 * See github.com/Alex-duzhichao/fabric/test/chaincodes/AutoVendor/chaincode/main.go for details
 */
package directdep

import (
	"github.com/Alex-duzhichao/fabric/test/chaincodes/AutoVendor/indirectdep"
)

func PointlessFunction() {
	// delegate to our indirect dependency
	indirectdep.PointlessFunction()
}
