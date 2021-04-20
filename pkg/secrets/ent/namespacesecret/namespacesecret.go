// Code generated by entc, DO NOT EDIT.

package namespacesecret

const (
	// Label holds the string label denoting the namespacesecret type in the database.
	Label = "namespace_secret"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNs holds the string denoting the ns field in the database.
	FieldNs = "ns"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSecret holds the string denoting the secret field in the database.
	FieldSecret = "secret"
	// Table holds the table name of the namespacesecret in the database.
	Table = "namespace_secrets"
)

// Columns holds all SQL columns for namespacesecret fields.
var Columns = []string{
	FieldID,
	FieldNs,
	FieldName,
	FieldSecret,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}