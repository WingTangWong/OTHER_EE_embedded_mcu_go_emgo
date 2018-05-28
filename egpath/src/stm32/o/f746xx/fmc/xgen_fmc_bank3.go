package fmc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f746xx/mmap"
)

type FMC_Bank3_Periph struct {
	PCR  RPCR
	SR   RSR
	PMEM RPMEM
	PATT RPATT
	_    uint32
	ECCR RECCR
}

func (p *FMC_Bank3_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var FMC_Bank3 = (*FMC_Bank3_Periph)(unsafe.Pointer(uintptr(mmap.FMC_Bank3_R_BASE)))

type PCR uint32

func (b PCR) Field(mask PCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PCR) J(v int) PCR {
	return PCR(bits.MakeField32(v, uint32(mask)))
}

type RPCR struct{ mmio.U32 }

func (r *RPCR) Bits(mask PCR) PCR     { return PCR(r.U32.Bits(uint32(mask))) }
func (r *RPCR) StoreBits(mask, b PCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPCR) SetBits(mask PCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RPCR) ClearBits(mask PCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPCR) Load() PCR             { return PCR(r.U32.Load()) }
func (r *RPCR) Store(b PCR)           { r.U32.Store(uint32(b)) }

func (r *RPCR) AtomicStoreBits(mask, b PCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPCR) AtomicSetBits(mask PCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPCR) AtomicClearBits(mask PCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPCR struct{ mmio.UM32 }

func (rm RMPCR) Load() PCR   { return PCR(rm.UM32.Load()) }
func (rm RMPCR) Store(b PCR) { rm.UM32.Store(uint32(b)) }

func (p *FMC_Bank3_Periph) PWAITEN() RMPCR {
	return RMPCR{mmio.UM32{&p.PCR.U32, uint32(PWAITEN)}}
}

func (p *FMC_Bank3_Periph) PBKEN() RMPCR {
	return RMPCR{mmio.UM32{&p.PCR.U32, uint32(PBKEN)}}
}

func (p *FMC_Bank3_Periph) PTYP() RMPCR {
	return RMPCR{mmio.UM32{&p.PCR.U32, uint32(PTYP)}}
}

func (p *FMC_Bank3_Periph) PWID() RMPCR {
	return RMPCR{mmio.UM32{&p.PCR.U32, uint32(PWID)}}
}

func (p *FMC_Bank3_Periph) ECCEN() RMPCR {
	return RMPCR{mmio.UM32{&p.PCR.U32, uint32(ECCEN)}}
}

func (p *FMC_Bank3_Periph) TCLR() RMPCR {
	return RMPCR{mmio.UM32{&p.PCR.U32, uint32(TCLR)}}
}

func (p *FMC_Bank3_Periph) TAR() RMPCR {
	return RMPCR{mmio.UM32{&p.PCR.U32, uint32(TAR)}}
}

func (p *FMC_Bank3_Periph) ECCPS() RMPCR {
	return RMPCR{mmio.UM32{&p.PCR.U32, uint32(ECCPS)}}
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

func (p *FMC_Bank3_Periph) IRS() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(IRS)}}
}

func (p *FMC_Bank3_Periph) ILS() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(ILS)}}
}

func (p *FMC_Bank3_Periph) IFS() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(IFS)}}
}

func (p *FMC_Bank3_Periph) IREN() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(IREN)}}
}

func (p *FMC_Bank3_Periph) ILEN() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(ILEN)}}
}

func (p *FMC_Bank3_Periph) IFEN() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(IFEN)}}
}

func (p *FMC_Bank3_Periph) FEMPT() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(FEMPT)}}
}

type PMEM uint32

func (b PMEM) Field(mask PMEM) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PMEM) J(v int) PMEM {
	return PMEM(bits.MakeField32(v, uint32(mask)))
}

type RPMEM struct{ mmio.U32 }

func (r *RPMEM) Bits(mask PMEM) PMEM    { return PMEM(r.U32.Bits(uint32(mask))) }
func (r *RPMEM) StoreBits(mask, b PMEM) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPMEM) SetBits(mask PMEM)      { r.U32.SetBits(uint32(mask)) }
func (r *RPMEM) ClearBits(mask PMEM)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPMEM) Load() PMEM             { return PMEM(r.U32.Load()) }
func (r *RPMEM) Store(b PMEM)           { r.U32.Store(uint32(b)) }

func (r *RPMEM) AtomicStoreBits(mask, b PMEM) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPMEM) AtomicSetBits(mask PMEM)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPMEM) AtomicClearBits(mask PMEM)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPMEM struct{ mmio.UM32 }

func (rm RMPMEM) Load() PMEM   { return PMEM(rm.UM32.Load()) }
func (rm RMPMEM) Store(b PMEM) { rm.UM32.Store(uint32(b)) }

func (p *FMC_Bank3_Periph) MEMSET3() RMPMEM {
	return RMPMEM{mmio.UM32{&p.PMEM.U32, uint32(MEMSET3)}}
}

func (p *FMC_Bank3_Periph) MEMWAIT3() RMPMEM {
	return RMPMEM{mmio.UM32{&p.PMEM.U32, uint32(MEMWAIT3)}}
}

func (p *FMC_Bank3_Periph) MEMHOLD3() RMPMEM {
	return RMPMEM{mmio.UM32{&p.PMEM.U32, uint32(MEMHOLD3)}}
}

func (p *FMC_Bank3_Periph) MEMHIZ3() RMPMEM {
	return RMPMEM{mmio.UM32{&p.PMEM.U32, uint32(MEMHIZ3)}}
}

type PATT uint32

func (b PATT) Field(mask PATT) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PATT) J(v int) PATT {
	return PATT(bits.MakeField32(v, uint32(mask)))
}

type RPATT struct{ mmio.U32 }

func (r *RPATT) Bits(mask PATT) PATT    { return PATT(r.U32.Bits(uint32(mask))) }
func (r *RPATT) StoreBits(mask, b PATT) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPATT) SetBits(mask PATT)      { r.U32.SetBits(uint32(mask)) }
func (r *RPATT) ClearBits(mask PATT)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPATT) Load() PATT             { return PATT(r.U32.Load()) }
func (r *RPATT) Store(b PATT)           { r.U32.Store(uint32(b)) }

func (r *RPATT) AtomicStoreBits(mask, b PATT) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPATT) AtomicSetBits(mask PATT)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPATT) AtomicClearBits(mask PATT)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPATT struct{ mmio.UM32 }

func (rm RMPATT) Load() PATT   { return PATT(rm.UM32.Load()) }
func (rm RMPATT) Store(b PATT) { rm.UM32.Store(uint32(b)) }

func (p *FMC_Bank3_Periph) ATTSET3() RMPATT {
	return RMPATT{mmio.UM32{&p.PATT.U32, uint32(ATTSET3)}}
}

func (p *FMC_Bank3_Periph) ATTWAIT3() RMPATT {
	return RMPATT{mmio.UM32{&p.PATT.U32, uint32(ATTWAIT3)}}
}

func (p *FMC_Bank3_Periph) ATTHOLD3() RMPATT {
	return RMPATT{mmio.UM32{&p.PATT.U32, uint32(ATTHOLD3)}}
}

func (p *FMC_Bank3_Periph) ATTHIZ3() RMPATT {
	return RMPATT{mmio.UM32{&p.PATT.U32, uint32(ATTHIZ3)}}
}

type ECCR uint32

func (b ECCR) Field(mask ECCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ECCR) J(v int) ECCR {
	return ECCR(bits.MakeField32(v, uint32(mask)))
}

type RECCR struct{ mmio.U32 }

func (r *RECCR) Bits(mask ECCR) ECCR    { return ECCR(r.U32.Bits(uint32(mask))) }
func (r *RECCR) StoreBits(mask, b ECCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RECCR) SetBits(mask ECCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RECCR) ClearBits(mask ECCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RECCR) Load() ECCR             { return ECCR(r.U32.Load()) }
func (r *RECCR) Store(b ECCR)           { r.U32.Store(uint32(b)) }

func (r *RECCR) AtomicStoreBits(mask, b ECCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RECCR) AtomicSetBits(mask ECCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RECCR) AtomicClearBits(mask ECCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMECCR struct{ mmio.UM32 }

func (rm RMECCR) Load() ECCR   { return ECCR(rm.UM32.Load()) }
func (rm RMECCR) Store(b ECCR) { rm.UM32.Store(uint32(b)) }

func (p *FMC_Bank3_Periph) ECC3() RMECCR {
	return RMECCR{mmio.UM32{&p.ECCR.U32, uint32(ECC3)}}
}
