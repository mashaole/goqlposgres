package domain

import (
	"context"
	"errors"
	"fmt"
	"goqlposgress/middleware"
	"goqlposgress/models"
)

// CreateMeetup is the resolver for the createMeetup field.
func (d *Domain) CreateMeetup(ctx context.Context, input models.NewMeetup) (*models.Meetup, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, ErrUnauthenticated
	}
	if len(input.Name) < 3 {
		return nil, errors.New("name not long enough")
	}

	if len(input.Description) < 3 {
		return nil, errors.New("description not long enough")
	}
	meetup := &models.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      currentUser.ID,
	}
	return d.MeetupsRepo.CreateMeetup(meetup)
}

// UpdateMeetup is the resolver for the updateMeetup field.
func (d *Domain) UpdateMeetup(ctx context.Context, id string, input models.UpdateMeetup) (*models.Meetup, error) {
	meetup, err := d.MeetupsRepo.GetByID(id)
	if err != nil || meetup == nil {
		return nil, errors.New("meetup not exist")
	}
	didUpdate := false
	if input.Name != nil {
		if len(*input.Name) < 3 {
			return nil, errors.New("name is not long enough")
		}
		meetup.Name = *input.Name
		didUpdate = true
	}
	if input.Description != nil {
		if len(*input.Description) < 3 {
			return nil, errors.New("Description is not long enough")
		}
		meetup.Description = *input.Description
		didUpdate = true
	}
	if !didUpdate {
		return nil, errors.New("no update done")
	}
	meetup, err = d.MeetupsRepo.Update(meetup)
	if err != nil {
		return nil, fmt.Errorf("error whiel updating meetup: %v", err)
	}
	return meetup, nil
}

// DeleteMeetup is the resolver for the deleteMeetup field.
func (d *Domain) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	meetup, err := d.MeetupsRepo.GetByID(id)
	if err != nil || meetup == nil {
		return false, errors.New("meetup not exist")
	}
	err = d.MeetupsRepo.Delete(meetup)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup: %v", err)
	}
	return true, nil
}
