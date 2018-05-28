package fsmc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f10x_hd/mmap"
)

type FSMC_Bank4_Periph struct {
	PCR4  RPCR4
	SR4   RSR4
	PMEM4 RPMEM4
	PATT4 RPATT4
	PIO4  RPIO4
}

func (p *FSMC_Bank4_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var FSMC_Bank4 = (*FSMC_Bank4_Periph)(unsafe.Pointer(uintptr(mmap.FSMC_Bank4_R_BASE)))

type PCR4 uint32

func (b PCR4) Field(mask PCR4) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PCR4) J(v int) PCR4 {
	return PCR4(bits.MakeField32(v, uint32(mask)))
}

type RPCR4 struct{ mmio.U32 }

func (r *RPCR4) Bits(mask PCR4) PCR4    { return PCR4(r.U32.Bits(uint32(mask))) }
func (r *RPCR4) StoreBits(mask, b PCR4) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPCR4) SetBits(mask PCR4)      { r.U32.SetBits(uint32(mask)) }
func (r *RPCR4) ClearBits(mask PCR4)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPCR4) Load() PCR4             { return PCR4(r.U32.Load()) }
func (r *RPCR4) Store(b PCR4)           { r.U32.Store(uint32(b)) }

func (r *RPCR4) AtomicStoreBits(mask, b PCR4) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPCR4) AtomicSetBits(mask PCR4)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPCR4) AtomicClearBits(mask PCR4)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPCR4 struct{ mmio.UM32 }

func (rm RMPCR4) Load() PCR4   { return PCR4(rm.UM32.Load()) }
func (rm RMPCR4) Store(b PCR4) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank4_Periph) PWAITEN() RMPCR4 {
	return RMPCR4{mmio.UM32{&p.PCR4.U32, uint32(PWAITEN)}}
}

func (p *FSMC_Bank4_Periph) PBKEN() RMPCR4 {
	return RMPCR4{mmio.UM32{&p.PCR4.U32, uint32(PBKEN)}}
}

func (p *FSMC_Bank4_Periph) PTYP() RMPCR4 {
	return RMPCR4{mmio.UM32{&p.PCR4.U32, uint32(PTYP)}}
}

func (p *FSMC_Bank4_Periph) PWID() RMPCR4 {
	return RMPCR4{mmio.UM32{&p.PCR4.U32, uint32(PWID)}}
}

func (p *FSMC_Bank4_Periph) ECCEN() RMPCR4 {
	return RMPCR4{mmio.UM32{&p.PCR4.U32, uint32(ECCEN)}}
}

func (p *FSMC_Bank4_Periph) TCLR() RMPCR4 {
	return RMPCR4{mmio.UM32{&p.PCR4.U32, uint32(TCLR)}}
}

func (p *FSMC_Bank4_Periph) TAR() RMPCR4 {
	return RMPCR4{mmio.UM32{&p.PCR4.U32, uint32(TAR)}}
}

func (p *FSMC_Bank4_Periph) ECCPS() RMPCR4 {
	return RMPCR4{mmio.UM32{&p.PCR4.U32, uint32(ECCPS)}}
}

type SR4 uint32

func (b SR4) Field(mask SR4) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR4) J(v int) SR4 {
	return SR4(bits.MakeField32(v, uint32(mask)))
}

type RSR4 struct{ mmio.U32 }

func (r *RSR4) Bits(mask SR4) SR4     { return SR4(r.U32.Bits(uint32(mask))) }
func (r *RSR4) StoreBits(mask, b SR4) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSR4) SetBits(mask SR4)      { r.U32.SetBits(uint32(mask)) }
func (r *RSR4) ClearBits(mask SR4)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSR4) Load() SR4             { return SR4(r.U32.Load()) }
func (r *RSR4) Store(b SR4)           { r.U32.Store(uint32(b)) }

func (r *RSR4) AtomicStoreBits(mask, b SR4) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSR4) AtomicSetBits(mask SR4)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSR4) AtomicClearBits(mask SR4)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSR4 struct{ mmio.UM32 }

func (rm RMSR4) Load() SR4   { return SR4(rm.UM32.Load()) }
func (rm RMSR4) Store(b SR4) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank4_Periph) IRS() RMSR4 {
	return RMSR4{mmio.UM32{&p.SR4.U32, uint32(IRS)}}
}

func (p *FSMC_Bank4_Periph) ILS() RMSR4 {
	return RMSR4{mmio.UM32{&p.SR4.U32, uint32(ILS)}}
}

func (p *FSMC_Bank4_Periph) IFS() RMSR4 {
	return RMSR4{mmio.UM32{&p.SR4.U32, uint32(IFS)}}
}

func (p *FSMC_Bank4_Periph) IREN() RMSR4 {
	return RMSR4{mmio.UM32{&p.SR4.U32, uint32(IREN)}}
}

func (p *FSMC_Bank4_Periph) ILEN() RMSR4 {
	return RMSR4{mmio.UM32{&p.SR4.U32, uint32(ILEN)}}
}

func (p *FSMC_Bank4_Periph) IFEN() RMSR4 {
	return RMSR4{mmio.UM32{&p.SR4.U32, uint32(IFEN)}}
}

func (p *FSMC_Bank4_Periph) FEMPT() RMSR4 {
	return RMSR4{mmio.UM32{&p.SR4.U32, uint32(FEMPT)}}
}

type PMEM4 uint32

func (b PMEM4) Field(mask PMEM4) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PMEM4) J(v int) PMEM4 {
	return PMEM4(bits.MakeField32(v, uint32(mask)))
}

type RPMEM4 struct{ mmio.U32 }

func (r *RPMEM4) Bits(mask PMEM4) PMEM4   { return PMEM4(r.U32.Bits(uint32(mask))) }
func (r *RPMEM4) StoreBits(mask, b PMEM4) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPMEM4) SetBits(mask PMEM4)      { r.U32.SetBits(uint32(mask)) }
func (r *RPMEM4) ClearBits(mask PMEM4)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPMEM4) Load() PMEM4             { return PMEM4(r.U32.Load()) }
func (r *RPMEM4) Store(b PMEM4)           { r.U32.Store(uint32(b)) }

func (r *RPMEM4) AtomicStoreBits(mask, b PMEM4) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPMEM4) AtomicSetBits(mask PMEM4)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPMEM4) AtomicClearBits(mask PMEM4)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPMEM4 struct{ mmio.UM32 }

func (rm RMPMEM4) Load() PMEM4   { return PMEM4(rm.UM32.Load()) }
func (rm RMPMEM4) Store(b PMEM4) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank4_Periph) MEMSET4() RMPMEM4 {
	return RMPMEM4{mmio.UM32{&p.PMEM4.U32, uint32(MEMSET4)}}
}

func (p *FSMC_Bank4_Periph) MEMWAIT4() RMPMEM4 {
	return RMPMEM4{mmio.UM32{&p.PMEM4.U32, uint32(MEMWAIT4)}}
}

func (p *FSMC_Bank4_Periph) MEMHOLD4() RMPMEM4 {
	return RMPMEM4{mmio.UM32{&p.PMEM4.U32, uint32(MEMHOLD4)}}
}

func (p *FSMC_Bank4_Periph) MEMHIZ4() RMPMEM4 {
	return RMPMEM4{mmio.UM32{&p.PMEM4.U32, uint32(MEMHIZ4)}}
}

type PATT4 uint32

func (b PATT4) Field(mask PATT4) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PATT4) J(v int) PATT4 {
	return PATT4(bits.MakeField32(v, uint32(mask)))
}

type RPATT4 struct{ mmio.U32 }

func (r *RPATT4) Bits(mask PATT4) PATT4   { return PATT4(r.U32.Bits(uint32(mask))) }
func (r *RPATT4) StoreBits(mask, b PATT4) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPATT4) SetBits(mask PATT4)      { r.U32.SetBits(uint32(mask)) }
func (r *RPATT4) ClearBits(mask PATT4)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPATT4) Load() PATT4             { return PATT4(r.U32.Load()) }
func (r *RPATT4) Store(b PATT4)           { r.U32.Store(uint32(b)) }

func (r *RPATT4) AtomicStoreBits(mask, b PATT4) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPATT4) AtomicSetBits(mask PATT4)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPATT4) AtomicClearBits(mask PATT4)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPATT4 struct{ mmio.UM32 }

func (rm RMPATT4) Load() PATT4   { return PATT4(rm.UM32.Load()) }
func (rm RMPATT4) Store(b PATT4) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank4_Periph) ATTSET4() RMPATT4 {
	return RMPATT4{mmio.UM32{&p.PATT4.U32, uint32(ATTSET4)}}
}

func (p *FSMC_Bank4_Periph) ATTWAIT4() RMPATT4 {
	return RMPATT4{mmio.UM32{&p.PATT4.U32, uint32(ATTWAIT4)}}
}

func (p *FSMC_Bank4_Periph) ATTHOLD4() RMPATT4 {
	return RMPATT4{mmio.UM32{&p.PATT4.U32, uint32(ATTHOLD4)}}
}

func (p *FSMC_Bank4_Periph) ATTHIZ4() RMPATT4 {
	return RMPATT4{mmio.UM32{&p.PATT4.U32, uint32(ATTHIZ4)}}
}

type PIO4 uint32

func (b PIO4) Field(mask PIO4) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PIO4) J(v int) PIO4 {
	return PIO4(bits.MakeField32(v, uint32(mask)))
}

type RPIO4 struct{ mmio.U32 }

func (r *RPIO4) Bits(mask PIO4) PIO4    { return PIO4(r.U32.Bits(uint32(mask))) }
func (r *RPIO4) StoreBits(mask, b PIO4) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPIO4) SetBits(mask PIO4)      { r.U32.SetBits(uint32(mask)) }
func (r *RPIO4) ClearBits(mask PIO4)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPIO4) Load() PIO4             { return PIO4(r.U32.Load()) }
func (r *RPIO4) Store(b PIO4)           { r.U32.Store(uint32(b)) }

func (r *RPIO4) AtomicStoreBits(mask, b PIO4) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPIO4) AtomicSetBits(mask PIO4)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPIO4) AtomicClearBits(mask PIO4)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPIO4 struct{ mmio.UM32 }

func (rm RMPIO4) Load() PIO4   { return PIO4(rm.UM32.Load()) }
func (rm RMPIO4) Store(b PIO4) { rm.UM32.Store(uint32(b)) }

func (p *FSMC_Bank4_Periph) IOSET4() RMPIO4 {
	return RMPIO4{mmio.UM32{&p.PIO4.U32, uint32(IOSET4)}}
}

func (p *FSMC_Bank4_Periph) IOWAIT4() RMPIO4 {
	return RMPIO4{mmio.UM32{&p.PIO4.U32, uint32(IOWAIT4)}}
}

func (p *FSMC_Bank4_Periph) IOHOLD4() RMPIO4 {
	return RMPIO4{mmio.UM32{&p.PIO4.U32, uint32(IOHOLD4)}}
}

func (p *FSMC_Bank4_Periph) IOHIZ4() RMPIO4 {
	return RMPIO4{mmio.UM32{&p.PIO4.U32, uint32(IOHIZ4)}}
}
