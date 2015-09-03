package arm

import (
	cs "github.com/bnagy/gapstone"
	uc "github.com/unicorn-engine/unicorn/bindings/go/unicorn"

	"../../models"
)

var Arch = &models.Arch{
	Bits:    32,
	Radare:  "arm",
	CS_ARCH: cs.CS_ARCH_ARM,
	CS_MODE: cs.CS_MODE_ARM,
	UC_ARCH: uc.UC_ARCH_ARM,
	UC_MODE: uc.UC_MODE_ARM,
	SP:      uc.UC_ARM_REG_SP,
}