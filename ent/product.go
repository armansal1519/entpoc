// Code generated by ent, DO NOT EDIT.

package ent

import (
	"entpoc/ent/product"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Product is the model entity for the Product schema.
type Product struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// TableName0 holds the value of the "table_name0" field.
	TableName0 string `json:"table_name0,omitempty"`
	// TableName1 holds the value of the "table_name1" field.
	TableName1 string `json:"table_name1,omitempty"`
	// TableName2 holds the value of the "table_name2" field.
	TableName2 string `json:"table_name2,omitempty"`
	// TableName3 holds the value of the "table_name3" field.
	TableName3 string `json:"table_name3,omitempty"`
	// TableName4 holds the value of the "table_name4" field.
	TableName4 string `json:"table_name4,omitempty"`
	// TableName5 holds the value of the "table_name5" field.
	TableName5 string `json:"table_name5,omitempty"`
	// TableName6 holds the value of the "table_name6" field.
	TableName6 string `json:"table_name6,omitempty"`
	// TableName7 holds the value of the "table_name7" field.
	TableName7 string `json:"table_name7,omitempty"`
	// TableName8 holds the value of the "table_name8" field.
	TableName8 string `json:"table_name8,omitempty"`
	// TableName9 holds the value of the "table_name9" field.
	TableName9 string `json:"table_name9,omitempty"`
	// TableName10 holds the value of the "table_name10" field.
	TableName10 string `json:"table_name10,omitempty"`
	// TableName11 holds the value of the "table_name11" field.
	TableName11  string `json:"table_name11,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Product) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case product.FieldID:
			values[i] = new(sql.NullInt64)
		case product.FieldName, product.FieldTableName0, product.FieldTableName1, product.FieldTableName2, product.FieldTableName3, product.FieldTableName4, product.FieldTableName5, product.FieldTableName6, product.FieldTableName7, product.FieldTableName8, product.FieldTableName9, product.FieldTableName10, product.FieldTableName11:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Product fields.
func (pr *Product) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case product.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case product.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case product.FieldTableName0:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field table_name0", values[i])
			} else if value.Valid {
				pr.TableName0 = value.String
			}
		case product.FieldTableName1:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field table_name1", values[i])
			} else if value.Valid {
				pr.TableName1 = value.String
			}
		case product.FieldTableName2:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field table_name2", values[i])
			} else if value.Valid {
				pr.TableName2 = value.String
			}
		case product.FieldTableName3:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field table_name3", values[i])
			} else if value.Valid {
				pr.TableName3 = value.String
			}
		case product.FieldTableName4:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field table_name4", values[i])
			} else if value.Valid {
				pr.TableName4 = value.String
			}
		case product.FieldTableName5:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field table_name5", values[i])
			} else if value.Valid {
				pr.TableName5 = value.String
			}
		case product.FieldTableName6:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field table_name6", values[i])
			} else if value.Valid {
				pr.TableName6 = value.String
			}
		case product.FieldTableName7:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field table_name7", values[i])
			} else if value.Valid {
				pr.TableName7 = value.String
			}
		case product.FieldTableName8:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field table_name8", values[i])
			} else if value.Valid {
				pr.TableName8 = value.String
			}
		case product.FieldTableName9:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field table_name9", values[i])
			} else if value.Valid {
				pr.TableName9 = value.String
			}
		case product.FieldTableName10:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field table_name10", values[i])
			} else if value.Valid {
				pr.TableName10 = value.String
			}
		case product.FieldTableName11:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field table_name11", values[i])
			} else if value.Valid {
				pr.TableName11 = value.String
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Product.
// This includes values selected through modifiers, order, etc.
func (pr *Product) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// Update returns a builder for updating this Product.
// Note that you need to call Product.Unwrap() before calling this method if this Product
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Product) Update() *ProductUpdateOne {
	return NewProductClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Product entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Product) Unwrap() *Product {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Product is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Product) String() string {
	var builder strings.Builder
	builder.WriteString("Product(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("table_name0=")
	builder.WriteString(pr.TableName0)
	builder.WriteString(", ")
	builder.WriteString("table_name1=")
	builder.WriteString(pr.TableName1)
	builder.WriteString(", ")
	builder.WriteString("table_name2=")
	builder.WriteString(pr.TableName2)
	builder.WriteString(", ")
	builder.WriteString("table_name3=")
	builder.WriteString(pr.TableName3)
	builder.WriteString(", ")
	builder.WriteString("table_name4=")
	builder.WriteString(pr.TableName4)
	builder.WriteString(", ")
	builder.WriteString("table_name5=")
	builder.WriteString(pr.TableName5)
	builder.WriteString(", ")
	builder.WriteString("table_name6=")
	builder.WriteString(pr.TableName6)
	builder.WriteString(", ")
	builder.WriteString("table_name7=")
	builder.WriteString(pr.TableName7)
	builder.WriteString(", ")
	builder.WriteString("table_name8=")
	builder.WriteString(pr.TableName8)
	builder.WriteString(", ")
	builder.WriteString("table_name9=")
	builder.WriteString(pr.TableName9)
	builder.WriteString(", ")
	builder.WriteString("table_name10=")
	builder.WriteString(pr.TableName10)
	builder.WriteString(", ")
	builder.WriteString("table_name11=")
	builder.WriteString(pr.TableName11)
	builder.WriteByte(')')
	return builder.String()
}

// Products is a parsable slice of Product.
type Products []*Product
