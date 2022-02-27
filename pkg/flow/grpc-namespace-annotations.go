package flow

import (
	"bytes"
	"context"
	"errors"
	"io"
	"time"

	"github.com/direktiv/direktiv/pkg/flow/ent"
	entnote "github.com/direktiv/direktiv/pkg/flow/ent/annotation"
	"github.com/direktiv/direktiv/pkg/flow/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (flow *flow) NamespaceAnnotation(ctx context.Context, req *grpc.NamespaceAnnotationRequest) (*grpc.NamespaceAnnotationResponse, error) {

	flow.sugar.Debugf("Handling gRPC request: %s", this())

	nsc := flow.db.Namespace

	d, err := flow.traverseToNamespaceAnnotation(ctx, nsc, req.GetNamespace(), req.GetKey())
	if err != nil {
		return nil, err
	}

	var resp grpc.NamespaceAnnotationResponse

	resp.Namespace = d.ns().Name
	resp.Key = d.annotation.Name
	resp.CreatedAt = timestamppb.New(d.annotation.CreatedAt)
	resp.UpdatedAt = timestamppb.New(d.annotation.UpdatedAt)
	resp.Checksum = d.annotation.Hash
	resp.TotalSize = int64(d.annotation.Size)

	if resp.TotalSize > parcelSize {
		return nil, status.Error(codes.ResourceExhausted, "annotation too large to return without using the parcelling API")
	}

	resp.Data = d.annotation.Data

	return &resp, nil

}

func (flow *flow) NamespaceAnnotationParcels(req *grpc.NamespaceAnnotationRequest, srv grpc.Flow_NamespaceAnnotationParcelsServer) error {

	flow.sugar.Debugf("Handling gRPC request: %s", this())

	ctx := srv.Context()

	nsc := flow.db.Namespace

	d, err := flow.traverseToNamespaceAnnotation(ctx, nsc, req.GetNamespace(), req.GetKey())
	if err != nil {
		return err
	}

	rdr := bytes.NewReader(d.annotation.Data)

	for {

		resp := new(grpc.NamespaceAnnotationResponse)

		resp.Namespace = d.ns().Name
		resp.Key = d.annotation.Name
		resp.CreatedAt = timestamppb.New(d.annotation.CreatedAt)
		resp.UpdatedAt = timestamppb.New(d.annotation.UpdatedAt)
		resp.Checksum = d.annotation.Hash
		resp.TotalSize = int64(d.annotation.Size)

		buf := new(bytes.Buffer)
		k, err := io.CopyN(buf, rdr, parcelSize)
		if err != nil {

			if errors.Is(err, io.EOF) {
				err = nil
			}

			if err == nil && k == 0 {
				if resp.TotalSize == 0 {
					resp.Data = buf.Bytes()
					err = srv.Send(resp)
					if err != nil {
						return err
					}
				}
				return nil
			}

			if err != nil {
				return err
			}

		}

		resp.Data = buf.Bytes()

		err = srv.Send(resp)
		if err != nil {
			return err
		}

	}

}

func annotationsOrder(p *pagination) ent.AnnotationPaginateOption {

	field := ent.AnnotationOrderFieldName
	direction := ent.OrderDirectionAsc

	if p.order != nil {

		if x := p.order.Field; x != "" && x == "NAME" {
			field = ent.AnnotationOrderFieldName
		}

		if x := p.order.Direction; x != "" && x == "DESC" {
			direction = ent.OrderDirectionDesc
		}

	}

	return ent.WithAnnotationOrder(&ent.AnnotationOrder{
		Direction: direction,
		Field:     field,
	})

}

func annotationsFilter(p *pagination) ent.AnnotationPaginateOption {

	if p.filter == nil {
		return nil
	}

	filter := p.filter.Val

	return ent.WithAnnotationFilter(func(query *ent.AnnotationQuery) (*ent.AnnotationQuery, error) {

		if filter == "" {
			return query, nil
		}

		field := p.filter.Field
		if field == "" {
			return query, nil
		}

		switch field {
		case "NAME":

			ftype := p.filter.Type
			if ftype == "" {
				return query, nil
			}

			switch ftype {
			case "CONTAINS":
				return query.Where(entnote.NameContains(filter)), nil
			}
		}

		return query, nil

	})

}

func (flow *flow) NamespaceAnnotations(ctx context.Context, req *grpc.NamespaceAnnotationsRequest) (*grpc.NamespaceAnnotationsResponse, error) {

	flow.sugar.Debugf("Handling gRPC request: %s", this())

	p, err := getPagination(req.Pagination)
	if err != nil {
		return nil, err
	}

	opts := []ent.AnnotationPaginateOption{}
	opts = append(opts, annotationsOrder(p))
	filter := annotationsFilter(p)
	if filter != nil {
		opts = append(opts, filter)
	}

	nsc := flow.db.Namespace
	ns, err := flow.getNamespace(ctx, nsc, req.GetNamespace())
	if err != nil {
		return nil, err
	}

	query := ns.QueryAnnotations()
	cx, err := query.Paginate(ctx, p.After(), p.First(), p.Before(), p.Last(), opts...)
	if err != nil {
		return nil, err
	}

	var resp grpc.NamespaceAnnotationsResponse

	resp.Namespace = ns.Name

	err = atob(cx, &resp.Annotations)
	if err != nil {
		return nil, err
	}

	for i := range cx.Edges {

		edge := cx.Edges[i]
		annotation := edge.Node

		v := resp.Annotations.Edges[i].Node
		v.Checksum = annotation.Hash
		v.CreatedAt = timestamppb.New(annotation.CreatedAt)
		v.Size = int64(annotation.Size)
		v.UpdatedAt = timestamppb.New(annotation.UpdatedAt)

	}

	return &resp, nil

}

func (flow *flow) NamespaceAnnotationsStream(req *grpc.NamespaceAnnotationsRequest, srv grpc.Flow_NamespaceAnnotationsStreamServer) error {

	flow.sugar.Debugf("Handling gRPC request: %s", this())

	ctx := srv.Context()
	phash := ""
	nhash := ""

	p, err := getPagination(req.Pagination)
	if err != nil {
		return err
	}

	opts := []ent.AnnotationPaginateOption{}
	opts = append(opts, annotationsOrder(p))
	filter := annotationsFilter(p)
	if filter != nil {
		opts = append(opts, filter)
	}

	nsc := flow.db.Namespace
	ns, err := flow.getNamespace(ctx, nsc, req.GetNamespace())
	if err != nil {
		return err
	}

	sub := flow.pubsub.SubscribeNamespaceAnnotations(ns)
	defer flow.cleanup(sub.Close)

resend:

	query := ns.QueryAnnotations()
	cx, err := query.Paginate(ctx, p.After(), p.First(), p.Before(), p.Last(), opts...)
	if err != nil {
		return err
	}

	resp := new(grpc.NamespaceAnnotationsResponse)

	resp.Namespace = ns.Name

	err = atob(cx, &resp.Annotations)
	if err != nil {
		return err
	}

	for i := range cx.Edges {

		edge := cx.Edges[i]
		annotation := edge.Node

		v := resp.Annotations.Edges[i].Node
		v.Checksum = annotation.Hash
		v.CreatedAt = timestamppb.New(annotation.CreatedAt)
		v.Size = int64(annotation.Size)
		v.UpdatedAt = timestamppb.New(annotation.UpdatedAt)

	}

	nhash = checksum(resp)
	if nhash != phash {
		err = srv.Send(resp)
		if err != nil {
			return err
		}
	}
	phash = nhash

	more := sub.Wait(ctx)
	if !more {
		return nil
	}

	goto resend

}

func (flow *flow) SetNamespaceAnnotation(ctx context.Context, req *grpc.SetNamespaceAnnotationRequest) (*grpc.SetNamespaceAnnotationResponse, error) {

	flow.sugar.Debugf("Handling gRPC request: %s", this())

	tx, err := flow.db.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer rollback(tx)

	nsc := tx.Namespace
	annotationc := tx.Annotation

	ns, err := flow.getNamespace(ctx, nsc, req.GetNamespace())
	if err != nil {
		return nil, err
	}

	var annotation *ent.Annotation

	key := req.GetKey()

	var newAnnotation bool
	annotation, newAnnotation, err = flow.SetAnnotation(ctx, annotationc, ns, key, req.GetData())
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	if newAnnotation {
		flow.logToNamespace(ctx, time.Now(), ns, "Created namespace annotation '%s'.", key)
	} else {
		flow.logToNamespace(ctx, time.Now(), ns, "Updated namespace annotation '%s'.", key)
	}

	flow.pubsub.NotifyNamespaceAnnotations(ns)

	var resp grpc.SetNamespaceAnnotationResponse

	resp.Namespace = ns.Name
	resp.Key = key
	resp.CreatedAt = timestamppb.New(annotation.CreatedAt)
	resp.UpdatedAt = timestamppb.New(annotation.UpdatedAt)
	resp.Checksum = annotation.Hash
	resp.TotalSize = int64(annotation.Size)

	return &resp, nil

}

type annotationQuerier interface {
	QueryAnnotations() *ent.AnnotationQuery
}

func (flow *flow) SetAnnotation(ctx context.Context, annotationc *ent.AnnotationClient, q annotationQuerier, key string, data []byte) (*ent.Annotation, bool, error) {

	hash, err := computeHash(data)
	if err != nil {
		flow.sugar.Error(err)
	}

	var annotation *ent.Annotation
	var newAnnotation bool

	annotation, err = q.QueryAnnotations().Where(entnote.NameEQ(key)).Only(ctx)

	if err != nil {

		if !IsNotFound(err) {
			return nil, false, err
		}

		query := annotationc.Create().SetSize(len(data)).SetHash(hash).SetData(data).SetName(key)

		switch q.(type) {
		case *ent.Namespace:
			query = query.SetNamespace(q.(*ent.Namespace))
		case *ent.Workflow:
			query = query.SetWorkflow(q.(*ent.Workflow))
		case *ent.Instance:
			query = query.SetInstance(q.(*ent.Instance))
		default:
			panic(errors.New("bad querier"))
		}

		annotation, err = query.Save(ctx)
		if err != nil {
			return nil, false, err
		}

		newAnnotation = true

	} else {

		query := annotation.Update().SetSize(len(data)).SetHash(hash).SetData(data)

		annotation, err = query.Save(ctx)
		if err != nil {
			return nil, false, err
		}

	}

	return annotation, newAnnotation, err

}

func (flow *flow) SetNamespaceAnnotationParcels(srv grpc.Flow_SetNamespaceAnnotationParcelsServer) error {

	flow.sugar.Debugf("Handling gRPC request: %s", this())

	ctx := srv.Context()

	req, err := srv.Recv()
	if err != nil {
		return err
	}

	namespace := req.GetNamespace()
	key := req.GetKey()

	totalSize := int(req.GetTotalSize())

	buf := new(bytes.Buffer)

	for {

		_, err = io.Copy(buf, bytes.NewReader(req.Data))
		if err != nil {
			return err
		}

		if req.TotalSize <= 0 {
			if buf.Len() >= totalSize {
				break
			}
		}

		req, err = srv.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		if req.TotalSize <= 0 {
			if buf.Len() >= totalSize {
				break
			}
		} else {
			if req == nil {
				break
			}
		}

		if int(req.GetTotalSize()) != totalSize {
			return errors.New("totalSize changed mid stream")
		}

	}

	if buf.Len() > totalSize {
		return errors.New("received more data than expected")
	}

	tx, err := flow.db.Tx(ctx)
	if err != nil {
		return err
	}
	defer rollback(tx)

	nsc := tx.Namespace
	annotationc := tx.Annotation

	ns, err := flow.getNamespace(ctx, nsc, namespace)
	if err != nil {
		return err
	}

	var annotation *ent.Annotation

	var newAnnotation bool
	annotation, newAnnotation, err = flow.SetAnnotation(ctx, annotationc, ns, key, buf.Bytes())
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	if newAnnotation {
		flow.logToNamespace(ctx, time.Now(), ns, "Created namespace annotation '%s'.", key)
	} else {
		flow.logToNamespace(ctx, time.Now(), ns, "Updated namespace annotation '%s'.", key)
	}

	flow.pubsub.NotifyNamespaceAnnotations(ns)

	var resp grpc.SetNamespaceAnnotationResponse

	resp.Namespace = ns.Name
	resp.Key = key
	resp.CreatedAt = timestamppb.New(annotation.CreatedAt)
	resp.UpdatedAt = timestamppb.New(annotation.UpdatedAt)
	resp.Checksum = annotation.Hash
	resp.TotalSize = int64(annotation.Size)

	err = srv.SendAndClose(&resp)
	if err != nil {
		return err
	}

	return nil

}

func (flow *flow) DeleteNamespaceAnnotation(ctx context.Context, req *grpc.DeleteNamespaceAnnotationRequest) (*emptypb.Empty, error) {

	flow.sugar.Debugf("Handling gRPC request: %s", this())

	tx, err := flow.db.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer rollback(tx)

	nsc := tx.Namespace

	d, err := flow.traverseToNamespaceAnnotation(ctx, nsc, req.GetNamespace(), req.GetKey())
	if err != nil {
		return nil, err
	}

	annotationc := tx.Annotation

	err = annotationc.DeleteOne(d.annotation).Exec(ctx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	flow.logToNamespace(ctx, time.Now(), d.ns(), "Deleted namespace annotation '%s'.", d.annotation.Name)
	flow.pubsub.NotifyNamespaceAnnotations(d.ns())

	var resp emptypb.Empty

	return &resp, nil

}

func (flow *flow) RenameNamespaceAnnotation(ctx context.Context, req *grpc.RenameNamespaceAnnotationRequest) (*grpc.RenameNamespaceAnnotationResponse, error) {

	flow.sugar.Debugf("Handling gRPC request: %s", this())

	tx, err := flow.db.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer rollback(tx)

	nsc := tx.Namespace
	d, err := flow.traverseToNamespaceAnnotation(ctx, nsc, req.GetNamespace(), req.GetOld())
	if err != nil {
		return nil, err
	}

	annotation, err := d.annotation.Update().SetName(req.GetNew()).Save(ctx)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	flow.logToNamespace(ctx, time.Now(), d.ns(), "Renamed namespace annotation from '%s' to '%s'.", req.GetOld(), req.GetNew())
	flow.pubsub.NotifyNamespaceAnnotations(d.ns())

	var resp grpc.RenameNamespaceAnnotationResponse

	resp.Checksum = d.annotation.Hash
	resp.CreatedAt = timestamppb.New(d.annotation.CreatedAt)
	resp.Key = annotation.Name
	resp.Namespace = d.ns().Name
	resp.TotalSize = int64(d.annotation.Size)
	resp.UpdatedAt = timestamppb.New(d.annotation.UpdatedAt)

	return &resp, nil

}
