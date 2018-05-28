package can

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f746xx/mmap"
)

type CAN_TxMailBox_Periph struct {
	TIR  RTIR
	TDTR RTDTR
	TDLR RTDLR
	TDHR RTDHR
}

func (p *CAN_TxMailBox_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type TIR uint32

func (b TIR) Field(mask TIR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TIR) J(v int) TIR {
	return TIR(bits.MakeField32(v, uint32(mask)))
}

type RTIR struct{ mmio.U32 }

func (r *RTIR) Bits(mask TIR) TIR     { return TIR(r.U32.Bits(uint32(mask))) }
func (r *RTIR) StoreBits(mask, b TIR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTIR) SetBits(mask TIR)      { r.U32.SetBits(uint32(mask)) }
func (r *RTIR) ClearBits(mask TIR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTIR) Load() TIR             { return TIR(r.U32.Load()) }
func (r *RTIR) Store(b TIR)           { r.U32.Store(uint32(b)) }

func (r *RTIR) AtomicStoreBits(mask, b TIR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RTIR) AtomicSetBits(mask TIR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RTIR) AtomicClearBits(mask TIR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMTIR struct{ mmio.UM32 }

func (rm RMTIR) Load() TIR   { return TIR(rm.UM32.Load()) }
func (rm RMTIR) Store(b TIR) { rm.UM32.Store(uint32(b)) }

type TDTR uint32

func (b TDTR) Field(mask TDTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TDTR) J(v int) TDTR {
	return TDTR(bits.MakeField32(v, uint32(mask)))
}

type RTDTR struct{ mmio.U32 }

func (r *RTDTR) Bits(mask TDTR) TDTR    { return TDTR(r.U32.Bits(uint32(mask))) }
func (r *RTDTR) StoreBits(mask, b TDTR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTDTR) SetBits(mask TDTR)      { r.U32.SetBits(uint32(mask)) }
func (r *RTDTR) ClearBits(mask TDTR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTDTR) Load() TDTR             { return TDTR(r.U32.Load()) }
func (r *RTDTR) Store(b TDTR)           { r.U32.Store(uint32(b)) }

func (r *RTDTR) AtomicStoreBits(mask, b TDTR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RTDTR) AtomicSetBits(mask TDTR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RTDTR) AtomicClearBits(mask TDTR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMTDTR struct{ mmio.UM32 }

func (rm RMTDTR) Load() TDTR   { return TDTR(rm.UM32.Load()) }
func (rm RMTDTR) Store(b TDTR) { rm.UM32.Store(uint32(b)) }

type TDLR uint32

func (b TDLR) Field(mask TDLR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TDLR) J(v int) TDLR {
	return TDLR(bits.MakeField32(v, uint32(mask)))
}

type RTDLR struct{ mmio.U32 }

func (r *RTDLR) Bits(mask TDLR) TDLR    { return TDLR(r.U32.Bits(uint32(mask))) }
func (r *RTDLR) StoreBits(mask, b TDLR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTDLR) SetBits(mask TDLR)      { r.U32.SetBits(uint32(mask)) }
func (r *RTDLR) ClearBits(mask TDLR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTDLR) Load() TDLR             { return TDLR(r.U32.Load()) }
func (r *RTDLR) Store(b TDLR)           { r.U32.Store(uint32(b)) }

func (r *RTDLR) AtomicStoreBits(mask, b TDLR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RTDLR) AtomicSetBits(mask TDLR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RTDLR) AtomicClearBits(mask TDLR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMTDLR struct{ mmio.UM32 }

func (rm RMTDLR) Load() TDLR   { return TDLR(rm.UM32.Load()) }
func (rm RMTDLR) Store(b TDLR) { rm.UM32.Store(uint32(b)) }

type TDHR uint32

func (b TDHR) Field(mask TDHR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TDHR) J(v int) TDHR {
	return TDHR(bits.MakeField32(v, uint32(mask)))
}

type RTDHR struct{ mmio.U32 }

func (r *RTDHR) Bits(mask TDHR) TDHR    { return TDHR(r.U32.Bits(uint32(mask))) }
func (r *RTDHR) StoreBits(mask, b TDHR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RTDHR) SetBits(mask TDHR)      { r.U32.SetBits(uint32(mask)) }
func (r *RTDHR) ClearBits(mask TDHR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RTDHR) Load() TDHR             { return TDHR(r.U32.Load()) }
func (r *RTDHR) Store(b TDHR)           { r.U32.Store(uint32(b)) }

func (r *RTDHR) AtomicStoreBits(mask, b TDHR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RTDHR) AtomicSetBits(mask TDHR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RTDHR) AtomicClearBits(mask TDHR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMTDHR struct{ mmio.UM32 }

func (rm RMTDHR) Load() TDHR   { return TDHR(rm.UM32.Load()) }
func (rm RMTDHR) Store(b TDHR) { rm.UM32.Store(uint32(b)) }
