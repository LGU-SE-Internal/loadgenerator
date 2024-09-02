package behaviors

const ()

var AdvancedSearchChain *Chain

func init() {
	AdvancedSearchChain = NewChain(
		NewFuncNode(VerifyCode, "VerifyCode"),
		NewFuncNode(LoginBasic, "LoginBasic"),
		NewFuncNode(QueryUser, "QueryUser"),
		NewFuncNode(TravelPlan)
	)
}
