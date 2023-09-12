// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AgentsColumns holds the columns for the "agents" table.
	AgentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
	}
	// AgentsTable holds the schema information for the "agents" table.
	AgentsTable = &schema.Table{
		Name:       "agents",
		Columns:    AgentsColumns,
		PrimaryKey: []*schema.Column{AgentsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "agent_id",
				Unique:  true,
				Columns: []*schema.Column{AgentsColumns[0]},
			},
		},
	}
	// StoragesColumns holds the columns for the "storages" table.
	StoragesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "owner", Type: field.TypeString, Size: 50},
		{Name: "type", Type: field.TypeInt32, Default: 0},
		{Name: "name", Type: field.TypeString, Size: 50},
		{Name: "cid", Type: field.TypeString, Size: 80},
		{Name: "size", Type: field.TypeInt32},
		{Name: "last_modify", Type: field.TypeTime},
		{Name: "parent_id", Type: field.TypeString, Size: 80},
	}
	// StoragesTable holds the schema information for the "storages" table.
	StoragesTable = &schema.Table{
		Name:       "storages",
		Columns:    StoragesColumns,
		PrimaryKey: []*schema.Column{StoragesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "storage_owner",
				Unique:  false,
				Columns: []*schema.Column{StoragesColumns[1]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "country_call_coding", Type: field.TypeString, Size: 8},
		{Name: "telephone_number", Type: field.TypeString, Size: 50},
		{Name: "password", Type: field.TypeString},
		{Name: "create_date", Type: field.TypeTime},
		{Name: "last_login_date", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_id",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[0]},
			},
			{
				Name:    "user_country_call_coding_telephone_number",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[1], UsersColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AgentsTable,
		StoragesTable,
		UsersTable,
	}
)

func init() {
}
