// +build f030x8

package gpio

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f030x8/mmap"
)

type GPIO_Periph struct {
	MODER   RMODER
	OTYPER  ROTYPER
	OSPEEDR ROSPEEDR
	PUPDR   RPUPDR
	IDR     RIDR
	ODR     RODR
	BSRR    RBSRR
	LCKR    RLCKR
	AFR     [2]RAFR
	BRR     RBRR
}

func (p *GPIO_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

//emgo:const
var GPIOA = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOA_BASE)))

//emgo:const
var GPIOB = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOB_BASE)))

//emgo:const
var GPIOC = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOC_BASE)))

//emgo:const
var GPIOD = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOD_BASE)))

//emgo:const
var GPIOF = (*GPIO_Periph)(unsafe.Pointer(uintptr(mmap.GPIOF_BASE)))

type MODER uint32

func (b MODER) Field(mask MODER) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask MODER) J(v int) MODER {
	return MODER(bits.MakeField32(v, uint32(mask)))
}

type RMODER struct{ mmio.U32 }

func (r *RMODER) Bits(mask MODER) MODER   { return MODER(r.U32.Bits(uint32(mask))) }
func (r *RMODER) StoreBits(mask, b MODER) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RMODER) SetBits(mask MODER)      { r.U32.SetBits(uint32(mask)) }
func (r *RMODER) ClearBits(mask MODER)    { r.U32.ClearBits(uint32(mask)) }
func (r *RMODER) Load() MODER             { return MODER(r.U32.Load()) }
func (r *RMODER) Store(b MODER)           { r.U32.Store(uint32(b)) }

func (r *RMODER) AtomicStoreBits(mask, b MODER) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RMODER) AtomicSetBits(mask MODER)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RMODER) AtomicClearBits(mask MODER)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMMODER struct{ mmio.UM32 }

func (rm RMMODER) Load() MODER   { return MODER(rm.UM32.Load()) }
func (rm RMMODER) Store(b MODER) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) MODER0() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER0)}}
}

func (p *GPIO_Periph) MODER1() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER1)}}
}

func (p *GPIO_Periph) MODER2() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER2)}}
}

func (p *GPIO_Periph) MODER3() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER3)}}
}

func (p *GPIO_Periph) MODER4() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER4)}}
}

func (p *GPIO_Periph) MODER5() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER5)}}
}

func (p *GPIO_Periph) MODER6() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER6)}}
}

func (p *GPIO_Periph) MODER7() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER7)}}
}

func (p *GPIO_Periph) MODER8() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER8)}}
}

func (p *GPIO_Periph) MODER9() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER9)}}
}

func (p *GPIO_Periph) MODER10() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER10)}}
}

func (p *GPIO_Periph) MODER11() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER11)}}
}

func (p *GPIO_Periph) MODER12() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER12)}}
}

func (p *GPIO_Periph) MODER13() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER13)}}
}

func (p *GPIO_Periph) MODER14() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER14)}}
}

func (p *GPIO_Periph) MODER15() RMMODER {
	return RMMODER{mmio.UM32{&p.MODER.U32, uint32(MODER15)}}
}

type OTYPER uint32

func (b OTYPER) Field(mask OTYPER) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OTYPER) J(v int) OTYPER {
	return OTYPER(bits.MakeField32(v, uint32(mask)))
}

type ROTYPER struct{ mmio.U32 }

func (r *ROTYPER) Bits(mask OTYPER) OTYPER  { return OTYPER(r.U32.Bits(uint32(mask))) }
func (r *ROTYPER) StoreBits(mask, b OTYPER) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ROTYPER) SetBits(mask OTYPER)      { r.U32.SetBits(uint32(mask)) }
func (r *ROTYPER) ClearBits(mask OTYPER)    { r.U32.ClearBits(uint32(mask)) }
func (r *ROTYPER) Load() OTYPER             { return OTYPER(r.U32.Load()) }
func (r *ROTYPER) Store(b OTYPER)           { r.U32.Store(uint32(b)) }

func (r *ROTYPER) AtomicStoreBits(mask, b OTYPER) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *ROTYPER) AtomicSetBits(mask OTYPER)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ROTYPER) AtomicClearBits(mask OTYPER)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMOTYPER struct{ mmio.UM32 }

func (rm RMOTYPER) Load() OTYPER   { return OTYPER(rm.UM32.Load()) }
func (rm RMOTYPER) Store(b OTYPER) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) OT_0() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_0)}}
}

func (p *GPIO_Periph) OT_1() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_1)}}
}

func (p *GPIO_Periph) OT_2() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_2)}}
}

func (p *GPIO_Periph) OT_3() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_3)}}
}

func (p *GPIO_Periph) OT_4() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_4)}}
}

func (p *GPIO_Periph) OT_5() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_5)}}
}

func (p *GPIO_Periph) OT_6() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_6)}}
}

func (p *GPIO_Periph) OT_7() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_7)}}
}

func (p *GPIO_Periph) OT_8() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_8)}}
}

func (p *GPIO_Periph) OT_9() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_9)}}
}

func (p *GPIO_Periph) OT_10() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_10)}}
}

func (p *GPIO_Periph) OT_11() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_11)}}
}

func (p *GPIO_Periph) OT_12() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_12)}}
}

func (p *GPIO_Periph) OT_13() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_13)}}
}

func (p *GPIO_Periph) OT_14() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_14)}}
}

func (p *GPIO_Periph) OT_15() RMOTYPER {
	return RMOTYPER{mmio.UM32{&p.OTYPER.U32, uint32(OT_15)}}
}

type OSPEEDR uint32

func (b OSPEEDR) Field(mask OSPEEDR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OSPEEDR) J(v int) OSPEEDR {
	return OSPEEDR(bits.MakeField32(v, uint32(mask)))
}

type ROSPEEDR struct{ mmio.U32 }

func (r *ROSPEEDR) Bits(mask OSPEEDR) OSPEEDR { return OSPEEDR(r.U32.Bits(uint32(mask))) }
func (r *ROSPEEDR) StoreBits(mask, b OSPEEDR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ROSPEEDR) SetBits(mask OSPEEDR)      { r.U32.SetBits(uint32(mask)) }
func (r *ROSPEEDR) ClearBits(mask OSPEEDR)    { r.U32.ClearBits(uint32(mask)) }
func (r *ROSPEEDR) Load() OSPEEDR             { return OSPEEDR(r.U32.Load()) }
func (r *ROSPEEDR) Store(b OSPEEDR)           { r.U32.Store(uint32(b)) }

func (r *ROSPEEDR) AtomicStoreBits(mask, b OSPEEDR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *ROSPEEDR) AtomicSetBits(mask OSPEEDR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *ROSPEEDR) AtomicClearBits(mask OSPEEDR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMOSPEEDR struct{ mmio.UM32 }

func (rm RMOSPEEDR) Load() OSPEEDR   { return OSPEEDR(rm.UM32.Load()) }
func (rm RMOSPEEDR) Store(b OSPEEDR) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) OSPEEDR0() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR0)}}
}

func (p *GPIO_Periph) OSPEEDR1() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR1)}}
}

func (p *GPIO_Periph) OSPEEDR2() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR2)}}
}

func (p *GPIO_Periph) OSPEEDR3() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR3)}}
}

func (p *GPIO_Periph) OSPEEDR4() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR4)}}
}

func (p *GPIO_Periph) OSPEEDR5() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR5)}}
}

func (p *GPIO_Periph) OSPEEDR6() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR6)}}
}

func (p *GPIO_Periph) OSPEEDR7() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR7)}}
}

func (p *GPIO_Periph) OSPEEDR8() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR8)}}
}

func (p *GPIO_Periph) OSPEEDR9() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR9)}}
}

func (p *GPIO_Periph) OSPEEDR10() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR10)}}
}

func (p *GPIO_Periph) OSPEEDR11() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR11)}}
}

func (p *GPIO_Periph) OSPEEDR12() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR12)}}
}

func (p *GPIO_Periph) OSPEEDR13() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR13)}}
}

func (p *GPIO_Periph) OSPEEDR14() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR14)}}
}

func (p *GPIO_Periph) OSPEEDR15() RMOSPEEDR {
	return RMOSPEEDR{mmio.UM32{&p.OSPEEDR.U32, uint32(OSPEEDR15)}}
}

type PUPDR uint32

func (b PUPDR) Field(mask PUPDR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask PUPDR) J(v int) PUPDR {
	return PUPDR(bits.MakeField32(v, uint32(mask)))
}

type RPUPDR struct{ mmio.U32 }

func (r *RPUPDR) Bits(mask PUPDR) PUPDR   { return PUPDR(r.U32.Bits(uint32(mask))) }
func (r *RPUPDR) StoreBits(mask, b PUPDR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RPUPDR) SetBits(mask PUPDR)      { r.U32.SetBits(uint32(mask)) }
func (r *RPUPDR) ClearBits(mask PUPDR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RPUPDR) Load() PUPDR             { return PUPDR(r.U32.Load()) }
func (r *RPUPDR) Store(b PUPDR)           { r.U32.Store(uint32(b)) }

func (r *RPUPDR) AtomicStoreBits(mask, b PUPDR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RPUPDR) AtomicSetBits(mask PUPDR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RPUPDR) AtomicClearBits(mask PUPDR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMPUPDR struct{ mmio.UM32 }

func (rm RMPUPDR) Load() PUPDR   { return PUPDR(rm.UM32.Load()) }
func (rm RMPUPDR) Store(b PUPDR) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) PUPDR0() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR0)}}
}

func (p *GPIO_Periph) PUPDR1() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR1)}}
}

func (p *GPIO_Periph) PUPDR2() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR2)}}
}

func (p *GPIO_Periph) PUPDR3() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR3)}}
}

func (p *GPIO_Periph) PUPDR4() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR4)}}
}

func (p *GPIO_Periph) PUPDR5() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR5)}}
}

func (p *GPIO_Periph) PUPDR6() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR6)}}
}

func (p *GPIO_Periph) PUPDR7() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR7)}}
}

func (p *GPIO_Periph) PUPDR8() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR8)}}
}

func (p *GPIO_Periph) PUPDR9() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR9)}}
}

func (p *GPIO_Periph) PUPDR10() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR10)}}
}

func (p *GPIO_Periph) PUPDR11() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR11)}}
}

func (p *GPIO_Periph) PUPDR12() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR12)}}
}

func (p *GPIO_Periph) PUPDR13() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR13)}}
}

func (p *GPIO_Periph) PUPDR14() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR14)}}
}

func (p *GPIO_Periph) PUPDR15() RMPUPDR {
	return RMPUPDR{mmio.UM32{&p.PUPDR.U32, uint32(PUPDR15)}}
}

type IDR uint32

func (b IDR) Field(mask IDR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IDR) J(v int) IDR {
	return IDR(bits.MakeField32(v, uint32(mask)))
}

type RIDR struct{ mmio.U32 }

func (r *RIDR) Bits(mask IDR) IDR     { return IDR(r.U32.Bits(uint32(mask))) }
func (r *RIDR) StoreBits(mask, b IDR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RIDR) SetBits(mask IDR)      { r.U32.SetBits(uint32(mask)) }
func (r *RIDR) ClearBits(mask IDR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RIDR) Load() IDR             { return IDR(r.U32.Load()) }
func (r *RIDR) Store(b IDR)           { r.U32.Store(uint32(b)) }

func (r *RIDR) AtomicStoreBits(mask, b IDR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RIDR) AtomicSetBits(mask IDR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RIDR) AtomicClearBits(mask IDR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMIDR struct{ mmio.UM32 }

func (rm RMIDR) Load() IDR   { return IDR(rm.UM32.Load()) }
func (rm RMIDR) Store(b IDR) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) V0() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(V0)}}
}

func (p *GPIO_Periph) V1() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(V1)}}
}

func (p *GPIO_Periph) V2() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(V2)}}
}

func (p *GPIO_Periph) V3() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(V3)}}
}

func (p *GPIO_Periph) V4() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(V4)}}
}

func (p *GPIO_Periph) V5() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(V5)}}
}

func (p *GPIO_Periph) V6() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(V6)}}
}

func (p *GPIO_Periph) V7() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(V7)}}
}

func (p *GPIO_Periph) V8() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(V8)}}
}

func (p *GPIO_Periph) V9() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(V9)}}
}

func (p *GPIO_Periph) V10() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(V10)}}
}

func (p *GPIO_Periph) V11() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(V11)}}
}

func (p *GPIO_Periph) V12() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(V12)}}
}

func (p *GPIO_Periph) V13() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(V13)}}
}

func (p *GPIO_Periph) V14() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(V14)}}
}

func (p *GPIO_Periph) V15() RMIDR {
	return RMIDR{mmio.UM32{&p.IDR.U32, uint32(V15)}}
}

type ODR uint32

func (b ODR) Field(mask ODR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ODR) J(v int) ODR {
	return ODR(bits.MakeField32(v, uint32(mask)))
}

type RODR struct{ mmio.U32 }

func (r *RODR) Bits(mask ODR) ODR     { return ODR(r.U32.Bits(uint32(mask))) }
func (r *RODR) StoreBits(mask, b ODR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RODR) SetBits(mask ODR)      { r.U32.SetBits(uint32(mask)) }
func (r *RODR) ClearBits(mask ODR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RODR) Load() ODR             { return ODR(r.U32.Load()) }
func (r *RODR) Store(b ODR)           { r.U32.Store(uint32(b)) }

func (r *RODR) AtomicStoreBits(mask, b ODR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RODR) AtomicSetBits(mask ODR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RODR) AtomicClearBits(mask ODR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMODR struct{ mmio.UM32 }

func (rm RMODR) Load() ODR   { return ODR(rm.UM32.Load()) }
func (rm RMODR) Store(b ODR) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) V0() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(V0)}}
}

func (p *GPIO_Periph) V1() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(V1)}}
}

func (p *GPIO_Periph) V2() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(V2)}}
}

func (p *GPIO_Periph) V3() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(V3)}}
}

func (p *GPIO_Periph) V4() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(V4)}}
}

func (p *GPIO_Periph) V5() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(V5)}}
}

func (p *GPIO_Periph) V6() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(V6)}}
}

func (p *GPIO_Periph) V7() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(V7)}}
}

func (p *GPIO_Periph) V8() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(V8)}}
}

func (p *GPIO_Periph) V9() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(V9)}}
}

func (p *GPIO_Periph) V10() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(V10)}}
}

func (p *GPIO_Periph) V11() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(V11)}}
}

func (p *GPIO_Periph) V12() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(V12)}}
}

func (p *GPIO_Periph) V13() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(V13)}}
}

func (p *GPIO_Periph) V14() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(V14)}}
}

func (p *GPIO_Periph) V15() RMODR {
	return RMODR{mmio.UM32{&p.ODR.U32, uint32(V15)}}
}

type BSRR uint32

func (b BSRR) Field(mask BSRR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BSRR) J(v int) BSRR {
	return BSRR(bits.MakeField32(v, uint32(mask)))
}

type RBSRR struct{ mmio.U32 }

func (r *RBSRR) Bits(mask BSRR) BSRR    { return BSRR(r.U32.Bits(uint32(mask))) }
func (r *RBSRR) StoreBits(mask, b BSRR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RBSRR) SetBits(mask BSRR)      { r.U32.SetBits(uint32(mask)) }
func (r *RBSRR) ClearBits(mask BSRR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RBSRR) Load() BSRR             { return BSRR(r.U32.Load()) }
func (r *RBSRR) Store(b BSRR)           { r.U32.Store(uint32(b)) }

func (r *RBSRR) AtomicStoreBits(mask, b BSRR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RBSRR) AtomicSetBits(mask BSRR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RBSRR) AtomicClearBits(mask BSRR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMBSRR struct{ mmio.UM32 }

func (rm RMBSRR) Load() BSRR   { return BSRR(rm.UM32.Load()) }
func (rm RMBSRR) Store(b BSRR) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) BS_0() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS_0)}}
}

func (p *GPIO_Periph) BS_1() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS_1)}}
}

func (p *GPIO_Periph) BS_2() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS_2)}}
}

func (p *GPIO_Periph) BS_3() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS_3)}}
}

func (p *GPIO_Periph) BS_4() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS_4)}}
}

func (p *GPIO_Periph) BS_5() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS_5)}}
}

func (p *GPIO_Periph) BS_6() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS_6)}}
}

func (p *GPIO_Periph) BS_7() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS_7)}}
}

func (p *GPIO_Periph) BS_8() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS_8)}}
}

func (p *GPIO_Periph) BS_9() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS_9)}}
}

func (p *GPIO_Periph) BS_10() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS_10)}}
}

func (p *GPIO_Periph) BS_11() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS_11)}}
}

func (p *GPIO_Periph) BS_12() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS_12)}}
}

func (p *GPIO_Periph) BS_13() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS_13)}}
}

func (p *GPIO_Periph) BS_14() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS_14)}}
}

func (p *GPIO_Periph) BS_15() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BS_15)}}
}

func (p *GPIO_Periph) BR_0() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR_0)}}
}

func (p *GPIO_Periph) BR_1() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR_1)}}
}

func (p *GPIO_Periph) BR_2() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR_2)}}
}

func (p *GPIO_Periph) BR_3() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR_3)}}
}

func (p *GPIO_Periph) BR_4() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR_4)}}
}

func (p *GPIO_Periph) BR_5() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR_5)}}
}

func (p *GPIO_Periph) BR_6() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR_6)}}
}

func (p *GPIO_Periph) BR_7() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR_7)}}
}

func (p *GPIO_Periph) BR_8() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR_8)}}
}

func (p *GPIO_Periph) BR_9() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR_9)}}
}

func (p *GPIO_Periph) BR_10() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR_10)}}
}

func (p *GPIO_Periph) BR_11() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR_11)}}
}

func (p *GPIO_Periph) BR_12() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR_12)}}
}

func (p *GPIO_Periph) BR_13() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR_13)}}
}

func (p *GPIO_Periph) BR_14() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR_14)}}
}

func (p *GPIO_Periph) BR_15() RMBSRR {
	return RMBSRR{mmio.UM32{&p.BSRR.U32, uint32(BR_15)}}
}

type LCKR uint32

func (b LCKR) Field(mask LCKR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask LCKR) J(v int) LCKR {
	return LCKR(bits.MakeField32(v, uint32(mask)))
}

type RLCKR struct{ mmio.U32 }

func (r *RLCKR) Bits(mask LCKR) LCKR    { return LCKR(r.U32.Bits(uint32(mask))) }
func (r *RLCKR) StoreBits(mask, b LCKR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RLCKR) SetBits(mask LCKR)      { r.U32.SetBits(uint32(mask)) }
func (r *RLCKR) ClearBits(mask LCKR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RLCKR) Load() LCKR             { return LCKR(r.U32.Load()) }
func (r *RLCKR) Store(b LCKR)           { r.U32.Store(uint32(b)) }

func (r *RLCKR) AtomicStoreBits(mask, b LCKR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RLCKR) AtomicSetBits(mask LCKR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RLCKR) AtomicClearBits(mask LCKR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMLCKR struct{ mmio.UM32 }

func (rm RMLCKR) Load() LCKR   { return LCKR(rm.UM32.Load()) }
func (rm RMLCKR) Store(b LCKR) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) LCK0() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK0)}}
}

func (p *GPIO_Periph) LCK1() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK1)}}
}

func (p *GPIO_Periph) LCK2() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK2)}}
}

func (p *GPIO_Periph) LCK3() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK3)}}
}

func (p *GPIO_Periph) LCK4() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK4)}}
}

func (p *GPIO_Periph) LCK5() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK5)}}
}

func (p *GPIO_Periph) LCK6() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK6)}}
}

func (p *GPIO_Periph) LCK7() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK7)}}
}

func (p *GPIO_Periph) LCK8() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK8)}}
}

func (p *GPIO_Periph) LCK9() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK9)}}
}

func (p *GPIO_Periph) LCK10() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK10)}}
}

func (p *GPIO_Periph) LCK11() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK11)}}
}

func (p *GPIO_Periph) LCK12() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK12)}}
}

func (p *GPIO_Periph) LCK13() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK13)}}
}

func (p *GPIO_Periph) LCK14() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK14)}}
}

func (p *GPIO_Periph) LCK15() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCK15)}}
}

func (p *GPIO_Periph) LCKK() RMLCKR {
	return RMLCKR{mmio.UM32{&p.LCKR.U32, uint32(LCKK)}}
}

type AFR uint32

func (b AFR) Field(mask AFR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask AFR) J(v int) AFR {
	return AFR(bits.MakeField32(v, uint32(mask)))
}

type RAFR struct{ mmio.U32 }

func (r *RAFR) Bits(mask AFR) AFR     { return AFR(r.U32.Bits(uint32(mask))) }
func (r *RAFR) StoreBits(mask, b AFR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RAFR) SetBits(mask AFR)      { r.U32.SetBits(uint32(mask)) }
func (r *RAFR) ClearBits(mask AFR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RAFR) Load() AFR             { return AFR(r.U32.Load()) }
func (r *RAFR) Store(b AFR)           { r.U32.Store(uint32(b)) }

func (r *RAFR) AtomicStoreBits(mask, b AFR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RAFR) AtomicSetBits(mask AFR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RAFR) AtomicClearBits(mask AFR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMAFR struct{ mmio.UM32 }

func (rm RMAFR) Load() AFR   { return AFR(rm.UM32.Load()) }
func (rm RMAFR) Store(b AFR) { rm.UM32.Store(uint32(b)) }

type BRR uint32

func (b BRR) Field(mask BRR) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BRR) J(v int) BRR {
	return BRR(bits.MakeField32(v, uint32(mask)))
}

type RBRR struct{ mmio.U32 }

func (r *RBRR) Bits(mask BRR) BRR     { return BRR(r.U32.Bits(uint32(mask))) }
func (r *RBRR) StoreBits(mask, b BRR) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RBRR) SetBits(mask BRR)      { r.U32.SetBits(uint32(mask)) }
func (r *RBRR) ClearBits(mask BRR)    { r.U32.ClearBits(uint32(mask)) }
func (r *RBRR) Load() BRR             { return BRR(r.U32.Load()) }
func (r *RBRR) Store(b BRR)           { r.U32.Store(uint32(b)) }

func (r *RBRR) AtomicStoreBits(mask, b BRR) { r.U32.AtomicStoreBits(uint32(mask), uint32(b)) }
func (r *RBRR) AtomicSetBits(mask BRR)      { r.U32.AtomicSetBits(uint32(mask)) }
func (r *RBRR) AtomicClearBits(mask BRR)    { r.U32.AtomicClearBits(uint32(mask)) }

type RMBRR struct{ mmio.UM32 }

func (rm RMBRR) Load() BRR   { return BRR(rm.UM32.Load()) }
func (rm RMBRR) Store(b BRR) { rm.UM32.Store(uint32(b)) }

func (p *GPIO_Periph) BR_0() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR_0)}}
}

func (p *GPIO_Periph) BR_1() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR_1)}}
}

func (p *GPIO_Periph) BR_2() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR_2)}}
}

func (p *GPIO_Periph) BR_3() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR_3)}}
}

func (p *GPIO_Periph) BR_4() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR_4)}}
}

func (p *GPIO_Periph) BR_5() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR_5)}}
}

func (p *GPIO_Periph) BR_6() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR_6)}}
}

func (p *GPIO_Periph) BR_7() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR_7)}}
}

func (p *GPIO_Periph) BR_8() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR_8)}}
}

func (p *GPIO_Periph) BR_9() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR_9)}}
}

func (p *GPIO_Periph) BR_10() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR_10)}}
}

func (p *GPIO_Periph) BR_11() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR_11)}}
}

func (p *GPIO_Periph) BR_12() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR_12)}}
}

func (p *GPIO_Periph) BR_13() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR_13)}}
}

func (p *GPIO_Periph) BR_14() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR_14)}}
}

func (p *GPIO_Periph) BR_15() RMBRR {
	return RMBRR{mmio.UM32{&p.BRR.U32, uint32(BR_15)}}
}
