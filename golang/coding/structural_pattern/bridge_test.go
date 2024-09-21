package structural_pattern

import "testing"

/*
Bridge:
When we have two factors that can change independently, and we need the combination of them,
use bridge to connect them rather than implementing a class to couple them.

3*4 > 3+4
*/

func TestBridge(t *testing.T) {
	tv := &TV{}
	basicRemote := &BasicRemote{device: tv}
	basicRemote.Power()
	basicRemote.VolumeUp()
	basicRemote.VolumeDown()

	radio := &Radio{}
	advancedRemote := &AdvancedRemote{device: radio}
	advancedRemote.Power()
	advancedRemote.VolumeUp()
	advancedRemote.VolumeDown()
}

type Remote interface {
	Power()
	VolumeUp()
	VolumeDown()
}

type Device interface {
	On()
	Off()
	IsOn() bool
	SetVolume(volume int)
	GetVolume() int
}

type BasicRemote struct {
	device Device
}

func (b *BasicRemote) Power() {
	if b.device.IsOn() {
		b.device.Off()
		return
	}
	b.device.On()
}

func (b *BasicRemote) VolumeUp() {
	b.device.SetVolume(b.device.GetVolume() + 1)
}

func (b *BasicRemote) VolumeDown() {
	b.device.SetVolume(b.device.GetVolume() - 1)
}

type AdvancedRemote struct {
	device      Device
	volumeUpCnt int
}

func (a *AdvancedRemote) Power() {
	if a.device.IsOn() {
		a.device.Off()
		return
	}
	a.device.On()
}

func (a *AdvancedRemote) VolumeUp() {
	a.volumeUpCnt++
	if a.volumeUpCnt < 2 {
		a.device.SetVolume(a.device.GetVolume() + 1)
		return
	}
	a.device.SetVolume(a.device.GetVolume() + 5)
}

func (a *AdvancedRemote) VolumeDown() {
	a.volumeUpCnt = 0
	a.device.SetVolume(a.device.GetVolume() - 1)
}

type TV struct {
	isOn   bool
	volume int
}

func (t *TV) On() {
	t.isOn = true
}

func (t *TV) Off() {
	t.isOn = false
}

func (t *TV) IsOn() bool {
	return t.isOn
}

func (t *TV) SetVolume(volume int) {
	t.volume = volume
}

func (t *TV) GetVolume() int {
	return t.volume
}

type Radio struct {
	isOn   bool
	volume int
}

func (r *Radio) On() {
	r.isOn = true
}

func (r *Radio) Off() {
	r.isOn = false
}

func (r *Radio) IsOn() bool {
	return r.isOn
}

func (r *Radio) SetVolume(volume int) {
	r.volume = volume
}

func (r *Radio) GetVolume() int {
	return r.volume
}
