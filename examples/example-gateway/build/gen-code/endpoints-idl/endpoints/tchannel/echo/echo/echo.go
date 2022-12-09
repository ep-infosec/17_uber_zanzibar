// Code generated by thriftrw v1.29.2. DO NOT EDIT.
// @generated

package echo

import (
	errors "errors"
	fmt "fmt"
	strings "strings"

	stream "go.uber.org/thriftrw/protocol/stream"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
)

// Echo_Echo_Args represents the arguments for the Echo.echo function.
//
// The arguments for echo are sent and received over the wire as this struct.
type Echo_Echo_Args struct {
	Msg string `json:"msg,required"`
}

// ToWire translates a Echo_Echo_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//	x, err := v.ToWire()
//	if err != nil {
//	  return err
//	}
//
//	if err := binaryProtocol.Encode(x, writer); err != nil {
//	  return err
//	}
func (v *Echo_Echo_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Msg), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Echo_Echo_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Echo_Echo_Args struct
// from the provided intermediate representation.
//
//	x, err := binaryProtocol.Decode(reader, wire.TStruct)
//	if err != nil {
//	  return nil, err
//	}
//
//	var v Echo_Echo_Args
//	if err := v.FromWire(x); err != nil {
//	  return nil, err
//	}
//	return &v, nil
func (v *Echo_Echo_Args) FromWire(w wire.Value) error {
	var err error

	msgIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Msg, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				msgIsSet = true
			}
		}
	}

	if !msgIsSet {
		return errors.New("field Msg of Echo_Echo_Args is required")
	}

	return nil
}

// Encode serializes a Echo_Echo_Args struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a Echo_Echo_Args struct could not be encoded.
func (v *Echo_Echo_Args) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TBinary}); err != nil {
		return err
	}
	if err := sw.WriteString(v.Msg); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	return sw.WriteStructEnd()
}

// Decode deserializes a Echo_Echo_Args struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a Echo_Echo_Args struct could not be generated from the wire
// representation.
func (v *Echo_Echo_Args) Decode(sr stream.Reader) error {

	msgIsSet := false

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 1 && fh.Type == wire.TBinary:
			v.Msg, err = sr.ReadString()
			if err != nil {
				return err
			}
			msgIsSet = true
		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	if !msgIsSet {
		return errors.New("field Msg of Echo_Echo_Args is required")
	}

	return nil
}

// String returns a readable string representation of a Echo_Echo_Args
// struct.
func (v *Echo_Echo_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Msg: %v", v.Msg)
	i++

	return fmt.Sprintf("Echo_Echo_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Echo_Echo_Args match the
// provided Echo_Echo_Args.
//
// This function performs a deep comparison.
func (v *Echo_Echo_Args) Equals(rhs *Echo_Echo_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.Msg == rhs.Msg) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Echo_Echo_Args.
func (v *Echo_Echo_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("msg", v.Msg)
	return err
}

// GetMsg returns the value of Msg if it is set or its
// zero value if it is unset.
func (v *Echo_Echo_Args) GetMsg() (o string) {
	if v != nil {
		o = v.Msg
	}
	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "echo" for this struct.
func (v *Echo_Echo_Args) MethodName() string {
	return "echo"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Echo_Echo_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Echo_Echo_Helper provides functions that aid in handling the
// parameters and return values of the Echo.echo
// function.
var Echo_Echo_Helper = struct {
	// Args accepts the parameters of echo in-order and returns
	// the arguments struct for the function.
	Args func(
		msg string,
	) *Echo_Echo_Args

	// IsException returns true if the given error can be thrown
	// by echo.
	//
	// An error can be thrown by echo only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for echo
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// echo into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by echo
	//
	//   value, err := echo(args)
	//   result, err := Echo_Echo_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from echo: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(string, error) (*Echo_Echo_Result, error)

	// UnwrapResponse takes the result struct for echo
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if echo threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := Echo_Echo_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Echo_Echo_Result) (string, error)
}{}

func init() {
	Echo_Echo_Helper.Args = func(
		msg string,
	) *Echo_Echo_Args {
		return &Echo_Echo_Args{
			Msg: msg,
		}
	}

	Echo_Echo_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	Echo_Echo_Helper.WrapResponse = func(success string, err error) (*Echo_Echo_Result, error) {
		if err == nil {
			return &Echo_Echo_Result{Success: &success}, nil
		}

		return nil, err
	}
	Echo_Echo_Helper.UnwrapResponse = func(result *Echo_Echo_Result) (success string, err error) {

		if result.Success != nil {
			success = *result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// Echo_Echo_Result represents the result of a Echo.echo function call.
//
// The result of a echo execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type Echo_Echo_Result struct {
	// Value returned by echo after a successful execution.
	Success *string `json:"success,omitempty"`
}

// ToWire translates a Echo_Echo_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//	x, err := v.ToWire()
//	if err != nil {
//	  return err
//	}
//
//	if err := binaryProtocol.Encode(x, writer); err != nil {
//	  return err
//	}
func (v *Echo_Echo_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = wire.NewValueString(*(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("Echo_Echo_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a Echo_Echo_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Echo_Echo_Result struct
// from the provided intermediate representation.
//
//	x, err := binaryProtocol.Decode(reader, wire.TStruct)
//	if err != nil {
//	  return nil, err
//	}
//
//	var v Echo_Echo_Result
//	if err := v.FromWire(x); err != nil {
//	  return nil, err
//	}
//	return &v, nil
func (v *Echo_Echo_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Success = &x
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Echo_Echo_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// Encode serializes a Echo_Echo_Result struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a Echo_Echo_Result struct could not be encoded.
func (v *Echo_Echo_Result) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if v.Success != nil {
		if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 0, Type: wire.TBinary}); err != nil {
			return err
		}
		if err := sw.WriteString(*(v.Success)); err != nil {
			return err
		}
		if err := sw.WriteFieldEnd(); err != nil {
			return err
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}

	if count != 1 {
		return fmt.Errorf("Echo_Echo_Result should have exactly one field: got %v fields", count)
	}

	return sw.WriteStructEnd()
}

// Decode deserializes a Echo_Echo_Result struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a Echo_Echo_Result struct could not be generated from the wire
// representation.
func (v *Echo_Echo_Result) Decode(sr stream.Reader) error {

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 0 && fh.Type == wire.TBinary:
			var x string
			x, err = sr.ReadString()
			v.Success = &x
			if err != nil {
				return err
			}

		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Echo_Echo_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Echo_Echo_Result
// struct.
func (v *Echo_Echo_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", *(v.Success))
		i++
	}

	return fmt.Sprintf("Echo_Echo_Result{%v}", strings.Join(fields[:i], ", "))
}

func _String_EqualsPtr(lhs, rhs *string) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this Echo_Echo_Result match the
// provided Echo_Echo_Result.
//
// This function performs a deep comparison.
func (v *Echo_Echo_Result) Equals(rhs *Echo_Echo_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_String_EqualsPtr(v.Success, rhs.Success) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Echo_Echo_Result.
func (v *Echo_Echo_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		enc.AddString("success", *v.Success)
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *Echo_Echo_Result) GetSuccess() (o string) {
	if v != nil && v.Success != nil {
		return *v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *Echo_Echo_Result) IsSetSuccess() bool {
	return v != nil && v.Success != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "echo" for this struct.
func (v *Echo_Echo_Result) MethodName() string {
	return "echo"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Echo_Echo_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
