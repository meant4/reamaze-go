package reamaze

const responseTemplatesEndpoint string = "/api/v1/response_templates"

type GetResponseTemplatesResponse struct {
	ResponseTemplates []struct {
		ID                    int    `json:"id"`
		Name                  string `json:"name"`
		Body                  string `json:"body"`
		ResponseTemplateGroup struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"response_template_group"`
	} `json:"response_templates"`
}

type GetResponseTemplateResponse struct {
	ID                    int    `json:"id"`
	Name                  string `json:"name"`
	Body                  string `json:"body"`
	ResponseTemplateGroup struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"response_template_group"`
}

type CreateResponseTemplateRequest struct {
	ResponseTemplate struct {
		Name       string `json:"name"`
		Body       string `json:"body"`
		IsPersonal bool   `json:"is_personal"`
	} `json:"response_template"`
}

type CreateResponseTemplateResponse GetResponseTemplateResponse
type UpdateResponseTemplateResponse GetResponseTemplateResponse
type UpdateResponseTemplateRequest CreateResponseTemplateRequest
