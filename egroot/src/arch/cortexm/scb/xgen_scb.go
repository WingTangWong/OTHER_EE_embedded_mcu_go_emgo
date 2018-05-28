package scb

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"
)

type SCB_Periph struct {
	CPUID RCPUID
	ICSR  RICSR
	VTOR  RVTOR
	AIRCR RAIRCR
	SCR   RSCR
	CCR   RCCR
	SHPR1 RSHPR1
	SHPR2 RSHPR2
	SHPR3 RSHPR3
	SHCSR RSHCSR
	CFSR  RCFSR
	HFSR  RHFSR
	_     uint32
	MMFR  RMMFR
	BFAR  RBFAR
	AFSR  RAFSR
}

func (p *SCB_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var SCB = (*SCB_Periph)(unsafe.Pointer(uintptr(0xE000ED00)))

type CPUID uint32

func (b CPUID) Field(mask CPUID) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CPUID) J(v int) CPUID {
	return CPUID(bits.MakeField32(v, uint32(mask)))
}

type RCPUID struct{ mmio.U32 }

func (r *RCPUID) Bits(mask CPUID) CPUID   { return CPUID(r.U32.Bits(uint32(mask))) }
func (r *RCPUID) StoreBits(mask, b CPUID) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCPUID) SetBits(mask CPUID)      { r.U32.SetBits(uint32(mask)) }
func (r *RCPUID) ClearBits(mask CPUID)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCPUID) Load() CPUID             { return CPUID(r.U32.Load()) }
func (r *RCPUID) Store(b CPUID)           { r.U32.Store(uint32(b)) }

func (r *RCPUID) AtomicStoreBits(mask, b CPUID) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCPUID) AtomicSetBits(mask CPUID)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCPUID) AtomicClearBits(mask CPUID)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCPUID struct{ mmio.UM32 }

func (rm RMCPUID) Load() CPUID   { return CPUID(rm.UM32.Load()) }
func (rm RMCPUID) Store(b CPUID) { rm.UM32.Store(uint32(b)) }

func (p *SCB_Periph) Revision() RMCPUID {
	return RMCPUID{mmio.UM32{&p.CPUID.U32, uint32(Revision)}}
}

func (p *SCB_Periph) PartNo() RMCPUID {
	return RMCPUID{mmio.UM32{&p.CPUID.U32, uint32(PartNo)}}
}

func (p *SCB_Periph) Constant() RMCPUID {
	return RMCPUID{mmio.UM32{&p.CPUID.U32, uint32(Constant)}}
}

func (p *SCB_Periph) Variant() RMCPUID {
	return RMCPUID{mmio.UM32{&p.CPUID.U32, uint32(Variant)}}
}

func (p *SCB_Periph) Implementer() RMCPUID {
	return RMCPUID{mmio.UM32{&p.CPUID.U32, uint32(Implementer)}}
}

type ICSR uint32

func (b ICSR) Field(mask ICSR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ICSR) J(v int) ICSR {
	return ICSR(bits.MakeField32(v, uint32(mask)))
}

type RICSR struct{ mmio.U32 }

func (r *RICSR) Bits(mask ICSR) ICSR    { return ICSR(r.U32.Bits(uint32(mask))) }
func (r *RICSR) StoreBits(mask, b ICSR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RICSR) SetBits(mask ICSR)      { r.U32.SetBits(uint32(mask)) }
func (r *RICSR) ClearBits(mask ICSR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RICSR) Load() ICSR             { return ICSR(r.U32.Load()) }
func (r *RICSR) Store(b ICSR)           { r.U32.Store(uint32(b)) }

func (r *RICSR) AtomicStoreBits(mask, b ICSR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RICSR) AtomicSetBits(mask ICSR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RICSR) AtomicClearBits(mask ICSR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMICSR struct{ mmio.UM32 }

func (rm RMICSR) Load() ICSR   { return ICSR(rm.UM32.Load()) }
func (rm RMICSR) Store(b ICSR) { rm.UM32.Store(uint32(b)) }

func (p *SCB_Periph) VECTACTIVE() RMICSR {
	return RMICSR{mmio.UM32{&p.ICSR.U32, uint32(VECTACTIVE)}}
}

func (p *SCB_Periph) RETTOBASE() RMICSR {
	return RMICSR{mmio.UM32{&p.ICSR.U32, uint32(RETTOBASE)}}
}

func (p *SCB_Periph) VECTPENDING() RMICSR {
	return RMICSR{mmio.UM32{&p.ICSR.U32, uint32(VECTPENDING)}}
}

func (p *SCB_Periph) ISRPENDING() RMICSR {
	return RMICSR{mmio.UM32{&p.ICSR.U32, uint32(ISRPENDING)}}
}

func (p *SCB_Periph) PENDSTCLR() RMICSR {
	return RMICSR{mmio.UM32{&p.ICSR.U32, uint32(PENDSTCLR)}}
}

func (p *SCB_Periph) PENDSTSET() RMICSR {
	return RMICSR{mmio.UM32{&p.ICSR.U32, uint32(PENDSTSET)}}
}

func (p *SCB_Periph) PENDSVCLR() RMICSR {
	return RMICSR{mmio.UM32{&p.ICSR.U32, uint32(PENDSVCLR)}}
}

func (p *SCB_Periph) PENDSVSET() RMICSR {
	return RMICSR{mmio.UM32{&p.ICSR.U32, uint32(PENDSVSET)}}
}

func (p *SCB_Periph) NMIPENDSET() RMICSR {
	return RMICSR{mmio.UM32{&p.ICSR.U32, uint32(NMIPENDSET)}}
}

type VTOR uint32

func (b VTOR) Field(mask VTOR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask VTOR) J(v int) VTOR {
	return VTOR(bits.MakeField32(v, uint32(mask)))
}

type RVTOR struct{ mmio.U32 }

func (r *RVTOR) Bits(mask VTOR) VTOR    { return VTOR(r.U32.Bits(uint32(mask))) }
func (r *RVTOR) StoreBits(mask, b VTOR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RVTOR) SetBits(mask VTOR)      { r.U32.SetBits(uint32(mask)) }
func (r *RVTOR) ClearBits(mask VTOR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RVTOR) Load() VTOR             { return VTOR(r.U32.Load()) }
func (r *RVTOR) Store(b VTOR)           { r.U32.Store(uint32(b)) }

func (r *RVTOR) AtomicStoreBits(mask, b VTOR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RVTOR) AtomicSetBits(mask VTOR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RVTOR) AtomicClearBits(mask VTOR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMVTOR struct{ mmio.UM32 }

func (rm RMVTOR) Load() VTOR   { return VTOR(rm.UM32.Load()) }
func (rm RMVTOR) Store(b VTOR) { rm.UM32.Store(uint32(b)) }

func (p *SCB_Periph) TBLOFF() RMVTOR {
	return RMVTOR{mmio.UM32{&p.VTOR.U32, uint32(TBLOFF)}}
}

type AIRCR uint32

func (b AIRCR) Field(mask AIRCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AIRCR) J(v int) AIRCR {
	return AIRCR(bits.MakeField32(v, uint32(mask)))
}

type RAIRCR struct{ mmio.U32 }

func (r *RAIRCR) Bits(mask AIRCR) AIRCR   { return AIRCR(r.U32.Bits(uint32(mask))) }
func (r *RAIRCR) StoreBits(mask, b AIRCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAIRCR) SetBits(mask AIRCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RAIRCR) ClearBits(mask AIRCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RAIRCR) Load() AIRCR             { return AIRCR(r.U32.Load()) }
func (r *RAIRCR) Store(b AIRCR)           { r.U32.Store(uint32(b)) }

func (r *RAIRCR) AtomicStoreBits(mask, b AIRCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RAIRCR) AtomicSetBits(mask AIRCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAIRCR) AtomicClearBits(mask AIRCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMAIRCR struct{ mmio.UM32 }

func (rm RMAIRCR) Load() AIRCR   { return AIRCR(rm.UM32.Load()) }
func (rm RMAIRCR) Store(b AIRCR) { rm.UM32.Store(uint32(b)) }

func (p *SCB_Periph) VECTRESET() RMAIRCR {
	return RMAIRCR{mmio.UM32{&p.AIRCR.U32, uint32(VECTRESET)}}
}

func (p *SCB_Periph) VECTCLRACTIVE() RMAIRCR {
	return RMAIRCR{mmio.UM32{&p.AIRCR.U32, uint32(VECTCLRACTIVE)}}
}

func (p *SCB_Periph) SYSRESETREQ() RMAIRCR {
	return RMAIRCR{mmio.UM32{&p.AIRCR.U32, uint32(SYSRESETREQ)}}
}

func (p *SCB_Periph) PRIGROUP() RMAIRCR {
	return RMAIRCR{mmio.UM32{&p.AIRCR.U32, uint32(PRIGROUP)}}
}

func (p *SCB_Periph) ENDIANNESS() RMAIRCR {
	return RMAIRCR{mmio.UM32{&p.AIRCR.U32, uint32(ENDIANNESS)}}
}

func (p *SCB_Periph) VECTKEY() RMAIRCR {
	return RMAIRCR{mmio.UM32{&p.AIRCR.U32, uint32(VECTKEY)}}
}

type SCR uint32

func (b SCR) Field(mask SCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SCR) J(v int) SCR {
	return SCR(bits.MakeField32(v, uint32(mask)))
}

type RSCR struct{ mmio.U32 }

func (r *RSCR) Bits(mask SCR) SCR     { return SCR(r.U32.Bits(uint32(mask))) }
func (r *RSCR) StoreBits(mask, b SCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSCR) SetBits(mask SCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSCR) ClearBits(mask SCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSCR) Load() SCR             { return SCR(r.U32.Load()) }
func (r *RSCR) Store(b SCR)           { r.U32.Store(uint32(b)) }

func (r *RSCR) AtomicStoreBits(mask, b SCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSCR) AtomicSetBits(mask SCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSCR) AtomicClearBits(mask SCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSCR struct{ mmio.UM32 }

func (rm RMSCR) Load() SCR   { return SCR(rm.UM32.Load()) }
func (rm RMSCR) Store(b SCR) { rm.UM32.Store(uint32(b)) }

func (p *SCB_Periph) SLEEPONEXIT() RMSCR {
	return RMSCR{mmio.UM32{&p.SCR.U32, uint32(SLEEPONEXIT)}}
}

func (p *SCB_Periph) SLEEPDEEP() RMSCR {
	return RMSCR{mmio.UM32{&p.SCR.U32, uint32(SLEEPDEEP)}}
}

func (p *SCB_Periph) SEVONPEND() RMSCR {
	return RMSCR{mmio.UM32{&p.SCR.U32, uint32(SEVONPEND)}}
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

func (p *SCB_Periph) NONBASETHRDENA() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(NONBASETHRDENA)}}
}

func (p *SCB_Periph) USERSETMPEND() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(USERSETMPEND)}}
}

func (p *SCB_Periph) UNALIGN_TRP() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(UNALIGN_TRP)}}
}

func (p *SCB_Periph) DIV_0_TRP() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(DIV_0_TRP)}}
}

func (p *SCB_Periph) BFHFNMIGN() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(BFHFNMIGN)}}
}

func (p *SCB_Periph) STKALIGN() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(STKALIGN)}}
}

func (p *SCB_Periph) DC() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(DC)}}
}

func (p *SCB_Periph) IC() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(IC)}}
}

func (p *SCB_Periph) BP() RMCCR {
	return RMCCR{mmio.UM32{&p.CCR.U32, uint32(BP)}}
}

type SHPR1 uint32

func (b SHPR1) Field(mask SHPR1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SHPR1) J(v int) SHPR1 {
	return SHPR1(bits.MakeField32(v, uint32(mask)))
}

type RSHPR1 struct{ mmio.U32 }

func (r *RSHPR1) Bits(mask SHPR1) SHPR1   { return SHPR1(r.U32.Bits(uint32(mask))) }
func (r *RSHPR1) StoreBits(mask, b SHPR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSHPR1) SetBits(mask SHPR1)      { r.U32.SetBits(uint32(mask)) }
func (r *RSHPR1) ClearBits(mask SHPR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSHPR1) Load() SHPR1             { return SHPR1(r.U32.Load()) }
func (r *RSHPR1) Store(b SHPR1)           { r.U32.Store(uint32(b)) }

func (r *RSHPR1) AtomicStoreBits(mask, b SHPR1) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSHPR1) AtomicSetBits(mask SHPR1)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSHPR1) AtomicClearBits(mask SHPR1)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSHPR1 struct{ mmio.UM32 }

func (rm RMSHPR1) Load() SHPR1   { return SHPR1(rm.UM32.Load()) }
func (rm RMSHPR1) Store(b SHPR1) { rm.UM32.Store(uint32(b)) }

func (p *SCB_Periph) PRI_MemManage() RMSHPR1 {
	return RMSHPR1{mmio.UM32{&p.SHPR1.U32, uint32(PRI_MemManage)}}
}

func (p *SCB_Periph) PRI_BusFault() RMSHPR1 {
	return RMSHPR1{mmio.UM32{&p.SHPR1.U32, uint32(PRI_BusFault)}}
}

func (p *SCB_Periph) PRI_UsageFault() RMSHPR1 {
	return RMSHPR1{mmio.UM32{&p.SHPR1.U32, uint32(PRI_UsageFault)}}
}

type SHPR2 uint32

func (b SHPR2) Field(mask SHPR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SHPR2) J(v int) SHPR2 {
	return SHPR2(bits.MakeField32(v, uint32(mask)))
}

type RSHPR2 struct{ mmio.U32 }

func (r *RSHPR2) Bits(mask SHPR2) SHPR2   { return SHPR2(r.U32.Bits(uint32(mask))) }
func (r *RSHPR2) StoreBits(mask, b SHPR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSHPR2) SetBits(mask SHPR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RSHPR2) ClearBits(mask SHPR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSHPR2) Load() SHPR2             { return SHPR2(r.U32.Load()) }
func (r *RSHPR2) Store(b SHPR2)           { r.U32.Store(uint32(b)) }

func (r *RSHPR2) AtomicStoreBits(mask, b SHPR2) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSHPR2) AtomicSetBits(mask SHPR2)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSHPR2) AtomicClearBits(mask SHPR2)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSHPR2 struct{ mmio.UM32 }

func (rm RMSHPR2) Load() SHPR2   { return SHPR2(rm.UM32.Load()) }
func (rm RMSHPR2) Store(b SHPR2) { rm.UM32.Store(uint32(b)) }

func (p *SCB_Periph) PRI_SVCall() RMSHPR2 {
	return RMSHPR2{mmio.UM32{&p.SHPR2.U32, uint32(PRI_SVCall)}}
}

type SHPR3 uint32

func (b SHPR3) Field(mask SHPR3) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SHPR3) J(v int) SHPR3 {
	return SHPR3(bits.MakeField32(v, uint32(mask)))
}

type RSHPR3 struct{ mmio.U32 }

func (r *RSHPR3) Bits(mask SHPR3) SHPR3   { return SHPR3(r.U32.Bits(uint32(mask))) }
func (r *RSHPR3) StoreBits(mask, b SHPR3) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSHPR3) SetBits(mask SHPR3)      { r.U32.SetBits(uint32(mask)) }
func (r *RSHPR3) ClearBits(mask SHPR3)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSHPR3) Load() SHPR3             { return SHPR3(r.U32.Load()) }
func (r *RSHPR3) Store(b SHPR3)           { r.U32.Store(uint32(b)) }

func (r *RSHPR3) AtomicStoreBits(mask, b SHPR3) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSHPR3) AtomicSetBits(mask SHPR3)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSHPR3) AtomicClearBits(mask SHPR3)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSHPR3 struct{ mmio.UM32 }

func (rm RMSHPR3) Load() SHPR3   { return SHPR3(rm.UM32.Load()) }
func (rm RMSHPR3) Store(b SHPR3) { rm.UM32.Store(uint32(b)) }

func (p *SCB_Periph) PRI_PendSV() RMSHPR3 {
	return RMSHPR3{mmio.UM32{&p.SHPR3.U32, uint32(PRI_PendSV)}}
}

func (p *SCB_Periph) PRI_SysTick() RMSHPR3 {
	return RMSHPR3{mmio.UM32{&p.SHPR3.U32, uint32(PRI_SysTick)}}
}

type SHCSR uint32

func (b SHCSR) Field(mask SHCSR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SHCSR) J(v int) SHCSR {
	return SHCSR(bits.MakeField32(v, uint32(mask)))
}

type RSHCSR struct{ mmio.U32 }

func (r *RSHCSR) Bits(mask SHCSR) SHCSR   { return SHCSR(r.U32.Bits(uint32(mask))) }
func (r *RSHCSR) StoreBits(mask, b SHCSR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSHCSR) SetBits(mask SHCSR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSHCSR) ClearBits(mask SHCSR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSHCSR) Load() SHCSR             { return SHCSR(r.U32.Load()) }
func (r *RSHCSR) Store(b SHCSR)           { r.U32.Store(uint32(b)) }

func (r *RSHCSR) AtomicStoreBits(mask, b SHCSR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSHCSR) AtomicSetBits(mask SHCSR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSHCSR) AtomicClearBits(mask SHCSR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSHCSR struct{ mmio.UM32 }

func (rm RMSHCSR) Load() SHCSR   { return SHCSR(rm.UM32.Load()) }
func (rm RMSHCSR) Store(b SHCSR) { rm.UM32.Store(uint32(b)) }

func (p *SCB_Periph) MEMFAULTACT() RMSHCSR {
	return RMSHCSR{mmio.UM32{&p.SHCSR.U32, uint32(MEMFAULTACT)}}
}

func (p *SCB_Periph) BUSFAULTACT() RMSHCSR {
	return RMSHCSR{mmio.UM32{&p.SHCSR.U32, uint32(BUSFAULTACT)}}
}

func (p *SCB_Periph) USGFAULTACT() RMSHCSR {
	return RMSHCSR{mmio.UM32{&p.SHCSR.U32, uint32(USGFAULTACT)}}
}

func (p *SCB_Periph) SVCALLACT() RMSHCSR {
	return RMSHCSR{mmio.UM32{&p.SHCSR.U32, uint32(SVCALLACT)}}
}

func (p *SCB_Periph) MONITORACT() RMSHCSR {
	return RMSHCSR{mmio.UM32{&p.SHCSR.U32, uint32(MONITORACT)}}
}

func (p *SCB_Periph) PENDSVACT() RMSHCSR {
	return RMSHCSR{mmio.UM32{&p.SHCSR.U32, uint32(PENDSVACT)}}
}

func (p *SCB_Periph) SYSTICKACT() RMSHCSR {
	return RMSHCSR{mmio.UM32{&p.SHCSR.U32, uint32(SYSTICKACT)}}
}

func (p *SCB_Periph) USGFAULTPENDED() RMSHCSR {
	return RMSHCSR{mmio.UM32{&p.SHCSR.U32, uint32(USGFAULTPENDED)}}
}

func (p *SCB_Periph) MEMFAULTPENDED() RMSHCSR {
	return RMSHCSR{mmio.UM32{&p.SHCSR.U32, uint32(MEMFAULTPENDED)}}
}

func (p *SCB_Periph) BUSFAULTPENDED() RMSHCSR {
	return RMSHCSR{mmio.UM32{&p.SHCSR.U32, uint32(BUSFAULTPENDED)}}
}

func (p *SCB_Periph) SVCALLPENDED() RMSHCSR {
	return RMSHCSR{mmio.UM32{&p.SHCSR.U32, uint32(SVCALLPENDED)}}
}

func (p *SCB_Periph) MEMFAULTENA() RMSHCSR {
	return RMSHCSR{mmio.UM32{&p.SHCSR.U32, uint32(MEMFAULTENA)}}
}

func (p *SCB_Periph) BUSFAULTENA() RMSHCSR {
	return RMSHCSR{mmio.UM32{&p.SHCSR.U32, uint32(BUSFAULTENA)}}
}

func (p *SCB_Periph) USGFAULTENA() RMSHCSR {
	return RMSHCSR{mmio.UM32{&p.SHCSR.U32, uint32(USGFAULTENA)}}
}

type CFSR uint32

func (b CFSR) Field(mask CFSR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CFSR) J(v int) CFSR {
	return CFSR(bits.MakeField32(v, uint32(mask)))
}

type RCFSR struct{ mmio.U32 }

func (r *RCFSR) Bits(mask CFSR) CFSR    { return CFSR(r.U32.Bits(uint32(mask))) }
func (r *RCFSR) StoreBits(mask, b CFSR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCFSR) SetBits(mask CFSR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCFSR) ClearBits(mask CFSR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCFSR) Load() CFSR             { return CFSR(r.U32.Load()) }
func (r *RCFSR) Store(b CFSR)           { r.U32.Store(uint32(b)) }

func (r *RCFSR) AtomicStoreBits(mask, b CFSR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCFSR) AtomicSetBits(mask CFSR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCFSR) AtomicClearBits(mask CFSR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCFSR struct{ mmio.UM32 }

func (rm RMCFSR) Load() CFSR   { return CFSR(rm.UM32.Load()) }
func (rm RMCFSR) Store(b CFSR) { rm.UM32.Store(uint32(b)) }

func (p *SCB_Periph) IACCVIOL() RMCFSR {
	return RMCFSR{mmio.UM32{&p.CFSR.U32, uint32(IACCVIOL)}}
}

func (p *SCB_Periph) DACCVIOL() RMCFSR {
	return RMCFSR{mmio.UM32{&p.CFSR.U32, uint32(DACCVIOL)}}
}

func (p *SCB_Periph) MUNSTKERR() RMCFSR {
	return RMCFSR{mmio.UM32{&p.CFSR.U32, uint32(MUNSTKERR)}}
}

func (p *SCB_Periph) MSTKERR() RMCFSR {
	return RMCFSR{mmio.UM32{&p.CFSR.U32, uint32(MSTKERR)}}
}

func (p *SCB_Periph) MLSPERR() RMCFSR {
	return RMCFSR{mmio.UM32{&p.CFSR.U32, uint32(MLSPERR)}}
}

func (p *SCB_Periph) MMARVALID() RMCFSR {
	return RMCFSR{mmio.UM32{&p.CFSR.U32, uint32(MMARVALID)}}
}

func (p *SCB_Periph) IBUSERR() RMCFSR {
	return RMCFSR{mmio.UM32{&p.CFSR.U32, uint32(IBUSERR)}}
}

func (p *SCB_Periph) PRECISERR() RMCFSR {
	return RMCFSR{mmio.UM32{&p.CFSR.U32, uint32(PRECISERR)}}
}

func (p *SCB_Periph) IMPRECISERR() RMCFSR {
	return RMCFSR{mmio.UM32{&p.CFSR.U32, uint32(IMPRECISERR)}}
}

func (p *SCB_Periph) UNSTKERR() RMCFSR {
	return RMCFSR{mmio.UM32{&p.CFSR.U32, uint32(UNSTKERR)}}
}

func (p *SCB_Periph) STKERR() RMCFSR {
	return RMCFSR{mmio.UM32{&p.CFSR.U32, uint32(STKERR)}}
}

func (p *SCB_Periph) LSPERR() RMCFSR {
	return RMCFSR{mmio.UM32{&p.CFSR.U32, uint32(LSPERR)}}
}

func (p *SCB_Periph) BFARVALID() RMCFSR {
	return RMCFSR{mmio.UM32{&p.CFSR.U32, uint32(BFARVALID)}}
}

func (p *SCB_Periph) UNDEFINSTR() RMCFSR {
	return RMCFSR{mmio.UM32{&p.CFSR.U32, uint32(UNDEFINSTR)}}
}

func (p *SCB_Periph) INVSTATE() RMCFSR {
	return RMCFSR{mmio.UM32{&p.CFSR.U32, uint32(INVSTATE)}}
}

func (p *SCB_Periph) INVPC() RMCFSR {
	return RMCFSR{mmio.UM32{&p.CFSR.U32, uint32(INVPC)}}
}

func (p *SCB_Periph) NOCP() RMCFSR {
	return RMCFSR{mmio.UM32{&p.CFSR.U32, uint32(NOCP)}}
}

func (p *SCB_Periph) UNALIGNED() RMCFSR {
	return RMCFSR{mmio.UM32{&p.CFSR.U32, uint32(UNALIGNED)}}
}

func (p *SCB_Periph) DIVBYZERO() RMCFSR {
	return RMCFSR{mmio.UM32{&p.CFSR.U32, uint32(DIVBYZERO)}}
}

type HFSR uint32

func (b HFSR) Field(mask HFSR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HFSR) J(v int) HFSR {
	return HFSR(bits.MakeField32(v, uint32(mask)))
}

type RHFSR struct{ mmio.U32 }

func (r *RHFSR) Bits(mask HFSR) HFSR    { return HFSR(r.U32.Bits(uint32(mask))) }
func (r *RHFSR) StoreBits(mask, b HFSR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RHFSR) SetBits(mask HFSR)      { r.U32.SetBits(uint32(mask)) }
func (r *RHFSR) ClearBits(mask HFSR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RHFSR) Load() HFSR             { return HFSR(r.U32.Load()) }
func (r *RHFSR) Store(b HFSR)           { r.U32.Store(uint32(b)) }

func (r *RHFSR) AtomicStoreBits(mask, b HFSR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RHFSR) AtomicSetBits(mask HFSR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RHFSR) AtomicClearBits(mask HFSR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMHFSR struct{ mmio.UM32 }

func (rm RMHFSR) Load() HFSR   { return HFSR(rm.UM32.Load()) }
func (rm RMHFSR) Store(b HFSR) { rm.UM32.Store(uint32(b)) }

func (p *SCB_Periph) VECTTBL() RMHFSR {
	return RMHFSR{mmio.UM32{&p.HFSR.U32, uint32(VECTTBL)}}
}

func (p *SCB_Periph) FORCED() RMHFSR {
	return RMHFSR{mmio.UM32{&p.HFSR.U32, uint32(FORCED)}}
}

func (p *SCB_Periph) DEBUGEVT() RMHFSR {
	return RMHFSR{mmio.UM32{&p.HFSR.U32, uint32(DEBUGEVT)}}
}

type MMFR uint32

func (b MMFR) Field(mask MMFR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask MMFR) J(v int) MMFR {
	return MMFR(bits.MakeField32(v, uint32(mask)))
}

type RMMFR struct{ mmio.U32 }

func (r *RMMFR) Bits(mask MMFR) MMFR    { return MMFR(r.U32.Bits(uint32(mask))) }
func (r *RMMFR) StoreBits(mask, b MMFR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RMMFR) SetBits(mask MMFR)      { r.U32.SetBits(uint32(mask)) }
func (r *RMMFR) ClearBits(mask MMFR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RMMFR) Load() MMFR             { return MMFR(r.U32.Load()) }
func (r *RMMFR) Store(b MMFR)           { r.U32.Store(uint32(b)) }

func (r *RMMFR) AtomicStoreBits(mask, b MMFR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RMMFR) AtomicSetBits(mask MMFR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RMMFR) AtomicClearBits(mask MMFR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMMMFR struct{ mmio.UM32 }

func (rm RMMMFR) Load() MMFR   { return MMFR(rm.UM32.Load()) }
func (rm RMMMFR) Store(b MMFR) { rm.UM32.Store(uint32(b)) }

type BFAR uint32

func (b BFAR) Field(mask BFAR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BFAR) J(v int) BFAR {
	return BFAR(bits.MakeField32(v, uint32(mask)))
}

type RBFAR struct{ mmio.U32 }

func (r *RBFAR) Bits(mask BFAR) BFAR    { return BFAR(r.U32.Bits(uint32(mask))) }
func (r *RBFAR) StoreBits(mask, b BFAR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RBFAR) SetBits(mask BFAR)      { r.U32.SetBits(uint32(mask)) }
func (r *RBFAR) ClearBits(mask BFAR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RBFAR) Load() BFAR             { return BFAR(r.U32.Load()) }
func (r *RBFAR) Store(b BFAR)           { r.U32.Store(uint32(b)) }

func (r *RBFAR) AtomicStoreBits(mask, b BFAR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RBFAR) AtomicSetBits(mask BFAR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RBFAR) AtomicClearBits(mask BFAR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMBFAR struct{ mmio.UM32 }

func (rm RMBFAR) Load() BFAR   { return BFAR(rm.UM32.Load()) }
func (rm RMBFAR) Store(b BFAR) { rm.UM32.Store(uint32(b)) }

type AFSR uint32

func (b AFSR) Field(mask AFSR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AFSR) J(v int) AFSR {
	return AFSR(bits.MakeField32(v, uint32(mask)))
}

type RAFSR struct{ mmio.U32 }

func (r *RAFSR) Bits(mask AFSR) AFSR    { return AFSR(r.U32.Bits(uint32(mask))) }
func (r *RAFSR) StoreBits(mask, b AFSR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAFSR) SetBits(mask AFSR)      { r.U32.SetBits(uint32(mask)) }
func (r *RAFSR) ClearBits(mask AFSR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RAFSR) Load() AFSR             { return AFSR(r.U32.Load()) }
func (r *RAFSR) Store(b AFSR)           { r.U32.Store(uint32(b)) }

func (r *RAFSR) AtomicStoreBits(mask, b AFSR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RAFSR) AtomicSetBits(mask AFSR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAFSR) AtomicClearBits(mask AFSR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMAFSR struct{ mmio.UM32 }

func (rm RMAFSR) Load() AFSR   { return AFSR(rm.UM32.Load()) }
func (rm RMAFSR) Store(b AFSR) { rm.UM32.Store(uint32(b)) }
