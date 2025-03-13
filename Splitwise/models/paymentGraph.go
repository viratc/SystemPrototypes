package models

type PaymentGraph struct {
	Graph map[string]BalanceMap
}

func NewPaymentGraph(graph map[string]BalanceMap) *PaymentGraph {
	return &PaymentGraph{Graph: graph}
}

func (p *PaymentGraph) GetPaymentGraph() map[string]BalanceMap {
	return p.Graph
}
