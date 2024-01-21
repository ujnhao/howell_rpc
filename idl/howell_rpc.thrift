namespace go coder.hao.howell_rpc 

enum CpsType {
    Unknow = 0
}

struct CpsRebateDiscounts {
    1:optional string Id 
    2:optional i64 AppId
    3:optional string Name
    4:optional CpsType CpsType   
    5:optional string JumpLink 
    6:optional string Extra  
    7:optional i32 Status
}

struct CreateCpsRebateDiscountsRequest {
    1: required CpsRebateDiscounts CRDEntity
}

struct CreateCpsRebateDiscountsResponse {
    1: optional string EntityId 
}

service HowellRpcService {
    CreateCpsRebateDiscountsResponse CreateCpsRebateDiscounts(1: CreateCpsRebateDiscountsRequest req)
}
