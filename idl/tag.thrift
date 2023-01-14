namespace go api

enum Code {
     Success         = 1
     ParamInvalid    = 2
     DBErr           = 3
}

struct Tag {
    1: i32 tag_id
    2: string name
    3: i8  state
}

struct GetTagListRequest {
    1: string  name 
	2: i8  state 
    3: i64 page 
    4: i64 page_size
}

struct TagListResponse{
   1: Code code
   2: string msg
   3: list<Tag> tags
   4: i64 totoal
}

service TagService {
    TagListResponse GetTagList(1:GetTagListRequest req)
}