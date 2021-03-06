// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"

	"mojo/public/go/application"
	"mojo/public/go/bindings"
	"mojo/public/go/system"
	"mojo/public/interfaces/bindings/mojom_types"

	"mojom/v23clientproxy"

	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/options"
	"v.io/v23/security"
	"v.io/v23/vom"
	"v.io/x/mojo/proxy/util"
	"v.io/x/mojo/transcoder"
	"v.io/x/ref/runtime/factories/roaming"
)

//#include "mojo/public/c/system/handle.h"
import "C"

type v23HeaderReceiver struct {
	delegate    *delegate
	v23Name     string
	ifaceSig    mojom_types.MojomInterface
	desc        map[string]mojom_types.UserDefinedType
	serviceName string
	handle      system.MessagePipeHandle
}

func (r *v23HeaderReceiver) SetupClientProxy(v23Name string, ifaceSig mojom_types.MojomInterface, desc map[string]mojom_types.UserDefinedType, serviceName string, handle system.MessagePipeHandle) (err error) {
	log := r.delegate.ctx
	log.Infof("[server] In SetupProxy(%s, %v, %v, %s, %v)", v23Name, ifaceSig, desc, serviceName, handle)
	r.v23Name = v23Name
	r.ifaceSig = ifaceSig
	r.desc = desc
	r.serviceName = serviceName
	r.handle = handle

	go func() {
		connector := bindings.NewConnector(r.handle, bindings.GetAsyncWaiter())

		// Read generic calls in a loop
		receiver := &messageReceiver{
			header:    r,
			ctx:       r.delegate.ctx,
			connector: connector,
		}
		stub := bindings.NewStub(connector, receiver)
		for {
			if err := stub.ServeRequest(); err != nil {
				connectionError, ok := err.(*bindings.ConnectionError)
				if !ok || !connectionError.Closed() {
					log.Errorf("%v", err)
				}
				break
			}
		}
		r.delegate.stubs = append(r.delegate.stubs, stub)
	}()
	return nil
}

type messageReceiver struct {
	header    *v23HeaderReceiver
	ctx       *context.T
	connector *bindings.Connector
}

func (s *messageReceiver) Accept(message *bindings.Message) (err error) {
	if _, ok := s.header.ifaceSig.Methods[message.Header.Type]; !ok {
		return fmt.Errorf("Method had index %d, but interface only has %d methods",
			message.Header.Type, len(s.header.ifaceSig.Methods))
	}

	methodSig := s.header.ifaceSig.Methods[message.Header.Type]
	methodName := *methodSig.DeclData.ShortName
	// Should we perform validation of flags like generated methods?
	// Does this handle 0-arg methods?

	messageBytes := message.Payload

	if methodSig.ResponseParams == nil {
		// This may be complicated to support because of a race on the server side.
		// When we send the no-response message through mojo we don't know when it has been received
		// and we can close the pipe (one is opened per incoming message).

		// Note that because the client expects no response, it will not receive the following
		// error.
		return fmt.Errorf("this error not received")
	}

	response, err := s.call(s.header.v23Name, methodName, messageBytes, methodSig.Parameters, methodSig.ResponseParams)
	if err != nil {
		return err
	}

	// TODO(alexfandrianto): This assumes that bindings.Encoder has the method
	// WriteRawBytes. We will need to add this to Mojo ourselves.
	// func (e *Encoder) WriteRawBytes(data []byte) {
	// 	first := e.end
	// 	e.claimData(align(len(data), defaultAlignment))
	// 	copy(e.buf[first:], data)
	// }
	//
	// See: https://codereview.chromium.org/1416433002/

	responseHeader := bindings.MessageHeader{
		Type:      message.Header.Type,
		Flags:     bindings.MessageIsResponseFlag,
		RequestId: message.Header.RequestId,
	}
	// responseMessage, err := bindings.EncodeMessage(responseHeader, byteCopyingPayload(response))
	// if err != nil {
	// 	return err
	// }
	// return s.connector.WriteMessage(responseMessage)

	// TODO(alexfandrianto): Replace this block with the above.
	encoder := bindings.NewEncoder()
	if err := responseHeader.Encode(encoder); err != nil {
		return err
	}
	if bytes, handles, err := encoder.Data(); err != nil {
		return err
	} else {
		// response is our payload; append to the end of our slice.
		bytes = append(bytes, response...)

		// This is analogous to bindings.newMessage
		responseMessage := &bindings.Message{
			Header:  responseHeader,
			Bytes:   bytes,
			Handles: handles,
			Payload: response,
		}
		return s.connector.WriteMessage(responseMessage)
	}
}

func (s *messageReceiver) call(name, method string, value []byte, inParamsType mojom_types.MojomStruct, outParamsType *mojom_types.MojomStruct) ([]byte, error) {
	s.ctx.Infof("server: %s.%s: %#v", name, method, inParamsType)
	inVType, err := transcoder.MojomStructToVDLType(inParamsType, s.header.desc)
	if err != nil {
		return nil, err
	}

	// Decode the vom.RawBytes from the mojom bytes and mojom type.
	target := util.StructSplitTarget()
	if err := transcoder.FromMojo(target, value, inVType); err != nil {
		return nil, fmt.Errorf("transcoder.FromMojo failed: %v", err)
	}

	// inVdlValue is a struct, but we need to send []interface.
	inargs := target.Fields()
	inargsIfc := make([]interface{}, len(inargs))
	for i := range inargs {
		inargsIfc[i] = inargs[i]
	}

	// We know that the v23serverproxy will give us back a bunch of
	// data in []interface{}. so we'll want to decode them into *vom.RawBytes.
	s.ctx.Infof("%s %v", method, outParamsType)
	var numParams int
	if outParamsType != nil {
		numParams = len(outParamsType.Fields)
	}
	outargs := make([]*vom.RawBytes, numParams)
	outptrs := make([]interface{}, len(outargs))
	for i := range outargs {
		outptrs[i] = &outargs[i]
	}

	// Now, run the call without any authorization.
	if err := v23.GetClient(s.ctx).Call(s.ctx, name, method, inargsIfc, outptrs, options.ServerAuthorizer{security.AllowEveryone()}); err != nil {
		return nil, err
	}

	if outParamsType == nil {
		return nil, nil
	}

	outVType, err := transcoder.MojomStructToVDLType(*outParamsType, s.header.desc)
	if err != nil {
		return nil, err
	}

	toMojoTarget := transcoder.ToMojomTarget()
	if err := util.JoinRawBytesAsStruct(toMojoTarget, outVType, outargs); err != nil {
		return nil, err
	}
	return toMojoTarget.Bytes(), nil
}

type delegate struct {
	ctx      *context.T
	stubs    []*bindings.Stub
	shutdown v23.Shutdown
}

func (delegate *delegate) Initialize(context application.Context) {
	roaming.SetArgs(context)
	ctx, shutdown := v23.Init()
	delegate.ctx = ctx
	delegate.shutdown = shutdown
	ctx.Infof("delegate.Initialize...")
}

func (delegate *delegate) Create(request v23clientproxy.V23ClientProxy_Request) {
	headerReceiver := &v23HeaderReceiver{delegate: delegate}
	v23Stub := v23clientproxy.NewV23ClientProxyStub(request, headerReceiver, bindings.GetAsyncWaiter())
	delegate.stubs = append(delegate.stubs, v23Stub)

	go func() {
		// Read header message
		if err := v23Stub.ServeRequest(); err != nil {
			connectionError, ok := err.(*bindings.ConnectionError)
			if !ok || !connectionError.Closed() {
				delegate.ctx.Errorf("%v", err)
			}
			return
		}
	}()
}

func (delegate *delegate) AcceptConnection(connection *application.Connection) {
	delegate.ctx.Infof("delegate.AcceptConnection...")
	connection.ProvideServices(&v23clientproxy.V23ClientProxy_ServiceFactory{delegate})
}

func (delegate *delegate) Quit() {
	delegate.ctx.Infof("delegate.Quit...")
	for _, stub := range delegate.stubs {
		stub.Close()
	}
	delegate.shutdown()
}

//export MojoMain
func MojoMain(handle C.MojoHandle) C.MojoResult {
	application.Run(&delegate{}, system.MojoHandle(handle))
	return C.MOJO_RESULT_OK
}

func main() {
}
