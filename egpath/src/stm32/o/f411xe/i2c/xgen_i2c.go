package i2c

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f411xe/mmap"
)

type I2C_Periph struct {
	CR1   RCR1
	CR2   RCR2
	OAR1  ROAR1
	OAR2  ROAR2
	DR    RDR
	SR1   RSR1
	SR2   RSR2
	CCR   RCCR
	TRISE RTRISE
	FLTR  RFLTR
}

func (p *I2C_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var I2C1 = (*I2C_Periph)(unsafe.Pointer(uintptr(mmap.I2C1_BASE)))

//emgo:const
var I2C2 = (*I2C_Periph)(unsafe.Pointer(uintptr(mmap.I2C2_BASE)))

//emgo:const
var I2C3 = (*I2C_Periph)(unsafe.Pointer(uintptr(mmap.I2C3_BASE)))

type CR1 uint32

func (b CR1) Field(mask CR1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR1) J(v int) CR1 {
	return CR1(bits.MakeField32(v, uint32(mask)))
}

type RCR1 struct{ mmio.U32 }

func (r *RCR1) Bits(mask CR1) CR1     { return CR1(r.U32.Bits(uint32(mask))) }
func (r *RCR1) StoreBits(mask, b CR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCR1) SetBits(mask CR1)      { r.U32.SetBits(uint32(mask)) }
func (r *RCR1) ClearBits(mask CR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCR1) Load() CR1             { return CR1(r.U32.Load()) }
func (r *RCR1) Store(b CR1)           { r.U32.Store(uint32(b)) }

func (r *RCR1) AtomicStoreBits(mask, b CR1) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCR1) AtomicSetBits(mask CR1)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCR1) AtomicClearBits(mask CR1)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCR1 struct{ mmio.UM32 }

func (rm RMCR1) Load() CR1   { return CR1(rm.UM32.Load()) }
func (rm RMCR1) Store(b CR1) { rm.UM32.Store(uint32(b)) }

func (p *I2C_Periph) PE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(PE)}}
}

func (p *I2C_Periph) SMBUS() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(SMBUS)}}
}

func (p *I2C_Periph) SMBTYPE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(SMBTYPE)}}
}

func (p *I2C_Periph) ENARP() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(ENARP)}}
}

func (p *I2C_Periph) ENPEC() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(ENPEC)}}
}

func (p *I2C_Periph) ENGC() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(ENGC)}}
}

func (p *I2C_Periph) NOSTRETCH() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(NOSTRETCH)}}
}

func (p *I2C_Periph) START() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(START)}}
}

func (p *I2C_Periph) STOP() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(STOP)}}
}

func (p *I2C_Periph) ACK() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(ACK)}}
}

func (p *I2C_Periph) POS() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(POS)}}
}

func (p *I2C_Periph) PEC() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(PEC)}}
}

func (p *I2C_Periph) ALERT() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(ALERT)}}
}

func (p *I2C_Periph) SWRST() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(SWRST)}}
}

type CR2 uint32

func (b CR2) Field(mask CR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR2) J(v int) CR2 {
	return CR2(bits.MakeField32(v, uint32(mask)))
}

type RCR2 struct{ mmio.U32 }

func (r *RCR2) Bits(mask CR2) CR2     { return CR2(r.U32.Bits(uint32(mask))) }
func (r *RCR2) StoreBits(mask, b CR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCR2) SetBits(mask CR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RCR2) ClearBits(mask CR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCR2) Load() CR2             { return CR2(r.U32.Load()) }
func (r *RCR2) Store(b CR2)           { r.U32.Store(uint32(b)) }

func (r *RCR2) AtomicStoreBits(mask, b CR2) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCR2) AtomicSetBits(mask CR2)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCR2) AtomicClearBits(mask CR2)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCR2 struct{ mmio.UM32 }

func (rm RMCR2) Load() CR2   { return CR2(rm.UM32.Load()) }
func (rm RMCR2) Store(b CR2) { rm.UM32.Store(uint32(b)) }

func (p *I2C_Periph) FREQ() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(FREQ)}}
}

func (p *I2C_Periph) ITERREN() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(ITERREN)}}
}

func (p *I2C_Periph) ITEVTEN() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(ITEVTEN)}}
}

func (p *I2C_Periph) ITBUFEN() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(ITBUFEN)}}
}

func (p *I2C_Periph) DMAEN() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(DMAEN)}}
}

func (p *I2C_Periph) LAST() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(LAST)}}
}

type OAR1 uint32

func (b OAR1) Field(mask OAR1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OAR1) J(v int) OAR1 {
	return OAR1(bits.MakeField32(v, uint32(mask)))
}

type ROAR1 struct{ mmio.U32 }

func (r *ROAR1) Bits(mask OAR1) OAR1    { return OAR1(r.U32.Bits(uint32(mask))) }
func (r *ROAR1) StoreBits(mask, b OAR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ROAR1) SetBits(mask OAR1)      { r.U32.SetBits(uint32(mask)) }
func (r *ROAR1) ClearBits(mask OAR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *ROAR1) Load() OAR1             { return OAR1(r.U32.Load()) }
func (r *ROAR1) Store(b OAR1)           { r.U32.Store(uint32(b)) }

func (r *ROAR1) AtomicStoreBits(mask, b OAR1) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *ROAR1) AtomicSetBits(mask OAR1)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ROAR1) AtomicClearBits(mask OAR1)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMOAR1 struct{ mmio.UM32 }

func (rm RMOAR1) Load() OAR1   { return OAR1(rm.UM32.Load()) }
func (rm RMOAR1) Store(b OAR1) { rm.UM32.Store(uint32(b)) }

func (p *I2C_Periph) ADD1_7() RMOAR1 {
	return RMOAR1{mmio.UM32{&p.OAR1.U32, uint32(ADD1_7)}}
}

func (p *I2C_Periph) ADD8_9() RMOAR1 {
	return RMOAR1{mmio.UM32{&p.OAR1.U32, uint32(ADD8_9)}}
}

func (p *I2C_Periph) ADD0() RMOAR1 {
	return RMOAR1{mmio.UM32{&p.OAR1.U32, uint32(ADD0)}}
}

func (p *I2C_Periph) ADDMODE() RMOAR1 {
	return RMOAR1{mmio.UM32{&p.OAR1.U32, uint32(ADDMODE)}}
}

type OAR2 uint32

func (b OAR2) Field(mask OAR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OAR2) J(v int) OAR2 {
	return OAR2(bits.MakeField32(v, uint32(mask)))
}

type ROAR2 struct{ mmio.U32 }

func (r *ROAR2) Bits(mask OAR2) OAR2    { return OAR2(r.U32.Bits(uint32(mask))) }
func (r *ROAR2) StoreBits(mask, b OAR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ROAR2) SetBits(mask OAR2)      { r.U32.SetBits(uint32(mask)) }
func (r *ROAR2) ClearBits(mask OAR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *ROAR2) Load() OAR2             { return OAR2(r.U32.Load()) }
func (r *ROAR2) Store(b OAR2)           { r.U32.Store(uint32(b)) }

func (r *ROAR2) AtomicStoreBits(mask, b OAR2) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *ROAR2) AtomicSetBits(mask OAR2)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ROAR2) AtomicClearBits(mask OAR2)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMOAR2 struct{ mmio.UM32 }

func (rm RMOAR2) Load() OAR2   { return OAR2(rm.UM32.Load()) }
func (rm RMOAR2) Store(b OAR2) { rm.UM32.Store(uint32(b)) }

func (p *I2C_Periph) ENDUAL() RMOAR2 {
	return RMOAR2{mmio.UM32{&p.OAR2.U32, uint32(ENDUAL)}}
}

func (p *I2C_Periph) SECADD1_7() RMOAR2 {
	return RMOAR2{mmio.UM32{&p.OAR2.U32, uint32(SECADD1_7)}}
}

type DR uint32

func (b DR) Field(mask DR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DR) J(v int) DR {
	return DR(bits.MakeField32(v, uint32(mask)))
}

type RDR struct{ mmio.U32 }

func (r *RDR) Bits(mask DR) DR      { return DR(r.U32.Bits(uint32(mask))) }
func (r *RDR) StoreBits(mask, b DR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDR) SetBits(mask DR)      { r.U32.SetBits(uint32(mask)) }
func (r *RDR) ClearBits(mask DR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDR) Load() DR             { return DR(r.U32.Load()) }
func (r *RDR) Store(b DR)           { r.U32.Store(uint32(b)) }

func (r *RDR) AtomicStoreBits(mask, b DR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDR) AtomicSetBits(mask DR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDR) AtomicClearBits(mask DR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDR struct{ mmio.UM32 }

func (rm RMDR) Load() DR   { return DR(rm.UM32.Load()) }
func (rm RMDR) Store(b DR) { rm.UM32.Store(uint32(b)) }

type SR1 uint32

func (b SR1) Field(mask SR1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR1) J(v int) SR1 {
	return SR1(bits.MakeField32(v, uint32(mask)))
}

type RSR1 struct{ mmio.U32 }

func (r *RSR1) Bits(mask SR1) SR1     { return SR1(r.U32.Bits(uint32(mask))) }
func (r *RSR1) StoreBits(mask, b SR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSR1) SetBits(mask SR1)      { r.U32.SetBits(uint32(mask)) }
func (r *RSR1) ClearBits(mask SR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSR1) Load() SR1             { return SR1(r.U32.Load()) }
func (r *RSR1) Store(b SR1)           { r.U32.Store(uint32(b)) }

func (r *RSR1) AtomicStoreBits(mask, b SR1) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSR1) AtomicSetBits(mask SR1)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSR1) AtomicClearBits(mask SR1)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSR1 struct{ mmio.UM32 }

func (rm RMSR1) Load() SR1   { return SR1(rm.UM32.Load()) }
func (rm RMSR1) Store(b SR1) { rm.UM32.Store(uint32(b)) }

func (p *I2C_Periph) SB() RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(SB)}}
}

func (p *I2C_Periph) ADDR() RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(ADDR)}}
}

func (p *I2C_Periph) BTF() RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(BTF)}}
}

func (p *I2C_Periph) ADD10() RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(ADD10)}}
}

func (p *I2C_Periph) STOPF() RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(STOPF)}}
}

func (p *I2C_Periph) RXNE() RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(RXNE)}}
}

func (p *I2C_Periph) TXE() RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(TXE)}}
}

func (p *I2C_Periph) BERR() RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(BERR)}}
}

func (p *I2C_Periph) ARLO() RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(ARLO)}}
}

func (p *I2C_Periph) AF() RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(AF)}}
}

func (p *I2C_Periph) OVR() RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(OVR)}}
}

func (p *I2C_Periph) PECERR() RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(PECERR)}}
}

func (p *I2C_Periph) TIMEOUT() RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(TIMEOUT)}}
}

func (p *I2C_Periph) SMBALERT() RMSR1 {
	return RMSR1{mmio.UM32{&p.SR1.U32, uint32(SMBALERT)}}
}

type SR2 uint32

func (b SR2) Field(mask SR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR2) J(v int) SR2 {
	return SR2(bits.MakeField32(v, uint32(mask)))
}

type RSR2 struct{ mmio.U32 }

func (r *RSR2) Bits(mask SR2) SR2     { return SR2(r.U32.Bits(uint32(mask))) }
func (r *RSR2) StoreBits(mask, b SR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSR2) SetBits(mask SR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RSR2) ClearBits(mask SR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSR2) Load() SR2             { return SR2(r.U32.Load()) }
func (r *RSR2) Store(b SR2)           { r.U32.Store(uint32(b)) }

func (r *RSR2) AtomicStoreBits(mask, b SR2) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSR2) AtomicSetBits(mask SR2)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSR2) AtomicClearBits(mask SR2)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSR2 struct{ mmio.UM32 }

func (rm RMSR2) Load() SR2   { return SR2(rm.UM32.Load()) }
func (rm RMSR2) Store(b SR2) { rm.UM32.Store(uint32(b)) }

func (p *I2C_Periph) MSL() RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(MSL)}}
}

func (p *I2C_Periph) BUSY() RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(BUSY)}}
}

func (p *I2C_Periph) TRA() RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(TRA)}}
}

func (p *I2C_Periph) GENCALL() RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(GENCALL)}}
}

func (p *I2C_Periph) SMBDEFAULT() RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(SMBDEFAULT)}}
}

func (p *I2C_Periph) SMBHOST() RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(SMBHOST)}}
}

func (p *I2C_Periph) DUALF() RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(DUALF)}}
}

func (p *I2C_Periph) PECVAL() RMSR2 {
	return RMSR2{mmio.UM32{&p.SR2.U32, uint32(PECVAL)}}
}

type CCR uint32

func (b CCR) Field(mask CCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CCR) J(v int) CCR {
	return CCR(bits.MakeField32(v, uint32(mask)))
}

type RCCR struct{ mmio.U32 }

func (r *RCCR) Bits(mask CCR) CCR     { return CCR(r.U32.Bits(uint32(mask))) }
func (r *RCCR) StoreBits(mask, b CCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCCR) SetBits(mask CCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCCR) ClearBits(mask CCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCCR) Load() CCR             { return CCR(r.U32.Load()) }
func (r *RCCR) Store(b CCR)           { r.U32.Store(uint32(b)) }

func (r *RCCR) AtomicStoreBits(mask, b CCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCCR) AtomicSetBits(mask CCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCCR) AtomicClearBits(mask CCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCCR struct{ mmio.UM32 }

func (rm RMCCR) Load() CCR   { return CCR(rm.UM32.Load()) }
func (rm RMCCR) Store(b CCR) { rm.UM32.Store(uint32(b)) }

func (p *I2C_Periph) CCRVAL() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(CCRVAL)}}
}

func (p *I2C_Periph) DUTY() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(DUTY)}}
}

func (p *I2C_Periph) FS() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(FS)}}
}

type TRISE uint32

func (b TRISE) Field(mask TRISE) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TRISE) J(v int) TRISE {
	return TRISE(bits.MakeField32(v, uint32(mask)))
}

type RTRISE struct{ mmio.U32 }

func (r *RTRISE) Bits(mask TRISE) TRISE   { return TRISE(r.U32.Bits(uint32(mask))) }
func (r *RTRISE) StoreBits(mask, b TRISE) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTRISE) SetBits(mask TRISE)      { r.U32.SetBits(uint32(mask)) }
func (r *RTRISE) ClearBits(mask TRISE)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTRISE) Load() TRISE             { return TRISE(r.U32.Load()) }
func (r *RTRISE) Store(b TRISE)           { r.U32.Store(uint32(b)) }

func (r *RTRISE) AtomicStoreBits(mask, b TRISE) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RTRISE) AtomicSetBits(mask TRISE)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RTRISE) AtomicClearBits(mask TRISE)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMTRISE struct{ mmio.UM32 }

func (rm RMTRISE) Load() TRISE   { return TRISE(rm.UM32.Load()) }
func (rm RMTRISE) Store(b TRISE) { rm.UM32.Store(uint32(b)) }

type FLTR uint32

func (b FLTR) Field(mask FLTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FLTR) J(v int) FLTR {
	return FLTR(bits.MakeField32(v, uint32(mask)))
}

type RFLTR struct{ mmio.U32 }

func (r *RFLTR) Bits(mask FLTR) FLTR    { return FLTR(r.U32.Bits(uint32(mask))) }
func (r *RFLTR) StoreBits(mask, b FLTR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RFLTR) SetBits(mask FLTR)      { r.U32.SetBits(uint32(mask)) }
func (r *RFLTR) ClearBits(mask FLTR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RFLTR) Load() FLTR             { return FLTR(r.U32.Load()) }
func (r *RFLTR) Store(b FLTR)           { r.U32.Store(uint32(b)) }

func (r *RFLTR) AtomicStoreBits(mask, b FLTR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RFLTR) AtomicSetBits(mask FLTR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RFLTR) AtomicClearBits(mask FLTR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMFLTR struct{ mmio.UM32 }

func (rm RMFLTR) Load() FLTR   { return FLTR(rm.UM32.Load()) }
func (rm RMFLTR) Store(b FLTR) { rm.UM32.Store(uint32(b)) }

func (p *I2C_Periph) DNF() RMFLTR {
	return RMFLTR{mmio.UM32{&p.FLTR.U32, uint32(DNF)}}
}

func (p *I2C_Periph) ANOFF() RMFLTR {
	return RMFLTR{mmio.UM32{&p.FLTR.U32, uint32(ANOFF)}}
}
