package onewire

type Master struct {
	Driver
}

// SendByte sends byte b on 1-wire bus.
func (m *Master) SendByte(b byte) error {
	for i := 0; i < 8; i++ {
		if err := m.SendBit(b & 1); err != nil {
			return err
		}
		b >>= 1
	}
	return nil
}

// ReadByte reads byte from 1-wire bus.
func (m *Master) ReadByte() (byte, error) {
	var b int
	for i := uint(0); i < 8; i++ {
		bit, err := m.RecvBit()
		if err != nil {
			return 0, err
		}
		b += int(bit) << i
	}
	return byte(b), nil
}
