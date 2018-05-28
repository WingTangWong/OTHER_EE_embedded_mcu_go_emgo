package sai

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type SAI_Block_Periph struct {
	CR1   RCR1
	CR2   RCR2
	FRCR  RFRCR
	SLOTR RSLOTR
	IMR   RIMR
	SR    RSR
	CLRFR RCLRFR
	DR    RDR
}

func (p *SAI_Block_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var SAI1_Block_A = (*SAI_Block_Periph)(unsafe.Pointer(uintptr(mmap.SAI1_Block_A_BASE)))

//emgo:const
var SAI1_Block_B = (*SAI_Block_Periph)(unsafe.Pointer(uintptr(mmap.SAI1_Block_B_BASE)))

//emgo:const
var SAI2_Block_A = (*SAI_Block_Periph)(unsafe.Pointer(uintptr(mmap.SAI2_Block_A_BASE)))

//emgo:const
var SAI2_Block_B = (*SAI_Block_Periph)(unsafe.Pointer(uintptr(mmap.SAI2_Block_B_BASE)))

type CR1 uint32

func (b CR1) Field(mask CR1) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR1) J(v int) CR1 {
	return CR1(bits.MakeField32(v, uint32(mask)))
}

type RCR1 struct{ mmio.U32 }

func (r *RCR1) Bits(mask CR1) CR1     { return CR1(r.U32.Bits(uint32(mask))) }
func (r *RCR1) StoreBits(mask, b CR1) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCR1) SetBits(mask CR1)      { r.U32.SetBits(uint32(mask)) }
func (r *RCR1) ClearBits(mask CR1)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCR1) Load() CR1             { return CR1(r.U32.Load()) }
func (r *RCR1) Store(b CR1)           { r.U32.Store(uint32(b)) }

func (r *RCR1) AtomicStoreBits(mask, b CR1) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCR1) AtomicSetBits(mask CR1)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCR1) AtomicClearBits(mask CR1)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCR1 struct{ mmio.UM32 }

func (rm RMCR1) Load() CR1   { return CR1(rm.UM32.Load()) }
func (rm RMCR1) Store(b CR1) { rm.UM32.Store(uint32(b)) }

type CR2 uint32

func (b CR2) Field(mask CR2) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR2) J(v int) CR2 {
	return CR2(bits.MakeField32(v, uint32(mask)))
}

type RCR2 struct{ mmio.U32 }

func (r *RCR2) Bits(mask CR2) CR2     { return CR2(r.U32.Bits(uint32(mask))) }
func (r *RCR2) StoreBits(mask, b CR2) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCR2) SetBits(mask CR2)      { r.U32.SetBits(uint32(mask)) }
func (r *RCR2) ClearBits(mask CR2)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCR2) Load() CR2             { return CR2(r.U32.Load()) }
func (r *RCR2) Store(b CR2)           { r.U32.Store(uint32(b)) }

func (r *RCR2) AtomicStoreBits(mask, b CR2) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCR2) AtomicSetBits(mask CR2)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCR2) AtomicClearBits(mask CR2)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCR2 struct{ mmio.UM32 }

func (rm RMCR2) Load() CR2   { return CR2(rm.UM32.Load()) }
func (rm RMCR2) Store(b CR2) { rm.UM32.Store(uint32(b)) }

type FRCR uint32

func (b FRCR) Field(mask FRCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FRCR) J(v int) FRCR {
	return FRCR(bits.MakeField32(v, uint32(mask)))
}

type RFRCR struct{ mmio.U32 }

func (r *RFRCR) Bits(mask FRCR) FRCR    { return FRCR(r.U32.Bits(uint32(mask))) }
func (r *RFRCR) StoreBits(mask, b FRCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RFRCR) SetBits(mask FRCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RFRCR) ClearBits(mask FRCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RFRCR) Load() FRCR             { return FRCR(r.U32.Load()) }
func (r *RFRCR) Store(b FRCR)           { r.U32.Store(uint32(b)) }

func (r *RFRCR) AtomicStoreBits(mask, b FRCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RFRCR) AtomicSetBits(mask FRCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RFRCR) AtomicClearBits(mask FRCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMFRCR struct{ mmio.UM32 }

func (rm RMFRCR) Load() FRCR   { return FRCR(rm.UM32.Load()) }
func (rm RMFRCR) Store(b FRCR) { rm.UM32.Store(uint32(b)) }

type SLOTR uint32

func (b SLOTR) Field(mask SLOTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SLOTR) J(v int) SLOTR {
	return SLOTR(bits.MakeField32(v, uint32(mask)))
}

type RSLOTR struct{ mmio.U32 }

func (r *RSLOTR) Bits(mask SLOTR) SLOTR   { return SLOTR(r.U32.Bits(uint32(mask))) }
func (r *RSLOTR) StoreBits(mask, b SLOTR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RSLOTR) SetBits(mask SLOTR)      { r.U32.SetBits(uint32(mask)) }
func (r *RSLOTR) ClearBits(mask SLOTR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RSLOTR) Load() SLOTR             { return SLOTR(r.U32.Load()) }
func (r *RSLOTR) Store(b SLOTR)           { r.U32.Store(uint32(b)) }

func (r *RSLOTR) AtomicStoreBits(mask, b SLOTR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RSLOTR) AtomicSetBits(mask SLOTR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RSLOTR) AtomicClearBits(mask SLOTR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMSLOTR struct{ mmio.UM32 }

func (rm RMSLOTR) Load() SLOTR   { return SLOTR(rm.UM32.Load()) }
func (rm RMSLOTR) Store(b SLOTR) { rm.UM32.Store(uint32(b)) }

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

type CLRFR uint32

func (b CLRFR) Field(mask CLRFR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CLRFR) J(v int) CLRFR {
	return CLRFR(bits.MakeField32(v, uint32(mask)))
}

type RCLRFR struct{ mmio.U32 }

func (r *RCLRFR) Bits(mask CLRFR) CLRFR   { return CLRFR(r.U32.Bits(uint32(mask))) }
func (r *RCLRFR) StoreBits(mask, b CLRFR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCLRFR) SetBits(mask CLRFR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCLRFR) ClearBits(mask CLRFR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCLRFR) Load() CLRFR             { return CLRFR(r.U32.Load()) }
func (r *RCLRFR) Store(b CLRFR)           { r.U32.Store(uint32(b)) }

func (r *RCLRFR) AtomicStoreBits(mask, b CLRFR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCLRFR) AtomicSetBits(mask CLRFR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCLRFR) AtomicClearBits(mask CLRFR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCLRFR struct{ mmio.UM32 }

func (rm RMCLRFR) Load() CLRFR   { return CLRFR(rm.UM32.Load()) }
func (rm RMCLRFR) Store(b CLRFR) { rm.UM32.Store(uint32(b)) }

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
