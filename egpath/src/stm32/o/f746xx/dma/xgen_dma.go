package dma

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f746xx/mmap"
)

type DMA_Periph struct {
	LISR  RLISR
	HISR  RHISR
	LIFCR RLIFCR
	HIFCR RHIFCR
}

func (p *DMA_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var DMA1 = (*DMA_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_BASE)))

//emgo:const
var DMA2 = (*DMA_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_BASE)))

type LISR uint32

func (b LISR) Field(mask LISR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask LISR) J(v int) LISR {
	return LISR(bits.MakeField32(v, uint32(mask)))
}

type RLISR struct{ mmio.U32 }

func (r *RLISR) Bits(mask LISR) LISR    { return LISR(r.U32.Bits(uint32(mask))) }
func (r *RLISR) StoreBits(mask, b LISR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RLISR) SetBits(mask LISR)      { r.U32.SetBits(uint32(mask)) }
func (r *RLISR) ClearBits(mask LISR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RLISR) Load() LISR             { return LISR(r.U32.Load()) }
func (r *RLISR) Store(b LISR)           { r.U32.Store(uint32(b)) }

func (r *RLISR) AtomicStoreBits(mask, b LISR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RLISR) AtomicSetBits(mask LISR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RLISR) AtomicClearBits(mask LISR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMLISR struct{ mmio.UM32 }

func (rm RMLISR) Load() LISR   { return LISR(rm.UM32.Load()) }
func (rm RMLISR) Store(b LISR) { rm.UM32.Store(uint32(b)) }

func (p *DMA_Periph) TCIF3() RMLISR {
	return RMLISR{mmio.UM32{&p.LISR.U32, uint32(TCIF3)}}
}

func (p *DMA_Periph) HTIF3() RMLISR {
	return RMLISR{mmio.UM32{&p.LISR.U32, uint32(HTIF3)}}
}

func (p *DMA_Periph) TEIF3() RMLISR {
	return RMLISR{mmio.UM32{&p.LISR.U32, uint32(TEIF3)}}
}

func (p *DMA_Periph) DMEIF3() RMLISR {
	return RMLISR{mmio.UM32{&p.LISR.U32, uint32(DMEIF3)}}
}

func (p *DMA_Periph) FEIF3() RMLISR {
	return RMLISR{mmio.UM32{&p.LISR.U32, uint32(FEIF3)}}
}

func (p *DMA_Periph) TCIF2() RMLISR {
	return RMLISR{mmio.UM32{&p.LISR.U32, uint32(TCIF2)}}
}

func (p *DMA_Periph) HTIF2() RMLISR {
	return RMLISR{mmio.UM32{&p.LISR.U32, uint32(HTIF2)}}
}

func (p *DMA_Periph) TEIF2() RMLISR {
	return RMLISR{mmio.UM32{&p.LISR.U32, uint32(TEIF2)}}
}

func (p *DMA_Periph) DMEIF2() RMLISR {
	return RMLISR{mmio.UM32{&p.LISR.U32, uint32(DMEIF2)}}
}

func (p *DMA_Periph) FEIF2() RMLISR {
	return RMLISR{mmio.UM32{&p.LISR.U32, uint32(FEIF2)}}
}

func (p *DMA_Periph) TCIF1() RMLISR {
	return RMLISR{mmio.UM32{&p.LISR.U32, uint32(TCIF1)}}
}

func (p *DMA_Periph) HTIF1() RMLISR {
	return RMLISR{mmio.UM32{&p.LISR.U32, uint32(HTIF1)}}
}

func (p *DMA_Periph) TEIF1() RMLISR {
	return RMLISR{mmio.UM32{&p.LISR.U32, uint32(TEIF1)}}
}

func (p *DMA_Periph) DMEIF1() RMLISR {
	return RMLISR{mmio.UM32{&p.LISR.U32, uint32(DMEIF1)}}
}

func (p *DMA_Periph) FEIF1() RMLISR {
	return RMLISR{mmio.UM32{&p.LISR.U32, uint32(FEIF1)}}
}

func (p *DMA_Periph) TCIF0() RMLISR {
	return RMLISR{mmio.UM32{&p.LISR.U32, uint32(TCIF0)}}
}

func (p *DMA_Periph) HTIF0() RMLISR {
	return RMLISR{mmio.UM32{&p.LISR.U32, uint32(HTIF0)}}
}

func (p *DMA_Periph) TEIF0() RMLISR {
	return RMLISR{mmio.UM32{&p.LISR.U32, uint32(TEIF0)}}
}

func (p *DMA_Periph) DMEIF0() RMLISR {
	return RMLISR{mmio.UM32{&p.LISR.U32, uint32(DMEIF0)}}
}

func (p *DMA_Periph) FEIF0() RMLISR {
	return RMLISR{mmio.UM32{&p.LISR.U32, uint32(FEIF0)}}
}

type HISR uint32

func (b HISR) Field(mask HISR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HISR) J(v int) HISR {
	return HISR(bits.MakeField32(v, uint32(mask)))
}

type RHISR struct{ mmio.U32 }

func (r *RHISR) Bits(mask HISR) HISR    { return HISR(r.U32.Bits(uint32(mask))) }
func (r *RHISR) StoreBits(mask, b HISR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RHISR) SetBits(mask HISR)      { r.U32.SetBits(uint32(mask)) }
func (r *RHISR) ClearBits(mask HISR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RHISR) Load() HISR             { return HISR(r.U32.Load()) }
func (r *RHISR) Store(b HISR)           { r.U32.Store(uint32(b)) }

func (r *RHISR) AtomicStoreBits(mask, b HISR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RHISR) AtomicSetBits(mask HISR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RHISR) AtomicClearBits(mask HISR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMHISR struct{ mmio.UM32 }

func (rm RMHISR) Load() HISR   { return HISR(rm.UM32.Load()) }
func (rm RMHISR) Store(b HISR) { rm.UM32.Store(uint32(b)) }

func (p *DMA_Periph) TCIF7() RMHISR {
	return RMHISR{mmio.UM32{&p.HISR.U32, uint32(TCIF7)}}
}

func (p *DMA_Periph) HTIF7() RMHISR {
	return RMHISR{mmio.UM32{&p.HISR.U32, uint32(HTIF7)}}
}

func (p *DMA_Periph) TEIF7() RMHISR {
	return RMHISR{mmio.UM32{&p.HISR.U32, uint32(TEIF7)}}
}

func (p *DMA_Periph) DMEIF7() RMHISR {
	return RMHISR{mmio.UM32{&p.HISR.U32, uint32(DMEIF7)}}
}

func (p *DMA_Periph) FEIF7() RMHISR {
	return RMHISR{mmio.UM32{&p.HISR.U32, uint32(FEIF7)}}
}

func (p *DMA_Periph) TCIF6() RMHISR {
	return RMHISR{mmio.UM32{&p.HISR.U32, uint32(TCIF6)}}
}

func (p *DMA_Periph) HTIF6() RMHISR {
	return RMHISR{mmio.UM32{&p.HISR.U32, uint32(HTIF6)}}
}

func (p *DMA_Periph) TEIF6() RMHISR {
	return RMHISR{mmio.UM32{&p.HISR.U32, uint32(TEIF6)}}
}

func (p *DMA_Periph) DMEIF6() RMHISR {
	return RMHISR{mmio.UM32{&p.HISR.U32, uint32(DMEIF6)}}
}

func (p *DMA_Periph) FEIF6() RMHISR {
	return RMHISR{mmio.UM32{&p.HISR.U32, uint32(FEIF6)}}
}

func (p *DMA_Periph) TCIF5() RMHISR {
	return RMHISR{mmio.UM32{&p.HISR.U32, uint32(TCIF5)}}
}

func (p *DMA_Periph) HTIF5() RMHISR {
	return RMHISR{mmio.UM32{&p.HISR.U32, uint32(HTIF5)}}
}

func (p *DMA_Periph) TEIF5() RMHISR {
	return RMHISR{mmio.UM32{&p.HISR.U32, uint32(TEIF5)}}
}

func (p *DMA_Periph) DMEIF5() RMHISR {
	return RMHISR{mmio.UM32{&p.HISR.U32, uint32(DMEIF5)}}
}

func (p *DMA_Periph) FEIF5() RMHISR {
	return RMHISR{mmio.UM32{&p.HISR.U32, uint32(FEIF5)}}
}

func (p *DMA_Periph) TCIF4() RMHISR {
	return RMHISR{mmio.UM32{&p.HISR.U32, uint32(TCIF4)}}
}

func (p *DMA_Periph) HTIF4() RMHISR {
	return RMHISR{mmio.UM32{&p.HISR.U32, uint32(HTIF4)}}
}

func (p *DMA_Periph) TEIF4() RMHISR {
	return RMHISR{mmio.UM32{&p.HISR.U32, uint32(TEIF4)}}
}

func (p *DMA_Periph) DMEIF4() RMHISR {
	return RMHISR{mmio.UM32{&p.HISR.U32, uint32(DMEIF4)}}
}

func (p *DMA_Periph) FEIF4() RMHISR {
	return RMHISR{mmio.UM32{&p.HISR.U32, uint32(FEIF4)}}
}

type LIFCR uint32

func (b LIFCR) Field(mask LIFCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask LIFCR) J(v int) LIFCR {
	return LIFCR(bits.MakeField32(v, uint32(mask)))
}

type RLIFCR struct{ mmio.U32 }

func (r *RLIFCR) Bits(mask LIFCR) LIFCR   { return LIFCR(r.U32.Bits(uint32(mask))) }
func (r *RLIFCR) StoreBits(mask, b LIFCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RLIFCR) SetBits(mask LIFCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RLIFCR) ClearBits(mask LIFCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RLIFCR) Load() LIFCR             { return LIFCR(r.U32.Load()) }
func (r *RLIFCR) Store(b LIFCR)           { r.U32.Store(uint32(b)) }

func (r *RLIFCR) AtomicStoreBits(mask, b LIFCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RLIFCR) AtomicSetBits(mask LIFCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RLIFCR) AtomicClearBits(mask LIFCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMLIFCR struct{ mmio.UM32 }

func (rm RMLIFCR) Load() LIFCR   { return LIFCR(rm.UM32.Load()) }
func (rm RMLIFCR) Store(b LIFCR) { rm.UM32.Store(uint32(b)) }

func (p *DMA_Periph) CTCIF3() RMLIFCR {
	return RMLIFCR{mmio.UM32{&p.LIFCR.U32, uint32(CTCIF3)}}
}

func (p *DMA_Periph) CHTIF3() RMLIFCR {
	return RMLIFCR{mmio.UM32{&p.LIFCR.U32, uint32(CHTIF3)}}
}

func (p *DMA_Periph) CTEIF3() RMLIFCR {
	return RMLIFCR{mmio.UM32{&p.LIFCR.U32, uint32(CTEIF3)}}
}

func (p *DMA_Periph) CDMEIF3() RMLIFCR {
	return RMLIFCR{mmio.UM32{&p.LIFCR.U32, uint32(CDMEIF3)}}
}

func (p *DMA_Periph) CFEIF3() RMLIFCR {
	return RMLIFCR{mmio.UM32{&p.LIFCR.U32, uint32(CFEIF3)}}
}

func (p *DMA_Periph) CTCIF2() RMLIFCR {
	return RMLIFCR{mmio.UM32{&p.LIFCR.U32, uint32(CTCIF2)}}
}

func (p *DMA_Periph) CHTIF2() RMLIFCR {
	return RMLIFCR{mmio.UM32{&p.LIFCR.U32, uint32(CHTIF2)}}
}

func (p *DMA_Periph) CTEIF2() RMLIFCR {
	return RMLIFCR{mmio.UM32{&p.LIFCR.U32, uint32(CTEIF2)}}
}

func (p *DMA_Periph) CDMEIF2() RMLIFCR {
	return RMLIFCR{mmio.UM32{&p.LIFCR.U32, uint32(CDMEIF2)}}
}

func (p *DMA_Periph) CFEIF2() RMLIFCR {
	return RMLIFCR{mmio.UM32{&p.LIFCR.U32, uint32(CFEIF2)}}
}

func (p *DMA_Periph) CTCIF1() RMLIFCR {
	return RMLIFCR{mmio.UM32{&p.LIFCR.U32, uint32(CTCIF1)}}
}

func (p *DMA_Periph) CHTIF1() RMLIFCR {
	return RMLIFCR{mmio.UM32{&p.LIFCR.U32, uint32(CHTIF1)}}
}

func (p *DMA_Periph) CTEIF1() RMLIFCR {
	return RMLIFCR{mmio.UM32{&p.LIFCR.U32, uint32(CTEIF1)}}
}

func (p *DMA_Periph) CDMEIF1() RMLIFCR {
	return RMLIFCR{mmio.UM32{&p.LIFCR.U32, uint32(CDMEIF1)}}
}

func (p *DMA_Periph) CFEIF1() RMLIFCR {
	return RMLIFCR{mmio.UM32{&p.LIFCR.U32, uint32(CFEIF1)}}
}

func (p *DMA_Periph) CTCIF0() RMLIFCR {
	return RMLIFCR{mmio.UM32{&p.LIFCR.U32, uint32(CTCIF0)}}
}

func (p *DMA_Periph) CHTIF0() RMLIFCR {
	return RMLIFCR{mmio.UM32{&p.LIFCR.U32, uint32(CHTIF0)}}
}

func (p *DMA_Periph) CTEIF0() RMLIFCR {
	return RMLIFCR{mmio.UM32{&p.LIFCR.U32, uint32(CTEIF0)}}
}

func (p *DMA_Periph) CDMEIF0() RMLIFCR {
	return RMLIFCR{mmio.UM32{&p.LIFCR.U32, uint32(CDMEIF0)}}
}

func (p *DMA_Periph) CFEIF0() RMLIFCR {
	return RMLIFCR{mmio.UM32{&p.LIFCR.U32, uint32(CFEIF0)}}
}

type HIFCR uint32

func (b HIFCR) Field(mask HIFCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HIFCR) J(v int) HIFCR {
	return HIFCR(bits.MakeField32(v, uint32(mask)))
}

type RHIFCR struct{ mmio.U32 }

func (r *RHIFCR) Bits(mask HIFCR) HIFCR   { return HIFCR(r.U32.Bits(uint32(mask))) }
func (r *RHIFCR) StoreBits(mask, b HIFCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RHIFCR) SetBits(mask HIFCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RHIFCR) ClearBits(mask HIFCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RHIFCR) Load() HIFCR             { return HIFCR(r.U32.Load()) }
func (r *RHIFCR) Store(b HIFCR)           { r.U32.Store(uint32(b)) }

func (r *RHIFCR) AtomicStoreBits(mask, b HIFCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RHIFCR) AtomicSetBits(mask HIFCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RHIFCR) AtomicClearBits(mask HIFCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMHIFCR struct{ mmio.UM32 }

func (rm RMHIFCR) Load() HIFCR   { return HIFCR(rm.UM32.Load()) }
func (rm RMHIFCR) Store(b HIFCR) { rm.UM32.Store(uint32(b)) }

func (p *DMA_Periph) CTCIF7() RMHIFCR {
	return RMHIFCR{mmio.UM32{&p.HIFCR.U32, uint32(CTCIF7)}}
}

func (p *DMA_Periph) CHTIF7() RMHIFCR {
	return RMHIFCR{mmio.UM32{&p.HIFCR.U32, uint32(CHTIF7)}}
}

func (p *DMA_Periph) CTEIF7() RMHIFCR {
	return RMHIFCR{mmio.UM32{&p.HIFCR.U32, uint32(CTEIF7)}}
}

func (p *DMA_Periph) CDMEIF7() RMHIFCR {
	return RMHIFCR{mmio.UM32{&p.HIFCR.U32, uint32(CDMEIF7)}}
}

func (p *DMA_Periph) CFEIF7() RMHIFCR {
	return RMHIFCR{mmio.UM32{&p.HIFCR.U32, uint32(CFEIF7)}}
}

func (p *DMA_Periph) CTCIF6() RMHIFCR {
	return RMHIFCR{mmio.UM32{&p.HIFCR.U32, uint32(CTCIF6)}}
}

func (p *DMA_Periph) CHTIF6() RMHIFCR {
	return RMHIFCR{mmio.UM32{&p.HIFCR.U32, uint32(CHTIF6)}}
}

func (p *DMA_Periph) CTEIF6() RMHIFCR {
	return RMHIFCR{mmio.UM32{&p.HIFCR.U32, uint32(CTEIF6)}}
}

func (p *DMA_Periph) CDMEIF6() RMHIFCR {
	return RMHIFCR{mmio.UM32{&p.HIFCR.U32, uint32(CDMEIF6)}}
}

func (p *DMA_Periph) CFEIF6() RMHIFCR {
	return RMHIFCR{mmio.UM32{&p.HIFCR.U32, uint32(CFEIF6)}}
}

func (p *DMA_Periph) CTCIF5() RMHIFCR {
	return RMHIFCR{mmio.UM32{&p.HIFCR.U32, uint32(CTCIF5)}}
}

func (p *DMA_Periph) CHTIF5() RMHIFCR {
	return RMHIFCR{mmio.UM32{&p.HIFCR.U32, uint32(CHTIF5)}}
}

func (p *DMA_Periph) CTEIF5() RMHIFCR {
	return RMHIFCR{mmio.UM32{&p.HIFCR.U32, uint32(CTEIF5)}}
}

func (p *DMA_Periph) CDMEIF5() RMHIFCR {
	return RMHIFCR{mmio.UM32{&p.HIFCR.U32, uint32(CDMEIF5)}}
}

func (p *DMA_Periph) CFEIF5() RMHIFCR {
	return RMHIFCR{mmio.UM32{&p.HIFCR.U32, uint32(CFEIF5)}}
}

func (p *DMA_Periph) CTCIF4() RMHIFCR {
	return RMHIFCR{mmio.UM32{&p.HIFCR.U32, uint32(CTCIF4)}}
}

func (p *DMA_Periph) CHTIF4() RMHIFCR {
	return RMHIFCR{mmio.UM32{&p.HIFCR.U32, uint32(CHTIF4)}}
}

func (p *DMA_Periph) CTEIF4() RMHIFCR {
	return RMHIFCR{mmio.UM32{&p.HIFCR.U32, uint32(CTEIF4)}}
}

func (p *DMA_Periph) CDMEIF4() RMHIFCR {
	return RMHIFCR{mmio.UM32{&p.HIFCR.U32, uint32(CDMEIF4)}}
}

func (p *DMA_Periph) CFEIF4() RMHIFCR {
	return RMHIFCR{mmio.UM32{&p.HIFCR.U32, uint32(CFEIF4)}}
}
