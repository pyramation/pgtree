package nodes

// ObjectTypeKeyword maps the ObjectType enum to the SQL name keyword.
var ObjectTypeKeyword = map[ObjectType]string{
	OBJECT_ACCESS_METHOD:   "ACCESS METHOD",
	OBJECT_AGGREGATE:       "AGGREGATE",
	OBJECT_AMOP:            "AMOP",
	OBJECT_AMPROC:          "AMPROC",
	OBJECT_ATTRIBUTE:       "ATTRIBUTE",
	OBJECT_CAST:            "CAST",
	OBJECT_COLUMN:          "COLUMN",
	OBJECT_COLLATION:       "COLLATION",
	OBJECT_CONVERSION:      "CONVERSION",
	OBJECT_DATABASE:        "DATABASE",
	OBJECT_DEFAULT:         "DEFAULT",
	OBJECT_DEFACL:          "DEFAULT ACL",
	OBJECT_DOMAIN:          "DOMAIN",
	OBJECT_DOMCONSTRAINT:   "DOMCONSTRAINT",
	OBJECT_EVENT_TRIGGER:   "EVENT TRIGGER",
	OBJECT_EXTENSION:       "EXTENSION",
	OBJECT_FDW:             "FOREIGN DATA WRAPPER",
	OBJECT_FOREIGN_SERVER:  "SERVER",
	OBJECT_FOREIGN_TABLE:   "FOREIGN TABLE",
	OBJECT_FUNCTION:        "FUNCTION",
	OBJECT_INDEX:           "INDEX",
	OBJECT_LANGUAGE:        "LANGUAGE",
	OBJECT_LARGEOBJECT:     "LARGEOBJECT",
	OBJECT_MATVIEW:         "MATERIALIZED VIEW",
	OBJECT_OPCLASS:         "OPERATOR CLASS",
	OBJECT_OPERATOR:        "OPERATOR",
	OBJECT_OPFAMILY:        "OPERATOR FAMILY",
	OBJECT_POLICY:          "POLICY",
	OBJECT_PROCEDURE:       "PROCEDURE",
	OBJECT_PUBLICATION:     "PUBLICATION",
	OBJECT_PUBLICATION_REL: "OBJECT_PUBLICATION_REL",
	OBJECT_ROLE:            "ROLE",
	OBJECT_ROUTINE:         "ROUTINE",
	OBJECT_RULE:            "RULE",
	OBJECT_SCHEMA:          "SCHEMA",
	OBJECT_SEQUENCE:        "SEQUENCE",
	OBJECT_SUBSCRIPTION:    "SUBSCRIPTION",
	OBJECT_STATISTIC_EXT:   "STATISTIC",
	OBJECT_TABCONSTRAINT:   "TABLE CONSTRAINT",
	OBJECT_TABLE:           "TABLE",
	OBJECT_TABLESPACE:      "TABLESPACE",
	OBJECT_TRANSFORM:       "TRANSFORM",
	OBJECT_TRIGGER:         "TRIGGER",
	OBJECT_TSCONFIGURATION: "TEXT SEARCH CONFIGURATION",
	OBJECT_TSDICTIONARY:    "TEXT SEARCH DICTIONARY",
	OBJECT_TSPARSER:        "TEXT SEARCH PARSER",
	OBJECT_TSTEMPLATE:      "TEXT SEARCH TEMPLATE",
	OBJECT_TYPE:            "TYPE",
	OBJECT_USER_MAPPING:    "USER MAPPING",
	OBJECT_VIEW:            "VIEW",
}

// SQLValueFunctionOpName maps SQLValueFunctionOp to sql standard function name.
var SQLValueFunctionOpName = map[SQLValueFunctionOp]string{
	SVFOP_CURRENT_DATE:        "current_date",
	SVFOP_CURRENT_TIME:        "current_time",
	SVFOP_CURRENT_TIME_N:      "current_time",
	SVFOP_CURRENT_TIMESTAMP:   "current_timestamp",
	SVFOP_CURRENT_TIMESTAMP_N: "current_timestamp",
	SVFOP_LOCALTIME:           "localtime",
	SVFOP_LOCALTIME_N:         "localtime",
	SVFOP_LOCALTIMESTAMP:      "localtimestamp",
	SVFOP_LOCALTIMESTAMP_N:    "localtimestamp",
	SVFOP_CURRENT_ROLE:        "current_role",
	SVFOP_CURRENT_USER:        "current_user",
	SVFOP_USER:                "user",
	SVFOP_SESSION_USER:        "session_user",
	SVFOP_CURRENT_CATALOG:     "current_catalog",
	SVFOP_CURRENT_SCHEMA:      "current_schema",
}

// LockClauseStrengthKeyword maps LockClauseStrength enums to sql keyword.
var LockClauseStrengthKeyword = map[LockClauseStrength]string{
	LCS_NONE:           "",
	LCS_FORKEYSHARE:    "KEY SHARE",
	LCS_FORSHARE:       "SHARE",
	LCS_FORNOKEYUPDATE: "NO KEY UPDATE",
	LCS_FORUPDATE:      "UPDATE",
}

// LockMode enumeration of lock modes.
type LockMode uint8

// LockModes.
const (
	LockModeAccessShare          LockMode = 1
	LockModeRowShare             LockMode = 2
	LockModeRowExclusive         LockMode = 3
	LockModeShareUpdateExclusive LockMode = 4
	LockModeShare                LockMode = 5
	LockModeShareRowExclusive    LockMode = 6
	LockModeExclusive            LockMode = 7
	LockModeAccessExclusive      LockMode = 8
)

// LockModeKeyword maps LockMode enums to sql keyword.
var LockModeKeyword = map[LockMode]string{
	LockModeAccessShare:          "IN ACCESS SHARE MODE",
	LockModeRowShare:             "IN ROW SHARE MODE",
	LockModeRowExclusive:         "IN ROW EXCLUSIVE MODE",
	LockModeShareUpdateExclusive: "IN SHARE UPDATE EXCLUSIVE MODE",
	LockModeShare:                "IN SHARE MODE",
	LockModeShareRowExclusive:    "IN SHARE ROW EXCLUSIVE MODE",
	LockModeExclusive:            "IN EXCLUSIVE MODE",
	LockModeAccessExclusive:      "IN ACCESS EXCLUSIVE MODE",
}

// CmdTypeKeyword maps CmdType enums to sql keyword.
var CmdTypeKeyword = map[CmdType]string{
	CMD_SELECT: "SELECT",
	CMD_UPDATE: "UPDATE",
	CMD_INSERT: "INSERT",
	CMD_DELETE: "DELETE",
}

// SetOpKeyword maps SetOperation enums to sql keyword.
var SetOpKeyword = map[SetOperation]string{
	SETOP_UNION:     "UNION",
	SETOP_INTERSECT: "INTERSECT",
	SETOP_EXCEPT:    "EXCEPT",
}

// ConstrTypeKeyword maps ConstrType enums to sql keyword.
var ConstrTypeKeyword = map[ConstrType]string{
	CONSTR_NULL:      "NULL",
	CONSTR_NOTNULL:   "NOT NULL",
	CONSTR_DEFAULT:   "DEFAULT",
	CONSTR_CHECK:     "CHECK",
	CONSTR_PRIMARY:   "PRIMARY KEY",
	CONSTR_UNIQUE:    "UNIQUE",
	CONSTR_EXCLUSION: "EXCLUDE",
	CONSTR_FOREIGN:   "FOREIGN KEY",
	CONSTR_GENERATED: "GENERATED",
	CONSTR_IDENTITY:  "GENERATED",
}

// AlterTableCommand maps AlterTableType enums to sql commands.
var AlterTableCommand = map[AlterTableType]string{
	AT_AddColumn:                 "ADD",
	AT_ColumnDefault:             "ALTER",
	AT_DropNotNull:               "ALTER",
	AT_SetNotNull:                "ALTER",
	AT_SetStatistics:             "ALTER",
	AT_SetOptions:                "ALTER",
	AT_ResetOptions:              "ALTER",
	AT_SetStorage:                "ALTER",
	AT_DropColumn:                "DROP",
	AT_AddIndex:                  "ADD INDEX",
	AT_AddConstraint:             "ADD",
	AT_AlterConstraint:           "ALTER CONSTRAINT",
	AT_ValidateConstraint:        "VALIDATE CONSTRAINT",
	AT_DropConstraint:            "DROP CONSTRAINT",
	AT_AlterColumnType:           "ALTER",
	AT_AlterColumnGenericOptions: "ALTER",
	AT_ChangeOwner:               "OWNER TO",
	AT_SetRelOptions:             "SET",
	AT_ResetRelOptions:           "RESET",
}

// AlterTableOption maps AlterTableType enums to sql command options.
var AlterTableOption = map[AlterTableType]string{
	AT_ColumnDefault:             "SET DEFAULT",
	AT_DropNotNull:               "DROP NOT NULL",
	AT_SetNotNull:                "SET NOT NULL",
	AT_SetStatistics:             "SET STATISTICS",
	AT_SetOptions:                "SET",
	AT_ResetOptions:              "RESET",
	AT_SetStorage:                "SET STORAGE",
	AT_AlterColumnType:           "TYPE",
	AT_AlterColumnGenericOptions: "OPTIONS",
}

// Defines the operator types used in Op Class definitions.
const (
	OperatorItemType = 1
	FunctionItemType = 2
)

// PgTypeNameToKeyword maps internal type names to sql names.
var PgTypeNameToKeyword = map[string]string{
	"bool":        "boolean",
	"int2":        "smallint",
	"int4":        "int",
	"int8":        "bigint",
	"real":        "real",
	"float4":      "real",
	"float8":      "double precision",
	"time":        "time",
	"timetz":      "time with time zone",
	"timestamp":   "timestamp",
	"timestamptz": "timestamp with time zone",
	"interval":    "interval",
}

// ContraintGeneratedWhenToKeyword maps Constraint GeneratedWhen clauses to keywords.
var ContraintGeneratedWhenToKeyword = map[string]string{
	"a": "ALWAYS",
	"d": "BY DEFAULT",
}
