package i2c

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type I2C_Periph struct {
	CR1      RCR1
	CR2      RCR2
	OAR1     ROAR1
	OAR2     ROAR2
	TIMINGR  RTIMINGR
	TIMEOUTR RTIMEOUTR
	ISR      RISR
	ICR      RICR
	PECR     RPECR
	RXDR     RRXDR
	TXDR     RTXDR
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

func (p *I2C_Periph) TXIE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(TXIE)}}
}

func (p *I2C_Periph) RXIE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(RXIE)}}
}

func (p *I2C_Periph) ADDRIE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(ADDRIE)}}
}

func (p *I2C_Periph) NACKIE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(NACKIE)}}
}

func (p *I2C_Periph) STOPIE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(STOPIE)}}
}

func (p *I2C_Periph) TCIE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(TCIE)}}
}

func (p *I2C_Periph) ERRIE() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(ERRIE)}}
}

func (p *I2C_Periph) DNF() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(DNF)}}
}

func (p *I2C_Periph) ANFOFF() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(ANFOFF)}}
}

func (p *I2C_Periph) SWRST() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(SWRST)}}
}

func (p *I2C_Periph) TXDMAEN() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(TXDMAEN)}}
}

func (p *I2C_Periph) RXDMAEN() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(RXDMAEN)}}
}

func (p *I2C_Periph) SBC() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(SBC)}}
}

func (p *I2C_Periph) NOSTRETCH() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(NOSTRETCH)}}
}

func (p *I2C_Periph) WUPEN() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(WUPEN)}}
}

func (p *I2C_Periph) GCEN() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(GCEN)}}
}

func (p *I2C_Periph) SMBHEN() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(SMBHEN)}}
}

func (p *I2C_Periph) SMBDEN() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(SMBDEN)}}
}

func (p *I2C_Periph) ALERTEN() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(ALERTEN)}}
}

func (p *I2C_Periph) PECEN() RMCR1 {
	return RMCR1{mmio.UM32{&p.CR1.U32, uint32(PECEN)}}
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

func (p *I2C_Periph) SADD() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(SADD)}}
}

func (p *I2C_Periph) RD_WRN() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(RD_WRN)}}
}

func (p *I2C_Periph) ADD10() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(ADD10)}}
}

func (p *I2C_Periph) HEAD10R() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(HEAD10R)}}
}

func (p *I2C_Periph) START() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(START)}}
}

func (p *I2C_Periph) STOP() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(STOP)}}
}

func (p *I2C_Periph) NACK() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(NACK)}}
}

func (p *I2C_Periph) NBYTES() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(NBYTES)}}
}

func (p *I2C_Periph) RELOAD() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(RELOAD)}}
}

func (p *I2C_Periph) AUTOEND() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(AUTOEND)}}
}

func (p *I2C_Periph) PECBYTE() RMCR2 {
	return RMCR2{mmio.UM32{&p.CR2.U32, uint32(PECBYTE)}}
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

func (p *I2C_Periph) OA1() RMOAR1 {
	return RMOAR1{mmio.UM32{&p.OAR1.U32, uint32(OA1)}}
}

func (p *I2C_Periph) OA1MODE() RMOAR1 {
	return RMOAR1{mmio.UM32{&p.OAR1.U32, uint32(OA1MODE)}}
}

func (p *I2C_Periph) OA1EN() RMOAR1 {
	return RMOAR1{mmio.UM32{&p.OAR1.U32, uint32(OA1EN)}}
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

func (p *I2C_Periph) OA2() RMOAR2 {
	return RMOAR2{mmio.UM32{&p.OAR2.U32, uint32(OA2)}}
}

func (p *I2C_Periph) OA2MSK() RMOAR2 {
	return RMOAR2{mmio.UM32{&p.OAR2.U32, uint32(OA2MSK)}}
}

func (p *I2C_Periph) OA2EN() RMOAR2 {
	return RMOAR2{mmio.UM32{&p.OAR2.U32, uint32(OA2EN)}}
}

type TIMINGR uint32

func (b TIMINGR) Field(mask TIMINGR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TIMINGR) J(v int) TIMINGR {
	return TIMINGR(bits.MakeField32(v, uint32(mask)))
}

type RTIMINGR struct{ mmio.U32 }

func (r *RTIMINGR) Bits(mask TIMINGR) TIMINGR { return TIMINGR(r.U32.Bits(uint32(mask))) }
func (r *RTIMINGR) StoreBits(mask, b TIMINGR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTIMINGR) SetBits(mask TIMINGR)      { r.U32.SetBits(uint32(mask)) }
func (r *RTIMINGR) ClearBits(mask TIMINGR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTIMINGR) Load() TIMINGR             { return TIMINGR(r.U32.Load()) }
func (r *RTIMINGR) Store(b TIMINGR)           { r.U32.Store(uint32(b)) }

func (r *RTIMINGR) AtomicStoreBits(mask, b TIMINGR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RTIMINGR) AtomicSetBits(mask TIMINGR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RTIMINGR) AtomicClearBits(mask TIMINGR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMTIMINGR struct{ mmio.UM32 }

func (rm RMTIMINGR) Load() TIMINGR   { return TIMINGR(rm.UM32.Load()) }
func (rm RMTIMINGR) Store(b TIMINGR) { rm.UM32.Store(uint32(b)) }

func (p *I2C_Periph) SCLL() RMTIMINGR {
	return RMTIMINGR{mmio.UM32{&p.TIMINGR.U32, uint32(SCLL)}}
}

func (p *I2C_Periph) SCLH() RMTIMINGR {
	return RMTIMINGR{mmio.UM32{&p.TIMINGR.U32, uint32(SCLH)}}
}

func (p *I2C_Periph) SDADEL() RMTIMINGR {
	return RMTIMINGR{mmio.UM32{&p.TIMINGR.U32, uint32(SDADEL)}}
}

func (p *I2C_Periph) SCLDEL() RMTIMINGR {
	return RMTIMINGR{mmio.UM32{&p.TIMINGR.U32, uint32(SCLDEL)}}
}

func (p *I2C_Periph) PRESC() RMTIMINGR {
	return RMTIMINGR{mmio.UM32{&p.TIMINGR.U32, uint32(PRESC)}}
}

type TIMEOUTR uint32

func (b TIMEOUTR) Field(mask TIMEOUTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TIMEOUTR) J(v int) TIMEOUTR {
	return TIMEOUTR(bits.MakeField32(v, uint32(mask)))
}

type RTIMEOUTR struct{ mmio.U32 }

func (r *RTIMEOUTR) Bits(mask TIMEOUTR) TIMEOUTR { return TIMEOUTR(r.U32.Bits(uint32(mask))) }
func (r *RTIMEOUTR) StoreBits(mask, b TIMEOUTR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTIMEOUTR) SetBits(mask TIMEOUTR)       { r.U32.SetBits(uint32(mask)) }
func (r *RTIMEOUTR) ClearBits(mask TIMEOUTR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RTIMEOUTR) Load() TIMEOUTR              { return TIMEOUTR(r.U32.Load()) }
func (r *RTIMEOUTR) Store(b TIMEOUTR)            { r.U32.Store(uint32(b)) }

func (r *RTIMEOUTR) AtomicStoreBits(mask, b TIMEOUTR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RTIMEOUTR) AtomicSetBits(mask TIMEOUTR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RTIMEOUTR) AtomicClearBits(mask TIMEOUTR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMTIMEOUTR struct{ mmio.UM32 }

func (rm RMTIMEOUTR) Load() TIMEOUTR   { return TIMEOUTR(rm.UM32.Load()) }
func (rm RMTIMEOUTR) Store(b TIMEOUTR) { rm.UM32.Store(uint32(b)) }

func (p *I2C_Periph) TIMEOUTA() RMTIMEOUTR {
	return RMTIMEOUTR{mmio.UM32{&p.TIMEOUTR.U32, uint32(TIMEOUTA)}}
}

func (p *I2C_Periph) TIDLE() RMTIMEOUTR {
	return RMTIMEOUTR{mmio.UM32{&p.TIMEOUTR.U32, uint32(TIDLE)}}
}

func (p *I2C_Periph) TIMOUTEN() RMTIMEOUTR {
	return RMTIMEOUTR{mmio.UM32{&p.TIMEOUTR.U32, uint32(TIMOUTEN)}}
}

func (p *I2C_Periph) TIMEOUTB() RMTIMEOUTR {
	return RMTIMEOUTR{mmio.UM32{&p.TIMEOUTR.U32, uint32(TIMEOUTB)}}
}

func (p *I2C_Periph) TEXTEN() RMTIMEOUTR {
	return RMTIMEOUTR{mmio.UM32{&p.TIMEOUTR.U32, uint32(TEXTEN)}}
}

type ISR uint32

func (b ISR) Field(mask ISR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ISR) J(v int) ISR {
	return ISR(bits.MakeField32(v, uint32(mask)))
}

type RISR struct{ mmio.U32 }

func (r *RISR) Bits(mask ISR) ISR     { return ISR(r.U32.Bits(uint32(mask))) }
func (r *RISR) StoreBits(mask, b ISR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RISR) SetBits(mask ISR)      { r.U32.SetBits(uint32(mask)) }
func (r *RISR) ClearBits(mask ISR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RISR) Load() ISR             { return ISR(r.U32.Load()) }
func (r *RISR) Store(b ISR)           { r.U32.Store(uint32(b)) }

func (r *RISR) AtomicStoreBits(mask, b ISR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RISR) AtomicSetBits(mask ISR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RISR) AtomicClearBits(mask ISR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMISR struct{ mmio.UM32 }

func (rm RMISR) Load() ISR   { return ISR(rm.UM32.Load()) }
func (rm RMISR) Store(b ISR) { rm.UM32.Store(uint32(b)) }

func (p *I2C_Periph) TXE() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TXE)}}
}

func (p *I2C_Periph) TXIS() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TXIS)}}
}

func (p *I2C_Periph) RXNE() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(RXNE)}}
}

func (p *I2C_Periph) ADDR() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(ADDR)}}
}

func (p *I2C_Periph) NACKF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(NACKF)}}
}

func (p *I2C_Periph) STOPF() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(STOPF)}}
}

func (p *I2C_Periph) TC() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TC)}}
}

func (p *I2C_Periph) TCR() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TCR)}}
}

func (p *I2C_Periph) BERR() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(BERR)}}
}

func (p *I2C_Periph) ARLO() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(ARLO)}}
}

func (p *I2C_Periph) OVR() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(OVR)}}
}

func (p *I2C_Periph) PECERR() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(PECERR)}}
}

func (p *I2C_Periph) TIMEOUT() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TIMEOUT)}}
}

func (p *I2C_Periph) ALERT() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(ALERT)}}
}

func (p *I2C_Periph) BUSY() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(BUSY)}}
}

func (p *I2C_Periph) DIR() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(DIR)}}
}

func (p *I2C_Periph) ADDCODE() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(ADDCODE)}}
}

type ICR uint32

func (b ICR) Field(mask ICR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ICR) J(v int) ICR {
	return ICR(bits.MakeField32(v, uint32(mask)))
}

type RICR struct{ mmio.U32 }

func (r *RICR) Bits(mask ICR) ICR     { return ICR(r.U32.Bits(uint32(mask))) }
func (r *RICR) StoreBits(mask, b ICR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RICR) SetBits(mask ICR)      { r.U32.SetBits(uint32(mask)) }
func (r *RICR) ClearBits(mask ICR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RICR) Load() ICR             { return ICR(r.U32.Load()) }
func (r *RICR) Store(b ICR)           { r.U32.Store(uint32(b)) }

func (r *RICR) AtomicStoreBits(mask, b ICR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RICR) AtomicSetBits(mask ICR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RICR) AtomicClearBits(mask ICR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMICR struct{ mmio.UM32 }

func (rm RMICR) Load() ICR   { return ICR(rm.UM32.Load()) }
func (rm RMICR) Store(b ICR) { rm.UM32.Store(uint32(b)) }

func (p *I2C_Periph) ADDRCF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(ADDRCF)}}
}

func (p *I2C_Periph) NACKCF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(NACKCF)}}
}

func (p *I2C_Periph) STOPCF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(STOPCF)}}
}

func (p *I2C_Periph) BERRCF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(BERRCF)}}
}

func (p *I2C_Periph) ARLOCF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(ARLOCF)}}
}

func (p *I2C_Periph) OVRCF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(OVRCF)}}
}

func (p *I2C_Periph) PECCF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(PECCF)}}
}

func (p *I2C_Periph) TIMOUTCF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(TIMOUTCF)}}
}

func (p *I2C_Periph) ALERTCF() RMICR {
	return RMICR{mmio.UM32{&p.ICR.U32, uint32(ALERTCF)}}
}

type PECR uint32

func (b PECR) Field(mask PECR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PECR) J(v int) PECR {
	return PECR(bits.MakeField32(v, uint32(mask)))
}

type RPECR struct{ mmio.U32 }

func (r *RPECR) Bits(mask PECR) PECR    { return PECR(r.U32.Bits(uint32(mask))) }
func (r *RPECR) StoreBits(mask, b PECR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPECR) SetBits(mask PECR)      { r.U32.SetBits(uint32(mask)) }
func (r *RPECR) ClearBits(mask PECR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPECR) Load() PECR             { return PECR(r.U32.Load()) }
func (r *RPECR) Store(b PECR)           { r.U32.Store(uint32(b)) }

func (r *RPECR) AtomicStoreBits(mask, b PECR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPECR) AtomicSetBits(mask PECR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPECR) AtomicClearBits(mask PECR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPECR struct{ mmio.UM32 }

func (rm RMPECR) Load() PECR   { return PECR(rm.UM32.Load()) }
func (rm RMPECR) Store(b PECR) { rm.UM32.Store(uint32(b)) }

func (p *I2C_Periph) PEC() RMPECR {
	return RMPECR{mmio.UM32{&p.PECR.U32, uint32(PEC)}}
}

type RXDR uint32

func (b RXDR) Field(mask RXDR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RXDR) J(v int) RXDR {
	return RXDR(bits.MakeField32(v, uint32(mask)))
}

type RRXDR struct{ mmio.U32 }

func (r *RRXDR) Bits(mask RXDR) RXDR    { return RXDR(r.U32.Bits(uint32(mask))) }
func (r *RRXDR) StoreBits(mask, b RXDR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRXDR) SetBits(mask RXDR)      { r.U32.SetBits(uint32(mask)) }
func (r *RRXDR) ClearBits(mask RXDR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RRXDR) Load() RXDR             { return RXDR(r.U32.Load()) }
func (r *RRXDR) Store(b RXDR)           { r.U32.Store(uint32(b)) }

func (r *RRXDR) AtomicStoreBits(mask, b RXDR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RRXDR) AtomicSetBits(mask RXDR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRXDR) AtomicClearBits(mask RXDR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMRXDR struct{ mmio.UM32 }

func (rm RMRXDR) Load() RXDR   { return RXDR(rm.UM32.Load()) }
func (rm RMRXDR) Store(b RXDR) { rm.UM32.Store(uint32(b)) }

func (p *I2C_Periph) RXDATA() RMRXDR {
	return RMRXDR{mmio.UM32{&p.RXDR.U32, uint32(RXDATA)}}
}

type TXDR uint32

func (b TXDR) Field(mask TXDR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TXDR) J(v int) TXDR {
	return TXDR(bits.MakeField32(v, uint32(mask)))
}

type RTXDR struct{ mmio.U32 }

func (r *RTXDR) Bits(mask TXDR) TXDR    { return TXDR(r.U32.Bits(uint32(mask))) }
func (r *RTXDR) StoreBits(mask, b TXDR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTXDR) SetBits(mask TXDR)      { r.U32.SetBits(uint32(mask)) }
func (r *RTXDR) ClearBits(mask TXDR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTXDR) Load() TXDR             { return TXDR(r.U32.Load()) }
func (r *RTXDR) Store(b TXDR)           { r.U32.Store(uint32(b)) }

func (r *RTXDR) AtomicStoreBits(mask, b TXDR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RTXDR) AtomicSetBits(mask TXDR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RTXDR) AtomicClearBits(mask TXDR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMTXDR struct{ mmio.UM32 }

func (rm RMTXDR) Load() TXDR   { return TXDR(rm.UM32.Load()) }
func (rm RMTXDR) Store(b TXDR) { rm.UM32.Store(uint32(b)) }

func (p *I2C_Periph) TXDATA() RMTXDR {
	return RMTXDR{mmio.UM32{&p.TXDR.U32, uint32(TXDATA)}}
}
