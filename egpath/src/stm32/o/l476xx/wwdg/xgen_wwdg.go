package wwdg

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type WWDG_Periph struct {
	CR  RCR
	CFR RCFR
	SR  RSR
}

func (p *WWDG_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var WWDG = (*WWDG_Periph)(unsafe.Pointer(uintptr(mmap.WWDG_BASE)))

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

func (p *WWDG_Periph) T() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(T)}}
}

func (p *WWDG_Periph) WDGA() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(WDGA)}}
}

type CFR uint32

func (b CFR) Field(mask CFR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CFR) J(v int) CFR {
	return CFR(bits.MakeField32(v, uint32(mask)))
}

type RCFR struct{ mmio.U32 }

func (r *RCFR) Bits(mask CFR) CFR     { return CFR(r.U32.Bits(uint32(mask))) }
func (r *RCFR) StoreBits(mask, b CFR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCFR) SetBits(mask CFR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCFR) ClearBits(mask CFR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCFR) Load() CFR             { return CFR(r.U32.Load()) }
func (r *RCFR) Store(b CFR)           { r.U32.Store(uint32(b)) }

func (r *RCFR) AtomicStoreBits(mask, b CFR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCFR) AtomicSetBits(mask CFR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCFR) AtomicClearBits(mask CFR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCFR struct{ mmio.UM32 }

func (rm RMCFR) Load() CFR   { return CFR(rm.UM32.Load()) }
func (rm RMCFR) Store(b CFR) { rm.UM32.Store(uint32(b)) }

func (p *WWDG_Periph) W() RMCFR {
	return RMCFR{mmio.UM32{&p.CFR.U32, uint32(W)}}
}

func (p *WWDG_Periph) WDGTB() RMCFR {
	return RMCFR{mmio.UM32{&p.CFR.U32, uint32(WDGTB)}}
}

func (p *WWDG_Periph) EWI() RMCFR {
	return RMCFR{mmio.UM32{&p.CFR.U32, uint32(EWI)}}
}

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

func (p *WWDG_Periph) EWIF() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(EWIF)}}
}
