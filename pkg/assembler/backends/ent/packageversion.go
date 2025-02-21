// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packagename"
	"github.com/guacsec/guac/pkg/assembler/backends/ent/packageversion"
	"github.com/guacsec/guac/pkg/assembler/graphql/model"
)

// PackageVersion is the model entity for the PackageVersion schema.
type PackageVersion struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// NameID holds the value of the "name_id" field.
	NameID int `json:"name_id,omitempty"`
	// Version holds the value of the "version" field.
	Version string `json:"version,omitempty"`
	// Subpath holds the value of the "subpath" field.
	Subpath string `json:"subpath,omitempty"`
	// Qualifiers holds the value of the "qualifiers" field.
	Qualifiers []model.PackageQualifier `json:"qualifiers,omitempty"`
	// A SHA1 of the qualifiers, subpath, version fields after sorting keys, used to ensure uniqueness of version records.
	Hash string `json:"hash,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PackageVersionQuery when eager-loading is set.
	Edges        PackageVersionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PackageVersionEdges holds the relations/edges for other nodes in the graph.
type PackageVersionEdges struct {
	// Name holds the value of the name edge.
	Name *PackageName `json:"name,omitempty"`
	// Occurrences holds the value of the occurrences edge.
	Occurrences []*Occurrence `json:"occurrences,omitempty"`
	// Sbom holds the value of the sbom edge.
	Sbom []*BillOfMaterials `json:"sbom,omitempty"`
	// EqualPackages holds the value of the equal_packages edge.
	EqualPackages []*PkgEqual `json:"equal_packages,omitempty"`
	// IncludedInSboms holds the value of the included_in_sboms edge.
	IncludedInSboms []*BillOfMaterials `json:"included_in_sboms,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
	// totalCount holds the count of the edges above.
	totalCount [5]map[string]int

	namedOccurrences     map[string][]*Occurrence
	namedSbom            map[string][]*BillOfMaterials
	namedEqualPackages   map[string][]*PkgEqual
	namedIncludedInSboms map[string][]*BillOfMaterials
}

// NameOrErr returns the Name value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PackageVersionEdges) NameOrErr() (*PackageName, error) {
	if e.loadedTypes[0] {
		if e.Name == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: packagename.Label}
		}
		return e.Name, nil
	}
	return nil, &NotLoadedError{edge: "name"}
}

// OccurrencesOrErr returns the Occurrences value or an error if the edge
// was not loaded in eager-loading.
func (e PackageVersionEdges) OccurrencesOrErr() ([]*Occurrence, error) {
	if e.loadedTypes[1] {
		return e.Occurrences, nil
	}
	return nil, &NotLoadedError{edge: "occurrences"}
}

// SbomOrErr returns the Sbom value or an error if the edge
// was not loaded in eager-loading.
func (e PackageVersionEdges) SbomOrErr() ([]*BillOfMaterials, error) {
	if e.loadedTypes[2] {
		return e.Sbom, nil
	}
	return nil, &NotLoadedError{edge: "sbom"}
}

// EqualPackagesOrErr returns the EqualPackages value or an error if the edge
// was not loaded in eager-loading.
func (e PackageVersionEdges) EqualPackagesOrErr() ([]*PkgEqual, error) {
	if e.loadedTypes[3] {
		return e.EqualPackages, nil
	}
	return nil, &NotLoadedError{edge: "equal_packages"}
}

// IncludedInSbomsOrErr returns the IncludedInSboms value or an error if the edge
// was not loaded in eager-loading.
func (e PackageVersionEdges) IncludedInSbomsOrErr() ([]*BillOfMaterials, error) {
	if e.loadedTypes[4] {
		return e.IncludedInSboms, nil
	}
	return nil, &NotLoadedError{edge: "included_in_sboms"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PackageVersion) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case packageversion.FieldQualifiers:
			values[i] = new([]byte)
		case packageversion.FieldID, packageversion.FieldNameID:
			values[i] = new(sql.NullInt64)
		case packageversion.FieldVersion, packageversion.FieldSubpath, packageversion.FieldHash:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PackageVersion fields.
func (pv *PackageVersion) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case packageversion.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pv.ID = int(value.Int64)
		case packageversion.FieldNameID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field name_id", values[i])
			} else if value.Valid {
				pv.NameID = int(value.Int64)
			}
		case packageversion.FieldVersion:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field version", values[i])
			} else if value.Valid {
				pv.Version = value.String
			}
		case packageversion.FieldSubpath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subpath", values[i])
			} else if value.Valid {
				pv.Subpath = value.String
			}
		case packageversion.FieldQualifiers:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field qualifiers", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pv.Qualifiers); err != nil {
					return fmt.Errorf("unmarshal field qualifiers: %w", err)
				}
			}
		case packageversion.FieldHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hash", values[i])
			} else if value.Valid {
				pv.Hash = value.String
			}
		default:
			pv.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PackageVersion.
// This includes values selected through modifiers, order, etc.
func (pv *PackageVersion) Value(name string) (ent.Value, error) {
	return pv.selectValues.Get(name)
}

// QueryName queries the "name" edge of the PackageVersion entity.
func (pv *PackageVersion) QueryName() *PackageNameQuery {
	return NewPackageVersionClient(pv.config).QueryName(pv)
}

// QueryOccurrences queries the "occurrences" edge of the PackageVersion entity.
func (pv *PackageVersion) QueryOccurrences() *OccurrenceQuery {
	return NewPackageVersionClient(pv.config).QueryOccurrences(pv)
}

// QuerySbom queries the "sbom" edge of the PackageVersion entity.
func (pv *PackageVersion) QuerySbom() *BillOfMaterialsQuery {
	return NewPackageVersionClient(pv.config).QuerySbom(pv)
}

// QueryEqualPackages queries the "equal_packages" edge of the PackageVersion entity.
func (pv *PackageVersion) QueryEqualPackages() *PkgEqualQuery {
	return NewPackageVersionClient(pv.config).QueryEqualPackages(pv)
}

// QueryIncludedInSboms queries the "included_in_sboms" edge of the PackageVersion entity.
func (pv *PackageVersion) QueryIncludedInSboms() *BillOfMaterialsQuery {
	return NewPackageVersionClient(pv.config).QueryIncludedInSboms(pv)
}

// Update returns a builder for updating this PackageVersion.
// Note that you need to call PackageVersion.Unwrap() before calling this method if this PackageVersion
// was returned from a transaction, and the transaction was committed or rolled back.
func (pv *PackageVersion) Update() *PackageVersionUpdateOne {
	return NewPackageVersionClient(pv.config).UpdateOne(pv)
}

// Unwrap unwraps the PackageVersion entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pv *PackageVersion) Unwrap() *PackageVersion {
	_tx, ok := pv.config.driver.(*txDriver)
	if !ok {
		panic("ent: PackageVersion is not a transactional entity")
	}
	pv.config.driver = _tx.drv
	return pv
}

// String implements the fmt.Stringer.
func (pv *PackageVersion) String() string {
	var builder strings.Builder
	builder.WriteString("PackageVersion(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pv.ID))
	builder.WriteString("name_id=")
	builder.WriteString(fmt.Sprintf("%v", pv.NameID))
	builder.WriteString(", ")
	builder.WriteString("version=")
	builder.WriteString(pv.Version)
	builder.WriteString(", ")
	builder.WriteString("subpath=")
	builder.WriteString(pv.Subpath)
	builder.WriteString(", ")
	builder.WriteString("qualifiers=")
	builder.WriteString(fmt.Sprintf("%v", pv.Qualifiers))
	builder.WriteString(", ")
	builder.WriteString("hash=")
	builder.WriteString(pv.Hash)
	builder.WriteByte(')')
	return builder.String()
}

// NamedOccurrences returns the Occurrences named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pv *PackageVersion) NamedOccurrences(name string) ([]*Occurrence, error) {
	if pv.Edges.namedOccurrences == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pv.Edges.namedOccurrences[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pv *PackageVersion) appendNamedOccurrences(name string, edges ...*Occurrence) {
	if pv.Edges.namedOccurrences == nil {
		pv.Edges.namedOccurrences = make(map[string][]*Occurrence)
	}
	if len(edges) == 0 {
		pv.Edges.namedOccurrences[name] = []*Occurrence{}
	} else {
		pv.Edges.namedOccurrences[name] = append(pv.Edges.namedOccurrences[name], edges...)
	}
}

// NamedSbom returns the Sbom named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pv *PackageVersion) NamedSbom(name string) ([]*BillOfMaterials, error) {
	if pv.Edges.namedSbom == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pv.Edges.namedSbom[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pv *PackageVersion) appendNamedSbom(name string, edges ...*BillOfMaterials) {
	if pv.Edges.namedSbom == nil {
		pv.Edges.namedSbom = make(map[string][]*BillOfMaterials)
	}
	if len(edges) == 0 {
		pv.Edges.namedSbom[name] = []*BillOfMaterials{}
	} else {
		pv.Edges.namedSbom[name] = append(pv.Edges.namedSbom[name], edges...)
	}
}

// NamedEqualPackages returns the EqualPackages named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pv *PackageVersion) NamedEqualPackages(name string) ([]*PkgEqual, error) {
	if pv.Edges.namedEqualPackages == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pv.Edges.namedEqualPackages[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pv *PackageVersion) appendNamedEqualPackages(name string, edges ...*PkgEqual) {
	if pv.Edges.namedEqualPackages == nil {
		pv.Edges.namedEqualPackages = make(map[string][]*PkgEqual)
	}
	if len(edges) == 0 {
		pv.Edges.namedEqualPackages[name] = []*PkgEqual{}
	} else {
		pv.Edges.namedEqualPackages[name] = append(pv.Edges.namedEqualPackages[name], edges...)
	}
}

// NamedIncludedInSboms returns the IncludedInSboms named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pv *PackageVersion) NamedIncludedInSboms(name string) ([]*BillOfMaterials, error) {
	if pv.Edges.namedIncludedInSboms == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pv.Edges.namedIncludedInSboms[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pv *PackageVersion) appendNamedIncludedInSboms(name string, edges ...*BillOfMaterials) {
	if pv.Edges.namedIncludedInSboms == nil {
		pv.Edges.namedIncludedInSboms = make(map[string][]*BillOfMaterials)
	}
	if len(edges) == 0 {
		pv.Edges.namedIncludedInSboms[name] = []*BillOfMaterials{}
	} else {
		pv.Edges.namedIncludedInSboms[name] = append(pv.Edges.namedIncludedInSboms[name], edges...)
	}
}

// PackageVersions is a parsable slice of PackageVersion.
type PackageVersions []*PackageVersion
