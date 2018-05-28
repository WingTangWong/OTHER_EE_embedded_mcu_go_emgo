package usb

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f746xx/mmap"
)

type USB_OTG_HostChannel_Periph struct {
	HCCHAR   RHCCHAR
	HCSPLT   RHCSPLT
	HCINT    RHCINT
	HCINTMSK RHCINTMSK
	HCTSIZ   RHCTSIZ
	HCDMA    RHCDMA
}

func (p *USB_OTG_HostChannel_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type HCCHAR uint32

func (b HCCHAR) Field(mask HCCHAR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HCCHAR) J(v int) HCCHAR {
	return HCCHAR(bits.MakeField32(v, uint32(mask)))
}

type RHCCHAR struct{ mmio.U32 }

func (r *RHCCHAR) Bits(mask HCCHAR) HCCHAR  { return HCCHAR(r.U32.Bits(uint32(mask))) }
func (r *RHCCHAR) StoreBits(mask, b HCCHAR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RHCCHAR) SetBits(mask HCCHAR)      { r.U32.SetBits(uint32(mask)) }
func (r *RHCCHAR) ClearBits(mask HCCHAR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RHCCHAR) Load() HCCHAR             { return HCCHAR(r.U32.Load()) }
func (r *RHCCHAR) Store(b HCCHAR)           { r.U32.Store(uint32(b)) }

func (r *RHCCHAR) AtomicStoreBits(mask, b HCCHAR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RHCCHAR) AtomicSetBits(mask HCCHAR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RHCCHAR) AtomicClearBits(mask HCCHAR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMHCCHAR struct{ mmio.UM32 }

func (rm RMHCCHAR) Load() HCCHAR   { return HCCHAR(rm.UM32.Load()) }
func (rm RMHCCHAR) Store(b HCCHAR) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_HostChannel_Periph) MPSIZ() RMHCCHAR {
	return RMHCCHAR{mmio.UM32{&p.HCCHAR.U32, uint32(MPSIZ)}}
}

func (p *USB_OTG_HostChannel_Periph) EPNUM() RMHCCHAR {
	return RMHCCHAR{mmio.UM32{&p.HCCHAR.U32, uint32(EPNUM)}}
}

func (p *USB_OTG_HostChannel_Periph) EPDIR() RMHCCHAR {
	return RMHCCHAR{mmio.UM32{&p.HCCHAR.U32, uint32(EPDIR)}}
}

func (p *USB_OTG_HostChannel_Periph) LSDEV() RMHCCHAR {
	return RMHCCHAR{mmio.UM32{&p.HCCHAR.U32, uint32(LSDEV)}}
}

func (p *USB_OTG_HostChannel_Periph) EPTYP() RMHCCHAR {
	return RMHCCHAR{mmio.UM32{&p.HCCHAR.U32, uint32(EPTYP)}}
}

func (p *USB_OTG_HostChannel_Periph) MC() RMHCCHAR {
	return RMHCCHAR{mmio.UM32{&p.HCCHAR.U32, uint32(MC)}}
}

func (p *USB_OTG_HostChannel_Periph) DAD() RMHCCHAR {
	return RMHCCHAR{mmio.UM32{&p.HCCHAR.U32, uint32(DAD)}}
}

func (p *USB_OTG_HostChannel_Periph) ODDFRM() RMHCCHAR {
	return RMHCCHAR{mmio.UM32{&p.HCCHAR.U32, uint32(ODDFRM)}}
}

func (p *USB_OTG_HostChannel_Periph) CHDIS() RMHCCHAR {
	return RMHCCHAR{mmio.UM32{&p.HCCHAR.U32, uint32(CHDIS)}}
}

func (p *USB_OTG_HostChannel_Periph) CHENA() RMHCCHAR {
	return RMHCCHAR{mmio.UM32{&p.HCCHAR.U32, uint32(CHENA)}}
}

type HCSPLT uint32

func (b HCSPLT) Field(mask HCSPLT) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HCSPLT) J(v int) HCSPLT {
	return HCSPLT(bits.MakeField32(v, uint32(mask)))
}

type RHCSPLT struct{ mmio.U32 }

func (r *RHCSPLT) Bits(mask HCSPLT) HCSPLT  { return HCSPLT(r.U32.Bits(uint32(mask))) }
func (r *RHCSPLT) StoreBits(mask, b HCSPLT) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RHCSPLT) SetBits(mask HCSPLT)      { r.U32.SetBits(uint32(mask)) }
func (r *RHCSPLT) ClearBits(mask HCSPLT)    { r.U32.ClearBits(uint32(mask)) }
func (r *RHCSPLT) Load() HCSPLT             { return HCSPLT(r.U32.Load()) }
func (r *RHCSPLT) Store(b HCSPLT)           { r.U32.Store(uint32(b)) }

func (r *RHCSPLT) AtomicStoreBits(mask, b HCSPLT) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RHCSPLT) AtomicSetBits(mask HCSPLT)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RHCSPLT) AtomicClearBits(mask HCSPLT)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMHCSPLT struct{ mmio.UM32 }

func (rm RMHCSPLT) Load() HCSPLT   { return HCSPLT(rm.UM32.Load()) }
func (rm RMHCSPLT) Store(b HCSPLT) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_HostChannel_Periph) PRTADDR() RMHCSPLT {
	return RMHCSPLT{mmio.UM32{&p.HCSPLT.U32, uint32(PRTADDR)}}
}

func (p *USB_OTG_HostChannel_Periph) HUBADDR() RMHCSPLT {
	return RMHCSPLT{mmio.UM32{&p.HCSPLT.U32, uint32(HUBADDR)}}
}

func (p *USB_OTG_HostChannel_Periph) XACTPOS() RMHCSPLT {
	return RMHCSPLT{mmio.UM32{&p.HCSPLT.U32, uint32(XACTPOS)}}
}

func (p *USB_OTG_HostChannel_Periph) COMPLSPLT() RMHCSPLT {
	return RMHCSPLT{mmio.UM32{&p.HCSPLT.U32, uint32(COMPLSPLT)}}
}

func (p *USB_OTG_HostChannel_Periph) SPLITEN() RMHCSPLT {
	return RMHCSPLT{mmio.UM32{&p.HCSPLT.U32, uint32(SPLITEN)}}
}

type HCINT uint32

func (b HCINT) Field(mask HCINT) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HCINT) J(v int) HCINT {
	return HCINT(bits.MakeField32(v, uint32(mask)))
}

type RHCINT struct{ mmio.U32 }

func (r *RHCINT) Bits(mask HCINT) HCINT   { return HCINT(r.U32.Bits(uint32(mask))) }
func (r *RHCINT) StoreBits(mask, b HCINT) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RHCINT) SetBits(mask HCINT)      { r.U32.SetBits(uint32(mask)) }
func (r *RHCINT) ClearBits(mask HCINT)    { r.U32.ClearBits(uint32(mask)) }
func (r *RHCINT) Load() HCINT             { return HCINT(r.U32.Load()) }
func (r *RHCINT) Store(b HCINT)           { r.U32.Store(uint32(b)) }

func (r *RHCINT) AtomicStoreBits(mask, b HCINT) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RHCINT) AtomicSetBits(mask HCINT)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RHCINT) AtomicClearBits(mask HCINT)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMHCINT struct{ mmio.UM32 }

func (rm RMHCINT) Load() HCINT   { return HCINT(rm.UM32.Load()) }
func (rm RMHCINT) Store(b HCINT) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_HostChannel_Periph) XFRC() RMHCINT {
	return RMHCINT{mmio.UM32{&p.HCINT.U32, uint32(XFRC)}}
}

func (p *USB_OTG_HostChannel_Periph) CHH() RMHCINT {
	return RMHCINT{mmio.UM32{&p.HCINT.U32, uint32(CHH)}}
}

func (p *USB_OTG_HostChannel_Periph) AHBERR() RMHCINT {
	return RMHCINT{mmio.UM32{&p.HCINT.U32, uint32(AHBERR)}}
}

func (p *USB_OTG_HostChannel_Periph) STALL() RMHCINT {
	return RMHCINT{mmio.UM32{&p.HCINT.U32, uint32(STALL)}}
}

func (p *USB_OTG_HostChannel_Periph) NAK() RMHCINT {
	return RMHCINT{mmio.UM32{&p.HCINT.U32, uint32(NAK)}}
}

func (p *USB_OTG_HostChannel_Periph) ACK() RMHCINT {
	return RMHCINT{mmio.UM32{&p.HCINT.U32, uint32(ACK)}}
}

func (p *USB_OTG_HostChannel_Periph) NYET() RMHCINT {
	return RMHCINT{mmio.UM32{&p.HCINT.U32, uint32(NYET)}}
}

func (p *USB_OTG_HostChannel_Periph) TXERR() RMHCINT {
	return RMHCINT{mmio.UM32{&p.HCINT.U32, uint32(TXERR)}}
}

func (p *USB_OTG_HostChannel_Periph) BBERR() RMHCINT {
	return RMHCINT{mmio.UM32{&p.HCINT.U32, uint32(BBERR)}}
}

func (p *USB_OTG_HostChannel_Periph) FRMOR() RMHCINT {
	return RMHCINT{mmio.UM32{&p.HCINT.U32, uint32(FRMOR)}}
}

func (p *USB_OTG_HostChannel_Periph) DTERR() RMHCINT {
	return RMHCINT{mmio.UM32{&p.HCINT.U32, uint32(DTERR)}}
}

type HCINTMSK uint32

func (b HCINTMSK) Field(mask HCINTMSK) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HCINTMSK) J(v int) HCINTMSK {
	return HCINTMSK(bits.MakeField32(v, uint32(mask)))
}

type RHCINTMSK struct{ mmio.U32 }

func (r *RHCINTMSK) Bits(mask HCINTMSK) HCINTMSK { return HCINTMSK(r.U32.Bits(uint32(mask))) }
func (r *RHCINTMSK) StoreBits(mask, b HCINTMSK)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RHCINTMSK) SetBits(mask HCINTMSK)       { r.U32.SetBits(uint32(mask)) }
func (r *RHCINTMSK) ClearBits(mask HCINTMSK)     { r.U32.ClearBits(uint32(mask)) }
func (r *RHCINTMSK) Load() HCINTMSK              { return HCINTMSK(r.U32.Load()) }
func (r *RHCINTMSK) Store(b HCINTMSK)            { r.U32.Store(uint32(b)) }

func (r *RHCINTMSK) AtomicStoreBits(mask, b HCINTMSK) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RHCINTMSK) AtomicSetBits(mask HCINTMSK)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RHCINTMSK) AtomicClearBits(mask HCINTMSK)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMHCINTMSK struct{ mmio.UM32 }

func (rm RMHCINTMSK) Load() HCINTMSK   { return HCINTMSK(rm.UM32.Load()) }
func (rm RMHCINTMSK) Store(b HCINTMSK) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_HostChannel_Periph) XFRCM() RMHCINTMSK {
	return RMHCINTMSK{mmio.UM32{&p.HCINTMSK.U32, uint32(XFRCM)}}
}

func (p *USB_OTG_HostChannel_Periph) CHHM() RMHCINTMSK {
	return RMHCINTMSK{mmio.UM32{&p.HCINTMSK.U32, uint32(CHHM)}}
}

func (p *USB_OTG_HostChannel_Periph) AHBERR() RMHCINTMSK {
	return RMHCINTMSK{mmio.UM32{&p.HCINTMSK.U32, uint32(AHBERR)}}
}

func (p *USB_OTG_HostChannel_Periph) STALLM() RMHCINTMSK {
	return RMHCINTMSK{mmio.UM32{&p.HCINTMSK.U32, uint32(STALLM)}}
}

func (p *USB_OTG_HostChannel_Periph) NAKM() RMHCINTMSK {
	return RMHCINTMSK{mmio.UM32{&p.HCINTMSK.U32, uint32(NAKM)}}
}

func (p *USB_OTG_HostChannel_Periph) ACKM() RMHCINTMSK {
	return RMHCINTMSK{mmio.UM32{&p.HCINTMSK.U32, uint32(ACKM)}}
}

func (p *USB_OTG_HostChannel_Periph) NYET() RMHCINTMSK {
	return RMHCINTMSK{mmio.UM32{&p.HCINTMSK.U32, uint32(NYET)}}
}

func (p *USB_OTG_HostChannel_Periph) TXERRM() RMHCINTMSK {
	return RMHCINTMSK{mmio.UM32{&p.HCINTMSK.U32, uint32(TXERRM)}}
}

func (p *USB_OTG_HostChannel_Periph) BBERRM() RMHCINTMSK {
	return RMHCINTMSK{mmio.UM32{&p.HCINTMSK.U32, uint32(BBERRM)}}
}

func (p *USB_OTG_HostChannel_Periph) FRMORM() RMHCINTMSK {
	return RMHCINTMSK{mmio.UM32{&p.HCINTMSK.U32, uint32(FRMORM)}}
}

func (p *USB_OTG_HostChannel_Periph) DTERRM() RMHCINTMSK {
	return RMHCINTMSK{mmio.UM32{&p.HCINTMSK.U32, uint32(DTERRM)}}
}

type HCTSIZ uint32

func (b HCTSIZ) Field(mask HCTSIZ) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HCTSIZ) J(v int) HCTSIZ {
	return HCTSIZ(bits.MakeField32(v, uint32(mask)))
}

type RHCTSIZ struct{ mmio.U32 }

func (r *RHCTSIZ) Bits(mask HCTSIZ) HCTSIZ  { return HCTSIZ(r.U32.Bits(uint32(mask))) }
func (r *RHCTSIZ) StoreBits(mask, b HCTSIZ) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RHCTSIZ) SetBits(mask HCTSIZ)      { r.U32.SetBits(uint32(mask)) }
func (r *RHCTSIZ) ClearBits(mask HCTSIZ)    { r.U32.ClearBits(uint32(mask)) }
func (r *RHCTSIZ) Load() HCTSIZ             { return HCTSIZ(r.U32.Load()) }
func (r *RHCTSIZ) Store(b HCTSIZ)           { r.U32.Store(uint32(b)) }

func (r *RHCTSIZ) AtomicStoreBits(mask, b HCTSIZ) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RHCTSIZ) AtomicSetBits(mask HCTSIZ)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RHCTSIZ) AtomicClearBits(mask HCTSIZ)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMHCTSIZ struct{ mmio.UM32 }

func (rm RMHCTSIZ) Load() HCTSIZ   { return HCTSIZ(rm.UM32.Load()) }
func (rm RMHCTSIZ) Store(b HCTSIZ) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_HostChannel_Periph) XFRSIZ() RMHCTSIZ {
	return RMHCTSIZ{mmio.UM32{&p.HCTSIZ.U32, uint32(XFRSIZ)}}
}

func (p *USB_OTG_HostChannel_Periph) PKTCNT() RMHCTSIZ {
	return RMHCTSIZ{mmio.UM32{&p.HCTSIZ.U32, uint32(PKTCNT)}}
}

func (p *USB_OTG_HostChannel_Periph) DOPING() RMHCTSIZ {
	return RMHCTSIZ{mmio.UM32{&p.HCTSIZ.U32, uint32(DOPING)}}
}

func (p *USB_OTG_HostChannel_Periph) DPID() RMHCTSIZ {
	return RMHCTSIZ{mmio.UM32{&p.HCTSIZ.U32, uint32(DPID)}}
}

type HCDMA uint32

func (b HCDMA) Field(mask HCDMA) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HCDMA) J(v int) HCDMA {
	return HCDMA(bits.MakeField32(v, uint32(mask)))
}

type RHCDMA struct{ mmio.U32 }

func (r *RHCDMA) Bits(mask HCDMA) HCDMA   { return HCDMA(r.U32.Bits(uint32(mask))) }
func (r *RHCDMA) StoreBits(mask, b HCDMA) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RHCDMA) SetBits(mask HCDMA)      { r.U32.SetBits(uint32(mask)) }
func (r *RHCDMA) ClearBits(mask HCDMA)    { r.U32.ClearBits(uint32(mask)) }
func (r *RHCDMA) Load() HCDMA             { return HCDMA(r.U32.Load()) }
func (r *RHCDMA) Store(b HCDMA)           { r.U32.Store(uint32(b)) }

func (r *RHCDMA) AtomicStoreBits(mask, b HCDMA) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RHCDMA) AtomicSetBits(mask HCDMA)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RHCDMA) AtomicClearBits(mask HCDMA)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMHCDMA struct{ mmio.UM32 }

func (rm RMHCDMA) Load() HCDMA   { return HCDMA(rm.UM32.Load()) }
func (rm RMHCDMA) Store(b HCDMA) { rm.UM32.Store(uint32(b)) }

func (p *USB_OTG_HostChannel_Periph) DMAADDR() RMHCDMA {
	return RMHCDMA{mmio.UM32{&p.HCDMA.U32, uint32(DMAADDR)}}
}
