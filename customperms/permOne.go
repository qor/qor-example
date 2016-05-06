package customperms

import (
	"fmt"
	"log"

	"github.com/qor/roles"
)

type PermTest struct{}

//
func (p *PermTest) Allow(mode roles.PermissionMode, roles ...string) roles.Permission {
	panic(fmt.Sprintf("permTest: Allow(%s,%v)", mode, roles))
}

func (p *PermTest) Concat(roles.Permission) roles.Permission {
	panic("PermTest: Concat()")
}

func (p *PermTest) Deny(mode roles.PermissionMode, roles ...string) roles.Permission {
	panic(fmt.Sprintf("permTest: Deny(%s,%v)", mode, roles))
}

func (p *PermTest) HasPermission(mode roles.PermissionMode, roles ...string) bool {
	log.Printf("DBG: permTest: HasPermission(%v, %v)", mode, roles)
	return true
}

func (p *PermTest) Role() roles.Role {
	panic("permTest: Role()")
}
