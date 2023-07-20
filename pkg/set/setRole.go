package set

import (
	"github.com/Karthika-Rajagopal/authz0/pkg/authz0"
	"github.com/Karthika-Rajagopal/authz0/pkg/models"
)

func SetRole(filename string, role models.Role) error {
	template := authz0.FileToTemplate(filename)
	template.Roles = append(template.Roles, role)
	authz0.TemplateToFile(template, filename)
	return nil
}
