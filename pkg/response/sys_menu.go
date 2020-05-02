package response

// 菜单树信息响应, 字段含义见models.SysMenu
type MenuTreeResponseStruct struct {
	Name       string                   `json:"name"`
	Title      string                   `json:"title"`
	Icon       string                   `json:"icon"`
	Path       string                   `json:"path"`
	Redirect   string                   `json:"redirect"`
	Component  string                   `json:"component"`
	Permission string                   `json:"permission"`
	Sort       int                      `json:"sort"`
	Status     bool                     `json:"status"`
	Visible    bool                     `json:"visible"`
	Breadcrumb bool                     `json:"breadcrumb"`
	Children   []MenuTreeResponseStruct `json:"children"`
}