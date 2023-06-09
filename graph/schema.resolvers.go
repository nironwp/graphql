package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.30

import (
	"context"

	"github.com/google/uuid"
	"github.com/nironwp/graphql/graph/model"
	"github.com/nironwp/graphql/internal/helpers"
)

// Courses is the resolver for the courses field.
func (r *categoryResolver) Courses(ctx context.Context, obj *model.Category) ([]*model.Course, error) {
	courses, err := r.CourseDB.FindByCategoryId(obj.ID)

	if err != nil {
		return nil, err
	}

	var results []*model.Course

	for _, course := range courses {
		results = append(results, &model.Course{
			ID:          course.ID,
			Name:        course.Name,
			Description: &course.Description,
		})
	}

	return results, nil
}

// Category is the resolver for the category field.
func (r *courseResolver) Category(ctx context.Context, obj *model.Course) (*model.Category, error) {
	category, err := r.CategoryDB.FindByCourseID(obj.ID)

	if err != nil {
		return nil, err
	}

	return &model.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: &category.Description,
	}, nil
}

// AddCategory is the resolver for the addCategory field.
func (r *mutationResolver) AddCategory(ctx context.Context, category model.NewCategory) (*model.Category, error) {
	description := helpers.PointerToString(category.Description)
	id := uuid.New().String()
	cc, err := r.CategoryDB.Create(category.Name, description)
	if err != nil {
		return nil, err
	}

	return &model.Category{
		ID:          id,
		Name:        cc.Name,
		Description: &cc.Description,
	}, nil
}

// AddCourse is the resolver for the addCourse field.
func (r *mutationResolver) AddCourse(ctx context.Context, course model.NewCourse) (*model.Course, error) {
	cc, err := r.CourseDB.Create(course.Name, helpers.PointerToString(course.Description), course.CategoryID)

	if err != nil {
		return nil, err
	}
	return &model.Course{
		ID:          cc.ID,
		Name:        cc.Name,
		Description: &cc.Description,
	}, nil
}

// Categories is the resolver for the categories field.
func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	categories, err := r.CategoryDB.FindAll()

	if err != nil {
		return nil, err
	}

	var results []*model.Category

	for _, category := range categories {
		results = append(results, &model.Category{
			ID:          category.ID,
			Name:        category.Name,
			Description: &category.Description,
		})
	}

	return results, nil
}

// Courses is the resolver for the courses field.
func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	courses, err := r.CourseDB.FindAll()

	if err != nil {
		return nil, err
	}

	var results []*model.Course

	for _, course := range courses {
		results = append(results, &model.Course{
			ID:          course.ID,
			Name:        course.Name,
			Description: &course.Description,
		})
	}

	return results, nil
}

// Category returns CategoryResolver implementation.
func (r *Resolver) Category() CategoryResolver { return &categoryResolver{r} }

// Course returns CourseResolver implementation.
func (r *Resolver) Course() CourseResolver { return &courseResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type categoryResolver struct{ *Resolver }
type courseResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
