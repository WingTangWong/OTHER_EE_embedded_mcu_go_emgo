// +build l1xx_md

package pwr

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l1xx_md/mmap"
)

type PWR_Periph struct {
	CR  RCR
	CSR RCSR
}

func (p *PWR_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var PWR = (*PWR_Periph)(unsafe.Pointer(uintptr(mmap.PWR_BASE)))

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

func (p *PWR_Periph) LPSDSR() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(LPSDSR)}}
}

func (p *PWR_Periph) PDDS() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PDDS)}}
}

func (p *PWR_Periph) CWUF() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(CWUF)}}
}

func (p *PWR_Periph) CSBF() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(CSBF)}}
}

func (p *PWR_Periph) PVDE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PVDE)}}
}

func (p *PWR_Periph) PLS() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PLS)}}
}

func (p *PWR_Periph) DBP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBP)}}
}

func (p *PWR_Periph) ULP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(ULP)}}
}

func (p *PWR_Periph) FWU() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(FWU)}}
}

func (p *PWR_Periph) VOS() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(VOS)}}
}

func (p *PWR_Periph) LPRUN() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(LPRUN)}}
}

type CSR uint32

func (b CSR) Field(mask CSR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSR) J(v int) CSR {
	return CSR(bits.MakeField32(v, uint32(mask)))
}

type RCSR struct{ mmio.U32 }

func (r *RCSR) Bits(mask CSR) CSR     { return CSR(r.U32.Bits(uint32(mask))) }
func (r *RCSR) StoreBits(mask, b CSR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCSR) SetBits(mask CSR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCSR) ClearBits(mask CSR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCSR) Load() CSR             { return CSR(r.U32.Load()) }
func (r *RCSR) Store(b CSR)           { r.U32.Store(uint32(b)) }

func (r *RCSR) AtomicStoreBits(mask, b CSR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCSR) AtomicSetBits(mask CSR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCSR) AtomicClearBits(mask CSR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCSR struct{ mmio.UM32 }

func (rm RMCSR) Load() CSR   { return CSR(rm.UM32.Load()) }
func (rm RMCSR) Store(b CSR) { rm.UM32.Store(uint32(b)) }

func (p *PWR_Periph) WUF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(WUF)}}
}

func (p *PWR_Periph) SBF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(SBF)}}
}

func (p *PWR_Periph) PVDO() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(PVDO)}}
}

func (p *PWR_Periph) VREFINTRDYF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(VREFINTRDYF)}}
}

func (p *PWR_Periph) VOSF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(VOSF)}}
}

func (p *PWR_Periph) REGLPF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(REGLPF)}}
}

func (p *PWR_Periph) EWUP1() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(EWUP1)}}
}

func (p *PWR_Periph) EWUP2() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(EWUP2)}}
}

func (p *PWR_Periph) EWUP3() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(EWUP3)}}
}
