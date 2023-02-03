package worker

// Database Connection '''Pooling'''
// A concept where connections are reused instead of requiring each transaction to create a new connection.

// Idle Connection happens when database is no longer active.
// DB object is a pool of many database connections [ODBC, JDBC, RDBMS, etc] which contains both 'in-use' and 'idle' connections.
// A connection is marked as in-use when you are using it to perform a database task,
// such as executing a SQL statement or querying rows.
// When the task is complete the connection is marked as idle.
// https://techdocs.broadcom.com/us/en/symantec-security-software/identity-security/advanced-authentication/9-0/administrating/administrating-ca-strong-authentication/server-instances/configure-connectivity-settings/idle-database-connections.html

const DBMaxIdleConns = 3
const DBMaxConns = 50
const Workers = 100
