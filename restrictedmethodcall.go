package sqlpassctxcheck

type RestrictedMethodCall struct {
	ReceiverType        string
	MethodName          string
	AlternateMethodName string
}

type RestrictedMethodCallList []RestrictedMethodCall

func (r RestrictedMethodCallList) Get(receiverType, methodName string) (*RestrictedMethodCall, bool) {
	for _, restrictedMethodCall := range r {
		if restrictedMethodCall.ReceiverType == receiverType &&
			restrictedMethodCall.MethodName == methodName {
			return &restrictedMethodCall, true
		}
	}
	return nil, false
}

var SqlModuleRestrictedMethodMap = map[SqlPackage]RestrictedMethodCallList{
	SQL: {
		{
			ReceiverType:        "DB",
			MethodName:          "Begin",
			AlternateMethodName: "BeginTx",
		},
		{
			ReceiverType:        "DB",
			MethodName:          "Exec",
			AlternateMethodName: "ExecContext",
		},
		{
			ReceiverType:        "DB",
			MethodName:          "Ping",
			AlternateMethodName: "PingContext",
		},
		{
			ReceiverType:        "DB",
			MethodName:          "Prepare",
			AlternateMethodName: "PrepareContext",
		},
		{
			ReceiverType:        "DB",
			MethodName:          "Query",
			AlternateMethodName: "QueryContext",
		},
		{
			ReceiverType:        "DB",
			MethodName:          "QueryRow",
			AlternateMethodName: "QueryRowContext",
		},
		{
			ReceiverType:        "Stmt",
			MethodName:          "Exec",
			AlternateMethodName: "ExecContext",
		},
		{
			ReceiverType:        "Stmt",
			MethodName:          "Query",
			AlternateMethodName: "QueryContext",
		},
		{
			ReceiverType:        "Stmt",
			MethodName:          "QueryRow",
			AlternateMethodName: "QueryRowContext",
		},
		{
			ReceiverType:        "Tx",
			MethodName:          "Exec",
			AlternateMethodName: "ExecContext",
		},
		{
			ReceiverType:        "Tx",
			MethodName:          "Prepare",
			AlternateMethodName: "PrepareContext",
		},
		{
			ReceiverType:        "Tx",
			MethodName:          "Query",
			AlternateMethodName: "QueryContext",
		},
		{
			ReceiverType:        "Tx",
			MethodName:          "QueryRow",
			AlternateMethodName: "QueryRowContext",
		},
		{
			ReceiverType:        "Tx",
			MethodName:          "Stmt",
			AlternateMethodName: "StmtContext",
		},
	},
	SQLX: {
		{
			ReceiverType:        "DB",
			MethodName:          "Beginx",
			AlternateMethodName: "BeginTxx",
		},
		{
			ReceiverType:        "DB",
			MethodName:          "Get",
			AlternateMethodName: "GetContext",
		},
		{
			ReceiverType:        "DB",
			MethodName:          "MustBegin",
			AlternateMethodName: "MustBeginTx",
		},
		{
			ReceiverType:        "DB",
			MethodName:          "MustExec",
			AlternateMethodName: "MustExecContext",
		},
		{
			ReceiverType:        "DB",
			MethodName:          "NamedExec",
			AlternateMethodName: "NamedExecContext",
		},
		{
			ReceiverType:        "DB",
			MethodName:          "NamedQuery",
			AlternateMethodName: "NamedQueryContext",
		},
		{
			ReceiverType:        "DB",
			MethodName:          "PrepareNamed",
			AlternateMethodName: "PrepareNamedContext",
		},
		{
			ReceiverType:        "DB",
			MethodName:          "Preparex",
			AlternateMethodName: "PreparexContext",
		},
		{
			ReceiverType:        "DB",
			MethodName:          "QueryRowx",
			AlternateMethodName: "QueryRowxContext",
		},
		{
			ReceiverType:        "DB",
			MethodName:          "Queryx",
			AlternateMethodName: "QueryxContext",
		},
		{
			ReceiverType:        "DB",
			MethodName:          "Select",
			AlternateMethodName: "SelectContext",
		},
		{
			ReceiverType:        "NamedStmt",
			MethodName:          "Exec",
			AlternateMethodName: "ExecContext",
		},
		{
			ReceiverType:        "NamedStmt",
			MethodName:          "Get",
			AlternateMethodName: "GetContext",
		},
		{
			ReceiverType:        "NamedStmt",
			MethodName:          "MustExec",
			AlternateMethodName: "MustExecContext",
		},
		{
			ReceiverType:        "NamedStmt",
			MethodName:          "Query",
			AlternateMethodName: "QueryContext",
		},
		{
			ReceiverType:        "NamedStmt",
			MethodName:          "QueryRow",
			AlternateMethodName: "QueryRowContext",
		},
		{
			ReceiverType:        "NamedStmt",
			MethodName:          "QueryRowx",
			AlternateMethodName: "QueryRowxContext",
		},
		{
			ReceiverType:        "NamedStmt",
			MethodName:          "Queryx",
			AlternateMethodName: "QueryxContext",
		},
		{
			ReceiverType:        "NamedStmt",
			MethodName:          "Select",
			AlternateMethodName: "SelectContext",
		},
		{
			ReceiverType:        "Stmt",
			MethodName:          "Get",
			AlternateMethodName: "GetContext",
		},
		{
			ReceiverType:        "Stmt",
			MethodName:          "MustExec",
			AlternateMethodName: "MustExecContext",
		},
		{
			ReceiverType:        "Stmt",
			MethodName:          "QueryRowx",
			AlternateMethodName: "QueryRowxContext",
		},
		{
			ReceiverType:        "Stmt",
			MethodName:          "Queryx",
			AlternateMethodName: "QueryxContext",
		},
		{
			ReceiverType:        "Stmt",
			MethodName:          "Select",
			AlternateMethodName: "SelectContext",
		},
		{
			ReceiverType:        "Tx",
			MethodName:          "Get",
			AlternateMethodName: "GetContext",
		},
		{
			ReceiverType:        "Tx",
			MethodName:          "MustExec",
			AlternateMethodName: "MustExecContext",
		},
		{
			ReceiverType:        "Tx",
			MethodName:          "NamedExec",
			AlternateMethodName: "NamedExecContext",
		},
		{
			ReceiverType:        "Tx",
			MethodName:          "NamedStmt",
			AlternateMethodName: "NamedStmtContext",
		},
		{
			ReceiverType:        "Tx",
			MethodName:          "PrepareNamed",
			AlternateMethodName: "PrepareNamedContext",
		},
		{
			ReceiverType:        "Tx",
			MethodName:          "Preparex",
			AlternateMethodName: "PreparexContext",
		},
		{
			ReceiverType:        "Tx",
			MethodName:          "QueryRowx",
			AlternateMethodName: "QueryRowxContext",
		},
		{
			ReceiverType:        "Tx",
			MethodName:          "Queryx",
			AlternateMethodName: "QueryxContext",
		},
		{
			ReceiverType:        "Tx",
			MethodName:          "Select",
			AlternateMethodName: "SelectContext",
		},
		{
			ReceiverType:        "Tx",
			MethodName:          "Stmtx",
			AlternateMethodName: "StmtxContext",
		},
	},
}
