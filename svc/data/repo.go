package data

import (
	"fmt"

	"github.com/jasontconnell/repository/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PackageRepo interface {
	GetPackage() Package
	SavePackage(pkg *Package) error
}

type PackageMongoRepo struct {
	repository.Repository[Package]
}

func (r *PackageMongoRepo) GetPackage() Package {
	list, err := r.GetAll("packages")
	if err != nil || len(list) == 0 {
		return Package{}
	}

	return list[0]
}

func (r *PackageMongoRepo) SavePackage(pkg *Package) error {
	id, err := r.Save("packages", *pkg)
	if err != nil {
		return fmt.Errorf("couldn't save package %w", err)
	}
	pkg.id = id
	return nil
}

type Package struct {
	id      primitive.ObjectID `bson:"_id"`
	Message string             `bson:"message"`
}

func (p Package) GetId() primitive.ObjectID {
	return p.id
}

func (p *Package) SetId(id primitive.ObjectID) {
	p.id = id
}
