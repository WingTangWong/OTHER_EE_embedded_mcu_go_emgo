// +build l476xx

package comp

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type COMP_Common_Periph struct {
	CSR RCSR
}

func (p *COMP_Common_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var COMP12_COMMON = (*COMP_Common_Periph)(unsafe.Pointer(uintptr(mmap.COMP2_BASE)))

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
