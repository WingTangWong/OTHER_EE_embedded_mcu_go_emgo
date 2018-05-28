// +build l476xx

package syscfg

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type SYSCFG_Periph struct {
	MEMRMP RMEMRMP
	CFGR1  RCFGR1
	EXTICR [4]REXTICR
	SCSR   RSCSR
	CFGR2  RCFGR2
	SWPR   RSWPR
	SKR    RSKR
}

func (p *SYSCFG_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var SYSCFG = (*SYSCFG_Periph)(unsafe.Pointer(uintptr(mmap.SYSCFG_BASE)))

type MEMRMP uint32

func (b MEMRMP) Field(mask MEMRMP) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask MEMRMP) J(v int) MEMRMP {
	return MEMRMP(bits.MakeField32(v, uint32(mask)))
}

type RMEMRMP struct{ mmio.U32 }

func (r *RMEMRMP) Bits(mask MEMRMP) MEMRMP  { return MEMRMP(r.U32.Bits(uint32(mask))) }
func (r *RMEMRMP) StoreBits(mask, b MEMRMP) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RMEMRMP) SetBits(mask MEMRMP)      { r.U32.SetBits(uint32(mask)) }
func (r *RMEMRMP) ClearBits(mask MEMRMP)    { r.U32.ClearBits(uint32(mask)) }
func (r *RMEMRMP) Load() MEMRMP             { return MEMRMP(r.U32.Load()) }
func (r *RMEMRMP) Store(b MEMRMP)           { r.U32.Store(uint32(b)) }

func (r *RMEMRMP) AtomicStoreBits(mask, b MEMRMP) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RMEMRMP) AtomicSetBits(mask MEMRMP)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RMEMRMP) AtomicClearBits(mask MEMRMP)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMMEMRMP struct{ mmio.UM32 }

func (rm RMMEMRMP) Load() MEMRMP   { return MEMRMP(rm.UM32.Load()) }
func (rm RMMEMRMP) Store(b MEMRMP) { rm.UM32.Store(uint32(b)) }

func (p *SYSCFG_Periph) MEM_MODE() RMMEMRMP {
	return RMMEMRMP{mmio.UM32{&p.MEMRMP.U32, uint32(MEM_MODE)}}
}

func (p *SYSCFG_Periph) FB_MODE() RMMEMRMP {
	return RMMEMRMP{mmio.UM32{&p.MEMRMP.U32, uint32(FB_MODE)}}
}

type CFGR1 uint32

func (b CFGR1) Field(mask CFGR1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CFGR1) J(v int) CFGR1 {
	return CFGR1(bits.MakeField32(v, uint32(mask)))
}

type RCFGR1 struct{ mmio.U32 }

func (r *RCFGR1) Bits(mask CFGR1) CFGR1   { return CFGR1(r.U32.Bits(uint32(mask))) }
func (r *RCFGR1) StoreBits(mask, b CFGR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCFGR1) SetBits(mask CFGR1)      { r.U32.SetBits(uint32(mask)) }
func (r *RCFGR1) ClearBits(mask CFGR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCFGR1) Load() CFGR1             { return CFGR1(r.U32.Load()) }
func (r *RCFGR1) Store(b CFGR1)           { r.U32.Store(uint32(b)) }

func (r *RCFGR1) AtomicStoreBits(mask, b CFGR1) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCFGR1) AtomicSetBits(mask CFGR1)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCFGR1) AtomicClearBits(mask CFGR1)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCFGR1 struct{ mmio.UM32 }

func (rm RMCFGR1) Load() CFGR1   { return CFGR1(rm.UM32.Load()) }
func (rm RMCFGR1) Store(b CFGR1) { rm.UM32.Store(uint32(b)) }

func (p *SYSCFG_Periph) FWDIS() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(FWDIS)}}
}

func (p *SYSCFG_Periph) BOOSTEN() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(BOOSTEN)}}
}

func (p *SYSCFG_Periph) I2C_PB6_FMP() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(I2C_PB6_FMP)}}
}

func (p *SYSCFG_Periph) I2C_PB7_FMP() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(I2C_PB7_FMP)}}
}

func (p *SYSCFG_Periph) I2C_PB8_FMP() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(I2C_PB8_FMP)}}
}

func (p *SYSCFG_Periph) I2C_PB9_FMP() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(I2C_PB9_FMP)}}
}

func (p *SYSCFG_Periph) I2C1_FMP() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(I2C1_FMP)}}
}

func (p *SYSCFG_Periph) I2C2_FMP() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(I2C2_FMP)}}
}

func (p *SYSCFG_Periph) I2C3_FMP() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(I2C3_FMP)}}
}

func (p *SYSCFG_Periph) FPU_IE_0() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(FPU_IE_0)}}
}

func (p *SYSCFG_Periph) FPU_IE_1() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(FPU_IE_1)}}
}

func (p *SYSCFG_Periph) FPU_IE_2() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(FPU_IE_2)}}
}

func (p *SYSCFG_Periph) FPU_IE_3() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(FPU_IE_3)}}
}

func (p *SYSCFG_Periph) FPU_IE_4() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(FPU_IE_4)}}
}

func (p *SYSCFG_Periph) FPU_IE_5() RMCFGR1 {
	return RMCFGR1{mmio.UM32{&p.CFGR1.U32, uint32(FPU_IE_5)}}
}

type EXTICR uint32

func (b EXTICR) Field(mask EXTICR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask EXTICR) J(v int) EXTICR {
	return EXTICR(bits.MakeField32(v, uint32(mask)))
}

type REXTICR struct{ mmio.U32 }

func (r *REXTICR) Bits(mask EXTICR) EXTICR  { return EXTICR(r.U32.Bits(uint32(mask))) }
func (r *REXTICR) StoreBits(mask, b EXTICR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *REXTICR) SetBits(mask EXTICR)      { r.U32.SetBits(uint32(mask)) }
func (r *REXTICR) ClearBits(mask EXTICR)    { r.U32.ClearBits(uint32(mask)) }
func (r *REXTICR) Load() EXTICR             { return EXTICR(r.U32.Load()) }
func (r *REXTICR) Store(b EXTICR)           { r.U32.Store(uint32(b)) }

func (r *REXTICR) AtomicStoreBits(mask, b EXTICR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *REXTICR) AtomicSetBits(mask EXTICR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *REXTICR) AtomicClearBits(mask EXTICR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMEXTICR struct{ mmio.UM32 }

func (rm RMEXTICR) Load() EXTICR   { return EXTICR(rm.UM32.Load()) }
func (rm RMEXTICR) Store(b EXTICR) { rm.UM32.Store(uint32(b)) }

func (p *SYSCFG_Periph) EXTI0(n int) RMEXTICR {
	return RMEXTICR{mmio.UM32{&p.EXTICR[n].U32, uint32(EXTI0)}}
}

func (p *SYSCFG_Periph) EXTI1(n int) RMEXTICR {
	return RMEXTICR{mmio.UM32{&p.EXTICR[n].U32, uint32(EXTI1)}}
}

func (p *SYSCFG_Periph) EXTI2(n int) RMEXTICR {
	return RMEXTICR{mmio.UM32{&p.EXTICR[n].U32, uint32(EXTI2)}}
}

func (p *SYSCFG_Periph) EXTI3(n int) RMEXTICR {
	return RMEXTICR{mmio.UM32{&p.EXTICR[n].U32, uint32(EXTI3)}}
}

type SCSR uint32

func (b SCSR) Field(mask SCSR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SCSR) J(v int) SCSR {
	return SCSR(bits.MakeField32(v, uint32(mask)))
}

type RSCSR struct{ mmio.U32 }

func (r *RSCSR) Bits(mask SCSR) SCSR    { return SCSR(r.U32.Bits(uint32(mask))) }
func (r *RSCSR) StoreBits(mask, b SCSR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSCSR) SetBits(mask SCSR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSCSR) ClearBits(mask SCSR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSCSR) Load() SCSR             { return SCSR(r.U32.Load()) }
func (r *RSCSR) Store(b SCSR)           { r.U32.Store(uint32(b)) }

func (r *RSCSR) AtomicStoreBits(mask, b SCSR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSCSR) AtomicSetBits(mask SCSR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSCSR) AtomicClearBits(mask SCSR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSCSR struct{ mmio.UM32 }

func (rm RMSCSR) Load() SCSR   { return SCSR(rm.UM32.Load()) }
func (rm RMSCSR) Store(b SCSR) { rm.UM32.Store(uint32(b)) }

func (p *SYSCFG_Periph) SRAM2ER() RMSCSR {
	return RMSCSR{mmio.UM32{&p.SCSR.U32, uint32(SRAM2ER)}}
}

func (p *SYSCFG_Periph) SRAM2BSY() RMSCSR {
	return RMSCSR{mmio.UM32{&p.SCSR.U32, uint32(SRAM2BSY)}}
}

type CFGR2 uint32

func (b CFGR2) Field(mask CFGR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CFGR2) J(v int) CFGR2 {
	return CFGR2(bits.MakeField32(v, uint32(mask)))
}

type RCFGR2 struct{ mmio.U32 }

func (r *RCFGR2) Bits(mask CFGR2) CFGR2   { return CFGR2(r.U32.Bits(uint32(mask))) }
func (r *RCFGR2) StoreBits(mask, b CFGR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCFGR2) SetBits(mask CFGR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RCFGR2) ClearBits(mask CFGR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCFGR2) Load() CFGR2             { return CFGR2(r.U32.Load()) }
func (r *RCFGR2) Store(b CFGR2)           { r.U32.Store(uint32(b)) }

func (r *RCFGR2) AtomicStoreBits(mask, b CFGR2) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCFGR2) AtomicSetBits(mask CFGR2)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCFGR2) AtomicClearBits(mask CFGR2)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCFGR2 struct{ mmio.UM32 }

func (rm RMCFGR2) Load() CFGR2   { return CFGR2(rm.UM32.Load()) }
func (rm RMCFGR2) Store(b CFGR2) { rm.UM32.Store(uint32(b)) }

func (p *SYSCFG_Periph) CLL() RMCFGR2 {
	return RMCFGR2{mmio.UM32{&p.CFGR2.U32, uint32(CLL)}}
}

func (p *SYSCFG_Periph) SPL() RMCFGR2 {
	return RMCFGR2{mmio.UM32{&p.CFGR2.U32, uint32(SPL)}}
}

func (p *SYSCFG_Periph) PVDL() RMCFGR2 {
	return RMCFGR2{mmio.UM32{&p.CFGR2.U32, uint32(PVDL)}}
}

func (p *SYSCFG_Periph) ECCL() RMCFGR2 {
	return RMCFGR2{mmio.UM32{&p.CFGR2.U32, uint32(ECCL)}}
}

func (p *SYSCFG_Periph) SPF() RMCFGR2 {
	return RMCFGR2{mmio.UM32{&p.CFGR2.U32, uint32(SPF)}}
}

type SWPR uint32

func (b SWPR) Field(mask SWPR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SWPR) J(v int) SWPR {
	return SWPR(bits.MakeField32(v, uint32(mask)))
}

type RSWPR struct{ mmio.U32 }

func (r *RSWPR) Bits(mask SWPR) SWPR    { return SWPR(r.U32.Bits(uint32(mask))) }
func (r *RSWPR) StoreBits(mask, b SWPR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSWPR) SetBits(mask SWPR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSWPR) ClearBits(mask SWPR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSWPR) Load() SWPR             { return SWPR(r.U32.Load()) }
func (r *RSWPR) Store(b SWPR)           { r.U32.Store(uint32(b)) }

func (r *RSWPR) AtomicStoreBits(mask, b SWPR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSWPR) AtomicSetBits(mask SWPR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSWPR) AtomicClearBits(mask SWPR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSWPR struct{ mmio.UM32 }

func (rm RMSWPR) Load() SWPR   { return SWPR(rm.UM32.Load()) }
func (rm RMSWPR) Store(b SWPR) { rm.UM32.Store(uint32(b)) }

func (p *SYSCFG_Periph) PAGE0() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE0)}}
}

func (p *SYSCFG_Periph) PAGE1() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE1)}}
}

func (p *SYSCFG_Periph) PAGE2() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE2)}}
}

func (p *SYSCFG_Periph) PAGE3() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE3)}}
}

func (p *SYSCFG_Periph) PAGE4() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE4)}}
}

func (p *SYSCFG_Periph) PAGE5() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE5)}}
}

func (p *SYSCFG_Periph) PAGE6() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE6)}}
}

func (p *SYSCFG_Periph) PAGE7() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE7)}}
}

func (p *SYSCFG_Periph) PAGE8() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE8)}}
}

func (p *SYSCFG_Periph) PAGE9() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE9)}}
}

func (p *SYSCFG_Periph) PAGE10() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE10)}}
}

func (p *SYSCFG_Periph) PAGE11() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE11)}}
}

func (p *SYSCFG_Periph) PAGE12() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE12)}}
}

func (p *SYSCFG_Periph) PAGE13() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE13)}}
}

func (p *SYSCFG_Periph) PAGE14() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE14)}}
}

func (p *SYSCFG_Periph) PAGE15() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE15)}}
}

func (p *SYSCFG_Periph) PAGE16() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE16)}}
}

func (p *SYSCFG_Periph) PAGE17() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE17)}}
}

func (p *SYSCFG_Periph) PAGE18() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE18)}}
}

func (p *SYSCFG_Periph) PAGE19() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE19)}}
}

func (p *SYSCFG_Periph) PAGE20() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE20)}}
}

func (p *SYSCFG_Periph) PAGE21() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE21)}}
}

func (p *SYSCFG_Periph) PAGE22() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE22)}}
}

func (p *SYSCFG_Periph) PAGE23() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE23)}}
}

func (p *SYSCFG_Periph) PAGE24() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE24)}}
}

func (p *SYSCFG_Periph) PAGE25() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE25)}}
}

func (p *SYSCFG_Periph) PAGE26() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE26)}}
}

func (p *SYSCFG_Periph) PAGE27() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE27)}}
}

func (p *SYSCFG_Periph) PAGE28() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE28)}}
}

func (p *SYSCFG_Periph) PAGE29() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE29)}}
}

func (p *SYSCFG_Periph) PAGE30() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE30)}}
}

func (p *SYSCFG_Periph) PAGE31() RMSWPR {
	return RMSWPR{mmio.UM32{&p.SWPR.U32, uint32(PAGE31)}}
}

type SKR uint32

func (b SKR) Field(mask SKR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SKR) J(v int) SKR {
	return SKR(bits.MakeField32(v, uint32(mask)))
}

type RSKR struct{ mmio.U32 }

func (r *RSKR) Bits(mask SKR) SKR     { return SKR(r.U32.Bits(uint32(mask))) }
func (r *RSKR) StoreBits(mask, b SKR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSKR) SetBits(mask SKR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSKR) ClearBits(mask SKR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSKR) Load() SKR             { return SKR(r.U32.Load()) }
func (r *RSKR) Store(b SKR)           { r.U32.Store(uint32(b)) }

func (r *RSKR) AtomicStoreBits(mask, b SKR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSKR) AtomicSetBits(mask SKR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSKR) AtomicClearBits(mask SKR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSKR struct{ mmio.UM32 }

func (rm RMSKR) Load() SKR   { return SKR(rm.UM32.Load()) }
func (rm RMSKR) Store(b SKR) { rm.UM32.Store(uint32(b)) }

func (p *SYSCFG_Periph) KEY() RMSKR {
	return RMSKR{mmio.UM32{&p.SKR.U32, uint32(KEY)}}
}
