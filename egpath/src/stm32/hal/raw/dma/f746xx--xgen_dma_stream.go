// +build f746xx

package dma

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f746xx/mmap"
)

type DMA_Stream_Periph struct {
	CR   RCR
	NDTR RNDTR
	PAR  RPAR
	M0AR RM0AR
	M1AR RM1AR
	FCR  RFCR
}

func (p *DMA_Stream_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var DMA1_Stream0 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Stream0_BASE)))

//emgo:const
var DMA1_Stream1 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Stream1_BASE)))

//emgo:const
var DMA1_Stream2 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Stream2_BASE)))

//emgo:const
var DMA1_Stream3 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Stream3_BASE)))

//emgo:const
var DMA1_Stream4 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Stream4_BASE)))

//emgo:const
var DMA1_Stream5 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Stream5_BASE)))

//emgo:const
var DMA1_Stream6 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Stream6_BASE)))

//emgo:const
var DMA1_Stream7 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA1_Stream7_BASE)))

//emgo:const
var DMA2_Stream0 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Stream0_BASE)))

//emgo:const
var DMA2_Stream1 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Stream1_BASE)))

//emgo:const
var DMA2_Stream2 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Stream2_BASE)))

//emgo:const
var DMA2_Stream3 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Stream3_BASE)))

//emgo:const
var DMA2_Stream4 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Stream4_BASE)))

//emgo:const
var DMA2_Stream5 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Stream5_BASE)))

//emgo:const
var DMA2_Stream6 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Stream6_BASE)))

//emgo:const
var DMA2_Stream7 = (*DMA_Stream_Periph)(unsafe.Pointer(uintptr(mmap.DMA2_Stream7_BASE)))

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

func (p *DMA_Stream_Periph) CHSEL() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(CHSEL)}}
}

func (p *DMA_Stream_Periph) MBURST() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(MBURST)}}
}

func (p *DMA_Stream_Periph) PBURST() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PBURST)}}
}

func (p *DMA_Stream_Periph) CT() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(CT)}}
}

func (p *DMA_Stream_Periph) DBM() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DBM)}}
}

func (p *DMA_Stream_Periph) PL() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PL)}}
}

func (p *DMA_Stream_Periph) PINCOS() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PINCOS)}}
}

func (p *DMA_Stream_Periph) MSIZE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(MSIZE)}}
}

func (p *DMA_Stream_Periph) PSIZE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PSIZE)}}
}

func (p *DMA_Stream_Periph) MINC() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(MINC)}}
}

func (p *DMA_Stream_Periph) PINC() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PINC)}}
}

func (p *DMA_Stream_Periph) CIRC() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(CIRC)}}
}

func (p *DMA_Stream_Periph) DIR() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DIR)}}
}

func (p *DMA_Stream_Periph) PFCTRL() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PFCTRL)}}
}

func (p *DMA_Stream_Periph) TCIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TCIE)}}
}

func (p *DMA_Stream_Periph) HTIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(HTIE)}}
}

func (p *DMA_Stream_Periph) TEIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(TEIE)}}
}

func (p *DMA_Stream_Periph) DMEIE() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(DMEIE)}}
}

func (p *DMA_Stream_Periph) EN() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(EN)}}
}

type NDTR uint32

func (b NDTR) Field(mask NDTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask NDTR) J(v int) NDTR {
	return NDTR(bits.MakeField32(v, uint32(mask)))
}

type RNDTR struct{ mmio.U32 }

func (r *RNDTR) Bits(mask NDTR) NDTR    { return NDTR(r.U32.Bits(uint32(mask))) }
func (r *RNDTR) StoreBits(mask, b NDTR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RNDTR) SetBits(mask NDTR)      { r.U32.SetBits(uint32(mask)) }
func (r *RNDTR) ClearBits(mask NDTR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RNDTR) Load() NDTR             { return NDTR(r.U32.Load()) }
func (r *RNDTR) Store(b NDTR)           { r.U32.Store(uint32(b)) }

func (r *RNDTR) AtomicStoreBits(mask, b NDTR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RNDTR) AtomicSetBits(mask NDTR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RNDTR) AtomicClearBits(mask NDTR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMNDTR struct{ mmio.UM32 }

func (rm RMNDTR) Load() NDTR   { return NDTR(rm.UM32.Load()) }
func (rm RMNDTR) Store(b NDTR) { rm.UM32.Store(uint32(b)) }

type PAR uint32

func (b PAR) Field(mask PAR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PAR) J(v int) PAR {
	return PAR(bits.MakeField32(v, uint32(mask)))
}

type RPAR struct{ mmio.U32 }

func (r *RPAR) Bits(mask PAR) PAR     { return PAR(r.U32.Bits(uint32(mask))) }
func (r *RPAR) StoreBits(mask, b PAR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPAR) SetBits(mask PAR)      { r.U32.SetBits(uint32(mask)) }
func (r *RPAR) ClearBits(mask PAR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPAR) Load() PAR             { return PAR(r.U32.Load()) }
func (r *RPAR) Store(b PAR)           { r.U32.Store(uint32(b)) }

func (r *RPAR) AtomicStoreBits(mask, b PAR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPAR) AtomicSetBits(mask PAR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPAR) AtomicClearBits(mask PAR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPAR struct{ mmio.UM32 }

func (rm RMPAR) Load() PAR   { return PAR(rm.UM32.Load()) }
func (rm RMPAR) Store(b PAR) { rm.UM32.Store(uint32(b)) }

type M0AR uint32

func (b M0AR) Field(mask M0AR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask M0AR) J(v int) M0AR {
	return M0AR(bits.MakeField32(v, uint32(mask)))
}

type RM0AR struct{ mmio.U32 }

func (r *RM0AR) Bits(mask M0AR) M0AR    { return M0AR(r.U32.Bits(uint32(mask))) }
func (r *RM0AR) StoreBits(mask, b M0AR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RM0AR) SetBits(mask M0AR)      { r.U32.SetBits(uint32(mask)) }
func (r *RM0AR) ClearBits(mask M0AR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RM0AR) Load() M0AR             { return M0AR(r.U32.Load()) }
func (r *RM0AR) Store(b M0AR)           { r.U32.Store(uint32(b)) }

func (r *RM0AR) AtomicStoreBits(mask, b M0AR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RM0AR) AtomicSetBits(mask M0AR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RM0AR) AtomicClearBits(mask M0AR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMM0AR struct{ mmio.UM32 }

func (rm RMM0AR) Load() M0AR   { return M0AR(rm.UM32.Load()) }
func (rm RMM0AR) Store(b M0AR) { rm.UM32.Store(uint32(b)) }

type M1AR uint32

func (b M1AR) Field(mask M1AR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask M1AR) J(v int) M1AR {
	return M1AR(bits.MakeField32(v, uint32(mask)))
}

type RM1AR struct{ mmio.U32 }

func (r *RM1AR) Bits(mask M1AR) M1AR    { return M1AR(r.U32.Bits(uint32(mask))) }
func (r *RM1AR) StoreBits(mask, b M1AR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RM1AR) SetBits(mask M1AR)      { r.U32.SetBits(uint32(mask)) }
func (r *RM1AR) ClearBits(mask M1AR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RM1AR) Load() M1AR             { return M1AR(r.U32.Load()) }
func (r *RM1AR) Store(b M1AR)           { r.U32.Store(uint32(b)) }

func (r *RM1AR) AtomicStoreBits(mask, b M1AR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RM1AR) AtomicSetBits(mask M1AR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RM1AR) AtomicClearBits(mask M1AR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMM1AR struct{ mmio.UM32 }

func (rm RMM1AR) Load() M1AR   { return M1AR(rm.UM32.Load()) }
func (rm RMM1AR) Store(b M1AR) { rm.UM32.Store(uint32(b)) }

type FCR uint32

func (b FCR) Field(mask FCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FCR) J(v int) FCR {
	return FCR(bits.MakeField32(v, uint32(mask)))
}

type RFCR struct{ mmio.U32 }

func (r *RFCR) Bits(mask FCR) FCR     { return FCR(r.U32.Bits(uint32(mask))) }
func (r *RFCR) StoreBits(mask, b FCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RFCR) SetBits(mask FCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RFCR) ClearBits(mask FCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RFCR) Load() FCR             { return FCR(r.U32.Load()) }
func (r *RFCR) Store(b FCR)           { r.U32.Store(uint32(b)) }

func (r *RFCR) AtomicStoreBits(mask, b FCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RFCR) AtomicSetBits(mask FCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RFCR) AtomicClearBits(mask FCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMFCR struct{ mmio.UM32 }

func (rm RMFCR) Load() FCR   { return FCR(rm.UM32.Load()) }
func (rm RMFCR) Store(b FCR) { rm.UM32.Store(uint32(b)) }

func (p *DMA_Stream_Periph) FEIE() RMFCR {
	return RMFCR{mmio.UM32{&p.FCR.U32, uint32(FEIE)}}
}

func (p *DMA_Stream_Periph) FS() RMFCR {
	return RMFCR{mmio.UM32{&p.FCR.U32, uint32(FS)}}
}

func (p *DMA_Stream_Periph) DMDIS() RMFCR {
	return RMFCR{mmio.UM32{&p.FCR.U32, uint32(DMDIS)}}
}

func (p *DMA_Stream_Periph) FTH() RMFCR {
	return RMFCR{mmio.UM32{&p.FCR.U32, uint32(FTH)}}
}
