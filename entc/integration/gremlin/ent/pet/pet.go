// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package pet

const (
	// Label holds the string label denoting the pet type in the database.
	Label = "pet"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"

	// TeamInverseLabel holds the string label denoting the team inverse edge type in the database.
	TeamInverseLabel = "user_team"
	// OwnerInverseLabel holds the string label denoting the owner inverse edge type in the database.
	OwnerInverseLabel = "user_pets"
)