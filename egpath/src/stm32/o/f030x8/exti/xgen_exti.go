package exti

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f030x8/mmap"
)

type EXTI_Periph struct {
	IMR   RIMR
	EMR   REMR
	RTSR  RRTSR
	FTSR  RFTSR
	SWIER RSWIER
	PR    RPR
}

func (p *EXTI_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var EXTI = (*EXTI_Periph)(unsafe.Pointer(uintptr(mmap.EXTI_BASE)))

type IMR uint32

func (b IMR) Field(mask IMR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IMR) J(v int) IMR {
	return IMR(bits.MakeField32(v, uint32(mask)))
}

type RIMR struct{ mmio.U32 }

func (r *RIMR) Bits(mask IMR) IMR     { return IMR(r.U32.Bits(uint32(mask))) }
func (r *RIMR) StoreBits(mask, b IMR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RIMR) SetBits(mask IMR)      { r.U32.SetBits(uint32(mask)) }
func (r *RIMR) ClearBits(mask IMR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RIMR) Load() IMR             { return IMR(r.U32.Load()) }
func (r *RIMR) Store(b IMR)           { r.U32.Store(uint32(b)) }

func (r *RIMR) AtomicStoreBits(mask, b IMR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RIMR) AtomicSetBits(mask IMR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RIMR) AtomicClearBits(mask IMR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMIMR struct{ mmio.UM32 }

func (rm RMIMR) Load() IMR   { return IMR(rm.UM32.Load()) }
func (rm RMIMR) Store(b IMR) { rm.UM32.Store(uint32(b)) }

func (p *EXTI_Periph) IL0() RMIMR {
	return RMIMR{mmio.UM32{&p.IMR.U32, uint32(IL0)}}
}

func (p *EXTI_Periph) IL1() RMIMR {
	return RMIMR{mmio.UM32{&p.IMR.U32, uint32(IL1)}}
}

func (p *EXTI_Periph) IL2() RMIMR {
	return RMIMR{mmio.UM32{&p.IMR.U32, uint32(IL2)}}
}

func (p *EXTI_Periph) IL3() RMIMR {
	return RMIMR{mmio.UM32{&p.IMR.U32, uint32(IL3)}}
}

func (p *EXTI_Periph) IL4() RMIMR {
	return RMIMR{mmio.UM32{&p.IMR.U32, uint32(IL4)}}
}

func (p *EXTI_Periph) IL5() RMIMR {
	return RMIMR{mmio.UM32{&p.IMR.U32, uint32(IL5)}}
}

func (p *EXTI_Periph) IL6() RMIMR {
	return RMIMR{mmio.UM32{&p.IMR.U32, uint32(IL6)}}
}

func (p *EXTI_Periph) IL7() RMIMR {
	return RMIMR{mmio.UM32{&p.IMR.U32, uint32(IL7)}}
}

func (p *EXTI_Periph) IL8() RMIMR {
	return RMIMR{mmio.UM32{&p.IMR.U32, uint32(IL8)}}
}

func (p *EXTI_Periph) IL9() RMIMR {
	return RMIMR{mmio.UM32{&p.IMR.U32, uint32(IL9)}}
}

func (p *EXTI_Periph) IL10() RMIMR {
	return RMIMR{mmio.UM32{&p.IMR.U32, uint32(IL10)}}
}

func (p *EXTI_Periph) IL11() RMIMR {
	return RMIMR{mmio.UM32{&p.IMR.U32, uint32(IL11)}}
}

func (p *EXTI_Periph) IL12() RMIMR {
	return RMIMR{mmio.UM32{&p.IMR.U32, uint32(IL12)}}
}

func (p *EXTI_Periph) IL13() RMIMR {
	return RMIMR{mmio.UM32{&p.IMR.U32, uint32(IL13)}}
}

func (p *EXTI_Periph) IL14() RMIMR {
	return RMIMR{mmio.UM32{&p.IMR.U32, uint32(IL14)}}
}

func (p *EXTI_Periph) IL15() RMIMR {
	return RMIMR{mmio.UM32{&p.IMR.U32, uint32(IL15)}}
}

func (p *EXTI_Periph) IL17() RMIMR {
	return RMIMR{mmio.UM32{&p.IMR.U32, uint32(IL17)}}
}

func (p *EXTI_Periph) IL18() RMIMR {
	return RMIMR{mmio.UM32{&p.IMR.U32, uint32(IL18)}}
}

func (p *EXTI_Periph) IL19() RMIMR {
	return RMIMR{mmio.UM32{&p.IMR.U32, uint32(IL19)}}
}

func (p *EXTI_Periph) IL23() RMIMR {
	return RMIMR{mmio.UM32{&p.IMR.U32, uint32(IL23)}}
}

type EMR uint32

func (b EMR) Field(mask EMR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask EMR) J(v int) EMR {
	return EMR(bits.MakeField32(v, uint32(mask)))
}

type REMR struct{ mmio.U32 }

func (r *REMR) Bits(mask EMR) EMR     { return EMR(r.U32.Bits(uint32(mask))) }
func (r *REMR) StoreBits(mask, b EMR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *REMR) SetBits(mask EMR)      { r.U32.SetBits(uint32(mask)) }
func (r *REMR) ClearBits(mask EMR)    { r.U32.ClearBits(uint32(mask)) }
func (r *REMR) Load() EMR             { return EMR(r.U32.Load()) }
func (r *REMR) Store(b EMR)           { r.U32.Store(uint32(b)) }

func (r *REMR) AtomicStoreBits(mask, b EMR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *REMR) AtomicSetBits(mask EMR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *REMR) AtomicClearBits(mask EMR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMEMR struct{ mmio.UM32 }

func (rm RMEMR) Load() EMR   { return EMR(rm.UM32.Load()) }
func (rm RMEMR) Store(b EMR) { rm.UM32.Store(uint32(b)) }

func (p *EXTI_Periph) EL0() RMEMR {
	return RMEMR{mmio.UM32{&p.EMR.U32, uint32(EL0)}}
}

func (p *EXTI_Periph) EL1() RMEMR {
	return RMEMR{mmio.UM32{&p.EMR.U32, uint32(EL1)}}
}

func (p *EXTI_Periph) EL2() RMEMR {
	return RMEMR{mmio.UM32{&p.EMR.U32, uint32(EL2)}}
}

func (p *EXTI_Periph) EL3() RMEMR {
	return RMEMR{mmio.UM32{&p.EMR.U32, uint32(EL3)}}
}

func (p *EXTI_Periph) EL4() RMEMR {
	return RMEMR{mmio.UM32{&p.EMR.U32, uint32(EL4)}}
}

func (p *EXTI_Periph) EL5() RMEMR {
	return RMEMR{mmio.UM32{&p.EMR.U32, uint32(EL5)}}
}

func (p *EXTI_Periph) EL6() RMEMR {
	return RMEMR{mmio.UM32{&p.EMR.U32, uint32(EL6)}}
}

func (p *EXTI_Periph) EL7() RMEMR {
	return RMEMR{mmio.UM32{&p.EMR.U32, uint32(EL7)}}
}

func (p *EXTI_Periph) EL8() RMEMR {
	return RMEMR{mmio.UM32{&p.EMR.U32, uint32(EL8)}}
}

func (p *EXTI_Periph) EL9() RMEMR {
	return RMEMR{mmio.UM32{&p.EMR.U32, uint32(EL9)}}
}

func (p *EXTI_Periph) EL10() RMEMR {
	return RMEMR{mmio.UM32{&p.EMR.U32, uint32(EL10)}}
}

func (p *EXTI_Periph) EL11() RMEMR {
	return RMEMR{mmio.UM32{&p.EMR.U32, uint32(EL11)}}
}

func (p *EXTI_Periph) EL12() RMEMR {
	return RMEMR{mmio.UM32{&p.EMR.U32, uint32(EL12)}}
}

func (p *EXTI_Periph) EL13() RMEMR {
	return RMEMR{mmio.UM32{&p.EMR.U32, uint32(EL13)}}
}

func (p *EXTI_Periph) EL14() RMEMR {
	return RMEMR{mmio.UM32{&p.EMR.U32, uint32(EL14)}}
}

func (p *EXTI_Periph) EL15() RMEMR {
	return RMEMR{mmio.UM32{&p.EMR.U32, uint32(EL15)}}
}

func (p *EXTI_Periph) EL17() RMEMR {
	return RMEMR{mmio.UM32{&p.EMR.U32, uint32(EL17)}}
}

func (p *EXTI_Periph) EL18() RMEMR {
	return RMEMR{mmio.UM32{&p.EMR.U32, uint32(EL18)}}
}

func (p *EXTI_Periph) EL19() RMEMR {
	return RMEMR{mmio.UM32{&p.EMR.U32, uint32(EL19)}}
}

func (p *EXTI_Periph) EL23() RMEMR {
	return RMEMR{mmio.UM32{&p.EMR.U32, uint32(EL23)}}
}

type RTSR uint32

func (b RTSR) Field(mask RTSR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RTSR) J(v int) RTSR {
	return RTSR(bits.MakeField32(v, uint32(mask)))
}

type RRTSR struct{ mmio.U32 }

func (r *RRTSR) Bits(mask RTSR) RTSR    { return RTSR(r.U32.Bits(uint32(mask))) }
func (r *RRTSR) StoreBits(mask, b RTSR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RRTSR) SetBits(mask RTSR)      { r.U32.SetBits(uint32(mask)) }
func (r *RRTSR) ClearBits(mask RTSR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RRTSR) Load() RTSR             { return RTSR(r.U32.Load()) }
func (r *RRTSR) Store(b RTSR)           { r.U32.Store(uint32(b)) }

func (r *RRTSR) AtomicStoreBits(mask, b RTSR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RRTSR) AtomicSetBits(mask RTSR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RRTSR) AtomicClearBits(mask RTSR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMRTSR struct{ mmio.UM32 }

func (rm RMRTSR) Load() RTSR   { return RTSR(rm.UM32.Load()) }
func (rm RMRTSR) Store(b RTSR) { rm.UM32.Store(uint32(b)) }

func (p *EXTI_Periph) TR0() RMRTSR {
	return RMRTSR{mmio.UM32{&p.RTSR.U32, uint32(TR0)}}
}

func (p *EXTI_Periph) TR1() RMRTSR {
	return RMRTSR{mmio.UM32{&p.RTSR.U32, uint32(TR1)}}
}

func (p *EXTI_Periph) TR2() RMRTSR {
	return RMRTSR{mmio.UM32{&p.RTSR.U32, uint32(TR2)}}
}

func (p *EXTI_Periph) TR3() RMRTSR {
	return RMRTSR{mmio.UM32{&p.RTSR.U32, uint32(TR3)}}
}

func (p *EXTI_Periph) TR4() RMRTSR {
	return RMRTSR{mmio.UM32{&p.RTSR.U32, uint32(TR4)}}
}

func (p *EXTI_Periph) TR5() RMRTSR {
	return RMRTSR{mmio.UM32{&p.RTSR.U32, uint32(TR5)}}
}

func (p *EXTI_Periph) TR6() RMRTSR {
	return RMRTSR{mmio.UM32{&p.RTSR.U32, uint32(TR6)}}
}

func (p *EXTI_Periph) TR7() RMRTSR {
	return RMRTSR{mmio.UM32{&p.RTSR.U32, uint32(TR7)}}
}

func (p *EXTI_Periph) TR8() RMRTSR {
	return RMRTSR{mmio.UM32{&p.RTSR.U32, uint32(TR8)}}
}

func (p *EXTI_Periph) TR9() RMRTSR {
	return RMRTSR{mmio.UM32{&p.RTSR.U32, uint32(TR9)}}
}

func (p *EXTI_Periph) TR10() RMRTSR {
	return RMRTSR{mmio.UM32{&p.RTSR.U32, uint32(TR10)}}
}

func (p *EXTI_Periph) TR11() RMRTSR {
	return RMRTSR{mmio.UM32{&p.RTSR.U32, uint32(TR11)}}
}

func (p *EXTI_Periph) TR12() RMRTSR {
	return RMRTSR{mmio.UM32{&p.RTSR.U32, uint32(TR12)}}
}

func (p *EXTI_Periph) TR13() RMRTSR {
	return RMRTSR{mmio.UM32{&p.RTSR.U32, uint32(TR13)}}
}

func (p *EXTI_Periph) TR14() RMRTSR {
	return RMRTSR{mmio.UM32{&p.RTSR.U32, uint32(TR14)}}
}

func (p *EXTI_Periph) TR15() RMRTSR {
	return RMRTSR{mmio.UM32{&p.RTSR.U32, uint32(TR15)}}
}

func (p *EXTI_Periph) TR16() RMRTSR {
	return RMRTSR{mmio.UM32{&p.RTSR.U32, uint32(TR16)}}
}

func (p *EXTI_Periph) TR17() RMRTSR {
	return RMRTSR{mmio.UM32{&p.RTSR.U32, uint32(TR17)}}
}

func (p *EXTI_Periph) TR19() RMRTSR {
	return RMRTSR{mmio.UM32{&p.RTSR.U32, uint32(TR19)}}
}

type FTSR uint32

func (b FTSR) Field(mask FTSR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FTSR) J(v int) FTSR {
	return FTSR(bits.MakeField32(v, uint32(mask)))
}

type RFTSR struct{ mmio.U32 }

func (r *RFTSR) Bits(mask FTSR) FTSR    { return FTSR(r.U32.Bits(uint32(mask))) }
func (r *RFTSR) StoreBits(mask, b FTSR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RFTSR) SetBits(mask FTSR)      { r.U32.SetBits(uint32(mask)) }
func (r *RFTSR) ClearBits(mask FTSR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RFTSR) Load() FTSR             { return FTSR(r.U32.Load()) }
func (r *RFTSR) Store(b FTSR)           { r.U32.Store(uint32(b)) }

func (r *RFTSR) AtomicStoreBits(mask, b FTSR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RFTSR) AtomicSetBits(mask FTSR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RFTSR) AtomicClearBits(mask FTSR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMFTSR struct{ mmio.UM32 }

func (rm RMFTSR) Load() FTSR   { return FTSR(rm.UM32.Load()) }
func (rm RMFTSR) Store(b FTSR) { rm.UM32.Store(uint32(b)) }

func (p *EXTI_Periph) TF0() RMFTSR {
	return RMFTSR{mmio.UM32{&p.FTSR.U32, uint32(TF0)}}
}

func (p *EXTI_Periph) TF1() RMFTSR {
	return RMFTSR{mmio.UM32{&p.FTSR.U32, uint32(TF1)}}
}

func (p *EXTI_Periph) TF2() RMFTSR {
	return RMFTSR{mmio.UM32{&p.FTSR.U32, uint32(TF2)}}
}

func (p *EXTI_Periph) TF3() RMFTSR {
	return RMFTSR{mmio.UM32{&p.FTSR.U32, uint32(TF3)}}
}

func (p *EXTI_Periph) TF4() RMFTSR {
	return RMFTSR{mmio.UM32{&p.FTSR.U32, uint32(TF4)}}
}

func (p *EXTI_Periph) TF5() RMFTSR {
	return RMFTSR{mmio.UM32{&p.FTSR.U32, uint32(TF5)}}
}

func (p *EXTI_Periph) TF6() RMFTSR {
	return RMFTSR{mmio.UM32{&p.FTSR.U32, uint32(TF6)}}
}

func (p *EXTI_Periph) TF7() RMFTSR {
	return RMFTSR{mmio.UM32{&p.FTSR.U32, uint32(TF7)}}
}

func (p *EXTI_Periph) TF8() RMFTSR {
	return RMFTSR{mmio.UM32{&p.FTSR.U32, uint32(TF8)}}
}

func (p *EXTI_Periph) TF9() RMFTSR {
	return RMFTSR{mmio.UM32{&p.FTSR.U32, uint32(TF9)}}
}

func (p *EXTI_Periph) TF10() RMFTSR {
	return RMFTSR{mmio.UM32{&p.FTSR.U32, uint32(TF10)}}
}

func (p *EXTI_Periph) TF11() RMFTSR {
	return RMFTSR{mmio.UM32{&p.FTSR.U32, uint32(TF11)}}
}

func (p *EXTI_Periph) TF12() RMFTSR {
	return RMFTSR{mmio.UM32{&p.FTSR.U32, uint32(TF12)}}
}

func (p *EXTI_Periph) TF13() RMFTSR {
	return RMFTSR{mmio.UM32{&p.FTSR.U32, uint32(TF13)}}
}

func (p *EXTI_Periph) TF14() RMFTSR {
	return RMFTSR{mmio.UM32{&p.FTSR.U32, uint32(TF14)}}
}

func (p *EXTI_Periph) TF15() RMFTSR {
	return RMFTSR{mmio.UM32{&p.FTSR.U32, uint32(TF15)}}
}

func (p *EXTI_Periph) TF16() RMFTSR {
	return RMFTSR{mmio.UM32{&p.FTSR.U32, uint32(TF16)}}
}

func (p *EXTI_Periph) TF17() RMFTSR {
	return RMFTSR{mmio.UM32{&p.FTSR.U32, uint32(TF17)}}
}

func (p *EXTI_Periph) TF19() RMFTSR {
	return RMFTSR{mmio.UM32{&p.FTSR.U32, uint32(TF19)}}
}

type SWIER uint32

func (b SWIER) Field(mask SWIER) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SWIER) J(v int) SWIER {
	return SWIER(bits.MakeField32(v, uint32(mask)))
}

type RSWIER struct{ mmio.U32 }

func (r *RSWIER) Bits(mask SWIER) SWIER   { return SWIER(r.U32.Bits(uint32(mask))) }
func (r *RSWIER) StoreBits(mask, b SWIER) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSWIER) SetBits(mask SWIER)      { r.U32.SetBits(uint32(mask)) }
func (r *RSWIER) ClearBits(mask SWIER)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSWIER) Load() SWIER             { return SWIER(r.U32.Load()) }
func (r *RSWIER) Store(b SWIER)           { r.U32.Store(uint32(b)) }

func (r *RSWIER) AtomicStoreBits(mask, b SWIER) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSWIER) AtomicSetBits(mask SWIER)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSWIER) AtomicClearBits(mask SWIER)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSWIER struct{ mmio.UM32 }

func (rm RMSWIER) Load() SWIER   { return SWIER(rm.UM32.Load()) }
func (rm RMSWIER) Store(b SWIER) { rm.UM32.Store(uint32(b)) }

func (p *EXTI_Periph) SWI0() RMSWIER {
	return RMSWIER{mmio.UM32{&p.SWIER.U32, uint32(SWI0)}}
}

func (p *EXTI_Periph) SWI1() RMSWIER {
	return RMSWIER{mmio.UM32{&p.SWIER.U32, uint32(SWI1)}}
}

func (p *EXTI_Periph) SWI2() RMSWIER {
	return RMSWIER{mmio.UM32{&p.SWIER.U32, uint32(SWI2)}}
}

func (p *EXTI_Periph) SWI3() RMSWIER {
	return RMSWIER{mmio.UM32{&p.SWIER.U32, uint32(SWI3)}}
}

func (p *EXTI_Periph) SWI4() RMSWIER {
	return RMSWIER{mmio.UM32{&p.SWIER.U32, uint32(SWI4)}}
}

func (p *EXTI_Periph) SWI5() RMSWIER {
	return RMSWIER{mmio.UM32{&p.SWIER.U32, uint32(SWI5)}}
}

func (p *EXTI_Periph) SWI6() RMSWIER {
	return RMSWIER{mmio.UM32{&p.SWIER.U32, uint32(SWI6)}}
}

func (p *EXTI_Periph) SWI7() RMSWIER {
	return RMSWIER{mmio.UM32{&p.SWIER.U32, uint32(SWI7)}}
}

func (p *EXTI_Periph) SWI8() RMSWIER {
	return RMSWIER{mmio.UM32{&p.SWIER.U32, uint32(SWI8)}}
}

func (p *EXTI_Periph) SWI9() RMSWIER {
	return RMSWIER{mmio.UM32{&p.SWIER.U32, uint32(SWI9)}}
}

func (p *EXTI_Periph) SWI10() RMSWIER {
	return RMSWIER{mmio.UM32{&p.SWIER.U32, uint32(SWI10)}}
}

func (p *EXTI_Periph) SWI11() RMSWIER {
	return RMSWIER{mmio.UM32{&p.SWIER.U32, uint32(SWI11)}}
}

func (p *EXTI_Periph) SWI12() RMSWIER {
	return RMSWIER{mmio.UM32{&p.SWIER.U32, uint32(SWI12)}}
}

func (p *EXTI_Periph) SWI13() RMSWIER {
	return RMSWIER{mmio.UM32{&p.SWIER.U32, uint32(SWI13)}}
}

func (p *EXTI_Periph) SWI14() RMSWIER {
	return RMSWIER{mmio.UM32{&p.SWIER.U32, uint32(SWI14)}}
}

func (p *EXTI_Periph) SWI15() RMSWIER {
	return RMSWIER{mmio.UM32{&p.SWIER.U32, uint32(SWI15)}}
}

func (p *EXTI_Periph) SWI16() RMSWIER {
	return RMSWIER{mmio.UM32{&p.SWIER.U32, uint32(SWI16)}}
}

func (p *EXTI_Periph) SWI17() RMSWIER {
	return RMSWIER{mmio.UM32{&p.SWIER.U32, uint32(SWI17)}}
}

func (p *EXTI_Periph) SWI19() RMSWIER {
	return RMSWIER{mmio.UM32{&p.SWIER.U32, uint32(SWI19)}}
}

type PR uint32

func (b PR) Field(mask PR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PR) J(v int) PR {
	return PR(bits.MakeField32(v, uint32(mask)))
}

type RPR struct{ mmio.U32 }

func (r *RPR) Bits(mask PR) PR      { return PR(r.U32.Bits(uint32(mask))) }
func (r *RPR) StoreBits(mask, b PR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPR) SetBits(mask PR)      { r.U32.SetBits(uint32(mask)) }
func (r *RPR) ClearBits(mask PR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPR) Load() PR             { return PR(r.U32.Load()) }
func (r *RPR) Store(b PR)           { r.U32.Store(uint32(b)) }

func (r *RPR) AtomicStoreBits(mask, b PR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPR) AtomicSetBits(mask PR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPR) AtomicClearBits(mask PR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPR struct{ mmio.UM32 }

func (rm RMPR) Load() PR   { return PR(rm.UM32.Load()) }
func (rm RMPR) Store(b PR) { rm.UM32.Store(uint32(b)) }

func (p *EXTI_Periph) PIF0() RMPR {
	return RMPR{mmio.UM32{&p.PR.U32, uint32(PIF0)}}
}

func (p *EXTI_Periph) PIF1() RMPR {
	return RMPR{mmio.UM32{&p.PR.U32, uint32(PIF1)}}
}

func (p *EXTI_Periph) PIF2() RMPR {
	return RMPR{mmio.UM32{&p.PR.U32, uint32(PIF2)}}
}

func (p *EXTI_Periph) PIF3() RMPR {
	return RMPR{mmio.UM32{&p.PR.U32, uint32(PIF3)}}
}

func (p *EXTI_Periph) PIF4() RMPR {
	return RMPR{mmio.UM32{&p.PR.U32, uint32(PIF4)}}
}

func (p *EXTI_Periph) PIF5() RMPR {
	return RMPR{mmio.UM32{&p.PR.U32, uint32(PIF5)}}
}

func (p *EXTI_Periph) PIF6() RMPR {
	return RMPR{mmio.UM32{&p.PR.U32, uint32(PIF6)}}
}

func (p *EXTI_Periph) PIF7() RMPR {
	return RMPR{mmio.UM32{&p.PR.U32, uint32(PIF7)}}
}

func (p *EXTI_Periph) PIF8() RMPR {
	return RMPR{mmio.UM32{&p.PR.U32, uint32(PIF8)}}
}

func (p *EXTI_Periph) PIF9() RMPR {
	return RMPR{mmio.UM32{&p.PR.U32, uint32(PIF9)}}
}

func (p *EXTI_Periph) PIF10() RMPR {
	return RMPR{mmio.UM32{&p.PR.U32, uint32(PIF10)}}
}

func (p *EXTI_Periph) PIF11() RMPR {
	return RMPR{mmio.UM32{&p.PR.U32, uint32(PIF11)}}
}

func (p *EXTI_Periph) PIF12() RMPR {
	return RMPR{mmio.UM32{&p.PR.U32, uint32(PIF12)}}
}

func (p *EXTI_Periph) PIF13() RMPR {
	return RMPR{mmio.UM32{&p.PR.U32, uint32(PIF13)}}
}

func (p *EXTI_Periph) PIF14() RMPR {
	return RMPR{mmio.UM32{&p.PR.U32, uint32(PIF14)}}
}

func (p *EXTI_Periph) PIF15() RMPR {
	return RMPR{mmio.UM32{&p.PR.U32, uint32(PIF15)}}
}

func (p *EXTI_Periph) PIF16() RMPR {
	return RMPR{mmio.UM32{&p.PR.U32, uint32(PIF16)}}
}

func (p *EXTI_Periph) PIF17() RMPR {
	return RMPR{mmio.UM32{&p.PR.U32, uint32(PIF17)}}
}

func (p *EXTI_Periph) PIF19() RMPR {
	return RMPR{mmio.UM32{&p.PR.U32, uint32(PIF19)}}
}
