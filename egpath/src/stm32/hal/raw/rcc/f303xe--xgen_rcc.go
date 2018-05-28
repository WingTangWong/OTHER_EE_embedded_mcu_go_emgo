// +build f303xe

package rcc

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f303xe/mmap"
)

type RCC_Periph struct {
	CR       RCR
	CFGR     RCFGR
	CIR      RCIR
	APB2RSTR RAPB2RSTR
	APB1RSTR RAPB1RSTR
	AHBENR   RAHBENR
	APB2ENR  RAPB2ENR
	APB1ENR  RAPB1ENR
	BDCR     RBDCR
	CSR      RCSR
	AHBRSTR  RAHBRSTR
	CFGR2    RCFGR2
	CFGR3    RCFGR3
}

func (p *RCC_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var RCC = (*RCC_Periph)(unsafe.Pointer(uintptr(mmap.RCC_BASE)))

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

func (p *RCC_Periph) HSION() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(HSION)}}
}

func (p *RCC_Periph) HSIRDY() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(HSIRDY)}}
}

func (p *RCC_Periph) HSITRIM() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(HSITRIM)}}
}

func (p *RCC_Periph) HSICAL() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(HSICAL)}}
}

func (p *RCC_Periph) HSEON() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(HSEON)}}
}

func (p *RCC_Periph) HSERDY() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(HSERDY)}}
}

func (p *RCC_Periph) HSEBYP() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(HSEBYP)}}
}

func (p *RCC_Periph) CSSON() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(CSSON)}}
}

func (p *RCC_Periph) PLLON() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PLLON)}}
}

func (p *RCC_Periph) PLLRDY() RMCR {
	return RMCR{mmio.UM32{&p.CR.U32, uint32(PLLRDY)}}
}

type CFGR uint32

func (b CFGR) Field(mask CFGR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CFGR) J(v int) CFGR {
	return CFGR(bits.MakeField32(v, uint32(mask)))
}

type RCFGR struct{ mmio.U32 }

func (r *RCFGR) Bits(mask CFGR) CFGR    { return CFGR(r.U32.Bits(uint32(mask))) }
func (r *RCFGR) StoreBits(mask, b CFGR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCFGR) SetBits(mask CFGR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCFGR) ClearBits(mask CFGR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCFGR) Load() CFGR             { return CFGR(r.U32.Load()) }
func (r *RCFGR) Store(b CFGR)           { r.U32.Store(uint32(b)) }

func (r *RCFGR) AtomicStoreBits(mask, b CFGR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCFGR) AtomicSetBits(mask CFGR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCFGR) AtomicClearBits(mask CFGR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCFGR struct{ mmio.UM32 }

func (rm RMCFGR) Load() CFGR   { return CFGR(rm.UM32.Load()) }
func (rm RMCFGR) Store(b CFGR) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) SW() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(SW)}}
}

func (p *RCC_Periph) SWS() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(SWS)}}
}

func (p *RCC_Periph) HPRE() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(HPRE)}}
}

func (p *RCC_Periph) PPRE1() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(PPRE1)}}
}

func (p *RCC_Periph) PPRE2() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(PPRE2)}}
}

func (p *RCC_Periph) PLLSRC() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(PLLSRC)}}
}

func (p *RCC_Periph) PLLXTPRE() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(PLLXTPRE)}}
}

func (p *RCC_Periph) PLLMUL() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(PLLMUL)}}
}

func (p *RCC_Periph) USBPRE() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(USBPRE)}}
}

func (p *RCC_Periph) I2SSRC() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(I2SSRC)}}
}

func (p *RCC_Periph) MCO() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(MCO)}}
}

func (p *RCC_Periph) MCOPRE() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(MCOPRE)}}
}

func (p *RCC_Periph) PLLNODIV() RMCFGR {
	return RMCFGR{mmio.UM32{&p.CFGR.U32, uint32(PLLNODIV)}}
}

type CIR uint32

func (b CIR) Field(mask CIR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CIR) J(v int) CIR {
	return CIR(bits.MakeField32(v, uint32(mask)))
}

type RCIR struct{ mmio.U32 }

func (r *RCIR) Bits(mask CIR) CIR     { return CIR(r.U32.Bits(uint32(mask))) }
func (r *RCIR) StoreBits(mask, b CIR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCIR) SetBits(mask CIR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCIR) ClearBits(mask CIR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCIR) Load() CIR             { return CIR(r.U32.Load()) }
func (r *RCIR) Store(b CIR)           { r.U32.Store(uint32(b)) }

func (r *RCIR) AtomicStoreBits(mask, b CIR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCIR) AtomicSetBits(mask CIR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCIR) AtomicClearBits(mask CIR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCIR struct{ mmio.UM32 }

func (rm RMCIR) Load() CIR   { return CIR(rm.UM32.Load()) }
func (rm RMCIR) Store(b CIR) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) LSIRDYF() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(LSIRDYF)}}
}

func (p *RCC_Periph) LSERDYF() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(LSERDYF)}}
}

func (p *RCC_Periph) HSIRDYF() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(HSIRDYF)}}
}

func (p *RCC_Periph) HSERDYF() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(HSERDYF)}}
}

func (p *RCC_Periph) PLLRDYF() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(PLLRDYF)}}
}

func (p *RCC_Periph) CSSF() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(CSSF)}}
}

func (p *RCC_Periph) LSIRDYIE() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(LSIRDYIE)}}
}

func (p *RCC_Periph) LSERDYIE() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(LSERDYIE)}}
}

func (p *RCC_Periph) HSIRDYIE() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(HSIRDYIE)}}
}

func (p *RCC_Periph) HSERDYIE() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(HSERDYIE)}}
}

func (p *RCC_Periph) PLLRDYIE() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(PLLRDYIE)}}
}

func (p *RCC_Periph) LSIRDYC() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(LSIRDYC)}}
}

func (p *RCC_Periph) LSERDYC() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(LSERDYC)}}
}

func (p *RCC_Periph) HSIRDYC() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(HSIRDYC)}}
}

func (p *RCC_Periph) HSERDYC() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(HSERDYC)}}
}

func (p *RCC_Periph) PLLRDYC() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(PLLRDYC)}}
}

func (p *RCC_Periph) CSSC() RMCIR {
	return RMCIR{mmio.UM32{&p.CIR.U32, uint32(CSSC)}}
}

type APB2RSTR uint32

func (b APB2RSTR) Field(mask APB2RSTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB2RSTR) J(v int) APB2RSTR {
	return APB2RSTR(bits.MakeField32(v, uint32(mask)))
}

type RAPB2RSTR struct{ mmio.U32 }

func (r *RAPB2RSTR) Bits(mask APB2RSTR) APB2RSTR { return APB2RSTR(r.U32.Bits(uint32(mask))) }
func (r *RAPB2RSTR) StoreBits(mask, b APB2RSTR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAPB2RSTR) SetBits(mask APB2RSTR)       { r.U32.SetBits(uint32(mask)) }
func (r *RAPB2RSTR) ClearBits(mask APB2RSTR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RAPB2RSTR) Load() APB2RSTR              { return APB2RSTR(r.U32.Load()) }
func (r *RAPB2RSTR) Store(b APB2RSTR)            { r.U32.Store(uint32(b)) }

func (r *RAPB2RSTR) AtomicStoreBits(mask, b APB2RSTR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RAPB2RSTR) AtomicSetBits(mask APB2RSTR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAPB2RSTR) AtomicClearBits(mask APB2RSTR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMAPB2RSTR struct{ mmio.UM32 }

func (rm RMAPB2RSTR) Load() APB2RSTR   { return APB2RSTR(rm.UM32.Load()) }
func (rm RMAPB2RSTR) Store(b APB2RSTR) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) SYSCFGRST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(SYSCFGRST)}}
}

func (p *RCC_Periph) TIM1RST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(TIM1RST)}}
}

func (p *RCC_Periph) SPI1RST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(SPI1RST)}}
}

func (p *RCC_Periph) TIM8RST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(TIM8RST)}}
}

func (p *RCC_Periph) USART1RST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(USART1RST)}}
}

func (p *RCC_Periph) SPI4RST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(SPI4RST)}}
}

func (p *RCC_Periph) TIM15RST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(TIM15RST)}}
}

func (p *RCC_Periph) TIM16RST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(TIM16RST)}}
}

func (p *RCC_Periph) TIM17RST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(TIM17RST)}}
}

func (p *RCC_Periph) TIM20RST() RMAPB2RSTR {
	return RMAPB2RSTR{mmio.UM32{&p.APB2RSTR.U32, uint32(TIM20RST)}}
}

type APB1RSTR uint32

func (b APB1RSTR) Field(mask APB1RSTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB1RSTR) J(v int) APB1RSTR {
	return APB1RSTR(bits.MakeField32(v, uint32(mask)))
}

type RAPB1RSTR struct{ mmio.U32 }

func (r *RAPB1RSTR) Bits(mask APB1RSTR) APB1RSTR { return APB1RSTR(r.U32.Bits(uint32(mask))) }
func (r *RAPB1RSTR) StoreBits(mask, b APB1RSTR)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAPB1RSTR) SetBits(mask APB1RSTR)       { r.U32.SetBits(uint32(mask)) }
func (r *RAPB1RSTR) ClearBits(mask APB1RSTR)     { r.U32.ClearBits(uint32(mask)) }
func (r *RAPB1RSTR) Load() APB1RSTR              { return APB1RSTR(r.U32.Load()) }
func (r *RAPB1RSTR) Store(b APB1RSTR)            { r.U32.Store(uint32(b)) }

func (r *RAPB1RSTR) AtomicStoreBits(mask, b APB1RSTR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RAPB1RSTR) AtomicSetBits(mask APB1RSTR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAPB1RSTR) AtomicClearBits(mask APB1RSTR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMAPB1RSTR struct{ mmio.UM32 }

func (rm RMAPB1RSTR) Load() APB1RSTR   { return APB1RSTR(rm.UM32.Load()) }
func (rm RMAPB1RSTR) Store(b APB1RSTR) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) TIM2RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM2RST)}}
}

func (p *RCC_Periph) TIM3RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM3RST)}}
}

func (p *RCC_Periph) TIM4RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM4RST)}}
}

func (p *RCC_Periph) TIM6RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM6RST)}}
}

func (p *RCC_Periph) TIM7RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(TIM7RST)}}
}

func (p *RCC_Periph) WWDGRST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(WWDGRST)}}
}

func (p *RCC_Periph) SPI2RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(SPI2RST)}}
}

func (p *RCC_Periph) SPI3RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(SPI3RST)}}
}

func (p *RCC_Periph) USART2RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(USART2RST)}}
}

func (p *RCC_Periph) USART3RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(USART3RST)}}
}

func (p *RCC_Periph) UART4RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(UART4RST)}}
}

func (p *RCC_Periph) UART5RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(UART5RST)}}
}

func (p *RCC_Periph) I2C1RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(I2C1RST)}}
}

func (p *RCC_Periph) I2C2RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(I2C2RST)}}
}

func (p *RCC_Periph) USBRST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(USBRST)}}
}

func (p *RCC_Periph) CANRST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(CANRST)}}
}

func (p *RCC_Periph) PWRRST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(PWRRST)}}
}

func (p *RCC_Periph) DAC1RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(DAC1RST)}}
}

func (p *RCC_Periph) I2C3RST() RMAPB1RSTR {
	return RMAPB1RSTR{mmio.UM32{&p.APB1RSTR.U32, uint32(I2C3RST)}}
}

type AHBENR uint32

func (b AHBENR) Field(mask AHBENR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AHBENR) J(v int) AHBENR {
	return AHBENR(bits.MakeField32(v, uint32(mask)))
}

type RAHBENR struct{ mmio.U32 }

func (r *RAHBENR) Bits(mask AHBENR) AHBENR  { return AHBENR(r.U32.Bits(uint32(mask))) }
func (r *RAHBENR) StoreBits(mask, b AHBENR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAHBENR) SetBits(mask AHBENR)      { r.U32.SetBits(uint32(mask)) }
func (r *RAHBENR) ClearBits(mask AHBENR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RAHBENR) Load() AHBENR             { return AHBENR(r.U32.Load()) }
func (r *RAHBENR) Store(b AHBENR)           { r.U32.Store(uint32(b)) }

func (r *RAHBENR) AtomicStoreBits(mask, b AHBENR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RAHBENR) AtomicSetBits(mask AHBENR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAHBENR) AtomicClearBits(mask AHBENR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMAHBENR struct{ mmio.UM32 }

func (rm RMAHBENR) Load() AHBENR   { return AHBENR(rm.UM32.Load()) }
func (rm RMAHBENR) Store(b AHBENR) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) DMA1EN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(DMA1EN)}}
}

func (p *RCC_Periph) DMA2EN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(DMA2EN)}}
}

func (p *RCC_Periph) SRAMEN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(SRAMEN)}}
}

func (p *RCC_Periph) FLITFEN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(FLITFEN)}}
}

func (p *RCC_Periph) FMCEN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(FMCEN)}}
}

func (p *RCC_Periph) CRCEN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(CRCEN)}}
}

func (p *RCC_Periph) GPIOHEN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(GPIOHEN)}}
}

func (p *RCC_Periph) GPIOAEN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(GPIOAEN)}}
}

func (p *RCC_Periph) GPIOBEN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(GPIOBEN)}}
}

func (p *RCC_Periph) GPIOCEN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(GPIOCEN)}}
}

func (p *RCC_Periph) GPIODEN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(GPIODEN)}}
}

func (p *RCC_Periph) GPIOEEN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(GPIOEEN)}}
}

func (p *RCC_Periph) GPIOFEN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(GPIOFEN)}}
}

func (p *RCC_Periph) GPIOGEN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(GPIOGEN)}}
}

func (p *RCC_Periph) TSCEN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(TSCEN)}}
}

func (p *RCC_Periph) ADC12EN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(ADC12EN)}}
}

func (p *RCC_Periph) ADC34EN() RMAHBENR {
	return RMAHBENR{mmio.UM32{&p.AHBENR.U32, uint32(ADC34EN)}}
}

type APB2ENR uint32

func (b APB2ENR) Field(mask APB2ENR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB2ENR) J(v int) APB2ENR {
	return APB2ENR(bits.MakeField32(v, uint32(mask)))
}

type RAPB2ENR struct{ mmio.U32 }

func (r *RAPB2ENR) Bits(mask APB2ENR) APB2ENR { return APB2ENR(r.U32.Bits(uint32(mask))) }
func (r *RAPB2ENR) StoreBits(mask, b APB2ENR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAPB2ENR) SetBits(mask APB2ENR)      { r.U32.SetBits(uint32(mask)) }
func (r *RAPB2ENR) ClearBits(mask APB2ENR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RAPB2ENR) Load() APB2ENR             { return APB2ENR(r.U32.Load()) }
func (r *RAPB2ENR) Store(b APB2ENR)           { r.U32.Store(uint32(b)) }

func (r *RAPB2ENR) AtomicStoreBits(mask, b APB2ENR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RAPB2ENR) AtomicSetBits(mask APB2ENR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAPB2ENR) AtomicClearBits(mask APB2ENR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMAPB2ENR struct{ mmio.UM32 }

func (rm RMAPB2ENR) Load() APB2ENR   { return APB2ENR(rm.UM32.Load()) }
func (rm RMAPB2ENR) Store(b APB2ENR) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) SYSCFGEN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(SYSCFGEN)}}
}

func (p *RCC_Periph) TIM1EN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(TIM1EN)}}
}

func (p *RCC_Periph) SPI1EN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(SPI1EN)}}
}

func (p *RCC_Periph) TIM8EN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(TIM8EN)}}
}

func (p *RCC_Periph) USART1EN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(USART1EN)}}
}

func (p *RCC_Periph) SPI4EN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(SPI4EN)}}
}

func (p *RCC_Periph) TIM15EN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(TIM15EN)}}
}

func (p *RCC_Periph) TIM16EN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(TIM16EN)}}
}

func (p *RCC_Periph) TIM17EN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(TIM17EN)}}
}

func (p *RCC_Periph) TIM20EN() RMAPB2ENR {
	return RMAPB2ENR{mmio.UM32{&p.APB2ENR.U32, uint32(TIM20EN)}}
}

type APB1ENR uint32

func (b APB1ENR) Field(mask APB1ENR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask APB1ENR) J(v int) APB1ENR {
	return APB1ENR(bits.MakeField32(v, uint32(mask)))
}

type RAPB1ENR struct{ mmio.U32 }

func (r *RAPB1ENR) Bits(mask APB1ENR) APB1ENR { return APB1ENR(r.U32.Bits(uint32(mask))) }
func (r *RAPB1ENR) StoreBits(mask, b APB1ENR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAPB1ENR) SetBits(mask APB1ENR)      { r.U32.SetBits(uint32(mask)) }
func (r *RAPB1ENR) ClearBits(mask APB1ENR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RAPB1ENR) Load() APB1ENR             { return APB1ENR(r.U32.Load()) }
func (r *RAPB1ENR) Store(b APB1ENR)           { r.U32.Store(uint32(b)) }

func (r *RAPB1ENR) AtomicStoreBits(mask, b APB1ENR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RAPB1ENR) AtomicSetBits(mask APB1ENR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAPB1ENR) AtomicClearBits(mask APB1ENR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMAPB1ENR struct{ mmio.UM32 }

func (rm RMAPB1ENR) Load() APB1ENR   { return APB1ENR(rm.UM32.Load()) }
func (rm RMAPB1ENR) Store(b APB1ENR) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) TIM2EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(TIM2EN)}}
}

func (p *RCC_Periph) TIM3EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(TIM3EN)}}
}

func (p *RCC_Periph) TIM4EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(TIM4EN)}}
}

func (p *RCC_Periph) TIM6EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(TIM6EN)}}
}

func (p *RCC_Periph) TIM7EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(TIM7EN)}}
}

func (p *RCC_Periph) WWDGEN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(WWDGEN)}}
}

func (p *RCC_Periph) SPI2EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(SPI2EN)}}
}

func (p *RCC_Periph) SPI3EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(SPI3EN)}}
}

func (p *RCC_Periph) USART2EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(USART2EN)}}
}

func (p *RCC_Periph) USART3EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(USART3EN)}}
}

func (p *RCC_Periph) UART4EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(UART4EN)}}
}

func (p *RCC_Periph) UART5EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(UART5EN)}}
}

func (p *RCC_Periph) I2C1EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(I2C1EN)}}
}

func (p *RCC_Periph) I2C2EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(I2C2EN)}}
}

func (p *RCC_Periph) USBEN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(USBEN)}}
}

func (p *RCC_Periph) CANEN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(CANEN)}}
}

func (p *RCC_Periph) PWREN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(PWREN)}}
}

func (p *RCC_Periph) DAC1EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(DAC1EN)}}
}

func (p *RCC_Periph) I2C3EN() RMAPB1ENR {
	return RMAPB1ENR{mmio.UM32{&p.APB1ENR.U32, uint32(I2C3EN)}}
}

type BDCR uint32

func (b BDCR) Field(mask BDCR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BDCR) J(v int) BDCR {
	return BDCR(bits.MakeField32(v, uint32(mask)))
}

type RBDCR struct{ mmio.U32 }

func (r *RBDCR) Bits(mask BDCR) BDCR    { return BDCR(r.U32.Bits(uint32(mask))) }
func (r *RBDCR) StoreBits(mask, b BDCR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RBDCR) SetBits(mask BDCR)      { r.U32.SetBits(uint32(mask)) }
func (r *RBDCR) ClearBits(mask BDCR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RBDCR) Load() BDCR             { return BDCR(r.U32.Load()) }
func (r *RBDCR) Store(b BDCR)           { r.U32.Store(uint32(b)) }

func (r *RBDCR) AtomicStoreBits(mask, b BDCR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RBDCR) AtomicSetBits(mask BDCR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RBDCR) AtomicClearBits(mask BDCR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMBDCR struct{ mmio.UM32 }

func (rm RMBDCR) Load() BDCR   { return BDCR(rm.UM32.Load()) }
func (rm RMBDCR) Store(b BDCR) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) LSE() RMBDCR {
	return RMBDCR{mmio.UM32{&p.BDCR.U32, uint32(LSE)}}
}

func (p *RCC_Periph) LSEDRV() RMBDCR {
	return RMBDCR{mmio.UM32{&p.BDCR.U32, uint32(LSEDRV)}}
}

func (p *RCC_Periph) RTCSEL() RMBDCR {
	return RMBDCR{mmio.UM32{&p.BDCR.U32, uint32(RTCSEL)}}
}

func (p *RCC_Periph) RTCEN() RMBDCR {
	return RMBDCR{mmio.UM32{&p.BDCR.U32, uint32(RTCEN)}}
}

func (p *RCC_Periph) BDRST() RMBDCR {
	return RMBDCR{mmio.UM32{&p.BDCR.U32, uint32(BDRST)}}
}

type CSR uint32

func (b CSR) Field(mask CSR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CSR) J(v int) CSR {
	return CSR(bits.MakeField32(v, uint32(mask)))
}

type RCSR struct{ mmio.U32 }

func (r *RCSR) Bits(mask CSR) CSR     { return CSR(r.U32.Bits(uint32(mask))) }
func (r *RCSR) StoreBits(mask, b CSR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCSR) SetBits(mask CSR)      { r.U32.SetBits(uint32(mask)) }
func (r *RCSR) ClearBits(mask CSR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCSR) Load() CSR             { return CSR(r.U32.Load()) }
func (r *RCSR) Store(b CSR)           { r.U32.Store(uint32(b)) }

func (r *RCSR) AtomicStoreBits(mask, b CSR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCSR) AtomicSetBits(mask CSR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCSR) AtomicClearBits(mask CSR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCSR struct{ mmio.UM32 }

func (rm RMCSR) Load() CSR   { return CSR(rm.UM32.Load()) }
func (rm RMCSR) Store(b CSR) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) LSION() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(LSION)}}
}

func (p *RCC_Periph) LSIRDY() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(LSIRDY)}}
}

func (p *RCC_Periph) V18PWRRSTF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(V18PWRRSTF)}}
}

func (p *RCC_Periph) RMVF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(RMVF)}}
}

func (p *RCC_Periph) OBLRSTF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(OBLRSTF)}}
}

func (p *RCC_Periph) PINRSTF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(PINRSTF)}}
}

func (p *RCC_Periph) PORRSTF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(PORRSTF)}}
}

func (p *RCC_Periph) SFTRSTF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(SFTRSTF)}}
}

func (p *RCC_Periph) IWDGRSTF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(IWDGRSTF)}}
}

func (p *RCC_Periph) WWDGRSTF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(WWDGRSTF)}}
}

func (p *RCC_Periph) LPWRRSTF() RMCSR {
	return RMCSR{mmio.UM32{&p.CSR.U32, uint32(LPWRRSTF)}}
}

type AHBRSTR uint32

func (b AHBRSTR) Field(mask AHBRSTR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AHBRSTR) J(v int) AHBRSTR {
	return AHBRSTR(bits.MakeField32(v, uint32(mask)))
}

type RAHBRSTR struct{ mmio.U32 }

func (r *RAHBRSTR) Bits(mask AHBRSTR) AHBRSTR { return AHBRSTR(r.U32.Bits(uint32(mask))) }
func (r *RAHBRSTR) StoreBits(mask, b AHBRSTR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAHBRSTR) SetBits(mask AHBRSTR)      { r.U32.SetBits(uint32(mask)) }
func (r *RAHBRSTR) ClearBits(mask AHBRSTR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RAHBRSTR) Load() AHBRSTR             { return AHBRSTR(r.U32.Load()) }
func (r *RAHBRSTR) Store(b AHBRSTR)           { r.U32.Store(uint32(b)) }

func (r *RAHBRSTR) AtomicStoreBits(mask, b AHBRSTR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RAHBRSTR) AtomicSetBits(mask AHBRSTR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAHBRSTR) AtomicClearBits(mask AHBRSTR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMAHBRSTR struct{ mmio.UM32 }

func (rm RMAHBRSTR) Load() AHBRSTR   { return AHBRSTR(rm.UM32.Load()) }
func (rm RMAHBRSTR) Store(b AHBRSTR) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) FMCRST() RMAHBRSTR {
	return RMAHBRSTR{mmio.UM32{&p.AHBRSTR.U32, uint32(FMCRST)}}
}

func (p *RCC_Periph) GPIOHRST() RMAHBRSTR {
	return RMAHBRSTR{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOHRST)}}
}

func (p *RCC_Periph) GPIOARST() RMAHBRSTR {
	return RMAHBRSTR{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOARST)}}
}

func (p *RCC_Periph) GPIOBRST() RMAHBRSTR {
	return RMAHBRSTR{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOBRST)}}
}

func (p *RCC_Periph) GPIOCRST() RMAHBRSTR {
	return RMAHBRSTR{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOCRST)}}
}

func (p *RCC_Periph) GPIODRST() RMAHBRSTR {
	return RMAHBRSTR{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIODRST)}}
}

func (p *RCC_Periph) GPIOERST() RMAHBRSTR {
	return RMAHBRSTR{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOERST)}}
}

func (p *RCC_Periph) GPIOFRST() RMAHBRSTR {
	return RMAHBRSTR{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOFRST)}}
}

func (p *RCC_Periph) GPIOGRST() RMAHBRSTR {
	return RMAHBRSTR{mmio.UM32{&p.AHBRSTR.U32, uint32(GPIOGRST)}}
}

func (p *RCC_Periph) TSCRST() RMAHBRSTR {
	return RMAHBRSTR{mmio.UM32{&p.AHBRSTR.U32, uint32(TSCRST)}}
}

func (p *RCC_Periph) ADC12RST() RMAHBRSTR {
	return RMAHBRSTR{mmio.UM32{&p.AHBRSTR.U32, uint32(ADC12RST)}}
}

func (p *RCC_Periph) ADC34RST() RMAHBRSTR {
	return RMAHBRSTR{mmio.UM32{&p.AHBRSTR.U32, uint32(ADC34RST)}}
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

func (p *RCC_Periph) PREDIV() RMCFGR2 {
	return RMCFGR2{mmio.UM32{&p.CFGR2.U32, uint32(PREDIV)}}
}

func (p *RCC_Periph) ADCPRE12() RMCFGR2 {
	return RMCFGR2{mmio.UM32{&p.CFGR2.U32, uint32(ADCPRE12)}}
}

func (p *RCC_Periph) ADCPRE34() RMCFGR2 {
	return RMCFGR2{mmio.UM32{&p.CFGR2.U32, uint32(ADCPRE34)}}
}

type CFGR3 uint32

func (b CFGR3) Field(mask CFGR3) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CFGR3) J(v int) CFGR3 {
	return CFGR3(bits.MakeField32(v, uint32(mask)))
}

type RCFGR3 struct{ mmio.U32 }

func (r *RCFGR3) Bits(mask CFGR3) CFGR3   { return CFGR3(r.U32.Bits(uint32(mask))) }
func (r *RCFGR3) StoreBits(mask, b CFGR3) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RCFGR3) SetBits(mask CFGR3)      { r.U32.SetBits(uint32(mask)) }
func (r *RCFGR3) ClearBits(mask CFGR3)    { r.U32.ClearBits(uint32(mask)) }
func (r *RCFGR3) Load() CFGR3             { return CFGR3(r.U32.Load()) }
func (r *RCFGR3) Store(b CFGR3)           { r.U32.Store(uint32(b)) }

func (r *RCFGR3) AtomicStoreBits(mask, b CFGR3) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RCFGR3) AtomicSetBits(mask CFGR3)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RCFGR3) AtomicClearBits(mask CFGR3)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMCFGR3 struct{ mmio.UM32 }

func (rm RMCFGR3) Load() CFGR3   { return CFGR3(rm.UM32.Load()) }
func (rm RMCFGR3) Store(b CFGR3) { rm.UM32.Store(uint32(b)) }

func (p *RCC_Periph) USART1SW() RMCFGR3 {
	return RMCFGR3{mmio.UM32{&p.CFGR3.U32, uint32(USART1SW)}}
}

func (p *RCC_Periph) I2CSW() RMCFGR3 {
	return RMCFGR3{mmio.UM32{&p.CFGR3.U32, uint32(I2CSW)}}
}

func (p *RCC_Periph) TIMSW() RMCFGR3 {
	return RMCFGR3{mmio.UM32{&p.CFGR3.U32, uint32(TIMSW)}}
}

func (p *RCC_Periph) TIM2SW() RMCFGR3 {
	return RMCFGR3{mmio.UM32{&p.CFGR3.U32, uint32(TIM2SW)}}
}

func (p *RCC_Periph) TIM34SW() RMCFGR3 {
	return RMCFGR3{mmio.UM32{&p.CFGR3.U32, uint32(TIM34SW)}}
}

func (p *RCC_Periph) USART2SW() RMCFGR3 {
	return RMCFGR3{mmio.UM32{&p.CFGR3.U32, uint32(USART2SW)}}
}

func (p *RCC_Periph) USART3SW() RMCFGR3 {
	return RMCFGR3{mmio.UM32{&p.CFGR3.U32, uint32(USART3SW)}}
}

func (p *RCC_Periph) UART4SW() RMCFGR3 {
	return RMCFGR3{mmio.UM32{&p.CFGR3.U32, uint32(UART4SW)}}
}

func (p *RCC_Periph) UART5SW() RMCFGR3 {
	return RMCFGR3{mmio.UM32{&p.CFGR3.U32, uint32(UART5SW)}}
}
