package fsmc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f10x_hd/mmap"
)

type FSMC_Bank1E_Periph struct {
	BWTR [7]RBWTR
}

func (p *FSMC_Bank1E_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var FSMC_Bank1E = (*FSMC_Bank1E_Periph)(unsafe.Pointer(uintptr(mmap.FSMC_Bank1E_R_BASE)))

type BWTR uint32

func (b BWTR) Field(mask BWTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BWTR) J(v int) BWTR {
	return BWTR(bits.MakeField32(v, uint32(mask)))
}

type RBWTR struct{ mmio.U32 }

func (r *RBWTR) Bits(mask BWTR) BWTR    { return BWTR(r.U32.Bits(uint32(mask))) }
func (r *RBWTR) StoreBits(mask, b BWTR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RBWTR) SetBits(mask BWTR)      { r.U32.SetBits(uint32(mask)) }
func (r *RBWTR) ClearBits(mask BWTR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RBWTR) Load() BWTR             { return BWTR(r.U32.Load()) }
func (r *RBWTR) Store(b BWTR)           { r.U32.Store(uint32(b)) }

func (r *RBWTR) AtomicStoreBits(mask, b BWTR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RBWTR) AtomicSetBits(mask BWTR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RBWTR) AtomicClearBits(mask BWTR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMBWTR struct{ mmio.UM32 }

func (rm RMBWTR) Load() BWTR   { return BWTR(rm.UM32.Load()) }
func (rm RMBWTR) Store(b BWTR) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank1E_Periph) EADDSET(n int) RMBWTR {
	return RMBWTR{mmio.UM32{&p.BWTR[n].U32, uint32(EADDSET)}}
}

func (p *FSMC_Bank1E_Periph) EADDHLD(n int) RMBWTR {
	return RMBWTR{mmio.UM32{&p.BWTR[n].U32, uint32(EADDHLD)}}
}

func (p *FSMC_Bank1E_Periph) EDATAST(n int) RMBWTR {
	return RMBWTR{mmio.UM32{&p.BWTR[n].U32, uint32(EDATAST)}}
}

func (p *FSMC_Bank1E_Periph) ECLKDIV(n int) RMBWTR {
	return RMBWTR{mmio.UM32{&p.BWTR[n].U32, uint32(ECLKDIV)}}
}

func (p *FSMC_Bank1E_Periph) EDATLAT(n int) RMBWTR {
	return RMBWTR{mmio.UM32{&p.BWTR[n].U32, uint32(EDATLAT)}}
}

func (p *FSMC_Bank1E_Periph) EACCMOD(n int) RMBWTR {
	return RMBWTR{mmio.UM32{&p.BWTR[n].U32, uint32(EACCMOD)}}
}
