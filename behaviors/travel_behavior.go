package behaviors

const ()

var TravelChain *Chain

func init() {
	TravelChain = NewChain(NewFuncNode(func(context *Context) (*NodeResult, error) {
		return nil, nil
	}, "DummyTravelChain"))
	LoginChain.AddNextChain(NewChain(NewFuncNode(LoginAdmin, "LoginAdmin")), 1)
}
