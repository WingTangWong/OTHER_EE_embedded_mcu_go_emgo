package flash

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type FLASH_Periph struct {
	ACR       RACR
	PDKEYR    RPDKEYR
	KEYR      RKEYR
	OPTKEYR   ROPTKEYR
	SR        RSR
	CR        RCR
	ECCR      RECCR
	RESERVED1 RRESERVED1
	OPTR      ROPTR
	PCROP1SR  RPCROP1SR
	PCROP1ER  RPCROP1ER
	WRP1AR    RWRP1AR
	WRP1BR    RWRP1BR
	_         [4]uint32
	PCROP2SR  RPCROP2SR
	PCROP2ER  RPCROP2ER
	WRP2AR    RWRP2AR
	WRP2BR    RWRP2BR
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

func (p *FLASH_Periph) PRFTEN() RMACR {
	return RMACR{mmio.UM32{&p.ACR.U32, uint32(PRFTEN)}}
}

func (p *FLASH_Periph) ICEN() RMACR {
	return RMACR{mmio.UM32{&p.ACR.U32, uint32(ICEN)}}
}

func (p *FLASH_Periph) DCEN() RMACR {
	return RMACR{mmio.UM32{&p.ACR.U32, uint32(DCEN)}}
}

func (p *FLASH_Periph) ICRST() RMACR {
	return RMACR{mmio.UM32{&p.ACR.U32, uint32(ICRST)}}
}

func (p *FLASH_Periph) DCRST() RMACR {
	return RMACR{mmio.UM32{&p.ACR.U32, uint32(DCRST)}}
}

func (p *FLASH_Periph) RUN_PD() RMACR {
	return RMACR{mmio.UM32{&p.ACR.U32, uint32(RUN_PD)}}
}

func (p *FLASH_Periph) SLEEP_PD() RMACR {
	return RMACR{mmio.UM32{&p.ACR.U32, uint32(SLEEP_PD)}}
}

type PDKEYR uint32

func (b PDKEYR) Field(mask PDKEYR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PDKEYR) J(v int) PDKEYR {
	return PDKEYR(bits.MakeField32(v, uint32(mask)))
}

type RPDKEYR struct{ mmio.U32 }

func (r *RPDKEYR) Bits(mask PDKEYR) PDKEYR  { return PDKEYR(r.U32.Bits(uint32(mask))) }
func (r *RPDKEYR) StoreBits(mask, b PDKEYR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPDKEYR) SetBits(mask PDKEYR)      { r.U32.SetBits(uint32(mask)) }
func (r *RPDKEYR) ClearBits(mask PDKEYR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPDKEYR) Load() PDKEYR             { return PDKEYR(r.U32.Load()) }
func (r *RPDKEYR) Store(b PDKEYR)           { r.U32.Store(uint32(b)) }

func (r *RPDKEYR) AtomicStoreBits(mask, b PDKEYR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPDKEYR) AtomicSetBits(mask PDKEYR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPDKEYR) AtomicClearBits(mask PDKEYR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPDKEYR struct{ mmio.UM32 }

func (rm RMPDKEYR) Load() PDKEYR   { return PDKEYR(rm.UM32.Load()) }
func (rm RMPDKEYR) Store(b PDKEYR) { rm.UM32.Store(uint32(b)) }

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

func (p *FLASH_Periph) EOP() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(EOP)}}
}

func (p *FLASH_Periph) OPERR() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(OPERR)}}
}

func (p *FLASH_Periph) PROGERR() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(PROGERR)}}
}

func (p *FLASH_Periph) WRPERR() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(WRPERR)}}
}

func (p *FLASH_Periph) PGAERR() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(PGAERR)}}
}

func (p *FLASH_Periph) SIZERR() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(SIZERR)}}
}

func (p *FLASH_Periph) PGSERR() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(PGSERR)}}
}

func (p *FLASH_Periph) MISERR() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(MISERR)}}
}

func (p *FLASH_Periph) FASTERR() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(FASTERR)}}
}

func (p *FLASH_Periph) RDERR() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(RDERR)}}
}

func (p *FLASH_Periph) OPTVERR() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(OPTVERR)}}
}

func (p *FLASH_Periph) BSY() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(BSY)}}
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

func (p *FLASH_Periph) MER1() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(MER1)}}
}

func (p *FLASH_Periph) PNB() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PNB)}}
}

func (p *FLASH_Periph) BKER() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(BKER)}}
}

func (p *FLASH_Periph) MER2() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(MER2)}}
}

func (p *FLASH_Periph) STRT() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(STRT)}}
}

func (p *FLASH_Periph) OPTSTRT() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(OPTSTRT)}}
}

func (p *FLASH_Periph) FSTPG() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(FSTPG)}}
}

func (p *FLASH_Periph) EOPIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(EOPIE)}}
}

func (p *FLASH_Periph) ERRIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ERRIE)}}
}

func (p *FLASH_Periph) RDERRIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(RDERRIE)}}
}

func (p *FLASH_Periph) OBL_LAUNCH() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(OBL_LAUNCH)}}
}

func (p *FLASH_Periph) OPTLOCK() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(OPTLOCK)}}
}

func (p *FLASH_Periph) LOCK() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(LOCK)}}
}

type ECCR uint32

func (b ECCR) Field(mask ECCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ECCR) J(v int) ECCR {
	return ECCR(bits.MakeField32(v, uint32(mask)))
}

type RECCR struct{ mmio.U32 }

func (r *RECCR) Bits(mask ECCR) ECCR    { return ECCR(r.U32.Bits(uint32(mask))) }
func (r *RECCR) StoreBits(mask, b ECCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RECCR) SetBits(mask ECCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RECCR) ClearBits(mask ECCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RECCR) Load() ECCR             { return ECCR(r.U32.Load()) }
func (r *RECCR) Store(b ECCR)           { r.U32.Store(uint32(b)) }

func (r *RECCR) AtomicStoreBits(mask, b ECCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RECCR) AtomicSetBits(mask ECCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RECCR) AtomicClearBits(mask ECCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMECCR struct{ mmio.UM32 }

func (rm RMECCR) Load() ECCR   { return ECCR(rm.UM32.Load()) }
func (rm RMECCR) Store(b ECCR) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) ADDR_ECC() RMECCR {
	return RMECCR{mmio.UM32{&p.ECCR.U32, uint32(ADDR_ECC)}}
}

func (p *FLASH_Periph) BK_ECC() RMECCR {
	return RMECCR{mmio.UM32{&p.ECCR.U32, uint32(BK_ECC)}}
}

func (p *FLASH_Periph) SYSF_ECC() RMECCR {
	return RMECCR{mmio.UM32{&p.ECCR.U32, uint32(SYSF_ECC)}}
}

func (p *FLASH_Periph) ECCIE() RMECCR {
	return RMECCR{mmio.UM32{&p.ECCR.U32, uint32(ECCIE)}}
}

func (p *FLASH_Periph) ECCC() RMECCR {
	return RMECCR{mmio.UM32{&p.ECCR.U32, uint32(ECCC)}}
}

func (p *FLASH_Periph) ECCD() RMECCR {
	return RMECCR{mmio.UM32{&p.ECCR.U32, uint32(ECCD)}}
}

type RESERVED1 uint32

func (b RESERVED1) Field(mask RESERVED1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RESERVED1) J(v int) RESERVED1 {
	return RESERVED1(bits.MakeField32(v, uint32(mask)))
}

type RRESERVED1 struct{ mmio.U32 }

func (r *RRESERVED1) Bits(mask RESERVED1) RESERVED1 { return RESERVED1(r.U32.Bits(uint32(mask))) }
func (r *RRESERVED1) StoreBits(mask, b RESERVED1)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRESERVED1) SetBits(mask RESERVED1)        { r.U32.SetBits(uint32(mask)) }
func (r *RRESERVED1) ClearBits(mask RESERVED1)      { r.U32.ClearBits(uint32(mask)) }
func (r *RRESERVED1) Load() RESERVED1               { return RESERVED1(r.U32.Load()) }
func (r *RRESERVED1) Store(b RESERVED1)             { r.U32.Store(uint32(b)) }

func (r *RRESERVED1) AtomicStoreBits(mask, b RESERVED1) {
	r.U32.AtomicStoreBits(uint32(mask), uint32(b))
}
func (r *RRESERVED1) AtomicSetBits(mask RESERVED1)   { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRESERVED1) AtomicClearBits(mask RESERVED1) { r.U32.AtomicClearBits(uint32(mask)) }

type RMRESERVED1 struct{ mmio.UM32 }

func (rm RMRESERVED1) Load() RESERVED1   { return RESERVED1(rm.UM32.Load()) }
func (rm RMRESERVED1) Store(b RESERVED1) { rm.UM32.Store(uint32(b)) }

type OPTR uint32

func (b OPTR) Field(mask OPTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OPTR) J(v int) OPTR {
	return OPTR(bits.MakeField32(v, uint32(mask)))
}

type ROPTR struct{ mmio.U32 }

func (r *ROPTR) Bits(mask OPTR) OPTR    { return OPTR(r.U32.Bits(uint32(mask))) }
func (r *ROPTR) StoreBits(mask, b OPTR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ROPTR) SetBits(mask OPTR)      { r.U32.SetBits(uint32(mask)) }
func (r *ROPTR) ClearBits(mask OPTR)    { r.U32.ClearBits(uint32(mask)) }
func (r *ROPTR) Load() OPTR             { return OPTR(r.U32.Load()) }
func (r *ROPTR) Store(b OPTR)           { r.U32.Store(uint32(b)) }

func (r *ROPTR) AtomicStoreBits(mask, b OPTR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *ROPTR) AtomicSetBits(mask OPTR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ROPTR) AtomicClearBits(mask OPTR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMOPTR struct{ mmio.UM32 }

func (rm RMOPTR) Load() OPTR   { return OPTR(rm.UM32.Load()) }
func (rm RMOPTR) Store(b OPTR) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) RDP() RMOPTR {
	return RMOPTR{mmio.UM32{&p.OPTR.U32, uint32(RDP)}}
}

func (p *FLASH_Periph) BOR_LEV() RMOPTR {
	return RMOPTR{mmio.UM32{&p.OPTR.U32, uint32(BOR_LEV)}}
}

func (p *FLASH_Periph) nRST_STOP() RMOPTR {
	return RMOPTR{mmio.UM32{&p.OPTR.U32, uint32(nRST_STOP)}}
}

func (p *FLASH_Periph) nRST_STDBY() RMOPTR {
	return RMOPTR{mmio.UM32{&p.OPTR.U32, uint32(nRST_STDBY)}}
}

func (p *FLASH_Periph) nRST_SHDW() RMOPTR {
	return RMOPTR{mmio.UM32{&p.OPTR.U32, uint32(nRST_SHDW)}}
}

func (p *FLASH_Periph) IWDG_SW() RMOPTR {
	return RMOPTR{mmio.UM32{&p.OPTR.U32, uint32(IWDG_SW)}}
}

func (p *FLASH_Periph) IWDG_STOP() RMOPTR {
	return RMOPTR{mmio.UM32{&p.OPTR.U32, uint32(IWDG_STOP)}}
}

func (p *FLASH_Periph) IWDG_STDBY() RMOPTR {
	return RMOPTR{mmio.UM32{&p.OPTR.U32, uint32(IWDG_STDBY)}}
}

func (p *FLASH_Periph) WWDG_SW() RMOPTR {
	return RMOPTR{mmio.UM32{&p.OPTR.U32, uint32(WWDG_SW)}}
}

func (p *FLASH_Periph) BFB2() RMOPTR {
	return RMOPTR{mmio.UM32{&p.OPTR.U32, uint32(BFB2)}}
}

func (p *FLASH_Periph) DUALBANK() RMOPTR {
	return RMOPTR{mmio.UM32{&p.OPTR.U32, uint32(DUALBANK)}}
}

func (p *FLASH_Periph) nBOOT1() RMOPTR {
	return RMOPTR{mmio.UM32{&p.OPTR.U32, uint32(nBOOT1)}}
}

func (p *FLASH_Periph) SRAM2_PE() RMOPTR {
	return RMOPTR{mmio.UM32{&p.OPTR.U32, uint32(SRAM2_PE)}}
}

func (p *FLASH_Periph) SRAM2_RST() RMOPTR {
	return RMOPTR{mmio.UM32{&p.OPTR.U32, uint32(SRAM2_RST)}}
}

type PCROP1SR uint32

func (b PCROP1SR) Field(mask PCROP1SR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PCROP1SR) J(v int) PCROP1SR {
	return PCROP1SR(bits.MakeField32(v, uint32(mask)))
}

type RPCROP1SR struct{ mmio.U32 }

func (r *RPCROP1SR) Bits(mask PCROP1SR) PCROP1SR { return PCROP1SR(r.U32.Bits(uint32(mask))) }
func (r *RPCROP1SR) StoreBits(mask, b PCROP1SR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPCROP1SR) SetBits(mask PCROP1SR)       { r.U32.SetBits(uint32(mask)) }
func (r *RPCROP1SR) ClearBits(mask PCROP1SR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RPCROP1SR) Load() PCROP1SR              { return PCROP1SR(r.U32.Load()) }
func (r *RPCROP1SR) Store(b PCROP1SR)            { r.U32.Store(uint32(b)) }

func (r *RPCROP1SR) AtomicStoreBits(mask, b PCROP1SR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPCROP1SR) AtomicSetBits(mask PCROP1SR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPCROP1SR) AtomicClearBits(mask PCROP1SR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPCROP1SR struct{ mmio.UM32 }

func (rm RMPCROP1SR) Load() PCROP1SR   { return PCROP1SR(rm.UM32.Load()) }
func (rm RMPCROP1SR) Store(b PCROP1SR) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) PCROP1_STRT() RMPCROP1SR {
	return RMPCROP1SR{mmio.UM32{&p.PCROP1SR.U32, uint32(PCROP1_STRT)}}
}

type PCROP1ER uint32

func (b PCROP1ER) Field(mask PCROP1ER) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PCROP1ER) J(v int) PCROP1ER {
	return PCROP1ER(bits.MakeField32(v, uint32(mask)))
}

type RPCROP1ER struct{ mmio.U32 }

func (r *RPCROP1ER) Bits(mask PCROP1ER) PCROP1ER { return PCROP1ER(r.U32.Bits(uint32(mask))) }
func (r *RPCROP1ER) StoreBits(mask, b PCROP1ER)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPCROP1ER) SetBits(mask PCROP1ER)       { r.U32.SetBits(uint32(mask)) }
func (r *RPCROP1ER) ClearBits(mask PCROP1ER)     { r.U32.ClearBits(uint32(mask)) }
func (r *RPCROP1ER) Load() PCROP1ER              { return PCROP1ER(r.U32.Load()) }
func (r *RPCROP1ER) Store(b PCROP1ER)            { r.U32.Store(uint32(b)) }

func (r *RPCROP1ER) AtomicStoreBits(mask, b PCROP1ER) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPCROP1ER) AtomicSetBits(mask PCROP1ER)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPCROP1ER) AtomicClearBits(mask PCROP1ER)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPCROP1ER struct{ mmio.UM32 }

func (rm RMPCROP1ER) Load() PCROP1ER   { return PCROP1ER(rm.UM32.Load()) }
func (rm RMPCROP1ER) Store(b PCROP1ER) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) PCROP1_END() RMPCROP1ER {
	return RMPCROP1ER{mmio.UM32{&p.PCROP1ER.U32, uint32(PCROP1_END)}}
}

func (p *FLASH_Periph) PCROP_RDP() RMPCROP1ER {
	return RMPCROP1ER{mmio.UM32{&p.PCROP1ER.U32, uint32(PCROP_RDP)}}
}

type WRP1AR uint32

func (b WRP1AR) Field(mask WRP1AR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRP1AR) J(v int) WRP1AR {
	return WRP1AR(bits.MakeField32(v, uint32(mask)))
}

type RWRP1AR struct{ mmio.U32 }

func (r *RWRP1AR) Bits(mask WRP1AR) WRP1AR  { return WRP1AR(r.U32.Bits(uint32(mask))) }
func (r *RWRP1AR) StoreBits(mask, b WRP1AR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RWRP1AR) SetBits(mask WRP1AR)      { r.U32.SetBits(uint32(mask)) }
func (r *RWRP1AR) ClearBits(mask WRP1AR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RWRP1AR) Load() WRP1AR             { return WRP1AR(r.U32.Load()) }
func (r *RWRP1AR) Store(b WRP1AR)           { r.U32.Store(uint32(b)) }

func (r *RWRP1AR) AtomicStoreBits(mask, b WRP1AR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RWRP1AR) AtomicSetBits(mask WRP1AR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RWRP1AR) AtomicClearBits(mask WRP1AR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMWRP1AR struct{ mmio.UM32 }

func (rm RMWRP1AR) Load() WRP1AR   { return WRP1AR(rm.UM32.Load()) }
func (rm RMWRP1AR) Store(b WRP1AR) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) WRP1A_STRT() RMWRP1AR {
	return RMWRP1AR{mmio.UM32{&p.WRP1AR.U32, uint32(WRP1A_STRT)}}
}

func (p *FLASH_Periph) WRP1A_END() RMWRP1AR {
	return RMWRP1AR{mmio.UM32{&p.WRP1AR.U32, uint32(WRP1A_END)}}
}

type WRP1BR uint32

func (b WRP1BR) Field(mask WRP1BR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRP1BR) J(v int) WRP1BR {
	return WRP1BR(bits.MakeField32(v, uint32(mask)))
}

type RWRP1BR struct{ mmio.U32 }

func (r *RWRP1BR) Bits(mask WRP1BR) WRP1BR  { return WRP1BR(r.U32.Bits(uint32(mask))) }
func (r *RWRP1BR) StoreBits(mask, b WRP1BR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RWRP1BR) SetBits(mask WRP1BR)      { r.U32.SetBits(uint32(mask)) }
func (r *RWRP1BR) ClearBits(mask WRP1BR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RWRP1BR) Load() WRP1BR             { return WRP1BR(r.U32.Load()) }
func (r *RWRP1BR) Store(b WRP1BR)           { r.U32.Store(uint32(b)) }

func (r *RWRP1BR) AtomicStoreBits(mask, b WRP1BR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RWRP1BR) AtomicSetBits(mask WRP1BR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RWRP1BR) AtomicClearBits(mask WRP1BR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMWRP1BR struct{ mmio.UM32 }

func (rm RMWRP1BR) Load() WRP1BR   { return WRP1BR(rm.UM32.Load()) }
func (rm RMWRP1BR) Store(b WRP1BR) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) WRP1B_STRT() RMWRP1BR {
	return RMWRP1BR{mmio.UM32{&p.WRP1BR.U32, uint32(WRP1B_STRT)}}
}

func (p *FLASH_Periph) WRP1B_END() RMWRP1BR {
	return RMWRP1BR{mmio.UM32{&p.WRP1BR.U32, uint32(WRP1B_END)}}
}

type PCROP2SR uint32

func (b PCROP2SR) Field(mask PCROP2SR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PCROP2SR) J(v int) PCROP2SR {
	return PCROP2SR(bits.MakeField32(v, uint32(mask)))
}

type RPCROP2SR struct{ mmio.U32 }

func (r *RPCROP2SR) Bits(mask PCROP2SR) PCROP2SR { return PCROP2SR(r.U32.Bits(uint32(mask))) }
func (r *RPCROP2SR) StoreBits(mask, b PCROP2SR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPCROP2SR) SetBits(mask PCROP2SR)       { r.U32.SetBits(uint32(mask)) }
func (r *RPCROP2SR) ClearBits(mask PCROP2SR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RPCROP2SR) Load() PCROP2SR              { return PCROP2SR(r.U32.Load()) }
func (r *RPCROP2SR) Store(b PCROP2SR)            { r.U32.Store(uint32(b)) }

func (r *RPCROP2SR) AtomicStoreBits(mask, b PCROP2SR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPCROP2SR) AtomicSetBits(mask PCROP2SR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPCROP2SR) AtomicClearBits(mask PCROP2SR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPCROP2SR struct{ mmio.UM32 }

func (rm RMPCROP2SR) Load() PCROP2SR   { return PCROP2SR(rm.UM32.Load()) }
func (rm RMPCROP2SR) Store(b PCROP2SR) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) PCROP2_STRT() RMPCROP2SR {
	return RMPCROP2SR{mmio.UM32{&p.PCROP2SR.U32, uint32(PCROP2_STRT)}}
}

type PCROP2ER uint32

func (b PCROP2ER) Field(mask PCROP2ER) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PCROP2ER) J(v int) PCROP2ER {
	return PCROP2ER(bits.MakeField32(v, uint32(mask)))
}

type RPCROP2ER struct{ mmio.U32 }

func (r *RPCROP2ER) Bits(mask PCROP2ER) PCROP2ER { return PCROP2ER(r.U32.Bits(uint32(mask))) }
func (r *RPCROP2ER) StoreBits(mask, b PCROP2ER)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPCROP2ER) SetBits(mask PCROP2ER)       { r.U32.SetBits(uint32(mask)) }
func (r *RPCROP2ER) ClearBits(mask PCROP2ER)     { r.U32.ClearBits(uint32(mask)) }
func (r *RPCROP2ER) Load() PCROP2ER              { return PCROP2ER(r.U32.Load()) }
func (r *RPCROP2ER) Store(b PCROP2ER)            { r.U32.Store(uint32(b)) }

func (r *RPCROP2ER) AtomicStoreBits(mask, b PCROP2ER) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPCROP2ER) AtomicSetBits(mask PCROP2ER)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPCROP2ER) AtomicClearBits(mask PCROP2ER)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPCROP2ER struct{ mmio.UM32 }

func (rm RMPCROP2ER) Load() PCROP2ER   { return PCROP2ER(rm.UM32.Load()) }
func (rm RMPCROP2ER) Store(b PCROP2ER) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) PCROP2_END() RMPCROP2ER {
	return RMPCROP2ER{mmio.UM32{&p.PCROP2ER.U32, uint32(PCROP2_END)}}
}

type WRP2AR uint32

func (b WRP2AR) Field(mask WRP2AR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRP2AR) J(v int) WRP2AR {
	return WRP2AR(bits.MakeField32(v, uint32(mask)))
}

type RWRP2AR struct{ mmio.U32 }

func (r *RWRP2AR) Bits(mask WRP2AR) WRP2AR  { return WRP2AR(r.U32.Bits(uint32(mask))) }
func (r *RWRP2AR) StoreBits(mask, b WRP2AR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RWRP2AR) SetBits(mask WRP2AR)      { r.U32.SetBits(uint32(mask)) }
func (r *RWRP2AR) ClearBits(mask WRP2AR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RWRP2AR) Load() WRP2AR             { return WRP2AR(r.U32.Load()) }
func (r *RWRP2AR) Store(b WRP2AR)           { r.U32.Store(uint32(b)) }

func (r *RWRP2AR) AtomicStoreBits(mask, b WRP2AR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RWRP2AR) AtomicSetBits(mask WRP2AR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RWRP2AR) AtomicClearBits(mask WRP2AR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMWRP2AR struct{ mmio.UM32 }

func (rm RMWRP2AR) Load() WRP2AR   { return WRP2AR(rm.UM32.Load()) }
func (rm RMWRP2AR) Store(b WRP2AR) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) WRP2A_STRT() RMWRP2AR {
	return RMWRP2AR{mmio.UM32{&p.WRP2AR.U32, uint32(WRP2A_STRT)}}
}

func (p *FLASH_Periph) WRP2A_END() RMWRP2AR {
	return RMWRP2AR{mmio.UM32{&p.WRP2AR.U32, uint32(WRP2A_END)}}
}

type WRP2BR uint32

func (b WRP2BR) Field(mask WRP2BR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask WRP2BR) J(v int) WRP2BR {
	return WRP2BR(bits.MakeField32(v, uint32(mask)))
}

type RWRP2BR struct{ mmio.U32 }

func (r *RWRP2BR) Bits(mask WRP2BR) WRP2BR  { return WRP2BR(r.U32.Bits(uint32(mask))) }
func (r *RWRP2BR) StoreBits(mask, b WRP2BR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RWRP2BR) SetBits(mask WRP2BR)      { r.U32.SetBits(uint32(mask)) }
func (r *RWRP2BR) ClearBits(mask WRP2BR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RWRP2BR) Load() WRP2BR             { return WRP2BR(r.U32.Load()) }
func (r *RWRP2BR) Store(b WRP2BR)           { r.U32.Store(uint32(b)) }

func (r *RWRP2BR) AtomicStoreBits(mask, b WRP2BR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RWRP2BR) AtomicSetBits(mask WRP2BR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RWRP2BR) AtomicClearBits(mask WRP2BR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMWRP2BR struct{ mmio.UM32 }

func (rm RMWRP2BR) Load() WRP2BR   { return WRP2BR(rm.UM32.Load()) }
func (rm RMWRP2BR) Store(b WRP2BR) { rm.UM32.Store(uint32(b)) }

func (p *FLASH_Periph) WRP2B_STRT() RMWRP2BR {
	return RMWRP2BR{mmio.UM32{&p.WRP2BR.U32, uint32(WRP2B_STRT)}}
}

func (p *FLASH_Periph) WRP2B_END() RMWRP2BR {
	return RMWRP2BR{mmio.UM32{&p.WRP2BR.U32, uint32(WRP2B_END)}}
}
