package requests

type Success struct {
    Ok bool
}

type ErrorMessage struct {
    Ok bool
    ErrorMessage string
}

type BuyRequest struct {
    Name string
    Email string
    Num int
}

/*
用户登录接口
*/
type UserLogin struct {
    UserName string
    Password string
}
