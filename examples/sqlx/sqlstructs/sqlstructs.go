// Code generated by sqlcodegen, DO NOT EDIT.

package sqlstructs

/*
	SELECT id FROM users LIMIT 5
*/
type IdInt struct {
	Id int `db:"id" json:"id"`
}

/*
	SELECT firebase_auth_uid FROM users LIMIT 5
*/
type Foo struct {
	FirebaseAuthUid string `db:"firebase_auth_uid" json:"firebase_auth_uid"`
}

/*
	SELECT *
	FROM
	users
	LIMIT 5
*/
type IdIntFirebaseAuthUidString struct {
	Id              int    `db:"id" json:"id"`
	FirebaseAuthUid string `db:"firebase_auth_uid" json:"firebase_auth_uid"`
}
