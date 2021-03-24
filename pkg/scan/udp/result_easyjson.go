// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package udp

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
	icmp "github.com/v-byte-cpu/sx/pkg/scan/icmp"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonD3b49167DecodeGithubComVByteCpuSxPkgScanUdp(in *jlexer.Lexer, out *ScanResult) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "scan":
			out.ScanType = string(in.String())
		case "ip":
			out.IP = string(in.String())
		case "icmp":
			if in.IsNull() {
				in.Skip()
				out.ICMP = nil
			} else {
				if out.ICMP == nil {
					out.ICMP = new(icmp.Response)
				}
				easyjsonD3b49167DecodeGithubComVByteCpuSxPkgScanIcmp(in, out.ICMP)
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD3b49167EncodeGithubComVByteCpuSxPkgScanUdp(out *jwriter.Writer, in ScanResult) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"scan\":"
		out.RawString(prefix[1:])
		out.String(string(in.ScanType))
	}
	{
		const prefix string = ",\"ip\":"
		out.RawString(prefix)
		out.String(string(in.IP))
	}
	{
		const prefix string = ",\"icmp\":"
		out.RawString(prefix)
		if in.ICMP == nil {
			out.RawString("null")
		} else {
			easyjsonD3b49167EncodeGithubComVByteCpuSxPkgScanIcmp(out, *in.ICMP)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ScanResult) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD3b49167EncodeGithubComVByteCpuSxPkgScanUdp(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ScanResult) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD3b49167EncodeGithubComVByteCpuSxPkgScanUdp(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ScanResult) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD3b49167DecodeGithubComVByteCpuSxPkgScanUdp(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ScanResult) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD3b49167DecodeGithubComVByteCpuSxPkgScanUdp(l, v)
}
func easyjsonD3b49167DecodeGithubComVByteCpuSxPkgScanIcmp(in *jlexer.Lexer, out *icmp.Response) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "Type":
			out.Type = uint8(in.Uint8())
		case "Code":
			out.Code = uint8(in.Uint8())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD3b49167EncodeGithubComVByteCpuSxPkgScanIcmp(out *jwriter.Writer, in icmp.Response) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Type\":"
		out.RawString(prefix[1:])
		out.Uint8(uint8(in.Type))
	}
	{
		const prefix string = ",\"Code\":"
		out.RawString(prefix)
		out.Uint8(uint8(in.Code))
	}
	out.RawByte('}')
}
