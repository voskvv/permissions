package database

import (
	"context"
	"io"

	"git.containerum.net/ch/permissions/pkg/model"
	kubeClientModel "github.com/containerum/kube-client/pkg/model"
)

type AccessWithLabel struct {
	model.Permission `pg:",override"`

	Label string `sql:"label"`
}

type DB interface {
	UserAccesses(ctx context.Context, userID string) ([]AccessWithLabel, error)
	SetUserAccesses(ctx context.Context, userID string, level kubeClientModel.AccessLevel) error
	SetNamespaceAccess(ctx context.Context, ns model.Namespace, accessLevel kubeClientModel.AccessLevel, toUserID string) error
	DeleteNamespaceAccess(ctx context.Context, ns model.Namespace, userID string) error

	NamespaceByID(ctx context.Context, userID, id string) (ret model.NamespaceWithPermissions, err error)
	NamespacePermissions(ctx context.Context, ns *model.NamespaceWithPermissions) error
	UserNamespaces(ctx context.Context, userID string, filter NamespaceFilter) (ret []model.NamespaceWithPermissions, err error)
	AllNamespaces(ctx context.Context, filter NamespaceFilter) (ret []model.Namespace, err error)
	CreateNamespace(ctx context.Context, namespace *model.Namespace) error
	RenameNamespace(ctx context.Context, namespace *model.Namespace, newLabel string) error
	ResizeNamespace(ctx context.Context, namespace model.Namespace) error
	DeleteNamespace(ctx context.Context, namespace *model.Namespace) error
	DeleteAllUserNamespaces(ctx context.Context, userID string) (deleted []model.Namespace, err error)

	Transactional(fn func(tx DB) error) error

	io.Closer
}
