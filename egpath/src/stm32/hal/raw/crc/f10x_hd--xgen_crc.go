// +build f10x_hd

package crc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f10x_hd/mmap"
)

type CRC_Periph struct {
	DR  RDR
	IDR RIDR
	_   uint8
	_   uint16
	CR  RCR
}

func (p *CRC_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var CRC = (*CRC_Periph)(unsafe.Pointer(uintptr(mmap.CRC_BASE)))

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

type IDR uint8

func (b IDR) Field(mask IDR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IDR) J(v int) IDR {
	return IDR(bits.MakeField32(v, uint32(mask)))
}

type RIDR struct{ mmio.U8 }

func (r *RIDR) Bits(mask IDR) IDR     { return IDR(r.U8.Bits(uint8(mask))) }
func (r *RIDR) StoreBits(mask, b IDR) { r.U8.StoreBits(uint8(mask), uint8(b)) }
func (r *RIDR) SetBits(mask IDR)      { r.U8.SetBits(uint8(mask)) }
func (r *RIDR) ClearBits(mask IDR)    { r.U8.ClearBits(uint8(mask)) }
func (r *RIDR) Load() IDR             { return IDR(r.U8.Load()) }
func (r *RIDR) Store(b IDR)           { r.U8.Store(uint8(b)) }

type RMIDR struct{ mmio.UM8 }

func (rm RMIDR) Load() IDR   { return IDR(rm.UM8.Load()) }
func (rm RMIDR) Store(b IDR) { rm.UM8.Store(uint8(b)) }

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

func (p *CRC_Periph) RESET() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(RESET)}}
}
