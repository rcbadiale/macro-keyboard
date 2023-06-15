package models

type ButtonModel struct {
	Name        string   `json:"name" binding:"required"`
	ActionChain []string `json:"action_chain"`
}
