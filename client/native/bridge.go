package native

import (
	"time"

	"github.com/CyCoreSystems/ari"
)

// Bridge provides the ARI Bridge accessors for the native client
type Bridge struct {
	client *Client
}

// Create creates a bridge and returns the lazy handle for the bridge
func (b *Bridge) Create(id string, t string, name string) (bh ari.BridgeHandle, err error) {

	req := struct {
		ID   string `json:"bridgeId,omitempty"`
		Type string `json:"type,omitempty"`
		Name string `json:"name,omitempty"`
	}{
		ID:   id,
		Type: t,
		Name: name,
	}

	err = b.client.post("/bridges/"+id, &req, nil)
	if err != nil {
		return
	}

	bh = b.Get(id)
	return
}

// Get gets the lazy handle for the given bridge id
func (b *Bridge) Get(id string) ari.BridgeHandle {
	return NewBridgeHandle(id, b)
}

// List lists the current bridges and returns a list of lazy handles
func (b *Bridge) List() (bx []ari.BridgeHandle, err error) {
	var bridges = []struct {
		ID string `json:"id"`
	}{}

	err = b.client.get("/bridges", &bridges)
	for _, i := range bridges {
		bx = append(bx, b.Get(i.ID))
	}
	return
}

// Data returns the details of a bridge
// Equivalent to Get /bridges/{bridgeId}
func (b *Bridge) Data(id string) (bd *ari.BridgeData, err error) {
	bd = &ari.BridgeData{}
	err = b.client.get("/bridges/"+id, bd)
	if err != nil {
		bd = nil
		err = dataGetError(err, "bridge", "%v", id)
	}
	return
}

// AddChannel adds a channel to a bridge
// Equivalent to Post /bridges/{id}/addChannel
func (b *Bridge) AddChannel(bridgeID string, channelID string) (err error) {

	type request struct {
		ChannelID string `json:"channel"`
		Role      string `json:"role,omitempty"`
	}

	req := request{channelID, ""}
	err = b.client.post("/bridges/"+bridgeID+"/addChannel", nil, &req)
	return
}

// RemoveChannel removes the specified channel from a bridge
// Equivalent to Post /bridges/{id}/removeChannel
func (b *Bridge) RemoveChannel(id string, channelID string) (err error) {
	req := struct {
		ChannelID string `json:"channel"`
	}{
		ChannelID: channelID,
	}

	//pass request
	err = b.client.post("/bridges/"+id+"/removeChannel", nil, &req)
	return
}

// Delete shuts down a bridge. If any channels are in this bridge,
// they will be removed and resume whatever they were doing beforehand.
// This means that the channels themselves are not deleted.
// Equivalent to DELETE /bridges/{id}
func (b *Bridge) Delete(id string) (err error) {
	err = b.client.del("/bridges/"+id, nil, "")
	return
}

// Play attempts to play the given mediaURI on the bridge, using the playbackID
// as the identifier to the created playback handle
func (b *Bridge) Play(id string, playbackID string, mediaURI string) (ph ari.PlaybackHandle, err error) {
	resp := make(map[string]interface{})
	type request struct {
		Media string `json:"media"`
	}
	req := request{mediaURI}
	err = b.client.post("/bridges/"+id+"/play/"+playbackID, &resp, &req)
	ph = b.client.Playback().Get(playbackID)
	return
}

// Record attempts to record audio on the bridge, using name as the identifier for
// the created live recording handle
func (b *Bridge) Record(id string, name string, opts *ari.RecordingOptions) (rh ari.LiveRecordingHandle, err error) {

	if opts == nil {
		opts = &ari.RecordingOptions{}
	}

	resp := make(map[string]interface{})
	type request struct {
		Name        string `json:"name"`
		Format      string `json:"format"`
		MaxDuration int    `json:"maxDurationSeconds"`
		MaxSilence  int    `json:"maxSilenceSeconds"`
		IfExists    string `json:"ifExists,omitempty"`
		Beep        bool   `json:"beep"`
		TerminateOn string `json:"terminateOn,omitempty"`
	}
	req := request{
		Name:        name,
		Format:      opts.Format,
		MaxDuration: int(opts.MaxDuration / time.Second),
		MaxSilence:  int(opts.MaxSilence / time.Second),
		IfExists:    opts.Exists,
		Beep:        opts.Beep,
		TerminateOn: opts.Terminate,
	}
	err = b.client.post("/bridges/"+id+"/record", &resp, &req)
	if err != nil {
		rh = b.client.LiveRecording().Get(name)
	}
	return
}

// Subscribe creates an event subscription for events related to the given
// bridge␃entity
func (b *Bridge) Subscribe(id string, n ...string) ari.Subscription {
	var ns nativeSubscription

	ns.events = make(chan ari.Event, 10)
	ns.closeChan = make(chan struct{})

	bridgeHandle := b.Get(id)

	go func() {
		sub := b.client.Bus().Subscribe(n...)
		defer sub.Cancel()
		for {

			select {
			case <-ns.closeChan:
				ns.closeChan = nil
				return
			case evt := <-sub.Events():
				//TODO: do we want to send in events on the bridge for a specific channel?
				if bridgeHandle.Match(evt) {
					ns.events <- evt
				}
			}
		}
	}()

	return &ns
}

// NewBridgeHandle creates a new bridge handle
func NewBridgeHandle(id string, b *Bridge) ari.BridgeHandle {
	return &BridgeHandle{
		id: id,
		b:  b,
	}
}

// BridgeHandle is the handle to a bridge for performing operations
type BridgeHandle struct {
	id string
	b  *Bridge
}

// ID returns the identifier for the bridge
func (bh *BridgeHandle) ID() string {
	return bh.id
}

// AddChannel adds a channel to the bridge
func (bh *BridgeHandle) AddChannel(channelID string) (err error) {
	err = bh.b.AddChannel(bh.id, channelID)
	return
}

// RemoveChannel removes a channel from the bridge
func (bh *BridgeHandle) RemoveChannel(channelID string) (err error) {
	err = bh.b.RemoveChannel(bh.id, channelID)
	return
}

// Delete deletes the bridge
func (bh *BridgeHandle) Delete() (err error) {
	err = bh.b.Delete(bh.id)
	return
}

// Data gets the bridge data
func (bh *BridgeHandle) Data() (bd *ari.BridgeData, err error) {
	bd, err = bh.b.Data(bh.id)
	return
}

// Play initiates playback of the specified media uri
// to the bridge, returning the Playback handle
func (bh *BridgeHandle) Play(id string, mediaURI string) (ph ari.PlaybackHandle, err error) {
	ph, err = bh.b.Play(bh.id, id, mediaURI)
	return
}

// Record records the bridge to the given filename
func (bh *BridgeHandle) Record(name string, opts *ari.RecordingOptions) (rh ari.LiveRecordingHandle, err error) {
	rh, err = bh.b.Record(bh.id, name, opts)
	return
}

/* FIXME(scm): what is this for?
// Playback returns the playback transport
func (bh *BridgeHandle) Playback() Playback {
	if pb, ok := bh.b.(Playbacker); ok {
		return pb.Playback()
	}
	return nil
}
*/

// Subscribe creates a subscription to the list of events
func (bh *BridgeHandle) Subscribe(n ...string) ari.Subscription {
	if bh == nil {
		return nil
	}
	return bh.b.Subscribe(bh.id, n...)
}

// Match returns true if the event matches the bridge
func (bh *BridgeHandle) Match(e ari.Event) bool {
	bridgeEvent, ok := e.(ari.BridgeEvent)
	if !ok {
		return false
	}
	bridgeIDs := bridgeEvent.GetBridgeIDs()
	for _, i := range bridgeIDs {
		if i == bh.id {
			return true
		}
	}
	return false
}
