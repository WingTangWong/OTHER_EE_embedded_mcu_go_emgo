// +build f303xe

package flash

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f303xe/mmap"
)

type FLASH_Periph struct {
	ACR     RACR
	KEYR    RKEYR
	OPTKEYR ROPTKEYR
	SR      RSR
	CR      RCR
	AR      RAR
	_       uint32
	OBR     ROBR
	WRPR    RWRPR
}

func (p *FLASH_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var FLASH = (*FLASH_Periph)(unsafe.Pointer(uintptr(mmap.FLASH_R_BASE)))

type ACR uint32

func (b ACR) Field(mask ACR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ACR) J(v int) ACR {
	return ACR(bits.MakeField32(v, uint32(mask)))
}

type RACR struct{ mmio.U32 }

func (r *RACR) Bits(mask ACR) ACR     { return ACR(r.U32.Bits(uint32(mask))) }
func (r *RACR) StoreBits(mask, b ACR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RACR) SetBits(mask ACR)      { r.U32.SetBits(uint32(mask)) }
func (r *RACR) ClearBits(mask ACR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RACR) Load() ACR             { return ACR(r.U32.Load()) }
func (r *RACR) Store(b ACR)           { r.U32.Store(uint32(b)) }

func (r *RACR) AtomicStoreBits(mask, b ACR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RACR) AtomicSetBits(mask ACR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RACR) AtomicClearBits(mask ACR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMACR struct{ mmio.UM32 }

func (rm RMACR) Load() ACR   { return ACR(rm.UM32.Load()) }
func (rm RMACR) Store(b ACR) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) LATENCY() RMACR {
	return RMACR{mmio.UM32{&p.ACR.U32, uint32(LATENCY)}}
}

func (p *FLASH_Periph) HLFCYA() RMACR {
	return RMACR{mmio.UM32{&p.ACR.U32, uint32(HLFCYA)}}
}

func (p *FLASH_Periph) PRFTBE() RMACR {
	return RMACR{mmio.UM32{&p.ACR.U32, uint32(PRFTBE)}}
}

func (p *FLASH_Periph) PRFTBS() RMACR {
	return RMACR{mmio.UM32{&p.ACR.U32, uint32(PRFTBS)}}
}

type KEYR uint32

func (b KEYR) Field(mask KEYR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask KEYR) J(v int) KEYR {
	return KEYR(bits.MakeField32(v, uint32(mask)))
}

type RKEYR struct{ mmio.U32 }

func (r *RKEYR) Bits(mask KEYR) KEYR    { return KEYR(r.U32.Bits(uint32(mask))) }
func (r *RKEYR) StoreBits(mask, b KEYR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RKEYR) SetBits(mask KEYR)      { r.U32.SetBits(uint32(mask)) }
func (r *RKEYR) ClearBits(mask KEYR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RKEYR) Load() KEYR             { return KEYR(r.U32.Load()) }
func (r *RKEYR) Store(b KEYR)           { r.U32.Store(uint32(b)) }

func (r *RKEYR) AtomicStoreBits(mask, b KEYR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RKEYR) AtomicSetBits(mask KEYR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RKEYR) AtomicClearBits(mask KEYR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMKEYR struct{ mmio.UM32 }

func (rm RMKEYR) Load() KEYR   { return KEYR(rm.UM32.Load()) }
func (rm RMKEYR) Store(b KEYR) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) FKEYR() RMKEYR {
	return RMKEYR{mmio.UM32{&p.KEYR.U32, uint32(FKEYR)}}
}

type OPTKEYR uint32

func (b OPTKEYR) Field(mask OPTKEYR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OPTKEYR) J(v int) OPTKEYR {
	return OPTKEYR(bits.MakeField32(v, uint32(mask)))
}

type ROPTKEYR struct{ mmio.U32 }

func (r *ROPTKEYR) Bits(mask OPTKEYR) OPTKEYR { return OPTKEYR(r.U32.Bits(uint32(mask))) }
func (r *ROPTKEYR) StoreBits(mask, b OPTKEYR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ROPTKEYR) SetBits(mask OPTKEYR)      { r.U32.SetBits(uint32(mask)) }
func (r *ROPTKEYR) ClearBits(mask OPTKEYR)    { r.U32.ClearBits(uint32(mask)) }
func (r *ROPTKEYR) Load() OPTKEYR             { return OPTKEYR(r.U32.Load()) }
func (r *ROPTKEYR) Store(b OPTKEYR)           { r.U32.Store(uint32(b)) }

func (r *ROPTKEYR) AtomicStoreBits(mask, b OPTKEYR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *ROPTKEYR) AtomicSetBits(mask OPTKEYR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ROPTKEYR) AtomicClearBits(mask OPTKEYR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMOPTKEYR struct{ mmio.UM32 }

func (rm RMOPTKEYR) Load() OPTKEYR   { return OPTKEYR(rm.UM32.Load()) }
func (rm RMOPTKEYR) Store(b OPTKEYR) { rm.UM32.Store(uint32(b)) }

type SR uint32

func (b SR) Field(mask SR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR) J(v int) SR {
	return SR(bits.MakeField32(v, uint32(mask)))
}

type RSR struct{ mmio.U32 }

func (r *RSR) Bits(mask SR) SR      { return SR(r.U32.Bits(uint32(mask))) }
func (r *RSR) StoreBits(mask, b SR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSR) SetBits(mask SR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSR) ClearBits(mask SR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSR) Load() SR             { return SR(r.U32.Load()) }
func (r *RSR) Store(b SR)           { r.U32.Store(uint32(b)) }

func (r *RSR) AtomicStoreBits(mask, b SR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSR) AtomicSetBits(mask SR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSR) AtomicClearBits(mask SR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSR struct{ mmio.UM32 }

func (rm RMSR) Load() SR   { return SR(rm.UM32.Load()) }
func (rm RMSR) Store(b SR) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) BSY() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(BSY)}}
}

func (p *FLASH_Periph) PGERR() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(PGERR)}}
}

func (p *FLASH_Periph) WRPERR() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(WRPERR)}}
}

func (p *FLASH_Periph) EOP() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(EOP)}}
}

type CR uint32

func (b CR) Field(mask CR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR) J(v int) CR {
	return CR(bits.MakeField32(v, uint32(mask)))
}

type RCR struct{ mmio.U32 }

func (r *RCR) Bits(mask CR) CR      { return CR(r.U32.Bits(uint32(mask))) }
func (r *RCR) StoreBits(mask, b CR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCR) SetBits(mask CR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCR) ClearBits(mask CR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCR) Load() CR             { return CR(r.U32.Load()) }
func (r *RCR) Store(b CR)           { r.U32.Store(uint32(b)) }

func (r *RCR) AtomicStoreBits(mask, b CR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCR) AtomicSetBits(mask CR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCR) AtomicClearBits(mask CR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCR struct{ mmio.UM32 }

func (rm RMCR) Load() CR   { return CR(rm.UM32.Load()) }
func (rm RMCR) Store(b CR) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) PG() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PG)}}
}

func (p *FLASH_Periph) PER() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PER)}}
}

func (p *FLASH_Periph) MER() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(MER)}}
}

func (p *FLASH_Periph) OPTPG() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(OPTPG)}}
}

func (p *FLASH_Periph) OPTER() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(OPTER)}}
}

func (p *FLASH_Periph) STRT() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(STRT)}}
}

func (p *FLASH_Periph) LOCK() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(LOCK)}}
}

func (p *FLASH_Periph) OPTWRE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(OPTWRE)}}
}

func (p *FLASH_Periph) ERRIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ERRIE)}}
}

func (p *FLASH_Periph) EOPIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(EOPIE)}}
}

func (p *FLASH_Periph) OBL_LAUNCH() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(OBL_LAUNCH)}}
}

type AR uint32

func (b AR) Field(mask AR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AR) J(v int) AR {
	return AR(bits.MakeField32(v, uint32(mask)))
}

type RAR struct{ mmio.U32 }

func (r *RAR) Bits(mask AR) AR      { return AR(r.U32.Bits(uint32(mask))) }
func (r *RAR) StoreBits(mask, b AR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAR) SetBits(mask AR)      { r.U32.SetBits(uint32(mask)) }
func (r *RAR) ClearBits(mask AR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RAR) Load() AR             { return AR(r.U32.Load()) }
func (r *RAR) Store(b AR)           { r.U32.Store(uint32(b)) }

func (r *RAR) AtomicStoreBits(mask, b AR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RAR) AtomicSetBits(mask AR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAR) AtomicClearBits(mask AR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMAR struct{ mmio.UM32 }

func (rm RMAR) Load() AR   { return AR(rm.UM32.Load()) }
func (rm RMAR) Store(b AR) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) FAR() RMAR {
	return RMAR{mmio.UM32{&p.AR.U32, uint32(FAR)}}
}

type OBR uint32

func (b OBR) Field(mask OBR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OBR) J(v int) OBR {
	return OBR(bits.MakeField32(v, uint32(mask)))
}

type ROBR struct{ mmio.U32 }

func (r *ROBR) Bits(mask OBR) OBR     { return OBR(r.U32.Bits(uint32(mask))) }
func (r *ROBR) StoreBits(mask, b OBR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ROBR) SetBits(mask OBR)      { r.U32.SetBits(uint32(mask)) }
func (r *ROBR) ClearBits(mask OBR)    { r.U32.ClearBits(uint32(mask)) }
func (r *ROBR) Load() OBR             { return OBR(r.U32.Load()) }
func (r *ROBR) Store(b OBR)           { r.U32.Store(uint32(b)) }

func (r *ROBR) AtomicStoreBits(mask, b OBR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *ROBR) AtomicSetBits(mask OBR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ROBR) AtomicClearBits(mask OBR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMOBR struct{ mmio.UM32 }

func (rm RMOBR) Load() OBR   { return OBR(rm.UM32.Load()) }
func (rm RMOBR) Store(b OBR) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) OPTERR() RMOBR {
	return RMOBR{mmio.UM32{&p.OBR.U32, uint32(OPTERR)}}
}

func (p *FLASH_Periph) RDPRT() RMOBR {
	return RMOBR{mmio.UM32{&p.OBR.U32, uint32(RDPRT)}}
}

func (p *FLASH_Periph) USER() RMOBR {
	return RMOBR{mmio.UM32{&p.OBR.U32, uint32(USER)}}
}

func (p *FLASH_Periph) DATA0() RMOBR {
	return RMOBR{mmio.UM32{&p.OBR.U32, uint32(DATA0)}}
}

func (p *FLASH_Periph) DATA1() RMOBR {
	return RMOBR{mmio.UM32{&p.OBR.U32, uint32(DATA1)}}
}

type WRPR uint32

func (b WRPR) Field(mask WRPR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRPR) J(v int) WRPR {
	return WRPR(bits.MakeField32(v, uint32(mask)))
}

type RWRPR struct{ mmio.U32 }

func (r *RWRPR) Bits(mask WRPR) WRPR    { return WRPR(r.U32.Bits(uint32(mask))) }
func (r *RWRPR) StoreBits(mask, b WRPR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RWRPR) SetBits(mask WRPR)      { r.U32.SetBits(uint32(mask)) }
func (r *RWRPR) ClearBits(mask WRPR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RWRPR) Load() WRPR             { return WRPR(r.U32.Load()) }
func (r *RWRPR) Store(b WRPR)           { r.U32.Store(uint32(b)) }

func (r *RWRPR) AtomicStoreBits(mask, b WRPR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RWRPR) AtomicSetBits(mask WRPR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RWRPR) AtomicClearBits(mask WRPR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMWRPR struct{ mmio.UM32 }

func (rm RMWRPR) Load() WRPR   { return WRPR(rm.UM32.Load()) }
func (rm RMWRPR) Store(b WRPR) { rm.UM32.Store(uint32(b)) }
