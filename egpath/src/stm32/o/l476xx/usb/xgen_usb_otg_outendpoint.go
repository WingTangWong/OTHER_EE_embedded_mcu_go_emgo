package usb

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l476xx/mmap"
)

type USB_OTG_OUTEndpoint_Periph struct {
	DOEPCTL  RDOEPCTL
	_        uint32
	DOEPINT  RDOEPINT
	_        uint32
	DOEPTSIZ RDOEPTSIZ
	DOEPDMA  RDOEPDMA
}

func (p *USB_OTG_OUTEndpoint_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type DOEPCTL uint32

func (b DOEPCTL) Field(mask DOEPCTL) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DOEPCTL) J(v int) DOEPCTL {
	return DOEPCTL(bits.MakeField32(v, uint32(mask)))
}

type RDOEPCTL struct{ mmio.U32 }

func (r *RDOEPCTL) Bits(mask DOEPCTL) DOEPCTL { return DOEPCTL(r.U32.Bits(uint32(mask))) }
func (r *RDOEPCTL) StoreBits(mask, b DOEPCTL) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDOEPCTL) SetBits(mask DOEPCTL)      { r.U32.SetBits(uint32(mask)) }
func (r *RDOEPCTL) ClearBits(mask DOEPCTL)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDOEPCTL) Load() DOEPCTL             { return DOEPCTL(r.U32.Load()) }
func (r *RDOEPCTL) Store(b DOEPCTL)           { r.U32.Store(uint32(b)) }

func (r *RDOEPCTL) AtomicStoreBits(mask, b DOEPCTL) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDOEPCTL) AtomicSetBits(mask DOEPCTL)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDOEPCTL) AtomicClearBits(mask DOEPCTL)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDOEPCTL struct{ mmio.UM32 }

func (rm RMDOEPCTL) Load() DOEPCTL   { return DOEPCTL(rm.UM32.Load()) }
func (rm RMDOEPCTL) Store(b DOEPCTL) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_OUTEndpoint_Periph) MPSIZ() RMDOEPCTL {
	return RMDOEPCTL{mmio.UM32{&p.DOEPCTL.U32, uint32(MPSIZ)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) USBAEP() RMDOEPCTL {
	return RMDOEPCTL{mmio.UM32{&p.DOEPCTL.U32, uint32(USBAEP)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) NAKSTS() RMDOEPCTL {
	return RMDOEPCTL{mmio.UM32{&p.DOEPCTL.U32, uint32(NAKSTS)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) SD0PID_SEVNFRM() RMDOEPCTL {
	return RMDOEPCTL{mmio.UM32{&p.DOEPCTL.U32, uint32(SD0PID_SEVNFRM)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) SODDFRM() RMDOEPCTL {
	return RMDOEPCTL{mmio.UM32{&p.DOEPCTL.U32, uint32(SODDFRM)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) EPTYP() RMDOEPCTL {
	return RMDOEPCTL{mmio.UM32{&p.DOEPCTL.U32, uint32(EPTYP)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) SNPM() RMDOEPCTL {
	return RMDOEPCTL{mmio.UM32{&p.DOEPCTL.U32, uint32(SNPM)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) STALL() RMDOEPCTL {
	return RMDOEPCTL{mmio.UM32{&p.DOEPCTL.U32, uint32(STALL)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) CNAK() RMDOEPCTL {
	return RMDOEPCTL{mmio.UM32{&p.DOEPCTL.U32, uint32(CNAK)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) SNAK() RMDOEPCTL {
	return RMDOEPCTL{mmio.UM32{&p.DOEPCTL.U32, uint32(SNAK)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) EPDIS() RMDOEPCTL {
	return RMDOEPCTL{mmio.UM32{&p.DOEPCTL.U32, uint32(EPDIS)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) EPENA() RMDOEPCTL {
	return RMDOEPCTL{mmio.UM32{&p.DOEPCTL.U32, uint32(EPENA)}}
}

type DOEPINT uint32

func (b DOEPINT) Field(mask DOEPINT) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DOEPINT) J(v int) DOEPINT {
	return DOEPINT(bits.MakeField32(v, uint32(mask)))
}

type RDOEPINT struct{ mmio.U32 }

func (r *RDOEPINT) Bits(mask DOEPINT) DOEPINT { return DOEPINT(r.U32.Bits(uint32(mask))) }
func (r *RDOEPINT) StoreBits(mask, b DOEPINT) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDOEPINT) SetBits(mask DOEPINT)      { r.U32.SetBits(uint32(mask)) }
func (r *RDOEPINT) ClearBits(mask DOEPINT)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDOEPINT) Load() DOEPINT             { return DOEPINT(r.U32.Load()) }
func (r *RDOEPINT) Store(b DOEPINT)           { r.U32.Store(uint32(b)) }

func (r *RDOEPINT) AtomicStoreBits(mask, b DOEPINT) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDOEPINT) AtomicSetBits(mask DOEPINT)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDOEPINT) AtomicClearBits(mask DOEPINT)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDOEPINT struct{ mmio.UM32 }

func (rm RMDOEPINT) Load() DOEPINT   { return DOEPINT(rm.UM32.Load()) }
func (rm RMDOEPINT) Store(b DOEPINT) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_OUTEndpoint_Periph) XFRC() RMDOEPINT {
	return RMDOEPINT{mmio.UM32{&p.DOEPINT.U32, uint32(XFRC)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) EPDISD() RMDOEPINT {
	return RMDOEPINT{mmio.UM32{&p.DOEPINT.U32, uint32(EPDISD)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) STUP() RMDOEPINT {
	return RMDOEPINT{mmio.UM32{&p.DOEPINT.U32, uint32(STUP)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) OTEPDIS() RMDOEPINT {
	return RMDOEPINT{mmio.UM32{&p.DOEPINT.U32, uint32(OTEPDIS)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) B2BSTUP() RMDOEPINT {
	return RMDOEPINT{mmio.UM32{&p.DOEPINT.U32, uint32(B2BSTUP)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) NYET() RMDOEPINT {
	return RMDOEPINT{mmio.UM32{&p.DOEPINT.U32, uint32(NYET)}}
}

type DOEPTSIZ uint32

func (b DOEPTSIZ) Field(mask DOEPTSIZ) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DOEPTSIZ) J(v int) DOEPTSIZ {
	return DOEPTSIZ(bits.MakeField32(v, uint32(mask)))
}

type RDOEPTSIZ struct{ mmio.U32 }

func (r *RDOEPTSIZ) Bits(mask DOEPTSIZ) DOEPTSIZ { return DOEPTSIZ(r.U32.Bits(uint32(mask))) }
func (r *RDOEPTSIZ) StoreBits(mask, b DOEPTSIZ)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDOEPTSIZ) SetBits(mask DOEPTSIZ)       { r.U32.SetBits(uint32(mask)) }
func (r *RDOEPTSIZ) ClearBits(mask DOEPTSIZ)     { r.U32.ClearBits(uint32(mask)) }
func (r *RDOEPTSIZ) Load() DOEPTSIZ              { return DOEPTSIZ(r.U32.Load()) }
func (r *RDOEPTSIZ) Store(b DOEPTSIZ)            { r.U32.Store(uint32(b)) }

func (r *RDOEPTSIZ) AtomicStoreBits(mask, b DOEPTSIZ) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDOEPTSIZ) AtomicSetBits(mask DOEPTSIZ)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDOEPTSIZ) AtomicClearBits(mask DOEPTSIZ)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDOEPTSIZ struct{ mmio.UM32 }

func (rm RMDOEPTSIZ) Load() DOEPTSIZ   { return DOEPTSIZ(rm.UM32.Load()) }
func (rm RMDOEPTSIZ) Store(b DOEPTSIZ) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_OUTEndpoint_Periph) XFRSIZ() RMDOEPTSIZ {
	return RMDOEPTSIZ{mmio.UM32{&p.DOEPTSIZ.U32, uint32(XFRSIZ)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) PKTCNT() RMDOEPTSIZ {
	return RMDOEPTSIZ{mmio.UM32{&p.DOEPTSIZ.U32, uint32(PKTCNT)}}
}

func (p *USB_OTG_OUTEndpoint_Periph) STUPCNT() RMDOEPTSIZ {
	return RMDOEPTSIZ{mmio.UM32{&p.DOEPTSIZ.U32, uint32(STUPCNT)}}
}

type DOEPDMA uint32

func (b DOEPDMA) Field(mask DOEPDMA) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DOEPDMA) J(v int) DOEPDMA {
	return DOEPDMA(bits.MakeField32(v, uint32(mask)))
}

type RDOEPDMA struct{ mmio.U32 }

func (r *RDOEPDMA) Bits(mask DOEPDMA) DOEPDMA { return DOEPDMA(r.U32.Bits(uint32(mask))) }
func (r *RDOEPDMA) StoreBits(mask, b DOEPDMA) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RDOEPDMA) SetBits(mask DOEPDMA)      { r.U32.SetBits(uint32(mask)) }
func (r *RDOEPDMA) ClearBits(mask DOEPDMA)    { r.U32.ClearBits(uint32(mask)) }
func (r *RDOEPDMA) Load() DOEPDMA             { return DOEPDMA(r.U32.Load()) }
func (r *RDOEPDMA) Store(b DOEPDMA)           { r.U32.Store(uint32(b)) }

func (r *RDOEPDMA) AtomicStoreBits(mask, b DOEPDMA) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RDOEPDMA) AtomicSetBits(mask DOEPDMA)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RDOEPDMA) AtomicClearBits(mask DOEPDMA)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMDOEPDMA struct{ mmio.UM32 }

func (rm RMDOEPDMA) Load() DOEPDMA   { return DOEPDMA(rm.UM32.Load()) }
func (rm RMDOEPDMA) Store(b DOEPDMA) { rm.UM32.Store(uint32(b)) }
