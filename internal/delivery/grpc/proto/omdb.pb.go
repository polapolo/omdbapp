// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: omdb.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MovieSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keyword string `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	Page    int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *MovieSearchRequest) Reset() {
	*x = MovieSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieSearchRequest) ProtoMessage() {}

func (x *MovieSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_omdb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieSearchRequest.ProtoReflect.Descriptor instead.
func (*MovieSearchRequest) Descriptor() ([]byte, []int) {
	return file_omdb_proto_rawDescGZIP(), []int{0}
}

func (x *MovieSearchRequest) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

func (x *MovieSearchRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type MovieSearchResponseDataSearch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Year   string `protobuf:"bytes,2,opt,name=year,proto3" json:"year,omitempty"`
	ImdbId string `protobuf:"bytes,3,opt,name=imdb_id,json=imdbId,proto3" json:"imdb_id,omitempty"`
	Type   string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Poster string `protobuf:"bytes,5,opt,name=poster,proto3" json:"poster,omitempty"`
}

func (x *MovieSearchResponseDataSearch) Reset() {
	*x = MovieSearchResponseDataSearch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieSearchResponseDataSearch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieSearchResponseDataSearch) ProtoMessage() {}

func (x *MovieSearchResponseDataSearch) ProtoReflect() protoreflect.Message {
	mi := &file_omdb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieSearchResponseDataSearch.ProtoReflect.Descriptor instead.
func (*MovieSearchResponseDataSearch) Descriptor() ([]byte, []int) {
	return file_omdb_proto_rawDescGZIP(), []int{1}
}

func (x *MovieSearchResponseDataSearch) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MovieSearchResponseDataSearch) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *MovieSearchResponseDataSearch) GetImdbId() string {
	if x != nil {
		return x.ImdbId
	}
	return ""
}

func (x *MovieSearchResponseDataSearch) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MovieSearchResponseDataSearch) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

type MovieSearchResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Search       []*MovieSearchResponseDataSearch `protobuf:"bytes,1,rep,name=search,proto3" json:"search,omitempty"`
	TotalResults string                           `protobuf:"bytes,2,opt,name=total_results,json=totalResults,proto3" json:"total_results,omitempty"`
	Response     string                           `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
	Error        string                           `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *MovieSearchResponseData) Reset() {
	*x = MovieSearchResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieSearchResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieSearchResponseData) ProtoMessage() {}

func (x *MovieSearchResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_omdb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieSearchResponseData.ProtoReflect.Descriptor instead.
func (*MovieSearchResponseData) Descriptor() ([]byte, []int) {
	return file_omdb_proto_rawDescGZIP(), []int{2}
}

func (x *MovieSearchResponseData) GetSearch() []*MovieSearchResponseDataSearch {
	if x != nil {
		return x.Search
	}
	return nil
}

func (x *MovieSearchResponseData) GetTotalResults() string {
	if x != nil {
		return x.TotalResults
	}
	return ""
}

func (x *MovieSearchResponseData) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *MovieSearchResponseData) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type MovieSearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *MovieSearchResponseData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Error string                   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *MovieSearchResponse) Reset() {
	*x = MovieSearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieSearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieSearchResponse) ProtoMessage() {}

func (x *MovieSearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_omdb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieSearchResponse.ProtoReflect.Descriptor instead.
func (*MovieSearchResponse) Descriptor() ([]byte, []int) {
	return file_omdb_proto_rawDescGZIP(), []int{3}
}

func (x *MovieSearchResponse) GetData() *MovieSearchResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *MovieSearchResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type MovieDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MovieDetailRequest) Reset() {
	*x = MovieDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieDetailRequest) ProtoMessage() {}

func (x *MovieDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_omdb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieDetailRequest.ProtoReflect.Descriptor instead.
func (*MovieDetailRequest) Descriptor() ([]byte, []int) {
	return file_omdb_proto_rawDescGZIP(), []int{4}
}

func (x *MovieDetailRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type MovieDetailResponseDataRating struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Value  string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MovieDetailResponseDataRating) Reset() {
	*x = MovieDetailResponseDataRating{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieDetailResponseDataRating) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieDetailResponseDataRating) ProtoMessage() {}

func (x *MovieDetailResponseDataRating) ProtoReflect() protoreflect.Message {
	mi := &file_omdb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieDetailResponseDataRating.ProtoReflect.Descriptor instead.
func (*MovieDetailResponseDataRating) Descriptor() ([]byte, []int) {
	return file_omdb_proto_rawDescGZIP(), []int{5}
}

func (x *MovieDetailResponseDataRating) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *MovieDetailResponseDataRating) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type MovieDetailResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string                           `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Year       string                           `protobuf:"bytes,2,opt,name=year,proto3" json:"year,omitempty"`
	Rated      string                           `protobuf:"bytes,3,opt,name=rated,proto3" json:"rated,omitempty"`
	Released   string                           `protobuf:"bytes,4,opt,name=released,proto3" json:"released,omitempty"`
	Runtime    string                           `protobuf:"bytes,5,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Genre      string                           `protobuf:"bytes,6,opt,name=genre,proto3" json:"genre,omitempty"`
	Director   string                           `protobuf:"bytes,7,opt,name=director,proto3" json:"director,omitempty"`
	Writer     string                           `protobuf:"bytes,8,opt,name=writer,proto3" json:"writer,omitempty"`
	Actors     string                           `protobuf:"bytes,9,opt,name=actors,proto3" json:"actors,omitempty"`
	Plot       string                           `protobuf:"bytes,10,opt,name=plot,proto3" json:"plot,omitempty"`
	Language   string                           `protobuf:"bytes,11,opt,name=language,proto3" json:"language,omitempty"`
	Country    string                           `protobuf:"bytes,12,opt,name=country,proto3" json:"country,omitempty"`
	Awards     string                           `protobuf:"bytes,13,opt,name=awards,proto3" json:"awards,omitempty"`
	Poster     string                           `protobuf:"bytes,14,opt,name=poster,proto3" json:"poster,omitempty"`
	Ratings    []*MovieDetailResponseDataRating `protobuf:"bytes,15,rep,name=ratings,proto3" json:"ratings,omitempty"`
	Metascore  string                           `protobuf:"bytes,16,opt,name=metascore,proto3" json:"metascore,omitempty"`
	ImdbRating string                           `protobuf:"bytes,17,opt,name=imdb_rating,json=imdbRating,proto3" json:"imdb_rating,omitempty"`
	ImdbVotes  string                           `protobuf:"bytes,18,opt,name=imdb_votes,json=imdbVotes,proto3" json:"imdb_votes,omitempty"`
	ImdbId     string                           `protobuf:"bytes,19,opt,name=imdb_id,json=imdbId,proto3" json:"imdb_id,omitempty"`
	Type       string                           `protobuf:"bytes,20,opt,name=type,proto3" json:"type,omitempty"`
	Dvd        string                           `protobuf:"bytes,21,opt,name=dvd,proto3" json:"dvd,omitempty"`
	BoxOffice  string                           `protobuf:"bytes,22,opt,name=box_office,json=boxOffice,proto3" json:"box_office,omitempty"`
	Production string                           `protobuf:"bytes,23,opt,name=production,proto3" json:"production,omitempty"`
	Website    string                           `protobuf:"bytes,24,opt,name=website,proto3" json:"website,omitempty"`
	Response   string                           `protobuf:"bytes,25,opt,name=response,proto3" json:"response,omitempty"`
	Error      string                           `protobuf:"bytes,26,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *MovieDetailResponseData) Reset() {
	*x = MovieDetailResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieDetailResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieDetailResponseData) ProtoMessage() {}

func (x *MovieDetailResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_omdb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieDetailResponseData.ProtoReflect.Descriptor instead.
func (*MovieDetailResponseData) Descriptor() ([]byte, []int) {
	return file_omdb_proto_rawDescGZIP(), []int{6}
}

func (x *MovieDetailResponseData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MovieDetailResponseData) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *MovieDetailResponseData) GetRated() string {
	if x != nil {
		return x.Rated
	}
	return ""
}

func (x *MovieDetailResponseData) GetReleased() string {
	if x != nil {
		return x.Released
	}
	return ""
}

func (x *MovieDetailResponseData) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *MovieDetailResponseData) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *MovieDetailResponseData) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

func (x *MovieDetailResponseData) GetWriter() string {
	if x != nil {
		return x.Writer
	}
	return ""
}

func (x *MovieDetailResponseData) GetActors() string {
	if x != nil {
		return x.Actors
	}
	return ""
}

func (x *MovieDetailResponseData) GetPlot() string {
	if x != nil {
		return x.Plot
	}
	return ""
}

func (x *MovieDetailResponseData) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *MovieDetailResponseData) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *MovieDetailResponseData) GetAwards() string {
	if x != nil {
		return x.Awards
	}
	return ""
}

func (x *MovieDetailResponseData) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

func (x *MovieDetailResponseData) GetRatings() []*MovieDetailResponseDataRating {
	if x != nil {
		return x.Ratings
	}
	return nil
}

func (x *MovieDetailResponseData) GetMetascore() string {
	if x != nil {
		return x.Metascore
	}
	return ""
}

func (x *MovieDetailResponseData) GetImdbRating() string {
	if x != nil {
		return x.ImdbRating
	}
	return ""
}

func (x *MovieDetailResponseData) GetImdbVotes() string {
	if x != nil {
		return x.ImdbVotes
	}
	return ""
}

func (x *MovieDetailResponseData) GetImdbId() string {
	if x != nil {
		return x.ImdbId
	}
	return ""
}

func (x *MovieDetailResponseData) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *MovieDetailResponseData) GetDvd() string {
	if x != nil {
		return x.Dvd
	}
	return ""
}

func (x *MovieDetailResponseData) GetBoxOffice() string {
	if x != nil {
		return x.BoxOffice
	}
	return ""
}

func (x *MovieDetailResponseData) GetProduction() string {
	if x != nil {
		return x.Production
	}
	return ""
}

func (x *MovieDetailResponseData) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *MovieDetailResponseData) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

func (x *MovieDetailResponseData) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type MovieDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  *MovieDetailResponseData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Error string                   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *MovieDetailResponse) Reset() {
	*x = MovieDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_omdb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MovieDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MovieDetailResponse) ProtoMessage() {}

func (x *MovieDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_omdb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MovieDetailResponse.ProtoReflect.Descriptor instead.
func (*MovieDetailResponse) Descriptor() ([]byte, []int) {
	return file_omdb_proto_rawDescGZIP(), []int{7}
}

func (x *MovieDetailResponse) GetData() *MovieDetailResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *MovieDetailResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_omdb_proto protoreflect.FileDescriptor

var file_omdb_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6f, 0x6d, 0x64, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x12, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x8e, 0x01, 0x0a, 0x1d, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79,
	0x65, 0x61, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x6d, 0x64, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x22, 0xae, 0x01, 0x0a, 0x17, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x3c, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x5f, 0x0a, 0x13, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x32, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x24, 0x0a, 0x12, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x4d, 0x0a, 0x1d, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0xd3, 0x05, 0x0a, 0x17, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6c, 0x6f, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x6c, 0x6f, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x12, 0x3e, 0x0a,
	0x07, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x6d, 0x64, 0x62, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x69, 0x6d, 0x64, 0x62, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a,
	0x69, 0x6d, 0x64, 0x62, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x69, 0x6d, 0x64, 0x62, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x69,
	0x6d, 0x64, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d,
	0x64, 0x62, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x76, 0x64, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x76, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f,
	0x78, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x62, 0x6f, 0x78, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62,
	0x73, 0x69, 0x74, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73,
	0x69, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x5f, 0x0a, 0x13, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x96, 0x01, 0x0a, 0x04, 0x4f, 0x4d, 0x44, 0x42, 0x12,
	0x46, 0x0a, 0x0b, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x19,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0b, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_omdb_proto_rawDescOnce sync.Once
	file_omdb_proto_rawDescData = file_omdb_proto_rawDesc
)

func file_omdb_proto_rawDescGZIP() []byte {
	file_omdb_proto_rawDescOnce.Do(func() {
		file_omdb_proto_rawDescData = protoimpl.X.CompressGZIP(file_omdb_proto_rawDescData)
	})
	return file_omdb_proto_rawDescData
}

var file_omdb_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_omdb_proto_goTypes = []interface{}{
	(*MovieSearchRequest)(nil),            // 0: proto.MovieSearchRequest
	(*MovieSearchResponseDataSearch)(nil), // 1: proto.MovieSearchResponseDataSearch
	(*MovieSearchResponseData)(nil),       // 2: proto.MovieSearchResponseData
	(*MovieSearchResponse)(nil),           // 3: proto.MovieSearchResponse
	(*MovieDetailRequest)(nil),            // 4: proto.MovieDetailRequest
	(*MovieDetailResponseDataRating)(nil), // 5: proto.MovieDetailResponseDataRating
	(*MovieDetailResponseData)(nil),       // 6: proto.MovieDetailResponseData
	(*MovieDetailResponse)(nil),           // 7: proto.MovieDetailResponse
}
var file_omdb_proto_depIdxs = []int32{
	1, // 0: proto.MovieSearchResponseData.search:type_name -> proto.MovieSearchResponseDataSearch
	2, // 1: proto.MovieSearchResponse.data:type_name -> proto.MovieSearchResponseData
	5, // 2: proto.MovieDetailResponseData.ratings:type_name -> proto.MovieDetailResponseDataRating
	6, // 3: proto.MovieDetailResponse.data:type_name -> proto.MovieDetailResponseData
	0, // 4: proto.OMDB.MovieSearch:input_type -> proto.MovieSearchRequest
	4, // 5: proto.OMDB.MovieDetail:input_type -> proto.MovieDetailRequest
	3, // 6: proto.OMDB.MovieSearch:output_type -> proto.MovieSearchResponse
	7, // 7: proto.OMDB.MovieDetail:output_type -> proto.MovieDetailResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_omdb_proto_init() }
func file_omdb_proto_init() {
	if File_omdb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_omdb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieSearchRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_omdb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieSearchResponseDataSearch); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_omdb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieSearchResponseData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_omdb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieSearchResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_omdb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieDetailRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_omdb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieDetailResponseDataRating); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_omdb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieDetailResponseData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_omdb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MovieDetailResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_omdb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_omdb_proto_goTypes,
		DependencyIndexes: file_omdb_proto_depIdxs,
		MessageInfos:      file_omdb_proto_msgTypes,
	}.Build()
	File_omdb_proto = out.File
	file_omdb_proto_rawDesc = nil
	file_omdb_proto_goTypes = nil
	file_omdb_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OMDBClient is the client API for OMDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OMDBClient interface {
	MovieSearch(ctx context.Context, in *MovieSearchRequest, opts ...grpc.CallOption) (*MovieSearchResponse, error)
	MovieDetail(ctx context.Context, in *MovieDetailRequest, opts ...grpc.CallOption) (*MovieDetailResponse, error)
}

type oMDBClient struct {
	cc grpc.ClientConnInterface
}

func NewOMDBClient(cc grpc.ClientConnInterface) OMDBClient {
	return &oMDBClient{cc}
}

func (c *oMDBClient) MovieSearch(ctx context.Context, in *MovieSearchRequest, opts ...grpc.CallOption) (*MovieSearchResponse, error) {
	out := new(MovieSearchResponse)
	err := c.cc.Invoke(ctx, "/proto.OMDB/MovieSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oMDBClient) MovieDetail(ctx context.Context, in *MovieDetailRequest, opts ...grpc.CallOption) (*MovieDetailResponse, error) {
	out := new(MovieDetailResponse)
	err := c.cc.Invoke(ctx, "/proto.OMDB/MovieDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OMDBServer is the server API for OMDB service.
type OMDBServer interface {
	MovieSearch(context.Context, *MovieSearchRequest) (*MovieSearchResponse, error)
	MovieDetail(context.Context, *MovieDetailRequest) (*MovieDetailResponse, error)
}

// UnimplementedOMDBServer can be embedded to have forward compatible implementations.
type UnimplementedOMDBServer struct {
}

func (*UnimplementedOMDBServer) MovieSearch(context.Context, *MovieSearchRequest) (*MovieSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MovieSearch not implemented")
}
func (*UnimplementedOMDBServer) MovieDetail(context.Context, *MovieDetailRequest) (*MovieDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MovieDetail not implemented")
}

func RegisterOMDBServer(s *grpc.Server, srv OMDBServer) {
	s.RegisterService(&_OMDB_serviceDesc, srv)
}

func _OMDB_MovieSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovieSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OMDBServer).MovieSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OMDB/MovieSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OMDBServer).MovieSearch(ctx, req.(*MovieSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OMDB_MovieDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MovieDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OMDBServer).MovieDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.OMDB/MovieDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OMDBServer).MovieDetail(ctx, req.(*MovieDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OMDB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.OMDB",
	HandlerType: (*OMDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MovieSearch",
			Handler:    _OMDB_MovieSearch_Handler,
		},
		{
			MethodName: "MovieDetail",
			Handler:    _OMDB_MovieDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "omdb.proto",
}