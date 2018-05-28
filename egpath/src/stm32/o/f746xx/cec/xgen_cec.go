package cec

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f746xx/mmap"
)

type CEC_Periph struct {
	CR   RCR
	CFGR RCFGR
	TXDR RTXDR
	RXDR RRXDR
	ISR  RISR
	IER  RIER
}

func (p *CEC_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var CEC = (*CEC_Periph)(unsafe.Pointer(uintptr(mmap.CEC_BASE)))

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

func (p *CEC_Periph) CECEN() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(CECEN)}}
}

func (p *CEC_Periph) TXSOM() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TXSOM)}}
}

func (p *CEC_Periph) TXEOM() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TXEOM)}}
}

type CFGR uint32

func (b CFGR) Field(mask CFGR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CFGR) J(v int) CFGR {
	return CFGR(bits.MakeField32(v, uint32(mask)))
}

type RCFGR struct{ mmio.U32 }

func (r *RCFGR) Bits(mask CFGR) CFGR    { return CFGR(r.U32.Bits(uint32(mask))) }
func (r *RCFGR) StoreBits(mask, b CFGR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCFGR) SetBits(mask CFGR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCFGR) ClearBits(mask CFGR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCFGR) Load() CFGR             { return CFGR(r.U32.Load()) }
func (r *RCFGR) Store(b CFGR)           { r.U32.Store(uint32(b)) }

func (r *RCFGR) AtomicStoreBits(mask, b CFGR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCFGR) AtomicSetBits(mask CFGR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCFGR) AtomicClearBits(mask CFGR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCFGR struct{ mmio.UM32 }

func (rm RMCFGR) Load() CFGR   { return CFGR(rm.UM32.Load()) }
func (rm RMCFGR) Store(b CFGR) { rm.UM32.Store(uint32(b)) }

func (p *CEC_Periph) SFT() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(SFT)}}
}

func (p *CEC_Periph) RXTOL() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(RXTOL)}}
}

func (p *CEC_Periph) BRESTP() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(BRESTP)}}
}

func (p *CEC_Periph) BREGEN() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(BREGEN)}}
}

func (p *CEC_Periph) LBPEGEN() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(LBPEGEN)}}
}

func (p *CEC_Periph) BRDNOGEN() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(BRDNOGEN)}}
}

func (p *CEC_Periph) SFTOPT() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(SFTOPT)}}
}

func (p *CEC_Periph) OAR() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(OAR)}}
}

func (p *CEC_Periph) LSTN() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(LSTN)}}
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

func (p *CEC_Periph) TXD() RMTXDR {
	return RMTXDR{mmio.UM32{&p.TXDR.U32, uint32(TXD)}}
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

func (p *CEC_Periph) RXBR() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(RXBR)}}
}

func (p *CEC_Periph) RXEND() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(RXEND)}}
}

func (p *CEC_Periph) RXOVR() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(RXOVR)}}
}

func (p *CEC_Periph) BRE() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(BRE)}}
}

func (p *CEC_Periph) SBPE() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(SBPE)}}
}

func (p *CEC_Periph) LBPE() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(LBPE)}}
}

func (p *CEC_Periph) RXACKE() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(RXACKE)}}
}

func (p *CEC_Periph) ARBLST() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(ARBLST)}}
}

func (p *CEC_Periph) TXBR() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TXBR)}}
}

func (p *CEC_Periph) TXEND() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TXEND)}}
}

func (p *CEC_Periph) TXUDR() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TXUDR)}}
}

func (p *CEC_Periph) TXERR() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TXERR)}}
}

func (p *CEC_Periph) TXACKE() RMISR {
	return RMISR{mmio.UM32{&p.ISR.U32, uint32(TXACKE)}}
}

type IER uint32

func (b IER) Field(mask IER) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IER) J(v int) IER {
	return IER(bits.MakeField32(v, uint32(mask)))
}

type RIER struct{ mmio.U32 }

func (r *RIER) Bits(mask IER) IER     { return IER(r.U32.Bits(uint32(mask))) }
func (r *RIER) StoreBits(mask, b IER) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RIER) SetBits(mask IER)      { r.U32.SetBits(uint32(mask)) }
func (r *RIER) ClearBits(mask IER)    { r.U32.ClearBits(uint32(mask)) }
func (r *RIER) Load() IER             { return IER(r.U32.Load()) }
func (r *RIER) Store(b IER)           { r.U32.Store(uint32(b)) }

func (r *RIER) AtomicStoreBits(mask, b IER) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RIER) AtomicSetBits(mask IER)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RIER) AtomicClearBits(mask IER)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMIER struct{ mmio.UM32 }

func (rm RMIER) Load() IER   { return IER(rm.UM32.Load()) }
func (rm RMIER) Store(b IER) { rm.UM32.Store(uint32(b)) }

func (p *CEC_Periph) RXBRIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(RXBRIE)}}
}

func (p *CEC_Periph) RXENDIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(RXENDIE)}}
}

func (p *CEC_Periph) RXOVRIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(RXOVRIE)}}
}

func (p *CEC_Periph) BREIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(BREIE)}}
}

func (p *CEC_Periph) SBPEIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(SBPEIE)}}
}

func (p *CEC_Periph) LBPEIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(LBPEIE)}}
}

func (p *CEC_Periph) RXACKEIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(RXACKEIE)}}
}

func (p *CEC_Periph) ARBLSTIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(ARBLSTIE)}}
}

func (p *CEC_Periph) TXBRIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(TXBRIE)}}
}

func (p *CEC_Periph) TXENDIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(TXENDIE)}}
}

func (p *CEC_Periph) TXUDRIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(TXUDRIE)}}
}

func (p *CEC_Periph) TXERRIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(TXERRIE)}}
}

func (p *CEC_Periph) TXACKEIE() RMIER {
	return RMIER{mmio.UM32{&p.IER.U32, uint32(TXACKEIE)}}
}
