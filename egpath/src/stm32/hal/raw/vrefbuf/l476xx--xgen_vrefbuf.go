// +build l476xx

package vrefbuf

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type VREFBUF_Periph struct {
	CSR RCSR
	CCR RCCR
}

func (p *VREFBUF_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var VREFBUF = (*VREFBUF_Periph)(unsafe.Pointer(uintptr(mmap.VREFBUF_BASE)))

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

func (p *VREFBUF_Periph) ENVR() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(ENVR)}}
}

func (p *VREFBUF_Periph) HIZ() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(HIZ)}}
}

func (p *VREFBUF_Periph) VRS() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(VRS)}}
}

func (p *VREFBUF_Periph) VRR() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(VRR)}}
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

func (p *VREFBUF_Periph) TRIM() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(TRIM)}}
}
