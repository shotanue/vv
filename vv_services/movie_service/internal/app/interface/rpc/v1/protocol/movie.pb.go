// Code generated by protoc-gen-go. DO NOT EDIT.
// source: movie.proto

package protocol

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// --------------------------
// types
// --------------------------
type MovieSummary struct {
	Adult                bool     `protobuf:"varint,1,opt,name=adult,proto3" json:"adult,omitempty"`
	BackdropPath         string   `protobuf:"bytes,2,opt,name=backdrop_path,json=backdropPath,proto3" json:"backdrop_path,omitempty"`
	GenreIds             []uint64 `protobuf:"varint,3,rep,packed,name=genre_ids,json=genreIds,proto3" json:"genre_ids,omitempty"`
	Id                   uint64   `protobuf:"varint,4,opt,name=id,proto3" json:"id,omitempty"`
	OriginalLanguage     string   `protobuf:"bytes,5,opt,name=original_language,json=originalLanguage,proto3" json:"original_language,omitempty"`
	OriginalTitle        string   `protobuf:"bytes,6,opt,name=original_title,json=originalTitle,proto3" json:"original_title,omitempty"`
	Overview             string   `protobuf:"bytes,7,opt,name=overview,proto3" json:"overview,omitempty"`
	PosterPath           string   `protobuf:"bytes,8,opt,name=poster_path,json=posterPath,proto3" json:"poster_path,omitempty"`
	ReleaseDate          string   `protobuf:"bytes,9,opt,name=release_date,json=releaseDate,proto3" json:"release_date,omitempty"`
	Title                string   `protobuf:"bytes,10,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MovieSummary) Reset()         { *m = MovieSummary{} }
func (m *MovieSummary) String() string { return proto.CompactTextString(m) }
func (*MovieSummary) ProtoMessage()    {}
func (*MovieSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{0}
}

func (m *MovieSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MovieSummary.Unmarshal(m, b)
}
func (m *MovieSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MovieSummary.Marshal(b, m, deterministic)
}
func (m *MovieSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MovieSummary.Merge(m, src)
}
func (m *MovieSummary) XXX_Size() int {
	return xxx_messageInfo_MovieSummary.Size(m)
}
func (m *MovieSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_MovieSummary.DiscardUnknown(m)
}

var xxx_messageInfo_MovieSummary proto.InternalMessageInfo

func (m *MovieSummary) GetAdult() bool {
	if m != nil {
		return m.Adult
	}
	return false
}

func (m *MovieSummary) GetBackdropPath() string {
	if m != nil {
		return m.BackdropPath
	}
	return ""
}

func (m *MovieSummary) GetGenreIds() []uint64 {
	if m != nil {
		return m.GenreIds
	}
	return nil
}

func (m *MovieSummary) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MovieSummary) GetOriginalLanguage() string {
	if m != nil {
		return m.OriginalLanguage
	}
	return ""
}

func (m *MovieSummary) GetOriginalTitle() string {
	if m != nil {
		return m.OriginalTitle
	}
	return ""
}

func (m *MovieSummary) GetOverview() string {
	if m != nil {
		return m.Overview
	}
	return ""
}

func (m *MovieSummary) GetPosterPath() string {
	if m != nil {
		return m.PosterPath
	}
	return ""
}

func (m *MovieSummary) GetReleaseDate() string {
	if m != nil {
		return m.ReleaseDate
	}
	return ""
}

func (m *MovieSummary) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type Genre struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Genre) Reset()         { *m = Genre{} }
func (m *Genre) String() string { return proto.CompactTextString(m) }
func (*Genre) ProtoMessage()    {}
func (*Genre) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{1}
}

func (m *Genre) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Genre.Unmarshal(m, b)
}
func (m *Genre) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Genre.Marshal(b, m, deterministic)
}
func (m *Genre) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Genre.Merge(m, src)
}
func (m *Genre) XXX_Size() int {
	return xxx_messageInfo_Genre.Size(m)
}
func (m *Genre) XXX_DiscardUnknown() {
	xxx_messageInfo_Genre.DiscardUnknown(m)
}

var xxx_messageInfo_Genre proto.InternalMessageInfo

func (m *Genre) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Genre) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type MovieDetail struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Genres               []*Genre `protobuf:"bytes,3,rep,name=genres,proto3" json:"genres,omitempty"`
	Homepage             string   `protobuf:"bytes,4,opt,name=homepage,proto3" json:"homepage,omitempty"`
	PosterPath           string   `protobuf:"bytes,5,opt,name=poster_path,json=posterPath,proto3" json:"poster_path,omitempty"`
	ReleaseDate          string   `protobuf:"bytes,6,opt,name=release_date,json=releaseDate,proto3" json:"release_date,omitempty"`
	BackdropPath         string   `protobuf:"bytes,7,opt,name=backdrop_path,json=backdropPath,proto3" json:"backdrop_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MovieDetail) Reset()         { *m = MovieDetail{} }
func (m *MovieDetail) String() string { return proto.CompactTextString(m) }
func (*MovieDetail) ProtoMessage()    {}
func (*MovieDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{2}
}

func (m *MovieDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MovieDetail.Unmarshal(m, b)
}
func (m *MovieDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MovieDetail.Marshal(b, m, deterministic)
}
func (m *MovieDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MovieDetail.Merge(m, src)
}
func (m *MovieDetail) XXX_Size() int {
	return xxx_messageInfo_MovieDetail.Size(m)
}
func (m *MovieDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_MovieDetail.DiscardUnknown(m)
}

var xxx_messageInfo_MovieDetail proto.InternalMessageInfo

func (m *MovieDetail) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MovieDetail) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MovieDetail) GetGenres() []*Genre {
	if m != nil {
		return m.Genres
	}
	return nil
}

func (m *MovieDetail) GetHomepage() string {
	if m != nil {
		return m.Homepage
	}
	return ""
}

func (m *MovieDetail) GetPosterPath() string {
	if m != nil {
		return m.PosterPath
	}
	return ""
}

func (m *MovieDetail) GetReleaseDate() string {
	if m != nil {
		return m.ReleaseDate
	}
	return ""
}

func (m *MovieDetail) GetBackdropPath() string {
	if m != nil {
		return m.BackdropPath
	}
	return ""
}

// --------------------------
// requests and responses
// --------------------------
type HotMoviesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HotMoviesRequest) Reset()         { *m = HotMoviesRequest{} }
func (m *HotMoviesRequest) String() string { return proto.CompactTextString(m) }
func (*HotMoviesRequest) ProtoMessage()    {}
func (*HotMoviesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{3}
}

func (m *HotMoviesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HotMoviesRequest.Unmarshal(m, b)
}
func (m *HotMoviesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HotMoviesRequest.Marshal(b, m, deterministic)
}
func (m *HotMoviesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HotMoviesRequest.Merge(m, src)
}
func (m *HotMoviesRequest) XXX_Size() int {
	return xxx_messageInfo_HotMoviesRequest.Size(m)
}
func (m *HotMoviesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HotMoviesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HotMoviesRequest proto.InternalMessageInfo

type HotMoviesResponse struct {
	Summaries            []*MovieSummary `protobuf:"bytes,1,rep,name=summaries,proto3" json:"summaries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *HotMoviesResponse) Reset()         { *m = HotMoviesResponse{} }
func (m *HotMoviesResponse) String() string { return proto.CompactTextString(m) }
func (*HotMoviesResponse) ProtoMessage()    {}
func (*HotMoviesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{4}
}

func (m *HotMoviesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HotMoviesResponse.Unmarshal(m, b)
}
func (m *HotMoviesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HotMoviesResponse.Marshal(b, m, deterministic)
}
func (m *HotMoviesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HotMoviesResponse.Merge(m, src)
}
func (m *HotMoviesResponse) XXX_Size() int {
	return xxx_messageInfo_HotMoviesResponse.Size(m)
}
func (m *HotMoviesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HotMoviesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HotMoviesResponse proto.InternalMessageInfo

func (m *HotMoviesResponse) GetSummaries() []*MovieSummary {
	if m != nil {
		return m.Summaries
	}
	return nil
}

type ComingSoonMoviesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComingSoonMoviesRequest) Reset()         { *m = ComingSoonMoviesRequest{} }
func (m *ComingSoonMoviesRequest) String() string { return proto.CompactTextString(m) }
func (*ComingSoonMoviesRequest) ProtoMessage()    {}
func (*ComingSoonMoviesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{5}
}

func (m *ComingSoonMoviesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComingSoonMoviesRequest.Unmarshal(m, b)
}
func (m *ComingSoonMoviesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComingSoonMoviesRequest.Marshal(b, m, deterministic)
}
func (m *ComingSoonMoviesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComingSoonMoviesRequest.Merge(m, src)
}
func (m *ComingSoonMoviesRequest) XXX_Size() int {
	return xxx_messageInfo_ComingSoonMoviesRequest.Size(m)
}
func (m *ComingSoonMoviesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ComingSoonMoviesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ComingSoonMoviesRequest proto.InternalMessageInfo

type ComingSoonMoviesResponse struct {
	Summaries            []*MovieSummary `protobuf:"bytes,1,rep,name=summaries,proto3" json:"summaries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ComingSoonMoviesResponse) Reset()         { *m = ComingSoonMoviesResponse{} }
func (m *ComingSoonMoviesResponse) String() string { return proto.CompactTextString(m) }
func (*ComingSoonMoviesResponse) ProtoMessage()    {}
func (*ComingSoonMoviesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{6}
}

func (m *ComingSoonMoviesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComingSoonMoviesResponse.Unmarshal(m, b)
}
func (m *ComingSoonMoviesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComingSoonMoviesResponse.Marshal(b, m, deterministic)
}
func (m *ComingSoonMoviesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComingSoonMoviesResponse.Merge(m, src)
}
func (m *ComingSoonMoviesResponse) XXX_Size() int {
	return xxx_messageInfo_ComingSoonMoviesResponse.Size(m)
}
func (m *ComingSoonMoviesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ComingSoonMoviesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ComingSoonMoviesResponse proto.InternalMessageInfo

func (m *ComingSoonMoviesResponse) GetSummaries() []*MovieSummary {
	if m != nil {
		return m.Summaries
	}
	return nil
}

type MovieDetailRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MovieDetailRequest) Reset()         { *m = MovieDetailRequest{} }
func (m *MovieDetailRequest) String() string { return proto.CompactTextString(m) }
func (*MovieDetailRequest) ProtoMessage()    {}
func (*MovieDetailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{7}
}

func (m *MovieDetailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MovieDetailRequest.Unmarshal(m, b)
}
func (m *MovieDetailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MovieDetailRequest.Marshal(b, m, deterministic)
}
func (m *MovieDetailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MovieDetailRequest.Merge(m, src)
}
func (m *MovieDetailRequest) XXX_Size() int {
	return xxx_messageInfo_MovieDetailRequest.Size(m)
}
func (m *MovieDetailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MovieDetailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MovieDetailRequest proto.InternalMessageInfo

type MovieDetailResponse struct {
	Detail               *MovieDetail `protobuf:"bytes,1,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *MovieDetailResponse) Reset()         { *m = MovieDetailResponse{} }
func (m *MovieDetailResponse) String() string { return proto.CompactTextString(m) }
func (*MovieDetailResponse) ProtoMessage()    {}
func (*MovieDetailResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fde087a4194eda75, []int{8}
}

func (m *MovieDetailResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MovieDetailResponse.Unmarshal(m, b)
}
func (m *MovieDetailResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MovieDetailResponse.Marshal(b, m, deterministic)
}
func (m *MovieDetailResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MovieDetailResponse.Merge(m, src)
}
func (m *MovieDetailResponse) XXX_Size() int {
	return xxx_messageInfo_MovieDetailResponse.Size(m)
}
func (m *MovieDetailResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MovieDetailResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MovieDetailResponse proto.InternalMessageInfo

func (m *MovieDetailResponse) GetDetail() *MovieDetail {
	if m != nil {
		return m.Detail
	}
	return nil
}

func init() {
	proto.RegisterType((*MovieSummary)(nil), "protocol.MovieSummary")
	proto.RegisterType((*Genre)(nil), "protocol.Genre")
	proto.RegisterType((*MovieDetail)(nil), "protocol.MovieDetail")
	proto.RegisterType((*HotMoviesRequest)(nil), "protocol.HotMoviesRequest")
	proto.RegisterType((*HotMoviesResponse)(nil), "protocol.HotMoviesResponse")
	proto.RegisterType((*ComingSoonMoviesRequest)(nil), "protocol.ComingSoonMoviesRequest")
	proto.RegisterType((*ComingSoonMoviesResponse)(nil), "protocol.ComingSoonMoviesResponse")
	proto.RegisterType((*MovieDetailRequest)(nil), "protocol.MovieDetailRequest")
	proto.RegisterType((*MovieDetailResponse)(nil), "protocol.MovieDetailResponse")
}

func init() { proto.RegisterFile("movie.proto", fileDescriptor_fde087a4194eda75) }

var fileDescriptor_fde087a4194eda75 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x55, 0xb2, 0x34, 0x4b, 0x6e, 0xba, 0xd1, 0x99, 0x01, 0xa6, 0x03, 0xd1, 0x05, 0x21, 0x2a,
	0x4d, 0xf4, 0xa1, 0xf0, 0x0f, 0xa8, 0x04, 0x43, 0x20, 0x4d, 0x19, 0x2f, 0x3c, 0x55, 0x5e, 0x73,
	0x95, 0x5a, 0x24, 0x71, 0x88, 0xdd, 0x22, 0xfe, 0x1d, 0x7f, 0x85, 0xdf, 0xc1, 0x0b, 0x8a, 0xe3,
	0xa4, 0x1f, 0xd9, 0xe0, 0x81, 0xa7, 0xf6, 0x9e, 0x73, 0x7c, 0x8f, 0xef, 0xf1, 0x0d, 0x04, 0x99,
	0x58, 0x73, 0x9c, 0x14, 0xa5, 0x50, 0x82, 0x78, 0xfa, 0x67, 0x21, 0xd2, 0xf0, 0xa7, 0x0d, 0xfd,
	0x4f, 0x15, 0x73, 0xbd, 0xca, 0x32, 0x56, 0xfe, 0x20, 0xa7, 0xd0, 0x63, 0xf1, 0x2a, 0x55, 0xd4,
	0x1a, 0x59, 0x63, 0x2f, 0xaa, 0x0b, 0xf2, 0x1c, 0x8e, 0x6e, 0xd8, 0xe2, 0x6b, 0x5c, 0x8a, 0x62,
	0x5e, 0x30, 0xb5, 0xa4, 0xf6, 0xc8, 0x1a, 0xfb, 0x51, 0xbf, 0x01, 0xaf, 0x98, 0x5a, 0x92, 0x33,
	0xf0, 0x13, 0xcc, 0x4b, 0x9c, 0xf3, 0x58, 0xd2, 0x83, 0xd1, 0xc1, 0xd8, 0x89, 0x3c, 0x0d, 0x5c,
	0xc6, 0x92, 0x1c, 0x83, 0xcd, 0x63, 0xea, 0x8c, 0xac, 0xb1, 0x13, 0xd9, 0x3c, 0x26, 0x17, 0x70,
	0x22, 0x4a, 0x9e, 0xf0, 0x9c, 0xa5, 0xf3, 0x94, 0xe5, 0xc9, 0x8a, 0x25, 0x48, 0x7b, 0xba, 0xeb,
	0xa0, 0x21, 0x3e, 0x1a, 0x9c, 0xbc, 0x80, 0xe3, 0x56, 0xac, 0xb8, 0x4a, 0x91, 0xba, 0x5a, 0x79,
	0xd4, 0xa0, 0x9f, 0x2b, 0x90, 0x0c, 0xc1, 0x13, 0x6b, 0x2c, 0xd7, 0x1c, 0xbf, 0xd3, 0x43, 0x2d,
	0x68, 0x6b, 0xf2, 0x0c, 0x82, 0x42, 0x48, 0x85, 0x65, 0x7d, 0x7f, 0x4f, 0xd3, 0x50, 0x43, 0xfa,
	0xf6, 0xe7, 0xd0, 0x2f, 0x31, 0x45, 0x26, 0x71, 0x1e, 0x33, 0x85, 0xd4, 0xd7, 0x8a, 0xc0, 0x60,
	0x33, 0xa6, 0xb0, 0xca, 0xa6, 0x76, 0x07, 0xcd, 0xd5, 0x45, 0x78, 0x01, 0xbd, 0x77, 0xd5, 0x94,
	0x66, 0x44, 0xab, 0x1d, 0x91, 0x80, 0x93, 0xb3, 0x0c, 0x4d, 0x56, 0xfa, 0x7f, 0xf8, 0xcb, 0x82,
	0x40, 0xe7, 0x3d, 0x43, 0xc5, 0x78, 0xda, 0x39, 0xd3, 0x5a, 0xd8, 0x5b, 0x16, 0xe4, 0x25, 0xb8,
	0x3a, 0xc8, 0x3a, 0xd6, 0x60, 0x7a, 0x6f, 0xd2, 0x3c, 0xe0, 0x44, 0x5b, 0x47, 0x86, 0xae, 0x12,
	0x58, 0x8a, 0x0c, 0x8b, 0x2a, 0x4c, 0xa7, 0x4e, 0xa0, 0xa9, 0xf7, 0x13, 0xe8, 0xfd, 0x33, 0x01,
	0xb7, 0x9b, 0x40, 0x67, 0x0f, 0x0e, 0xbb, 0x7b, 0x10, 0x12, 0x18, 0xbc, 0x17, 0x4a, 0x4f, 0x29,
	0x23, 0xfc, 0xb6, 0x42, 0xa9, 0xc2, 0x4b, 0x38, 0xd9, 0xc2, 0x64, 0x21, 0x72, 0x89, 0xe4, 0x0d,
	0xf8, 0x52, 0xaf, 0x1d, 0x47, 0x49, 0x2d, 0x3d, 0xd9, 0xc3, 0xcd, 0x64, 0xdb, 0x6b, 0x19, 0x6d,
	0x84, 0xe1, 0x63, 0x78, 0xf4, 0x56, 0x64, 0x3c, 0x4f, 0xae, 0x85, 0xc8, 0x77, 0x5d, 0xae, 0x80,
	0x76, 0xa9, 0xff, 0x32, 0x3b, 0x05, 0xb2, 0xf5, 0x5c, 0x8d, 0xcf, 0x0c, 0xee, 0xef, 0xa0, 0xc6,
	0xe2, 0x15, 0xb8, 0xb1, 0x46, 0xf4, 0x83, 0x06, 0xd3, 0x07, 0x7b, 0xfd, 0x8d, 0xdc, 0x88, 0xa6,
	0xbf, 0xad, 0xe6, 0xdb, 0xab, 0x76, 0x74, 0x81, 0x64, 0x06, 0x7e, 0x1b, 0x12, 0x19, 0x6e, 0x0e,
	0xef, 0xa7, 0x39, 0x3c, 0xbb, 0x95, 0x33, 0xb7, 0xf8, 0x02, 0x83, 0xfd, 0x10, 0xc8, 0xf9, 0xe6,
	0xc0, 0x1d, 0xd9, 0x0d, 0xc3, 0xbf, 0x49, 0x4c, 0xeb, 0x0f, 0xbb, 0xcb, 0xfb, 0xe4, 0xf6, 0xf9,
	0x4c, 0xc3, 0xa7, 0x77, 0xb0, 0x75, 0xaf, 0x1b, 0x57, 0xb3, 0xaf, 0xff, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xcd, 0x25, 0x87, 0x35, 0x99, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MovieServiceClient is the client API for MovieService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MovieServiceClient interface {
	HotMovies(ctx context.Context, in *HotMoviesRequest, opts ...grpc.CallOption) (*HotMoviesResponse, error)
	ComingSoonMovies(ctx context.Context, in *ComingSoonMoviesRequest, opts ...grpc.CallOption) (*ComingSoonMoviesResponse, error)
	MovieDetail(ctx context.Context, in *MovieDetailRequest, opts ...grpc.CallOption) (*MovieDetailResponse, error)
}

type movieServiceClient struct {
	cc *grpc.ClientConn
}

func NewMovieServiceClient(cc *grpc.ClientConn) MovieServiceClient {
	return &movieServiceClient{cc}
}

func (c *movieServiceClient) HotMovies(ctx context.Context, in *HotMoviesRequest, opts ...grpc.CallOption) (*HotMoviesResponse, error) {
	out := new(HotMoviesResponse)
	err := c.cc.Invoke(ctx, "/protocol.MovieService/HotMovies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieServiceClient) ComingSoonMovies(ctx context.Context, in *ComingSoonMoviesRequest, opts ...grpc.CallOption) (*ComingSoonMoviesResponse, error) {
	out := new(ComingSoonMoviesResponse)
	err := c.cc.Invoke(ctx, "/protocol.MovieService/ComingSoonMovies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *movieServiceClient) MovieDetail(ctx context.Context, in *MovieDetailRequest, opts ...grpc.CallOption) (*MovieDetailResponse, error) {
	out := new(MovieDetailResponse)
	err := c.cc.Invoke(ctx, "/protocol.MovieService/MovieDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MovieServiceServer is the server API for MovieService service.
type MovieServiceServer interface {
	HotMovies(context.Context, *HotMoviesRequest) (*HotMoviesResponse, error)
	ComingSoonMovies(context.Context, *ComingSoonMoviesRequest) (*ComingSoonMoviesResponse, error)
	MovieDetail(context.Context, *MovieDetailRequest) (*MovieDetailResponse, error)
}

// UnimplementedMovieServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMovieServiceServer struct {
}

func (*UnimplementedMovieServiceServer) HotMovies(ctx context.Context, req *HotMoviesRequest) (*HotMoviesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HotMovies not implemented")
}
func (*UnimplementedMovieServiceServer) ComingSoonMovies(ctx context.Context, req *ComingSoonMoviesRequest) (*ComingSoonMoviesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ComingSoonMovies not implemented")
}
func (*UnimplementedMovieServiceServer) MovieDetail(ctx context.Context, req *MovieDetailRequest) (*MovieDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MovieDetail not implemented")
}

func RegisterMovieServiceServer(s *grpc.Server, srv MovieServiceServer) {
	s.RegisterService(&_MovieService_serviceDesc, srv)
}

func _MovieService_HotMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HotMoviesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).HotMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.MovieService/HotMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).HotMovies(ctx, req.(*HotMoviesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieService_ComingSoonMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ComingSoonMoviesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).ComingSoonMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.MovieService/ComingSoonMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).ComingSoonMovies(ctx, req.(*ComingSoonMoviesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MovieService_MovieDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovieDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MovieServiceServer).MovieDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol.MovieService/MovieDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MovieServiceServer).MovieDetail(ctx, req.(*MovieDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MovieService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.MovieService",
	HandlerType: (*MovieServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HotMovies",
			Handler:    _MovieService_HotMovies_Handler,
		},
		{
			MethodName: "ComingSoonMovies",
			Handler:    _MovieService_ComingSoonMovies_Handler,
		},
		{
			MethodName: "MovieDetail",
			Handler:    _MovieService_MovieDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "movie.proto",
}