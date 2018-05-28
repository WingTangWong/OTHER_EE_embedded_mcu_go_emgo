package dac

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f40_41xxx/mmap"
)

type DAC_Periph struct {
	CR      RCR
	SWTRIGR RSWTRIGR
	DHR12R1 RDHR12R1
	DHR12L1 RDHR12L1
	DHR8R1  RDHR8R1
	DHR12R2 RDHR12R2
	DHR12L2 RDHR12L2
	DHR8R2  RDHR8R2
	DHR12RD RDHR12RD
	DHR12LD RDHR12LD
	DHR8RD  RDHR8RD
	DOR1    RDOR1
	DOR2    RDOR2
	SR      RSR
}

func (p *DAC_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var DAC = (*DAC_Periph)(unsafe.Pointer(uintptr(mmap.DAC_BASE)))

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

func (p *DAC_Periph) EN1() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(EN1)}}
}

func (p *DAC_Periph) BOFF1() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(BOFF1)}}
}

func (p *DAC_Periph) TEN1() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TEN1)}}
}

func (p *DAC_Periph) TSEL1() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TSEL1)}}
}

func (p *DAC_Periph) WAVE1() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(WAVE1)}}
}

func (p *DAC_Periph) MAMP1() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(MAMP1)}}
}

func (p *DAC_Periph) DMAEN1() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DMAEN1)}}
}

func (p *DAC_Periph) DMAUDRIE1() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DMAUDRIE1)}}
}

func (p *DAC_Periph) EN2() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(EN2)}}
}

func (p *DAC_Periph) BOFF2() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(BOFF2)}}
}

func (p *DAC_Periph) TEN2() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TEN2)}}
}

func (p *DAC_Periph) TSEL2() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TSEL2)}}
}

func (p *DAC_Periph) WAVE2() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(WAVE2)}}
}

func (p *DAC_Periph) MAMP2() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(MAMP2)}}
}

func (p *DAC_Periph) DMAEN2() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DMAEN2)}}
}

func (p *DAC_Periph) DMAUDRIE2() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DMAUDRIE2)}}
}

type SWTRIGR uint32

func (b SWTRIGR) Field(mask SWTRIGR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SWTRIGR) J(v int) SWTRIGR {
	return SWTRIGR(bits.MakeField32(v, uint32(mask)))
}

type RSWTRIGR struct{ mmio.U32 }

func (r *RSWTRIGR) Bits(mask SWTRIGR) SWTRIGR { return SWTRIGR(r.U32.Bits(uint32(mask))) }
func (r *RSWTRIGR) StoreBits(mask, b SWTRIGR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSWTRIGR) SetBits(mask SWTRIGR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSWTRIGR) ClearBits(mask SWTRIGR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSWTRIGR) Load() SWTRIGR             { return SWTRIGR(r.U32.Load()) }
func (r *RSWTRIGR) Store(b SWTRIGR)           { r.U32.Store(uint32(b)) }

func (r *RSWTRIGR) AtomicStoreBits(mask, b SWTRIGR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSWTRIGR) AtomicSetBits(mask SWTRIGR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSWTRIGR) AtomicClearBits(mask SWTRIGR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSWTRIGR struct{ mmio.UM32 }

func (rm RMSWTRIGR) Load() SWTRIGR   { return SWTRIGR(rm.UM32.Load()) }
func (rm RMSWTRIGR) Store(b SWTRIGR) { rm.UM32.Store(uint32(b)) }

func (p *DAC_Periph) SWTRIG1() RMSWTRIGR {
	return RMSWTRIGR{mmio.UM32{&p.SWTRIGR.U32, uint32(SWTRIG1)}}
}

func (p *DAC_Periph) SWTRIG2() RMSWTRIGR {
	return RMSWTRIGR{mmio.UM32{&p.SWTRIGR.U32, uint32(SWTRIG2)}}
}

type DHR12R1 uint32

func (b DHR12R1) Field(mask DHR12R1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DHR12R1) J(v int) DHR12R1 {
	return DHR12R1(bits.MakeField32(v, uint32(mask)))
}

type RDHR12R1 struct{ mmio.U32 }

func (r *RDHR12R1) Bits(mask DHR12R1) DHR12R1 { return DHR12R1(r.U32.Bits(uint32(mask))) }
func (r *RDHR12R1) StoreBits(mask, b DHR12R1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDHR12R1) SetBits(mask DHR12R1)      { r.U32.SetBits(uint32(mask)) }
func (r *RDHR12R1) ClearBits(mask DHR12R1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDHR12R1) Load() DHR12R1             { return DHR12R1(r.U32.Load()) }
func (r *RDHR12R1) Store(b DHR12R1)           { r.U32.Store(uint32(b)) }

func (r *RDHR12R1) AtomicStoreBits(mask, b DHR12R1) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDHR12R1) AtomicSetBits(mask DHR12R1)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDHR12R1) AtomicClearBits(mask DHR12R1)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDHR12R1 struct{ mmio.UM32 }

func (rm RMDHR12R1) Load() DHR12R1   { return DHR12R1(rm.UM32.Load()) }
func (rm RMDHR12R1) Store(b DHR12R1) { rm.UM32.Store(uint32(b)) }

func (p *DAC_Periph) DACC1DHR() RMDHR12R1 {
	return RMDHR12R1{mmio.UM32{&p.DHR12R1.U32, uint32(DACC1DHR)}}
}

type DHR12L1 uint32

func (b DHR12L1) Field(mask DHR12L1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DHR12L1) J(v int) DHR12L1 {
	return DHR12L1(bits.MakeField32(v, uint32(mask)))
}

type RDHR12L1 struct{ mmio.U32 }

func (r *RDHR12L1) Bits(mask DHR12L1) DHR12L1 { return DHR12L1(r.U32.Bits(uint32(mask))) }
func (r *RDHR12L1) StoreBits(mask, b DHR12L1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDHR12L1) SetBits(mask DHR12L1)      { r.U32.SetBits(uint32(mask)) }
func (r *RDHR12L1) ClearBits(mask DHR12L1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDHR12L1) Load() DHR12L1             { return DHR12L1(r.U32.Load()) }
func (r *RDHR12L1) Store(b DHR12L1)           { r.U32.Store(uint32(b)) }

func (r *RDHR12L1) AtomicStoreBits(mask, b DHR12L1) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDHR12L1) AtomicSetBits(mask DHR12L1)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDHR12L1) AtomicClearBits(mask DHR12L1)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDHR12L1 struct{ mmio.UM32 }

func (rm RMDHR12L1) Load() DHR12L1   { return DHR12L1(rm.UM32.Load()) }
func (rm RMDHR12L1) Store(b DHR12L1) { rm.UM32.Store(uint32(b)) }

func (p *DAC_Periph) DACC1DHR() RMDHR12L1 {
	return RMDHR12L1{mmio.UM32{&p.DHR12L1.U32, uint32(DACC1DHR)}}
}

type DHR8R1 uint32

func (b DHR8R1) Field(mask DHR8R1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DHR8R1) J(v int) DHR8R1 {
	return DHR8R1(bits.MakeField32(v, uint32(mask)))
}

type RDHR8R1 struct{ mmio.U32 }

func (r *RDHR8R1) Bits(mask DHR8R1) DHR8R1  { return DHR8R1(r.U32.Bits(uint32(mask))) }
func (r *RDHR8R1) StoreBits(mask, b DHR8R1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDHR8R1) SetBits(mask DHR8R1)      { r.U32.SetBits(uint32(mask)) }
func (r *RDHR8R1) ClearBits(mask DHR8R1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDHR8R1) Load() DHR8R1             { return DHR8R1(r.U32.Load()) }
func (r *RDHR8R1) Store(b DHR8R1)           { r.U32.Store(uint32(b)) }

func (r *RDHR8R1) AtomicStoreBits(mask, b DHR8R1) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDHR8R1) AtomicSetBits(mask DHR8R1)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDHR8R1) AtomicClearBits(mask DHR8R1)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDHR8R1 struct{ mmio.UM32 }

func (rm RMDHR8R1) Load() DHR8R1   { return DHR8R1(rm.UM32.Load()) }
func (rm RMDHR8R1) Store(b DHR8R1) { rm.UM32.Store(uint32(b)) }

func (p *DAC_Periph) DACC1DHR() RMDHR8R1 {
	return RMDHR8R1{mmio.UM32{&p.DHR8R1.U32, uint32(DACC1DHR)}}
}

type DHR12R2 uint32

func (b DHR12R2) Field(mask DHR12R2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DHR12R2) J(v int) DHR12R2 {
	return DHR12R2(bits.MakeField32(v, uint32(mask)))
}

type RDHR12R2 struct{ mmio.U32 }

func (r *RDHR12R2) Bits(mask DHR12R2) DHR12R2 { return DHR12R2(r.U32.Bits(uint32(mask))) }
func (r *RDHR12R2) StoreBits(mask, b DHR12R2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDHR12R2) SetBits(mask DHR12R2)      { r.U32.SetBits(uint32(mask)) }
func (r *RDHR12R2) ClearBits(mask DHR12R2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDHR12R2) Load() DHR12R2             { return DHR12R2(r.U32.Load()) }
func (r *RDHR12R2) Store(b DHR12R2)           { r.U32.Store(uint32(b)) }

func (r *RDHR12R2) AtomicStoreBits(mask, b DHR12R2) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDHR12R2) AtomicSetBits(mask DHR12R2)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDHR12R2) AtomicClearBits(mask DHR12R2)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDHR12R2 struct{ mmio.UM32 }

func (rm RMDHR12R2) Load() DHR12R2   { return DHR12R2(rm.UM32.Load()) }
func (rm RMDHR12R2) Store(b DHR12R2) { rm.UM32.Store(uint32(b)) }

func (p *DAC_Periph) DACC2DHR() RMDHR12R2 {
	return RMDHR12R2{mmio.UM32{&p.DHR12R2.U32, uint32(DACC2DHR)}}
}

type DHR12L2 uint32

func (b DHR12L2) Field(mask DHR12L2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DHR12L2) J(v int) DHR12L2 {
	return DHR12L2(bits.MakeField32(v, uint32(mask)))
}

type RDHR12L2 struct{ mmio.U32 }

func (r *RDHR12L2) Bits(mask DHR12L2) DHR12L2 { return DHR12L2(r.U32.Bits(uint32(mask))) }
func (r *RDHR12L2) StoreBits(mask, b DHR12L2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDHR12L2) SetBits(mask DHR12L2)      { r.U32.SetBits(uint32(mask)) }
func (r *RDHR12L2) ClearBits(mask DHR12L2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDHR12L2) Load() DHR12L2             { return DHR12L2(r.U32.Load()) }
func (r *RDHR12L2) Store(b DHR12L2)           { r.U32.Store(uint32(b)) }

func (r *RDHR12L2) AtomicStoreBits(mask, b DHR12L2) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDHR12L2) AtomicSetBits(mask DHR12L2)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDHR12L2) AtomicClearBits(mask DHR12L2)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDHR12L2 struct{ mmio.UM32 }

func (rm RMDHR12L2) Load() DHR12L2   { return DHR12L2(rm.UM32.Load()) }
func (rm RMDHR12L2) Store(b DHR12L2) { rm.UM32.Store(uint32(b)) }

func (p *DAC_Periph) DACC2DHR() RMDHR12L2 {
	return RMDHR12L2{mmio.UM32{&p.DHR12L2.U32, uint32(DACC2DHR)}}
}

type DHR8R2 uint32

func (b DHR8R2) Field(mask DHR8R2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DHR8R2) J(v int) DHR8R2 {
	return DHR8R2(bits.MakeField32(v, uint32(mask)))
}

type RDHR8R2 struct{ mmio.U32 }

func (r *RDHR8R2) Bits(mask DHR8R2) DHR8R2  { return DHR8R2(r.U32.Bits(uint32(mask))) }
func (r *RDHR8R2) StoreBits(mask, b DHR8R2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDHR8R2) SetBits(mask DHR8R2)      { r.U32.SetBits(uint32(mask)) }
func (r *RDHR8R2) ClearBits(mask DHR8R2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDHR8R2) Load() DHR8R2             { return DHR8R2(r.U32.Load()) }
func (r *RDHR8R2) Store(b DHR8R2)           { r.U32.Store(uint32(b)) }

func (r *RDHR8R2) AtomicStoreBits(mask, b DHR8R2) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDHR8R2) AtomicSetBits(mask DHR8R2)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDHR8R2) AtomicClearBits(mask DHR8R2)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDHR8R2 struct{ mmio.UM32 }

func (rm RMDHR8R2) Load() DHR8R2   { return DHR8R2(rm.UM32.Load()) }
func (rm RMDHR8R2) Store(b DHR8R2) { rm.UM32.Store(uint32(b)) }

func (p *DAC_Periph) DACC2DHR() RMDHR8R2 {
	return RMDHR8R2{mmio.UM32{&p.DHR8R2.U32, uint32(DACC2DHR)}}
}

type DHR12RD uint32

func (b DHR12RD) Field(mask DHR12RD) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DHR12RD) J(v int) DHR12RD {
	return DHR12RD(bits.MakeField32(v, uint32(mask)))
}

type RDHR12RD struct{ mmio.U32 }

func (r *RDHR12RD) Bits(mask DHR12RD) DHR12RD { return DHR12RD(r.U32.Bits(uint32(mask))) }
func (r *RDHR12RD) StoreBits(mask, b DHR12RD) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDHR12RD) SetBits(mask DHR12RD)      { r.U32.SetBits(uint32(mask)) }
func (r *RDHR12RD) ClearBits(mask DHR12RD)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDHR12RD) Load() DHR12RD             { return DHR12RD(r.U32.Load()) }
func (r *RDHR12RD) Store(b DHR12RD)           { r.U32.Store(uint32(b)) }

func (r *RDHR12RD) AtomicStoreBits(mask, b DHR12RD) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDHR12RD) AtomicSetBits(mask DHR12RD)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDHR12RD) AtomicClearBits(mask DHR12RD)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDHR12RD struct{ mmio.UM32 }

func (rm RMDHR12RD) Load() DHR12RD   { return DHR12RD(rm.UM32.Load()) }
func (rm RMDHR12RD) Store(b DHR12RD) { rm.UM32.Store(uint32(b)) }

func (p *DAC_Periph) DACC1DHR() RMDHR12RD {
	return RMDHR12RD{mmio.UM32{&p.DHR12RD.U32, uint32(DACC1DHR)}}
}

func (p *DAC_Periph) DACC2DHR() RMDHR12RD {
	return RMDHR12RD{mmio.UM32{&p.DHR12RD.U32, uint32(DACC2DHR)}}
}

type DHR12LD uint32

func (b DHR12LD) Field(mask DHR12LD) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DHR12LD) J(v int) DHR12LD {
	return DHR12LD(bits.MakeField32(v, uint32(mask)))
}

type RDHR12LD struct{ mmio.U32 }

func (r *RDHR12LD) Bits(mask DHR12LD) DHR12LD { return DHR12LD(r.U32.Bits(uint32(mask))) }
func (r *RDHR12LD) StoreBits(mask, b DHR12LD) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDHR12LD) SetBits(mask DHR12LD)      { r.U32.SetBits(uint32(mask)) }
func (r *RDHR12LD) ClearBits(mask DHR12LD)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDHR12LD) Load() DHR12LD             { return DHR12LD(r.U32.Load()) }
func (r *RDHR12LD) Store(b DHR12LD)           { r.U32.Store(uint32(b)) }

func (r *RDHR12LD) AtomicStoreBits(mask, b DHR12LD) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDHR12LD) AtomicSetBits(mask DHR12LD)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDHR12LD) AtomicClearBits(mask DHR12LD)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDHR12LD struct{ mmio.UM32 }

func (rm RMDHR12LD) Load() DHR12LD   { return DHR12LD(rm.UM32.Load()) }
func (rm RMDHR12LD) Store(b DHR12LD) { rm.UM32.Store(uint32(b)) }

func (p *DAC_Periph) DACC1DHR() RMDHR12LD {
	return RMDHR12LD{mmio.UM32{&p.DHR12LD.U32, uint32(DACC1DHR)}}
}

func (p *DAC_Periph) DACC2DHR() RMDHR12LD {
	return RMDHR12LD{mmio.UM32{&p.DHR12LD.U32, uint32(DACC2DHR)}}
}

type DHR8RD uint32

func (b DHR8RD) Field(mask DHR8RD) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DHR8RD) J(v int) DHR8RD {
	return DHR8RD(bits.MakeField32(v, uint32(mask)))
}

type RDHR8RD struct{ mmio.U32 }

func (r *RDHR8RD) Bits(mask DHR8RD) DHR8RD  { return DHR8RD(r.U32.Bits(uint32(mask))) }
func (r *RDHR8RD) StoreBits(mask, b DHR8RD) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDHR8RD) SetBits(mask DHR8RD)      { r.U32.SetBits(uint32(mask)) }
func (r *RDHR8RD) ClearBits(mask DHR8RD)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDHR8RD) Load() DHR8RD             { return DHR8RD(r.U32.Load()) }
func (r *RDHR8RD) Store(b DHR8RD)           { r.U32.Store(uint32(b)) }

func (r *RDHR8RD) AtomicStoreBits(mask, b DHR8RD) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDHR8RD) AtomicSetBits(mask DHR8RD)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDHR8RD) AtomicClearBits(mask DHR8RD)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDHR8RD struct{ mmio.UM32 }

func (rm RMDHR8RD) Load() DHR8RD   { return DHR8RD(rm.UM32.Load()) }
func (rm RMDHR8RD) Store(b DHR8RD) { rm.UM32.Store(uint32(b)) }

func (p *DAC_Periph) DACC1DHR() RMDHR8RD {
	return RMDHR8RD{mmio.UM32{&p.DHR8RD.U32, uint32(DACC1DHR)}}
}

func (p *DAC_Periph) DACC2DHR() RMDHR8RD {
	return RMDHR8RD{mmio.UM32{&p.DHR8RD.U32, uint32(DACC2DHR)}}
}

type DOR1 uint32

func (b DOR1) Field(mask DOR1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DOR1) J(v int) DOR1 {
	return DOR1(bits.MakeField32(v, uint32(mask)))
}

type RDOR1 struct{ mmio.U32 }

func (r *RDOR1) Bits(mask DOR1) DOR1    { return DOR1(r.U32.Bits(uint32(mask))) }
func (r *RDOR1) StoreBits(mask, b DOR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDOR1) SetBits(mask DOR1)      { r.U32.SetBits(uint32(mask)) }
func (r *RDOR1) ClearBits(mask DOR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDOR1) Load() DOR1             { return DOR1(r.U32.Load()) }
func (r *RDOR1) Store(b DOR1)           { r.U32.Store(uint32(b)) }

func (r *RDOR1) AtomicStoreBits(mask, b DOR1) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDOR1) AtomicSetBits(mask DOR1)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDOR1) AtomicClearBits(mask DOR1)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDOR1 struct{ mmio.UM32 }

func (rm RMDOR1) Load() DOR1   { return DOR1(rm.UM32.Load()) }
func (rm RMDOR1) Store(b DOR1) { rm.UM32.Store(uint32(b)) }

func (p *DAC_Periph) DACC1DOR() RMDOR1 {
	return RMDOR1{mmio.UM32{&p.DOR1.U32, uint32(DACC1DOR)}}
}

type DOR2 uint32

func (b DOR2) Field(mask DOR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DOR2) J(v int) DOR2 {
	return DOR2(bits.MakeField32(v, uint32(mask)))
}

type RDOR2 struct{ mmio.U32 }

func (r *RDOR2) Bits(mask DOR2) DOR2    { return DOR2(r.U32.Bits(uint32(mask))) }
func (r *RDOR2) StoreBits(mask, b DOR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDOR2) SetBits(mask DOR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RDOR2) ClearBits(mask DOR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDOR2) Load() DOR2             { return DOR2(r.U32.Load()) }
func (r *RDOR2) Store(b DOR2)           { r.U32.Store(uint32(b)) }

func (r *RDOR2) AtomicStoreBits(mask, b DOR2) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDOR2) AtomicSetBits(mask DOR2)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDOR2) AtomicClearBits(mask DOR2)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDOR2 struct{ mmio.UM32 }

func (rm RMDOR2) Load() DOR2   { return DOR2(rm.UM32.Load()) }
func (rm RMDOR2) Store(b DOR2) { rm.UM32.Store(uint32(b)) }

func (p *DAC_Periph) DACC2DOR() RMDOR2 {
	return RMDOR2{mmio.UM32{&p.DOR2.U32, uint32(DACC2DOR)}}
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

func (p *DAC_Periph) DMAUDR1() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(DMAUDR1)}}
}

func (p *DAC_Periph) DMAUDR2() RMSR {
	return RMSR{mmio.UM32{&p.SR.U32, uint32(DMAUDR2)}}
}
