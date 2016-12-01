package usb

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"mmio"
	"unsafe"

	"stm32/o/f746xx/mmap"
)

type USB_OTG_Global_Periph struct {
	GOTGCTL            GOTGCTL
	GOTGINT            GOTGINT
	GAHBCFG            GAHBCFG
	GUSBCFG            GUSBCFG
	GRSTCTL            GRSTCTL
	GINTSTS            GINTSTS
	GINTMSK            GINTMSK
	GRXSTSR            GRXSTSR
	GRXSTSP            GRXSTSP
	GRXFSIZ            GRXFSIZ
	DIEPTXF0_HNPTXFSIZ DIEPTXF0_HNPTXFSIZ
	HNPTXSTS           HNPTXSTS
	_                  [2]uint32
	GCCFG              GCCFG
	CID                CID
	_                  [3]uint32
	GHWCFG3            GHWCFG3
	_                  uint32
	GLPMCFG            GLPMCFG
	GPWRDN             GPWRDN
	GDFIFOCFG          GDFIFOCFG
	GADPCTL            GADPCTL
	_                  [39]uint32
	HPTXFSIZ           HPTXFSIZ
	DIEPTXF            [15]DIEPTXF
}

func (p *USB_OTG_Global_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

type GOTGCTL_Bits uint32

type GOTGCTL struct{ mmio.U32 }

func (r *GOTGCTL) Bits(mask GOTGCTL_Bits) GOTGCTL_Bits { return GOTGCTL_Bits(r.U32.Bits(uint32(mask))) }
func (r *GOTGCTL) StoreBits(mask, b GOTGCTL_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *GOTGCTL) SetBits(mask GOTGCTL_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *GOTGCTL) ClearBits(mask GOTGCTL_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *GOTGCTL) Load() GOTGCTL_Bits                  { return GOTGCTL_Bits(r.U32.Load()) }
func (r *GOTGCTL) Store(b GOTGCTL_Bits)                { r.U32.Store(uint32(b)) }

type GOTGCTL_Mask struct{ mmio.UM32 }

func (rm GOTGCTL_Mask) Load() GOTGCTL_Bits   { return GOTGCTL_Bits(rm.UM32.Load()) }
func (rm GOTGCTL_Mask) Store(b GOTGCTL_Bits) { rm.UM32.Store(uint32(b)) }

type GOTGINT_Bits uint32

type GOTGINT struct{ mmio.U32 }

func (r *GOTGINT) Bits(mask GOTGINT_Bits) GOTGINT_Bits { return GOTGINT_Bits(r.U32.Bits(uint32(mask))) }
func (r *GOTGINT) StoreBits(mask, b GOTGINT_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *GOTGINT) SetBits(mask GOTGINT_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *GOTGINT) ClearBits(mask GOTGINT_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *GOTGINT) Load() GOTGINT_Bits                  { return GOTGINT_Bits(r.U32.Load()) }
func (r *GOTGINT) Store(b GOTGINT_Bits)                { r.U32.Store(uint32(b)) }

type GOTGINT_Mask struct{ mmio.UM32 }

func (rm GOTGINT_Mask) Load() GOTGINT_Bits   { return GOTGINT_Bits(rm.UM32.Load()) }
func (rm GOTGINT_Mask) Store(b GOTGINT_Bits) { rm.UM32.Store(uint32(b)) }

type GAHBCFG_Bits uint32

type GAHBCFG struct{ mmio.U32 }

func (r *GAHBCFG) Bits(mask GAHBCFG_Bits) GAHBCFG_Bits { return GAHBCFG_Bits(r.U32.Bits(uint32(mask))) }
func (r *GAHBCFG) StoreBits(mask, b GAHBCFG_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *GAHBCFG) SetBits(mask GAHBCFG_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *GAHBCFG) ClearBits(mask GAHBCFG_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *GAHBCFG) Load() GAHBCFG_Bits                  { return GAHBCFG_Bits(r.U32.Load()) }
func (r *GAHBCFG) Store(b GAHBCFG_Bits)                { r.U32.Store(uint32(b)) }

type GAHBCFG_Mask struct{ mmio.UM32 }

func (rm GAHBCFG_Mask) Load() GAHBCFG_Bits   { return GAHBCFG_Bits(rm.UM32.Load()) }
func (rm GAHBCFG_Mask) Store(b GAHBCFG_Bits) { rm.UM32.Store(uint32(b)) }

type GUSBCFG_Bits uint32

type GUSBCFG struct{ mmio.U32 }

func (r *GUSBCFG) Bits(mask GUSBCFG_Bits) GUSBCFG_Bits { return GUSBCFG_Bits(r.U32.Bits(uint32(mask))) }
func (r *GUSBCFG) StoreBits(mask, b GUSBCFG_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *GUSBCFG) SetBits(mask GUSBCFG_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *GUSBCFG) ClearBits(mask GUSBCFG_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *GUSBCFG) Load() GUSBCFG_Bits                  { return GUSBCFG_Bits(r.U32.Load()) }
func (r *GUSBCFG) Store(b GUSBCFG_Bits)                { r.U32.Store(uint32(b)) }

type GUSBCFG_Mask struct{ mmio.UM32 }

func (rm GUSBCFG_Mask) Load() GUSBCFG_Bits   { return GUSBCFG_Bits(rm.UM32.Load()) }
func (rm GUSBCFG_Mask) Store(b GUSBCFG_Bits) { rm.UM32.Store(uint32(b)) }

type GRSTCTL_Bits uint32

type GRSTCTL struct{ mmio.U32 }

func (r *GRSTCTL) Bits(mask GRSTCTL_Bits) GRSTCTL_Bits { return GRSTCTL_Bits(r.U32.Bits(uint32(mask))) }
func (r *GRSTCTL) StoreBits(mask, b GRSTCTL_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *GRSTCTL) SetBits(mask GRSTCTL_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *GRSTCTL) ClearBits(mask GRSTCTL_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *GRSTCTL) Load() GRSTCTL_Bits                  { return GRSTCTL_Bits(r.U32.Load()) }
func (r *GRSTCTL) Store(b GRSTCTL_Bits)                { r.U32.Store(uint32(b)) }

type GRSTCTL_Mask struct{ mmio.UM32 }

func (rm GRSTCTL_Mask) Load() GRSTCTL_Bits   { return GRSTCTL_Bits(rm.UM32.Load()) }
func (rm GRSTCTL_Mask) Store(b GRSTCTL_Bits) { rm.UM32.Store(uint32(b)) }

type GINTSTS_Bits uint32

type GINTSTS struct{ mmio.U32 }

func (r *GINTSTS) Bits(mask GINTSTS_Bits) GINTSTS_Bits { return GINTSTS_Bits(r.U32.Bits(uint32(mask))) }
func (r *GINTSTS) StoreBits(mask, b GINTSTS_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *GINTSTS) SetBits(mask GINTSTS_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *GINTSTS) ClearBits(mask GINTSTS_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *GINTSTS) Load() GINTSTS_Bits                  { return GINTSTS_Bits(r.U32.Load()) }
func (r *GINTSTS) Store(b GINTSTS_Bits)                { r.U32.Store(uint32(b)) }

type GINTSTS_Mask struct{ mmio.UM32 }

func (rm GINTSTS_Mask) Load() GINTSTS_Bits   { return GINTSTS_Bits(rm.UM32.Load()) }
func (rm GINTSTS_Mask) Store(b GINTSTS_Bits) { rm.UM32.Store(uint32(b)) }

type GINTMSK_Bits uint32

type GINTMSK struct{ mmio.U32 }

func (r *GINTMSK) Bits(mask GINTMSK_Bits) GINTMSK_Bits { return GINTMSK_Bits(r.U32.Bits(uint32(mask))) }
func (r *GINTMSK) StoreBits(mask, b GINTMSK_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *GINTMSK) SetBits(mask GINTMSK_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *GINTMSK) ClearBits(mask GINTMSK_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *GINTMSK) Load() GINTMSK_Bits                  { return GINTMSK_Bits(r.U32.Load()) }
func (r *GINTMSK) Store(b GINTMSK_Bits)                { r.U32.Store(uint32(b)) }

type GINTMSK_Mask struct{ mmio.UM32 }

func (rm GINTMSK_Mask) Load() GINTMSK_Bits   { return GINTMSK_Bits(rm.UM32.Load()) }
func (rm GINTMSK_Mask) Store(b GINTMSK_Bits) { rm.UM32.Store(uint32(b)) }

type GRXSTSR_Bits uint32

type GRXSTSR struct{ mmio.U32 }

func (r *GRXSTSR) Bits(mask GRXSTSR_Bits) GRXSTSR_Bits { return GRXSTSR_Bits(r.U32.Bits(uint32(mask))) }
func (r *GRXSTSR) StoreBits(mask, b GRXSTSR_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *GRXSTSR) SetBits(mask GRXSTSR_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *GRXSTSR) ClearBits(mask GRXSTSR_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *GRXSTSR) Load() GRXSTSR_Bits                  { return GRXSTSR_Bits(r.U32.Load()) }
func (r *GRXSTSR) Store(b GRXSTSR_Bits)                { r.U32.Store(uint32(b)) }

type GRXSTSR_Mask struct{ mmio.UM32 }

func (rm GRXSTSR_Mask) Load() GRXSTSR_Bits   { return GRXSTSR_Bits(rm.UM32.Load()) }
func (rm GRXSTSR_Mask) Store(b GRXSTSR_Bits) { rm.UM32.Store(uint32(b)) }

type GRXSTSP_Bits uint32

type GRXSTSP struct{ mmio.U32 }

func (r *GRXSTSP) Bits(mask GRXSTSP_Bits) GRXSTSP_Bits { return GRXSTSP_Bits(r.U32.Bits(uint32(mask))) }
func (r *GRXSTSP) StoreBits(mask, b GRXSTSP_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *GRXSTSP) SetBits(mask GRXSTSP_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *GRXSTSP) ClearBits(mask GRXSTSP_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *GRXSTSP) Load() GRXSTSP_Bits                  { return GRXSTSP_Bits(r.U32.Load()) }
func (r *GRXSTSP) Store(b GRXSTSP_Bits)                { r.U32.Store(uint32(b)) }

type GRXSTSP_Mask struct{ mmio.UM32 }

func (rm GRXSTSP_Mask) Load() GRXSTSP_Bits   { return GRXSTSP_Bits(rm.UM32.Load()) }
func (rm GRXSTSP_Mask) Store(b GRXSTSP_Bits) { rm.UM32.Store(uint32(b)) }

type GRXFSIZ_Bits uint32

type GRXFSIZ struct{ mmio.U32 }

func (r *GRXFSIZ) Bits(mask GRXFSIZ_Bits) GRXFSIZ_Bits { return GRXFSIZ_Bits(r.U32.Bits(uint32(mask))) }
func (r *GRXFSIZ) StoreBits(mask, b GRXFSIZ_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *GRXFSIZ) SetBits(mask GRXFSIZ_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *GRXFSIZ) ClearBits(mask GRXFSIZ_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *GRXFSIZ) Load() GRXFSIZ_Bits                  { return GRXFSIZ_Bits(r.U32.Load()) }
func (r *GRXFSIZ) Store(b GRXFSIZ_Bits)                { r.U32.Store(uint32(b)) }

type GRXFSIZ_Mask struct{ mmio.UM32 }

func (rm GRXFSIZ_Mask) Load() GRXFSIZ_Bits   { return GRXFSIZ_Bits(rm.UM32.Load()) }
func (rm GRXFSIZ_Mask) Store(b GRXFSIZ_Bits) { rm.UM32.Store(uint32(b)) }

type DIEPTXF0_HNPTXFSIZ_Bits uint32

type DIEPTXF0_HNPTXFSIZ struct{ mmio.U32 }

func (r *DIEPTXF0_HNPTXFSIZ) Bits(mask DIEPTXF0_HNPTXFSIZ_Bits) DIEPTXF0_HNPTXFSIZ_Bits {
	return DIEPTXF0_HNPTXFSIZ_Bits(r.U32.Bits(uint32(mask)))
}
func (r *DIEPTXF0_HNPTXFSIZ) StoreBits(mask, b DIEPTXF0_HNPTXFSIZ_Bits) {
	r.U32.StoreBits(uint32(mask), uint32(b))
}
func (r *DIEPTXF0_HNPTXFSIZ) SetBits(mask DIEPTXF0_HNPTXFSIZ_Bits)   { r.U32.SetBits(uint32(mask)) }
func (r *DIEPTXF0_HNPTXFSIZ) ClearBits(mask DIEPTXF0_HNPTXFSIZ_Bits) { r.U32.ClearBits(uint32(mask)) }
func (r *DIEPTXF0_HNPTXFSIZ) Load() DIEPTXF0_HNPTXFSIZ_Bits {
	return DIEPTXF0_HNPTXFSIZ_Bits(r.U32.Load())
}
func (r *DIEPTXF0_HNPTXFSIZ) Store(b DIEPTXF0_HNPTXFSIZ_Bits) { r.U32.Store(uint32(b)) }

type DIEPTXF0_HNPTXFSIZ_Mask struct{ mmio.UM32 }

func (rm DIEPTXF0_HNPTXFSIZ_Mask) Load() DIEPTXF0_HNPTXFSIZ_Bits {
	return DIEPTXF0_HNPTXFSIZ_Bits(rm.UM32.Load())
}
func (rm DIEPTXF0_HNPTXFSIZ_Mask) Store(b DIEPTXF0_HNPTXFSIZ_Bits) { rm.UM32.Store(uint32(b)) }

type HNPTXSTS_Bits uint32

type HNPTXSTS struct{ mmio.U32 }

func (r *HNPTXSTS) Bits(mask HNPTXSTS_Bits) HNPTXSTS_Bits {
	return HNPTXSTS_Bits(r.U32.Bits(uint32(mask)))
}
func (r *HNPTXSTS) StoreBits(mask, b HNPTXSTS_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *HNPTXSTS) SetBits(mask HNPTXSTS_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *HNPTXSTS) ClearBits(mask HNPTXSTS_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *HNPTXSTS) Load() HNPTXSTS_Bits             { return HNPTXSTS_Bits(r.U32.Load()) }
func (r *HNPTXSTS) Store(b HNPTXSTS_Bits)           { r.U32.Store(uint32(b)) }

type HNPTXSTS_Mask struct{ mmio.UM32 }

func (rm HNPTXSTS_Mask) Load() HNPTXSTS_Bits   { return HNPTXSTS_Bits(rm.UM32.Load()) }
func (rm HNPTXSTS_Mask) Store(b HNPTXSTS_Bits) { rm.UM32.Store(uint32(b)) }

type GCCFG_Bits uint32

type GCCFG struct{ mmio.U32 }

func (r *GCCFG) Bits(mask GCCFG_Bits) GCCFG_Bits { return GCCFG_Bits(r.U32.Bits(uint32(mask))) }
func (r *GCCFG) StoreBits(mask, b GCCFG_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *GCCFG) SetBits(mask GCCFG_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *GCCFG) ClearBits(mask GCCFG_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *GCCFG) Load() GCCFG_Bits                { return GCCFG_Bits(r.U32.Load()) }
func (r *GCCFG) Store(b GCCFG_Bits)              { r.U32.Store(uint32(b)) }

type GCCFG_Mask struct{ mmio.UM32 }

func (rm GCCFG_Mask) Load() GCCFG_Bits   { return GCCFG_Bits(rm.UM32.Load()) }
func (rm GCCFG_Mask) Store(b GCCFG_Bits) { rm.UM32.Store(uint32(b)) }

type CID_Bits uint32

type CID struct{ mmio.U32 }

func (r *CID) Bits(mask CID_Bits) CID_Bits { return CID_Bits(r.U32.Bits(uint32(mask))) }
func (r *CID) StoreBits(mask, b CID_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CID) SetBits(mask CID_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CID) ClearBits(mask CID_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CID) Load() CID_Bits              { return CID_Bits(r.U32.Load()) }
func (r *CID) Store(b CID_Bits)            { r.U32.Store(uint32(b)) }

type CID_Mask struct{ mmio.UM32 }

func (rm CID_Mask) Load() CID_Bits   { return CID_Bits(rm.UM32.Load()) }
func (rm CID_Mask) Store(b CID_Bits) { rm.UM32.Store(uint32(b)) }

type GHWCFG3_Bits uint32

type GHWCFG3 struct{ mmio.U32 }

func (r *GHWCFG3) Bits(mask GHWCFG3_Bits) GHWCFG3_Bits { return GHWCFG3_Bits(r.U32.Bits(uint32(mask))) }
func (r *GHWCFG3) StoreBits(mask, b GHWCFG3_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *GHWCFG3) SetBits(mask GHWCFG3_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *GHWCFG3) ClearBits(mask GHWCFG3_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *GHWCFG3) Load() GHWCFG3_Bits                  { return GHWCFG3_Bits(r.U32.Load()) }
func (r *GHWCFG3) Store(b GHWCFG3_Bits)                { r.U32.Store(uint32(b)) }

type GHWCFG3_Mask struct{ mmio.UM32 }

func (rm GHWCFG3_Mask) Load() GHWCFG3_Bits   { return GHWCFG3_Bits(rm.UM32.Load()) }
func (rm GHWCFG3_Mask) Store(b GHWCFG3_Bits) { rm.UM32.Store(uint32(b)) }

type GLPMCFG_Bits uint32

type GLPMCFG struct{ mmio.U32 }

func (r *GLPMCFG) Bits(mask GLPMCFG_Bits) GLPMCFG_Bits { return GLPMCFG_Bits(r.U32.Bits(uint32(mask))) }
func (r *GLPMCFG) StoreBits(mask, b GLPMCFG_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *GLPMCFG) SetBits(mask GLPMCFG_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *GLPMCFG) ClearBits(mask GLPMCFG_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *GLPMCFG) Load() GLPMCFG_Bits                  { return GLPMCFG_Bits(r.U32.Load()) }
func (r *GLPMCFG) Store(b GLPMCFG_Bits)                { r.U32.Store(uint32(b)) }

type GLPMCFG_Mask struct{ mmio.UM32 }

func (rm GLPMCFG_Mask) Load() GLPMCFG_Bits   { return GLPMCFG_Bits(rm.UM32.Load()) }
func (rm GLPMCFG_Mask) Store(b GLPMCFG_Bits) { rm.UM32.Store(uint32(b)) }

type GPWRDN_Bits uint32

type GPWRDN struct{ mmio.U32 }

func (r *GPWRDN) Bits(mask GPWRDN_Bits) GPWRDN_Bits { return GPWRDN_Bits(r.U32.Bits(uint32(mask))) }
func (r *GPWRDN) StoreBits(mask, b GPWRDN_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *GPWRDN) SetBits(mask GPWRDN_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *GPWRDN) ClearBits(mask GPWRDN_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *GPWRDN) Load() GPWRDN_Bits                 { return GPWRDN_Bits(r.U32.Load()) }
func (r *GPWRDN) Store(b GPWRDN_Bits)               { r.U32.Store(uint32(b)) }

type GPWRDN_Mask struct{ mmio.UM32 }

func (rm GPWRDN_Mask) Load() GPWRDN_Bits   { return GPWRDN_Bits(rm.UM32.Load()) }
func (rm GPWRDN_Mask) Store(b GPWRDN_Bits) { rm.UM32.Store(uint32(b)) }

type GDFIFOCFG_Bits uint32

type GDFIFOCFG struct{ mmio.U32 }

func (r *GDFIFOCFG) Bits(mask GDFIFOCFG_Bits) GDFIFOCFG_Bits {
	return GDFIFOCFG_Bits(r.U32.Bits(uint32(mask)))
}
func (r *GDFIFOCFG) StoreBits(mask, b GDFIFOCFG_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *GDFIFOCFG) SetBits(mask GDFIFOCFG_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *GDFIFOCFG) ClearBits(mask GDFIFOCFG_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *GDFIFOCFG) Load() GDFIFOCFG_Bits             { return GDFIFOCFG_Bits(r.U32.Load()) }
func (r *GDFIFOCFG) Store(b GDFIFOCFG_Bits)           { r.U32.Store(uint32(b)) }

type GDFIFOCFG_Mask struct{ mmio.UM32 }

func (rm GDFIFOCFG_Mask) Load() GDFIFOCFG_Bits   { return GDFIFOCFG_Bits(rm.UM32.Load()) }
func (rm GDFIFOCFG_Mask) Store(b GDFIFOCFG_Bits) { rm.UM32.Store(uint32(b)) }

type GADPCTL_Bits uint32

type GADPCTL struct{ mmio.U32 }

func (r *GADPCTL) Bits(mask GADPCTL_Bits) GADPCTL_Bits { return GADPCTL_Bits(r.U32.Bits(uint32(mask))) }
func (r *GADPCTL) StoreBits(mask, b GADPCTL_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *GADPCTL) SetBits(mask GADPCTL_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *GADPCTL) ClearBits(mask GADPCTL_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *GADPCTL) Load() GADPCTL_Bits                  { return GADPCTL_Bits(r.U32.Load()) }
func (r *GADPCTL) Store(b GADPCTL_Bits)                { r.U32.Store(uint32(b)) }

type GADPCTL_Mask struct{ mmio.UM32 }

func (rm GADPCTL_Mask) Load() GADPCTL_Bits   { return GADPCTL_Bits(rm.UM32.Load()) }
func (rm GADPCTL_Mask) Store(b GADPCTL_Bits) { rm.UM32.Store(uint32(b)) }

type HPTXFSIZ_Bits uint32

type HPTXFSIZ struct{ mmio.U32 }

func (r *HPTXFSIZ) Bits(mask HPTXFSIZ_Bits) HPTXFSIZ_Bits {
	return HPTXFSIZ_Bits(r.U32.Bits(uint32(mask)))
}
func (r *HPTXFSIZ) StoreBits(mask, b HPTXFSIZ_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *HPTXFSIZ) SetBits(mask HPTXFSIZ_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *HPTXFSIZ) ClearBits(mask HPTXFSIZ_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *HPTXFSIZ) Load() HPTXFSIZ_Bits             { return HPTXFSIZ_Bits(r.U32.Load()) }
func (r *HPTXFSIZ) Store(b HPTXFSIZ_Bits)           { r.U32.Store(uint32(b)) }

type HPTXFSIZ_Mask struct{ mmio.UM32 }

func (rm HPTXFSIZ_Mask) Load() HPTXFSIZ_Bits   { return HPTXFSIZ_Bits(rm.UM32.Load()) }
func (rm HPTXFSIZ_Mask) Store(b HPTXFSIZ_Bits) { rm.UM32.Store(uint32(b)) }

type DIEPTXF_Bits uint32

type DIEPTXF struct{ mmio.U32 }

func (r *DIEPTXF) Bits(mask DIEPTXF_Bits) DIEPTXF_Bits { return DIEPTXF_Bits(r.U32.Bits(uint32(mask))) }
func (r *DIEPTXF) StoreBits(mask, b DIEPTXF_Bits)      { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DIEPTXF) SetBits(mask DIEPTXF_Bits)           { r.U32.SetBits(uint32(mask)) }
func (r *DIEPTXF) ClearBits(mask DIEPTXF_Bits)         { r.U32.ClearBits(uint32(mask)) }
func (r *DIEPTXF) Load() DIEPTXF_Bits                  { return DIEPTXF_Bits(r.U32.Load()) }
func (r *DIEPTXF) Store(b DIEPTXF_Bits)                { r.U32.Store(uint32(b)) }

type DIEPTXF_Mask struct{ mmio.UM32 }

func (rm DIEPTXF_Mask) Load() DIEPTXF_Bits   { return DIEPTXF_Bits(rm.UM32.Load()) }
func (rm DIEPTXF_Mask) Store(b DIEPTXF_Bits) { rm.UM32.Store(uint32(b)) }