package db

import "github.com/hashicorp/go-memdb"

func GetDbSchema() *memdb.DBSchema {
	return &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"endPoint": &memdb.TableSchema{
				Name: "endPoint",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Id"},
					},
				},
			},
			"health": &memdb.TableSchema{
				Name: "health",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "EndPointId"},
					},
				},
			},
		},
	}
}
